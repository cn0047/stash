package testutils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/to-com/wp/internal/common"
	"github.com/to-com/wp/internal/dto"
	"github.com/to-com/wp/internal/mocks"
	"log"
	"os"
	"testing"
)

func ReadFileAsString(t *testing.T, file string) string {
	t.Helper()

	data, err := os.ReadFile(file)

	if err != nil {
		t.Fail()
	}

	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, data)

	if err != nil {
		log.Printf("unable to read json file at %s: %v", file, err)

		return ""
	}

	return buffer.String()
}

func BuildCtx() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, common.CtxKeyEnv, "dev")
	ctx = context.WithValue(ctx, common.CtxKeyToken, "myToken")

	return ctx
}

func responseSchedulesMatchRequest(waveReq dto.WaveRequest, waveResp dto.WaveResponse) bool {
	expectedSchedulesLen := 0
	if waveReq.PrelimTime != "" {
		expectedSchedulesLen++
	}
	if waveReq.DeltaTime != "" {
		expectedSchedulesLen++
	}
	if len(waveResp.Schedules) != expectedSchedulesLen {
		return false
	}
	for _, schedule := range waveResp.Schedules {
		if schedule.ScheduleType == dto.Prelim && schedule.ScheduleTime != waveReq.PrelimTime {
			return false
		}
		if schedule.ScheduleType == dto.Delta && schedule.ScheduleTime != waveReq.DeltaTime {
			return false
		}
	}

	return true
}

func waveMatch(waveReq dto.WaveRequest, waveResp dto.WaveResponse) bool {
	return waveReq.Cutoff == waveResp.Cutoff &&
		waveReq.FromTime == waveResp.FromTime &&
		waveReq.ToTime == waveResp.ToTime &&
		responseSchedulesMatchRequest(waveReq, waveResp)
}

func findWaveByCutoff(wpReq dto.wpRequest, cutoff string) *dto.WaveRequest {
	for _, wave := range wpReq.Waves {
		if wave.Cutoff == cutoff {
			return &wave
		}
	}

	return nil
}

func Validatewp(t *testing.T, wpRequest dto.wpRequest, wpResponse dto.wpResponse) {
	assert.Equal(t, fmt.Sprintf("%v", mocks.GetUserPermissionMock()["user_id"]), wpResponse.CreatedBy)
	assert.NotEmpty(t, wpResponse.CreatedAt)

	assert.Equal(t, len(wpRequest.Waves), len(wpResponse.Waves))

	for _, wave := range wpResponse.Waves {
		w := findWaveByCutoff(wpRequest, wave.Cutoff)
		require.NotEmptyf(t, w, "wave %v from response does not match any request waves", wave)

		if w != nil {
			assert.Truef(t, waveMatch(*w, wave), "wave %v from response does not match wave from request %v", wave, &w)
		}
	}
}

type contextMatcher struct {
	ctx *context.Context
}

func (ctxm contextMatcher) Matches(x any) bool {
	ctx, ok := x.(context.Context)

	if !ok {
		return false
	}

	return common.GetCtxEnv(*ctxm.ctx) == common.GetCtxEnv(ctx) && common.GetCtxToken(*ctxm.ctx) == common.GetCtxToken(ctx)
}

func (ctxm contextMatcher) String() string {
	return fmt.Sprintf("context matched %v", ctxm.ctx)
}

func EqContext(expectedContext context.Context) gomock.Matcher {
	return contextMatcher{&expectedContext}
}
