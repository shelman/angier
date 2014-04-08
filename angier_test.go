package angier

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTransfer(t *testing.T) {

	Convey("With two structs of the same type", t, func() {

		type s struct {
			A int
			B string
			C bool
		}

		Convey("transferring from one to the other should copy over all of"+
			" the fields", func() {

			src := &s{
				A: 2,
				B: "hi",
				C: true,
			}

			dest := &s{}

			So(Transfer(src, dest), ShouldBeNil)
			So(src, ShouldResemble, dest)

		})

	})
}
