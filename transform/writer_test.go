package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"reflect"
	"testing"
)

func TestWriteModel(t *testing.T) {
	// Arrange
	input := "ffd702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 063a 0743ce5b 40 20f5"
	output := Cat048Model{
		SacSic: &SourceIdentifier{
			Sac: 8,
			Sic: 54,
		},
		AircraftAddress:        "490D01",
		AircraftIdentification: "NJE834H ",
		TimeOfDay:              34102.640625,
		RhoTheta: &PolarPosition{
			Rho:   148.77734375,
			Theta: 2.1174999999999997,
		},
		CartesianXY: nil,
		FlightLevel: &FL{
			V:     "code_validated",
			G:     "default",
			Level: 180,
		},
		RadarPlotCharacteristics: &PlotCharacteristics{
			SRR: 2,
			SAM: -73,
		},
		Mode3ACode: &Mode3A{
			Squawk: "4423",
			V:      "code_validated",
			G:      "default",
			L:      "code_derived_from_transponder",
		},
		TrackNumber: 1594,
		TrackVelocity: &Velocity{
			GroundSpeed: 0.113464065,
			Heading:     290.5485,
		},
		TrackStatus: &Status{
			CNF: "confirmed_track",
			RAD: "ssr_modes_track",
			DOU: "normal_confidence",
			MAH: "no_horizontal_man_sensed",
			CDM: "maintaining",
		},
		BDSRegisterData: nil,
		ComACASCapabilityFlightStatus: &ACASCapaFlightStatus{
			COM:  "comm_a_and_comm_b_capability",
			STAT: "no_alert_no_spi_aircraft_airborne",
			SI:   "si_code_capable",
			MSSC: "yes",
			ARC:  "25_ft_resolution",
			AIC:  "yes",
			B1A:  "1",
			B1B:  "5",
		},
	}

	uap048 := uap.Cat048V127
	data, _ := goasterix.HexStringToByte(input)
	rec := goasterix.NewRecord()
	_, err := rec.Decode(data, uap048)
	cat048Model := new(Cat048Model)

	// Act
	WriteModel(cat048Model, *rec)

	// Assert
	if err != nil {
		t.Errorf("FAIL: error = %v; Expected: %v", err, nil)
	} else {
		t.Logf("SUCCESS: error: %v; Expected: %v", err, nil)
	}
	in := reflect.ValueOf(*cat048Model)
	out := reflect.ValueOf(output)
	typeOfS := in.Type()

	for i := 0; i < in.NumField(); i++ {
		if reflect.DeepEqual(in.Field(i).Interface(), out.Field(i).Interface()) == false {
			t.Errorf("FAIL: %s - %v; Expected: %v", typeOfS.Field(i).Name, in.Field(i).Interface(), out.Field(i).Interface())
		} else {
			t.Logf("SUCCESS: %s - %v; Expected: %v", typeOfS.Field(i).Name, in.Field(i).Interface(), out.Field(i).Interface())
		}
	}
}

