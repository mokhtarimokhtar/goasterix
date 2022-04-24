package uap

// Cat030ArtasV70 User Application Profile
// version 7.0
var Cat030ArtasV70 = StandardUAP{
	Name:     "ARTAS",
	Category: 30,
	Version:  7.0,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I030/010",
			Description: "SERVER IDENTIFICATION TAG",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I030/015",
			Description: "USER NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         3,
			DataItem:    "I030/030",
			Description: "SERVICE IDENTIFICATION",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         4,
			DataItem:    "I030/035",
			Description: "TYPE OF MESSAGE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         5,
			DataItem:    "I030/040",
			Description: "TRACK NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         6,
			DataItem:    "I030/070",
			Description: "TIME OF LAST UPDATE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         7,
			DataItem:    "I030/170",
			Description: "TRACK AGES",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         8,
			DataItem:    "I030/100",
			Description: "CALCULATED TRACK POSITION (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         9,
			DataItem:    "I030/180",
			Description: "CALCULATED TRACK VELOCITY (POLAR)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         10,
			DataItem:    "I030/181",
			Description: "CALCULATED TRACK VELOCITY (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         11,
			DataItem:    "I030/060",
			Description: "TRACK MODE 3/A",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         12,
			DataItem:    "I030/150",
			Description: "MEASURED TRACK MODE C",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         13,
			DataItem:    "I030/130",
			Description: "CALCULATED TRACK ALTITUDE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         14,
			DataItem:    "I030/160",
			Description: "CALCULATED TRACK FLIGHT LEVEL",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         15,
			DataItem:    "I030/080",
			Description: "ARTAS TRACK STATUS",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         16,
			DataItem:    "I030/090",
			Description: "ARTAS TRACK QUALITY",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         17,
			DataItem:    "I030/200",
			Description: "MODE OF FLIGHT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         18,
			DataItem:    "I030/220",
			Description: "CALCULATED RATE OF CLIMB/DESCENT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         19,
			DataItem:    "I030/240",
			Description: "CALCULATED RATE OF TURN",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         20,
			DataItem:    "I030/290",
			Description: "PLOT AGES",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         21,
			DataItem:    "I030/260",
			Description: "RADAR IDENTIFICATION TAG",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         22,
			DataItem:    "I030/360",
			Description: "MEASURED POSITION",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         23,
			DataItem:    "I030/140",
			Description: "LAST MEASURED MODE C",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         24,
			DataItem:    "I030/340",
			Description: "LAST MEASURED MODE 3/A",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         25,
			DataItem:    "I030/RE",
			Description: "RESERVED EXPANSION DATA FIELD",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         26,
			DataItem:    "I030/390",
			Description: "FPPS IDENTIFICATION TAG",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         27,
			DataItem:    "I030/400",
			Description: "CALLSIGN",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 7,
			},
		},
		{
			FRN:         28,
			DataItem:    "I030/410",
			Description: "PLN NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         29,
			DataItem:    "I030/440",
			Description: "DEPARTURE AIRPORT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         30,
			DataItem:    "I030/450",
			Description: "DESTINATION AIRPORT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         31,
			DataItem:    "I030/435",
			Description: "CATEGORY OF TURBULENCE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         32,
			DataItem:    "I030/430",
			Description: "TYPE OF AIRCRAFT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         33,
			DataItem:    "I030/460",
			Description: "ALLOCATED SSR CODES",
			Type:        Repetitive,
			SizeItem: SizeField{
				ForRepetitive: 2,
			},
		},
		{
			FRN:         34,
			DataItem:    "I030/480",
			Description: "CURRENT CLEARED FLIGHT LEVEL",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         35,
			DataItem:    "I030/420",
			Description: "FLIGHT CATEGORY",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         36,
			DataItem:    "I030/490",
			Description: "CURRENT CONTROL POSITION",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         37,
			DataItem:    "I030/020",
			Description: "TIME OF MESSAGE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         38,
			DataItem:    "I030/382",
			Description: "AIRCRAFT ADDRESS",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         39,
			DataItem:    "I030/384",
			Description: "AIRCRAFT IDENTIFICATION",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 6,
			},
		},
		{
			FRN:         40,
			DataItem:    "I030/386",
			Description: "COMMUNICATIONS CAPABILITY AND FLIGHT STATUS",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         41,
			DataItem:    "I030/110",
			Description: "ESTIMATED ACCURACY OF TRACK POSITION (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         42,
			DataItem:    "I030/190",
			Description: "ESTIMATED ACCURACY OF TRACK VELOCITY (POLAR)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         43,
			DataItem:    "I030/191",
			Description: "ESTIMATED ACCURACY OF TRACK VELOCITY (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         44,
			DataItem:    "I030/135",
			Description: "ESTIMATED ACCURACY OF TRACK ALTITUDE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         45,
			DataItem:    "I030/165",
			Description: "ESTIMATED ACCURACY OF CALCULATED TRACK FLIGHT LEVEL",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         46,
			DataItem:    "I030/230",
			Description: "ESTIMATED ACCURACY OF RATE OF CLIMB/DESCENT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         47,
			DataItem:    "I030/250",
			Description: "ESTIMATED ACCURACY OF RATE OF TURN",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         48,
			DataItem:    "I030/210",
			Description: "MODE OF FLIGHT PROBABILITIES",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         49,
			DataItem:    "I030/120",
			Description: "TRACK MODE 2 CODE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         50,
			DataItem:    "I030/050",
			Description: "ARTAS TRACK NUMBER",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   3,
				ForExtendedSecondary: 3,
			},
		},
		{
			FRN:         51,
			DataItem:    "I030/270",
			Description: "LOCAL TRACK NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         52,
			DataItem:    "I030/370",
			Description: "MEASURED 3-D HEIGHT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN: 53, DataItem: "NA", Type: Spare,
		},
		{
			FRN: 54, DataItem: "NA", Type: Spare,
		},
		{
			FRN: 55, DataItem: "NA", Type: Spare,
		},
		{
			FRN:      56,
			DataItem: "I030/RE",
			Type:     RE,
		},
	},
}

