package goasterix

var CatForTest = StandardUAP{
	Name:     "catfortest_0.1",
	Category: 26, // not exist
	Version:  0.1,
	DataItems: []Item{
		&Fixed{
			Base: Base{
				FRN:         1,
				DataItem:    "I026/010",
				Description: "Fixed type field for test",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Extended{
			Base: Base{
				FRN:         2,
				DataItem:    "I026/020",
				Description: "Extended type field for test",
				Type:        ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 2,
		},
		&Explicit{
			Base: Base{
				FRN:         3,
				DataItem:    "I026/030",
				Description: "Explicit type field for test",
				Type:        ExplicitField,
			},
		},
		&Repetitive{
			Base: Base{
				FRN:         4,
				DataItem:    "I026/040",
				Description: "Repetitive type field for test",
				Type:        RepetitiveField,
			},
			SubItemSize: 3,
		},
		&Spare{
			Base{
				FRN: 5,
			},
		},
		&Compound{
			Base: Base{
				FRN:         6,
				DataItem:    "I026/060",
				Description: "Compound type field for test",
				Type:        CompoundField,
			},
			Secondary: []Item{
				&Fixed{
					Base: Base{
						FRN:         1,
						DataItem:    "Compound/001",
						Description: "Compound Fixed type field for test",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Spare{Base{FRN: 2}},
				&Extended{
					Base: Base{
						FRN:         3,
						DataItem:    "Compound/003",
						Description: "Compound Extended type field for test",
						Type:        ExtendedField,
					},
					PrimaryItemSize:   1,
					SecondaryItemSize: 1,
				},
				&Repetitive{
					Base: Base{
						FRN:         4,
						DataItem:    "Compound/004",
						Description: "Compound Repetitive type field for test",
						Type:        RepetitiveField,
					},
					SubItemSize: 2,
				},
				&Explicit{
					Base: Base{
						FRN:         5,
						DataItem:    "Compound/005",
						Description: "Compound Explicit type field for test",
						Type:        ExplicitField,
					},
				},
			},
		},
		&ReservedExpansion{
			Base: Base{
				FRN:         7,
				DataItem:    "RE",
				Description: "Reserved Expansion type field for test",
				Type:        REField,
			},
		},
		&SpecialPurpose{
			Base: Base{
				FRN:         8,
				DataItem:    "SP",
				Description: "Special Purpose type field for test",
				Type:        SPField,
			},
		},
		&Spare{
			Base{
				FRN: 9,
			},
		},
	},
}
