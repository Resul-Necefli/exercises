package main

import "fmt"

func ttx(a, b int) func() int {
	c := a + b
	return func() int {
		c++
		return c
	}

}

func context() func() int {

	x := 0
	return func() int {
		x++
		return x
	}

}

func main() {

	resizster := ttx(44, 55)
	for i := 0; i < 5; i++ {

		fmt.Println(resizster())
	}

	f := context()
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("values:", f())
	fmt.Println("------------------------------------------")
	// burada  gordymuz kimi  herbir  deyer ferqli deyiskenlerde saxlandigi  kimi  ferqli  adreslerde  tutulut
	// bunun ucun g  ile  zengi funksiyaya etdikde o yeniden baslayir
	g := context()
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())
	fmt.Println("values:", g())

}
