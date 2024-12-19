package main

import "fmt"

func main() {
	dontknow := 2344
	i := 1
	for know := 1; dontknow > 0; know *= 2 {
		fmt.Printf("Знают - %d, не знают - %d. День %d.\n", know, dontknow, i)
		i++
		dontknow -= know * 2
	}
}

/*func main() {
	sum := 0
	for А := 0; А < 100; А++ {
		sum = 0
		for Б := 0; Б < 100; Б++ {
			for В := 0; В < 100; В++ {
				for Г := 0; Г < 100; Г++ {
					sum += А*100 + Б*10 + В
					sum += Б*100 + В*10 + Г
					sum += В*100 + Г*10 + А
					sum += Г*100 + А*10 + Б
					if sum == 2442 {
						fmt.Println(А, "+", Б, "+", В, "+", Г, "=", А+Б+В+Г)
					}
				}
			}
		}
	}
	fmt.Println("Done!")
}*/
