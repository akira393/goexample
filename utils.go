package main

import "fmt"

var (
	Version  string
	Revison  string
	BuildTag string
)

func PrintMeta() {
	fmt.Println(Version, Revison, BuildTag)
}
