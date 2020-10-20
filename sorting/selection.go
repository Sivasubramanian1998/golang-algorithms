package main

import "fmt"

func main() {
    var a[10] int
    var n, i, d, pos, t int
    fmt.Scan(&n)
    for i=0; i<n; i++ {
        fmt.Scan(&a[i])
    }
    for i=0; i < n-1; i++ {
        pos = i;
    for d = i+1; d<n; d++ {
        if a[pos] > a[d]{
            pos = d;
        }
    }
    if pos != i {
        t = a[i]
        a[i] = a[pos]
        a[pos] = t
    }
}

fmt.Println("selection sort in asc order")
for i=0; i<n; i++ {
    fmt.Println("%d\n", a[i]);
}
}
