package spec

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/cn007b/servicetemplate/test"
)

func TestSpecForMFCConfig(t *testing.T) {
	Convey("Given: test config", t, func() {
		Convey("When: get config", func() {
			Convey("Then: no error", func() {
				So(test.GetApp(t).Config.Env, ShouldEqual, "test")
			})
		})
	})
}
