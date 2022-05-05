package model

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/commbds"
	"github.com/mokhtarimokhtar/goasterix/item"
	"strconv"
	"strings"
)

type TargetReport struct {
	TYP    string `json:"typ"`
	SIM    string `json:"sim"`
	RDP    string `json:"rdp"`
	SPI    string `json:"spi"`
	RAB    string `json:"rab"`
	TST    string `json:"tst,omitempty"`
	ERR    string `json:"err,omitempty"`
	XPP    string `json:"xpp,omitempty"`
	ME     string `json:"me,omitempty"`
	MI     string `json:"mi,omitempty"`
	FOEFRI string `json:"foefri,omitempty"`
}

type PolarPosition struct {
	Rho   float64 `json:"rho"`
	Theta float64 `json:"theta"`
}

type Cat048Model struct {
	SacSic                        *SourceIdentifier     `json:"sourceIdentifier,omitempty"`
	AircraftAddress               string                `json:"aircraftAddress,omitempty"`
	AircraftIdentification        string                `json:"aircraftIdentification,omitempty"`
	TimeOfDay                     float64               `json:"timeOfDay,omitempty"`
	RhoTheta                      *PolarPosition        `json:"rhoTheta,omitempty"`
	CartesianXY                   *CartesianXYPosition  `json:"cartesianXY,omitempty"`
	FlightLevel                   *FL                   `json:"flightLevel,omitempty"`
	RadarPlotCharacteristics      *PlotCharacteristics  `json:"radarPlotCharacteristics,omitempty"`
	Mode3ACode                    *Mode3A               `json:"mode3ACode,omitempty"`
	TrackNumber                   uint16                `json:"trackNumber,omitempty"`
	TrackVelocity                 *Velocity             `json:"trackVelocity,omitempty"`
	TrackStatus                   *Status               `json:"trackStatus,omitempty"`
	BDSRegisterData               []*commbds.Bds        `json:"bdsRegisterData,omitempty"`
	ComACASCapabilityFlightStatus *ACASCapaFlightStatus `json:"comAcasCapabilityFlightStatus,omitempty"`
}

func (m Cat048Model) String() string {
	data, _ := json.Marshal(m)
	return string(data)
}

