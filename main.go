package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: makemask [rom file, rom file,...]")
	}
	bootdata, err := Asset("data/boot.6102")
	if err != nil {
		panic(err)
	}
	for arg := range flag.Args() {
		name := flag.Arg(arg)
		fmt.Printf("Masking %s.\n", name)
		f, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.WriteAt(bootdata, 0x40)
	}
}
