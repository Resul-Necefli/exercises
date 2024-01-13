package defaultn

import (
	"fmt"
)

type Factories struct {
	name string
}

type Company struct {
	Name      string
	Factories []*Factories
}
type Monday struct {
	last_name string
	firs_name string
	age       int
	job       bool
	salary    float32
	Company   *Company
}

func Red() {
	fmt.Println("i am  like  red")
}

func Compy() {

	ssh := []string{
		"İri Qabaritli Transformatorlar Zavodu",
		"ATEF Elektromaş Zavodu",
		"ATEF Transformator Zavodu",
		"Yüksək Gərginlikli Avadanlıqlar Zavodu",
		"Metal Konstruksiya Zavodu",
		"Tökmə Məmulatları Zavodu",
		"Kabel Məhsulları Zavodu",
		"Dəqiq Mexaniki Emal Zavodu",
		"Bərk İzolyasiya Zavodu",
	}

	desr := Monday{
		last_name: "jhon",
		firs_name: "clack",
		age:       24,
		job:       true,
		salary:    2.200,
		Company: &Company{
			Name:      "RRRR",
			Factories: CreateFactories(ssh),
		},
	}
	fmt.Println(desr)

}

func CreateFactories(names []string) []*Factories {
	var factories []*Factories

	for _, name := range names {
		factory := &Factories{name: name}
		factories = append(factories, factory)
		fmt.Println(factory.name)
	}

	return factories
}
