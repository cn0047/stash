package mocks

import (
	"fmt"
	"gopkg.in/h2non/gock.v1"
)

type MockLocationInfo struct {
	RetailerID string
	MfcID      string
	EnvType    string
	Timezone   string
}

func GetMockLocationInfo() MockLocationInfo {
	return MockLocationInfo{
		RetailerID: "fake-retailer",
		MfcID:      "fake-mfc",
		EnvType:    "dev",
		Timezone:   "Asia/Dubai",
	}
}

func GetHeaders() map[string]string {
	locationInfo := GetMockLocationInfo()
	return map[string]string{
		"X-Env-Type":    locationInfo.EnvType,
		"X-Retailer-Id": locationInfo.RetailerID,
		"X-Token":       "token",
	}
}

func GetUserPermissionMock() map[string]string {
	return map[string]string{
		"user_id": "test-user",
	}
}

func GetIspsEnabledConfigValueMock(enabled bool) []map[string]any {
	return []map[string]any{
		{
			"categories":        []string{"isps", "pickerman"},
			"name":              "ISPS_ENABLED",
			"value":             enabled,
			"location-code-tom": GetMockLocationInfo().MfcID,
			"value-type":        "boolean",
		},
	}
}

func GetTscLocationInfoMock(timezone string) []map[string]any {
	return []map[string]any{
		{
			"mfc-ref-code": GetMockLocationInfo().MfcID,
			"timezone":     timezone,
		},
	}
}

func GetTscConfigErrorResponseMock() map[string]any {
	return map[string]any{
		"description": "These location codes you entered are not MFCs: " + GetMockLocationInfo().MfcID,
	}
}

func GetTscLocationsErrorResponseMock() map[string]any {
	return map[string]any{
		"description": "service is not available",
	}
}

var locationInfo = GetMockLocationInfo()

var tscPath = "^/api/v1/configuration/config-items$"
var tscLocationsPath = "^/api/v1/locations$"

func MockTscServiceIspsEnabledValue(URLPattern string, enabled bool) {
	tscBaseURL := fmt.Sprintf(URLPattern, locationInfo.RetailerID, locationInfo.EnvType)

	gock.New(tscBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		Get(tscPath).
		MatchParam("location-codes", locationInfo.MfcID).
		MatchParam("level", "mfc").
		MatchParam("categories", "isps").
		Reply(200).
		JSON(GetIspsEnabledConfigValueMock(enabled))
}

func MockTscLocationInfo(URLPattern string) {
	tscBaseURL := fmt.Sprintf(URLPattern, locationInfo.RetailerID, locationInfo.EnvType)

	gock.New(tscBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		Get(tscLocationsPath).
		Reply(200).
		JSON(GetTscLocationInfoMock(GetMockLocationInfo().Timezone))
}

func MockFailTscGetConfigItemsResponse(URLPattern string) {
	tscBaseURL := fmt.Sprintf(URLPattern, locationInfo.RetailerID, locationInfo.EnvType)

	gock.New(tscBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		Get(tscPath).
		MatchParam("location-codes", locationInfo.MfcID).
		MatchParam("level", "mfc").
		MatchParam("categories", "isps").
		Reply(400).
		JSON(GetTscConfigErrorResponseMock())
}

func MockFailTscLocationInfo(URLPattern string) {
	tscBaseURL := fmt.Sprintf(URLPattern, locationInfo.RetailerID, locationInfo.EnvType)

	gock.New(tscBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		Get(tscLocationsPath).
		Reply(500).
		JSON(GetTscLocationsErrorResponseMock())
}
