package main

import "fmt"

func main() {
    var a[10] int
    var n, c, d, swap int
    fmt.Scan(&n)
    for c=0; c<n; c++ {
        fmt.Scan(&a[c])
    }
    for c=0; c<n-1; c++ {
        for d=0; d<n-c-1; d++ {
            if a[d] > a[d+1] {
                swap = a[d]
                a[d] = a[d+1]
                a[d+1] = swap
            }
        }
    }
    
    fmt.Println("bubble sort in asc order")
    for c=0; c<n; c++ {
        fmt.Println("%d\n", a[c])
    }
}