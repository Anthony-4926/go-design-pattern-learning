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
// 抽象的工厂
type PhoneFactory interface {
	ProductPhone() Phone
}

type XiaoMiFactory struct{}

func (f XiaoMiFactory) ProductPhone() Phone {
	return new(XiaoMi)
}

type HuaWeiFactory struct{}

func (f HuaWeiFactory) ProductPhone() Phone {
	return new(HuaWei)
}

// ------------------------main-------------------------------
func main() {
	// 生产小米手机
	var xiaoMiFactory PhoneFactory
	xiaoMiFactory = new(XiaoMiFactory)
	xiaoMi := xiaoMiFactory.ProductPhone()
	xiaoMi.Call()

	//	生产华为手机
	var huaWeiFactory PhoneFactory
	huaWeiFactory = new(HuaWeiFactory)
	huaWei := huaWeiFactory.ProductPhone()
	huaWei.Call()
}