// Cat030ArtasV62 User Application Profile
// version 6.2
var Cat030ArtasV62 = StandardUAP{
	Name:     "ARTAS",
	Category: 30,
	Version:  6.2,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I030/010",
			Description: "SERVER IDENTIFICATION TAG",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I030/015",
			Description: "USER NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         3,
			DataItem:    "I030/030",
			Description: "SERVICE IDENTIFICATION",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         4,
			DataItem:    "I030/035",
			Description: "TYPE OF MESSAGE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         5,
			DataItem:    "I030/040",
			Description: "TRACK NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         6,
			DataItem:    "I030/070",
			Description: "TIME OF LAST UPDATE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         7,
			DataItem:    "I030/170",
			Description: "TRACK AGES",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         8,
			DataItem:    "I030/100",
			Description: "CALCULATED TRACK POSITION (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         9,
			DataItem:    "I030/180",
			Description: "CALCULATED TRACK VELOCITY (POLAR)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         10,
			DataItem:    "I030/181",
			Description: "CALCULATED TRACK VELOCITY (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         11,
			DataItem:    "I030/060",
			Description: "TRACK MODE 3/A",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         12,
			DataItem:    "I030/150",
			Description: "MEASURED TRACK MODE C",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         13,
			DataItem:    "I030/130",
			Description: "CALCULATED TRACK ALTITUDE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         14,
			DataItem:    "I030/160",
			Description: "CALCULATED TRACK FLIGHT LEVEL",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         15,
			DataItem:    "I030/080",
			Description: "ARTAS TRACK STATUS",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   1,
				ForExtendedSecondary: 1,
			},
		},
		{
			FRN:         16,
			DataItem:    "I030/090",
			Description: "ARTAS TRACK QUALITY",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         17,
			DataItem:    "I030/200",
			Description: "MODE OF FLIGHT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         18,
			DataItem:    "I030/220",
			Description: "CALCULATED RATE OF CLIMB/DESCENT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         19,
			DataItem:    "I030/240",
			Description: "CALCULATED RATE OF TURN",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         20,
			DataItem:    "I030/290",
			Description: "PLOT AGES",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         21,
			DataItem:    "I030/260",
			Description: "RADAR IDENTIFICATION TAG",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         22,
			DataItem:    "I030/360",
			Description: "MEASURED POSITION",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         23,
			DataItem:    "I030/140",
			Description: "LAST MEASURED MODE C",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         24,
			DataItem:    "I030/340",
			Description: "LAST MEASURED MODE 3/A",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         25,
			DataItem:    "I030/RE",
			Description: "RESERVED EXPANSION DATA FIELD",
			Type:        RE,
		},
		{
			FRN:         26,
			DataItem:    "I030/390",
			Description: "FPPS IDENTIFICATION TAG",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         27,
			DataItem:    "I030/400",
			Description: "CALLSIGN",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 7,
			},
		},
		{
			FRN:         28,
			DataItem:    "I030/410",
			Description: "PLN NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         29,
			DataItem:    "I030/440",
			Description: "DEPARTURE AIRPORT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         30,
			DataItem:    "I030/450",
			Description: "DESTINATION AIRPORT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         31,
			DataItem:    "I030/435",
			Description: "CATEGORY OF TURBULENCE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         32,
			DataItem:    "I030/430",
			Description: "TYPE OF AIRCRAFT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         33,
			DataItem:    "I030/460",
			Description: "ALLOCATED SSR CODES",
			Type:        Repetitive,
			SizeItem: SizeField{
				ForRepetitive: 2,
			},
		},
		{
			FRN:         34,
			DataItem:    "I030/480",
			Description: "CURRENT CLEARED FLIGHT LEVEL",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         35,
			DataItem:    "I030/420",
			Description: "FLIGHT CATEGORY",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         36,
			DataItem:    "I030/490",
			Description: "CURRENT CONTROL POSITION",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         37,
			DataItem:    "I030/020",
			Description: "TIME OF MESSAGE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         38,
			DataItem:    "I030/382",
			Description: "AIRCRAFT ADDRESS",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         39,
			DataItem:    "I030/384",
			Description: "AIRCRAFT IDENTIFICATION",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 6,
			},
		},
		{
			FRN:         40,
			DataItem:    "I030/386",
			Description: "COMMUNICATIONS CAPABILITY AND FLIGHT STATUS",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         41,
			DataItem:    "I030/110",
			Description: "ESTIMATED ACCURACY OF TRACK POSITION (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         42,
			DataItem:    "I030/190",
			Description: "ESTIMATED ACCURACY OF TRACK VELOCITY (POLAR)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         43,
			DataItem:    "I030/191",
			Description: "ESTIMATED ACCURACY OF TRACK VELOCITY (CARTESIAN)",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 4,
			},
		},
		{
			FRN:         44,
			DataItem:    "I030/135",
			Description: "ESTIMATED ACCURACY OF TRACK ALTITUDE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         45,
			DataItem:    "I030/165",
			Description: "ESTIMATED ACCURACY OF CALCULATED TRACK FLIGHT LEVEL",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         46,
			DataItem:    "I030/230",
			Description: "ESTIMATED ACCURACY OF RATE OF CLIMB/DESCENT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         47,
			DataItem:    "I030/250",
			Description: "ESTIMATED ACCURACY OF RATE OF TURN",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 1,
			},
		},
		{
			FRN:         48,
			DataItem:    "I030/210",
			Description: "MODE OF FLIGHT PROBABILITIES",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 3,
			},
		},
		{
			FRN:         49,
			DataItem:    "I030/120",
			Description: "TRACK MODE 2 CODE",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         50,
			DataItem:    "I030/050",
			Description: "ARTAS TRACK NUMBER",
			Type:        Extended,
			SizeItem: SizeField{
				ForExtendedPrimary:   3,
				ForExtendedSecondary: 3,
			},
		},
		{
			FRN:         51,
			DataItem:    "I030/270",
			Description: "LOCAL TRACK NUMBER",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN:         52,
			DataItem:    "I030/370",
			Description: "MEASURED 3-D HEIGHT",
			Type:        Fixed,
			SizeItem: SizeField{
				ForFixed: 2,
			},
		},
		{
			FRN: 53, DataItem: "NA", Type: Spare,
		},
		{
			FRN: 54, DataItem: "NA", Type: Spare,
		},
		{
			FRN: 55, DataItem: "NA", Type: Spare,
		},
		{
			FRN: 56, DataItem: "NA", Type: Spare,
		},
	},
}
