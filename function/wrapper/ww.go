package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//  burada  bir txt fayli  yardiriq  ve  bunu  yaradarken  biz  bir eror  birde    faylin ozunu  deyiskende  saxlayiriq
	dd, err := os.Create("option.txt")
	if err != nil {
		log.Fatal(err)
	}

	//  fayli  sonda baglayiriq  ki  lazim  olan  butun emeliyatlar  yerine  yetrile bilsin  buda  defer()   funksiyasi  yerine yetirir
	defer dd.Close()
	fmt.Println("fayil yaradildi")
	//  indi  ise  biz  oz  yartdigmiz  faylin icne  bir seyler  elave  edek  burda elave  etmek  ucun  istfade etdiyimiz metod writestring()
	//  bize  bir  eror ve  birde  []byte  dilimi  qaytarir  bunun ucunde  byten dilmini saxlmriq yeni tutmuruq her  hansisa deyskenden
	//  cunki  ne  gonderdiymizi  bilirik  bunu  etmek  menasiz  olardi
	_, err = dd.WriteString("hello gopher! my name is James Bond")
	if err != nil {
		log.Fatalf("fayla melumatin oturlmesinde  nese  xeta bas verdi %v\n", err)
		return
	}

	fmt.Println("")
	fmt.Println("fayla melumatlar ugurla otruldu")
	//  burada  faylimizi  funksiyamiza  arqument olaraq  gonderek
	reader, err := readFile("option.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reader)         //  bu  cixsda  biz  byte cixsini  yeni  gore  bilirik yeni     102,105  kimi
	fmt.Println(string(reader)) //  o byte  cixsini  stringe oxuya  bileceyimiz dile  bu  formada  cevririk  standart library  de  string()
	//  bunu  nece  edir  gosterlip

}

//	yaratdigimiz  bu  func  her  hansisa  bir  faylin icini  oxuyur  ve  bize  bir  eror  ve  birde  byte dilmi qaytarir
//
// ve  bizde  bunu  muxtelif  deyiskenlerde  saxlyiriq  ve  return  ediriy
func readFile(fn string) ([]byte, error) {

	xb, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("fayil oxunarken xeta bas verdi%v\n", err) // sehvi yoxlayaq

	}
	return xb, nil //  byte xb    eror  nil  qaytaraq
}



//--------------------------------------------------------------------------------------------------------------------------------------


package main

import (
	"fmt"
	"time"
)

// OriginalFunction, orginal  olan bir fonksiyon
func OriginalFunction() {
	fmt.Println("Orijinal fonksiyon çalişti.")
}

// Logger, başka bir fonksiyonu sargılayan ve loglama islerini elave eden bir "wrapper" fonksiyonu
func Logger(f func()) {
	fmt.Println("Fonksiyon çağrildi:", time.Now())
	f() // Orijinal fonksiyonu çağıriq
	 fmt.Println("Fonksiyon çağrildiktan sonra:", time.Now())
}

func main() {
	// Logger fonksiyonunu  istifade  ederek  OriginalFunction  funksiyasini sargilayiriq
	Logger(OriginalFunction)
}


/*
Bu nə edir? : Sarma funksiyaları funksiyaya əlavə funksionallıq əlavə etmək üçün istifadə olunur. 
Bu funksionallıq orijinal funksiyanın işini dəyişdirmədən və ya təsir etmədən sonrakı əməliyyatları yerinə yetirməyə imkan verir.

Niyə istifadə olunur? Sarmalayıcı funksiyalar müxtəlif funksiyaları birləşdirərək, 
həmçinin təkrar istifadə edilə bilən və modul kod strukturu yaratmaqla mürəkkəb əməliyyatları sadələşdirmək üçün istifadə olunur.

Necə istifadə etməli? : Funksiya başqa funksiya tərəfindən çağırıldıqda, 
çağırılan funksiyanın ətrafına "sarğı" funksiyası yerləşdirilə bilər. 
Bu sarma funksiyası çağırılan funksiyanı çağırmadan əvvəl və ya sonra müəyyən əməliyyatları yerinə yetirə bilər.

Nümunə İstifadəsi : Məsələn, siz giriş əməliyyatı əlavə etmək üçün "sarğı" funksiyasından istifadə edə bilərsiniz. 
Orijinal funksiyanı çağırmadan əvvəl və sonra giriş əməliyyatlarını yerinə yetirən funksiya orijinal funksiyanın ətrafına sarıla bilər. 
Beləliklə, hər hansı bir funksiyanı çağırarkən qeyd avtomatik olaraq həyata keçirilə bilər.

*/