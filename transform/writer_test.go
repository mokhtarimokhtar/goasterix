package transform

import (
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"github.com/mokhtarimokhtar/goasterix/util"
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
	data, _ := util.HexStringToByte(input)
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
	input := "c71b3b6bc1810000000022ff2102428a117f90060121450a4075756dcb4b6dcb4b31f314120dab05f04000000781dd286dcb4c15a674c596a00303"
	output := []byte(`{"aircraftOperationStatus":{"ra":"TCAS II or ACAS RA not active","tc":"no capability for Trajectory Change Reports","ts":"no capability to support Target State Reports","arv":"no capability to generate ARV-reports","cdtia":"CDTI not operational","nottcas":"TCAS not operational","sa":"Single Antenna only"},"DataSourceIdentification":{"sac":0,"sic":0},"EmitterCategory":"75000 lbs < medium a/c < 300000 lbs","TargetReportDescriptor":{"atp":"24-Bit ICAO address","arc":"25ft","rc":"Default","rab":"Report from target transponder"},"Mode3ACode":{"a2":4,"a1":8,"b2":128,"b1":1,"c4":1,"c2":2},"TimeOfMessageReceptionForPosition":56214.5859375,"TimeOfMessageReceptionForVelocity":56214.5859375,"TimeOfReportTransmission":56214.59375,"TargetAddress":"407575","QualityIndicators":{"nucrornacv":1,"nucpornic":8,"fx":{"nicbaro":1,"sil":3,"nacp":9,"fx":{"sils":"flight-hour","sda":2,"gva":2}}},"PositionWGS84":{"latitude":49.21396007879,"longitude":3.17800967302},"PositionWGS84HighRes":{"latitude":49.214513010960005,"longitude":3.17804837592},"FlightLevel":380,"BarometricVerticalRate":{"re":"Value in defined range"},"AirborneGroundVector":{"re":"Value in defined range","groundspeed":0.11724853515625,"trackangle":311.0009765625},"TargetIdentification":"EZY41YZ ","TargetStatus":{"icf":"No intent change active","lnav":"LNAV Mode not engaged","ps":"No emergency/not reported","ss":"No condition reported"},"MOPSVersion":{"vns":"supported","vn":"ED102A/DO-260B","ltt":"1090 es"}}`)

	uap021 := uap.Cat021v10
	data, _ := util.HexStringToByte(input)
	rec := new(goasterix.Record)
	_, _ = rec.Decode(data, uap021)
	cat048Model := new(Cat021Model)

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
		t.Logf("SUCCESS: %s;", recJson)
	}
}

func TestWriteModelXML(t *testing.T) {
	// Arrange
	input := "fff702 0836 429b52 a0 94c70181 0913 02d0 6002b7 490d01 38a178cf4220 02e79a5d27a00c0060a3280030a4000040 063a 0743ce5b 40 20f5"
	output := []byte(`<Cat048Model><SacSic><sac>8</sac><sic>54</sic></SacSic><AircraftAddress>490D01</AircraftAddress><AircraftIdentification>NJE834H </AircraftIdentification><TimeOfDay>34102.640625</TimeOfDay><RhoTheta><Rho>148.77734375</Rho><Theta>2.1174999999999997</Theta></RhoTheta><FlightLevel><V>code_validated</V><G>default</G><Level>180</Level></FlightLevel><RadarPlotCharacteristics><SRL>0</SRL><SRR>2</SRR><SAM>-73</SAM><PRL>0</PRL><PAM>0</PAM><RPD>0</RPD><APD>0</APD></RadarPlotCharacteristics><Mode3ACode><Squawk>4423</Squawk><V>code_validated</V><G>default</G><L>code_derived_from_transponder</L></Mode3ACode><TrackNumber>1594</TrackNumber><TrackVelocity><GroundSpeed>0.113464065</GroundSpeed><Heading>290.5485</Heading></TrackVelocity><TrackStatus><CNF>confirmed_track</CNF><RAD>ssr_modes_track</RAD><DOU>normal_confidence</DOU><MAH>no_horizontal_man_sensed</MAH><CDM>maintaining</CDM><TRE></TRE><GHO></GHO><SUP></SUP><TCC></TCC></TrackStatus><BDSRegisterData><TransponderRegisterNumber>60</TransponderRegisterNumber><Code60><MagneticHeading>-68</MagneticHeading><MagneticHeadingStatus>true</MagneticHeadingStatus><IndicatedAirspeed>302</IndicatedAirspeed><IndicatedAirspeedStatus>true</IndicatedAirspeedStatus><Mach>0.632</Mach><MachStatus>true</MachStatus><BarometricAltitudeRate>32</BarometricAltitudeRate><BarometricAltitudeRateStatus>true</BarometricAltitudeRateStatus><InertialVerticalVelocity>0</InertialVerticalVelocity><InertialVerticalVelocityStatus>true</InertialVerticalVelocityStatus></Code60></BDSRegisterData><BDSRegisterData><TransponderRegisterNumber>40</TransponderRegisterNumber><Code40><MCPSelectAltitudeStatus>true</MCPSelectAltitudeStatus><MCPSelectAltitude>18000</MCPSelectAltitude><FMSSelectAltitudeStatus>false</FMSSelectAltitudeStatus><FMSSelectAltitude>0</FMSSelectAltitude><BarometricPressureSettingStatus>true</BarometricPressureSettingStatus><BarometricPressureSetting>1013</BarometricPressureSetting><MCPModeBitsStatus>false</MCPModeBitsStatus><VNAVMode>0</VNAVMode><ALTHOLDMode>0</ALTHOLDMode><APPROACHMode>0</APPROACHMode><TargetAltSourceBitsStatus>false</TargetAltSourceBitsStatus><TargetAltSourceBits>0</TargetAltSourceBits></Code40></BDSRegisterData><ComACASCapabilityFlightStatus><COM>comm_a_and_comm_b_capability</COM><STAT>no_alert_no_spi_aircraft_airborne</STAT><SI>si_code_capable</SI><MSSC>yes</MSSC><ARC>25_ft_resolution</ARC><AIC>yes</AIC><B1A>1</B1A><B1B>5</B1B></ComACASCapabilityFlightStatus></Cat048Model>`)

	uap048 := uap.Cat048V127
	data, _ := util.HexStringToByte(input)
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
