// panic recover anlayisi

/*
	panic bildiyimiz kimi icnde oldugu funksiyada caxnasmalar yaradir ve  proqrami dayandiraraq  bize bir xeta verir

xetananin location  ve hansi setirde oldugunda bildirir   bes biz panic()  eden funksayani  xeta oldugunu eslinde bilrmisik kimi
ve ya bele bir xeta gozlentimizin  oldugunu var sayaraq bu xetani bu formada gostererek proqramin ana main() funk icndeki
diger funksiyalar ve ya para metrleri isletmesini bacara bilerikmi?  beli  recover() funksiyasida tam bu isi gorur
onun istfadesi asgadaki kimdir
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("proqram basladi")
	hello()
	fmt.Println("proqram bitdi")

}

func hello() {
	// recover defer ile isletmeliyik  niye ?    cunki  kod compiler olunarken bilirikki butun kod bir seyfeye yiglir elbilki bir a4
	//vereqine yiglirmis kimi dusnek bu zaman proqram artix bilirki panic()  var  ve  eyni zamanda recover() panic  tutur   amma  defer oldugunan
	// recover() qarsisinda  o qalir axira ve diger funksiyalar isleyir  ve sondada recover() panic() oldugun ozu bildrir ve funksia dayanir
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("xeta ele alindi", err)
		}
	}()

	testingo()
}

func testingo() {
	_, err := os.Open("sourche.txt")
	if err != nil {
		panic(err)
	}

}
