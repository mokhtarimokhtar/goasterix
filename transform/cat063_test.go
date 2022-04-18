package transform

import (
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix/util"
	"reflect"
	"testing"

	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

func TestCat063Model_ToJsonRecord(t *testing.T) {
	// Arrange
	//bff0090c79387308290000000012000000000000000000000000
	//input := "b0 0836 429b71 0801" //FSPEC : 10110000 => B0 f6 0836 429b71 0801 10940028200094008000

	input := "bff0 090c 79387308290000000012000000000000000000000000"

	output := []byte(`{"dataSourceIdentifier":{"sac":9,"sic":12},"timeOfMessage":62064.8984375,"sensorIdentifier":{"sac":8,"sic":41},"sensorConfigStatus":{"con":"operational","psr":"psr_go","ssr":"ssr_go","mds":"mode_s_go","ads":"ads_go","mlt":"mlt_go"},"timeStampingBias":0,"modeSRangeGainAndBias":{"srg":0.00018,"srb":0},"psrRangeGainAndBias":{"prg":0,"prb":0}}`)

	uap063 := uap.Cat063V16
	data, _ := util.HexStringToByte(input)
	rec := goasterix.NewRecord()
	_, err := rec.Decode(data, uap063)

	model := new(Cat063Model)
	model.write(*rec)

	// Act
	recJson, _ := json.Marshal(model)

	// Assert
	if err != nil {
		t.Errorf("MsgFailInValue: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("MsgSuccessInValue: error: %v; Expected: %v", err, nil)
	}

	if reflect.DeepEqual(recJson, output) == false {
		t.Errorf("MsgFailInValue: %s; \nExpected: %s", recJson, output)
	} else {
		t.Logf("MsgSuccessInValue: %s; Expected: %s", recJson, output)
	}
}

func TestExtractSensorStatus(t *testing.T) {
	// Arrange
	type testCase struct {
		Name   string
		input  goasterix.Extended
		output SensorStatus
	}

	dataset := []testCase{
		{
			Name: "testcase 1",
			input: goasterix.Extended{
				Primary:   []byte{0x01},
				Secondary: []byte{0x00},
			},
			output: SensorStatus{
				CON: "operational",
				PSR: "psr_go",
				SSR: "ssr_go",
				MDS: "mode_s_go",
				ADS: "ads_go",
				MLT: "mlt_go",
				OPS: "system_released_for_operationnal_use",
				ODP: "default_no_overload",
				OXT: "default_no_overload",
				MSC: "monitoring_system_connected",
				TSV: "valid",
				NPW: "default",
			},
		},
		{
			Name: "testcase 2",
			input: goasterix.Extended{
				Primary:   []byte{0x7f},
				Secondary: []byte{0xfe},
			},
			output: SensorStatus{
				CON: "degraded",
				PSR: "psr_nogo",
				SSR: "ssr_nogo",
				MDS: "mode_s_nogo",
				ADS: "ads_nogo",
				MLT: "mlt_nogo",
				OPS: "operationnal_use_of_system_inhibited",
				ODP: "overload_in_dp",
				OXT: "overload_in_transmission_subsystem",
				MSC: "monitoring_system_disconnected",
				TSV: "invalid",
				NPW: "no_plot_being_received",
			},
		},
		{
			Name: "testcase 3",
			input: goasterix.Extended{
				Primary: []byte{0x80},
			},
			output: SensorStatus{
				CON: "initialization",
				PSR: "psr_go",
				SSR: "ssr_go",
				MDS: "mode_s_go",
				ADS: "ads_go",
				MLT: "mlt_go",
			},
		},
		{
			Name: "testcase 4",
			input: goasterix.Extended{
				Primary: []byte{0xc0},
			},
			output: SensorStatus{
				CON: "not_currently_connected",
				PSR: "psr_go",
				SSR: "ssr_go",
				MDS: "mode_s_go",
				ADS: "ads_go",
				MLT: "mlt_go",
			},
		},
	}

	for _, row := range dataset {
		// Act
		res := extractSensorStatus(row.input)

		// Assert
		if reflect.DeepEqual(res, row.output) == false {
			t.Errorf(util.MsgFailInValue, row.Name, res, row.output)
		} else {
			t.Logf(util.MsgSuccessInValue, row.Name, res, row.output)
		}
	}
}
