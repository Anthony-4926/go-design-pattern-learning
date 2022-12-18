# 适用场景

- 系统只需要一个实例对象，如系统要求提供一个唯一的序列号生成器或资源管理器，或者需要考虑资源消耗太大而只允许创建一个对象。

- 客户调用类的单个实例只允许使用一个公共访问点，除了该公共访问点，不能通过其他途径访问该实例。

# 代码实现

## 饿汉式

饿汉式在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存。饿汉式的好处是，不会出现线程并发创建，导致多个单例的出现，但是缺点是如果这个单例对象在业务逻辑没有被使用，也会客观的创建一块内存对象。

```go
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

```

```
单例的方法
```

与饿汉式相对应的是懒汉式

## 懒汉式-推演过程

单例模式要解决的问题是：**保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用。**

推演过程如下

### 只在实例为nil时创建实例

```go
package main

import (
	"fmt"
)

// 不可导出
type singleton struct{}

func (s *singleton) DoSomeThing() {
	fmt.Println("单例的方法")
}

var instance *singleton

func GetInstance() *singleton {
	//只有首次GetInstance()方法被调用，才会生成这个单例的实例
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func main() {
	s := GetInstance()
	s.DoSomeThing()
}

```

### 线程安全的单例模式

上面的“懒汉式”实现是非线程安全的设计方式，也就是如果多个线程或者协程同时首次调用`GetInstance()`方法有概率导致多个实例被创建，则违背了单例的设计初衷。那么在上面的基础上进行修改，可以利用`Sync.Mutex`进行加锁，保证线程安全。这种线程安全的写法，有个最大的缺点就是每次调用该方法时都需要进行锁操作，在性能上相对不高效，具体的实现改进如下：

```go
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

var lock sync.Mutex

func GetInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()
	//只有首次GetInstance()方法被调用，才会生成这个单例的实例
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func main() {
	s := GetInstance()
	s.DoSomeThing()
}

```

### 高性能的单例模式

上面代码虽然解决了线程安全，但是每次调用`GetInstance()`都要加锁会极大影响性能。所以接下来可以借助`sync/atomic`来进行内存的状态存留来做互斥。`atomic`就可以自动加载和设置标记，代码如下：

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 不可导出
type singleton struct{}

func (s *singleton) DoSomeThing() {
	fmt.Println("单例的方法")
}

var instance *singleton

// 标记
var initialized uint32
var lock sync.Mutex

func GetInstance() *singleton {
	//如果标记为被设置，直接返回，不加锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	lock.Lock()
	defer lock.Unlock()
	//只有首次GetInstance()方法被调用，才会生成这个单例的实例
	if instance == nil {
		instance = new(singleton)
		//设置标记位
		atomic.StoreUint32(&initialized, 1)
	}
	return instance
}

func main() {
	s := GetInstance()
	s.DoSomeThing()
}

```

### go特性的单例模式

上述的实现其实Golang有个方法已经帮助开发者实现完成，就是`Once`模块

```go
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

```

