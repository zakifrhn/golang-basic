package module

import "fmt"

type DeretNumber struct {
	num int
}

type IfDeret interface {
	GetPrima() int
	GetGenapGanjil() int
	GetFibonacci() int
}

func CreateNumber(num int) *DeretNumber {
	return &DeretNumber{
		num: num,
	}
}

func (dn DeretNumber) GetPrima() int {
	for i := 2; i <= dn.num; i++ {
		isPrime := true

		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Println(i)
		}
	}

	return dn.num
}

func (dn DeretNumber) GetGenapGanjil() int {
	for i := 0; i <= dn.num; i++ {
		if i%2 == 0 {
			fmt.Println("ini adalah bilangan genap:", i)
			fmt.Print("\n")
		} else if i%2 == 1 {
			fmt.Println("ini adalah bilangan ganjil:", i)
			fmt.Print("\n")
		}
	}
	return dn.num
}

func (dn DeretNumber) GetFibonacci() int {
	var n1, n2, n3 int = 0, 0, 1

	for i := 0; i <= dn.num; i++ {
		fmt.Printf("fibonacci-%d: %v \n", i, n2)
		n1 = n2 + n3
		n2 = n3
		n3 = n1
	}
	return dn.num
}
