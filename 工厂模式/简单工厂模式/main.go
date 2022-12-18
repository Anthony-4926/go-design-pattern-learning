package main

import "fmt"

// ---------------------phone--------------------------------
// 抽象的产品
type Phone interface {
	Call()
}

type XiaoMi struct{}

func (phone *XiaoMi) Call() {
	fmt.Println("小米打电话")
}

type HuaWei struct{}

func (phone *HuaWei) Call() {
	fmt.Println("华为打电话")
}

// ----------------------PhoneFactory-------------------------
type PhoneFactory struct{}

func (*PhoneFactory) ProductPhone(model string) Phone {
	var phone Phone
	if model == "HuaWei" {
		phone = new(HuaWei)
	} else if model == "XiaoMi" {
		phone = new(XiaoMi)
	}
	return phone
}

func NewPhoneFactory() *PhoneFactory {
	return new(PhoneFactory)
}

// ------------------------main-------------------------------
func main() {
	phoneFactory := NewPhoneFactory()

	phone6 := phoneFactory.ProductPhone("XiaoMi")
	phone6.Call()

	phone13 := phoneFactory.ProductPhone("HuaWei")
	phone13.Call()

}
