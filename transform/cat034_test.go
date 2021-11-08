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
		res, err := messageType(row.input)

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
