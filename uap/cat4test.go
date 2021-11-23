package uap

// Cat4Test User Application Profile
// Specific for testing
var Cat4Test = StandardUAP{
	Name:     "cat4test_0.1",
	Category: 26, // not exist
	Version:  0.1,
	Items: []DataField{
		{
			FRN:      1,
			DataItem: "I026/001",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:      2,
			DataItem: "I026/002",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:      3,
			DataItem: "I026/003",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, Size: 1},
						7: {NameType: Spare},
						6: {NameType: Extended, PrimarySize: 1, SecondarySize: 1},
						5: {NameType: Spare},
						4: {NameType: Repetitive, Size: 2},
						3: {NameType: Spare},
						2: {NameType: Explicit},
					},
				},
			},
		},
		{
			FRN:      4,
			DataItem: "I026/004",
			Type: TypeField{
				NameType: Repetitive,
				Size:     2,
			},
		},
		{
			FRN:      5,
			DataItem: "I026/005",
			Type: TypeField{
				NameType: Explicit,
			},
		},
		{
			FRN:      6,
			DataItem: "I026/006",
			Type: TypeField{
				NameType: RFS,
			},
		},
		{
			FRN: 7, DataItem: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 8, DataItem: "SP-Data Item", Type: TypeField{NameType: SP},
		},
		/*
			{
				FRN:      9,
				DataItem: "I026/008",
				Type:     TypeField{NameType: Compound2},
				CompoundItems: &[]DataField{
					{
						FRN:         1,
						DataItem:    "Compound001",
						Description: "",
						Type:        TypeField{NameType: Fixed, Size: 1},
					},
					{
						FRN:  2,
						Type: TypeField{NameType: Spare},
					},
					{
						FRN:         3,
						DataItem:    "Compound003",
						Description: "",
						Type:        TypeField{NameType: Extended, PrimarySize: 1, SecondarySize: 1},
					},
					{
						FRN:  4,
						Type: TypeField{NameType: Spare},
					},
					{
						FRN:         5,
						DataItem:    "Compound005",
						Description: "",
						Type:        TypeField{NameType: Repetitive, Size: 2},
					},
					{
						FRN:  6,
						Type: TypeField{NameType: Spare},
					},
					{
						FRN:         7,
						DataItem:    "Compound005",
						Description: "",
						Type:        TypeField{NameType: Explicit, Size: 2},
					},
				},
			},

		*/
	},
}
