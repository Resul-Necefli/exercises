package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	/*burada vurgulamali oldgumuz novbeti sey biz shevlerimizi muyen bir .txt uzantili fayila yonelde bileriy ve bu sayede sehvimiz terminala yox
	  o fayla yoneldilecekdir  bununcunun biz   log.SetOutput()metodundadn   istfade edirik bu metod sayesinde bizsehv mesajini fayla yazmis olurq
	  bu metod (w io.writer)  interfeysini  istfade edir
	*/

	fn, err := os.Create("file.txt")

	if err != nil {
		log.Println(err)
	} else {
		log.SetOutput(fn)
	}

}
func main() {

	_, err := os.Open("filess.txt")
	if err != nil {
		log.Println("xeta deyeri error:", err)
	}

	telskop, err := os.OpenFile("tg.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Println(err)
	}
	log.SetOutput(telskop)
	//burada olmayan fayliacaraq biz sehv mesajni aliriq ve bunu muxtelif metodlardan istfade ederek ekrana gonedrmis oluruq
	// bu metodlarin ferqlerine baxaq
	_, err = os.Open("sdd.txt")
	if err != nil {

		log.Fatal(err)   // proqrami Exit 1 cixisi ile sonlandirir ve bize sehvin tarixi baresinde melumat verir
		log.Fatalln(err) // eynidir
		log.Println(err) // eynidir
		log.Panic(err)   // proqrami Exit 1 cixisi ile sonlandirir ve bize  sehvin oldugu folder haqda ve setir haqda melumat otrur
		panic(err)       // sadece  olaraq caxnasma yardir ve  shevin location u  ve bas verdiyi setir haqda melumat verir  hemcinin panic sehvlerini
		// recover() metodu ile tuta bilrdikki hansiki orda proqram sehv olsda davam edirdi
		eror := fmt.Errorf(" error happend : %e", err.Error())
		fmt.Println(eror)

	}
}
