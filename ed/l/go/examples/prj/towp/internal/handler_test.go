package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal/business/validator"
	"github.com/to-com/wp/internal/common"
	"github.com/to-com/wp/internal/errors"
	"github.com/to-com/wp/internal/service/auth"
	"github.com/to-com/wp/internal/testutils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/internal/dto"
	"github.com/to-com/wp/internal/mocks"
)

var locationInfo = mocks.GetMockLocationInfo()
var wpBasePath = fmt.Sprintf("/v1/retailers/%s/mfcs/%s/wp",
	locationInfo.RetailerID, locationInfo.MfcID)
var wpBasePath2 = "/v1/retailers/fake-retailer-2/mfcs/fake-mfc-2/wp"
var getTriggersBasePath = "/v1/triggers"
var generateTriggersPath = "/v1/triggers:generate"
var fireTriggersPath = "/v1/triggers:fire"

type wpSuite struct {
	suite.Suite

	testApp *TestApplication
}

func TestHandlerwpSuite(t *testing.T) {
	suite.Run(t, new(wpSuite))
}

func (s *wpSuite) SetupTest() {
	s.testApp = GetTestApp(s.T())
}

func GetTestApp(t *testing.T) *TestApplication {
	testApp := &TestApplication{}

	cfg, err := config.Load()
	if err != nil {
		t.Errorf("unable to load config for tests, error: %v", err)
		t.FailNow()
	}

	logger := foundation.NewLogger()

	testApp.ctrl = gomock.NewController(t)
	testApp.mockBusiness = mocks.NewMockBusiness(testApp.ctrl)
	testApp.mockAuth = mocks.NewMockAuthentication(testApp.ctrl)

	httpHandler := NewHTTPHandler(logger, testApp.mockBusiness, testApp.mockAuth)

	testApp.App = &Application{
		cfg:     cfg,
		logger:  logger,
		handler: httpHandler,
	}

	return testApp
}

func (s *wpSuite) TearDownTest() {
	s.testApp.Shutdown()
}

func (s *wpSuite) TestGetPostwpErrorResponsesFromAuth() {
	headers := mocks.GetHeaders()

	s.T().Run("Create wave plan with no response from auth service", func(t *testing.T) {
		s.testApp.mockAuth.
			EXPECT().
			CheckUser(gomock.Any(), gomock.Any()).
			Return("", fmt.Errorf("%w: %s", errors.ErrHTTPUnauthorized, auth.AuthErr))

		// No calls to business.Getwp are expected
		s.testApp.mockBusiness.
			EXPECT().
			Createwp(gomock.Any(), gomock.Any()).
			Times(0)

		response := httptest.NewRecorder()
		wpRequestString := []byte(testutils.ReadFileAsString(t, "testdata/wpNoSchedules.json"))
		request := httptest.NewRequest(http.MethodPost, wpBasePath,
			bytes.NewBuffer(wpRequestString))

		for header, value := range headers {
			request.Header.Add(header, value)
		}
		s.testApp.App.routes().ServeHTTP(response, request)

		require.Equal(t, http.StatusUnauthorized, response.Code)

		expectedResponse := errors.ResponseWithErrors{
			Errors: []errors.Error{
				{
					Msg: fmt.Sprintf("%v: %s", errors.ErrHTTPUnauthorized, auth.AuthErr),
				},
			},
		}
		var resp errors.ResponseWithErrors
		err := json.Unmarshal(response.Body.Bytes(), &resp)
		require.Nil(t, err)
		assert.Equal(t, expectedResponse, resp)
	})
}

func (s *wpSuite) TestGetwpWhenBusinessReturnsError() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	s.testApp.mockBusiness.
		EXPECT().
		Getwp(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID), gomock.Eq(locationInfo.MfcID)).
		Return(dto.wpResponse{}, fmt.Errorf("fail to get wave plan")).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, wpBasePath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	s.Equal(http.StatusInternalServerError, response.Code)
	s.Contains(response.Body.String(), "fail to get wave plan")
}

