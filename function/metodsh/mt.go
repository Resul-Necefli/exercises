/*
metodlar  structlara  yartgimiz  her  hansi  type   aid  olar  biler
yazlis   formasi  funksiyalara  benzeyir  ve  onlar  muyyen  tipe aid olurlar

func (a user) tenni(){}    burada   tenni() metoddur  ve o user tipi uzerinde  insa  edilip


*/

package main

import (
	"errors"
	"fmt"
)

// user structunu tanimlayaq
type user struct {
	name     string
	username string
	age      int
	password string
	email    string
}

// user structu bir novdur  ve  bu novun ozune mexsus bir  metodunu yaradaq  bu metod  bu struct uzerine insa edilir hal hazirda
//  yeni biz bu metod vastesi ile  o structun uzerinde her hansisa bir emeliyyatlari yerine yetirme gucune sahbik
/*
metodlar var olma  sebebi  gonu daha  cox hemde  oop(obyek yonumlu olmasina getrip cixara bilir)  meselen biz  bir struct yardiriq ve
ondan isdediyimiz qeder  istfade  edirik  ve isdediymiz  qederde  metodlarini yarda  bilirik  meselen  sifreni yoxlayan metod  ve ya  name yenileyen metod

*/
// burada  bir metod yartdiq ve bize  bir xeta qaytarmasini ona  emr  etdik
func (u user) passwordValid() error {

	//  burada hemin struct icnde olan o structa gelen  sifrenin  6  herfli olup olmamasini  istfadecinin  duzgun parol teyin
	// edip etmemesini yoxlayriq
	if len(u.password) < 6 {
		return errors.New("sifrenizde  en az 6 eded reqem olmaldir")
	}

	return nil

}

func main() {
	//structumza deyerleri otururuk
	p1 := user{
		name:     "James",
		username: "Bond",
		age:      42,
		password: "12346",
		email:    "Jamesbond42@gmail.com",
	}
	// ve metodumuzu cagririq  tebiki struct vasitesi ile cagrilir cunki  o struct uzerine insa edilmis bir metoddur
	err := p1.passwordValid()
	// metod bir sehv qaytarir ve bizde  onu   saxlayiriq
	if err != nil {
		fmt.Println(err, p1.password)
	} else {
		fmt.Println("sifreniz dogrulandi")
	}

	fmt.Println(p1.name, p1.username, p1.age, p1.password, p1.email)

}
