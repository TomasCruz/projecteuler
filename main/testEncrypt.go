package main

import (
	"crypto/cipher"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

func main() {
	var inputFlag string
	flag.StringVar(&inputFlag, "kind", "e", "a string var")
	flag.Parse()

	joinedArgs := strings.Join(flag.Args(), " ")

	var gcm cipher.AEAD
	var err error
	var result string

	if gcm, err = projecteuler.GenerateGcm(); err != nil {
		log.Fatalln(err)
	}

	switch inputFlag {
	case "e":
		if result, err = projecteuler.EncryptString(joinedArgs, gcm); err != nil {
			log.Fatalln(err)
		}
	default:
		if result, err = projecteuler.DecryptString(joinedArgs, gcm); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println(result)
}