func (s *wpSuite) TestGetwpWhenNoPlanExist() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	s.testApp.mockBusiness.
		EXPECT().
		Getwp(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID), gomock.Eq(locationInfo.MfcID)).
		Return(dto.wpResponse{}, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, wpBasePath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	s.Equal(http.StatusNoContent, response.Code)
	s.Empty(response.Body.String())
}

func (s *wpSuite) TestGetwpNoSchedules() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	// Mock business.Getwp response
	wpResponseString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpNoSchedulesResponse.json"))
	var expectedResponse dto.wpResponse
	err := json.Unmarshal(wpResponseString, &expectedResponse)
	if err != nil {
		s.T().Fatalf("fail lo unmarshal expected wave plan response, err: %v", err)
	}
	s.testApp.mockBusiness.
		EXPECT().
		Getwp(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID), gomock.Eq(locationInfo.MfcID)).
		Return(expectedResponse, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, wpBasePath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	var actualResponse dto.wpResponse
	err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
	if err != nil {
		s.T().Fatalf("fail lo unmarshal actual wave plan response, err: %v", err)
	}
	s.Equal(http.StatusOK, response.Code)
	s.Equal(expectedResponse, actualResponse)
}

func (s *wpSuite) TestGetwpWithSchedules() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	// Mock business.Getwp response
	wpResponseString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedulesResponse.json"))
	var expectedResponse dto.wpResponse
	err := json.Unmarshal(wpResponseString, &expectedResponse)
	if err != nil {
		s.T().Fatalf("fail lo unmarshal expected wave plan response, err: %v", err)
	}
	s.testApp.mockBusiness.
		EXPECT().
		Getwp(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID), gomock.Eq(locationInfo.MfcID)).
		Return(expectedResponse, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, wpBasePath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	var actualResponse dto.wpResponse
	err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
	if err != nil {
		s.T().Fatalf("fail lo unmarshal actual wave plan response, err: %v", err)
	}
	s.Equal(http.StatusOK, response.Code)
	s.Equal(expectedResponse, actualResponse)
}

func (s *wpSuite) TestPostwpWithNoSchedules() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)
	ctx = context.WithValue(ctx, common.CtxKeyToken, headers["X-Token"])

	s.testApp.mockAuth.
		EXPECT().
		CheckUser(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID)).
		Return(mocks.GetUserPermissionMock()["user_id"], nil).
		Times(1)

	wpRequestString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpNoSchedules.json"))

	var wpRequest dto.wpRequest
	err := json.Unmarshal(wpRequestString, &wpRequest)
	require.Nil(s.T(), err, "fail to unmarshal test request, err: %v", err)
	wpRequest.RetailerID = locationInfo.RetailerID
	wpRequest.MfcID = locationInfo.MfcID
	wpRequest.UserID = mocks.GetUserPermissionMock()["user_id"]

	wpResponseString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpNoSchedulesResponse.json"))
	var expectedResponse dto.wpResponse
	err = json.Unmarshal(wpResponseString, &expectedResponse)
	require.Nil(s.T(), err, "fail to unmarshal expected response, err: %v", err)

	// Mock business.Createwp response
	s.testApp.mockBusiness.
		EXPECT().
		Createwp(testutils.EqContext(ctx), gomock.Eq(wpRequest)).
		Return(expectedResponse, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer(wpRequestString))

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	require.Equal(s.T(), http.StatusCreated, response.Code)

	var actualResponse dto.wpResponse
	err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
	require.Nil(s.T(), err, "fail lo unmarshal actual wave plan response, err: %v", err)

	s.Equal(http.StatusCreated, response.Code)
	s.Equal(expectedResponse, actualResponse)
}

