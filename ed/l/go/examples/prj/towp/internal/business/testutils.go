package business

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/to-com/wp/internal/dto"
	"reflect"
)

type pubsubwpCreatedMatcher struct {
	event dto.wpCreatedEvent
}

type pubsubFireTriggerMatcher struct {
	event dto.TriggersFiredEvent
}

func (pm pubsubwpCreatedMatcher) Matches(x any) bool {
	arg, ok := x.(dto.wpCreatedEvent)

	if !ok {
		return false
	}

	return reflect.DeepEqual(arg.wp, pm.event.wp) && !arg.CreatedAt.IsZero()
}

func (pm pubsubFireTriggerMatcher) Matches(x any) bool {
	arg, ok := x.(dto.TriggersFiredEvent)

	if !ok {
		return false
	}

	return reflect.DeepEqual(arg.Cutoffs, pm.event.Cutoffs) && arg.ScheduleType == pm.event.ScheduleType
}

func (pm pubsubwpCreatedMatcher) String() string {
	return fmt.Sprintf("msg matched %v", pm.event)
}

func (pm pubsubFireTriggerMatcher) String() string {
	return fmt.Sprintf("msg matched %v", pm.event)
}

func EqwpCreatedPubsubEvent(expectedEvent dto.wpCreatedEvent) gomock.Matcher {
	return pubsubwpCreatedMatcher{expectedEvent}
}

func EqPubsubFireTriggerEvent(expectedEvent dto.TriggersFiredEvent) gomock.Matcher {
	return pubsubFireTriggerMatcher{expectedEvent}
}