func (m *Cat048Model) write(rec goasterix.IRecord) {
	dataItems := rec.GetItems()
	for _, dataItem := range dataItems {
		switch dataItem.GetFrn() {
		case 1:
			// decode sac sic
			tmp := new(SourceIdentifier)
			sub := dataItem.GetSubItems()
			tmp.Sac = sub[0].Data[0]
			tmp.Sic = sub[1].Data[0]
			m.SacSic = tmp
		case 2:
			// decode timeOfDay
			sub := dataItem.GetSubItems()
			m.TimeOfDay, _ = timeOfDay(sub[0].Data)

		//case 3:
		//	// decode timeOfDay
		//	sub := dataItem.GetSubItems()
		//	fmt.Println(sub)

		case 4:
			// rhoTheta returns a slice [Rho,Theta] of float64,
			// Rho NM (1 bit = 1/256 NM). Theta deg (1 bit = approx. 0.0055°)
			// Measured position of an aircraft in local polar co-ordinates.
			sub := dataItem.GetSubItems()
			rt := new(PolarPosition)
			rt.Rho = float64(uint16(sub[0].Data[0])<<8+uint16(sub[0].Data[1])) / 256
			rt.Theta = float64(uint16(sub[1].Data[0])<<8+uint16(sub[1].Data[1])) * 0.0055
			m.RhoTheta = rt
		case 5:
			sub := dataItem.GetSubItems()
			m.Mode3ACode = getMode3ACode(sub)
		case 6:
			// decode Flight Level
			sub := dataItem.GetSubItems()
			m.FlightLevel = getFlightLevel(sub)
		case 7:
			// decode Radar Plot Characteristics
			sub := dataItem.GetSubItems()
			m.RadarPlotCharacteristics = getRadarPlotCharacteristics(sub)
		case 8:
			// decode AircraftAddress
			// AircraftAddress returns the hexadecimal code in string format.
			// Aircraft address (24-bits Mode S address) assigned uniquely to each aircraft.
			sub := dataItem.GetSubItems()
			m.AircraftAddress = strings.ToUpper(hex.EncodeToString(sub[0].Data))
		case 9:
			// decode Aircraft Identification
			sub := dataItem.GetSubItems()
			m.AircraftIdentification, _ = GetModeSIdentification(sub)
		case 10:
			sub := dataItem.GetSubItems()
			m.BDSRegisterData, _ = getModeSMBData(sub)
		case 11:
			// trackNumber returns an integer (Identification of a track).
			// An integer value representing a unique reference to a track record within a particular track file.
			sub := dataItem.GetSubItems()
			m.TrackNumber = uint16(sub[0].Data[0])<<8 + uint16(sub[0].Data[1])
		case 12:
			// cartesianXY returns a slice [X,Y] of float64 NM (1 bit = 1/128 NM) Max range = ±256 NM.
			// Calculated position of an aircraft in cartesianXY co-ordinates.
			sub := dataItem.GetSubItems()
			pos := new(CartesianXYPosition)
			pos.X = float64(int16(sub[0].Data[0])<<8+int16(sub[0].Data[1])) / 128
			pos.Y = float64(int16(sub[1].Data[0])<<8+int16(sub[1].Data[1])) / 128
			m.CartesianXY = pos
		case 13:
			// trackVelocity returns a slice [GroundSpeed,Heading] of float64.
			// GroundSpeed returns float64 NM/s (1 bit = 2^-14 NM/s).
			// Heading returns a float64 deg (1 bit = approx. 0.0055°).
			// Calculated track Velocity expressed in polar co-ordinates.
			sub := dataItem.GetSubItems()
			v := new(Velocity)
			v.GroundSpeed = float64(uint16(sub[0].Data[0])<<8+uint16(sub[0].Data[1])) * 0.000061035
			v.Heading = float64(uint16(sub[1].Data[0])<<8+uint16(sub[1].Data[1])) * 0.0055
			m.TrackVelocity = v
		case 14:
			// decode Track Status
			sub := dataItem.GetSubItems()
			m.TrackStatus = getTrackStatus(sub)

		case 21:
			// decode Communications/ACAS Capability and Flight Status
			sub := dataItem.GetSubItems()
			m.ComACASCapabilityFlightStatus = comACASCapabilityFlightStatus(sub)
		}
	}
}

type ACASCapaFlightStatus struct {
	COM  string `json:"com"`
	STAT string `json:"stat"`
	SI   string `json:"si"`
	MSSC string `json:"mssc"`
	ARC  string `json:"arc"`
	AIC  string `json:"aic"`
	B1A  string `json:"b1a"`
	B1B  string `json:"b1b"`
}

