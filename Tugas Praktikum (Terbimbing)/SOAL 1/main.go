package main

import "fmt"

func main() {
	var a, b, i int

	fmt.Scan(&a, &b)

	for i = a; i <= b; i++ {
		if i != a {
			fmt.Print(" ")
		}
		fmt.Print(i)
	}

	fmt.Println()
}

// program TampilkanBarisBilangan
// Kamus
//     a, b, i : integer

// Algoritma
//     input (a, b)
//     for i <- a to b do
//         if i != a then
//             output (" ")
//         endif
//         output (i)
//     endfor
//     output
// endprogram