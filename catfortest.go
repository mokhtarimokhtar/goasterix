package goasterix

import "github.com/mokhtarimokhtar/goasterix/uap"

var Cat4ForTest = []Item{
	&Fixed{
		Base: Base{
			FRN:         1,
			DataItem:    "I026/010",
			Description: "Data Source Identifier",
			Type:        uap.Fixed,
		},
		Size: 2,
		SubItems: []SubItem{
			&SubItemFromTo{
				Name: "SAC",
				Type: uap.FromTo,
				From: 16,
				To:   9,
			},
			&SubItemFromTo{
				Name: "SIC",
				Type: uap.FromTo,
				From: 8,
				To:   1,
			},
		},
	},
	&Fixed{
		Base: Base{
			FRN:         2,
			DataItem:    "I026/055",
			Description: "Mode-1 Code",
			Type:        uap.Fixed,
		},
		Size: 1,
		SubItems: []SubItem{
			&SubItemBit{
				Name: "V",
				Type: uap.Bit,
				Pos:  8,
			},
			&SubItemBit{
				Name: "G",
				Type: uap.Bit,
				Pos:  7,
			},
			&SubItemBit{
				Name: "L",
				Type: uap.Bit,
				Pos:  6,
			},
			&SubItemFromTo{
				Name: "CODE",
				Type: uap.FromTo,
				From: 5,
				To:   1,
			},
		},
	},
}
