// //  structlar

// package main

// import (
// 	"fmt"
// )

// type man struct {
// 	name string
// 	age  int
// 	job  string
// }
// type anonim struct {
// 	man
// 	keys bool
// 	name string
// }

// //	type   acar  sozu  sonra  struct  adi  qeyd  olunur  qloblada olan  struct
// //
// // muxtelif  funksiyalardan ona  zeng  etmek  ve  istifade  etmey olar
// type resul struct { //----------  bu  struct  bize  gosterirki  onlarin zero  deyeri  nil  kimi  teyin  olunur
// }

// // bir  struct yaradip  onun  daxilinde  muxtelif  tiplerden parametrler  teyin  ede  bilerik
// type domein struct {
// 	firs_Name string
// 	last_name string
// 	age       int
// }

// func main() {

// 	// yartigmiz structlara  func main()  daxilnde   deyerler  arqumentler  teyin  ederken  biz onu  her  hansisa  bir  deyiskende  tanimlamaliyiq
// 	//  ve  o   deyisken vasitesi  ile   cap   edrik  bu  dilin  spesifikasiyasinda  beledir bele  yardilip  sebebide  odurki  biz
// 	//  bir  struct yardip  ondan  muxtelif  deyiskenlere  declaration  edip   isdedyimiz  qeder  o  sturctdan  isdfade  ede  biliriy
// 	dd := resul{}
// 	fmt.Println(dd)

// 	//  burada  temel  struc  deyerlerin  verilmesine  aid  numnelerdir
// 	p := domein{
// 		firs_Name: "resul",
// 		last_name: "necefli",
// 		age:       24,
// 	}

// 	p1 := domein{
// 		firs_Name: "orxan",
// 		last_name: "neceflli",
// 		age:       30,
// 	}
// 	//  sintaksisi asagdaki  kimdir
// 	//  ic  ice  structlar
// 	//  burada  esa  nuans  odurki

// 	killer := anonim{
// 		man: man{
// 			name: "James  Bond", //  burdada  name  var
// 			age:  25,
// 			job:  "killer",
// 		},
// 		keys: true,
// 		name: "Resul  Necefli", // burada  name  var

// 	}

// 	killer.age = 45

// 	//  biz  normal  killer.name  yazsaq  proqram  man.name    nezere  almayacaq  uzde olani goturub yazdiracaq  buna  gorede  bele  seflere
// 	// yol  vermemek  ucun  biz    proqrama  hansi  name  isdedyimizi  deqiq  demeliyik     killer.man.name  kimi
// 	fmt.Printf("type %T  deyer %v\n", killer, killer)
// 	fmt.Println(killer.man.name, killer.name, killer.man.age)

// 	//  anonim  struc    bu  tip   structlar  yalniz  birdefe  istfade  edilir ve  func daxilnde  yardila  biler   eyer  bize  mueyyen  bir  yerlerde
// 	//  bir   sade  struct  lazimdrsa  ve  birdefe  istfade  edilecekse  anaonim  struct  en yaxsidir

// 	rr := struct {
// 		name string
// 		age  int
// 		job  bool
// 	}{
// 		name: "Rivali Geralt",
// 		age:  554,
// 		job:  true,
// 	}

// 	fmt.Println(p)
// 	fmt.Println(p1)
// 	fmt.Println("anonim  struc  deyerimz", rr)

// 	//  burada    car struc   mueyyen bir  deyiskende  tanimlayaraq deyerler teyin  ediriy
// 	//  numunede  gorunduyu  kimi  biz  car  struct vasitesi  ile engine  struc  icndeki  start() metoduna  zeng  ede  biliriy
// 	//
// 	avto := car{
// 		engine: engine{
// 			engene_now: "V12 DUBLEX"},
// 		name: "mercedes benz",
// 	}

// 	avto.Start()
// 	fmt.Println(avto)

// 	//  yartdigimiz  type  de  bir  deyisken  yardaq  ve  ona  42  int  deyerini  verek  artiq  onu  ekrana  yazdira  bileriy
// 	//  untmayaq  ki       struc, map, slices ,arry  ve  s   bunlarda  bir  type  dir
// 	var a result = 42
// 	fmt.Println(a)

// }

