package examples

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// command: go test -v singleton_test.go singleton.go
func TestSingleton(t *testing.T) {
	t.Run("线程非安全测试", testGetInstanceThreadUnsafe)
	t.Run("线程安全测试", testGetInstanceThreadSafety)
}

// 单线程测试
func testGetInstanceThreadUnsafe(t *testing.T) {
	// 普通测试
	instance1 := GetInstanceThreadUnsafe("yxx")
	instance2 := GetInstanceThreadUnsafe("张三")
	// 如果单列是一个指针的值
	if fmt.Sprintf("%p", instance1) != fmt.Sprintf("%p", instance2) {
		t.Fail()
	}
}

// 多线程测试
func testGetInstanceThreadSafety(t *testing.T) {

	var wg sync.WaitGroup
	presidentChan := make(chan *President, 1000)
	// 模拟 1000 个并发获取单例
	for i := 0; i < 1000; i++ {
		name := "名称" + strconv.Itoa(i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := GetInstanceThreadSafety(name)
			presidentChan <- instance
		}()

	}
	wg.Wait()

	close(presidentChan)

	temp := make(map[*President]*President)

	for v := range presidentChan {
		temp[v] = v
	}

	// 判断是否有重复的
	t.Logf("获取了 %d 个实例", len(temp))

	if len(temp) > 1 {
		t.Fail()
	}

}
