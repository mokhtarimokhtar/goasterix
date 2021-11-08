package uap

var Cat001V12 = StandardUAP{
	Name:     "cat001_1.2",
	Category: 1,
	Version:  1.2,
	Items: []DataField{
		{
			FRN: 1,
			DataItem: "I001/010",
			Type: TypeField{
				Name: Fixed,
				Size: 2,
			},
		},
		{
			FRN:         2,
			DataItem:    "I001/020",
			Conditional: true,
			Type: TypeField{
				Name: Extended,
				Size: 1,
			},
		},
	},
}

var Cat001TrackV12 = []DataField{
	{
		FRN: 3, DataItem: "I001/161", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 4, DataItem: "I001/040", Type: TypeField{Name: Fixed, Size: 4},
	},
	{
		FRN: 5, DataItem: "I001/042", Type: TypeField{Name: Fixed, Size: 4},
	},
	{
		FRN: 6, DataItem: "I001/200", Type: TypeField{Name: Fixed, Size: 4},
	},
	{
		FRN: 7, DataItem: "I001/070", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 8, DataItem: "I001/090", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 9, DataItem: "I001/141", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 10, DataItem: "I001/130", Type: TypeField{Name: Extended, Size: 1},
	},
	{
		FRN: 11, DataItem: "I001/131", Type: TypeField{Name: Fixed, Size: 1},
	},
	{
		FRN: 12, DataItem: "I001/120", Type: TypeField{Name: Fixed, Size: 1},
	},
	{
		FRN: 13, DataItem: "I001/170", Type: TypeField{Name: Extended, Size: 1},
	},
	{
		FRN: 14, DataItem: "I001/210", Type: TypeField{Name: Extended, Size: 1},
	},
	{
		FRN: 15, DataItem: "I001/050", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 16, DataItem: "I001/080", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 17, DataItem: "I001/100", Type: TypeField{Name: Fixed, Size: 4},
	},
	{
		FRN: 18, DataItem: "I001/060", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 19, DataItem: "I001/030", Type: TypeField{Name: Extended, Size: 1},
	},
	{
		FRN: 20, DataItem: "SP-Data Item", Type: TypeField{Name: SP},
	},
	{
		FRN: 21, DataItem: "Random Field Sequencing", Type: TypeField{Name: RFS},
	},
	{
		FRN: 22, DataItem: "I001/150", Type: TypeField{Name: Fixed, Size: 1},
	},
}

var Cat001PlotV12 = []DataField{
	{
		FRN: 3, DataItem: "I001/040", Type: TypeField{Name: Fixed, Size: 4},
	},
	{
		FRN: 4, DataItem: "I001/070", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 5, DataItem: "I001/090", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 6, DataItem: "I001/130", Type: TypeField{Name: Extended, Size: 1},
	},
	{
		FRN: 7, DataItem: "I001/141", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 8, DataItem: "I001/050", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 9, DataItem: "I001/120", Type: TypeField{Name: Fixed, Size: 1},
	},
	{
		FRN: 10, DataItem: "I001/131", Type: TypeField{Name: Fixed, Size: 1},
	},
	{
		FRN: 11, DataItem: "I001/080", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 12, DataItem: "I001/100", Type: TypeField{Name: Fixed, Size: 4},
	},
	{
		FRN: 13, DataItem: "I001/060", Type: TypeField{Name: Fixed, Size: 2},
	},
	{
		FRN: 14, DataItem: "I001/030", Type: TypeField{Name: Extended, Size: 1},
	},
	{
		FRN: 15, DataItem: "I001/150", Type: TypeField{Name: Fixed, Size: 1},
	},
	{
		FRN: 16, DataItem: "NA", Type: TypeField{Name: Spare},
	},
	{
		FRN: 16, DataItem: "NA", Type: TypeField{Name: Spare},
	},
	{
		FRN: 17, DataItem: "NA", Type: TypeField{Name: Spare},
	},
	{
		FRN: 18, DataItem: "NA", Type: TypeField{Name: Spare},
	},
	{
		FRN: 19, DataItem: "NA", Type: TypeField{Name: Spare},
	},
	{
		FRN: 20, DataItem: "SP-Data Item", Type: TypeField{Name: SP},
	},
	{
		FRN: 21, DataItem: "Random Field Sequencing", Type: TypeField{Name: RFS},
	},
}

