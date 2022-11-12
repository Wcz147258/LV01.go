package main

import "fmt"

func main() {
	l := 0
	b := 0
	e := 0
	var a [1000]int
	fmt.Scan(&l)
	fmt.Scan(&b)
	fmt.Scan(&e)
	for i := 0; i < l; i++ {
		fmt.Scan(&a[i])
	}
	for i := b; i < e; i++ {
		for j := i; j < e; j++ {
			tmp := 0
			if a[j] > a[j+1] {
				tmp = a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}

		}
	}
	for i := 0; i < l; i++ {
		fmt.Print(a[i], " ")
	}

}
