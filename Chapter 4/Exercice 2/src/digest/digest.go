// The digest program computes the hash (SHA256, SHA384 or SHA512) of stdin input
package main

import (
	"flag"
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	)
	
var algo = flag.String("d", "sha256", "Digest to be used")

func main() {
	flag.Parse()
	var input string
	
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Printf("Invalid input\n")
	} else {
		switch *algo {
			case "sha256":
				digest := sha256.Sum256(([]byte)(input))
				fmt.Printf("SHA256: %x\n", digest)	
			case "sha384":
				digest := sha512.Sum384(([]byte)(input))
				fmt.Printf("SHA384: %x\n", digest)	
			case "sha512":
				digest := sha512.Sum512(([]byte)(input))
				fmt.Printf("SHA512: %x\n", digest)	
			default:
				fmt.Printf("Invalid digest algorithm (%s)\n", *algo)
		}
	}
}