package spec

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/to-com/poc-td/test"
)

func TestSpecForMFCConfig(t *testing.T) {
	Convey("Given: test MFC Config", t, func() {
		Convey("When: get MFC Config", func() {
			Convey("Then: no error", func() {
				out, err := test.GetApp(t).MFCConfig.Get(getCtx("test", "", "test"))

				So(err, ShouldBeNil)
				So(out.MfcID, ShouldNotBeEmpty)
			})
		})
	})
}
