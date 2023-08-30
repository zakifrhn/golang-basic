package main

import (
	"fmt"
)

func WaitGP() {

	//numChan := make(chan int)

	wg.Add(3)

	var deret = DeretNumber{40}
	mt.Lock()
	go deret.GetGenapGanjil()

	mt.Lock()
	go deret.GetFibonacci()

	mt.Lock()
	go deret.GetPrima()

	// hasil1 := <-numChan
	// hasil2 := <-numChan
	// hasil3 := <-numChan
	// fmt.Println(hasil1)

	// fmt.Println(hasil2)

	// fmt.Println(hasil3)
	wg.Wait()
}

type DeretNumber struct {
	num int
}

func (dn DeretNumber) GetPrima() (ch chan int) {
	defer func() {
		mt.Unlock()
		wg.Done()
	}()
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

	return ch
}

func (dn DeretNumber) GetGenapGanjil() (ch chan int) {
	defer func() {
		mt.Unlock()
		wg.Done()
	}()
	for i := 1; i <= dn.num; i++ {
		if i%2 == 0 {
			fmt.Println("ini adalah bilangan genap:", i)
			fmt.Print("\n")
		} else if i%2 == 1 {
			fmt.Println("ini adalah bilangan ganjil:", i)
			fmt.Print("\n")
		}
	}
	return ch
}

func (dn DeretNumber) GetFibonacci() (ch chan int) {

	defer func() {
		mt.Unlock()
		wg.Done()
	}()

	var n1, n2, n3 int = 0, 0, 1

	for i := 0; i <= dn.num; i++ {
		fmt.Printf("fibonacci-%d: %v \n", i, n2)
		n1 = n2 + n3
		n2 = n3
		n3 = n1
	}
	return ch
}
