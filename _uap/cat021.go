package _uap

// Cat021v10 User Application Profile
// version 2.5
/* todo
var Cat021v10 = StandardUAP{
	Name:     "cat021_2.5",
	Category: 21,
	Version:  2.5,
	DataItems: []DataField{
		{
			FRN:      1,
			DataItemName: "I021/010",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         2,
			DataItemName:    "I021/040",
			Description: "Target Report Descriptor",
			Type: TypeField{
				NameType:      Extended,
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
		},
		{
			FRN:         3,
			DataItemName:    "I021/161",
			Description: "Track Number",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         4,
			DataItemName:    "I021/015",
			Description: "Service Identification",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         5,
			DataItemName:    "I021/071",
			Description: "Time of Applicability for Position",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     3,
			},
		},
		{
			FRN:         6,
			DataItemName:    "I021/130",
			Description: "Position in WGS-84 co-ordinates",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     6,
			},
		},
		{
			FRN:         7,
			DataItemName:    "I021/131",
			Description: "Position in WGS-84 co-ordinates, high res",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     8,
			},
		},
		{
			FRN:         8,
			DataItemName:    "I021/072",
			Description: "Time of Applicability for Velocity",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     3,
			},
		},
		{
			FRN:         9,
			DataItemName:    "I021/150",
			Description: "Air Speed",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         10,
			DataItemName:    "I021/151",
			Description: "True Air Speed",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         11,
			DataItemName:    "I021/080",
			Description: "Target Address",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     3,
			},
		},
		{
			FRN:         12,
			DataItemName:    "I021/073",
			Description: "Time of Message Reception of Position",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     3,
			},
		},
		{
			FRN:         13,
			DataItemName:    "I021/074",
			Description: "Time of Message Reception of Position-High Precision",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     4,
			},
		},
		{
			FRN:         14,
			DataItemName:    "I021/075",
			Description: "Time of Message Reception of Velocity",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     3,
			},
		},
		{
			FRN:         15,
			DataItemName:    "I021/076",
			Description: "Time of Message Reception of Velocity-High Precision",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     4,
			},
		},
		{
			FRN:         16,
			DataItemName:    "I021/140",
			Description: "Geometric Height",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         17,
			DataItemName:    "I021/090",
			Description: "Quality Indicators",
			Type: TypeField{
				NameType:      Extended,
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
		},
		{
			FRN:         18,
			DataItemName:    "I021/210",
			Description: "MOPS Version",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         19,
			DataItemName:    "I021/070",
			Description: "Mode 3/A Code",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         20,
			DataItemName:    "I021/230",
			Description: "Roll Angle",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         21,
			DataItemName:    "I021/145",
			Description: "Flight Level",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         22,
			DataItemName:    "I021/152",
			Description: "Magnetic Heading",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         23,
			DataItemName:    "I021/200",
			Description: "Target Status",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         24,
			DataItemName:    "I021/155",
			Description: "Barometric Vertical Rate",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         25,
			DataItemName:    "I021/157",
			Description: "Geometric Vertical Rate",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         26,
			DataItemName:    "I021/160",
			Description: "Airborne Ground Vector",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     4,
			},
		},
		{
			FRN:         27,
			DataItemName:    "I021/165",
			Description: "Track Angle Rate",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         28,
			DataItemName:    "I021/177",
			Description: "Time of Report Transmission",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     3,
			},
		},
		{
			FRN:         29,
			DataItemName:    "I021/170",
			Description: "Target Identification",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     6,
			},
		},
		{
			FRN:         30,
			DataItemName:    "I021/020",
			Description: "Emitter Category",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         31,
			DataItemName:    "I021/220",
			Description: "Met Information",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, SizeField: 2},
						7: {NameType: Fixed, SizeField: 2},
						6: {NameType: Fixed, SizeField: 2},
						5: {NameType: Fixed, SizeField: 1},
						4: {NameType: Spare},
						3: {NameType: Spare},
						2: {NameType: Spare},
					},
				},
			},
		},
		{
			FRN:         32,
			DataItemName:    "I021/146",
			Description: "Selected Altitude",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         33,
			DataItemName:    "I021/148",
			Description: "Final State Selected Altitude",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     2,
			},
		},
		{
			FRN:         34,
			DataItemName:    "I021/110",
			Description: "Trajectory Intent",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, SizeField: 1},
						7: {NameType: Repetitive, SizeField: 15},
						6: {NameType: Spare},
						5: {NameType: Spare},
						4: {NameType: Spare},
						3: {NameType: Spare},
						2: {NameType: Spare},
					},
				},
			},
		},
		{
			FRN:         35,
			DataItemName:    "I021/016",
			Description: "Service Management",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         36,
			DataItemName:    "I021/008",
			Description: "Aircraft Operational Status",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         37,
			DataItemName:    "I021/271",
			Description: "Surface Capabilities and Characteristics",
			Type: TypeField{
				NameType:      Extended,
				PrimaryItemSize:   1,
				SecondaryItemSize: 1,
			},
		},
		{
			FRN:         38,
			DataItemName:    "I021/132",
			Description: "Message Amplitude",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         38,
			DataItemName:    "I021/250",
			Description: "Mode S MB Data",
			Type: TypeField{
				NameType: Repetitive,
				SizeField:     8,
			},
		},
		{
			FRN:         40,
			DataItemName:    "I021/260",
			Description: "ACAS Resolution Advisory Report",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     7,
			},
		},
		{
			FRN:         41,
			DataItemName:    "I021/400",
			Description: "Receiver ID",
			Type: TypeField{
				NameType: Fixed,
				SizeField:     1,
			},
		},
		{
			FRN:         42,
			DataItemName:    "I021/295",
			Description: "Data Ages",
			Type: TypeField{
				NameType: Compound,
				Primary: &Primary{
					MetaField{
						8: {NameType: Fixed, SizeField: 1},
						7: {NameType: Fixed, SizeField: 1},
						6: {NameType: Fixed, SizeField: 1},
						5: {NameType: Fixed, SizeField: 1},
						4: {NameType: Fixed, SizeField: 1},
						3: {NameType: Fixed, SizeField: 1},
						2: {NameType: Fixed, SizeField: 1},
					},
				},
			},
		},
		{
			FRN: 43, DataItemName: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 44, DataItemName: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 45, DataItemName: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 46, DataItemName: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN: 47, DataItemName: "NA", Type: TypeField{NameType: Spare},
		},
		{
			FRN:         48,
			DataItemName:    "RE-Data Item",
			Description: "Reserved Expansion Field",
			Type: TypeField{
				NameType: RE,
			},
		},
		{
			FRN:         49,
			DataItemName:    "SP-Data Item",
			Description: "Special Purpose Field",
			Type: TypeField{
				NameType: SP,
			},
		},
	},
}
*/
