package main

import "fmt"

// 不可导出
type singleton struct{}

func (s *singleton) DoSomeThing() {
	fmt.Println("单例的方法")
}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

func main() {
	s := GetInstance()
	s.DoSomeThing()
}
