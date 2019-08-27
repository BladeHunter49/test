package main

import (
	"log"
)

func main() {
	no := []int{1, 2}
	log.Print("Print no ", no, "\n")
	log.Println("Println no", no)
	log.Printf("Prinf no[%d,%d]", no[0], no[1])
	// log.Panicln("Panicln no", no)
	log.Fatalln("Fatalln no", no)
}