// comACASCapabilityFlightStatus returns a map of sting, COM, STAT, SI, MSSC, ARC, AIC, B1A, BB.
// COM is an integer of Communications capability of the transponder from 0 to 4.
// STAT is an integer of Flight Status from 0 to 5.
// SI is an integer of SI/II Transponder Capability 0 or 1.
// MSSC is an integer of Mode-S Specific Service Capability 0 or 1 (no or yes).
// ARC is an integer of Altitude reporting capability 0 or 1 (100ft or 25ft resolution)
// AIC is an integer of Aircraft identification capability 0 or 1 (no or yes).
// B1A is a byte of BDS 1,0 bit 16.
// B1B is a byte of BDS 1,0 bits 37/40.
// Communications capability of the transponder, capability of the on board ACAS equipment and flight Status.
// Ref: 5.2.23 Records Item I048/230, Communications/ACAS Capability and Flight Status.
func comACASCapabilityFlightStatus(subItems []item.SubItem) *ACASCapaFlightStatus {
	var a = new(ACASCapaFlightStatus)

	for _, sub := range subItems {
		switch sub.Name {
		case "COM":
			a.COM = GetValueOfStruct(sub, "no_communications_capability", "comm_a_and_comm_b_capability",
				"comm_a_and_comm_b_and_uplink_elm", "comm_a_and_comm_b_and_uplink_elm_and_downlink_elm",
				"level_5_transponder_capability", "not_assigned", "not_assigned", "not_assigned")
		case "STAT":
			a.STAT = GetValueOfStruct(sub, "no_alert_no_spi_aircraft_airborne", "no_alert_no_spi_aircraft_on_ground",
				"alert_no_spi_aircraft_airborne", "alert_no_spi_aircraft_on_ground", "alert_spi_aircraft_airborne_or_on_ground",
				"no_alert_spi_aircraft_airborne_or_on_ground", "not_assigned", "unknown")
		case "SI":
			a.SI = GetValueOfStruct(sub, "si_code_capable", "sii_code_capable")
		case "MSSC":
			a.MSSC = GetValueOfStruct(sub, "no", "yes")
		case "ARC":
			a.ARC = GetValueOfStruct(sub, "100_ft_resolution", "25_ft_resolution")
		case "AIC":
			a.AIC = GetValueOfStruct(sub, "no", "yes")
		case "B1A":
			a.B1A = GetValueOfStruct(sub, "0", "1")
		case "B1B":
			a.B1B = strconv.Itoa(int(sub.Data[0]))
		}
	}

	/*
		com := data[0] & 0xE0 >> 5
		switch com {
		case 0:
			a.COM = "no_communications_capability"
		case 1:
			a.COM = "comm_a_and_comm_b_capability"
		case 2:
			a.COM = "comm_a_and_comm_b_and_uplink_elm"
		case 3:
			a.COM = "comm_a_and_comm_b_and_uplink_elm_and_downlink_elm"
		case 4:
			a.COM = "level_5_transponder_capability"
		case 5, 6, 7:
			a.COM = "not_assigned"
		}

		stat := data[0] & 0x1C >> 2
		switch stat {
		case 0:
			a.STAT = "no_alert_no_spi_aircraft_airborne"
		case 1:
			a.STAT = "no_alert_no_spi_aircraft_on_ground"
		case 2:
			a.STAT = "alert_no_spi_aircraft_airborne"
		case 3:
			a.STAT = "alert_no_spi_aircraft_on_ground"
		case 4:
			a.STAT = "alert_spi_aircraft_airborne_or_on_ground"
		case 5:
			a.STAT = "no_alert_spi_aircraft_airborne_or_on_ground"
		case 6:
			a.STAT = "not_assigned"
		case 7:
			a.STAT = "unknown"
		}

		si := data[0] & 0x02 >> 1
		if si == 0 {
			a.SI = "si_code_capable"
		} else {
			a.SI = "sii_code_capable"
		}

		mssc := data[1] & 0x80 >> 7
		if mssc == 0 {
			a.MSSC = "no"
		} else {
			a.MSSC = "yes"
		}

		arc := data[1] & 0x40 >> 6
		if arc == 0 {
			a.ARC = "100_ft_resolution"
		} else {
			a.ARC = "25_ft_resolution"
		}

		aic := data[1] & 0x20 >> 5
		if aic == 0 {
			a.AIC = "no"
		} else {
			a.AIC = "yes"
		}

		b1a := data[1] & 0x10 >> 4
		if b1a == 0 {
			a.B1A = "0"
		} else {
			a.B1A = "1"
		}

		b1b := data[1] & 0x0F
		a.B1B = strconv.Itoa(int(b1b))
	*/

	return a
}

type CartesianXYPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Velocity struct {
	GroundSpeed float64 `json:"groundSpeed"`
	Heading     float64 `json:"heading"`
}

