package goasterix

import "github.com/mokhtarimokhtar/goasterix/item"

var CatForTest = item.StandardUAP{
	Name:     "catfortest_0.1",
	Category: 26, // not exist
	Version:  0.1,
	DataItems: []item.DataItem{
		&item.Fixed{
			Base: item.Base{
				FRN:          1,
				DataItemName: "I026/010",
				Description:  "Fixed type field for test",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				&item.SubItemFromTo{
					Name: "SAC",
					Type: item.FromToField,
					Pos:  item.BitPosition{From: 16, To: 9},
				},
				&item.SubItemFromTo{
					Name: "SIC",
					Type: item.FromToField,
					Pos:  item.BitPosition{From: 8, To: 1},
				},
			},
		},
		&item.Extended{
			Base: item.Base{
				FRN:          2,
				DataItemName: "I026/020",
				Description:  "Extended type field for test",
				Type:         item.ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 2,
			/*
				SubItems: []item.SubItem{
					&item.SubItemFromTo{
						Name: "TYP",
						Type: item.FromToField,
						From: 8,
						To:   6,
					},
					&item.SubItemBit{Name: "SIM", Type: item.BitField, Pos: 5},
					&item.SubItemBit{Name: "RDP", Type: item.BitField, Pos: 4},
					&item.SubItemBit{Name: "SPI", Type: item.BitField, Pos: 3},
					&item.SubItemBit{Name: "RAB", Type: item.BitField, Pos: 2},
					&item.SubItemBit{Name: "TST", Type: item.BitField, Pos: 8},
					&item.SubItemBit{Name: "ERR", Type: item.BitField, Pos: 7},
					&item.SubItemBit{Name: "XPP", Type: item.BitField, Pos: 6},
					&item.SubItemBit{Name: "ME", Type: item.BitField, Pos: 5},
					&item.SubItemBit{Name: "MI", Type: item.BitField, Pos: 4},
					&item.SubItemFromTo{
						Name: "FOE/FRI",
						Type: item.FromToField,
						From: 3,
						To:   2,
					},
				},
			*/
		},
		&item.Explicit{
			Base: item.Base{
				FRN:          3,
				DataItemName: "I026/030",
				Description:  "Explicit type field for test",
				Type:         item.ExplicitField,
			},
		},
		&item.Repetitive{
			Base: item.Base{
				FRN:          4,
				DataItemName: "I026/040",
				Description:  "Repetitive type field for test",
				Type:         item.RepetitiveField,
			},
			SubItemSize: 3,
			SubItems: []item.SubItem{
				&item.SubItemFromTo{
					Name: "DOP",
					Type: item.FromToField,
					Pos:  item.BitPosition{From: 24, To: 17},
				},
				&item.SubItemFromTo{
					Name: "AMB",
					Type: item.FromToField,
					Pos:  item.BitPosition{From: 16, To: 9},
				},
				&item.SubItemFromTo{
					Name: "FRQ",
					Type: item.FromToField,
					Pos:  item.BitPosition{From: 8, To: 1},
				},
			},
		},
		&item.Spare{Base: item.Base{FRN: 5}},
		&item.Compound{
			Base: item.Base{
				FRN:          6,
				DataItemName: "I026/060",
				Description:  "Compound type field for test",
				Type:         item.CompoundField,
			},
			Secondary: []item.DataItem{
				&item.Fixed{
					Base: item.Base{
						FRN:          1,
						DataItemName: "Compound/001",
						Description:  "Compound Fixed type field for test",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Spare{Base: item.Base{FRN: 2}},
				&item.Extended{
					Base: item.Base{
						FRN:          3,
						DataItemName: "Compound/003",
						Description:  "Compound Extended type field for test",
						Type:         item.ExtendedField,
					},
					PrimaryItemSize:   1,
					SecondaryItemSize: 1,
				},
				&item.Repetitive{
					Base: item.Base{
						FRN:          4,
						DataItemName: "Compound/004",
						Description:  "Compound Repetitive type field for test",
						Type:         item.RepetitiveField,
					},
					SubItemSize: 2,
				},
				&item.Explicit{
					Base: item.Base{
						FRN:          5,
						DataItemName: "Compound/005",
						Description:  "Compound Explicit type field for test",
						Type:         item.ExplicitField,
					},
				},
			},
		},
		&item.ReservedExpansion{
			Base: item.Base{
				FRN:          7,
				DataItemName: "RE",
				Description:  "Reserved Expansion type field for test",
				Type:         item.REField,
			},
		},
		&item.SpecialPurpose{
			Base: item.Base{
				FRN:          8,
				DataItemName: "SP",
				Description:  "Special Purpose type field for test",
				Type:         item.SPField,
			},
		},
		&item.Spare{Base: item.Base{FRN: 9}},
	},
}
