
// xeritler map[]

package main

import (

	"fmt"

)

func main() {

	// deyiskene   map  tipiniin   verilmesi
	var ssd map[string]int
	fmt.Println(ssd) //  bu  sekilde  olur  ve  bos   map[]  yeni  nil  deyerini  alir   yeni  bos  olaraq  tanimlanir  4
	//arxa  terfde  ram  de  hecbir  yer  ayrilmir
	hacker := map[string]int{
		"Resul": 42,
		"ORxan": 03,
	}
	fmt.Println(hacker)
	ssd = hacker
	fmt.Println()
	fmt.Println(ssd)
	//------------------------------------------------------
	//  indi  bu deyiskene  deyer  elave  edek
	ssd["resul Necefli"] = 45 //  sintaksis  xeatsi verir  cunki  eslinde   ortada  bir xerite  yoxdu  yeni  var  deyil  sadece  bu  deyisken her  hansisa  bir  xerite  tipli  deyiskeni  ozunde  saxliya biler

	fmt.Println(ssd)

	//  make() fonksuyonu  ile  map[]  yartmaq
	fg := make(map[string]bool) //   arxa  terfde   bir  xerite  var  parmetirleri var   amma  arqumetleri  yoxdur
	fg["order by"] = true       //  yartigmiz  xeriteye  deyeri elave  edey burda
	fmt.Println(fg["order by"])

	//-----------------

	//  xeritlerde  silme  islemi   delete()  funksiyasi
	fg["yaddas"] = false
	delete(fg, "order by")
	fmt.Println(fg)

	//  map  type  referans  vermesi  map[]  tipleride  deyerini  yox  referansini  paylasir silce  kimi

	laptop := map[string]int{
		"asus":    1500,
		"dell":    1000,
		"toshiba": 800,
		"hp":      1200,
		"casper":  700,
	}

	fmt.Printf("%#v\n", laptop)

	var notebook map[string]int

	notebook = laptop
	fmt.Printf("%#v\n", notebook)
	notebook["asus"] = 555
	fmt.Printf("%#v\n", notebook)
	fmt.Printf("%#v\n", laptop)
	//  ve  burada  biz  gorduyumuz  kimi   map[]  referansini  paylasir  notebook daki  deyer  deyisen kimi  laptop daki  deyerde  deyisir

	notebook = make(map[string]int) //burada  biz  deyiskenin  xeritesini  yartix  untmayaqki  iki  xeritede  yeni  notebook  ve laptop
	//  indi  artiq  bu  anan  etibaren  ayri  yari  xeritelere  referans  edirler     yeni   notebook  bi  xerite  taype  idi  ve  o  bir  xerite  deyildi
	//  referans  ede  bileceyi  bir  yer  yoxuydu  ama  make()  funksiyasi  ile  bunu yartix  artiq  o  oz  xeritesine  referans  verir
	//  bundan  sonrasinda  biz onu  yazdirdigimizda o  xerite  bos  olacaq map[]  sonra  deyer  teyin  etdik ve  ekranda  CPU :22   deyerini  yazdirdi
	//  ama  biz  golang  isleme  pirinsipne  uygun  olaraq   yuxardan  asgi  yeni main(){...}  scoplarinin  arasindan  sohbet    gedir   burada  yeniden
	//
	//notebook = laptop  bele  yazsak  o  andaki vezyetde  bunlar  ikisde  eyni   xeriteye  refrans  vermis  olacaq   ve  eyni  deyerlere  sahip  olacaqlar

	notebook["CPU"] = 22

	notebook = laptop
	fmt.Println(notebook)
	fmt.Println(laptop)

	//  her  hansi  bir  deyerin  xeritenin icnde  olup  olmadigini  yoxlayaq
	//  v --  acar   k--  deyer   temsil   edir
	v, ok := fg["helper"]

	if ok {
		fmt.Println("deyer vaar  ", v, ok)
	} else {
		fmt.Println("deyer yoxdur")
	}

	//  if else
	if v, ok := fg["order by"]; !ok {

		fmt.Println("bele  bir  qeydiyyat  movcud  deyil")
	} else {
		fmt.Println("axtaxrdiginiz qeydiyyat ", v, ok)
	}

}

  
//--------------------------------------------------------------------------