// getModeSMBData returns an array of map.
// Mode S Comm B data as extracted from the aircraft transponder.
func getModeSMBData(subItems []item.SubItem) ([]*commbds.Bds, error) {
	var err error
	var msb []*commbds.Bds
	msb = make([]*commbds.Bds, 0, len(subItems))

	for _, sub := range subItems {
		var data [8]byte
		copy(data[:], sub.Data) // convert slice to array of 8 bytes
		bds := new(commbds.Bds)
		err = bds.Decode(data)
		msb = append(msb, bds)
	}

	return msb, err
}

func GetModeSIdentification(subItems []item.SubItem) (string, error) {
	var err error
	var buf bytes.Buffer
	buf.Reset()

	for _, sub := range subItems {
		if str, found := TableIA5[sub.Data[0]]; found {
			buf.WriteString(str)
		} else {
			err = ErrCharUnknown
		}
	}

	/*
		ch1 := data[0] & 0xFC >> 2
		str1, found1 := TableIA5[ch1]

		ch2 := data[0]&0x03<<4 + data[1]&0xF0>>4
		str2, found2 := TableIA5[ch2]

		ch3 := data[1]&0x0F<<2 + data[2]&0xC0>>6
		str3, found3 := TableIA5[ch3]

		ch4 := data[2] & 0x3F
		str4, found4 := TableIA5[ch4]

		ch5 := data[3] & 0xFC >> 2
		str5, found5 := TableIA5[ch5]

		ch6 := data[3]&0x03<<4 + data[4]&0xF0>>4
		str6, found6 := TableIA5[ch6]

		ch7 := data[4]&0x0F<<2 + data[5]&0xC0>>6
		str7, found7 := TableIA5[ch7]

		ch8 := data[5] & 0x3F
		str8, found8 := TableIA5[ch8]

		if !found1 || !found2 || !found3 || !found4 || !found5 || !found6 || !found7 || !found8 {
			err = ErrCharUnknown
		}

		//s = strings.TrimSpace(str1 + str2 + str3 + str4 + str5 + str6 + str7 + str8)
		s = str1 + str2 + str3 + str4 + str5 + str6 + str7 + str8
		return s, err
	*/

	return buf.String(), err
}

type Status struct {
	CNF string `json:"cnf"`
	RAD string `json:"rad"`
	DOU string `json:"dou"`
	MAH string `json:"mah"`
	CDM string `json:"cdm"`
	TRE string `json:"tre,omitempty"`
	GHO string `json:"gho,omitempty"`
	SUP string `json:"sup,omitempty"`
	TCC string `json:"tcc,omitempty"`
}

