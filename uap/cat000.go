package uap

var Cat000 = StandardUAP{
	Name:     "cat000",
	Category: 0,
	Version:  0.0,
	Items: []DataField{
		{
			FRN:         1,
			DataItem:    "I048/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         3,
			DataItem:    "I048/020",
			Description: "Target Report Descriptor",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 2,
			},
		},
		{
			FRN:         10,
			DataItem:    "I048/250",
			Description: "Mode S MB Data",
			Type:        Repetitive,
			Size: Size{
				ForRepetitive: 8,
			},
		},
	},
}