func TestWriteModelJSON(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []byte(`{"sourceIdentifier":{"sac":8,"sic":54},"aircraftAddress":"490D01","aircraftIdentification":"NJE834H ","timeOfDay":34102.640625,"rhoTheta":{"rho":148.77734375,"theta":2.1174999999999997},"flightLevel":{"v":"code_validated","g":"default","level":180},"radarPlotCharacteristics":{"srr":2,"sam":-73},"mode3ACode":{"squawk":"4423","v":"code_validated","g":"default","l":"code_derived_from_transponder"},"trackNumber":1594,"trackVelocity":{"groundSpeed":0.113464065,"heading":290.5485},"trackStatus":{"cnf":"confirmed_track","rad":"ssr_modes_track","dou":"normal_confidence","mah":"no_horizontal_man_sensed","cdm":"maintaining"},"bdsRegisterData":[{"transponderRegisterNumber":"60","code60":{"magneticHeading":-68,"indicatedAirspeed":302,"mach":0.632,"barometricAltitudeRate":32}},{"transponderRegisterNumber":"40","code40":{"mcpSelectAltitude":18000,"barometricPressureSetting":1013}}],"comAcasCapabilityFlightStatus":{"com":"comm_a_and_comm_b_capability","stat":"no_alert_no_spi_aircraft_airborne","si":"si_code_capable","mssc":"yes","arc":"25_ft_resolution","aic":"yes","b1a":"1","b1b":"5"}}`)

	uap048 := uap.Cat048V127
	data, _ := goasterix.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, _ = rec.Decode(data, uap048)
	cat048Model := new(Cat048Model)

	// Act
	recJson, err := WriteModelJSON(cat048Model, *rec)

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

func TestWriteModelXML(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []byte(`<Cat048Model><SacSic><sac>8</sac><sic>54</sic></SacSic><AircraftAddress>490D01</AircraftAddress><AircraftIdentification>NJE834H </AircraftIdentification><TimeOfDay>34102.640625</TimeOfDay><RhoTheta><Rho>148.77734375</Rho><Theta>2.1174999999999997</Theta></RhoTheta><FlightLevel><V>code_validated</V><G>default</G><Level>180</Level></FlightLevel><RadarPlotCharacteristics><SRL>0</SRL><SRR>2</SRR><SAM>-73</SAM><PRL>0</PRL><PAM>0</PAM><RPD>0</RPD><APD>0</APD></RadarPlotCharacteristics><Mode3ACode><Squawk>4423</Squawk><V>code_validated</V><G>default</G><L>code_derived_from_transponder</L></Mode3ACode><TrackNumber>1594</TrackNumber><TrackVelocity><GroundSpeed>0.113464065</GroundSpeed><Heading>290.5485</Heading></TrackVelocity><TrackStatus><CNF>confirmed_track</CNF><RAD>ssr_modes_track</RAD><DOU>normal_confidence</DOU><MAH>no_horizontal_man_sensed</MAH><CDM>maintaining</CDM><TRE></TRE><GHO></GHO><SUP></SUP><TCC></TCC></TrackStatus><BDSRegisterData><TransponderRegisterNumber>60</TransponderRegisterNumber><Code60><MagneticHeading>-68</MagneticHeading><MagneticHeadingStatus>true</MagneticHeadingStatus><IndicatedAirspeed>302</IndicatedAirspeed><IndicatedAirspeedStatus>true</IndicatedAirspeedStatus><Mach>0.632</Mach><MachStatus>true</MachStatus><BarometricAltitudeRate>32</BarometricAltitudeRate><BarometricAltitudeRateStatus>true</BarometricAltitudeRateStatus><InertialVerticalVelocity>0</InertialVerticalVelocity><InertialVerticalVelocityStatus>true</InertialVerticalVelocityStatus></Code60></BDSRegisterData><BDSRegisterData><TransponderRegisterNumber>40</TransponderRegisterNumber><Code40><MCPSelectAltitudeStatus>true</MCPSelectAltitudeStatus><MCPSelectAltitude>18000</MCPSelectAltitude><FMSSelectAltitudeStatus>false</FMSSelectAltitudeStatus><FMSSelectAltitude>0</FMSSelectAltitude><BarometricPressureSettingStatus>true</BarometricPressureSettingStatus><BarometricPressureSetting>1013</BarometricPressureSetting><MCPModeBitsStatus>false</MCPModeBitsStatus><VNAVMode>0</VNAVMode><ALTHOLDMode>0</ALTHOLDMode><APPROACHMode>0</APPROACHMode><TargetAltSourceBitsStatus>false</TargetAltSourceBitsStatus><TargetAltSourceBits>0</TargetAltSourceBits></Code40></BDSRegisterData><ComACASCapabilityFlightStatus><COM>comm_a_and_comm_b_capability</COM><STAT>no_alert_no_spi_aircraft_airborne</STAT><SI>si_code_capable</SI><MSSC>yes</MSSC><ARC>25_ft_resolution</ARC><AIC>yes</AIC><B1A>1</B1A><B1B>5</B1B></ComACASCapabilityFlightStatus></Cat048Model>`)

	uap048 := uap.Cat048V127
	data, _ := goasterix.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, _ = rec.Decode(data, uap048)
	cat048Model := new(Cat048Model)

	// Act
	recJson, err := WriteModelXML(cat048Model, *rec)

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
