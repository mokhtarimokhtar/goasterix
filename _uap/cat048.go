package _uap

// Cat048V127 User Application Profile
// version 1.27
var Cat048V127 = StandardUAP{
	Name:     "cat048_1.27",
	Category: 48,
	Version:  1.27,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I048/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I048/140",
			Description: "Time-of-Day",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         3,
			DataItem:    "I048/020",
			Description: "Target Report Descriptor",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         4,
			DataItem:    "I048/040",
			Description: "Measured Position in Slant Polar Coordinates",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         5,
			DataItem:    "I048/070",
			Description: "Mode-3/A Code in Octal Representation",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         6,
			DataItem:    "I048/090",
			Description: "Flight Level in Binary Representation",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         7,
			DataItem:    "I048/130",
			Description: "Radar Plot Characteristics",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "SRL",
					Description: "SSR plot runlength",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         2,
					DataItem:    "SRR",
					Description: "Number of received replies",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         3,
					DataItem:    "SAM",
					Description: "Amplitude of received replies for M(SSR)",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         4,
					DataItem:    "PRL",
					Description: "PSR plot runlength",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "PAM",
					Description: "PSR amplitude",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         6,
					DataItem:    "RPD",
					Description: "Difference in Range between PSR and SSR plot",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
				{
					FRN:         7,
					DataItem:    "APD",
					Description: "Difference in Azimuth between PSR and SSR plot",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 1,
					},
				},
			},
		},
		// FX
		{
			FRN:         8,
			DataItem:    "I048/220",
			Description: "Aircraft Address",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         9,
			DataItem:    "I048/240",
			Description: "Aircraft Identification",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 6,
			},
		},
		{
			FRN:         10,
			DataItem:    "I048/250",
			Description: "Mode S MB Data",
			Type:        Repetitive,
			SizeItem: SizeField{
				ForRepetitive: 8,
			},
		},
		{
			FRN:         11,
			DataItem:    "I048/161",
			Description: "Track Number",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         12,
			DataItem:    "I048/042",
			Description: "Calculated Position in Cartesian Coordinates",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         13,
			DataItem:    "I048/200",
			Description: "Calculated Track Velocity in Polar Representation",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         14,
			DataItem:    "I048/170",
			Description: "Track Status",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         15,
			DataItem:    "I048/210",
			Description: "Track Quality",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         16,
			DataItem:    "I048/030",
			Description: "Warning/Error Conditions/Target Classification",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         17,
			DataItem:    "I048/080",
			Description: "Mode-3/A Code Confidence Indicator",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         18,
			DataItem:    "I048/100",
			Description: "Mode-C Code and Confidence Indicator",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         19,
			DataItem:    "I048/110",
			Description: "Height Measured by 3D Radar",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         20,
			DataItem:    "I048/120",
			Description: "Radial Doppler Speed",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "CAL",
					Description: "Calculated Doppler Speed",
					Type:        Fixed,
					SizeItem: SizeField{
						ForFixed: 2,
					},
				},
				{
					FRN:         2,
					DataItem:    "RDS",
					Description: "Raw Doppler Speed",
					Type:        Repetitive,
					SizeItem: SizeField{
						ForRepetitive: 2,
					},
				},
				{
					FRN:  3,
					Type: Spare,
				},
				{
					FRN:  4,
					Type: Spare,
				},
				{
					FRN:  5,
					Type: Spare,
				},
				{
					FRN:  6,
					Type: Spare,
				},
				{
					FRN:  7,
					Type: Spare,
				},
			},
		},
		{
			FRN:         21,
			DataItem:    "I048/230",
			Description: "Communications / ACAS Capability and Flight Status",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         22,
			DataItem:    "I048/260",
			Description: "ACAS Resolution Advisory Report",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 7,
			},
		},
		{
			FRN:         23,
			DataItem:    "I048/055",
			Description: "Mode-1 Code in Octal Representation",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         24,
			DataItem:    "I048/050",
			Description: "Mode-2 Code in Octal Representation",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         25,
			DataItem:    "I048/065",
			Description: "Mode-1 Code Confidence Indicator",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         26,
			DataItem:    "I048/060",
			Description: "Mode-2 Code Confidence Indicator",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         27,
			DataItem:    "SP-Data Item",
			Description: "Special Purpose Field",
			Type:        SP,
		},
		{
			FRN:         28,
			DataItem:    "RE-Data Item",
			Description: "Reserved Expansion Field",
			Type:        RE,
		},
	},
}
