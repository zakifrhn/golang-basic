package main

import "fmt"

func FiboOddEven() {
	wg.Add(2)

	cnl := make(chan int)

	var numbers = deretAngka{40}

	mt.RLock()
	go fibo(numbers, cnl)

	mt.Lock()
	go receiver(cnl)

	wg.Wait()
}

type deretAngka struct {
	num int
}

func fibo(da deretAngka, cl chan int) {
	defer func() {
		wg.Done()
	}()
	mt.RUnlock()

	//var result = 0
	for i := 2; i <= da.num; i++ {
		isPrime := true

		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			//fmt.Println(i)

			//mengirim data ke channel
			cl <- i
		}
	}
	close(cl)

}

func receiver(cl <-chan int) {
	defer func() {
		mt.Unlock()
		wg.Done()
	}()

	// for numFibo := range cl {
	// 	fmt.Println("ini adalah angka-angka fibonacci", numFibo)
	// }

	for numFibo := range cl {
		tmp := numFibo
		for i := numFibo; i <= tmp; i++ {
			// result := i
			if i%2 == 0 {
				fmt.Print("ini adalah bilangan genap:", i)
				fmt.Println("")
				// ca <- i
			} else if i%2 == 1 {
				fmt.Println("ini adalah bilangan ganjil:", i)
				fmt.Print("\n")
			}
		}
	}
}
