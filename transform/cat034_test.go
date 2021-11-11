package transform

import (
	"reflect"
	"testing"
)

func TestCat034Model_MessageType(t *testing.T) {
	// Arrange
	type dataTest struct {
		input  [1]byte
		output string
	}
	dataset := []dataTest{
		{[1]byte{0x00}, "undefined_message_type"},
		{[1]byte{0x01}, "north_marker_message"},
		{[1]byte{0x02}, "sector_crossing_message"},
		{[1]byte{0x03}, "geographical_filtering_message"},
		{[1]byte{0x04}, "jamming_strobe_message"},
		{[1]byte{0x05}, "solar_storm_message"},
		{[1]byte{0x06}, "ssr_jamming_strobe_message"},
		{[1]byte{0x07}, "mode_s_jamming_strobe_message"},
	}
	for _, row := range dataset {
		// Act
		res := messageType(row.input)

		// Assert
		if res != row.output {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat034Model_SystemConfiguration(t *testing.T) {
	// Arrange
	input := []byte{0x84, 0x00, 0x20, 0x00}
	output := SysConf{
		Com: &ComSysConf{
			Nogo:   "system_inhibited",
			Rdpc:   "radar_data_processor_chain1",
			Rdpr:   "default_situation",
			Ovlrdp: "no_overload",
			Ovlxmt: "no_overload",
			Msc:    "monitoring_system_connected",
			Tsv:    "time_source_valid",
		},
		Psr: nil,
		Ssr: nil,
		Mds: &MdsSysConf{
			Ant:    "antenna_1",
			ChAB:   "channel_a_only_selected",
			Ovlsur: "no_overload",
			Msc:    "monitoring_system_connected",
			Scf:    "channel_a_in_use",
			Dlf:    "channel_a_in_use",
			Ovlscf: "no_overload",
			Ovldlf: "no_overload",
		},
	}

	// Act
	res, err := systemConfiguration(input)

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

func TestCat034Model_MessageCountValues(t *testing.T) {
	// Arrange
	type dataTest struct {
		Name   string
		input  []byte
		output []MessageCounter
		err    error
	}
	dataset := []dataTest{
		{
			Name:  "testcase: rep 1, no_detection",
			input: []byte{0x01, 0x00, 0x0F},
			output: []MessageCounter{
				{
					Type:    "no_detection",
					Counter: 15,
				},
			},
			err: nil,
		},
		{
			Name: "testcase: rep 21, all type",
			input: []byte{0x15,
				0x00, 0x0F,
				0x08, 0x0F, // 00001 000
				0x10, 0x0F, // 00010 000
				0x18, 0x0F, // 00011 000
				0x20, 0x0F, // 00100 000
				0x28, 0x0F, // 00101 000
				0x30, 0x0F, // 00110 000
				0x38, 0x0F, // 00111 000
				0x40, 0x0F, // 01000 000
				0x48, 0x0F, // 01001 000
				0x50, 0x0F, // 01010 000
				0x58, 0x0F, // 01011 000
				0x60, 0x0F, // 01100 000
				0x68, 0x0F, // 01101 000
				0x70, 0x0F, // 01110 000
				0x78, 0x0F, // 01111 000
				0x80, 0x0F, // 10000 000
				0x88, 0x0F, // 10001 000
				0x90, 0x0F, // 10010 000
				0x98, 0x0F, // 10011 000
				0xa0, 0x0F, // 10100 000
				0xa8, 0x0F, // 10101 000
			},
			output: []MessageCounter{
				{Type: "no_detection", Counter: 15},
				{Type: "single_psr_target_reports", Counter: 15},
				{Type: "single_ssr_target_reports", Counter: 15},
				{Type: "ssr_psr_target_reports", Counter: 15},
				{Type: "single_all_call_target_reports", Counter: 15},
				{Type: "single_roll_call_target_reports", Counter: 15},
				{Type: "all_call_psr_target_reports", Counter: 15},
				{Type: "roll_call_psr_target_reports", Counter: 15},
				{Type: "filter_for_weather_data", Counter: 15},
				{Type: "filter_for_jamming_strobe", Counter: 15},
				{Type: "filter_for_psr_data", Counter: 15},
				{Type: "filter_for_ssr_mode_s_data", Counter: 15},
				{Type: "filter_for_ssr_mode_s_psr_data", Counter: 15},
				{Type: "filter_for_enhanced_surveillance_data", Counter: 15},
				{Type: "filter_for_psr_enhanced_surveillance", Counter: 15},
				{Type: "filter_for_psr_enhanced_surveillance_ssr_mode_s_data_not_in_area", Counter: 15},
				{Type: "filter_for_psr_enhanced_surveillance__all_ssr_mode_s_data", Counter: 15},
				{Type: "re_interrogations_per_sector", Counter: 15},
				{Type: "bds_swap_and_wrong_df_replies_per_sector", Counter: 15},
				{Type: "mode_ac_fruit_per_sector", Counter: 15},
				{Type: "mode_s_fruit_per_sector", Counter: 15},
			},
			err: nil,
		},
		{
			Name:  "testcase: rep 1, error unknown",
			input: []byte{0x01, 0xb8, 0x0d},
			output: []MessageCounter{
				{
					Type:    "unknown",
					Counter: 13,
				},
			},
			err: ErrTypeUnknown,
		},
	}

	for _, row := range dataset {
		// Act
		res, err := messageCountValues(row.input)

		// Assert
		if err != row.err {
			t.Errorf("FAIL: error = %v; Expected: %v", err, row.err)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, row.err)
		}

		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf("FAIL: %v; \nExpected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat034Model_GenericPolarWindow(t *testing.T) {
	// Arrange
	input := [8]byte{0x00, 0x00, 0x64, 0x00, 0x00, 0x00, 0x27, 0x10}
	output := GenericPolarWindow{
		RhoStart:   0,
		RhoEnd:     100,
		ThetaStart: 0,
		ThetaEnd:   55,
	}

	// Act
	res := genericPolarWindow(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}

}

func TestCat034Model_DataFilter(t *testing.T) {
	// Arrange
	type dataTest struct {
		input  [1]byte
		output string
	}
	dataset := []dataTest{
		{[1]byte{0x00}, "invalid_value"},
		{[1]byte{0x01}, "filter_weather_data"},
		{[1]byte{0x02}, "filter_jamming_strobe"},
		{[1]byte{0x03}, "filter_psr_data"},
		{[1]byte{0x04}, "filter_ssr_modes_data"},
		{[1]byte{0x05}, "filter_ssr_modes_psr_data"},
		{[1]byte{0x06}, "enhanced_surveillance_data"},
		{[1]byte{0x07}, "filter_psr_enhanced_surveillance_data"},
		{[1]byte{0x08}, "filter_psr_enhanced_surveillance_ssr_modes_data_not_in_area_prime_interest"},
		{[1]byte{0x09}, "filter_psr_enhanced_surveillance_all_ssr_modes_data"},
		{[1]byte{0x10}, "error_undefined"},
	}
	for _, row := range dataset {
		// Act
		res, err := dataFilter(row.input)

		// Assert
		if err != nil {
			t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
		} else {
			t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
		}

		if res != row.output {
			t.Errorf("FAIL: %v; Expected: %v", res, row.output)
		} else {
			t.Logf("SUCCESS: %v; Expected: %v", res, row.output)
		}
	}
}

func TestCat034Model_Position3DofDataSource(t *testing.T) {
	// Arrange
	input := [8]byte{0x00, 0x8a, 0x15, 0x50, 0xf8, 0x16, 0x23, 0x32}
	output := Pos3D{
		Latitude:  29.976482,
		Longitude: 31.131310,
		Height:    138,
	}

	// Act
	res := position3DofDataSource(input)

	// Assert
	if res != output {
		t.Errorf("FAIL: %v; Expected: %v", res, output)
	} else {
		t.Logf("SUCCESS: %v; Expected: %v", res, output)
	}
}

func TestCat034Model_SystemProcessingMode(t *testing.T) {
	// Arrange
	type dataTest struct {
		Name   string
		input  []byte
		output SysProcess
	}
	dataset := []dataTest{
		{
			Name:  "testcase 1: full subfield",
			input: []byte{0x9c, 0x00, 0x00, 0x00, 0x00},
			output: SysProcess{
				ComSysPro: &ComSysPro{
					Redrdp: "no_reduction_active",
					Redxmt: "no_reduction_active",
				},
				Psr: &PsrSysPro{
					Pol:    "linear_polarization",
					Redrad: "no_reduction_active",
					Stc:    "stcMap_1",
				},
				Ssr: &SsrSysPro{Redrad: "no_reduction_active"},
				Mds: &MdsSysPro{
					Redrad: "no_reduction_active",
					Clu:    "autonomous",
				},
			},
		},
		{
			Name:  "testcase 2: full subfield",
			input: []byte{0x9c, 0x12, 0x94, 0x20, 0x30},
			output: SysProcess{
				ComSysPro: &ComSysPro{
					Redrdp: "reduction_step_1_active",
					Redxmt: "reduction_step_1_active",
				},
				Psr: &PsrSysPro{
					Pol:    "circular_polarization",
					Redrad: "reduction_step_1_active",
					Stc:    "stcMap_2",
				},
				Ssr: &SsrSysPro{Redrad: "reduction_step_1_active"},
				Mds: &MdsSysPro{
					Redrad: "reduction_step_1_active",
					Clu:    "no_autonomous",
				},
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := systemProcessingMode(row.input)

		// Assert
		in := reflect.ValueOf(res)
		out := reflect.ValueOf(row.output)
		typeOfS := in.Type()

		for i := 0; i < in.NumField(); i++ {
			if reflect.DeepEqual(in.Field(i).Interface(), out.Field(i).Interface()) == false {
				t.Errorf("FAIL: %s: %s - %v; Expected: %v", row.Name, typeOfS.Field(i).Name, in.Field(i).Interface(), out.Field(i).Interface())
			} else {
				t.Logf("SUCCESS: %s - %v; Expected: %v", typeOfS.Field(i).Name, in.Field(i).Interface(), out.Field(i).Interface())
			}
		}

	}
}
