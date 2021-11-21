package uap

// Cat062V119 User Application Profile
// version 1.19
var Cat062V119 = StandardUAP{
	Name:     "cat062_1.19",
	Category: 62,
	Version:  1.19,
	Items: []DataField{
		{
			FRN:      1,
			DataItem: "I062/010",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:      2,
			DataItem: "NA",
			Type:     TypeField{NameType: Spare},
		},
		{
			FRN:      3,
			DataItem: "I062/015",
			Type: TypeField{
				NameType: Fixed,
				Size:     1,
			},
		},
		{
			FRN:      4,
			DataItem: "I062/070",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		{
			FRN:      5,
			DataItem: "I062/105",
			Type: TypeField{
				NameType: Fixed,
				Size:     8,
			},
		},
		{
			FRN:      6,
			DataItem: "I062/100",
			Type: TypeField{
				NameType: Fixed,
				Size:     6,
			},
		},
		{
			FRN:      7,
			DataItem: "I062/185",
			Type: TypeField{
				NameType: Fixed,
				Size:     4,
			},
		},
		{
			FRN:      8,
			DataItem: "I062/210",
			Type: TypeField{
				NameType: Fixed,
				Size:     3,
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:      9,
			DataItem: "I062/060",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:      10,
			DataItem: "I062/245",
			Type: TypeField{
				NameType: Fixed,
				Size:     7,
			},
		},
		{
			FRN:         11,
			DataItem:    "I062/380",
			Description: "Aircraft Derived Data",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        3,
							Item:        "ADR",
							Description: "Target Adress",
						},
						7: {
							NameType:    Fixed,
							Size:        6,
							Item:        "ID",
							Description: "Target Identification",
						},
						6: {
							NameType:    Fixed,
							Size:        2,
							Item:        "MHG",
							Description: "Magnetic Heading",
						},
						5: {
							NameType:    Fixed,
							Size:        2,
							Item:        "IAS",
							Description: "Indicated Airspeed",
						},
						4: {
							NameType:    Fixed,
							Size:        2,
							Item:        "TAS",
							Description: "True Airspeed",
						},
						3: {
							NameType:    Fixed,
							Size:        2,
							Item:        "SAL",
							Description: "Selected Altitude",
						},
						2: {
							NameType:    Fixed,
							Size:        2,
							Item:        "FSS",
							Description: "Final State Selected Altitude",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:      Extended,
							PrimarySize:   1,
							SecondarySize: 1,
							Item:          "TIS",
							Description:   "Trajectory Intent Status",
						},
						7: {
							NameType:    Repetitive,
							Size:        15,
							Item:        "TID",
							Description: "Trajectory Intent Data",
						},
						6: {
							NameType:    Fixed,
							Size:        2,
							Item:        "COM",
							Description: "Communications / ACAS",
						},
						5: {
							NameType:    Fixed,
							Size:        2,
							Item:        "SAB",
							Description: "Status Reported by ADS-B",
						},
						4: {
							NameType:    Fixed,
							Size:        7,
							Item:        "ACS",
							Description: "ACAS Resolution Advisory Reports",
						},
						3: {
							NameType:    Fixed,
							Size:        2,
							Item:        "BVR",
							Description: "Barometric Vertical Rate",
						},
						2: {
							NameType:    Fixed,
							Size:        2,
							Item:        "GVR",
							Description: "Geometric Vertical Rate",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        2,
							Item:        "RAN",
							Description: "Roll Angle",
						},
						7: {
							NameType:    Fixed,
							Size:        2,
							Item:        "TAR",
							Description: "Track Angle Rate",
						},
						6: {
							NameType:    Fixed,
							Size:        2,
							Item:        "TAN",
							Description: "Track Angle",
						},
						5: {
							NameType:    Fixed,
							Size:        2,
							Item:        "GSP",
							Description: "Ground Speed",
						},
						4: {
							NameType:    Fixed,
							Size:        1,
							Item:        "VUN",
							Description: "Velocity Uncertainty",
						},
						3: {
							NameType:    Fixed,
							Size:        8,
							Item:        "MET",
							Description: "Meteorological Data",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "EMC",
							Description: "Emitter Category",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        6,
							Item:        "POS",
							Description: "Position Data",
						},
						7: {
							NameType:    Fixed,
							Size:        2,
							Item:        "GAL",
							Description: "Geometric Altitude Data",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "PUN",
							Description: "Position Uncertainty Data",
						},
						5: {
							NameType:    Repetitive,
							Size:        8,
							Item:        "MB",
							Description: "Mode S MB Data",
						},
						4: {
							NameType:    Fixed,
							Size:        2,
							Item:        "IAR",
							Description: "Indicated Airspeed",
						},
						3: {
							NameType:    Fixed,
							Size:        2,
							Item:        "MAC",
							Description: "Mach Number",
						},
						2: {
							NameType:    Fixed,
							Size:        2,
							Item:        "BPS",
							Description: "Barometric Pressure Setting",
						},
						//  1:FX
					},
				},
			},
		},
		{
			FRN:      12,
			DataItem: "I062/040",
			Type: TypeField{
				NameType: Fixed,
				Size:     2,
			},
		},
		{
			FRN:      13,
			DataItem: "I062/080",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:         14,
			DataItem:    "I062/290",
			Description: "System Track Update Ages",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TRK",
							Description: "Track age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "PSR",
							Description: "PSR age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "SSR",
							Description: "SSR age",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MDS",
							Description: "Mode S age",
						},
						4: {
							NameType:    Fixed,
							Size:        2,
							Item:        "ADS",
							Description: "ADS-C age",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "ES",
							Description: "ADS-B Extended Squitter age",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "VDL",
							Description: "ADS-B VDL Mode 4 age",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "UAT",
							Description: "ADS-B UAT age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "LOP",
							Description: "Loop age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MLT",
							Description: "Multilateration age",
						},
						//bits-5/2 spare bits set to zero
						5: {
							NameType: Fixed,
							Size:     0,
						},
						4: {
							NameType: Fixed,
							Size:     0,
						},
						3: {
							NameType: Fixed,
							Size:     0,
						},
						2: {
							NameType: Fixed,
							Size:     0,
						},
						//  1:FX
					},
				},
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:      15,
			DataItem: "I062/200",
			Type: TypeField{
				NameType: Fixed,
				Size:     1},
		},
		{
			FRN:         16,
			DataItem:    "I062/295",
			Description: "Track Data Ages",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MFL",
							Description: "Measured Flight Level age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MD1",
							Description: "Mode 1 age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MD2",
							Description: "Mode 2 age",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MDA",
							Description: "Mode 3/A age",
						},
						4: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MD4",
							Description: "True Mode 4 age",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MD5",
							Description: "Mode 5 age",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MHG",
							Description: "Magnetic Heading age",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "IAS",
							Description: "Indicated Airspeed/Mach Nb age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TAS",
							Description: "True Airspeed age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "SAL",
							Description: "Selected ALtitude Age",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "FSS",
							Description: "Final State Slected Altitude Age",
						},
						4: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TID",
							Description: "Trajectory Intent Data age",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "COM",
							Description: "Communications / ACAS Capability and Flight Status age",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "SAB",
							Description: "Status Reported by ADS-B age",
						},
						//  1:FX

					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "ACS",
							Description: "ACAS Resolution Advisory Report age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "BVR",
							Description: "Barometric Vertical Rate age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "GVR",
							Description: "Geometric Vertical Rate age",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "RAN",
							Description: "Roll Angle age",
						},
						4: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TAR",
							Description: "Track Angle Rate age",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TAN",
							Description: "Track Angle age",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "GSP",
							Description: "Ground Speed age",
						},
						//  1:FX

					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "VUN",
							Description: "Velocity Uncertainity age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MET",
							Description: "Meteorological Data age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "EMC",
							Description: "Emitter Category age",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "POS",
							Description: "Position Data age",
						},
						4: {
							NameType:    Fixed,
							Size:        1,
							Item:        "GAL",
							Description: "Geometric Altitude Data age",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "PUN",
							Description: "Position Uncertainty Data age",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MB",
							Description: "Mode S MB Data age",
						},
						//  1:FX

					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "IAR",
							Description: "Indicated Airspeed Data age",
						},
						7: {
							NameType:    Fixed,
							Size:        1,
							Item:        "MAC",
							Description: "Mac Number Data age",
						},
						6: {
							NameType:    Fixed,
							Size:        1,
							Item:        "BPS",
							Description: "Barometric Pressure Setting Data age",
						},
						//bit-5/2 spare bits set to zero
						5: {
							NameType: Spare,
						},
						4: {
							NameType: Spare,
						},
						3: {
							NameType: Spare,
						},
						2: {
							NameType: Spare,
						},
						//  1:FX

					},
				},
			},
		},
		{
			FRN:      17,
			DataItem: "I062/136",
			Type: TypeField{
				NameType: Fixed,
				Size:     2},
		},
		{
			FRN:      18,
			DataItem: "I062/130",
			Type: TypeField{
				NameType: Fixed,
				Size:     2},
		},
		{
			FRN:      19,
			DataItem: "I062/135",
			Type: TypeField{
				NameType: Fixed,
				Size:     2},
		},
		{
			FRN:      20,
			DataItem: "I062/220",
			Type: TypeField{
				NameType: Fixed,
				Size:     2},
		},
		{
			FRN:         21,
			DataItem:    "I062/390",
			Description: "Flight Plan Related Data",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        2,
							Item:        "TAG",
							Description: "FPPS Identification Tag",
						},
						7: {
							NameType:    Fixed,
							Size:        7,
							Item:        "CSN",
							Description: "Callsign",
						},
						6: {
							NameType:    Fixed,
							Size:        4,
							Item:        "IFI",
							Description: "IFPS_FLIGHT_ID",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "FCT",
							Description: "Flight Category",
						},
						4: {
							NameType:    Fixed,
							Size:        4,
							Item:        "TAC",
							Description: "Type of Aircraft",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "WTC",
							Description: "Wake Turbulence Category",
						},
						2: {
							NameType:    Fixed,
							Size:        4,
							Item:        "DEP",
							Description: "Departure Airport",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        4,
							Item:        "DST",
							Description: "Destination Airport",
						},
						7: {
							NameType:    Fixed,
							Size:        3,
							Item:        "RDS",
							Description: "Runaway Designation",
						},
						6: {
							NameType:    Fixed,
							Size:        2,
							Item:        "CFL",
							Description: "Current Cleared Flight Level",
						},
						5: {
							NameType:    Fixed,
							Size:        2,
							Item:        "CTL",
							Description: "Current Control Position",
						},
						4: {
							NameType:    Repetitive,
							Size:        4,
							Item:        "TOD",
							Description: "Time Of Departure / Arrival",
						},
						3: {
							NameType:    Fixed,
							Size:        6,
							Item:        "AST",
							Description: "Aircraft Stand",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "STS",
							Description: "Stand Status",
						},
						//  1:FX

					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        7,
							Item:        "STD",
							Description: "Standard Instrument Departure",
						},
						7: {
							NameType:    Fixed,
							Size:        7,
							Item:        "STA",
							Description: "STandard Instrument Arrival",
						},
						6: {
							NameType:    Fixed,
							Size:        2,
							Item:        "PEM",
							Description: "Pre-emergency Mode 3/A code",
						},
						5: {
							NameType:    Fixed,
							Size:        7,
							Item:        "PEC",
							Description: "Pre-emergency Callsign",
						},
						//bits-4/2 Spare bits set to zero
						4: {
							NameType: Fixed,
							Size:     0,
						},
						3: {
							NameType: Fixed,
							Size:     0,
						},
						2: {
							NameType: Fixed,
							Size:     0,
						},
						//  1:FX
					},
				},
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:      22,
			DataItem: "I062/270",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   1,
				SecondarySize: 1,
			},
		},
		{
			FRN:      23,
			DataItem: "I062/300",
			Type: TypeField{
				NameType: Fixed,
				Size:     1},
		},
		{
			FRN:         24,
			DataItem:    "I062/110",
			Description: "Mode 5 Data reports & Extended Mode 1 Code",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "SUM",
							Description: "Mode 5 Summary",
						},
						7: {
							NameType:    Fixed,
							Size:        4,
							Item:        "PMN",
							Description: "Mode 5 PIN/ National Origin/ Mission Code",
						},
						6: {
							NameType:    Fixed,
							Size:        6,
							Item:        "POS",
							Description: "Mode 5 Reported Position",
						},
						5: {
							NameType:    Fixed,
							Size:        2,
							Item:        "GA",
							Description: "Mode 5 GNSS-derived Altitude",
						},
						4: {
							NameType:    Fixed,
							Size:        2,
							Item:        "EM1",
							Description: "Extended Mode 1 Code in Octal Representation",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TOS",
							Description: "Time Offset for POS and GA",
						},
						2: {
							NameType:    Fixed,
							Size:        1,
							Item:        "XP",
							Description: "X Pulse Presence",
						},
						//  1:FX
					},
				},
			},
		},
		{
			FRN:      25,
			DataItem: "I062/120",
			Type: TypeField{
				NameType: Fixed,
				Size:     2},
		},
		{
			FRN:      26,
			DataItem: "I062/510",
			Type: TypeField{
				NameType:      Extended,
				PrimarySize:   3,
				SecondarySize: 3,
			},
		},
		{
			FRN:         27,
			DataItem:    "I062/500",
			Description: "Estimated Accuracies",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        4,
							Item:        "APC",
							Description: "Accuracy Of Track Position (Cartesian)",
						},
						7: {
							NameType:    Fixed,
							Size:        2,
							Item:        "COV",
							Description: "XY Covariance",
						},
						6: {
							NameType:    Fixed,
							Size:        4,
							Item:        "APW",
							Description: "Estimated Accuracy Of Track Position (WGS-84)",
						},
						5: {
							NameType:    Fixed,
							Size:        1,
							Item:        "AGA",
							Description: "Estimated Accuracy Of Calculated Track Geometric",
						},
						4: {
							NameType:    Fixed,
							Size:        1,
							Item:        "ABA",
							Description: "Estimated Accuracy Of Calculated Track Barometric",
						},
						3: {
							NameType:    Fixed,
							Size:        2,
							Item:        "ATV",
							Description: "Estimated Accuracy Of Track Velocity (Cartesian)",
						},
						2: {
							NameType:    Fixed,
							Size:        2,
							Item:        "AA",
							Description: "Estimated Accuracy Of Acceleration (Cartesian)",
						},
						//  1:FX
					},
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        1,
							Item:        "ARC",
							Description: "Estimated Accuracy Of Rate Of Climb/Descent",
						},
						7: {
							NameType: Fixed,
							Size:     0,
						},
						6: {
							NameType: Fixed,
							Size:     0,
						},
						5: {
							NameType: Fixed,
							Size:     0,
						},
						4: {
							NameType: Fixed,
							Size:     0,
						},
						3: {
							NameType: Fixed,
							Size:     0,
						},
						2: {
							NameType: Fixed,
							Size:     0,
						},
						//  1:FX
					},
				},
			},
		},
		{
			FRN:         28,
			DataItem:    "I062/340",
			Description: "Measured Information",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {
							NameType:    Fixed,
							Size:        2,
							Item:        "SID",
							Description: "Sensor Identification",
						},
						7: {
							NameType:    Fixed,
							Size:        4,
							Item:        "POS",
							Description: "Measured Position",
						},
						6: {
							NameType:    Fixed,
							Size:        2,
							Item:        "HEI",
							Description: "Measured 3-D Height",
						},
						5: {
							NameType:    Fixed,
							Size:        2,
							Item:        "MDC",
							Description: "Last Measured Mode C code",
						},
						4: {
							NameType:    Fixed,
							Size:        2,
							Item:        "MDA",
							Description: "Last Measured Mode 3/A code",
						},
						3: {
							NameType:    Fixed,
							Size:        1,
							Item:        "TYP",
							Description: "Report Type",
						},
						2: {
							NameType: Fixed,
							Size:     0,
						},
						//  1:FX
					},
				},
			},
		},
		//FX : Field Extension Indicator
		{
			FRN:      29,
			DataItem: "NA",
			Type: TypeField{
				NameType: Spare},
		},
		{
			FRN:      30,
			DataItem: "NA",
			Type: TypeField{
				NameType: Spare},
		},
		{
			FRN:      31,
			DataItem: "NA",
			Type: TypeField{
				NameType: Spare},
		},
		{
			FRN:      32,
			DataItem: "NA",
			Type: TypeField{
				NameType: Spare},
		},
		{
			FRN:      33,
			DataItem: "NA",
			Type: TypeField{
				NameType: Spare},
		},
		{
			FRN:      34,
			DataItem: "RE-Data Item",
			Type: TypeField{
				NameType: RE,
			},
		},
		{
			FRN:      35,
			DataItem: "SP-Data Item",
			Type: TypeField{
				NameType: SP,
			},
		},
		//FX : Field Extension Indicator
	},
}
