package main

import "fmt"

func main() {


	parser := NewIniParser()

	
	fmt.Println(parser.IniSections["Colors"][0])

	
}