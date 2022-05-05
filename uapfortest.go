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
	I060REName = "I060RE"
	I060REDesc = "Reserved Expansion type field for test"
	I060SPName = "I060SP"
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
			SubItems: []item.SubItem{
				{
					Name: "SUB-A",
					Type: item.FromToField,
					From: 16, To: 9,
				},
				{
					Name: "SUB-B",
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
			SecondaryItemSize: 1,
			SubItems: []item.SubItem{
				{Name: "SUB-A", Type: item.FromToField, From: 8, To: 6},
				{Name: "SUB-B", Type: item.BitField, Bit: 5},
				{Name: "SUB-C", Type: item.BitField, Bit: 4},
				{Name: "SUB-D", Type: item.BitField, Bit: 3},
				{Name: "SUB-E", Type: item.BitField, Bit: 2},
				{Name: "FX", Type: item.BitField, Bit: 1},

				{Name: "SUB-F", Type: item.BitField, Bit: 8},
				{Name: "SUB-G", Type: item.BitField, Bit: 7},
				{Name: "SUB-H", Type: item.BitField, Bit: 6},
				{Name: "SUB-I", Type: item.BitField, Bit: 5},
				{Name: "SUB-J", Type: item.BitField, Bit: 4},
				{Name: "SUB-K", Type: item.FromToField, From: 3, To: 2},
				{Name: "FX", Type: item.BitField, Bit: 1},
			},
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
			SubItems: []item.SubItem{
				{
					Name: "SUB-A",
					Type: item.FromToField,
					From: 24, To: 17,
				},
				{
					Name: "SUB-B",
					Type: item.FromToField,
					From: 16, To: 9,
				},
				{
					Name: "SUB-C",
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
					SubItems: []item.SubItem{
						{
							Name: "SUB-A",
							Type: item.FromToField,
							From: 8, To: 1,
						},
					},
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
					SubItems: []item.SubItem{
						{Name: "SUB-A", Type: item.FromToField, From: 8, To: 2},
						{Name: "FX", Type: item.BitField, Bit: 1},

						{Name: "SUB-B", Type: item.BitField, Bit: 8},
						{Name: "SUB-C", Type: item.FromToField, From: 7, To: 2},
						{Name: "FX", Type: item.BitField, Bit: 1},
					},
				},
				&item.Repetitive{
					Base: item.Base{
						FRN:          item.FRN4,
						DataItemName: I060CName,
						Description:  I060CDesc,
						Type:         item.RepetitiveField,
					},
					SubItemSize: 2,
					SubItems: []item.SubItem{
						{
							Name: "SUB-A",
							Type: item.FromToField,
							From: 16, To: 1,
						},
					},
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
