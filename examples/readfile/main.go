package main

import (
	"fmt"
	"goasterix"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("../data/sample.ast")
	if err != nil {
		log.Fatalln(err)
	}

	w := new(goasterix.WrapperDataBlock) // wrapper of asterix datablock, it contains one or more datablocks
	_, err = w.Decode(data) // decode method one or more datablocks

	for _, dataB := range w.DataBlocks {
		fmt.Printf("Category: %v, Len: %v\n", dataB.Category, dataB.Len)
		for i, recs := range dataB.String() {
			fmt.Printf("+ Record n°%v\n", i+1)
			for _, rec := range recs {
				fmt.Println("-",rec)
			}
		}
	}
}