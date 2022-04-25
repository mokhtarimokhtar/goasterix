package goasterix

var Cat048V127 = StandardUAP{
	Name:     "cat048_1.27",
	Category: 48,
	Version:  1.27,
	DataItems: []Item{
		&Fixed{
			Base: Base{
				FRN:         1,
				DataItem:    "I048/010",
				Description: "Data Source Identifier",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Fixed{
			Base: Base{
				FRN:         2,
				DataItem:    "I048/140",
				Description: "Time-of-Day",
				Type:        FixedField,
			},
			Size: 3,
		},
		&Extended{
			Base: Base{
				FRN:         3,
				DataItem:    "I048/020",
				Description: "Target Report Descriptor",
				Type:        ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
		},
		&Fixed{
			Base: Base{
				FRN:         4,
				DataItem:    "I048/040",
				Description: "Measured Position in Slant Polar Coordinates",
				Type:        FixedField,
			},
			Size: 4,
		},
		&Fixed{
			Base: Base{
				FRN:         5,
				DataItem:    "I048/070",
				Description: "Mode-3/A Code in Octal Representation",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Fixed{
			Base: Base{
				FRN:         6,
				DataItem:    "I048/090",
				Description: "Flight Level in Binary Representation",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Compound{
			Base: Base{
				FRN:         7,
				DataItem:    "I048/130",
				Description: "Radar Plot Characteristics",
				Type:        CompoundField,
			},
			Secondary: []Item{
				&Fixed{
					Base: Base{
						FRN:         1,
						DataItem:    "SRL",
						Description: "SSR plot runlength",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Fixed{
					Base: Base{
						FRN:         2,
						DataItem:    "SRR",
						Description: "Number of received replies",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Fixed{
					Base: Base{
						FRN:         3,
						DataItem:    "SAM",
						Description: "Amplitude of received replies for M(SSR)",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Fixed{
					Base: Base{
						FRN:         4,
						DataItem:    "PRL",
						Description: "PSR plot runlength",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Fixed{
					Base: Base{
						FRN:         5,
						DataItem:    "PAM",
						Description: "PSR amplitude",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Fixed{
					Base: Base{
						FRN:         6,
						DataItem:    "RPD",
						Description: "Difference in Range between PSR and SSR plot",
						Type:        FixedField,
					},
					Size: 1,
				},
				&Fixed{
					Base: Base{
						FRN:         7,
						DataItem:    "APD",
						Description: "Difference in Azimuth between PSR and SSR plot",
						Type:        FixedField,
					},
					Size: 1,
				},
			},
		},
		// FX
		&Fixed{
			Base: Base{
				FRN:         8,
				DataItem:    "I048/220",
				Description: "Aircraft Address",
				Type:        FixedField,
			},
			Size: 3,
		},
		&Fixed{
			Base: Base{
				FRN:         9,
				DataItem:    "I048/240",
				Description: "Aircraft Identification",
				Type:        FixedField,
			},
			Size: 6,
		},
		&Repetitive{
			Base: Base{
				FRN:         10,
				DataItem:    "I048/250",
				Description: "Mode S MB Data",
				Type:        RepetitiveField,
			},
			SubItemSize: 8,
		},
		&Fixed{
			Base: Base{
				FRN:         11,
				DataItem:    "I048/161",
				Description: "Track Number",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Fixed{
			Base: Base{
				FRN:         12,
				DataItem:    "I048/042",
				Description: "Calculated Position in Cartesian Coordinates",
				Type:        FixedField,
			},
			Size: 4,
		},
		&Fixed{
			Base: Base{
				FRN:         13,
				DataItem:    "I048/200",
				Description: "Calculated Track Velocity in Polar Representation",
				Type:        FixedField,
			},
			Size: 4,
		},
		&Extended{
			Base: Base{
				FRN:         14,
				DataItem:    "I048/170",
				Description: "Track Status",
				Type:        ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
		},
		// FX
		&Fixed{
			Base: Base{
				FRN:         15,
				DataItem:    "I048/210",
				Description: "Track Quality",
				Type:        FixedField,
			},
			Size: 4,
		},
		&Extended{
			Base: Base{
				FRN:         16,
				DataItem:    "I048/030",
				Description: "Warning/Error Conditions/Target Classification",
				Type:        ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
		},
		&Fixed{
			Base: Base{
				FRN:         17,
				DataItem:    "I048/080",
				Description: "Mode-3/A Code Confidence Indicator",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Fixed{
			Base: Base{
				FRN:         18,
				DataItem:    "I048/100",
				Description: "Mode-C Code and Confidence Indicator",
				Type:        FixedField,
			},
			Size: 4,
		},
		&Fixed{
			Base: Base{
				FRN:         19,
				DataItem:    "I048/110",
				Description: "Height Measured by 3D Radar",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Compound{
			Base: Base{
				FRN:         20,
				DataItem:    "I048/120",
				Description: "Radial Doppler Speed",
				Type:        CompoundField,
			},
			Secondary: []Item{
				&Fixed{
					Base: Base{
						FRN:         1,
						DataItem:    "CAL",
						Description: "Calculated Doppler Speed",
						Type:        FixedField,
					},
					Size: 2,
				},
				&Repetitive{
					Base: Base{
						FRN:         2,
						DataItem:    "RDS",
						Description: "Raw Doppler Speed",
						Type:        RepetitiveField,
					},
					SubItemSize: 2,
				},
				&Spare{Base{FRN: 3}},
				&Spare{Base{FRN: 4}},
				&Spare{Base{FRN: 5}},
				&Spare{Base{FRN: 6}},
				&Spare{Base{FRN: 7}},
			},
		},
		&Fixed{
			Base: Base{
				FRN:         21,
				DataItem:    "I048/230",
				Description: "Communications / ACAS Capability and Flight Status",
				Type:        FixedField,
			},
			Size: 2,
		},
		// FX
		&Fixed{
			Base: Base{
				FRN:         22,
				DataItem:    "I048/260",
				Description: "ACAS Resolution Advisory Report",
				Type:        FixedField,
			},
			Size: 7,
		},
		&Fixed{
			Base: Base{
				FRN:         23,
				DataItem:    "I048/055",
				Description: "Mode-1 Code in Octal Representation",
				Type:        FixedField,
			},
			Size: 1,
		},
		&Fixed{
			Base: Base{
				FRN:         24,
				DataItem:    "I048/050",
				Description: "Mode-2 Code in Octal Representation",
				Type:        FixedField,
			},
			Size: 2,
		},
		&Fixed{
			Base: Base{
				FRN:         25,
				DataItem:    "I048/065",
				Description: "Mode-1 Code Confidence Indicator",
				Type:        FixedField,
			},
			Size: 1,
		},
		&Fixed{
			Base: Base{
				FRN:         26,
				DataItem:    "I048/060",
				Description: "Mode-2 Code Confidence Indicator",
				Type:        FixedField,
			},
			Size: 2,
		},
		&SpecialPurpose{
			Base: Base{
				FRN:         27,
				DataItem:    "SP-Data Item",
				Description: "Special Purpose Field",
				Type:        SPField,
			},
		},
		&ReservedExpansion{
			Base: Base{
				FRN:         28,
				DataItem:    "RE-Data Item",
				Description: "Reserved Expansion Field",
				Type:        REField,
			},
		},
		// FX
	},
}
