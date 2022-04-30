package goasterix

import "github.com/mokhtarimokhtar/goasterix/item"

const (
	I010Name   = "I026/010"
	I010Desc   = "Fixed type field for test"
	I020Name   = "I026/020"
	I020Desc   = "Extended type field for test"
	I030Name   = "I026/030"
	I030Desc   = "Explicit type field for test"
	I040Name   = "I026/040"
	I040Desc   = "Repetitive type field for test"
	I060Name   = "I026/060"
	I060Desc   = "Compound type field for test"
	I060AName  = "I026/060/A"
	I060ADesc  = "Compound Fixed type field for test"
	I060BName  = "I026/060/B"
	I060BDesc  = "Compound Extended type field for test"
	I060CName  = "I026/060/C"
	I060CDesc  = "Compound Repetitive type field for test"
	I060DName  = "I026/060/D"
	I060DDesc  = "Compound Explicit type field for test"
	I060REName = "I060REName"
	I060REDesc = "Reserved Expansion type field for test"
	I060SPName = "I060SPName"
	I060SPDesc = "Special Purpose type field for test"
)

var CatForTest = item.UAP{
	Name:     "catfortest_0.1",
	Category: 26, // not used
	Version:  0.1,
	DataItems: []item.DataItem{
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN1,
				DataItemName: I010Name,
				Description:  I010Desc,
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
				DataItemName: I020Name,
				Description:  I020Desc,
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
				DataItemName: I030Name,
				Description:  I030Desc,
				Type:         item.ExplicitField,
			},
		},
		&item.Repetitive{
			Base: item.Base{
				FRN:          item.FRN4,
				DataItemName: I040Name,
				Description:  I040Desc,
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
				DataItemName: I060Name,
				Description:  I060Desc,
				Type:         item.CompoundField,
			},
			Secondary: []item.DataItem{
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN1,
						DataItemName: I060AName,
						Description:  I060ADesc,
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Spare{Base: item.Base{FRN: item.FRN2}},
				&item.Extended{
					Base: item.Base{
						FRN:          item.FRN3,
						DataItemName: I060BName,
						Description:  I060BDesc,
						Type:         item.ExtendedField,
					},
					PrimaryItemSize:   1,
					SecondaryItemSize: 1,
				},
				&item.Repetitive{
					Base: item.Base{
						FRN:          item.FRN4,
						DataItemName: I060CName,
						Description:  I060CDesc,
						Type:         item.RepetitiveField,
					},
					SubItemSize: 2,
				},
				&item.Explicit{
					Base: item.Base{
						FRN:          item.FRN5,
						DataItemName: I060DName,
						Description:  I060DDesc,
						Type:         item.ExplicitField,
					},
				},
			},
		},
		&item.ReservedExpansion{
			Base: item.Base{
				FRN:          item.FRN7,
				DataItemName: I060REName,
				Description:  I060REDesc,
				Type:         item.REField,
			},
		},
		&item.SpecialPurpose{
			Base: item.Base{
				FRN:          item.FRN8,
				DataItemName: I060SPName,
				Description:  I060SPDesc,
				Type:         item.SPField,
			},
		},
		&item.Spare{Base: item.Base{FRN: item.FRN9}},
	},
}
