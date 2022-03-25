package main

import (
	"os"
	"txt2json/rest"
)

func erc(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	name, namej := rest.DefPath()

	inp, err := os.ReadFile(name)
	erc(err)

	ot, err := os.Create(namej)
	erc(err)

	d := ""

	str := make([]string, 0, 99999)

	for i := 0; i < len(inp); i++ {
		if inp[i] == 10 {
			str = append(str, d)
			d = ""
		} else {
			d += string(inp[i])
		}
	}

	ot.WriteString("{\n")

	for i := 0; i < len(str)-1; i++ {
		ot.WriteString(rest.String2Json(str[i]) + ",\n")
	}

	ot.WriteString(rest.String2Json(str[len(str)-1]) + "\n}")

	print("Done\n")
}
