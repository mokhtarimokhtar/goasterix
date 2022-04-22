package uap

// Cat255StrV51 User Application Profile
// version 5.1
// French ANSP specific category
var Cat255StrV51 = StandardUAP{
	Category: 255,
	Version:  5.1,
	DataItems: []DataField{
		{
			FRN:      1,
			DataItem: "I255/010",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "I255/020",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:      3,
			DataItem: "I255/030",
			Type:     Extended,
			Size: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:      4,
			DataItem: "I255/050",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:      5,
			DataItem: "I255/060",
			Type:     Repetitive,
			Size: SizeField{
				ForRepetitive: 1,
			},
		},
		{
			FRN:      6,
			DataItem: "I255/070",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 9,
			},
		},
		{
			FRN:      7,
			DataItem: "I255/040",
			Type:     Repetitive,
			Size: SizeField{
				ForRepetitive: 10,
			},
		},
	},
}
