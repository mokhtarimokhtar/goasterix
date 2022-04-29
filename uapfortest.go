package goasterix

import "github.com/mokhtarimokhtar/goasterix/item"

var CatForTest = item.UAP{
	Name:     "catfortest_0.1",
	Category: 26, // not exist
	Version:  0.1,
	DataItems: []item.DataItem{
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN1,
				DataItemName: "I026/010",
				Description:  "Fixed type field for test",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItemBits{
				{
					Name: "SAC",
					Type: item.FromToField,
					From: 16, To: 9,
				},
				{
					Name: "SIC",
					Type: item.FromToField,
					From: 8, To: 1,
				},
			},
		},
		&item.Extended{
			Base: item.Base{
				FRN:          item.FRN2,
				DataItemName: "I026/020",
				Description:  "Extended type field for test",
				Type:         item.ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 2,
			/*
				PrimarySubItems: []item.SubItem{
					&item.SubItemFromTo{
						Name: "TYP",
						Type: item.FromToField,
						Pos:  item.BitPosition{From: 8, To: 6},
					},
					&item.SubItemBit{Name: "SIM", Type: item.BitField, Pos: item.BitPosition{Bit: 5}},
					&item.SubItemBit{Name: "RDP", Type: item.BitField, Pos: item.BitPosition{Bit: 4}},
					&item.SubItemBit{Name: "SPI", Type: item.BitField, Pos: item.BitPosition{Bit: 3}},
					&item.SubItemBit{Name: "RAB", Type: item.BitField, Pos: item.BitPosition{Bit: 2}},
				},
				SecondarySubItems: []item.SubItem{
					&item.SubItemBit{Name: "TST", Type: item.BitField, Pos: item.BitPosition{Bit: 8}},
					&item.SubItemBit{Name: "ERR", Type: item.BitField, Pos: item.BitPosition{Bit: 7}},
					&item.SubItemBit{Name: "XPP", Type: item.BitField, Pos: item.BitPosition{Bit: 6}},
					&item.SubItemBit{Name: "ME", Type: item.BitField, Pos: item.BitPosition{Bit: 5}},
					&item.SubItemBit{Name: "MI", Type: item.BitField, Pos: item.BitPosition{Bit: 4}},
					&item.SubItemFromTo{
						Name: "FOE/FRI",
						Type: item.FromToField,
						Pos:  item.BitPosition{From: 3, To: 2},
					},
				},
			*/
		},
		&item.Explicit{
			Base: item.Base{
				FRN:          item.FRN3,
				DataItemName: "I026/030",
				Description:  "Explicit type field for test",
				Type:         item.ExplicitField,
			},
		},
		&item.Repetitive{
			Base: item.Base{
				FRN:          item.FRN4,
				DataItemName: "I026/040",
				Description:  "Repetitive type field for test",
				Type:         item.RepetitiveField,
			},
			SubItemSize: 3,
			SubItems: []item.SubItemBits{
				{
					Name: "DOP",
					Type: item.FromToField,
					From: 24, To: 17,
				},
				{
					Name: "AMB",
					Type: item.FromToField,
					From: 16, To: 9,
				},
				{
					Name: "FRQ",
					Type: item.FromToField,
					From: 8, To: 1,
				},
			},
		},
		&item.Spare{Base: item.Base{FRN: item.FRN5}},
		&item.Compound{
			Base: item.Base{
				FRN:          item.FRN6,
				DataItemName: "I026/060",
				Description:  "Compound type field for test",
				Type:         item.CompoundField,
			},
			Secondary: []item.DataItem{
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN1,
						DataItemName: "Compound/001",
						Description:  "Compound Fixed type field for test",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Spare{Base: item.Base{FRN: item.FRN2}},
				&item.Extended{
					Base: item.Base{
						FRN:          item.FRN3,
						DataItemName: "Compound/003",
						Description:  "Compound Extended type field for test",
						Type:         item.ExtendedField,
					},
					PrimaryItemSize:   1,
					SecondaryItemSize: 1,
				},
				&item.Repetitive{
					Base: item.Base{
						FRN:          item.FRN4,
						DataItemName: "Compound/004",
						Description:  "Compound Repetitive type field for test",
						Type:         item.RepetitiveField,
					},
					SubItemSize: 2,
				},
				&item.Explicit{
					Base: item.Base{
						FRN:          item.FRN5,
						DataItemName: "Compound/005",
						Description:  "Compound Explicit type field for test",
						Type:         item.ExplicitField,
					},
				},
			},
		},
		&item.ReservedExpansion{
			Base: item.Base{
				FRN:          item.FRN7,
				DataItemName: "RE",
				Description:  "Reserved Expansion type field for test",
				Type:         item.REField,
			},
		},
		&item.SpecialPurpose{
			Base: item.Base{
				FRN:          item.FRN8,
				DataItemName: "SP",
				Description:  "Special Purpose type field for test",
				Type:         item.SPField,
			},
		},
		&item.Spare{Base: item.Base{FRN: item.FRN9}},
	},
}
