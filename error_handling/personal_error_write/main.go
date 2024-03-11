package main

import (
	"fmt"
)

// oz structimzi yardaraq hemcinin oz xusisi sehvlerimzde yarda bilerik  ve structumza isdeiymiz qeder field elave ede bilmeyimzide
// bizim xetalari daha asan bir sekilde idare de bilmeyimze serait yaradir
type calculator struct {
	asd float64
	fff float64
	err error
}

// yartigimiz struc Error metoduna sahib olmaldiki gonun error interfacesin emplament ede bilsin
func (c *calculator) Error() string {
	return fmt.Sprintf("kok funksiyasiyasina  gonderilen  deyerler \t %v \t %v \t %v\n", c.asd, c.fff, c.err)
}

func main() {
	a, err := sqrt(-44)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(a)
	}
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		nem := fmt.Errorf("kok funksiyasi menfi ededleri qebul etmir")
		return 0, &calculator{asd: a, fff: a, err: nem}
	} else {
		return a, nil
	}
}
