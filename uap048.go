package goasterix

import "github.com/mokhtarimokhtar/goasterix/item"

// Cat048V127 User Application Profile
// version 1.27

var Cat048V127 = item.UAP{
	Name:     "cat048_1.27",
	Category: 48,
	Version:  1.27,
	DataItems: []item.DataItem{
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN1,
				DataItemName: "I048/010",
				Description:  "Data Source Identifier",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItemBits{
				{
					Name: "SAC",
					Type: item.FromToField,
					From: 16, To: 9,
				},
				{
					Name: "SIC",
					Type: item.FromToField,
					From: 8, To: 1,
				},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN2,
				DataItemName: "I048/140",
				Description:  "Time-of-Day",
				Type:         item.FixedField,
			},
			Size: 3,
		},
		&item.Extended{
			Base: item.Base{
				FRN:          item.FRN3,
				DataItemName: "I048/020",
				Description:  "Target Report Descriptor",
				Type:         item.ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN4,
				DataItemName: "I048/040",
				Description:  "Measured Position in Slant Polar Coordinates",
				Type:         item.FixedField,
			},
			Size: 4,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN5,
				DataItemName: "I048/070",
				Description:  "Mode-3/A Code in Octal Representation",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN6,
				DataItemName: "I048/090",
				Description:  "Flight Level in Binary Representation",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.Compound{
			Base: item.Base{
				FRN:          item.FRN7,
				DataItemName: "I048/130",
				Description:  "Radar Plot Characteristics",
				Type:         item.CompoundField,
			},
			Secondary: []item.DataItem{
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN1,
						DataItemName: "SRL",
						Description:  "SSR plot runlength",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN2,
						DataItemName: "SRR",
						Description:  "Number of received replies",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN3,
						DataItemName: "SAM",
						Description:  "Amplitude of received replies for M(SSR)",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN4,
						DataItemName: "PRL",
						Description:  "PSR plot runlength",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN5,
						DataItemName: "PAM",
						Description:  "PSR amplitude",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN6,
						DataItemName: "RPD",
						Description:  "Difference in Range between PSR and SSR plot",
						Type:         item.FixedField,
					},
					Size: 1,
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN7,
						DataItemName: "APD",
						Description:  "Difference in Azimuth between PSR and SSR plot",
						Type:         item.FixedField,
					},
					Size: 1,
				},
			},
		},
		// FX
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN8,
				DataItemName: "I048/220",
				Description:  "Aircraft Address",
				Type:         item.FixedField,
			},
			Size: 3,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN9,
				DataItemName: "I048/240",
				Description:  "Aircraft Identification",
				Type:         item.FixedField,
			},
			Size: 6,
		},
		&item.Repetitive{
			Base: item.Base{
				FRN:          item.FRN10,
				DataItemName: "I048/250",
				Description:  "Mode S MB Data",
				Type:         item.RepetitiveField,
			},
			SubItemSize: 8,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN11,
				DataItemName: "I048/161",
				Description:  "Track Number",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN12,
				DataItemName: "I048/042",
				Description:  "Calculated Position in Cartesian Coordinates",
				Type:         item.FixedField,
			},
			Size: 4,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN13,
				DataItemName: "I048/200",
				Description:  "Calculated Track Velocity in Polar Representation",
				Type:         item.FixedField,
			},
			Size: 4,
		},
		&item.Extended{
			Base: item.Base{
				FRN:          item.FRN14,
				DataItemName: "I048/170",
				Description:  "Track Status",
				Type:         item.ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
		},
		// FX
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN15,
				DataItemName: "I048/210",
				Description:  "Track Quality",
				Type:         item.FixedField,
			},
			Size: 4,
		},
		&item.Extended{
			Base: item.Base{
				FRN:          item.FRN16,
				DataItemName: "I048/030",
				Description:  "Warning/Error Conditions/Target Classification",
				Type:         item.ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN17,
				DataItemName: "I048/080",
				Description:  "Mode-3/A Code Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN18,
				DataItemName: "I048/100",
				Description:  "Mode-C Code and Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 4,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN19,
				DataItemName: "I048/110",
				Description:  "Height Measured by 3D Radar",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.Compound{
			Base: item.Base{
				FRN:          item.FRN20,
				DataItemName: "I048/120",
				Description:  "Radial Doppler Speed",
				Type:         item.CompoundField,
			},
			Secondary: []item.DataItem{
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN1,
						DataItemName: "CAL",
						Description:  "Calculated Doppler Speed",
						Type:         item.FixedField,
					},
					Size: 2,
				},
				&item.Repetitive{
					Base: item.Base{
						FRN:          item.FRN2,
						DataItemName: "RDS",
						Description:  "Raw Doppler Speed",
						Type:         item.RepetitiveField,
					},
					SubItemSize: 6,
					SubItems: []item.SubItemBits{
						{
							Name: "DOP",
							Type: item.FromToField,
							From: 48, To: 33,
						},
						{
							Name: "AMB",
							Type: item.FromToField,
							From: 32, To: 17,
						},
						{
							Name: "FRQ",
							Type: item.FromToField,
							From: 16, To: 1,
						},
					},
				},
				&item.Spare{Base: item.Base{FRN: item.FRN3}},
				&item.Spare{Base: item.Base{FRN: item.FRN4}},
				&item.Spare{Base: item.Base{FRN: item.FRN5}},
				&item.Spare{Base: item.Base{FRN: item.FRN6}},
				&item.Spare{Base: item.Base{FRN: item.FRN7}},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN21,
				DataItemName: "I048/230",
				Description:  "Communications / ACAS Capability and Flight Status",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		// FX
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN22,
				DataItemName: "I048/260",
				Description:  "ACAS Resolution Advisory Report",
				Type:         item.FixedField,
			},
			Size: 7,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN23,
				DataItemName: "I048/055",
				Description:  "Mode-1 Code in Octal Representation",
				Type:         item.FixedField,
			},
			Size: 1,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN24,
				DataItemName: "I048/050",
				Description:  "Mode-2 Code in Octal Representation",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN25,
				DataItemName: "I048/065",
				Description:  "Mode-1 Code Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 1,
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN26,
				DataItemName: "I048/060",
				Description:  "Mode-2 Code Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 2,
		},
		&item.SpecialPurpose{
			Base: item.Base{
				FRN:          item.FRN27,
				DataItemName: "I060SPName-Data DataItemName",
				Description:  "Special Purpose Field",
				Type:         item.SPField,
			},
		},
		&item.ReservedExpansion{
			Base: item.Base{
				FRN:          item.FRN28,
				DataItemName: "I060REName-Data DataItemName",
				Description:  "Reserved Expansion Field",
				Type:         item.REField,
			},
		},
		// FX
	},
}
