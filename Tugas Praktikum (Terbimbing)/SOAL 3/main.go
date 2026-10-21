package main

import "fmt"

func main() {
	var n, i, hasil int

	fmt.Scan(&n)

	hasil = 1
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}

	fmt.Println(hasil)
}



// program Faktorial
// // Kamus
// //     n, i, hasil : integer

// // Algoritma
// //     input (n)
// //     hasil <- 1
// //     for i <- 1 to n do
// //         hasil <- hasil * i
// //     endfor
// //     output (hasil)
// // endprogram