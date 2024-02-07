/*interface bir*/
/*
package main

import "fmt"

type add interface {
	Speak()
}

func interfaceFunc(a add) {
	a.Speak()

}

type result struct {
	first_Name string
	last_name  string
}

func (r result) Speak() {
	fmt.Println("senin adin")
	fmt.Println(r.first_Name, r.last_name)
}

type book struct {
	title string
}

func (b book) Speak() {
	fmt.Println("giris deyerimiz", b.title)
}

func main() {

	rr := result{
		first_Name: "james",
		last_name:  "clarc",
	}

	bb := book{
		title: "THE WITCHER",
	}

	interfaceFunc(bb)
	interfaceFunc(rr)

}

*/

//--------------------------------------------------------------------------------------------------------------------------------------------------------

/*
package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}


func writeLog(s fmt.Stringer) {
	log.Println(s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var c count = 42

	writeLog(b)
	writeLog(c)

}

*/

/*`writeLog funksiyası və fmt.Stringer interfeysi:

writeLogFunksiya fmt.Stringeröz interfeysini həyata keçirən bir növü parametr kimi qəbul edir.
Go dilində interfeyslər verilmiş metodun bütün növləri üçün ümumiləşdirilmiş sxemi müəyyən edir.
 fmt.Stringerİnterfeys yalnız bir metoddan ibarətdir: String() string. Bu üsul dəyəri insan tərəfindən oxuna bilən sətirə çevirmək üçün istifadə olunur.
writeLog funksiyasının parametri:

writeLogFunksiya fmt.Stringeronun interfeysini həyata keçirən istənilən növü parametr kimi qəbul edir.
Yəni bu funksiya String()öz metodunu təyin edən və sətri qaytaran istənilən növü qəbul edə bilər.
Funksiya parametr kimi qəbul etdiyi növün String()metodunu çağırır və bu üsulla qaytarılan sətirdən istifadə edir.
Parametri necə əldə etmək olar:

writeLogFunksiya fmt.Stringeröz interfeysini həyata keçirən bir növü parametr kimi qəbul edir.
Məsələn, kimi bookvə ya tiplər bu funksiyaya ötürülə bilər, çünki onlar onun interfeysini həyata keçirirlər.countfmt.Stringer
Funksiya String()parametr kimi qəbul etdiyi növün metodunu çağıraraq sətri əldə edir.
 Bu, fmt.Stringeronun interfeysinin bir hissəsi olduğu üçün mümkündür. Bu üsul növün xüsusi simli təsvirini qaytarmalıdır.
Giriş üçün sətirin yazılması:

Funksiya String()parametr kimi qəbul etdiyi növün metodunu çağıraraq sətri əldə etdikdən sonra bu sətri jurnala yazır.
Jurnal log.Println()öz funksiyası vasitəsilə həyata keçirilir.
Bu, fmt.Stringeronun interfeysini həyata keçirən istənilən növ simli təsvirləri qeyd etmək üçün bir yoldur.
Nəticədə, writeLogfunksiya fmt.Stringerparametr kimi onun interfeysini həyata keçirən istənilən növü götürür.
Bu tip metodu çağıraraq String()sətri əldə edir və bu sətri jurnala yazır.
Bu, müxtəlif növ sətir təsvirlərini asanlıqla daxil etmək üçün istifadə edilə bilən bir üsuldur.





`*/
//---------------------------------------------------------------------------------------------------------------------------------------------------------

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

