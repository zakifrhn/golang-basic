package main

import "fmt"

func Random() {
	wg.Add(2)

	//contoh buffer channel
	data := make(chan int, 20)

	mt.RLock()
	go pengirim(data)

	mt.Lock()
	go penerima(data)

	wg.Wait()
}

func pengirim(msg chan int) {
	// defer wg.Done()

	mt.RUnlock()

	defer func() {
		//mt.RUnlock()
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("ini angka", i)

		msg <- i
	}

	close(msg)
}

func penerima(msg chan int) {

	defer wg.Done()
	mt.Unlock()

	for receiverData := range msg {
		fmt.Println("angka yang diterima adalah", receiverData)
	}
}
