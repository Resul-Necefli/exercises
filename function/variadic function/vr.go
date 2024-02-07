//  variadik  funksiyalar

package main

import "fmt"

/*
variadic funksiyalar qeyri-mehdud sayda  eyni tipde olan arqumenti qebul  ede  bilen funksiyalardir
biz  bu funksiyanin  daxilnde  elave  parametrler  teyin  edip  bura  arqumentler  oture  bileriy
variadic   funksiyalarda   variadic  parametr  mutleq   sonuncu parametr olmalidir
ve bu  funksiyalar  aldigi  deyerleri  oz  daxilene  slice type  formatinda  otrur
funksiya  daxlinde  yalniz bir  eded  variadik  parametr  ola  biler  diyerleri  adi  bir  aqument  qebul eden  parametr  olur
*/
//  tester func yardaq
func tester(r int, s ...string) (string, int) {

	var name string
	for _, v := range s {
		name += v + "\n"

	}
	return name, r

}

func templatest(username ...string) string {

	var userss string
	for _, v := range username {

		userss += v + "\n"
	}
	return userss
}

func main() {
	//  func  2
	// burada  arqumentlerin  otrulme  sirasina  diqet  etmliyik  funksiya birnci  hansi  type  parametri  gozleyirse  bizde  o parametri  oturmeliyik
	//  birnci  olan  birnci  ikinci  olan ikinci  qebul  edilir
	funcss, func2 := tester(55, "Clarc", "desc")

	fmt.Println(func2, funcss)

	// func 1
	sdd := []string{"json", "james", "haadi"}
	//biz  arqument olaraq  variadic  funksialarin  daxilne  slices  de  oture bileriy amma
	//slicesin  sonun  da  ...  splat   operatorundan istfade  ediriy bu  slicesi acir  ve  onu {1,2,3}  bu  haldan  (1,2,3) bu  hala  cevirir
	username11 := templatest(sdd...)
	fmt.Println(username11)
}
