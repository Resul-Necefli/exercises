package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// file yaradaq  ve sehvlerimizi yolxayaq
	files, err := os.Create("Resul.txt")
	if err != nil {
		panic(err) // sehv varsa panic() prqorami sonlandir  deyirem
	}
	defer files.Close() // ve faylimi sonda baglmagi untmuramki kompuyterim yavlamasin cunki bu bi yox bir nece faylda ola bilerdi acdigim
	//  burda   bufere yazmaq ucun   bir nov yazi metodunu ele aliriq deye bilerem
	write := bufio.NewWriter(files)
	// daha sonra  o yartigimiz metodla  bufere bir yazi yaziriq  bufer bizim ram yadsda saxlanilan  yerimizdi muveqqetidir yeni adeten daha
	// suretli ve verimli olsun deye bezi isleri bura  sonra file gonderirik
	_, err = write.WriteString("salma menim adim resuldur ")
	if err != nil {
		log.Panic(err)
	}
	//  Flish()  metoduna zeng ediremki sendeki yazini yaz  fayla oda  elave edir  fayla
	err = write.Flush()
	if err != nil {
		panic(err)
	}
	// burada  biz yartdigimiz  ve icne yazi yazdigimiz fayli acip baxmaq isdeyirik
	sddd, err := os.Open("Resul.txt")
	if err != nil {
		log.Fatal(err) //  her  defesinde sehvleri yoxlmagimiz lazmdirki sehv olduqda mesaj versin ve sehvin hardan qaynaqlandigini bilek
		//  cunki bir proqramda  heden cox sehv qaytaran metodlar ola biler
	}
	//  ve acdigim  fayli baglmagi untmuruq performans  cehetden buna diqett edilmelidir
	defer sddd.Close()
	//  burda   io.ReadALL() funksiyasi istfade ederek acdigimiz faylin icndeki her nedise onu oxuya bilerik
	menimfaylim, err := io.ReadAll(sddd)
	if err != nil {
		panic(err)
	}
	// io.ReadALL()   funksiyasi   bize oxunan fayli byte  formatinde qaytardigi  ucun biz  onu string() metodu ile stringe cevririki
	// oxuya bildiyimiz formmatda olsun
	fmt.Println(string(menimfaylim))

}
