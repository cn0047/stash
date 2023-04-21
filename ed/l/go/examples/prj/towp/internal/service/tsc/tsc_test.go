package tsc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal/testutils"
)

const retailerID, mfcID = "fake-retailer", "fake-mfc"

func prepareTSC(t *testing.T) *Service {
	t.Helper()

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unable to load config for testing TSC service, error: %v", err)
	}
	logger := foundation.NewLogger()
	httpClient := &http.Client{}

	return New(cfg, logger, httpClient)
}

func TestTSCService(t *testing.T) {
	serviceCatalog := prepareTSC(t)

	t.Run("when tsc service called and request failed", func(t *testing.T) {
		srv := makeServer("request failed", http.StatusBadRequest)
		defer srv.Close()

		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s"

		ctx := testutils.BuildCtx()
		_, err := serviceCatalog.InStorePickingEnabled(ctx, retailerID, mfcID)

		assert.NotEmpty(t, err)
	})

	t.Run("when tsc service called and success response received and config turned on", func(t *testing.T) {
		f, _ := os.ReadFile("testdata/tsc_isps_enabled_config.json")
		srv := makeServer(string(f), http.StatusOK)
		defer srv.Close()

		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"

		ctx := testutils.BuildCtx()
		ispsEnabled, _ := serviceCatalog.InStorePickingEnabled(ctx, retailerID, mfcID)

		assert.Equal(t, true, ispsEnabled)
	})

	t.Run("when tsc service called and success response received and config turned off", func(t *testing.T) {
		f, _ := os.ReadFile("testdata/tsc_isps_disabled_config.json")

		srv := makeServer(string(f), http.StatusOK)
		defer srv.Close()

		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"

		ctx := testutils.BuildCtx()
		ispsEnabled, _ := serviceCatalog.InStorePickingEnabled(ctx, retailerID, mfcID)

		assert.Equal(t, false, ispsEnabled)
	})

	t.Run("when tsc service called and config not found", func(t *testing.T) {
		srv := makeServer("[]", http.StatusOK)
		defer srv.Close()

		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"

		ctx := testutils.BuildCtx()
		_, err := serviceCatalog.InStorePickingEnabled(ctx, retailerID, mfcID)

		assert.ErrorContains(t, err, "ISPS_ENABLED config is not found in service-catalog")
	})

	t.Run("when tsc service called and unable to serialize response", func(t *testing.T) {
		srv := makeServer(`"[]"`, http.StatusOK)
		defer srv.Close()

		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"

		ctx := testutils.BuildCtx()
		_, err := serviceCatalog.InStorePickingEnabled(ctx, retailerID, mfcID)

		assert.ErrorContains(t, err, "unable to unmarshal service-catalog config")
	})
}

func TestTSCServiceGetLocationInfo(t *testing.T) {
	serviceCatalog := prepareTSC(t)
	ctx := testutils.BuildCtx()

	t.Run("invalid host", func(t *testing.T) {
		srv := makeServer(``, http.StatusOK)
		serviceCatalog.cfg.TSCURLTemplate = "http://not-a-real-host/%s/%s/"
		defer srv.Close()

		_, err := serviceCatalog.GetLocationInfo(ctx, retailerID, mfcID)

		assert.ErrorContains(t, err, "failed to perform request")
	})

	t.Run("server error", func(t *testing.T) {
		srv := makeServer(``, http.StatusInternalServerError)
		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"
		defer srv.Close()

		_, err := serviceCatalog.GetLocationInfo(ctx, retailerID, mfcID)

		assert.ErrorContains(t, err, "got unexpected http status code")
	})

	t.Run("invalid payload", func(t *testing.T) {
		srv := makeServer(`error001`, http.StatusOK)
		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"
		defer srv.Close()

		_, err := serviceCatalog.GetLocationInfo(ctx, retailerID, mfcID)

		assert.ErrorContains(t, err, "failed to unmarshal response")
	})

	t.Run("timezone not found", func(t *testing.T) {
		srv := makeServer(`[]`, http.StatusOK)
		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"
		defer srv.Close()

		_, err := serviceCatalog.GetLocationInfo(ctx, retailerID, mfcID)

		assert.ErrorContains(t, err, "location info not found")
	})

	t.Run("simple success case", func(t *testing.T) {
		expectedTimezone := "Europe/Kiev"
		expectedLocationInfo := fmt.Sprintf(`[{"mfc-ref-code": "%s", "timezone": "%s"}]`, mfcID, expectedTimezone)
		srv := makeServer(expectedLocationInfo, http.StatusOK)
		serviceCatalog.cfg.TSCURLTemplate = srv.URL + "/%s/%s/"
		defer srv.Close()

		actual, err := serviceCatalog.GetLocationInfo(ctx, retailerID, mfcID)

		assert.Nil(t, err)
		assert.Equal(t, expectedTimezone, actual.Timezone)
	})
}

func makeServer(responseBody string, responseStatusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseStatusCode)
		_, _ = w.Write([]byte(responseBody))
	}))
}
