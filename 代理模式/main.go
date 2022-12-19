package main

import "fmt"

// ----------------------抽象的主题--------------------
type ApplyPatent interface {
	writeDoc()
}

// ----------------------具体的主题--------------------
type ApplyUtilityModelPatent struct {
}

func (ApplyUtilityModelPatent) writeDoc() {
	fmt.Println("撰写实用新型专利的创新文档")
}

type ApplyInventionPatent struct {
}

func (ApplyInventionPatent) writeDoc() {
	fmt.Println("撰写发明专利的创新文档")
}

// ----------------------具体的主题--------------------
type Proxy struct {
	Patent ApplyPatent
}

func (p Proxy) writeDoc() {
	fmt.Println("注册账号")

	//	调用
	p.Patent.writeDoc()

	fmt.Println("递交申请")
}

func main() {
	var proxy Proxy
	proxy = Proxy{
		Patent: ApplyInventionPatent{},
	}
	proxy.writeDoc()

	proxy = Proxy{
		Patent: ApplyUtilityModelPatent{},
	}
	proxy.writeDoc()
}
