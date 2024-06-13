package main

import (
	Tcef "github.com/Edw590TryCatch-go"
	"log"
)

func main() {
	Tcef.Tcef{
		Try: func() {
			var test *bool = nil
			*test = true
		},
		Catch: func(e Tcef.Exception) {
			log.Printf("Caught %v\n", e)
		},
		Else: func() {
			log.Println("If it didn't panic...") // Which it will, so this won't execute
		},
		Finally: func() {
			log.Println("In any case...")
		},
	}.Do()

	log.Println("--------------------------")

	// Panic ignored here. Just try and forget if it didn't work.
	Tcef.Tcef{
		Try: func() {
			panic(nil)
		},
	}.Do()

	log.Println("--------------------------")

	Tcef.Tcef{
		Try: func() {
			var test bool
			test = true
			if test {}
		},
		Else: func() {
			log.Println("If it didn't panic...")
		},
		Finally: func() {
			log.Println("In any case...")
		},
	}.Do()
}
