package main

import (
	"fmt"
	"sync"
)

// 不可导出
type singleton struct{}

func (s *singleton) DoSomeThing() {
	fmt.Println("单例的方法")
}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func main() {
	s := GetInstance()
	s.DoSomeThing()
}
