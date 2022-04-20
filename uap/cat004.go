package uap

// Cat004V112 User Application Profile
// version 1.12
var Cat004V112 = StandardUAP{
	Name:     "cat004_1.12",
	Category: 4,
	Version:  1.12,
	Items: []DataField{
		{
			FRN:         1,
			DataItem:    "I004/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I004/000",
			Description: "Message Type",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         3,
			DataItem:    "I004/015",
			Description: "SDPS Identifier",
			Type:        Repetitive,
			Size: Size{
				ForRepetitive: 2,
			},
		},
		{
			FRN:         4,
			DataItem:    "I004/020",
			Description: "Time Of Message",
			Type:        Fixed,
			Size: Size{
				ForFixed: 3,
			},
		},
		{
			FRN:         5,
			DataItem:    "I004/040",
			Description: "Alert Identifier",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         6,
			DataItem:    "I004/045",
			Description: "Alert Status",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         7,
			DataItem:    "I004/060",
			Description: "Safety Net Function & System Status",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		// FX
		{
			FRN:         8,
			DataItem:    "I004/030",
			Description: "Track Number 1",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         9,
			DataItem:    "I004/170",
			Description: "Aircraft Identification & Characteristics 1",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "AI1",
					Description: "Aircraft Identifier 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         2,
					DataItem:    "M31",
					Description: "Mode 3/A Code Aircraft 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         3,
					DataItem:    "CPW",
					Description: "Predicted Conflict Position 1 (WGS84)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 10,
					},
				},
				{
					FRN:         4,
					DataItem:    "CPC",
					Description: "Predicted Conflict Position 1 (Cartesian Coordinates)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 8,
					},
				},
				{
					FRN:         5,
					DataItem:    "TT1",
					Description: "Time to Threshold Aircraft 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         6,
					DataItem:    "DT1",
					Description: "Distance to Threshold Aircraft 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         7,
					DataItem:    "AC1",
					Description: "Aircraft Characteristics Aircraft 1",
					Type:        Extended,
					Size: Size{
						ForExtendedPrimary:   1,
						ForExtendedSecondary: 1,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "MS1",
					Description: "Mode S Identifier Aircraft 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         9,
					DataItem:    "FP1",
					Description: "Flight Plan Number Aircraft 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         10,
					DataItem:    "CF1",
					Description: "Cleared Flight Level Aircraft 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
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
					FRN:  13,
					Type: Spare,
				},
				{
					FRN:  14,
					Type: Spare,
				},
				// FX
			},
		},
		{
			FRN:         10,
			DataItem:    "I004/120",
			Description: "Conflict Characteristics",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "CN",
					Description: "Conflict Nature",
					Type:        Extended,
					Size: Size{
						ForExtendedPrimary:   1,
						ForExtendedSecondary: 1,
					},
				},
				{
					FRN:         2,
					DataItem:    "CC",
					Description: "Conflict Classification",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         3,
					DataItem:    "CP",
					Description: "Conflict Probability",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         4,
					DataItem:    "CD",
					Description: "Conflict Duration",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
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
				// FX
			},
		},
		{
			FRN:         11,
			DataItem:    "I004/070",
			Description: "Conflict Timing and Separation",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "TC",
					Description: "Time to Conflict",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         2,
					DataItem:    "TCA",
					Description: "Time to Closest Approach",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         3,
					DataItem:    "CHS",
					Description: "Current Horizontal Separation",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         4,
					DataItem:    "MHS",
					Description: "Estimated Minimum Horizontal Separation",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         5,
					DataItem:    "CVS",
					Description: "Current Vertical Separation",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         6,
					DataItem:    "MVS",
					Description: "Estimated Minimum Vertical Separation",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:  7,
					Type: Spare,
				},
				// FX
			},
		},
		{
			FRN:         12,
			DataItem:    "I004/076",
			Description: "Vertical Deviation",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         13,
			DataItem:    "I004/074",
			Description: "Longitudinal Deviation",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         14,
			DataItem:    "I004/075",
			Description: "Transversal Distance Deviation",
			Type:        Fixed,
			Size: Size{
				ForFixed: 3,
			},
		},
		// FX
		{
			FRN:         15,
			DataItem:    "I004/100",
			Description: "Area Definitions",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "AN",
					Description: "Area Name",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         2,
					DataItem:    "CAN",
					Description: "Crossing Area Name",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         3,
					DataItem:    "RT1",
					Description: "Runway/Taxiway Designator 1",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         4,
					DataItem:    "RT2",
					Description: "Runway/Taxiway Designator 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         5,
					DataItem:    "SB",
					Description: "Stop Bar Designator",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         6,
					DataItem:    "G",
					Description: "Gate Designator",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:  7,
					Type: Spare,
				},
				// FX
			},
		},
		{
			FRN:         16,
			DataItem:    "I004/035",
			Description: "Track Number 2",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         17,
			DataItem:    "I004/171",
			Description: "Aircraft Identification & Characteristics 2",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "AI2",
					Description: "Aircraft Identifier 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         2,
					DataItem:    "M32",
					Description: "Mode 3/A Code Aircraft 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         3,
					DataItem:    "CPW",
					Description: "Predicted Conflict Position 2 (WGS84)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 10,
					},
				},
				{
					FRN:         4,
					DataItem:    "CPC",
					Description: "Predicted Conflict Position 2 (Cartesian Coordinates)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 8,
					},
				},
				{
					FRN:         5,
					DataItem:    "TT2",
					Description: "Time to Threshold Aircraft 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         6,
					DataItem:    "DT2",
					Description: "Distance to Threshold Aircraft 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         7,
					DataItem:    "AC2",
					Description: "Aircraft Characteristics Aircraft 2",
					Type:        Extended,
					Size: Size{
						ForExtendedPrimary:   1,
						ForExtendedSecondary: 1,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "MS2",
					Description: "Mode S Identifier Aircraft 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         9,
					DataItem:    "FP2",
					Description: "Flight Plan Number Aircraft 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         10,
					DataItem:    "CF2",
					Description: "Cleared Flight Level Aircraft 2",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
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
					FRN:  13,
					Type: Spare,
				},
				{
					FRN:  14,
					Type: Spare,
				},
				// FX
			},
		},
		{
			FRN:         18,
			DataItem:    "I004/110",
			Description: "FDPS Sector Control Identifier",
			Type:        Repetitive,
			Size: Size{
				ForRepetitive: 2,
			},
		},
		{
			FRN:      19,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:         20,
			DataItem:    "RE",
			Description: "Reserved Expansion Field",
			Type:        RE,
		},
		{
			FRN:         21,
			DataItem:    "SP",
			Description: "Special Purpose Field",
			Type:        SP,
		},
		// FX
	},
}
