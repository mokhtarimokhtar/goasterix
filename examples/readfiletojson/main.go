package main

import (
	"fmt"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/model"
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
		// Parsing JSON datablock for each record

		if dataB.Category == 48 {
			for _, record := range dataB.Records {
				catModel := new(model.Cat048Model)
				catJson, _ := model.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		} else if dataB.Category == 34 {
			for _, record := range dataB.Records {
				catModel := new(model.Cat034Model)
				catJson, _ := model.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		} else if dataB.Category == 30 {
			for _, record := range dataB.Records {
				catModel := new(model.Cat030STRModel)
				catJson, _ := model.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		} else if dataB.Category == 255 {
			for _, record := range dataB.Records {
				catModel := new(model.Cat255STRModel)
				catJson, _ := model.WriteModelJSON(catModel, record.Items)
				fmt.Println(string(catJson))
			}
		}
	}
}