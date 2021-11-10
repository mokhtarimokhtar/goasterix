package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func TestCat048Model_ToJsonRecord(t *testing.T) {
	// Arrange
	// bds 02 e79a5d27a00c00 60 a3280030a40000 40
	//input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	input := "ffff02 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 00800080 0743ce5b 40 20f5"
	output := []byte(`{"sourceIdentifier":{"sac":8,"sic":54},"aircraftAddress":"490D01","aircraftIdentification":"NJE834H","timeOfDay":34102.640625,"rhoTheta":{"rho":148.77734375,"theta":2.1174999999999997},"cartesianXY":{"x":1,"y":1},"flightLevel":{"v":"code_validated","g":"default","level":180},"radarPlotCharacteristics":{"srr":2,"sam":-73},"mode3ACode":{"squawk":"4423","v":"code_validated","g":"default","l":"code_derived_from_transponder"},"trackNumber":1594,"trackVelocity":{"groundSpeed":0.113464065,"heading":290.5485},"trackStatus":{"cnf":"confirmed_track","rad":"ssr_modes_track","dou":"normal_confidence","mah":"no_horizontal_man_sensed","cdm":"maintaining"},"bdsRegisterData":[{"transponderRegisterNumber":"60","code60":{"magneticHeading":-68,"indicatedAirspeed":302,"mach":0.632,"barometricAltitudeRate":32}},{"transponderRegisterNumber":"40","code40":{"mcpSelectAltitude":18000,"barometricPressureSetting":1013}}],"comAcasCapabilityFlightStatus":{"com":"comm_a_and_comm_b_capability","stat":"no_alert_no_spi_aircraft_airborne","si":"si_code_capable","mssc":"yes","arc":"25_ft_resolution","aic":"yes","b1a":"1","b1b":"5"}}`)

	uap048 := uap.Cat048V127
	data, _ := goasterix.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, err := rec.Decode(data, uap048)

	cat048Model := new(Cat048Model)
	cat048Model.write(rec.Items)

	// Act
	recJson, _ := json.Marshal(cat048Model)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf("FAIL: %s; \nExpected: %s", recJson, output)
	} else {
		t.Logf("SUCCESS: %s; Expected: %s", recJson, output)
	}
}

