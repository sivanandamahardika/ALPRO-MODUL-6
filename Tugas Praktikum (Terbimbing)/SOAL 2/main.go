package main

import "fmt"

func main() {
	var a, b, i, hasil int

	fmt.Scan(&a, &b)

	hasil = 1
	for i = 1; i <= b; i++ {
		hasil = hasil * a
	}

	fmt.Println(hasil)
}



// program Pemangkatan
// Kamus
//     a, b, i, hasil : integer

// Algoritma
//     input (a, b)
//     hasil <- 1
//     for i <- 1 to b do
//         hasil <- hasil * a
//     endfor
//     output (hasil)
// endprogram