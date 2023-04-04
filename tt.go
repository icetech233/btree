package main

import "fmt"

func AA() {

	fmt.Println("go:build go1.18 +build go1.18")
}

const add = `
//go:build go1.18
// +build go1.18
`
