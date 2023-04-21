package internal

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal/business/validator"
	"github.com/to-com/wp/internal/dto"
	"github.com/to-com/wp/internal/testutils"
	"log"
	"os"
	"testing"
)

const masterBranch = "master"

func TestOutboundBackendPact(t *testing.T) {
	t.Setenv("CONSUMER_NAME", "outboundbff-wp")

	loadEnvVars()
	app := GetTestApp(t)
	go startApp(app)

	verifier := provider.NewVerifier()

	var wpNoSchedulesResponse dto.wpResponse
	respString := []byte(testutils.ReadFileAsString(t, "testdata/wpNoSchedulesResponse.json"))
	_ = json.Unmarshal(respString, &wpNoSchedulesResponse)

	var wpWithSchedulesResponse dto.wpResponse
	respString = []byte(testutils.ReadFileAsString(t, "testdata/wpWithSchedulesResponse.json"))
	_ = json.Unmarshal(respString, &wpWithSchedulesResponse)

	stateMappings := models.StateHandlers{
		"wave plan with no schedulers exists for retailer fake-retailer mfc fake-mfc": func(setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "Getwp" in business service.
				app.mockBusiness.
					EXPECT().
					Getwp(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(wpNoSchedulesResponse, nil)
			}

			return nil, nil
		},

		"wave plan with schedulers exists for retailer fake-retailer mfc fake-mfc": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "Getwp" in business service.
				app.mockBusiness.
					EXPECT().
					Getwp(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(wpWithSchedulesResponse, nil)
			}

			return nil, nil
		},

		"wave plan does not exist for retailer fake-retailer mfc fake-mfc": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "Getwp" in business service.
				app.mockBusiness.
					EXPECT().
					Getwp(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(dto.wpResponse{}, nil)
			}

			return nil, nil
		},

		"retailer fake-retailer, mfc fake-mfc exists, it supports isps": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "CheckUser" in auth service.
				app.mockAuth.
					EXPECT().
					CheckUser(gomock.Any(), gomock.Any()).
					Return("test-user", nil)

				// Mock "Createwp" in business service.
				var wp dto.wpRequest
				app.mockBusiness.
					EXPECT().
					Createwp(gomock.Any(), gomock.AssignableToTypeOf(wp)).
					Return(wpWithSchedulesResponse, nil)
			}

			return nil, nil
		},

		"retailer fake-retailer, mfc fake-mfc exists, it does not support isps": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "CheckUser" in auth service.
				app.mockAuth.
					EXPECT().
					CheckUser(gomock.Any(), gomock.Any()).
					Return("test-user", nil)
				// Mock "Createwp" in business service.
				var wp dto.wpRequest
				app.mockBusiness.
					EXPECT().
					Createwp(gomock.Any(), gomock.AssignableToTypeOf(wp)).
					Return(wpNoSchedulesResponse, nil)
			}

			return nil, nil
		},

		"simple error response case": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "CheckUser" in auth service.
				app.mockAuth.
					EXPECT().
					CheckUser(gomock.Any(), gomock.Any()).
					Return("test-user", nil)

				err := validator.ValidationError{
					WavesErrors: []*validator.WaveError{
						{
							Err: fmt.Errorf("configuration must have at least one wave"),
						},
					},
				}
				// Mock business.Createwp response
				app.mockBusiness.
					EXPECT().
					Createwp(gomock.Any(), gomock.Any()).
					Return(dto.wpResponse{}, &err)
			}

			return nil, nil
		},

		"complex error response case": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "CheckUser" in auth service.
				app.mockAuth.
					EXPECT().
					CheckUser(gomock.Any(), gomock.Any()).
					Return("test-user", nil)

				err := validator.ValidationError{
					WavesErrors: []*validator.WaveError{
						{
							Err:    fmt.Errorf("cannot save schedules since InStorePicking configuration turned off"),
							Cutoff: "14:00",
							ErrFields: []string{
								"prelim_picklist_schedule_time",
								"delta_picklist_schedule_time",
							},
						},
					},
				}
				// Mock business.Createwp response
				app.mockBusiness.
					EXPECT().
					Createwp(gomock.Any(), gomock.Any()).
					Return(dto.wpResponse{}, &err)
			}

			return nil, nil
		},

		"mixed errors response case": func(
			setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				// Mock "CheckUser" in auth service.
				app.mockAuth.
					EXPECT().
					CheckUser(gomock.Any(), gomock.Any()).
					Return("test-user", nil)

				err := validator.ValidationError{
					WavesErrors: []*validator.WaveError{
						{
							Err:    fmt.Errorf("cannot save schedules since InStorePicking configuration turned off"),
							Cutoff: "14:00",
							ErrFields: []string{
								"prelim_picklist_schedule_time",
								"delta_picklist_schedule_time",
							},
						},
						{
							Err: fmt.Errorf("wave configuration must cover a full 24-hour time period"),
						},
					},
				}
				// Mock business.Createwp response
				app.mockBusiness.
					EXPECT().
					Createwp(gomock.Any(), gomock.Any()).
					Return(dto.wpResponse{}, &err)
			}

			return nil, nil
		},
	}

	verifyRequest := provider.VerifyRequest{
		ProviderBaseURL:            "http://127.0.0.1:8080",
		Provider:                   "wp",
		ProviderBranch:             os.Getenv("PROVIDER_BRANCH"),
		ProviderVersion:            os.Getenv("PROVIDER_VERSION"),
		StateHandlers:              stateMappings,
		PublishVerificationResults: true,
	}

	consumerVersion := os.Getenv("CONSUMER_VERSION")
	switch consumerVersion {
	case "":
		loadEnvVars()
		t.Log("*** Provider verification tests are run locally. Tests results won't be published to pactflow")
		verifyRequest.ConsumerVersionSelectors = []provider.Selector{
			&provider.ConsumerVersionSelector{
				Consumer: os.Getenv("CONSUMER_NAME"),
				Branch:   masterBranch},
		}
		verifyRequest.ProviderBranch = "local"
		verifyRequest.ProviderVersion = "current"
		verifyRequest.PublishVerificationResults = false
	case masterBranch:
		t.Log("*** Provider verification tests are run in CI against `master` version of consumer. " +
			"Tests results will be published to pactflow")
		verifyRequest.ConsumerVersionSelectors = []provider.Selector{
			&provider.ConsumerVersionSelector{
				Consumer: os.Getenv("CONSUMER_NAME"),
				Branch:   masterBranch},
		}
	default:
		t.Log("*** Provider verification tests are run in CI against exact version of consumer. " +
			"Tests results will be published to pactflow")
		verifyRequest.PactURLs = []string{buildPactURL()}
	}
	verifyRequest.BrokerURL = os.Getenv("PACT_BROKER_URL")
	verifyRequest.BrokerToken = os.Getenv("PACT_BROKER_TOKEN")

	err := verifier.VerifyProvider(t, verifyRequest)

	assert.Empty(t, err)
	app.Shutdown()
}
func startApp(app *TestApplication) {
	logger := foundation.NewLogger()
	logger.Infow("startup", "status", "initializing application")

	if err := app.App.Serve(); err != nil {
		logger.Fatal(err)
	}
}

func loadEnvVars() {
	if os.Getenv("CONSUMER_VERSION") == "" {
		err := godotenv.Load("../.env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}

func buildPactURL() string {
	return fmt.Sprintf("%s/pacts/provider/wp/consumer/%s/version/%s",
		os.Getenv("PACT_BROKER_URL"),
		os.Getenv("CONSUMER_NAME"),
		os.Getenv("CONSUMER_VERSION"))
}