func (s *wpSuite) TestPostwpWithSchedules() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)
	ctx = context.WithValue(ctx, common.CtxKeyToken, "token")

	s.testApp.mockAuth.
		EXPECT().
		CheckUser(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID)).
		Return(mocks.GetUserPermissionMock()["user_id"], nil).
		Times(1)

	wpRequestString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedules.json"))

	var wpRequest dto.wpRequest
	err := json.Unmarshal(wpRequestString, &wpRequest)
	require.Nil(s.T(), err, "fail to unmarshal test request, err: %v", err)
	wpRequest.RetailerID = locationInfo.RetailerID
	wpRequest.MfcID = locationInfo.MfcID
	wpRequest.UserID = mocks.GetUserPermissionMock()["user_id"]

	wpResponseString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedulesResponse.json"))
	var expectedResponse dto.wpResponse
	err = json.Unmarshal(wpResponseString, &expectedResponse)
	require.Nil(s.T(), err, "fail to unmarshal expected response, err: %v", err)

	// Mock business.Createwp response
	s.testApp.mockBusiness.
		EXPECT().
		Createwp(testutils.EqContext(ctx), wpRequest).
		Return(expectedResponse, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer(wpRequestString))

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	require.Equal(s.T(), http.StatusCreated, response.Code)

	var actualResponse dto.wpResponse
	err = json.Unmarshal(response.Body.Bytes(), &actualResponse)
	require.Nil(s.T(), err, "fail lo unmarshal actual wave plan response, err: %v", err)

	s.Equal(expectedResponse, actualResponse)
}

func (s *wpSuite) TestPostwpWithInvalidRequestBody() {
	headers := mocks.GetHeaders()

	s.testApp.mockAuth.
		EXPECT().
		CheckUser(gomock.Any(), gomock.Any()).
		Return(mocks.GetUserPermissionMock()["user_id"], nil).
		Times(1)

	// No calls to business.Createwp are expected
	s.testApp.mockBusiness.
		EXPECT().
		Createwp(gomock.Any(), gomock.Any()).
		Times(0)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer([]byte("invalid request body")))

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	s.Equal(http.StatusBadRequest, response.Code)
}

func (s *wpSuite) TestPostwpWhenBusinessReturnsError() {
	headers := mocks.GetHeaders()

	s.testApp.mockAuth.
		EXPECT().
		CheckUser(gomock.Any(), gomock.Any()).
		Return(mocks.GetUserPermissionMock()["user_id"], nil).
		Times(1)
	// Mock business.Createwp response
	s.testApp.mockBusiness.
		EXPECT().
		Createwp(gomock.Any(), gomock.Any()).
		Return(dto.wpResponse{}, fmt.Errorf("business error")).
		Times(1)

	response := httptest.NewRecorder()
	wpRequestString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedules.json"))
	request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer(wpRequestString))

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	require.Equal(s.T(), http.StatusInternalServerError, response.Code)
	s.Equal(`{"errors":[{"message":"business error"}]}`, response.Body.String())
}

func (s *wpSuite) TestPostwpWhenBusinessReturnsValidationError() {
	headers := mocks.GetHeaders()

	s.testApp.mockAuth.
		EXPECT().
		CheckUser(gomock.Any(), gomock.Any()).
		Return(mocks.GetUserPermissionMock()["user_id"], nil).
		Times(1)

	err := validator.ValidationError{
		WavesErrors: []*validator.WaveError{
			{
				Err: fmt.Errorf("business error"),
			},
		},
	}
	// Mock business.Createwp response
	s.testApp.mockBusiness.
		EXPECT().
		Createwp(gomock.Any(), gomock.Any()).
		Return(dto.wpResponse{}, &err).
		Times(1)

	response := httptest.NewRecorder()
	wpRequestString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedules.json"))
	request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer(wpRequestString))

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	require.Equal(s.T(), http.StatusBadRequest, response.Code)
	s.Equal(`{"errors":[{"message":"business error","validation_error":true}]}`, response.Body.String())
}

