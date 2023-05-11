package main

import (
	"fmt"
	"sync"
	"time"
)

type F interface {
	f()
}

type S1 struct{}

// 使用值接收器的方法既可以通过值调用，也可以通过指针调用
func (s S1) f() {}

// 一个类型可以有值接收器方法集和指针接收器方法集
// 值接收器方法集是指针接收器方法集的子集
func (s *S1) fanother() {}

type S2 struct{}

// 如果希望接口方法修改基础数据，则必须使用指针传递(将对象指针赋值给接口变量)
// 使用指针接收器的方法只能通过指针或addressable values调用
func (s *S2) f() {}

type S3 struct{}

/* 零值Mutex是有效的
**/
type SMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}

/* 枚举类型
**/

type Operation int

const (
	Add Operation = iota + 1
	Subtract
	Multiply
)

/* 使用time处理时间
 **/
// time.Time 表达瞬时时间
func isActive(now, start, stop time.Time) bool {
	return (start.Before(now) || start.Equal(now) && now.Before(stop))
}

// time.Duration 表示时间段
func poll(delay time.Duration) {
	for {
		time.Sleep(delay)
	}
}

/*
Errors
错误类型:

错误包装: - 如果调用其他方法时出现错误，通常有三处处理方式可以选择
	1. 讲原始错误原样返回
	2. fmt.Errorf()搭配 %w 将错误添加进上下文返回 -  %w
	3. fmt.Errorf()搭配 %v 将错误添加进上下文返回

错误命名:
	对于存储为全局变量的错误值，根据是否导出，使用前缀Err 或 err

	对于自定义错误类型，请改用后缀Error

	无论调用方如何处理错误，它通常都应该只处理每个错误一次。
*/

func main() {
	// 编译器 Interface 合理性验证
	// var f1 F = S1{}
	// f1.f()

	// var f2 F = &S2{}
	// f2.f()

	// channel 的size要么是1，要么是无缓冲
	_ = make(chan int, 1)
	_ = make(chan int)
	var i interface
	i = 1
	_, ok := i.(string)
	if !ok {
		fmt.Println()
	}
}
