package main

import (
	"fmt"
)

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
	//	faktoryel()
	//	switchCase()
	getProductsWitForeach()

	fmt.Println(calculator(3, 5))
	odd, even := seperateOddEven(1, 3, 5, 76)
	fmt.Println(odd, even)
}

//variadic fonksiyon string args gibi eleman sayısı bağımsız dizi alır
func seperateOddEven(a ...int) (odd []int, even []int) {

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			even = append(even, a[i])
		} else {
			odd = append(odd, a[i])
		}
	}

	return
}

func calculator(a, b int) (int, float32) {
	carpim := a * b
	bolum := float32(a) / float32(b)
	return carpim, bolum
}

func faktoryel() {
	var faktoryel, j = 0, 1
	fmt.Println("Kaç Faktöryel=")
	fmt.Scan(&faktoryel)
	var result = 1
	for i := 1; i <= faktoryel; i++ {
		result *= i
	}

	//while döngüsü
	for j <= faktoryel {
		j++
	}

	fmt.Printf("Faktöryel :%v", result)
}

type Product struct {
	Name string
	Code string
}

func getProductsWitForeach() {
	var slice = []Product{}
	slice = append(slice, Product{Name: "Ürün1", Code: "Code1"})
	slice = append(slice, Product{Name: "Ürün2", Code: "Code2"})
	slice = append(slice, Product{Name: "Ürün3", Code: "Code3"})
	for i, e := range slice {
		fmt.Println(i, e)
	}

}

func switchCase() {
	fmt.Println("Gün Giriniz=")
	var gun = 0
	fmt.Scan(&gun)
	switch gun {
	case 1:
		fmt.Println("pazartesi")
	case 2:
		fmt.Println("salı")
	default:
		fmt.Println("Geçersiz")
	}
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
