package main

import "fmt"

var mapping = map[string]string{
	"ปา":   "pā",
	"ลิ":   "li",
	"ลี":   "lī",
	"ปาลิ": "pāli",
}

func main() {
	fmt.Println(mapping)
}
