package main

import (
	"fmt"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/transform"
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
		fmt.Printf("Category: %v, Len: %v\n", dataB.Category, dataB.Len)
		// Parsing JSON datablock for each record

		if dataB.Category == 48 {
			for _, record := range dataB.Records {
				catModel := new(transform.Cat048Model)
				transform.WriteModel(catModel, record.Items)
				catJson, _ := transform.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		} else if dataB.Category == 34 {
			for _, record := range dataB.Records {
				catModel := new(transform.Cat034Model)
				catJson, _ := transform.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		} else if dataB.Category == 30 {
			for _, record := range dataB.Records {
				catModel := new(transform.Cat030STRModel)
				catJson, _ := transform.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		} else if dataB.Category == 255 {
			for _, record := range dataB.Records {
				catModel := new(transform.Cat255STRModel)
				catJson, _ := transform.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		}
	}
}
