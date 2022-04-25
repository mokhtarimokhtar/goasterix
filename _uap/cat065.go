package _uap

// Cat065V15 User Application Profile
// version 1.19
var Cat065V15 = StandardUAP{
	Name:     "cat065_1.5",
	Category: 65,
	Version:  1.5,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I065/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I065/000",
			Description: "Message Type",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         3,
			DataItem:    "I065/015",
			Description: "Service Identification",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         4,
			DataItem:    "I065/030",
			Description: "Time Of Message",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         5,
			DataItem:    "I065/020",
			Description: "Batch Number",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         6,
			DataItem:    "I065/040",
			Description: "SDPS Configuration and Status",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         7,
			DataItem:    "I065/050",
			Description: "Service Status Report",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:  8,
			Type: Spare,
		},
		{
			FRN:  9,
			Type: Spare,
		},
		{
			FRN:  10,
			Type: Spare,
		},
		{
			FRN:  11,
			Type: Spare,
		},
		{
			FRN:  12,
			Type: Spare,
		},
		{
			FRN:         13,
			DataItem:    "RE-Data Item",
			Description: "Reserved Expansion Field",
			Type:        RE,
		},
		{
			FRN:         14,
			DataItem:    "SP-Data Item",
			Description: "Special Purpose Field",
			Type:        SP,
		},

		//FX : Field Extension Indicator
	},
}
