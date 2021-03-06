package uap

// Cat255StrV51 User Application Profile
// version 5.1
// French ANSP specific category
var Cat255StrV51 = StandardUAP{
	Category: 255,
	Version:  5.1,
	Items: []DataField{
		{
			FRN:      1,
			DataItem: "I255/010",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "I255/020",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 3,
			},
		},
		{
			FRN:      3,
			DataItem: "I255/030",
			Type:     Extended,
			Extended: ExtendedField{
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:      4,
			DataItem: "I255/050",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 4,
			},
		},
		{
			FRN:      5,
			DataItem: "I255/060",
			Type:     Repetitive,
			Repetitive: RepetitiveField{
				SubItemSize: 1,
			},
		},
		{
			FRN:      6,
			DataItem: "I255/070",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 9,
			},
		},
		{
			FRN:      7,
			DataItem: "I255/040",
			Type:     Repetitive,
			Repetitive: RepetitiveField{
				SubItemSize: 10,
			},
		},
	},
}
