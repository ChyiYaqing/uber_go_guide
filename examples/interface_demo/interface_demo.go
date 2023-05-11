package interface_demo

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

var f1 F = S1{}

// f2.f() 可以修改底层数据，给接口变量 f2 赋值时使用的是对象指针
var f2 F = &S2{}
