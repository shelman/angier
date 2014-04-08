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

	Convey("With two structs of different types", t, func() {

		type s1 struct {
			// two matching fields
			MatchOne int
			MatchTwo bool

			// two fields that don't appear in the other struct
			OnlySrcOne int
			OnlySrcTwo string

			// two fields with the same name but the wrong type
			WrongTypeOne int
			WrongTypeTwo bool
		}

		type s2 struct {
			// two matching fields
			MatchOne int
			MatchTwo bool

			// two fields that don't appear in the other struct
			OnlyDestOne int
			OnlyDestTwo string

			// two fields with the same name but the wrong type
			WrongTypeOne bool
			WrongTypeTwo int
		}

		Convey("transferring from one to the other should copy only the"+
			" fields with the same name and type, and leave the other fields"+
			" untouched", func() {

			src := &s1{
				MatchOne:     2,
				MatchTwo:     true,
				OnlySrcOne:   1,
				OnlySrcTwo:   "hi",
				WrongTypeOne: 3,
				WrongTypeTwo: true,
			}

			dest := &s2{}

			So(Transfer(src, dest), ShouldBeNil)
			So(dest, ShouldResemble, &s2{MatchOne: 2, MatchTwo: true})

		})

	})
}
