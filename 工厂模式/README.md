# 简单工厂模式

## 角色和职责

简单工厂模式并不属于GoF的23种设计模式。他是开发者自发认为的一种非常简易的设计模式，其角色和职责如下：

| 角色     | 职责                                                         |
| :------- | ------------------------------------------------------------ |
| 工厂     | 简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。 |
| 抽象产品 | 简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。 |
| 具体产品 | 简单工厂模式所创建的具体实例对象。                           |

## 适用场景

工厂类负责创建的对象比较少，客户只知道传入了工厂类的参数，对于始何创建对象（逻辑）不关心。

## 类图

<img src="http://imgbed4926.oss-cn-hangzhou.aliyuncs.com/img/UML 图 (1).jpg" alt="UML 图 (1)" height="350dp" />

## 代码实现

```go
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

```

```
小米打电话
华为打电话
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

---

# 工厂方法模式

## 角色和职责

| 角色     | 职责                                                 |
| :------- | ---------------------------------------------------- |
| 抽象工厂 | 工厂方法模式的核心，任何工厂类都必须实现这个接口。   |
| 具体工厂 | 具体工厂类是抽象工厂的一个实现，负责实例化产品对象。 |
| 抽象产品 | 工厂方法模式所创建的所有对象的父类。                 |
| 具体产品 | 具体工厂所创建的具体实例对象。                       |

## 适用场景

当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。

## 类图

<img src="http://imgbed4926.oss-cn-hangzhou.aliyuncs.com/img/image-20221216202913130.png" alt="image-20221216202913130" height="400dp" />

## 代码实现

```go
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

```

```
小米打电话
华为打电话
```

## 优点

- 一个调用者想创建一个对象，只要知道其名称就可以了。
- 扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以。
- 屏蔽产品的具体实现，调用者只关心产品的接口。

## 缺点

每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。这并不是什么好事。

---

# 抽象工厂模式

## 角色和职责

| 角色     | 职责                                                         |
| :------- | ------------------------------------------------------------ |
| 抽象工厂 | 它声明了一组用于创建一族产品的方法，每一个方法对应一种产品。 |
| 具体工厂 | 它实现了在抽象工厂中声明的创建产品的方法，生成一组具体产品，这些产品构成了一个产品族，每一个产品都位于某个产品等级结构中。 |
| 抽象产品 | 它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法。 |
| 具体产品 | 它定义具体工厂生产的具体产品对象，实现抽象产品接口中声明的业务方法。 |

## 适用场景

- 系统中有多于一个的产品族。而每次只使用其中某一产品族。可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。

- 产品等级结构稳定。设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。

## 类图

<img src="http://imgbed4926.oss-cn-hangzhou.aliyuncs.com/img/image-20221217203354179.png" alt="image-20221217203354179" height="400dp" />

## 代码实现

```go
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

```

```
中国产的小米打电话
中国产的华为打电话
印度产的小米打电话
印度产的华为打电话
```



## 优点

抽象工厂模式除了具有工厂方法模式的优点外，最主要的优点就是可以在类的内部对产品族进行约束。所谓的产品族，一般或多或少的都存在一定的关联，抽象工厂模式就可以在类内部对产品族的关联关系进行定义和描述，而不必专门引入一个新的类来进行管理。

- 易于交换产品系列，由于具体的工场类，在使用的时候只需要在应用中初始化一次，所以改变工厂就很简单，只需要改变具体地工厂就能使用对应的配置信息。
- 它让具体地创建过程和客户端分离，客户端通过他们的抽象接口操作实例，产品的具体类名也和具体工厂分离，不会出现在客户端代码中。

## 缺点

抽象工厂模式在于难于应付“新对象”的需求变动。难以支持新种类的产品。难以扩展抽象工厂以生产新种类的产品。这是因为抽象工厂几乎确定了可以被创建的产品集合，支持新种类的产品就需要扩展该工厂接口，这将涉及抽象工厂类及其所有子类的改变。

