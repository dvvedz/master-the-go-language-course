package main

import "myapp/packageone"

var myVar string = "myVar"

func main() {
	var blockVar string = "blockVar"

	// Variables
	// declare a package level variable for the main
	// package named myVar

	// decalre a block variable for the main function
	// called blockVar

	// declare a package level variable for the main function
	// package named PackageVar

	// create an exported function in packageone called PrintMe
	// in the main funtion, print out the values of myVar,
	// blockVar, PackageVar on one line using PrintMe function in packageone
	packageone.PrintMe(myVar, blockVar)
}
