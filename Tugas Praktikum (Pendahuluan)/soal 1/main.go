package main
import "fmt"

func main() {
	var n, i int
	var jam, menit, sisa, durasi int

	fmt.Scan(&n)
	durasi = 0
	for i = 1; i <= n; i++ {
		fmt.Scan(&menit)
		durasi = durasi + menit

	}
	jam = durasi / 60
	sisa = durasi % 60
	menit = sisa
	fmt.Println(jam, menit)
}

// program Durasi_Film
// Kamus
// 		n, i : integer
// 		jam, menit, sisa, durasi : integer
// Algoritma
// 		input (n)
// 		durasi <- 0
//		for i <- 1 to n do
// 			input (menit)
// 			durasi <- durasi + menit
// 		endfor
// 		jam <- durasi div 60
// 		sisa <- durasi mod 60
// 		menit <- sisa
// 		output (jam, menit)
// endprogram