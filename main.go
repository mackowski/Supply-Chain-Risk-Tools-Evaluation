package main

import (
	"fmt"
	obfuscated "github.com/mackowski/ObfuscatedMethodPackage"
)

func main() {
	// Example usage of the package
	fmt.Println("Using ObfuscatedMethodPackage")
	// Add your code here to use the package's functionality
	err := obfuscated.ConnectToEncodedAddress("bG9jYWxob3N0OjQ0NDQ=")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
