package main

import "fmt"

var mapping = map[string]string{
	"ปา":   "pā",
	"ลิ":   "li",
	"ลี":   "lī",
	"ปาลิ": "pāli",
	"สาธุ": "sādhu",
}

func main() {
	fmt.Println(mapping)
}
