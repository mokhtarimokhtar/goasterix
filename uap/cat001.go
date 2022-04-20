package uap

var Cat001V12 = StandardUAP{
	Name:     "cat001_1.2",
	Category: 1,
	Version:  1.2,
	Items: []DataField{
		{
			FRN:         1,
			DataItem:    "I001/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I001/020",
			Description: "Target Report Descriptor",
			Conditional: true,
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
	},
}
var Cat001PlotV12 = []DataField{
	{
		FRN:         3,
		DataItem:    "I001/040",
		Description: "Measured Position in Polar Coordinates",
		Type:        Fixed,
		Size: Size{
			ForFixed: 4,
		},
	},
	{
		FRN:         4,
		DataItem:    "I001/070",
		Description: "Mode-3/A Code in Octal Representation",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         5,
		DataItem:    "I001/090",
		Description: "Mode-C Code in Binary Representation",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         6,
		DataItem:    "I001/130",
		Description: "Radar Plot Characteristics",
		Type:        Extended,
		Size: Size{
			ForExtendedPrimary:   1,
			ForExtendedSecondary: 1,
		},
	},
	{
		FRN:         7,
		DataItem:    "I001/141",
		Description: "Truncated Time of Day",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         8,
		DataItem:    "I001/050",
		Description: "Mode-2 Code in Octal Representation",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         9,
		DataItem:    "I001/120",
		Description: "Measured Radial Doppler Speed",
		Type:        Fixed,
		Size: Size{
			ForFixed: 1,
		},
	},
	{
		FRN:         10,
		DataItem:    "I001/131",
		Description: "Received Power",
		Type:        Fixed,
		Size: Size{
			ForFixed: 1,
		},
	},
	{
		FRN:         11,
		DataItem:    "I001/080",
		Description: "Mode-3/A Code Confidence Indicator",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         12,
		DataItem:    "I001/100",
		Description: "Mode-C Code and Code Confidence Indicator",
		Type:        Fixed,
		Size: Size{
			ForFixed: 4,
		},
	},
	{
		FRN:         13,
		DataItem:    "I001/060",
		Description: "Mode-2 Code Confidence Indicator",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         14,
		DataItem:    "I001/030",
		Description: "Warning/Error Conditions",
		Type:        Extended,
		Size: Size{
			ForExtendedPrimary:   1,
			ForExtendedSecondary: 1,
		},
	},
	{
		FRN:         15,
		DataItem:    "I001/150",
		Description: "Presence of X-Pulse",
		Type:        Fixed,
		Size: Size{
			ForFixed: 1,
		},
	},
	{
		FRN:      16,
		DataItem: "NA",
		Type:     Spare,
	},
	{
		FRN:      16,
		DataItem: "NA",
		Type:     Spare,
	},
	{
		FRN:      17,
		DataItem: "NA",
		Type:     Spare,
	},
	{
		FRN:      18,
		DataItem: "NA",
		Type:     Spare,
	},
	{
		FRN:      19,
		DataItem: "NA",
		Type:     Spare,
	},
	{
		FRN:         20,
		DataItem:    "SP-Data Item",
		Description: "",
		Type:        SP,
	},
	{
		FRN:         21,
		DataItem:    "Random Field Sequencing",
		Description: "",
		Type:        RFS,
	},
}

var Cat001TrackV12 = []DataField{
	{
		FRN:         3,
		DataItem:    "I001/161",
		Description: "Track/Plot Number",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         4,
		DataItem:    "I001/040",
		Description: "Measured Position in Polar Coordinates",
		Type:        Fixed,
		Size: Size{
			ForFixed: 4,
		},
	},
	{
		FRN:         5,
		DataItem:    "I001/042",
		Description: "Calculated Position in Cartesian Coordinates",
		Type:        Fixed,
		Size: Size{
			ForFixed: 4,
		},
	},
	{
		FRN:         6,
		DataItem:    "I001/200",
		Description: "Calculated Track Velocity in polar Coordinates",
		Type:        Fixed,
		Size: Size{
			ForFixed: 4,
		},
	},
	{
		FRN:         7,
		DataItem:    "I001/070",
		Description: "Mode-3/A Code in Octal Representation",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         8,
		DataItem:    "I001/090",
		Description: "Mode-C Code in Binary Representation",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         9,
		DataItem:    "I001/141",
		Description: "Truncated Time of Day",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         10,
		DataItem:    "I001/130",
		Description: "Radar Plot Characteristics",
		Size: Size{
			ForExtendedPrimary:   1,
			ForExtendedSecondary: 1,
		},
	},
	{
		FRN:         11,
		DataItem:    "I001/131",
		Description: "Received Power",
		Type:        Fixed,
		Size: Size{
			ForFixed: 1,
		},
	},
	{
		FRN:         12,
		DataItem:    "I001/120",
		Description: "Measured Radial Doppler Speed",
		Type:        Fixed,
		Size: Size{
			ForFixed: 1,
		},
	},
	{
		FRN:         13,
		DataItem:    "I001/170",
		Description: "Track Status",
		Type:        Extended,
		Size: Size{
			ForExtendedPrimary:   1,
			ForExtendedSecondary: 1,
		},
	},
	{
		FRN:         14,
		DataItem:    "I001/210",
		Description: "Track Quality",
		Type:        Extended,
		Size: Size{
			ForExtendedPrimary:   1,
			ForExtendedSecondary: 1,
		},
	},
	{
		FRN:         15,
		DataItem:    "I001/050",
		Description: "Mode-2 Code in Octal Representation",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         16,
		DataItem:    "I001/080",
		Description: "Mode-3/A Code Confidence Indicator",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         17,
		DataItem:    "I001/100",
		Description: "Mode-C Code and Code Confidence Indicator",
		Type:        Fixed,
		Size: Size{
			ForFixed: 4,
		},
	},
	{
		FRN:         18,
		DataItem:    "I001/060",
		Description: "Mode-2 Code Confidence Indicator",
		Type:        Fixed,
		Size: Size{
			ForFixed: 2,
		},
	},
	{
		FRN:         19,
		DataItem:    "I001/030",
		Description: "Warning/Error Conditions",
		Type:        Extended,
		Size: Size{
			ForExtendedPrimary:   1,
			ForExtendedSecondary: 1,
		},
	},
	{
		FRN:         20,
		DataItem:    "SP-Data Item",
		Description: "Reserved for Special Purpose Indicator",
		Type:        SP,
	},
	{
		FRN:         21,
		DataItem:    "Random Field Sequencing",
		Description: "Reserved for RFS Indicator",
		Type:        RFS,
	},
	{
		FRN:         22,
		DataItem:    "I001/150",
		Description: "Presence of X-Pulse",
		Type:        Fixed,
		Size: Size{
			ForFixed: 1,
		},
	},
}
