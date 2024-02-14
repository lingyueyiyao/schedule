package service

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCore_CallSerialNumber(t *testing.T) {
	var core = NewCore()

	Convey("test shower for single, double and more, test barrel", t, func() {
		core.CallSerialNumber(201)
		core.CallSerialNumber(205)
		core.CallSerialNumber(222)
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)
		core.GetSerialNumber(TypeBarrel, 2)
		core.GetSerialNumber(TypeBarrel, 2)
		core.GetSerialNumber(TypeBarrel, 2)
		core.GetSerialNumber(TypeBarrel, 2)
		core.GetSerialNumber(TypeBarrel, 2)

		core.CallSerialNumber(205)
		core.CallSerialNumber(206)
		core.CallSerialNumber(207)
		core.CallSerialNumber(208)
		core.CallSerialNumber(205)
		core.CallSerialNumber(201)
		core.CallSerialNumber(202)
		core.CallSerialNumber(203)
		core.CallSerialNumber(221)
	})

	Convey("test three type room for single or double", t, func() {
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)

		core.CallSerialNumber(222)
		core.CallSerialNumber(221)
		core.CallSerialNumber(220)
	})

	Convey("test barrel room for single or double", t, func() {
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)

		core.CallSerialNumber(205)
		core.CallSerialNumber(206)
		core.CallSerialNumber(207)
	})
}

func TestCore_ShareRoom(t *testing.T) {
	var core = NewCore()

	Convey("test share room", t, func() {
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 1)

		err := core.ShareRoom(1, 4, 5)
		So(err, ShouldBeNil)
		core.CallSerialNumber(201)
		core.CallSerialNumber(221)
		core.CallSerialNumber(202)
		core.CallSerialNumber(222)
	})
}

func TestCore_DelSerialNumber(t *testing.T) {
	var core = NewCore()

	Convey("test del customer", t, func() {
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)
		core.GetSerialNumber(TypeBarrel, 2)
		core.GetSerialNumber(TypeBarrel, 2)

		err := core.DelSerialNumber(TypeShower, 4)
		So(err, ShouldNotBeNil)
		err = core.DelSerialNumber(TypeShower, 3)
		So(err, ShouldBeNil)
		core.CallSerialNumber(221)
		core.CallSerialNumber(201)
		core.CallSerialNumber(202)
		err = core.DelSerialNumber(TypeBarrel, 3)
		So(err, ShouldNotBeNil)
		err = core.DelSerialNumber(TypeBarrel, 1)
		So(err, ShouldBeNil)
		core.CallSerialNumber(205)
		core.CallSerialNumber(206)
	})
}

func TestCore_ExchangeBathType(t *testing.T) {
	var core = NewCore()

	Convey("test customer change bathType", t, func() {
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)
		core.GetSerialNumber(TypeBarrel, 2)
		core.GetSerialNumber(TypeBarrel, 2)

		ShouldEqual(len(core.waitQueue.showerQueue), 3)
		ShouldEqual(len(core.waitQueue.barrelQueue), 2)

		core.ExchangeBathType(TypeShower, TypeBarrel, 1)
		ShouldEqual(len(core.waitQueue.showerQueue), 2)
		ShouldEqual(len(core.waitQueue.barrelQueue), 3)

		core.ExchangeBathType(TypeBarrel, TypeShower, 2)
		ShouldEqual(len(core.waitQueue.showerQueue), 2)
		ShouldEqual(len(core.waitQueue.barrelQueue), 3)

		for _, c := range core.waitQueue.showerQueue {
			fmt.Print(c.serialNum, " ")
		}
		fmt.Println()
		for _, c := range core.waitQueue.barrelQueue {
			fmt.Print(c.serialNum, " ")
		}
	})
}

func TestCore_ChangeSerialNumAward(t *testing.T) {
	var core = NewCore()

	Convey("test change customer's serial number", t, func() {
		core.GetSerialNumber(TypeShower, 1)
		core.GetSerialNumber(TypeShower, 2)
		core.GetSerialNumber(TypeShower, 3)

		core.DelSerialNumber(TypeShower, 1)
		for _, c := range core.waitQueue.showerQueue {
			fmt.Println(c.serialNum, c.number, " ")
		}

		core.ChangeSerialNumAward(TypeShower, 3, 1)
		for _, c := range core.waitQueue.showerQueue {
			fmt.Println(c.serialNum, c.number, " ")
		}
	})
}
