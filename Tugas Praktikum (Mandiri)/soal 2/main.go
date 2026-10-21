package main

import "fmt"

func main() {
    var n, k, i int

    fmt.Scan(&n)

    for k = 1; k <= n; k++ {
        i = n + 1 - k
        fmt.Println("Request", k, "diproses dengan prioritas", i)
    }
}


// Program TelkomCloud
// Kamus
//     n, k, i : integer
// Algoritma
//     input(n)
//     for k <- 1 to n do
//         i <- n + 1 - k
//         output("Request ", k, " diproses dengan prioritas ", i)
//     endfor
// EndProgram