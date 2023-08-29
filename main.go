package main

import (
	"basic/module"
	"fmt"
	"math"
)

func main() {

	var numExam = module.CreateNumber(40)
	checkMethod(numExam)

	var resultRumus = module.CreateHitung(2, 3, 5)
	checkRumus(resultRumus)
	pembulatan(4.36)

}

func checkMethod(deret module.IfDeret) {
	//just only pick choose one and another please command it

	//fmt.Println(deret.GetFibonacci())
	//fmt.Println(deret.GetGenapGanjil())
	//fmt.Println(deret.GetPrima())
}

func checkRumus(calc module.Hitung) {
	//just only pick choose one and another please command it

	//fmt.Println(calc.Hitung2d())
	//fmt.Println(calc.Hitung3d())
}

func pembulatan(number float64) float64 {

	num := (number * 100) / 10
	fmt.Println(num)

	num1 := (math.Round(num)) / 10
	fmt.Println(num1)

	fmt.Printf("hasil desimal: %.2f", num1)

	return num
}