func (s *wpSuite) TestRequiredHeaders() {
	headers := mocks.GetHeaders()

	var xEnvType = "X-Env-Type"
	s.T().Run("X-Env-Type header is not provided. Get wave plan", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, wpBasePath2, nil)

		for header, value := range headers {
			if header == xEnvType {
				continue
			}
			request.Header.Add(header, value)
		}

		s.testApp.App.routes().ServeHTTP(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), ErrEmptyEnv.Error())
	})

	s.T().Run("X-Env-Type header is not provided. Post wave plan", func(t *testing.T) {
		wpRequestString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedules.json"))
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer(wpRequestString))

		for header, value := range headers {
			if header == xEnvType {
				continue
			}
			request.Header.Add(header, value)
		}

		s.testApp.App.routes().ServeHTTP(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), ErrEmptyEnv.Error())
	})

	s.T().Run("X-Token header is not provided. Post wave plan", func(t *testing.T) {
		wpRequestString := []byte(testutils.ReadFileAsString(s.T(), "testdata/wpWithSchedules.json"))
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, wpBasePath, bytes.NewBuffer(wpRequestString))

		for header, value := range headers {
			if header == "X-Token" {
				continue
			}
			request.Header.Add(header, value)
		}

		s.testApp.App.routes().ServeHTTP(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), ErrEmptyToken.Error())
	})

	s.T().Run("X-Token header is not provided. Generate triggers", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, generateTriggersPath, nil)

		for header, value := range headers {
			if header == xEnvType {
				continue
			}
			request.Header.Add(header, value)
		}

		s.testApp.App.routes().ServeHTTP(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), ErrEmptyEnv.Error())
	})

	s.T().Run("X-Token header is not provided. Fire triggers", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, fireTriggersPath, nil)

		for header, value := range headers {
			if header == xEnvType {
				continue
			}
			request.Header.Add(header, value)
		}

		s.testApp.App.routes().ServeHTTP(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		assert.Contains(t, response.Body.String(), ErrEmptyEnv.Error())
	})
}

func (s *wpSuite) TestGetTriggersEmptywp() {
	headers := mocks.GetHeaders()

	emptyResponse := `{"triggers":[]}`

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	var triggersResponse dto.GetTriggersResponse
	err := json.Unmarshal([]byte(emptyResponse), &triggersResponse)
	s.Empty(err)

	s.testApp.mockBusiness.
		EXPECT().
		GetTriggers(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID), gomock.Eq(locationInfo.MfcID)).
		Return(triggersResponse, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, getTriggersBasePath, nil)

	query := request.URL.Query()
	query.Add("retailerId", locationInfo.RetailerID)
	query.Add("mfcId", locationInfo.MfcID)
	request.URL.RawQuery = query.Encode()

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	s.Equal(http.StatusOK, response.Code)
	s.Equal(emptyResponse, response.Body.String())
}

func (s *wpSuite) TestGetTriggers() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	var triggersResponse dto.GetTriggersResponse
	err := json.Unmarshal([]byte(testutils.ReadFileAsString(s.T(), "testdata/triggerItems.json")), &triggersResponse)
	s.Empty(err)

	s.testApp.mockBusiness.
		EXPECT().
		GetTriggers(testutils.EqContext(ctx), gomock.Eq(locationInfo.RetailerID), gomock.Eq(locationInfo.MfcID)).
		Return(triggersResponse, nil).
		Times(1)
	// get triggers
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, getTriggersBasePath, nil)

	query := request.URL.Query()
	query.Add("retailerId", locationInfo.RetailerID)
	query.Add("mfcId", locationInfo.MfcID)
	request.URL.RawQuery = query.Encode()

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	var resp dto.GetTriggersResponse
	err = json.Unmarshal(response.Body.Bytes(), &resp)
	s.Empty(err)

	s.Equal(http.StatusOK, response.Code)
	s.Equal(len(resp.Triggers), 3)

	s.Equal(triggersResponse, resp)
}

