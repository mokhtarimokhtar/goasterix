package uap

// Cat063V16 User Application Profile
// version 1.6
var Cat063V16 = StandardUAP{
	Name:     "cat063_1.6",
	Category: 63,
	Version:  1.6,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I063/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I063/015",
			Description: "Service Identification",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         3,
			DataItem:    "I063/030",
			Description: "Time of Message",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         4,
			DataItem:    "I063/050",
			Description: "Sensor Identifier",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         5,
			DataItem:    "I063/060",
			Description: "Sensor Configuration and Status",
			Type:        Extended,
			Size: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         6,
			DataItem:    "I063/070",
			Description: "Time Stamping Bias",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         7,
			DataItem:    "I063/080",
			Description: "SSR/Mode S Range Gain and Bias",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 4,
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:         8,
			DataItem:    "I063/081",
			Description: "SSR/Mode S Azimuth Bias",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         9,
			DataItem:    "I063/090",
			Description: "PSR Range Gain and Bias",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         10,
			DataItem:    "I063/091",
			Description: "PSR Azimuth Bias",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         11,
			DataItem:    "I063/092",
			Description: "PSR Elevation Bias",
			Type:        Fixed,
			Size: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:      12,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      13,
			DataItem: "RE-Data Item",
			Type:     RE,
		},
		{
			FRN:      14,
			DataItem: "SP-Data Item",
			Type:     SP,
		},
		//FX : Field Extension Indicator
	},
}
