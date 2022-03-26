package rest

import "fmt"

func DefPath() (string, string) {
	
	var (
		name string
		named string
	)

	fmt.Print("File name from this directory: ")
	fmt.Scanln(&name)
	
	for i := 0; i < len(name); i++{
		if string(name[i]) == "."{
			break
		} else {
			named += string(name[i])
		}
	}

	return name, named + ".json"
}
