package main

import "fmt"

type Option struct {
	A string
	B string
	C int
}

type OptionFunc func(*Option)

func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}
func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
	}
}
func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}

var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)

func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}
func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

func main() {
	x := newOption("nazha", "小王子", 10)
	fmt.Println(x)
	x = newOption2()
	fmt.Println(x)
	x = newOption2(
		WithA("沙河娜扎"),
		//WithB(""),
		WithC(250),
	)
	fmt.Println(x)
}
