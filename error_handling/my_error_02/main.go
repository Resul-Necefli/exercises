package main

import (
	"errors"
	"fmt"
	"math"
)

// oz structumuzu yaradaq
type sqrtError struct {
	lat  float64
	long float64
	err  error
}

// yaratdigimiz structura Error() metodunu elave edirik ki o  error interfaceni  implement ede bilsin  ve bizde oz xeta mesajimizi
// oz novbesinde bura elave edirik
func (s *sqrtError) Error() string {

	return fmt.Sprintf("math error : %+v\t%+v \t %v", s.lat, s.long, s.err)
}

// biz error   istfade edeceyimiz  function yaradiriq   ve bu funksiya daxilinde isdafade edeceyiy
func sqrt(s float64) (float64, error) {

	if s < 0 {
		return 0, &sqrtError{long: 0, lat: 0, err: errors.New("negatif nummber")}
	}
	//
	return math.Sqrt(s), nil
}

func main() {
	// function qayidan deyeri
	result, err := sqrt(-4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	f, err := sqrt(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

	p1 := person{name: " ", last_name: " "}

	value, err := NameString(p1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

}

// oz sturktumuz yaridirq eyni sekilde
type person struct {
	name      string
	last_name string
	errr      error
}

// Error() herbir nov   error interfaysini islede biler bu bizim sehflerimiz haqda dahada yaxsi melumat vere bilmeyimizin bi yoludur sadece
func (p *person) Error() string {

	return fmt.Sprintf("name:   last_name:  %v  %v  %v\n ", p.last_name, p.name, p.errr)

}

func NameString(a person) (string, error) {
	if len(a.name) < 10 {
		return " ", &person{name: "", last_name: "", errr: errors.New("ad  soyad bolmesi bos qala bilmez ve herbiri 6 karkter olmalidir")}
	} else {
		return a.name, nil
	}
}
