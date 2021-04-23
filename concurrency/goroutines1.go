package main

import "time"

func main() {

	//eş zamanlılığı sağlar
	go f()
	time.Sleep(2000 * time.Millisecond)
}

func f() {
	for i := byte('a'); i <= byte('z'); i++ {
		println(string(i))
		time.Sleep(10 * time.Millisecond)
	}
}
