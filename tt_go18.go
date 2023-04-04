//go:build !go1.18
// +build !go1.18

package tree

import "fmt"

func AA() {

	fmt.Println("go:build !go1.18 +build !go1.18")
}
