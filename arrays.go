//Arrays 
package main

import "fmt"

func main() {
        var a [5]int
        fmt.Println("Empty array:",a)
        a[3] =9
        for i:=0;i<5;i++ {
            fmt.Println("After changes:",a[i])
        }
        fmt.Println("Length of Array:",len(a))
        b:= [5]float64 {0.6,15.4,96.7,12.7,0.1}
        fmt.Println(b)
        
        var c [5][5]int 
        for i:=0;i<5;i++ {
            for j:=0;j<5;j++ {
                c[i][j] =i+j
            }
        }
        fmt.Println(c)
}
