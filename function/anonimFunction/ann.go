package main

import "fmt"

// standart formada  bir  funksiya yazaq
func foo() {
	fmt.Println("when")
}
func main() {
	foo()
	ananomFunction := func() {
		fmt.Println("Geralt")
	}
	// ve funksiyaya zeng edek
	ananomFunction()

	//  indi  parametr almayan ama yerindece  zeng eden ananim funksiya  yazaq
	func() {
		fmt.Println("jenifer")
	}()
	//  anonim  funksiyani bir  parmetir verey  ve  onu yerindece  zeng edek
	func(s string) {
		fmt.Println("values", s)
	}("1636")

}
