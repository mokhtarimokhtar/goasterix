package uap

// CatForTest User Application Profile
// Specific for testing
var CatForTest = StandardUAP{
	Name:     "catfortest_0.1",
	Category: 26, // not exist
	Version:  0.1,
	DataItems: []DataField{
		{
			FRN:         1,
			DataItem:    "I026/010",
			Description: "Data Source Identifier",
			Type:        Fixed,
			SizeItem:    SizeField{ForFixed: 2},
			SubItems: []SubItem{
				{
					Name: "SAC",
					Type: FromTo,
					BitPosition: BitPosition{
						From: 16,
						To:   9,
					},
				},
				{
					Name: "SIC",
					Type: FromTo,
					BitPosition: BitPosition{
						From: 8,
						To:   1,
					},
				},
			},
		},
		{
			FRN:         2,
			DataItem:    "I026/010",
			Description: "Code",
			Type:        Fixed,
			SizeItem:    SizeField{ForFixed: 2},
			SubItems: []SubItem{
				{
					Name: "V",
					Type: Bit,
					BitPosition: BitPosition{
						Bit: 8,
					},
				},
				{
					Name: "G",
					Type: Bit,
					BitPosition: BitPosition{
						Bit: 7,
					},
				},
				{
					Name: "G",
					Type: Bit,
					BitPosition: BitPosition{
						Bit: 6,
					},
				},
				{
					Name: "Mode Code",
					Type: FromTo,
					BitPosition: BitPosition{
						From: 5,
						To:   1,
					},
				},
			},
		},

		/*
			{
				FRN:         1,
				DataItem:    "I026/001",
				Description: "Fixed type field for test",
				Type:        Fixed,
				SizeItem: SizeField{
					ForFixed: 2,
				},
			},
			{
				FRN:         2,
				DataItem:    "I026/002",
				Description: "Extended type field for test",
				Type:        Extended,
				SizeItem: SizeField{
					ForExtendedPrimary:   1,
					ForExtendedSecondary: 2,
				},
			},
			{
				FRN:         3,
				DataItem:    "I026/003",
				Description: "Explicit type field for test",
				Type:        Explicit,
			},
			{
				FRN:         4,
				DataItem:    "I026/004",
				Description: "Repetitive type field for test",
				Type:        Repetitive,
				SizeItem: SizeField{
					ForRepetitive: 2,
				},
			},
			{
				FRN:         5,
				DataItem:    "I026/005",
				Description: "Compound type field for test",
				Type:        Compound,
				Compound: []DataField{
					{
						FRN:         1,
						DataItem:    "Compound/001",
						Description: "Compound Fixed type field for test",
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
						FRN:         3,
						DataItem:    "Compound/003",
						Description: "Compound Extended type field for test",
						Type:        Extended,
						SizeItem: SizeField{
							ForExtendedPrimary:   1,
							ForExtendedSecondary: 1,
						},
					},
					{
						FRN:  4,
						Type: Spare,
					},
					{
						FRN:         5,
						DataItem:    "Compound/005",
						Description: "Compound Repetitive type field for test",
						Type:        Repetitive,
						SizeItem: SizeField{
							ForRepetitive: 2,
						},
					},
					{
						FRN:  6,
						Type: Spare,
					},
					{
						FRN:         7,
						DataItem:    "Compound/007",
						Description: "Compound Explicit type field for test",
						Type:        Explicit,
					},
					{
						FRN:         8,
						DataItem:    "Compound/008",
						Description: "Compound Fixed type field for test",
						Type:        Fixed,
						SizeItem: SizeField{
							ForFixed: 2,
						},
					},
				},
			},
			{
				FRN:         6,
				DataItem:    "I026/006",
				Description: "RFS(Random Field Sequencing) type field for test",
				Type:        RFS,
				RFS: []DataField{},
			},
			{
				FRN:      7,
				DataItem: "NA",
				Type:     Spare,
			},
			{
				FRN:         8,
				DataItem:    "RE",
				Description: "RE (Reserved Expansion Field) type field for test",
				Type:        RE,
			},
			{
				FRN:         9,
				DataItem:    "SP",
				Description: "SP (Special Purpose field) type field for test",
				Type:        SP,
			},
			{
				FRN:         10,
				DataItem:    "I026/010",
				Description: "Fixed type field for test",
				Conditional: true,
				Type:        Fixed,
				SizeItem: SizeField{
					ForFixed: 1,
				},
			},
		*/

	},
}

var Cat4TestTrack = []DataField{
	{
		FRN:         11,
		DataItem:    "I026/011",
		Description: "Fixed type field for test",
		Type:        Fixed,
		SizeItem: SizeField{
			ForFixed: 1,
		},
	},
	{
		FRN:         12,
		DataItem:    "I026/012",
		Description: "Fixed type field for test",
		Type:        Fixed,
		SizeItem: SizeField{
			ForFixed: 2,
		},
	},
}

var Cat4TestPlot = []DataField{
	{
		FRN:         11,
		DataItem:    "I026/011",
		Description: "Fixed type field for test",
		Type:        Fixed,
		SizeItem: SizeField{
			ForFixed: 3,
		},
	},
	{
		FRN:         12,
		DataItem:    "I026/012",
		Description: "Fixed type field for test",
		Type:        Fixed,
		SizeItem: SizeField{
			ForFixed: 1,
		},
	},
}
