package main

import "fmt"

func fibbonaci(n int) {
	currentNum := 0
	isNFibbonaci := false
	f1 := 0
	f2 := 1
	for i := 0; i < n-1; i++ {
		currentNum = f1 + f2
		f1 = f2
		f2 = currentNum
		if n == currentNum {
			isNFibbonaci = true
		}
	}
	if isNFibbonaci {
		fmt.Println(n)
		for i := 0; i < 9; i++ {
			currentNum = f1 + f2
			f1 = f2
			f2 = currentNum
			fmt.Println(currentNum)
		}
	} else {
		for i := 0; i < 10; i++ {
			currentNum = f1 + f2
			f1 = f2
			f2 = currentNum
			fmt.Println(currentNum)
		}
	}

}

func main() {
	num := 0
	fmt.Scan(&num)

	fibbonaci(num)
}
