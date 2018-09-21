package realtimelog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"time"

	"go-app/config"
	"go-app/service/ping"
)

func Ping(ctx context.Context, msg string) (r *http.Response, err error) {
	payload := map[string]string{
		"project":   config.ProjectID,
		"namespace": "health-check",
		"msg":       msg}

	j, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("[20180703-003] Failed to marshal payload, error: %v", err)
	}

	return ping.Exec(ctx, config.RealTimeLogURL, "application/json", bytes.NewBuffer(j))
}

func Pinging(ctx context.Context, msg string) (r map[int]*http.Response, err error) {
	r = make(map[int]*http.Response, config.RealTimeLogPingingThreshold)

	for i := 0; i < config.RealTimeLogPingingThreshold; i++ {
		res, err := Ping(ctx, msg+"-"+strconv.Itoa(i))
		if err == nil {
			r[i] = res
		} else {
			return r, fmt.Errorf("[20180703-004] Failed to perform pinging, error: %v", err)
		}

		time.Sleep(config.RealTimeLogPingingSleepLimit * time.Millisecond)
	}

	return r, nil
}