// MAP  exercises

//---------------------------------------------------------------------------

package main

import "fmt"

func main() {

	/*
		key  value
		`bond_james``shaken, not stirred`, `martinis`, `fast cars`
		`moneypenny_jenny``intelligence`, `lc`, `computer science`
		`no_dr``cats`, `ice cream`, `sunsets``  */

	film := map[string][]string{

		"bond_james":       {"haken", "not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "intelligence", "computer science"},
		"no_drcats":        {"ice  kerm", "sunset"},
	}
	fmt.Println(film)
	delete(film, "bond_james")

	for i, v := range film {
		fmt.Printf("key:%#v values:%#v\n", i, v)
	}

	for i, v := range film {

		fmt.Println(i)
		for s, d := range v {

			fmt.Println(s, d)
		}

	}

	/*

			key value
		`fleming_ian` `steaks`, `cigars`, `espionage`
	*/

	loopkey := make(map[string][]string)

	loopkey["fleming_ian"] = []string{"steaks", "cigars", "espionage"}
	for i, v := range loopkey {
		fmt.Println(i, " bura  daxildir asagdaki  filmler")

		for _, d := range v {

			fmt.Println(d)
		}
	}

	
	x := []string{"in", "my", "younger", "and", "more", "vulnerable", "years", "my", "father", "gave", "me", "some",
		"advice", "that", "i’ve", "been", "turning", "over", "in", "my", "mind", "ever", "since.", "whenever",
		"you", "feel", "like", "criticizing", "anyone,", "he", "told", "me,", "just", "remember", "that", "all",
		"the", "people", "in", "this", "world", "haven’t", "had", "the", "advantages", "that", "you’ve",
		"had.", "he", "didn’t", "say", "any", "more,", "but", "we’ve", "always", "been", "unusually",
		"communicative", "in", "a", "reserved", "way,", "and", "i", "understood", "that", "he", "meant",
		"a", "great", "deal", "more", "than", "that.", "in", "consequence,", "i’m", "inclined", "to",
		"reserve", "all", "judgements,", "a", "habit", "that", "has", "opened", "up", "many", "curious",
		"natures", "to", "me", "and", "also", "made", "me", "the", "victim", "of", "not", "a", "few",
		"veteran", "bores.", "the", "abnormal", "mind", "is", "quick", "to", "detect", "and", "attach",
		"itself", "to", "this", "quality", "when", "it", "appears", "in", "a", "normal", "person,", "and", "so",
		"it", "came", "about", "that", "in", "college", "i", "was", "unjustly", "accused", "of", "being", "a",
		"politician,", "because", "i", "was", "privy", "to", "the", "secret", "griefs", "of", "wild,", "unknown",
		"men.", "most", "of", "the", "confidences", "were", "unsought—frequently", "i", "have",
		"feigned", "sleep,", "preoccupation,", "or", "a", "hostile", "levity", "when", "i", "realized", "by",
		"some", "unmistakable", "sign", "that", "an", "intimate", "revelation", "was", "quivering", "on",
		"the", "horizon;", "for", "the", "intimate", "revelations", "of", "young", "men,", "or", "at", "least",
		"the", "terms", "in", "which", "they", "express", "them,", "are", "usually", "plagiaristic", "and",
		"marred", "by", "obvious", "suppressions.", "reserving", "judgements", "is", "a", "matter", "of",
		"infinite", "hope.", "i", "am", "still", "a", "little", "afraid", "of", "missing", "something", "if", "i",
		"forget", "that,", "as", "my", "father", "snobbishly", "suggested,", "and", "i", "snobbishly",
		"repeat,", "a", "sense", "of", "the", "fundamental", "decencies", "is", "parcelled", "out",
		"unequally", "at", "birth.", "and,", "after", "boasting", "this", "way", "of", "my", "tolerance,", "i",
		"come", "to", "the", "admission", "that", "it", "has", "a", "limit.", "conduct", "may", "be",
		"founded", "on", "the", "hard", "rock", "or", "the", "wet", "marshes,", "but", "after", "a", "certain",
		"point", "i", "don’t", "care", "what", "it’s", "founded", "on.", "when", "i", "came", "back", "from",
		"the", "east", "last", "autumn", "i", "felt", "that", "i", "wanted", "the", "world", "to", "be", "in",
		"uniform", "and", "at", "a", "sort", "of", "moral", "attention", "forever;", "i", "wanted", "no",
		"more", "riotous", "excursions", "with", "privileged", "glimpses", "into", "the", "human", "heart.",
		"only", "gatsby,", "the", "man", "who", "gives", "his", "name", "to", "this", "book,", "was",
		"exempt", "from", "my", "reaction—gatsby,", "who", "represented", "everything", "for", "which",
		"i", "have", "an", "unaffected", "scorn.", "if", "personality", "is", "an", "unbroken", "series", "of",
		"successful", "gestures,", "then", "there", "was", "something", "gorgeous", "about", "him,",
		"some", "heightened", "sensitivity", "to", "the", "promises", "of", "life,", "as", "if", "he", "were",
		"related", "to", "one", "of", "those", "intricate", "machines", "that", "register", "earthquakes",
		"ten", "thousand", "miles", "away.", "this", "responsiveness", "had", "nothing", "to", "do", "with",
		"that", "flabby", "impressionability", "which", "is", "dignified", "under", "the", "name", "of", "the",
		"creative", "temperament—it", "was", "an", "extraordinary", "gift", "for", "hope,", "a", "romantic",
		"readiness", "such", "as", "i", "have", "never", "found", "in", "any", "other", "person", "and",
		"which", "it", "is", "not", "likely", "i", "shall", "ever", "find", "again.", "no—gatsby", "turned", "out",
		"all", "right", "at", "the", "end;", "it", "is", "what", "preyed", "on", "gatsby,", "what", "foul", "dust",
		"floated", "in", "the", "wake", "of", "his", "dreams", "that", "temporarily", "closed", "out", "my",
		"interest", "in", "the", "abortive", "sorrows", "and", "short-winded", "elations", "of", "men.", "my",
		"family", "have", "been", "prominent,", "well-to-do", "people", "in", "this", "middle", "western",
		"city", "for", "three", "generations.", "the", "carraways", "are", "something", "of", "a", "clan,",
		"and", "we", "have", "a", "tradition", "that", "we’re", "descended", "from", "the", "dukes", "of",
		"buccleuch,", "but", "the", "actual", "founder", "of", "my", "line", "was", "my", "grandfather’s",
		"brother,", "who", "came", "here", "in", "fifty-one,", "sent", "a", "substitute", "to", "the", "civil",
		"war,", "and", "started", "the", "wholesale", "hardware", "business", "that", "my", "father",
		"carries", "on", "today.", "i", "never", "saw", "this", "great-uncle,", "but", "i’m", "supposed", "to",
		"look", "like", "him—with", "special", "reference", "to", "the", "rather", "hard-boiled", "painting",
		"that", "hangs", "in", "father’s", "office.", "i", "graduated", "from", "new", "haven", "in", "1915,",
		"just", "a", "quarter", "of", "a", "century", "after", "my", "father,", "and", "a", "little", "later", "i",
		"participated", "in", "that", "delayed", "teutonic", "migration", "known", "as", "the", "great",
		"war.", "i", "enjoyed", "the", "counter-raid", "so", "thoroughly", "that", "i", "came", "back",
		"restless.", "instead", "of", "being", "the", "warm", "centre", "of", "the", "world,", "the", "middle",
		"west", "now", "seemed", "like", "the", "ragged", "edge", "of", "the", "universe—so", "i",
		"decided", "to", "go", "east", "and", "learn", "the", "bond", "business.", "everybody", "i", "knew",
		"was", "in", "the", "bond", "business,", "so", "i", "supposed", "it", "could", "support", "one",
		"more", "single", "man.", "all", "my", "aunts", "and", "uncles", "talked", "it", "over", "as", "if",
		"they", "were", "choosing", "a", "prep", "school", "for", "me,", "and", "finally", "said,",
		"why—ye-es,", "with", "very", "grave,", "hesitant", "faces.", "father", "agreed", "to", "finance",
		"me", "for", "a", "year,", "and", "after", "various", "delays", "i", "came", "east,", "permanently,",
		"i", "thought,", "in", "the", "spring", "of", "twenty-two.", "the", "practical", "thing", "was", "to",
		"find", "rooms", "in", "the", "city,", "but", "it", "was", "a", "warm", "season,", "and", "i", "had",
		"just", "left", "a", "country", "of", "wide", "lawns", "and", "friendly", "trees,", "so", "when", "a",
		"young", "man", "at", "the", "office", "suggested", "that", "we", "take", "a", "house", "together",
		"in", "a", "commuting", "town,", "it", "sounded", "like", "a", "great", "idea.", "he", "found", "the",
		"house,", "a", "weather-beaten", "cardboard", "bungalow", "at", "eighty", "a", "month,", "but",
		"at", "the", "last", "minute", "the", "firm", "ordered", "him", "to", "washington,", "and", "i", "went",
		"out", "to", "the", "country", "alone.", "i", "had", "a", "dog—at", "least", "i", "had", "him", "for", "a",
		"few", "days", "until", "he", "ran", "away—and", "an", "old", "dodge", "and", "a", "finnish",
		"woman,", "who", "made", "my", "bed", "and", "cooked", "breakfast", "and", "muttered",
		"finnish", "wisdom", "to", "herself", "over", "the", "electric", "stove.", "it", "was", "lonely", "for",
		"a", "day", "or", "so", "until", "one", "morning", "some", "man,", "more", "recently", "arrived",
		"than", "i,", "stopped", "me", "on", "the", "road.", "how", "do", "you", "get", "to", "west", "egg",
		"village?", "he", "asked", "helplessly.", "i", "told", "him.", "and", "as", "i", "walked", "on", "i",
		"was", "lonely", "no", "longer.", "i", "was", "a", "guide,", "a", "pathfinder,", "an", "original",
		"settler.", "he", "had", "casually", "conferred", "on", "me", "the", "freedom", "of", "the",
		"neighbourhood.", "and", "so", "with", "the", "sunshine", "and", "the", "great", "bursts", "of",
		"leaves", "growing", "on", "the", "trees,", "just", "as", "things", "grow", "in", "fast", "movies,", "i",
		"had", "that", "familiar", "conviction", "that", "life", "was", "beginning", "over", "again", "with",
		"the", "summer.", "there", "was", "so", "much", "to", "read,", "for", "one", "thing,", "and", "so",
		"much", "fine", "health", "to", "be", "pulled", "down", "out", "of", "the", "young", "breath-giving",
		"air.", "i", "bought", "a", "dozen", "volumes", "on", "banking", "and", "credit", "and",
		"investment", "securities,", "and", "they", "stood", "on", "my", "shelf", "in", "red"}


		//  yuxardaki  deyerleri    dongu  ile  bir  bir  map  xertimize  elave  ediry ve  o  deyerlerin  adi  acar    sayi  ise  value  deyer  kimi
		//   qeyd  olunur  artix xeritede
	ma := map[string]int{}

	for _, v := range x {

		ma[v]++
	}

	//  burda  biz  artix  xeritemizin  icndeki  deyerleri  ve  acarlari ekrana  yazdiriq buda  bize  her sozu ve  deyeri  olaraqda  o  sozden 
	//nece  eded  varsa  qarsisinda  yazilacaqdir

	for k, v := range ma {
		fmt.Println(k, v)
	}

}