//  indi  oz novumuzu ve  onun metodunu  yardiriq  ve  budefe  yartigimiz  bufer tipli deyiskenden  isdifade  ederk  buradaki
//  deyeri o  bufere oturulmeni  gore bilerik

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	// bir  fayil acaq ve buradan qayidan deyeri ve  xetani  deyiskende  tutaq

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	//  acilan fayli sonda  baglamaliyiq ki   aciq qalmasin
	defer f.Close()
	// indi  ise  bu yartigimiz faylin icne  bir  deyer elave edek bunun  icin
	//write()  metodunu  istifade  edirik  bu metod her  hansi  buferin ve ya  fayilin icine deyerler  oturmeye  imkan  verir
	//  cunki  os paketininde  write  deye bir  metodu  var  ve   eyni metoda malikdirse  ve  eyni deyer  geri  dondurse biz
	// burada  bu metodu olan butun paketlere  bunu  uygulaya  bilirik
	//  ve  bu metod  func (*os.File).Write(b []byte) (n int, err error)    bizden  bir  byte  dilimi isdeyir  yeni  deyer  bu dilmin icnde olsun isdeyir
	//  ve  bize  byte  sayisini  ve  eyer  bunu  saya  bilmyeceyi  teqdirde  bir  erorr sef  dondurur
	d := []byte("Hello gopher")

	_, err = f.Write(d)
	if err != nil {
		log.Fatal(err)

	}
	//  artiq  acdigimiz  faylin icine  deyerimiz  oturduk
	//  burada  dd  adinda  bir  bufer  yardiriq  ve  bu dd  write()  metodu var  artiq  demeli bizim yuxarida yartigimiz
	//   oz metodumuzki var  writeOut()  person struct  metodu  o metodun icnde  parametr teyin etmisiy  w io.Writer  bu bir interface dir
	//  ve  bu interface   sadce  bir  write()  metodunu  taniyir bizim dd  deyiskenimzdede  bu metod oldugu  ucun demeli biz  ora giris ede  bilerik
	//  yeni  w parmetri bizim dd deyiskenini ozune  arqumet olaraq  qebul ede  biler
	var dd bytes.Buffer
	p := person{
		first: "James Bond",
	}

	p.writeOut(&dd)
	fmt.Println(p.writeOut(&dd)) // --  cixis  olaraq 10 <nil>  bunu  gorurk  bu  metod buferdeki byte  sayini ve  erorr  qaytarir sehv olmadigi
	//  ucun  nil yazsini goruruk
	//-----------------------------------------------------------------------------------------------------------------------------------------------------

	//   biz  indi  bir  bufer yardacagiq  ve  budefe  bu  buferin icne  otrdyumuz  deyeri oz  novumuzde istfade  ede bileceyik
	//  cunki   herki  novde eyni metoda  sahib olmus  olacaq

	b := bytes.NewBufferString("hello  go ")

	fmt.Println(b.String())

	fmt.Println(b.WriteString(" new write"))
	fmt.Println(b.String())

	//  buferi bu  formatda  yardilir  bufer  yadsda olan muveqqeti bir yerdir

	/*
			Məlumatların ötürülməsi sürətinin artırılması : Buferlər məlumatların daxil edilməsini və çıxışını daha səmərəli edə bilər.
		Məsələn, fayldan verilənləri oxuyarkən və ya fayla məlumat yazarkən verilənlər buferlər vasitəsilə daha böyük bloklarda oxuna və ya yazıla bilər.
		Bu, diskə girişləri azaltmaqla məlumat ötürmə sürətini artıra bilər.

		Sıxılma və ya Şifrələmə : Sıxılma və ya şifrələmə kimi əməliyyatlar zamanı məlumatların buferləşdirilməsi əməliyyatları daha səmərəli edə bilər.
		Böyük miqdarda məlumatı sıxışdırmaq və ya şifrələmək üçün buferlərdən istifadə etməklə əməliyyatlar daha sürətli həyata keçirilə bilər.

		Məlumatların emalı : Məlumatların işlənməsi tətbiqlərində məlumatları bufer etməklə əməliyyatların daha nizamlı və effektiv olmasını təmin edə bilərik.
		Məsələn, bir sıra verilənləri emal etmək üçün biz bu məlumatları əvvəlcə bufer etməklə daha asan emal edə bilərik.

		Məlumat Kommunikasiyası : Şəbəkə rabitəsi və ya verilənlər bazası əməliyyatlarında biz buferlər vasitəsilə məlumat axınına daha yaxşı nəzarət edə bilərik.
		Xüsusilə şəbəkə kommunikasiyalarında buferlər daxil olan məlumatları saxlamaq və daha böyük blok kimi emal etmək üçün istifadə edilə bilər.
		Bu, şəbəkə trafikini azalda və ünsiyyəti daha səmərəli edə bilər.

		Bir sözlə, buferlər məlumat axınlarını təşkil etmək, məlumatları müvəqqəti saxlamaq və məlumat ötürülməsini optimallaşdırmaq üçün istifadə olunur.
		 Bu, çox vaxt məlumatların işlənməsi, məlumatların ötürülməsi və faylların işlənməsi kimi ssenarilərdə mühüm rol oynayır.*/

}
