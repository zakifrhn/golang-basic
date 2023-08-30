package main

import "fmt"

func TryIt() {
	wg.Add(2)
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)

	go sum(a[0:len(a)/2], chn)
	go sum(a[(len(a)/2):8], chn)

	//menerima (receiver) hasil dari channel
	result1 := <-chn
	result2 := <-chn

	total := result1 + result2
	fmt.Println(total)

	wg.Wait()
}

func sum(d []int, chn chan int) {

	defer func() {
		wg.Done()
	}()

	//membuat penampung nilai utk penjumlahan nanti
	var tmp = 0
	for i := 0; i < len(d); i++ {

		//menjumlahkan semua angka di dalam slice array
		tmp += d[i]
	}

	//mengirim (sender) hasil (tmp) ke channel
	chn <- tmp
}
