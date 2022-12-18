package main

import "fmt"

// ---------------------抽象层--------------------------------
type XiaoMi interface {
	Call()
}
type HuaWei interface {
	Call()
}
type PhoneFactory interface {
	ProductXiaoMi() XiaoMi
	ProductHuaWei() HuaWei
}

// ---------------------实现层--------------------------------
type XiaoMiMadeInChina struct{}

func (phone *XiaoMiMadeInChina) Call() {
	fmt.Println("中国产的小米打电话")
}

type XiaoMiMadeInIndian struct{}

func (phone *XiaoMiMadeInIndian) Call() {
	fmt.Println("印度产的小米打电话")
}

type HuaWeiMadeInChina struct{}

func (phone *HuaWeiMadeInChina) Call() {
	fmt.Println("中国产的华为打电话")
}

type HuaWeiMadeInIndian struct{}

func (phone *HuaWeiMadeInIndian) Call() {
	fmt.Println("印度产的华为打电话")
}

type ChinaPhoneFactory struct{}

func (f ChinaPhoneFactory) ProductXiaoMi() XiaoMi {
	return new(XiaoMiMadeInChina)
}
func (f ChinaPhoneFactory) ProductHuaWei() HuaWei {
	return new(HuaWeiMadeInChina)
}

type IndianPhoneFactory struct{}

func (f IndianPhoneFactory) ProductXiaoMi() XiaoMi {
	return new(XiaoMiMadeInIndian)
}
func (f IndianPhoneFactory) ProductHuaWei() HuaWei {
	return new(HuaWeiMadeInIndian)
}

func main() {
	// 创建一个中国工厂
	var ChinaF PhoneFactory
	ChinaF = new(ChinaPhoneFactory)

	ChinaMi := ChinaF.ProductXiaoMi()
	ChinaMi.Call()

	ChinaHuaWei := ChinaF.ProductHuaWei()
	ChinaHuaWei.Call()

	// 创建一个印度工厂
	var IndianF PhoneFactory
	IndianF = new(IndianPhoneFactory)

	IndianMi := IndianF.ProductXiaoMi()
	IndianMi.Call()

	IndianHuaWei := IndianF.ProductHuaWei()
	IndianHuaWei.Call()

}
