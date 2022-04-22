package uap

// Cat032StrV70 User Application Profile
// version 5.1
// French ANSP specific category
var Cat032StrV70 = StandardUAP{
	Category: 32,
	Version:  7.0,
	DataItems: []DataField{
		{
			FRN:      1,
			DataItem: "I032/010",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "I032/020",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:      3,
			DataItem: "I032/060",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:      4,
			DataItem: "I032/070",
			Type:     Fixed,
			Size: SizeField{
				ForFixed: 15,
			},
		},
		{
			FRN:      5,
			DataItem: "I032/080",
			Type:     Extended,
			Size: SizeField{
				ForExtendedPrimary:   12,
				ForExtendedSecondary: 1,
			},
		},
	},
}
