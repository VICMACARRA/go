package main

import "fmt"

func main() {
	// 	edad:=0
	// kk:=&edad
	// 	fmt.Println(&kk)
	// 	//


	type persona struct {
		nombre string
		edad   int
	}

	type padres struct {
		madre persona
		padre persona
		hijos []persona
	}

	pedro := persona{"Pedro", 20}
	victor := persona{"Victor", 30}
	// fmt.Println(pedro)
	// fmt.Println(victor)
//change age funcion
	funcion := func(p *persona) {
		p.edad = p.edad + 100
	}
	 funcion(&pedro)
	 funcion(&victor)

	
	// fmt.Println(pedro)
	// fmt.Println(victor)

	FamiliaJuarez := padres{
		madre: persona{"Maria", 40},
		padre: persona{"Juan Jose", 50},
		hijos: []persona{pedro, victor, {"Sandra", 160}},
	}
	fmt.Println(FamiliaJuarez)
	type personaService interface {
		getName() string
		setName(n string)
		getAge() int
		setAge(a int)
	}


var n bool = true
fmt.Printf("%v, %T\n", n,n)
	

	

}