// //	biz Composition  nece  yarndigina  baxriq  burada    bu  termin muxtelif  sade  tiplerin  bir-birinin   uzerine  yiglaraq
// //	murekkeb  tipler   yarda  bilmesdir
// //	meslen  biz  asagda   mator  ve  masin structu  yartdiq  ve    mator  structunu  masin  structunun  icne  yerlsdirdik
// //
// // ve  bizim  masin   car{}  struct  artiq  engine  struct  butun metodlarini  ve  ya  onun  parametrlerini   ehtiva ede  bilir

// type engine struct {
// 	engene_now string
// }

// func (e *engine) Start() {
// 	fmt.Println("mator mersedes")
// }

// type car struct {
// 	engine
// 	name string
// }

// type result int

// //  biz  go  da oz  type  yarda  bileriy  bunun  ucun  biz  type  acar  sozunu istfade  etmeliyik  global  seviyede  yeni
// //  artiq  menim  result  type  var  ve  o  yalniz  int  deyerler  depolaya  bilir
// //  eyer  projectmiz  ucun  daha  oxunaqli  bir  kod  isdeyrsinizse  oznuz  ucun  type  leri ozmuz  yarda  bilerik
// //  tebiki   Go  oz  type  var  ve  yetri  qeder  her bir  deyeri  ozunde  saxlayacaq  qeder  novleri  movcuddur
// //  buradan biz   gonun    oop ,  nesne   yonumlu  oldugu  aglimza  gelir   bu suala  go.blog   cavab  tapa  bileriy

/*
package main

import (
	"fmt"
)

type engine struct {
	motor bool
}

type car struct {
	engine
	model string
	rengi string
	sport bool
	turbo bool
}

func main() {

	p2 := struct {
		firs_Name  string
		friendname map[string]int
		fovDrinks  []string
	}{
		firs_Name: "Orxan",
		friendname: map[string]int{
			"Resul Necefli": 24,
			"Nebi  necefli": 33,
		},
		keys: true,
		name: "Resul  Necefli", // burada  name  var

		fovDrinks: []string{
			` arax , viski , dublex ,  kanyak ,  su `,
		},
	}
	fmt.Println(p2)

	for i, v := range p2.friendname {
		fmt.Println(p2.firs_Name, i, v)
		for _, c := range p2.fovDrinks {
			fmt.Println(p2.firs_Name, "sevilen ickiler", c)

		}
	}

	c1 := car{
		engine: engine{true},
		model:  "mercedes",
		rengi:  "qara",
		sport:  true,
		turbo:  true,
	}

	c2 := car{
		engine: engine{true},
		model:  "toyoto",
		rengi:  "qirmizi",
		sport:  true,
	}

	fmt.Println(c1.engine, c1.model, c1.motor, c1.sport, c1.turbo, c1.rengi)
	fmt.Println(c1)
	fmt.Println(c2.engine, c2.model, c2.motor, c2.sport, c2.turbo, c2.rengi)
	fmt.Println(c1.engine, c1.model, c1.motor, c1.sport, c1.turbo, c1.rengi)

type engine struct {
	engene_now string
}

*/

package main

import (
	"fmt"
)

type office struct {
	teacher string
}
type result struct {
	firs_Name string
	last_name string
	age       int
	job       bool
	salary    float32
	manager   bool
	office
}

func main() {

	c := result{"resul necefli", " orxan necefli", 45, true, 525.5, true, office{"orxan"}}

	d := result{
		firs_Name: "james",
		last_name: "bond",
		age:       52,
		job:       true,
		salary:    455.2,
		manager:   true,
		office:    office{teacher: "ORxan mellim"},
	}
	fmt.Println(d)
	fmt.Println(c)

	var ds result
	ds.firs_Name = "resul  necefli"
	ds.last_name = "orxan  necefli"
	ds.job = true

	fmt.Println(ds)

	e := new(result)

	e = &result{
		firs_Name: "Resul",
		last_name: "necefli",
		age:       25,
		job:       true,
		salary:    55.3,
		manager:   false,
		office:    office{teacher: "Orxan"},
	}

	fmt.Println(e)
	fmt.Printf("deyerleri  yazdiraq  %+v\n", *&e.age)

}
