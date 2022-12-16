# 简单工厂模式

## 角色和职责

简单工厂模式并不属于GoF的23种设计模式。他是开发者自发认为的一种非常简易的设计模式，其角色和职责如下：

| 角色     | 职责                                                         |
| :------- | ------------------------------------------------------------ |
| 工厂     | 简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。 |
| 抽象产品 | 简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。 |
| 具体产品 | 简单工厂模式所创建的具体实例对象。                           |

## 适用范围

工厂类负责创建的对象比较少，客户只知道传入了工厂类的参数，对于始何创建对象（逻辑）不关心。

## 类图

<img src="http://imgbed4926.oss-cn-hangzhou.aliyuncs.com/img/UML 图 (1).jpg" alt="UML 图 (1)" height="350dp" />

## 代码实现

```go
package main

import "fmt"

// ---------------------phone--------------------------------
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

```

```
小米6打电话
小米13打电话
```

## 优点

- 工厂类含有必要的判断逻辑，可以决定在什么时候创建哪一个产品类的实例，客户端可以免除直接创建产品对象的责任，而仅仅“消费”产品；简单工厂模式通过这种做法实现了对责任的分割，它提供了专门的工厂类用于创建对象。
- 客户端无须知道所创建的具体产品类的类名，只需要知道具体产品类所对应的参数即可，对于一些复杂的类名，通过简单工厂模式可以减少使用者的记忆量。
- 通过引入配置文件，可以在不修改任何客户端代码的情况下更换和增加新的具体产品类，在一定程度上提高了系统的灵活性。

## 缺点

- 由于工厂类集中了所有产品创建逻辑，一旦不能正常工作，整个系统都要受到影响。
- 使用简单工厂模式将会增加系统中类的个数，在一定程序上增加了系统的复杂度和理解难度。
- 系统扩展困难，一旦添加新产品就不得不修改工厂逻辑，在产品类型较多时，有可能造成工厂逻辑过于复杂，不利于系统的扩展和维护。
- 简单工厂模式由于使用了静态工厂方法，造成工厂角色无法形成基于继承的等级结构。