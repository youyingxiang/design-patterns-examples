// 单例模式 (Singleton)
// 现实世界的例子
// 一次只能有一个国家的总统。无论何时打电话，都必须将同一位总统付诸行动。
// 简单来说:确保只创建特定类的一个对象。
// 单例模式实际上被认为是反模式，应该避免过度使用它。
// 它不一定是坏的，可能有一些有效的用例，但应谨慎使用，因为它在您的应用程序中引入了一个全局状态，并且在一个地方更改它可能会影响其他区域，
// 并且它可能变得非常难以调试。关于它们的另一个坏处是它使你的代码紧密耦合加上嘲弄单例可能很困难。

package examples

import (
	"sync"
)

// President 构建一个示例结构体
type President struct {
	Name string
}

// 设置一个私有变量作为每次要返回的单例
var once sync.Once
var instance1 *President

// GetInstanceThreadUnsafe 写一个可以获取单例的方法(懒汉方式)
// 存在线程安全问题，高并发时有可能创建多个对象
func GetInstanceThreadUnsafe(name string) *President {
	if instance1 == nil {
		instance1 = &President{Name: name}
	}
	return instance1
}

// 线程安全
var instance2 *President

// GetInstanceThreadSafety 写一个可以获取单例的方法(避免了每次加锁，提高代码效率)
func GetInstanceThreadSafety(name string) *President {
	once.Do(func() {
		instance2 = &President{Name: name}
	})
	return instance2
}
