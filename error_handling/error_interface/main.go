package main

import (
	"errors"
	"fmt"
)

// error  interface isleme yeri eslinde funkisyanin return olunarken   orada artiq interface implament olunur deye bilerik
// bu sadece  errors.New() arxa terfidi deye bileriki oradki erors paketi ve  new  metodu arxa terfde bu funksay veya  interfaceye dayanaraq
// isleyir
type errror interface {
	Error() string
}

type errortring struct {
	s string
}

func (e *errortring) Error() string {
	return e.s
}

/*
hersey burada baslayir deye bilerik errors paketinin New metodunu cagrirq ve o islemeye basiliyr icne verdiymiz stringi errorString

		struct s string  hisesine gedir    ve sonra  return olduqda   orda  error ki var   o eslinde  interfacedir yeni biz eslinde hem error hemde
		interface return etmis oluruq deye bilerem  ve  interface  tetiklenir ve islemeye basliyir   ve  icnde olan Error()string metodunu qaytarir
	     bu metodda  bizim struct  field mizi  return edirki buda bizim  sexsi yartigmiz xeta mesajidir
*/
func New(text string) error {
	return &errortring{text}
}

func main() {

	fmt.Println(sqrt(5))
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("kok alma funksiyasi mendi eded qebul etmir")
	} else {
		return f, nil
	}
}
