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
				Name: Fixed,
				Size: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "I026/002",
			Type: TypeField{
				Name: Extended,
				Size: 1,
			},
		},
		{
			FRN:      3,
			DataItem: "I026/003",
			Type: TypeField{
				Name: Compound,
				Meta: MetaField{
					8: {Name: Fixed, Size: 1},
					7: {Name: Spare},
					6: {Name: Extended, Size: 1},
					5: {Name: Spare},
					4: {Name: Repetitive, Size: 2},
					3: {Name: Spare},
					2: {Name: Explicit},
				},
			},
		},
		{
			FRN:      4,
			DataItem: "I026/004",
			Type: TypeField{
				Name: Repetitive,
				Size: 2,
			},
		},
		{
			FRN:      5,
			DataItem: "I026/005",
			Type: TypeField{
				Name: Explicit,
			},
		},
		{
			FRN:      6,
			DataItem: "I026/006",
			Type: TypeField{
				Name: RFS,
			},
		},
		{
			FRN: 7, DataItem: "NA", Type: TypeField{Name: Spare},
		},
		{
			FRN: 8, DataItem: "SP-Data Item", Type: TypeField{Name: SP},
		},
	},
}
