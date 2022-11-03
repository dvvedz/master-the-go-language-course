package packageone

import "fmt"

var PackageVar string = "PackageVar"

func PrintMe(mv string, bv string) {
	fmt.Println(mv, bv, PackageVar)
}
