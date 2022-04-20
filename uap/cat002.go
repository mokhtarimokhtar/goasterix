package uap

// Cat002V10 User Application Profile CAT002
// version 1.0
var Cat002V10 = StandardUAP{
	Category: 2,
	Version:  1.0,
	Items: []DataField{
		{
			FRN:         1,
			DataItem:    "I002/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I002/000",
			Description: "Message Type",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         3,
			DataItem:    "I002/020",
			Description: "Sector Number",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         4,
			DataItem:    "I002/030",
			Description: "Time of Day",
			Type:        Fixed,
			Size: Size{
				ForFixed: 3,
			},
		},
		{
			FRN:         5,
			DataItem:    "I002/041",
			Description: "Antenna Rotation Period",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         6,
			DataItem:    "I002/050",
			Description: "Station Configuration Status",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         7,
			DataItem:    "I002/060",
			Description: "Station Processing Mode",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         8,
			DataItem:    "I002/070",
			Description: "Plot Count Values",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         9,
			DataItem:    "I002/100",
			Description: "Dynamic Window - Type 1",
			Type:        Fixed,
			Size: Size{
				ForFixed: 8,
			},
		},
		{
			FRN:         10,
			DataItem:    "I002/090",
			Description: "Collimation Error",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         11,
			DataItem:    "I002/080",
			Description: "Warning/Error Conditions",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:      12,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:         13,
			DataItem:    "SP-Data Item",
			Description: "Reserved for SP Indicator",
			Type:        SP,
		},
		{
			FRN:         14,
			DataItem:    "Random Field Sequencing",
			Description: "Reserved for RFS Indicator",
			Type:        RFS,
		},
	},
}
