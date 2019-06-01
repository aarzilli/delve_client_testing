package main

type Iface interface {
	Method()
}

type IfaceImpl int

func (r IfaceImpl) Method() {
	println("hello world")
}

func main() {
	var iface Iface = IfaceImpl(0)
	iface.Method()
}
