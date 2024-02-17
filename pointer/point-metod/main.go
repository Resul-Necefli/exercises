package main

import "fmt"

type andPoint struct {
	first string
}

/*
burada metodumuz deyer sematikasi ile  irelleyir yeni burada deyerlerin bir  kopyasini alir ve oz isinigorur sonra o deyerler
silinir  bu o demekdirki  adam  icndeki structun kopyasin cixarip gotrur ozune deymir ona  mudaxile edede bilmez  ve  o deyer ler yigin (stack)
deyilen yerde saxlanilir hansiki  metod isin bitrenen sora hersey silinir
*/

func (adam andPoint) Myname() string {
	return adam.first
}

// func siya ise dusende burda  bas veren proses beledir  pointerName(sdk *andPoint,s string) string{}  yeni metod eslinde xususi bi novu olan
// func sayilir
/*
daxilde  sdk.first = s   eslinde     (*sdk).first =*s  olmalidir  go  bunu  bizim yerimize edir onun ucun bele rahat sekilde
gosterici qebul edicsine hem deyeri hemde adresi isdedymiz kimi otrsekde qebul edir  ozu lazim olan formata  cevirir cunki

*/
func (sdk *andPoint) pointerName(s string) string {
	sdk.first = s
	return sdk.first
}
func main() {
	h := andPoint{first: "Anda rio"}
	// deyer qebul eden metodumuz isledi ve esil structa temas ede bilmediyi ucun onun sadece bir kopyasini alip o kopya uzerinde emliyatllarini
	// yerine yetirecek  struct  esli deysmir
	fmt.Println("Mynames metodunan gelen deyer -->", h.Myname())
	// bu metod pointerName  gosterci  tipinde  oldugu ucun struct birbasa  ozune referans  edir bu sebebden dolayi oradaki deyreleri
	// deyisdre biler  ve bunuda burada  edirik asgida  structmuzun first bolmesine yeni deyerini gonderirik
	fmt.Println(h.pointerName("Resul necefli"))
	//  ve  burada birdaha  structun name  ekrana yazdirdigda gorurki artiq o deysilip cunki deyerini yox refransini paylasir
	fmt.Println("Mynames metodunan gelen deyer -->", h.Myname())

}