// getTrackStatus returns a struct of string, CNF, RAD, DOU, MAH, CDM id exist: TRE, GHO, SUP, TCC.
// Status of monoradar track (PSR and/or SSR updated).
func getTrackStatus(subItems []item.SubItem) *Status {
	var ts = new(Status)

	for _, sub := range subItems {
		switch sub.Name {
		case "CNF":
			ts.CNF = GetValueOfStruct(sub, "confirmed_track", "tentative_track")
		case "RAD":
			ts.RAD = GetValueOfStruct(sub, "combined_track", "psr_track", "ssr_modes_track", "invalid")
		case "DOU":
			ts.DOU = GetValueOfStruct(sub, "normal_confidence", "low_confidence")
		case "MAH":
			ts.MAH = GetValueOfStruct(sub, "no_horizontal_man_sensed", "horizontal_man_sensed")
		case "CDM":
			ts.CDM = GetValueOfStruct(sub, "maintaining", "climbing", "descending", "unknown")
		case "TRE":
			ts.TRE = GetValueOfStruct(sub, "track_still_alive", "end_of_track_lifetime")
		case "GHO":
			ts.GHO = GetValueOfStruct(sub, "true_target_track", "ghost_target_track")
		case "SUP":
			ts.SUP = GetValueOfStruct(sub, "no", "yes")
		case "TCC":
			ts.TCC = GetValueOfStruct(sub, "radar_plane", "slant_range_correction_used")
		}
	}

	/*
		if item.Primary[0]&0x80 != 0 {
			ts.CNF = "tentative_track"
		} else {
			ts.CNF = "confirmed_track"
		}

		tmp := item.Primary[0] & 0x60 >> 5
		switch tmp {
		case 0:
			ts.RAD = "combined_track"
		case 1:
			ts.RAD = "psr_track"
		case 2:
			ts.RAD = "ssr_modes_track"
		case 3:
			ts.RAD = "invalid"
		}

		if item.Primary[0]&0x10 != 0 {
			ts.DOU = "low_confidence"
		} else {
			ts.DOU = "normal_confidence"
		}

		if item.Primary[0]&0x08 != 0 {
			ts.MAH = "horizontal_man_sensed"
		} else {
			ts.MAH = "no_horizontal_man_sensed"
		}

		tmp = item.Primary[0] & 0x06 >> 1
		switch tmp {
		case 0:
			ts.CDM = "maintaining"
		case 1:
			ts.CDM = "climbing"
		case 2:
			ts.CDM = "descending"
		case 3:
			ts.CDM = "unknown"
		}

		if item.Secondary != nil {
			if item.Secondary[0]&0x80 != 0 {
				ts.TRE = "end_of_track_lifetime"
			} else {
				ts.TRE = "track_still_alive"
			}

			if item.Secondary[0]&0x40 != 0 {
				ts.GHO = "ghost_target_track"
			} else {
				ts.GHO = "true_target_track"
			}

			if item.Secondary[0]&0x20 != 0 {
				ts.SUP = "yes"
			} else {
				ts.SUP = "no"
			}

			if item.Secondary[0]&0x10 != 0 {
				ts.TCC = "slant_range_correction_used"
			} else {
				ts.TCC = "radar_plane"
			}
		}
	*/

	return ts
}

type Mode3A struct {
	Squawk string `json:"squawk"`
	V      string `json:"v"`
	G      string `json:"g"`
	L      string `json:"l"`
}

// getMode3ACode returns codes VGL in order.
// Squawk returns a string.
// It converts Mode-3/A reply in octal representation to a string.
// Mode-3/A code converted into octal representation.
// Ref: 5.2.10 Records Item I048/070, Mode-3/A TransponderRegisterNumber in Octal Representation.
func getMode3ACode(subItems []item.SubItem) *Mode3A {
	mode3A := new(Mode3A)

	for _, sub := range subItems {
		switch sub.Name {
		case "V":
			mode3A.V = GetValueOfStruct(sub, "code_validated", "code_not_validated")
		case "G":
			mode3A.G = GetValueOfStruct(sub, "default", "garbled_code")
		case "L":
			mode3A.L = GetValueOfStruct(sub, "code_derived_from_transponder", "code_not_extracted")
		case "Mode-3/A":
			tmp := uint16(sub.Data[0])&0x000F<<8 + uint16(sub.Data[1])&0x00FF
			mode3A.Squawk = strconv.FormatUint(uint64(tmp), 8)
		}
	}

	/*
		if data[0]&0x80 != 0 {
			mode3A.V = "code_not_validated"
		} else {
			mode3A.V = "code_validated"
		}

		if data[0]&0x40 != 0 {
			mode3A.G = "garbled_code"
		} else {
			mode3A.G = "default"
		}

		if data[0]&0x20 != 0 {
			mode3A.L = "code_not_extracted"
		} else {
			mode3A.L = "code_derived_from_transponder"
		}

		tmp := uint16(data[0])&0x000F<<8 + uint16(data[1])&0x00FF
		mode3A.Squawk = strconv.FormatUint(uint64(tmp), 8)
	*/

	return mode3A
}

// FL Flight Level, unit of altitude (expressed in 100's of feet)
type FL struct {
	V     string  `json:"v"`
	G     string  `json:"g"`
	Level float64 `json:"level"`
}

