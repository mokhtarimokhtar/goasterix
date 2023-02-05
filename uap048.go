package goasterix

import "github.com/mokhtarimokhtar/goasterix/item"

const (
	I048010Name = "I048/010"
	I048010Desc = "Data Source Identifier"
	I048140Name = "I048/140"
	I048140Desc = "Time of Day"
	I048020Name = "I048/020"
	I048020Desc = "Target Report Descriptor"

	I048130Name    = "I048/130"
	I048130Desc    = "Radar Plot Characteristics"
	I048130SRL     = "SRL"
	I048130SRLDesc = "SSR plot runlength"
	I048130SRR     = "SRR"
	I048130SRRDesc = "Number of received replies"
	I048130SAM     = "SAM"
	I048130SAMDesc = "Amplitude of received replies for M(SSR)"
	I048130PRL     = "PRL"
	I048130PRLDesc = "PSR plot runlength"
	I048130PAM     = "PAM"
	I048130PAMDesc = "PSR amplitude"
	I048130RPD     = "RPD"
	I048130RPDDesc = "Difference in Range between PSR and SSR plot"
	I048130APD     = "APD"
	I048130APDDesc = "Difference in Azimuth between PSR and SSR plot"

	I048240Name  = "I048/240"
	I048240Desc  = "Aircraft Identification"
	I048240Char1 = "char1"
	I048240Char2 = "char2"
	I048240Char3 = "char3"
	I048240Char4 = "char4"
	I048240Char5 = "char5"
	I048240Char6 = "char6"
	I048240Char7 = "char7"
	I048240Char8 = "char8"
)

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
				DataItemName: I048010Name,
				Description:  I048010Desc,
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
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
				DataItemName: I048140Name,
				Description:  I048140Desc,
				Type:         item.FixedField,
			},
			Size: 3,
			SubItems: []item.SubItem{
				{
					Type: item.FromToField,
					From: 24, To: 1,
				},
			},
		},
		&item.Extended{
			Base: item.Base{
				FRN:          item.FRN3,
				DataItemName: I048020Name,
				Description:  I048020Desc,
				Type:         item.ExtendedField,
			},
			PrimaryItemSize:   1,
			SecondaryItemSize: 1,
			SubItems: []item.SubItem{
				{Name: "TYP", Type: item.FromToField, From: 8, To: 6},
				{Name: "SIM", Type: item.BitField, Bit: 5},
				{Name: "RDP", Type: item.BitField, Bit: 4},
				{Name: "SPI", Type: item.BitField, Bit: 3},
				{Name: "RAB", Type: item.BitField, Bit: 2},
				{Name: "FX", Type: item.BitField, Bit: 1},
				{Name: "TST", Type: item.BitField, Bit: 8},
				{Name: "ERR", Type: item.BitField, Bit: 7},
				{Name: "XPP", Type: item.BitField, Bit: 6},
				{Name: "ME", Type: item.BitField, Bit: 5},
				{Name: "MI", Type: item.BitField, Bit: 4},
				{Name: "FOE/FRI", Type: item.FromToField, From: 3, To: 2},
				{Name: "FX", Type: item.BitField, Bit: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN4,
				DataItemName: "I048/040",
				Description:  "Measured Position in Slant Polar Coordinates",
				Type:         item.FixedField,
			},
			Size: 4,
			SubItems: []item.SubItem{
				{Name: "RHO", Type: item.FromToField, From: 32, To: 17},
				{Name: "THETA", Type: item.FromToField, From: 16, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN5,
				DataItemName: "I048/070",
				Description:  "Mode-3/A Code in Octal Representation",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "V", Type: item.BitField, Bit: 16},
				{Name: "G", Type: item.BitField, Bit: 15},
				{Name: "L", Type: item.BitField, Bit: 14},
				{Name: "Mode-3/A", Type: item.FromToField, From: 12, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN6,
				DataItemName: "I048/090",
				Description:  "Flight Level in Binary Representation",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "V", Type: item.BitField, Bit: 16},
				{Name: "G", Type: item.BitField, Bit: 15},
				{Name: "LEVEL", Type: item.FromToField, From: 14, To: 1},
			},
		},
		&item.Compound{
			Base: item.Base{
				FRN:          item.FRN7,
				DataItemName: I048130Name,
				Description:  I048130Desc,
				Type:         item.CompoundField,
			},
			Secondary: []item.DataItem{
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN1,
						DataItemName: I048130SRL,
						Description:  I048130SRLDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130SRL, Type: item.FromToField, From: 8, To: 1},
					},
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN2,
						DataItemName: I048130SRR,
						Description:  I048130SRRDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130SRR, Type: item.FromToField, From: 8, To: 1},
					},
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN3,
						DataItemName: I048130SAM,
						Description:  I048130SAMDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130SAM, Type: item.FromToField, From: 8, To: 1},
					},
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN4,
						DataItemName: I048130PRL,
						Description:  I048130PRLDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130PRL, Type: item.FromToField, From: 8, To: 1},
					},
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN5,
						DataItemName: I048130PAM,
						Description:  I048130PAMDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130PAM, Type: item.FromToField, From: 8, To: 1},
					},
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN6,
						DataItemName: I048130RPD,
						Description:  I048130RPDDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130RPD, Type: item.FromToField, From: 8, To: 1},
					},
				},
				&item.Fixed{
					Base: item.Base{
						FRN:          item.FRN7,
						DataItemName: I048130APD,
						Description:  I048130APDDesc,
						Type:         item.FixedField,
					},
					Size: 1,
					SubItems: []item.SubItem{
						{Name: I048130APD, Type: item.FromToField, From: 8, To: 1},
					},
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
			SubItems: []item.SubItem{
				{Name: "AIRCRAFT ADDRESS", Type: item.FromToField, From: 24, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN9,
				DataItemName: I048240Name,
				Description:  I048240Desc,
				Type:         item.FixedField,
			},
			Size: 6,
			SubItems: []item.SubItem{
				{Name: I048240Char1, Type: item.FromToField, From: 48, To: 43},
				{Name: I048240Char2, Type: item.FromToField, From: 42, To: 37},
				{Name: I048240Char3, Type: item.FromToField, From: 36, To: 31},
				{Name: I048240Char4, Type: item.FromToField, From: 30, To: 25},
				{Name: I048240Char5, Type: item.FromToField, From: 24, To: 19},
				{Name: I048240Char6, Type: item.FromToField, From: 18, To: 13},
				{Name: I048240Char7, Type: item.FromToField, From: 12, To: 7},
				{Name: I048240Char8, Type: item.FromToField, From: 6, To: 1},
			},
		},
		&item.Repetitive{
			Base: item.Base{
				FRN:          item.FRN10,
				DataItemName: "I048/250",
				Description:  "Mode S MB Data",
				Type:         item.RepetitiveField,
			},
			SubItemSize: 8,
			SubItems: []item.SubItem{
				{Name: "MBData", Type: item.FromToField, From: 64, To: 1},
				//{Name: "MBData", Type: item.FromToField, From: 64, To: 9},
				//{Name: "BDS1", Type: item.FromToField, From: 8, To: 5},
				//{Name: "BDS2", Type: item.FromToField, From: 4, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN11,
				DataItemName: "I048/161",
				Description:  "Track Number",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "TRACK NUMBER", Type: item.FromToField, From: 12, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN12,
				DataItemName: "I048/042",
				Description:  "Calculated Position in Cartesian Coordinates",
				Type:         item.FixedField,
			},
			Size: 4,
			SubItems: []item.SubItem{
				{Name: "X-Component", Type: item.FromToField, From: 32, To: 17},
				{Name: "Y-Component", Type: item.FromToField, From: 16, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN13,
				DataItemName: "I048/200",
				Description:  "Calculated Track Velocity in Polar Representation",
				Type:         item.FixedField,
			},
			Size: 4,
			SubItems: []item.SubItem{
				{Name: "GROUNDSPEED", Type: item.FromToField, From: 32, To: 17},
				{Name: "HEADING", Type: item.FromToField, From: 16, To: 1},
			},
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
			SubItems: []item.SubItem{
				{Name: "CNF", Type: item.BitField, Bit: 8},
				{Name: "RAD", Type: item.FromToField, From: 7, To: 6},
				{Name: "DOU", Type: item.BitField, Bit: 5},
				{Name: "MAH", Type: item.BitField, Bit: 4},
				{Name: "CDM", Type: item.FromToField, From: 3, To: 2},
				{Name: "FX", Type: item.BitField, Bit: 1},

				{Name: "TRE", Type: item.BitField, Bit: 8},
				{Name: "GHO", Type: item.BitField, Bit: 7},
				{Name: "SUP", Type: item.BitField, Bit: 6},
				{Name: "TCC", Type: item.BitField, Bit: 5},
				{Name: "FX", Type: item.BitField, Bit: 1},
			},
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
			SubItems: []item.SubItem{
				{Name: "SigmaX", Type: item.FromToField, From: 32, To: 25},
				{Name: "SigmaY", Type: item.FromToField, From: 24, To: 17},
				{Name: "SigmaV", Type: item.FromToField, From: 16, To: 9},
				{Name: "SigmaH", Type: item.FromToField, From: 8, To: 1},
			},
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
			SubItems: []item.SubItem{
				{Name: "Code", Type: item.FromToField, From: 8, To: 2},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN17,
				DataItemName: "I048/080",
				Description:  "Mode-3/A Code Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "QXi", Type: item.FromToField, From: 12, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN18,
				DataItemName: "I048/100",
				Description:  "Mode-C Code and Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 4,
			SubItems: []item.SubItem{
				{Name: "V", Type: item.BitField, Bit: 32},
				{Name: "G", Type: item.BitField, Bit: 31},
				{Name: "Mode-C", Type: item.FromToField, From: 28, To: 17},
				{Name: "QXi", Type: item.FromToField, From: 12, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN19,
				DataItemName: "I048/110",
				Description:  "Height Measured by 3D Radar",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "3D-Height", Type: item.FromToField, From: 14, To: 1},
			},
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
					SubItems: []item.SubItem{
						{Name: "D", Type: item.BitField, Bit: 16},
						{Name: "CAL", Type: item.FromToField, From: 10, To: 1},
					},
				},
				&item.Repetitive{
					Base: item.Base{
						FRN:          item.FRN2,
						DataItemName: "RDS",
						Description:  "Raw Doppler Speed",
						Type:         item.RepetitiveField,
					},
					SubItemSize: 6,
					SubItems: []item.SubItem{
						{Name: "DOP", Type: item.FromToField, From: 48, To: 33},
						{Name: "AMB", Type: item.FromToField, From: 32, To: 17},
						{Name: "FRQ", Type: item.FromToField, From: 16, To: 1},
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
			SubItems: []item.SubItem{
				{Name: "COM", Type: item.FromToField, From: 16, To: 14},
				{Name: "STAT", Type: item.FromToField, From: 13, To: 11},
				{Name: "SI", Type: item.BitField, Bit: 10},
				{Name: "MSSC", Type: item.BitField, Bit: 8},
				{Name: "ARC", Type: item.BitField, Bit: 7},
				{Name: "AIC", Type: item.BitField, Bit: 6},
				{Name: "B1A", Type: item.BitField, Bit: 5},
				{Name: "B1B", Type: item.FromToField, From: 4, To: 1},
			},
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
			SubItems: []item.SubItem{
				{Name: "ACAS-RA", Type: item.FromToField, From: 56, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN23,
				DataItemName: "I048/055",
				Description:  "Mode-1 Code in Octal Representation",
				Type:         item.FixedField,
			},
			Size: 1,
			SubItems: []item.SubItem{
				{Name: "V", Type: item.BitField, Bit: 8},
				{Name: "G", Type: item.BitField, Bit: 7},
				{Name: "L", Type: item.BitField, Bit: 6},
				{Name: "Mode-1", Type: item.FromToField, From: 5, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN24,
				DataItemName: "I048/050",
				Description:  "Mode-2 Code in Octal Representation",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "V", Type: item.BitField, Bit: 16},
				{Name: "G", Type: item.BitField, Bit: 15},
				{Name: "L", Type: item.BitField, Bit: 14},
				{Name: "Mode-2", Type: item.FromToField, From: 12, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN25,
				DataItemName: "I048/065",
				Description:  "Mode-1 Code Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 1,
			SubItems: []item.SubItem{
				{Name: "QXi", Type: item.FromToField, From: 5, To: 1},
			},
		},
		&item.Fixed{
			Base: item.Base{
				FRN:          item.FRN26,
				DataItemName: "I048/060",
				Description:  "Mode-2 Code Confidence Indicator",
				Type:         item.FixedField,
			},
			Size: 2,
			SubItems: []item.SubItem{
				{Name: "QXi", Type: item.FromToField, From: 12, To: 1},
			},
		},
		&item.SpecialPurpose{
			Base: item.Base{
				FRN:          item.FRN27,
				DataItemName: "SP-Data Item",
				Description:  "Special Purpose Field",
				Type:         item.SPField,
			},
		},
		&item.ReservedExpansion{
			Base: item.Base{
				FRN:          item.FRN28,
				DataItemName: "RE-Data Item",
				Description:  "Reserved Expansion Field",
				Type:         item.REField,
			},
		},
		// FX
	},
}
