package main

import (
	"sync"
)

var wg = sync.WaitGroup{}
var mt = sync.RWMutex{}

func main() {

	//WaitGP()
	//FiboOddEven()
	//TryIt()
	Random()
}