func TestCat048Model_RhoTheta(t *testing.T) {
	// Arrange
	input := [4]byte{0xFF, 0xFF, 0xFF, 0xFF}
	output := PolarPosition{Rho: float64(0xFFFF) / 256, Theta: float64(0xFFFF) * 0.0055}

	// Act
	res, err := rhoTheta(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

type DataTestMode3A struct {
	TestCaseName string
	input        [2]byte
	output       Mode3A
}

func TestCat048Model_Mode3ACodeVGL(t *testing.T) {
	// setup
	dataset := []DataTestMode3A{
		{
			TestCaseName: "validated",
			input:        [2]byte{0x1F, 0xFF},
			output: Mode3A{
				Squawk: "7777", // FFF = 111 111 111 111 = 7777
				V:      "code_validated",
				G:      "default",
				L:      "code_derived_from_transponder",
			},
		},
		{
			TestCaseName: "not validated",
			input:        [2]byte{0xEF, 0xFF},
			output: Mode3A{
				Squawk: "7777", // FFF = 111 111 111 111 = 7777
				V:      "code_not_validated",
				G:      "garbled_code",
				L:      "code_not_extracted",
			},
		},
	}
	for _, row := range dataset {
		// Arrange
		// Act
		res := mode3ACodeVGL(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

type DataTestFlightLevel struct {
	TestCaseName string
	input        [2]byte
	output       FL
}

func TestCat048Model_FlightLevel(t *testing.T) {
	// setup
	dataset := []DataTestFlightLevel{
		{
			TestCaseName: "validated",
			input:        [2]byte{0x00, 0xFF},
			output: FL{
				V:     "code_validated",
				G:     "default",
				Level: float64(uint16(0x00FF)) / 4,
			},
		},
		{
			TestCaseName: "not validated",
			input:        [2]byte{0xC0, 0xFF}, // 1100-1111 1111-1111
			output: FL{
				V:     "code_not_validated",
				G:     "garbled_code",
				Level: float64(uint16(0x00FF)) / 4,
			},
		},
	}
	for _, row := range dataset {
		// Arrange
		// Act
		res := flightLevel(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

type DataTestModeSI struct {
	TestCaseName string
	input        [6]byte
	output       string
	err          error
}

func TestCat048Model_ModeSIdentification(t *testing.T) {
	// setup
	dataset := []DataTestModeSI{
		{
			TestCaseName: "valid chars",
			input:        [6]byte{0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x20},
			output:       "AFR1234",
			err:          nil,
		},
		{
			TestCaseName: "chars 6 unknown",
			input:        [6]byte{0x04, 0x64, 0xB1, 0xCB, 0x3D, 0x3A},
			output:       "AFR1234",
			err:          ErrCharUnknown,
		},
		{
			TestCaseName: "chars unknown",
			input:        [6]byte{},
			output:       "",
			err:          ErrCharUnknown,
		},
	}

	for _, row := range dataset {
		// Arrange
		// Act
		s, err := modeSIdentification(row.input)

		// Assert
		if err != row.err {
			//if errors.Is(err, row.err) {
			t.Errorf("FAIL: error: %v; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}

		if s != row.output {
			t.Errorf("FAIL: s = %s; Expected: %s", s, row.output)
		} else {
			t.Logf("SUCCESS: s = %s; Expected: %s", s, row.output)
		}
	}
}

func TestCat048Model_TrackNumber(t *testing.T) {
	// Arrange
	input := [2]byte{0x0F, 0xFF}
	output := uint16(4095)

	// Act
	res, err := trackNumber(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func TestCat048Model_CartesianXY(t *testing.T) {
	// Arrange
	input := [4]byte{0x01, 0x00, 0xFF, 0x00}
	output := CartesianXYPosition{
		X: float64(int16(256)) / 128,
		Y: float64(int16(-256)) / 128,
	}

	// Act
	res, err := cartesianXY(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func TestCat048Model_TrackVelocity(t *testing.T) {
	// Arrange
	input := [4]byte{0x07, 0xc3, 0xdf, 0xc6}
	output := Velocity{
		GroundSpeed: 0.121276545,
		Heading:     315.073,
	}

	// Act
	res, err := trackVelocity(input)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(res, output) == false {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

type DataTestTrackStatus struct {
	TestCaseName string
	input        []byte
	output       Status
}

func TestCat048Model_TrackStatus(t *testing.T) {
	// setup
	dataset := []DataTestTrackStatus{
		{
			TestCaseName: "testcase 1",
			input:        []byte{0x41, 0x00},
			output: Status{
				CNF: "confirmed_track",
				RAD: "ssr_modes_track",
				DOU: "normal_confidence",
				MAH: "no_horizontal_man_sensed",
				CDM: "maintaining",
				TRE: "track_still_alive",
				GHO: "true_target_track",
				SUP: "no",
				TCC: "radar_plane",
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        []byte{0x9b, 0xF0},
			output: Status{
				CNF: "tentative_track",
				RAD: "combined_track",
				DOU: "low_confidence",
				MAH: "horizontal_man_sensed",
				CDM: "climbing",
				TRE: "end_of_track_lifetime",
				GHO: "ghost_target_track",
				SUP: "yes",
				TCC: "slant_range_correction_used",
			},
		},
		{
			TestCaseName: "testcase 3",
			input:        []byte{0x24}, // 0010-0100
			output: Status{
				CNF: "confirmed_track",
				RAD: "psr_track",
				DOU: "normal_confidence",
				MAH: "no_horizontal_man_sensed",
				CDM: "descending",
			},
		},
		{
			TestCaseName: "testcase 4",
			input:        []byte{0x66}, // 0110-0110
			output: Status{
				CNF: "confirmed_track",
				RAD: "invalid",
				DOU: "normal_confidence",
				MAH: "no_horizontal_man_sensed",
				CDM: "unknown",
			},
		},
	}
	for _, row := range dataset {
		// Arrange
		// Act
		res := trackStatus(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

type DataTestRadarPlot struct {
	TestCaseName string
	input        []byte
	output       PlotCharacteristics
}

func TestCat048Model_RadarPlotCharacteristics(t *testing.T) {
	// setup
	dataset := []DataTestRadarPlot{
		{
			TestCaseName: "full subfield",
			input:        []byte{0xFE, 0x64, 0x64, 0x64, 0x64, 0x64, 0x7F, 0x64},
			output: PlotCharacteristics{
				SRL: 4.3999999999999995,
				SRR: 100,
				SAM: 100,
				PRL: 4.3999999999999995,
				PAM: 100,
				RPD: 0.49609375,
				APD: 2.1972656,
			},
		},
	}
	for _, row := range dataset {
		// Arrange
		// Act
		res := radarPlotCharacteristics(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

type DataTestComACAS struct {
	TestCaseName string
	input        [2]byte
	output       ACASCapaFlightStatus
}

func TestCat048Model_ComACASCapabilityFlightStatus(t *testing.T) {
	// setup
	dataset := []DataTestComACAS{
		{
			TestCaseName: "testcase 1",
			input:        [2]byte{0x00, 0x00},
			output: ACASCapaFlightStatus{
				COM:  "no_communications_capability",
				STAT: "no_alert_no_spi_aircraft_airborne",
				SI:   "si_code_capable",
				MSSC: "no",
				ARC:  "100_ft_resolution",
				AIC:  "no",
				B1A:  "0",
				B1B:  "0",
			},
		},
		{
			TestCaseName: "testcase 2",
			input:        [2]byte{0x26, 0xFF}, //0010 0110 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "comm_a_and_comm_b_capability",
				STAT: "no_alert_no_spi_aircraft_on_ground",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
		{
			TestCaseName: "testcase 3",
			input:        [2]byte{0x4a, 0xFF}, //0100 1010 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "comm_a_and_comm_b_and_uplink_elm",
				STAT: "alert_no_spi_aircraft_airborne",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
		{
			TestCaseName: "testcase 4",
			input:        [2]byte{0x6E, 0xFF}, //0110 1110 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "comm_a_and_comm_b_and_uplink_elm_and_downlink_elm",
				STAT: "alert_no_spi_aircraft_on_ground",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
		{
			TestCaseName: "testcase 5",
			input:        [2]byte{0x92, 0xFF}, //1001 0010 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "level_5_transponder_capability",
				STAT: "alert_spi_aircraft_airborne_or_on_ground",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
		{
			TestCaseName: "testcase 6",
			input:        [2]byte{0xb6, 0xFF}, //1011 0110 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "not_assigned",
				STAT: "no_alert_spi_aircraft_airborne_or_on_ground",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
		{
			TestCaseName: "testcase 7",
			input:        [2]byte{0xda, 0xFF}, //1101 1010 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "not_assigned",
				STAT: "not_assigned",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
		{
			TestCaseName: "testcase 8",
			input:        [2]byte{0xfe, 0xFF}, //1111 1110 1111 1111
			output: ACASCapaFlightStatus{
				COM:  "not_assigned",
				STAT: "unknown",
				SI:   "sii_code_capable",
				MSSC: "yes",
				ARC:  "25_ft_resolution",
				AIC:  "yes",
				B1A:  "1",
				B1B:  "15",
			},
		},
	}
	for _, row := range dataset {
		// Arrange
		// Act
		res := comACASCapabilityFlightStatus(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}
