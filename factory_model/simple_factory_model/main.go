package main

import "fmt"

// ---------------------phone--------------------------------
// 抽象的产品
type Phone interface {
	Call()
}

type XiaoMi6 struct{}

func (phone *XiaoMi6) Call() {
	fmt.Println("小米6打电话")
}

type XiaoMi13 struct{}

func (phone *XiaoMi13) Call() {
	fmt.Println("小米13打电话")
}

// ----------------------PhoneFactory-------------------------
type PhoneFactory struct{}

func (*PhoneFactory) ProductPhone(model string) Phone {
	var phone Phone
	if model == "XiaoMi13" {
		phone = new(XiaoMi13)
	} else if model == "XiaoMi6" {
		phone = new(XiaoMi6)
	}
	return phone
}

func NewPhoneFactory() *PhoneFactory {
	return new(PhoneFactory)
}

// ------------------------main-------------------------------
func main() {
	phoneFactory := NewPhoneFactory()

	phone6 := phoneFactory.ProductPhone("XiaoMi6")
	phone6.Call()

	phone13 := phoneFactory.ProductPhone("XiaoMi13")
	phone13.Call()

}