// getFlightLevel returns a float64 (1 bit = 1/4 FL).
// Flight Level into binary representation converted in an integer (16bits).
func getFlightLevel(subItems []item.SubItem) *FL {
	var fl = new(FL)

	//fl.V = GetValueOfStruct(subItems[0], "code_validated", "code_not_validated")

	for _, sub := range subItems {
		switch sub.Name {
		case "V":
			fl.V = GetValueOfStruct(sub, "code_validated", "code_not_validated")
		case "G":
			fl.G = GetValueOfStruct(sub, "default", "garbled_code")
		case "LEVEL":
			fl.Level = float64((uint16(sub.Data[0])<<8+uint16(sub.Data[1]))&0x3FFF) / 4
		}
	}

	/*
		if data[0]&0x80 != 0 {
			fl.V = "code_not_validated"
		} else {
			fl.V = "code_validated"
		}
		if data[0]&0x40 != 0 {
			fl.G = "garbled_code"
		} else {
			fl.G = "default"
		}
		fl.Level = float64((uint16(data[0])<<8+uint16(data[1]))&0x3FFF) / 4
	*/
	return fl
}

type PlotCharacteristics struct {
	SRL float64 `json:"srl,omitempty"`
	SRR uint8   `json:"srr,omitempty"`
	SAM int8    `json:"sam,omitempty"`
	PRL float64 `json:"prl,omitempty"`
	PAM int8    `json:"pam,omitempty"`
	RPD float64 `json:"rpd,omitempty"`
	APD float64 `json:"apd,omitempty"`
}

// getRadarPlotCharacteristics returns a map of float64,
// It returns according to Primary Subfield (fspec).
// SRL: SSR Plot Runlength (1 bits = 0.044 dg).
// SSR: Number of Received Replies for (M)SSR.
// SAM: Amplitude of (M)SSR Reply 1 dBm,two's complement form.
// PRL: Primary Plot Runlength (1 bits = 0.044 dg).
// PAM: Amplitude of Primary Plot 1 dBm, two's complement form.
// RPD: Difference in Range between PSR and SSR plot (1 bit = 1/256 NM), two's complement form.
// APD: Difference in Azimuth between PSR and SSR plot, two's complement form.
// Additional information on the quality of the target report.
// 5.2.16 Records Item I048/130, Radar Plot Characteristics
func getRadarPlotCharacteristics(subItems []item.SubItem) *PlotCharacteristics {
	var rpc = new(PlotCharacteristics)

	for _, sub := range subItems {
		switch sub.Name {
		case goasterix.I048130SRL:
			rpc.SRL = float64(sub.Data[0]) * 0.044
		case goasterix.I048130SRR:
			rpc.SRR = sub.Data[0]
		case goasterix.I048130SAM:
			rpc.SAM = int8(sub.Data[0])
		case goasterix.I048130PRL:
			rpc.PRL = float64(sub.Data[0]) * 0.044
		case goasterix.I048130PAM:
			rpc.PAM = int8(sub.Data[0])
		case goasterix.I048130RPD:
			rpc.RPD = float64(int8(sub.Data[0])) / 256
		case goasterix.I048130APD:
			rpc.APD = float64(int8(sub.Data[0])) * 0.021972656
		}
	}

	/*
		for _, item := range cp.Secondary {
			switch item.Meta.FRN {
			case 1:
				rpc.SRL = float64(item.Fixed.Data[0]) * 0.044
			case 2:
				rpc.SRR = item.Fixed.Data[0]
			case 3:
				rpc.SAM = int8(item.Fixed.Data[0])
			case 4:
				rpc.PRL = float64(item.Fixed.Data[0]) * 0.044
			case 5:
				rpc.PAM = int8(item.Fixed.Data[0])
			case 6:
				rpc.RPD = float64(int8(item.Fixed.Data[0])) / 256
			case 7:
				rpc.APD = float64(int8(item.Fixed.Data[0])) * 0.021972656
			}
		}
	*/

	return rpc
}
