package rest

import "fmt"

func DefPath() (string, string) {
	var name string
	fmt.Print("File name from this directory: ")
	fmt.Scanln(&name)
	namj := name[:len(name)-3] + "json"
	return name, namj
}
