package uap

// Cat062V119 User Application Profile
// version 1.19
var Cat062V119 = StandardUAP{
	Name:     "cat062_1.19",
	Category: 62,
	Version:  1.19,
	Items: []DataField{
		{
			FRN:         1,
			DataItem:    "I062/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:      2,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:         3,
			DataItem:    "I062/015",
			Description: "Service Identification",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         4,
			DataItem:    "I062/070",
			Description: "Time Of Track Information",
			Type:        Fixed,
			Size: Size{
				ForFixed: 3,
			},
		},
		{
			FRN:         5,
			DataItem:    "I062/105",
			Description: "Calculated Track Position (WGS-84)",
			Type:        Fixed,
			Size: Size{
				ForFixed: 8,
			},
		},
		{
			FRN:         6,
			DataItem:    "I062/100",
			Description: "Calculated Track Position (Cartesian)",
			Type:        Fixed,
			Size: Size{
				ForFixed: 6,
			},
		},
		{
			FRN:         7,
			DataItem:    "I062/185",
			Description: "Calculated Track Velocity (Cartesian)",
			Type:        Fixed,
			Size: Size{
				ForFixed: 4,
			},
		},
		{
			FRN:         8,
			DataItem:    "I062/210",
			Description: "Calculated Acceleration (Cartesian)",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:         9,
			DataItem:    "I062/060",
			Description: "Track Mode 3/A Code",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         10,
			DataItem:    "I062/245",
			Description: "Target Identification",
			Type:        Fixed,
			Size: Size{
				ForFixed: 7,
			},
		},
		{
			FRN:         11,
			DataItem:    "I062/380",
			Description: "Aircraft Derived Data",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "ADR",
					Description: "Target Adress",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         2,
					DataItem:    "ID",
					Description: "Target Identification",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         3,
					DataItem:    "MHG",
					Description: "Magnetic Heading",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         4,
					DataItem:    "IAS",
					Description: "Indicated Airspeed",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         5,
					DataItem:    "TAS",
					Description: "True Airspeed",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         6,
					DataItem:    "SAL",
					Description: "Selected Altitude",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         7,
					DataItem:    "FSS",
					Description: "Final State Selected Altitude",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "TIS",
					Description: "Trajectory Intent Status",
					Type:        Extended,
					Size: Size{
						ForExtendedPrimary:   1,
						ForExtendedSecondary: 1,
					},
				},
				{
					FRN:         9,
					DataItem:    "TID",
					Description: "Trajectory Intent Data",
					Type:        Repetitive,
					Size: Size{
						ForRepetitive: 15,
					},
				},
				{
					FRN:         10,
					DataItem:    "COM",
					Description: "Communications / ACAS",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         11,
					DataItem:    "SAB",
					Description: "Status Reported by ADS-B",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         12,
					DataItem:    "ACS",
					Description: "ACAS Resolution Advisory Reports",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         13,
					DataItem:    "BVR",
					Description: "Barometric Vertical Rate",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         14,
					DataItem:    "GVR",
					Description: "Geometric Vertical Rate",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				// FX
				{
					FRN:         15,
					DataItem:    "RAN",
					Description: "Roll Angle",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         16,
					DataItem:    "TAR",
					Description: "Track Angle Rate",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         17,
					DataItem:    "TAN",
					Description: "Track Angle",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         18,
					DataItem:    "GSP",
					Description: "Ground Speed",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         19,
					DataItem:    "VUN",
					Description: "Velocity Uncertainty",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         20,
					DataItem:    "MET",
					Description: "Meteorological Data",
					Type:        Fixed,
					Size: Size{
						ForFixed: 8,
					},
				},
				{
					FRN:         21,
					DataItem:    "EMC",
					Description: "Emitter Category",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         22,
					DataItem:    "POS",
					Description: "Position Data",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         23,
					DataItem:    "GAL",
					Description: "Geometric Altitude Data",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         24,
					DataItem:    "PUN",
					Description: "Position Uncertainty Data",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         25,
					DataItem:    "MB",
					Description: "Mode S MB Data",
					Type:        Repetitive,
					Size: Size{
						ForRepetitive: 8,
					},
				},
				{
					FRN:         26,
					DataItem:    "IAR",
					Description: "Indicated Airspeed",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         27,
					DataItem:    "MAC",
					Description: "Mach Number",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         28,
					DataItem:    "BPS",
					Description: "Barometric Pressure Setting",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
			},
		},
		{
			FRN:         12,
			DataItem:    "I062/040",
			Description: "Track Number",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         13,
			DataItem:    "I062/080",
			Description: "Track Status",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         14,
			DataItem:    "I062/290",
			Description: "System Track Update Ages",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "TRK",
					Description: "Track age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         2,
					DataItem:    "PSR",
					Description: "PSR age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         3,
					DataItem:    "SSR",
					Description: "SSR age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         4,
					DataItem:    "MDS",
					Description: "Mode S age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "ADS",
					Description: "ADS-C age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         6,
					DataItem:    "ES",
					Description: "ADS-B Extended Squitter age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         7,
					DataItem:    "VDL",
					Description: "ADS-B VDL Mode 4 age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "UAT",
					Description: "ADS-B UAT age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         9,
					DataItem:    "LOP",
					Description: "Loop age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         10,
					DataItem:    "MLT",
					Description: "Multilateration age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
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
		//FX : Field Extension Indicator
		{
			FRN:         15,
			DataItem:    "I062/200",
			Description: "Mode of Movement",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         16,
			DataItem:    "I062/295",
			Description: "Track Data Ages",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "MFL",
					Description: "Measured Flight Level age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         2,
					DataItem:    "MD1",
					Description: "Mode 1 age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         3,
					DataItem:    "MD2",
					Description: "Mode 2 age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         4,
					DataItem:    "MDA",
					Description: "Mode 3/A age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "MD4",
					Description: "True Mode 4 age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         6,
					DataItem:    "MD5",
					Description: "Mode 5 age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         7,
					DataItem:    "MHG",
					Description: "Magnetic Heading age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "IAS",
					Description: "Indicated Airspeed/Mach Nb age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         9,
					DataItem:    "TAS",
					Description: "True Airspeed age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         10,
					DataItem:    "SAL",
					Description: "Selected ALtitude Age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         11,
					DataItem:    "FSS",
					Description: "Final State Slected Altitude Age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         12,
					DataItem:    "COM",
					Description: "Communications / ACAS Capability and Flight Status age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         13,
					DataItem:    "TID",
					Description: "Trajectory Intent Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         14,
					DataItem:    "SAB",
					Description: "Status Reported by ADS-B age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         15,
					DataItem:    "ACS",
					Description: "ACAS Resolution Advisory Report age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         16,
					DataItem:    "BVR",
					Description: "Barometric Vertical Rate age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         17,
					DataItem:    "GVR",
					Description: "Geometric Vertical Rate age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         18,
					DataItem:    "RAN",
					Description: "Roll Angle age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         19,
					DataItem:    "TAR",
					Description: "Track Angle Rate age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         20,
					DataItem:    "TAN",
					Description: "Track Angle age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         21,
					DataItem:    "GSP",
					Description: "Ground Speed age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         22,
					DataItem:    "VUN",
					Description: "Velocity Uncertainity age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         23,
					DataItem:    "MET",
					Description: "Meteorological Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         24,
					DataItem:    "EMC",
					Description: "Emitter Category age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         25,
					DataItem:    "POS",
					Description: "Position Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         26,
					DataItem:    "GAL",
					Description: "Geometric Altitude Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         27,
					DataItem:    "PUN",
					Description: "Position Uncertainty Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         28,
					DataItem:    "MB",
					Description: "Mode S MB Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         29,
					DataItem:    "IAR",
					Description: "Indicated Airspeed Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         30,
					DataItem:    "MAC",
					Description: "Mac Number Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         31,
					DataItem:    "BPS",
					Description: "Barometric Pressure Setting Data age",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:  32,
					Type: Spare,
				},
				{
					FRN:  33,
					Type: Spare,
				},
				{
					FRN:  34,
					Type: Spare,
				},
				{
					FRN:  35,
					Type: Spare,
				},
				// FX
			},
		},
		{
			FRN:         17,
			DataItem:    "I062/136",
			Description: "Measured Flight Level",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         18,
			DataItem:    "I062/130",
			Description: "Calculated Track Geometric Altitude",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         19,
			DataItem:    "I062/135",
			Description: "Calculated Track Barometric Altitude",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         20,
			DataItem:    "I062/220",
			Description: "Calculated Rate Of Climb/Descent",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         21,
			DataItem:    "I062/390",
			Description: "Flight Plan Related Data",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "TAG",
					Description: "FPPS Identification Tag",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         2,
					DataItem:    "CSN",
					Description: "Callsign",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         3,
					DataItem:    "IFI",
					Description: "IFPS_FLIGHT_ID",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         4,
					DataItem:    "FCT",
					Description: "Flight Category",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "TAC",
					Description: "Type of Aircraft",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         6,
					DataItem:    "WTC",
					Description: "Wake Turbulence Category",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         7,
					DataItem:    "DEP",
					Description: "Departure Airport",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "DST",
					Description: "Destination Airport",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         9,
					DataItem:    "RDS",
					Description: "Runaway Designation",
					Type:        Fixed,
					Size: Size{
						ForFixed: 3,
					},
				},
				{
					FRN:         10,
					DataItem:    "CFL",
					Description: "Current Cleared Flight Level",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         11,
					DataItem:    "CTL",
					Description: "Current Control Position",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         12,
					DataItem:    "TOD",
					Description: "Time Of Departure / Arrival",
					Type:        Repetitive,
					Size: Size{
						ForRepetitive: 4,
					},
				},
				{
					FRN:         13,
					DataItem:    "AST",
					Description: "Aircraft Stand",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         14,
					DataItem:    "STS",
					Description: "Stand Status",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				// FX
				{
					FRN:         15,
					DataItem:    "STD",
					Description: "Standard Instrument Departure",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         16,
					DataItem:    "STA",
					Description: "STandard Instrument Arrival",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:         17,
					DataItem:    "PEM",
					Description: "Pre-emergency Mode 3/A code",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         18,
					DataItem:    "PEC",
					Description: "Pre-emergency Callsign",
					Type:        Fixed,
					Size: Size{
						ForFixed: 7,
					},
				},
				{
					FRN:  19,
					Type: Spare,
				},
				{
					FRN:  20,
					Type: Spare,
				},
				{
					FRN:  21,
					Type: Spare,
				},
				// FX
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:         22,
			DataItem:    "I062/270",
			Description: "Target Size & Orientation",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         23,
			DataItem:    "I062/300",
			Description: "Vehicle Fleet Identification",
			Type:        Fixed,
			Size: Size{
				ForFixed: 1,
			},
		},
		{
			FRN:         24,
			DataItem:    "I062/110",
			Description: "Mode 5 Data reports & Extended Mode 1 Code",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "SUM",
					Description: "SMode 5 Summary",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         2,
					DataItem:    "PMN",
					Description: "Mode 5 PIN/ National Origin/ Mission Code",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         3,
					DataItem:    "POS",
					Description: "Mode 5 Reported Position",
					Type:        Fixed,
					Size: Size{
						ForFixed: 6,
					},
				},
				{
					FRN:         4,
					DataItem:    "GA",
					Description: "Mode 5 GNSS-derived Altitude",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         5,
					DataItem:    "EM1",
					Description: "Extended Mode 1 Code in Octal Representation",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         6,
					DataItem:    "TOS",
					Description: "Time Offset for POS and GA",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         7,
					DataItem:    "XP",
					Description: "X Pulse Presence",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
			},
		},
		{
			FRN:         25,
			DataItem:    "I062/120",
			Description: "Track Mode 2 Code",
			Type:        Fixed,
			Size: Size{
				ForFixed: 2,
			},
		},
		{
			FRN:         26,
			DataItem:    "I062/510",
			Description: "Composed Track Number",
			Type:        Extended,
			Size: Size{
				ForExtendedPrimary:   3,
				ForExtendedSecondary: 3,
			},
		},
		{
			FRN:         27,
			DataItem:    "I062/500",
			Description: "Estimated Accuracies",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "APC",
					Description: "Accuracy Of Track Position (Cartesian)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         2,
					DataItem:    "COV",
					Description: "XY Covariance",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         3,
					DataItem:    "APW",
					Description: "Estimated Accuracy Of Track Position (WGS-84)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         4,
					DataItem:    "AGA",
					Description: "Estimated Accuracy Of Calculated Track Geometric",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         5,
					DataItem:    "ABA",
					Description: "Estimated Accuracy Of Calculated Track Barometric",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:         6,
					DataItem:    "ATV",
					Description: "Estimated Accuracy Of Track Velocity (Cartesian)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         7,
					DataItem:    "AA",
					Description: "Estimated Accuracy Of Acceleration (Cartesian)",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				// FX
				{
					FRN:         8,
					DataItem:    "ARC",
					Description: "Estimated Accuracy Of Rate Of Climb/Descent",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
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
			FRN:         28,
			DataItem:    "I062/340",
			Description: "Measured Information",
			Type:        Compound,
			Compound: []DataField{
				{
					FRN:         1,
					DataItem:    "SID",
					Description: "Sensor Identification",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         2,
					DataItem:    "POS",
					Description: "Measured Position",
					Type:        Fixed,
					Size: Size{
						ForFixed: 4,
					},
				},
				{
					FRN:         3,
					DataItem:    "HEI",
					Description: "Measured 3-D Height",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         4,
					DataItem:    "MDC",
					Description: "Last Measured Mode C code",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         5,
					DataItem:    "MDA",
					Description: "Last Measured Mode 3/A code",
					Type:        Fixed,
					Size: Size{
						ForFixed: 2,
					},
				},
				{
					FRN:         6,
					DataItem:    "TYP",
					Description: "Report Type",
					Type:        Fixed,
					Size: Size{
						ForFixed: 1,
					},
				},
				{
					FRN:  7,
					Type: Spare,
				},
				// FX
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:      29,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      30,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      31,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      32,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      33,
			DataItem: "NA",
			Type:     Spare,
		},
		{
			FRN:      34,
			DataItem: "RE-Data Item",
			Type:     RE,
		},
		{
			FRN:      35,
			DataItem: "SP-Data Item",
			Type:     SP,
		},
		//FX : Field Extension Indicator
	},
}
