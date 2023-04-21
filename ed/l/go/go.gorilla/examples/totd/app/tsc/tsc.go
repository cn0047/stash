package tsc

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/to-com/poc-td/app/payload"
	"github.com/to-com/poc-td/app/util"
)

const (
	// BaseURLFormat holds format for service catalog URL.
	BaseURLFormat      = "https://service-catalog-%s-%s.tom.to.com"
	IsFlowRacksEnabled = "IS_FLOW_RACKS_ENABLED"
	OSRCategory        = "osr"
	OutboundCategory   = "outbound"
)

type FlowRackID string

// toServiceCatalog describes service catalog service.
type toServiceCatalog interface {
	// GetConfig gets config from service catalog (including OSR config + FlowRacks).
	GetConfig(ctx context.Context, input payload.GetConfigInput) (payload.MFCConfig, error)
}

// Service represents service catalog service.
type Service struct {
	httpClient *http.Client
	log        *zap.SugaredLogger
}

// NewService creates new service catalog instance.
func NewService(log *zap.SugaredLogger) toServiceCatalog {
	c := &http.Client{
		Transport: &http.Transport{
			DialContext:         (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
			TLSHandshakeTimeout: 2 * time.Second,
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
		},
		Timeout: 5 * time.Second,
	}

	s := &Service{
		httpClient: c,
		log:        log,
	}

	return s
}

// GetConfig {@inheritdoc}.
func (s *Service) GetConfig(ctx context.Context, input payload.GetConfigInput) (c payload.MFCConfig, err error) {
	var g errgroup.Group

	g.Go(func() error {
		conf, err := s.getOSRConfig(ctx, input)
		if err != nil {
			return fmt.Errorf("failed to get OSR config, err: %w", err)
		}

		c.ClientID = conf.ClientID
		c.Env = conf.Env
		c.MfcID = conf.MfcID
		c.UpdatedAt = conf.UpdatedAt
		c.ErrorRamp = conf.ErrorRamp
		c.Count = conf.Count
		c.Depth = conf.Depth
		c.Start = conf.Start
		c.IDGen = conf.IDGen
		c.LaneMapping = conf.LaneMapping
		c.ExpressLaneMapping = conf.ExpressLaneMapping

		return nil
	})

	g.Go(func() error {
		flowRacksEnabled, err := s.IsFlowRacksEnabled(ctx, input)
		if err != nil {
			return fmt.Errorf("failed to check IsFlowRacksEnabled, err: %w", err)
		}
		if !flowRacksEnabled {
			return nil
		}

		fr, err := s.getFlowRacksMapping(ctx, input)
		if err != nil {
			return fmt.Errorf("failed to get flowRacksMapping, err: %w", err)
		}

		c.FlowRacksMapping = fr

		return nil
	})

	err = g.Wait()
	if err != nil {
		return c, fmt.Errorf("failed to get TSC config, err: %w", err)
	}

	return c, nil
}

// getOSRConfig gets configs from service catalog with category OSR.
func (s *Service) getOSRConfig(
	ctx context.Context, input payload.GetConfigInput,
) (conf payload.MFCConfig, err error) {
	url := fmt.Sprintf(
		BaseURLFormat+"/api/v1/configuration/config-items?value-format=json&location-codes=%s&level=mfc&categories=%s",
		input.Retailer, input.Env, input.MFC, OSRCategory,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return conf, fmt.Errorf("failed to create new request, err: %w", err)
	}

	req.Header.Add("X-Token", input.Token)

	response, err := s.httpClient.Do(req)
	if err != nil {
		return conf, fmt.Errorf("failed to perform request, err: %w", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			s.log.With(zap.Error(err)).Error("failed to close response bode")
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return conf, fmt.Errorf("failed to read response body, err: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return conf, fmt.Errorf("an error occurred while retreiving OSR configs. Status code: %s, body: %s",
			response.Status, body)
	}

	responsePayload := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(body, &responsePayload); err != nil {
		return conf, fmt.Errorf("unable to decode OSR configurations from the Service Catalog")
	}

	conf, err = s.parseOSRConfigs(responsePayload)
	if err != nil {
		return conf, fmt.Errorf("failed to parse OSR configs, err: %w", err)
	}

	conf.ClientID = input.Retailer
	conf.Env = input.Env
	conf.MfcID = input.MFC

	return conf, nil
}

func (s *Service) getFlowRacksMapping(
	ctx context.Context, input payload.GetConfigInput,
) (m map[int64]string, err error) {
	url := fmt.Sprintf(
		BaseURLFormat+"/api/v1/flow-racks?mfc-tom-code=%s", input.Retailer, input.Env, input.MFC,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return m, fmt.Errorf("failed to create new request, err: %w", err)
	}

	req.Header.Add("X-Token", input.Token)

	response, err := s.httpClient.Do(req)
	if err != nil {
		return m, fmt.Errorf("failed to perform request, err: %w", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			s.log.With(zap.Error(err)).Error("failed to close response bode")
		}
	}()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return m, fmt.Errorf("failed to read response body, err: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return m, fmt.Errorf("got bad status code: %s, body: %s", response.Status, body)
	}

	responsePayload := make(map[string]interface{}, 0)
	if err := json.Unmarshal(body, &responsePayload); err != nil {
		return m, fmt.Errorf("failed to unmarshal response body: %s", body)
	}

	rawFlowRacks, ok := responsePayload["flow-racks"]
	if !ok {
		return m, fmt.Errorf("failed to find flow-racks in response body: %v", responsePayload)
	}
	resultFlowRacks, err := util.ConvertToMapWithInt64Keys[string, string](rawFlowRacks)
	if err != nil {
		return m, fmt.Errorf("failed to convert keys, err: %w", err)
	}

	return resultFlowRacks, nil
}

func (s *Service) parseOSRConfigs(payload []map[string]interface{}) (conf payload.MFCConfig, err error) {
	conf.LaneMapping = make(map[int64]int64)
	conf.ExpressLaneMapping = make(map[int64]int64)

	for _, item := range payload {
		switch item["name"] {
		case "dispatch_error_ramp":
			conf.ErrorRamp, err = getIntWithoutPrefix(item["value"], "DISPATCH")

		case "dispatch_ramp_count":
			conf.Count, err = getInt(item["value"])

		case "dispatch_ramp_depth":
			conf.Depth, err = getInt(item["value"])

		case "dispatch_start_ramp":
			conf.Start, err = getInt(item["value"])

		case "dispatch_ramp_id_gen":
			conf.IDGen, err = getString(item["value"])

		case "dispatch_express_lanes":
			strVal, isString := item["value"].(string)
			if isString {
				valSlice := make([]string, 0)
				err := json.Unmarshal([]byte(strVal), &valSlice)
				if err != nil {
					return conf, fmt.Errorf("failed to unmarshal TSC config dispatch_express_lanes: %s", strVal)
				}

				for i, v := range valSlice {
					idx := i + 1 // first value must be 1 not 0.
					index := int64(idx)
					intVal, err := getIntWithoutPrefix(v, "DISPATCH")
					if err != nil {
						return conf, fmt.Errorf("failed to get int for config dispatch_express_lanes: %#v", v)
					}
					conf.ExpressLaneMapping[index] = intVal
				}
			}
			interfaceVal, isInterface := item["value"].([]interface{})
			if isInterface {
				for i, v := range interfaceVal {
					idx := i + 1 // first value must be 1 not 0.
					index := int64(idx)
					strVal, ok := v.(string)
					if !ok {
						return conf, fmt.Errorf("failed to convert TSC config dispatch_express_lanes: %#v", item["value"])
					}
					intVal, err := getIntWithoutPrefix(strVal, "DISPATCH")
					if err != nil {
						return conf, fmt.Errorf("failed to get int value for config dispatch_express_lanes: %#v", v)
					}
					conf.ExpressLaneMapping[index] = intVal
				}
			}
			if !isString && !isInterface {
				return conf, fmt.Errorf("failed to parse TSC config dispatch_express_lanes: %#v", item["value"])
			}
		}
		if err != nil {
			return conf, fmt.Errorf(
				"failed to parse TSC config: %s, with value: %#v, err: %w", item["name"], item["value"], err,
			)
		}
	}

	return conf, nil
}

func getInt(item interface{}) (int64, error) {
	float64Val, ok := item.(float64)
	if ok {
		return int64(float64Val), nil
	}

	val, err := getString(item)
	if err != nil {
		return 0, fmt.Errorf("failed to get string for: %#v, err: %w", item, err)
	}

	int64Val, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert TSC config to int: %#v", item)
	}

	return int64Val, nil
}

func getIntWithoutPrefix(item interface{}, prefix string) (int64, error) {
	val, err := getString(item)
	if err != nil {
		return 0, fmt.Errorf("failed to get string for: %#v", item)
	}

	intVal, err := strconv.ParseInt(strings.TrimPrefix(val, prefix), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to convert TSC config to int: %#v", item)
	}

	return intVal, nil
}

func getString(item interface{}) (string, error) {
	val, ok := item.(string)
	if !ok {
		return "", fmt.Errorf("failed to convert TSC config to string: %#v", item)
	}

	val = strings.Replace(val, `\"`, ``, -1)
	val = strings.Replace(val, `"`, ``, -1)

	return val, nil
}

func (s *Service) IsFlowRacksEnabled(
	ctx context.Context, input payload.GetConfigInput,
) (res bool, err error) {

	url := fmt.Sprintf(
		BaseURLFormat+"/api/v1/configuration/config-items?location-codes=%s&level=mfc&categories=%s",
		input.Retailer, input.Env, input.MFC, OutboundCategory,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create new request for retrieving FlowRack configuration err: %w", err)
	}

	req.Header.Add("X-Token", input.Token)

	response, err := s.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to perform request for retrieving FlowRack configuration, err: %w", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			s.log.With(zap.Error(err)).Error("failed to close response bode")
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response body, err: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return false, fmt.Errorf("an error occurred while retreiving FlowRack configuration. Status code: %s, body: %s",
			response.Status, body)
	}

	responsePayload := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(body, &responsePayload); err != nil {
		return false, fmt.Errorf("unable to decode FlowRack configuration from the Service Catalog")
	}

	for _, item := range responsePayload {
		if item["name"] == IsFlowRacksEnabled {
			strVal, ok := item["value"].(string)
			if !ok {
				return false, fmt.Errorf("got unexpected value type: %#v", item["value"])
			}
			enabled := strVal == "true"
			return enabled, nil
		}
	}

	// Tolerate missing config, probably missing config means default false value,
	// anyway this case will be logged for visibility, for now it just won't fail with error.
	res = false
	s.log.With(
		zap.String("retailer", input.Retailer), zap.String("env", input.Env), zap.String("mfc", input.MFC),
	).Warn("failed to find FlowRack configuration")

	return res, nil
}

func (s *Service) GetFlowRackByRampID(
	ctx context.Context, input payload.GetFlowRackByRampIDInput,
) (FlowRackID, error) {

	url := fmt.Sprintf(
		BaseURLFormat+"/api/v1/flow-racks/dispatch-ramp/%s?mfc-tom-code=%s",
		input.Retailer, input.Env, input.RampID, input.MFC,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create new request for retrieving FlowRack address, err: %w", err)
	}

	req.Header.Add("X-Token", input.Token)

	response, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request for retrieving FlowRack address, err: %w", err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			s.log.With(zap.Error(err)).Error("failed to close response bode")
		}
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body, err: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return "", fmt.Errorf("an error occurred while retreiving FlowRack address. Status code: %s, body: %s",
			response.Status, body)
	}

	responsePayload := make(map[string]string, 1)
	if err := json.Unmarshal(body, &responsePayload); err != nil {
		return "", fmt.Errorf("unable to decode FlowRack address from the Service Catalog")
	}

	rawFlowRack := responsePayload["flow-rack-address"]

	return FlowRackID(rawFlowRack), nil
}
