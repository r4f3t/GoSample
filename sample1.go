package main

import "fmt"

const programIsmi = "İdeal Kilo Programı"

func main() {
	var boy float32 = 181
	var yas float32 = 25
	var cinsiyet string = "E"
	var katsayi float32 = 0

	if cinsiyet == "E" {
		katsayi = 0.9
	} else {
		katsayi = 0.8
	}
	var ideal = (boy - 100 + yas/10) * katsayi
	fmt.Println(programIsmi)
	fmt.Printf("Boy: %v Yaş: %v Cinsiyet: %v İdeal Kilo: %v", boy, yas, cinsiyet, ideal)

	arraysSlicesMaps()

}

func arraysSlicesMaps() {
	fmt.Println()
	var dizi [2]int
	dizi[0] = 21
	dizi[1] = 25
	fmt.Println(dizi)

	slice := make([]int, 2)
	slice[1] = 1
	slice = append(slice, 2, 3, 4, 5)

	var slice2 []int

	slice2 = slice[2:5]

	fmt.Println(slice)
	fmt.Println(slice2)

	var map1 = make(map[string]string)
	map1["Ad"] = "Rafet"
	map1["Soyad"] = "Parlak"

	fmt.Println(map1)

}
