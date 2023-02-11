package transform

type LengthWidth struct {
	Length string
	Width  string
}

var TableLWV1 = map[uint32]LengthWidth{
	uint32(0):  LengthWidth{Length: "L < 15", Width: "W < 11.5"},
	uint32(1):  LengthWidth{Length: "L < 15", Width: "W < 23"},
	uint32(2):  LengthWidth{Length: "L < 25", Width: "W < 28.5"},
	uint32(3):  LengthWidth{Length: "L < 25", Width: "W < 34"},
	uint32(4):  LengthWidth{Length: "L < 35", Width: "W < 33"},
	uint32(5):  LengthWidth{Length: "L < 35", Width: "W < 38"},
	uint32(6):  LengthWidth{Length: "L < 45", Width: "W < 39.5"},
	uint32(7):  LengthWidth{Length: "L < 45", Width: "W < 45"},
	uint32(8):  LengthWidth{Length: "L < 55", Width: "W < 45"},
	uint32(9):  LengthWidth{Length: "L < 55", Width: "W < 52"},
	uint32(10): LengthWidth{Length: "L < 65", Width: "W < 59.5"},
	uint32(11): LengthWidth{Length: "L < 65", Width: "W < 67"},
	uint32(12): LengthWidth{Length: "L < 75", Width: "W < 72.5"},
	uint32(13): LengthWidth{Length: "L < 75", Width: "W < 80"},
	uint32(14): LengthWidth{Length: "L < 85", Width: "W < 80"},
	uint32(15): LengthWidth{Length: "L < 85", Width: "W > 80"},
}

var TableLWV2 = map[uint32]LengthWidth{
	uint32(0):  LengthWidth{Length: "L < 15", Width: "W < 11.5"},
	uint32(1):  LengthWidth{Length: "L < 15", Width: "W < 23"},
	uint32(2):  LengthWidth{Length: "L < 25", Width: "W < 28.5"},
	uint32(3):  LengthWidth{Length: "L < 25", Width: "W < 34"},
	uint32(4):  LengthWidth{Length: "L < 35", Width: "W < 33"},
	uint32(5):  LengthWidth{Length: "L < 35", Width: "W < 38"},
	uint32(6):  LengthWidth{Length: "L < 45", Width: "W < 39.5"},
	uint32(7):  LengthWidth{Length: "L < 45", Width: "W < 45"},
	uint32(8):  LengthWidth{Length: "L < 55", Width: "W < 45"},
	uint32(9):  LengthWidth{Length: "L < 55", Width: "W < 52"},
	uint32(10): LengthWidth{Length: "L < 65", Width: "W < 59.5"},
	uint32(11): LengthWidth{Length: "L < 65", Width: "W < 67"},
	uint32(12): LengthWidth{Length: "L < 75", Width: "W < 72.5"},
	uint32(13): LengthWidth{Length: "L < 75", Width: "W < 80"},
	uint32(14): LengthWidth{Length: "L < 85", Width: "W < 80"},
	uint32(15): LengthWidth{Length: "L > 85", Width: "W > 80"},
}
