package main

import "fmt"

func main() {
    var n, i, h int

    fmt.Scan(&n)

    for i = 1; i <= n; i++ {
        h = 2*i - 1
        fmt.Println(h)
    }
}



// Program BilanganGanjil
// Kamus
//     n, i, h : integer
// Algoritma
//     input(n)
//     for i <- 1 to n do
//         h <- 2 * i - 1
//         output(h)
//     endfor
// EndProgram