func (s *wpSuite) TestGetTriggersReturnsError() {
	headers := mocks.GetHeaders()

	errorMessage := "get triggers error"

	s.testApp.mockBusiness.
		EXPECT().
		GetTriggers(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(dto.GetTriggersResponse{}, fmt.Errorf(errorMessage)).
		Times(1)

	// get triggers
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, getTriggersBasePath, nil)

	q := request.URL.Query()
	q.Add("retailerId", locationInfo.RetailerID)
	q.Add("mfcId", locationInfo.MfcID)
	request.URL.RawQuery = q.Encode()

	for header, value := range headers {
		request.Header.Add(header, value)
	}
	s.testApp.App.routes().ServeHTTP(response, request)

	var resp errors.ResponseWithErrors
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	s.Empty(err)

	expectedResponse := errors.ResponseWithErrors{
		Errors: []errors.Error{
			{
				Msg: errorMessage,
			},
		},
	}

	s.Equal(http.StatusInternalServerError, response.Code)
	s.Equal(resp, expectedResponse)
}

func (s *wpSuite) TestGenerateTriggers() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	generatedTriggers := dto.GenerateTriggersResponse{
		GeneratedTriggers: 2,
	}
	s.testApp.mockBusiness.
		EXPECT().
		GenerateTriggers(testutils.EqContext(ctx)).
		Return(generatedTriggers, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, generateTriggersPath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	var actual dto.GenerateTriggersResponse
	err := json.Unmarshal(response.Body.Bytes(), &actual)
	s.Empty(err)
	s.Equal(http.StatusOK, response.Code)
	s.Equal(generatedTriggers, actual)
}

func (s *wpSuite) TestGenerateTriggersReturnsError() {
	headers := mocks.GetHeaders()

	errorMessage := "generate triggers error"

	s.testApp.mockBusiness.
		EXPECT().
		GenerateTriggers(gomock.Any()).
		Return(dto.GenerateTriggersResponse{}, fmt.Errorf(errorMessage)).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, generateTriggersPath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	expectedResponse := errors.ResponseWithErrors{
		Errors: []errors.Error{
			{
				Msg: errorMessage,
			},
		},
	}
	var resp errors.ResponseWithErrors
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	s.Empty(err)
	s.Equal(http.StatusInternalServerError, response.Code)
	s.Equal(expectedResponse, resp)
}

func (s *wpSuite) TestFireTriggers() {
	headers := mocks.GetHeaders()

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, locationInfo.EnvType)

	firedTriggers := dto.FireTriggersResponse{
		Triggers: []dto.FireTriggerResponse{
			{
				RetailerID:     "fake-retailer",
				MfcID:          "fake-mfc",
				ScheduleID:     uuid.NewString(),
				ScheduleType:   "Prelim",
				CutoffDateTime: time.Now(),
				TriggerAt:      time.Now().Add(+2 * time.Hour),
				CreatedAt:      time.Now(),
				FiredAt:        nil,
			},
		},
	}
	firedTriggersPayload, err := json.Marshal(firedTriggers)
	s.Empty(err)

	s.testApp.mockBusiness.
		EXPECT().
		FireTriggers(testutils.EqContext(ctx)).
		Return(firedTriggers, nil).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, fireTriggersPath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	//actual := dto.FireTriggersResponse{Triggers: make([]dto.FireTriggerResponse, 1)}
	actual := dto.FireTriggersResponse{}
	err = json.Unmarshal(response.Body.Bytes(), &actual)
	s.Empty(err)
	s.Equal(http.StatusOK, response.Code)
	s.Equal(firedTriggersPayload, response.Body.Bytes())
}

func (s *wpSuite) TestFireTriggersReturnsError() {
	headers := mocks.GetHeaders()

	errorMessage := "fire triggers error"

	s.testApp.mockBusiness.
		EXPECT().
		FireTriggers(gomock.Any()).
		Return(dto.FireTriggersResponse{}, fmt.Errorf(errorMessage)).
		Times(1)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, fireTriggersPath, nil)

	for header, value := range headers {
		request.Header.Add(header, value)
	}

	s.testApp.App.routes().ServeHTTP(response, request)

	expectedResponse := errors.ResponseWithErrors{
		Errors: []errors.Error{
			{
				Msg: errorMessage,
			},
		},
	}
	var resp errors.ResponseWithErrors
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	s.Empty(err)
	s.Equal(http.StatusInternalServerError, response.Code)
	s.Equal(expectedResponse, resp)
}
