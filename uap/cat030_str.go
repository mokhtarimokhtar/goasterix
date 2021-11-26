package uap

// Cat030StrV51 User Application Profile
// version 5.1
// French ANSP specific category
var Cat030StrV51 = StandardUAP{
	Name:     "STR",
	Category: 30,
	Version:  5.1,
	Items: []DataField{
		{
			FRN:      1,
			DataItem: "I030/010",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      3,
			DataItem: "I030/050",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 3,
			},
		},
		{
			FRN:      4,
			DataItem: "I030/020",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 3,
			},
		},
		{
			FRN:      5,
			DataItem: "I030/080",
			Type:     Extended,
			Extended: ExtendedField{
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:      6,
			DataItem: "I030/060",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      7,
			DataItem: "I030/100",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 4,
			},
		},
		{
			FRN:      8,
			DataItem: "I030/090",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 1,
			},
		},
		{
			FRN:      9,
			DataItem: "I030/135",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      10,
			DataItem: "I030/136",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      11,
			DataItem: "I030/181",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 4,
			},
		},
		{
			FRN:      12,
			DataItem: "I030/200",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 1,
			},
		},
		{
			FRN:      13,
			DataItem: "I030/220",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      14,
			DataItem: "I030/SPE",
			Type:     Extended,
			Extended: ExtendedField{
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:      15,
			DataItem: "I030/260",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      16,
			DataItem: "I030/400",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 7,
			},
		},
		{
			FRN:      17,
			DataItem: "I030/410",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      18,
			DataItem: "I030/430",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 4,
			},
		},
		{
			FRN:      19,
			DataItem: "I030/435",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 1,
			},
		},
		{
			FRN:      20,
			DataItem: "I030/440",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 4,
			},
		},
		{
			FRN:      21,
			DataItem: "I030/450",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 4,
			},
		},
		{
			FRN:      22,
			DataItem: "I030/130",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 2,
			},
		},
		{
			FRN:      23,
			DataItem: "I030/382",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 3,
			},
		},
		{
			FRN:      24,
			DataItem: "I030/384",
			Type:     Fixed,
			Fixed: FixedField{
				Size: 6,
			},
		},
	},
}
