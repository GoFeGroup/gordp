package core

import "fmt"

func ShowLine(line []byte) {
	var i int
	for i = 0; i < len(line); i++ {
		fmt.Printf(" %02X", line[i])
	}
	for ; i < 16; i++ {
		fmt.Printf("   ")
	}
	fmt.Printf("\t")
	for i = 0; i < len(line); i++ {
		if line[i] >= ' ' && line[i] < 128 {
			fmt.Printf("%c", line[i])
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("\n")
}

func ShowHex(data []byte) {
	for {
		if len(data) <= 16 {
			break
		}
		ShowLine(data[:16])
		data = data[16:]
	}
	ShowLine(data)
}
