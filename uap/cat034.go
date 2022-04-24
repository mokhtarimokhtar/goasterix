package uap

// Cat034V127 User Application Profile CAT034
// version 1.27
var Cat034V127 = StandardUAP{
	Category: 34,
	Version:  1.27,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I034/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I034/000",
			Description: "Message Type",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         3,
			DataItem:    "I034/030",
			Description: "Time-of-Day",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         4,
			DataItem:    "I034/020",
			Description: "Sector Number",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         5,
			DataItem:    "I034/041",
			Description: "Antenna Rotation Period",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN: 6, DataItem: "I034/050",
			Description: "System Configuration and Status",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "COM",
					Description: "Common Part",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:  2,
					Type: Spare,
				},
				{
					FRN:  3,
					Type: Spare,
				},
				{
					FRN:         4,
					DataItem:    "PSR",
					Description: "Specific Status for PSR Sensor",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "SSR",
					Description: "Specific Status for SSR Sensor",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         6,
					DataItem:    "MDS",
					Description: "Specific Status for Mode S Sensor",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 2,
					},
				},
				{
					FRN:  7,
					Type: Spare,
				},
			},
		},
		{
			FRN: 7, DataItem: "I034/060",
			Description: "System Processing Mode",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "COM",
					Description: "Common Part",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:  2,
					Type: Spare,
				},
				{
					FRN:  3,
					Type: Spare,
				},
				{
					FRN:         4,
					DataItem:    "PSR",
					Description: "Specific Processing Mode information for PSR Sensor",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "SSR",
					Description: "Specific Processing Mode information for SSR Sensor",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         6,
					DataItem:    "MDS",
					Description: "Specific Processing Mode information for Mode S Sensor",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:  7,
					Type: Spare,
				},
			},
		},
		{
			FRN:         8,
			DataItem:    "I034/070",
			Description: "Message Count Values",
			Type:        Repetitive,
			SizeItem: SizeField{
				ForRepetitive: 2,
			},
		},
		{
			FRN:         9,
			DataItem:    "I034/100",
			Description: "Generic Polar Window",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 8,
			},
		},
		{
			FRN:         10,
			DataItem:    "I034/110",
			Description: "Data Filter",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         11,
			DataItem:    "I034/120",
			Description: "3D-Position of Data Source",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 8,
			},
		},
		{
			FRN:         12,
			DataItem:    "I034/090",
			Description: "Collimation Error",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
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
	},
}
