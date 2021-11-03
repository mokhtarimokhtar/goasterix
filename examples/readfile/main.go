package main

import (
	"fmt"
	"github.com/mokhtarimokhtar/goasterix"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("../data/sample.ast")
	if err != nil {
		log.Fatalln(err)
	}

	w := new(goasterix.WrapperDataBlock) // wrapper of asterix datablock, it contains one or more datablocks
	_, err = w.Decode(data)              // decode method one or more datablocks
	if err != nil {
		fmt.Println("ERROR Wrapper: ", err)
	}

	for _, dataB := range w.DataBlocks {
		// dataBlock contains one datablock = CAT + LEN + RECORD(S)
		fmt.Printf("Category: %v, Len: %v\n", dataB.Category, dataB.Len)
		for i, records := range dataB.String() {
			// records contains one or more records = N * items
			fmt.Printf("+ Record n°%v\n", i+1)
			for _, record := range records {
				fmt.Println("-", record) // it displays its items in Hexadecimal
			}
		}
	}
}
