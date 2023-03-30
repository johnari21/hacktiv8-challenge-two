package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %v\n", i)
	}

	for j := 0; j < 11; j++ {

		if j == 5 {

			for idx, val := range "САШАРВО" {
				fmt.Printf("character %#U starts at byte position %d\n", val, idx)
			}
		}

		fmt.Printf("Nilai j = %v\n", j)
	}
}
