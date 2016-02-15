package tutil

import (
	"testing"

	cv "github.com/glycerine/goconvey/convey"
)

func TestMmToHm(t *testing.T) {

	cv.Convey("Given the MmToHm() utility function in time_util.go", t, func() {

		cv.Convey("It should be the inverse of HmToMm()", func() {
			cv.So(MmToHm(HmToMm(930)), cv.ShouldEqual, 930)
			cv.So(MmToHm(HmToMm(931)), cv.ShouldEqual, 931)
			cv.So(MmToHm(HmToMm(1600)), cv.ShouldEqual, 1600)
			cv.So(MmToHm(HmToMm(1200)), cv.ShouldEqual, 1200)
		})
	})
}
