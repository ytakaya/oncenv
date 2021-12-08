package main

import "os"

func init() {
	os.Getenv("A")
}

type a string

func (a a) a() {
	os.Getenv("A") // want "os\\.Getenv\\(\\) is used in inappropriate place"
}

func main() {
	os.Getenv("A") // want "os\\.Getenv\\(\\) is used in inappropriate place"
}
