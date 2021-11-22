package transform

import (
	"encoding/hex"
	"errors"
	"github.com/mokhtarimokhtar/goasterix/commbds"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"strconv"
	"strings"
)

var (
	// ErrCharUnknown reports which not found equivalent International Alphabet 5 char.
	ErrCharUnknown = errors.New("[ASTERIX Error] char unknown")
)

// FL Flight Level, unit of altitude (expressed in 100's of feet)
type FL struct {
	V     string  `json:"v"`
	G     string  `json:"g"`
	Level float64 `json:"level"`
}

type Mode3A struct {
	Squawk string `json:"squawk"`
	V      string `json:"v"`
	G      string `json:"g"`
	L      string `json:"l"`
}

type Velocity struct {
	GroundSpeed float64 `json:"groundSpeed"`
	Heading     float64 `json:"heading"`
}

type PolarPosition struct {
	Rho   float64 `json:"rho"`
	Theta float64 `json:"theta"`
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

type Status struct {
	CNF string `json:"cnf,omitempty"`
	RAD string `json:"rad,omitempty"`
	DOU string `json:"dou,omitempty"`
	MAH string `json:"mah,omitempty"`
	CDM string `json:"cdm,omitempty"`
	TRE string `json:"tre,omitempty"`
	GHO string `json:"gho,omitempty"`
	SUP string `json:"sup,omitempty"`
	TCC string `json:"tcc,omitempty"`
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

// Write writes a single ASTERIX Record to Cat048Model.
// Items is a slice of Items DataField.
func (data *Cat048Model) write(items []uap.DataField) {
	for _, item := range items {
		switch item.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		case 2:
			// decode timeOfDay
			var payload [3]byte
			copy(payload[:], item.Payload[:])
			data.TimeOfDay, _ = timeOfDay(payload)
		// todo: case 3
		case 4:
			// decode PolarPosition
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp := rhoTheta(payload)
			data.RhoTheta = &tmp
		case 5:
			// decode Mode3aVGL
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := mode3ACodeVGL(payload)
			data.Mode3ACode = &tmp
		case 6:
			// decode Flight Level
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := flightLevel(payload)
			data.FlightLevel = &tmp
		case 7:
			// decode Radar Plot Characteristics
			tmp := radarPlotCharacteristics(item.Payload)
			data.RadarPlotCharacteristics = &tmp
		case 8:
			// decode AircraftAddress
			// AircraftAddress returns the hexadecimal code in string format.
			// Aircraft address (24-bits Mode S address) assigned uniquely to each aircraft.
			data.AircraftAddress = strings.ToUpper(hex.EncodeToString(item.Payload[:]))
		case 9:
			// decode Aircraft Identification
			var payload [6]byte
			copy(payload[:], item.Payload[:])
			data.AircraftIdentification, _ = modeSIdentification(payload)
		case 10:
			data.BDSRegisterData, _ = modeSMBData(item.Payload)
		case 11:
			// decode trackNumber
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			data.TrackNumber = trackNumber(payload)
		case 12:
			// decode Cartesian Coordinates
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := cartesianXY(payload)
			data.CartesianXY = &tmp
		case 13:
			// decode trackVelocity
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := trackVelocity(payload)
			data.TrackVelocity = &tmp
		case 14:
			// decode Track Status
			tmp := trackStatus(item.Payload[:])
			data.TrackStatus = &tmp
		// todo: case 15
		// todo: case 16
		// todo: case 17
		// todo: case 18
		// todo: case 19
		// todo: case 20
		case 21:
			// decode Communications/ACAS Capability and Flight Status
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := comACASCapabilityFlightStatus(payload)
			data.ComACASCapabilityFlightStatus = &tmp
		}
		// todo: case 22
		// todo: case 23
		// todo: case 24
		// todo: case 25
		// todo: case 26
		// todo: case 27
		// todo: case 28
	}
}

// rhoTheta returns a slice [Rho,Theta] of float64,
// Rho NM (1 bit = 1/256 NM). Theta deg (1 bit = approx. 0.0055°)
// Measured position of an aircraft in local polar co-ordinates.
func rhoTheta(data [4]byte) PolarPosition {
	var rt PolarPosition
	rt.Rho = float64(uint16(data[0])<<8+uint16(data[1])) / 256
	rt.Theta = float64(uint16(data[2])<<8+uint16(data[3])) * 0.0055
	return rt
}

// mode3ACodeVGL returns codes VGL in order.
// Squawk returns a string.
// It converts Mode-3/A reply in octal representation to a string.
// Mode-3/A code converted into octal representation.
// Ref: 5.2.10 Records Item I048/070, Mode-3/A TransponderRegisterNumber in Octal Representation.
func mode3ACodeVGL(data [2]byte) Mode3A {
	var mode3A Mode3A
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

	return mode3A
}

// flightLevel returns a float64 (1 bit = 1/4 FL).
// Flight Level into binary representation converted in an integer (16bits).
func flightLevel(data [2]byte) FL {
	var fl FL
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
	return fl
}

// radarPlotCharacteristics returns a map of float64,
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
func radarPlotCharacteristics(data []byte) PlotCharacteristics {
	var rpc PlotCharacteristics
	offset := 1

	if data[0]&0x80 != 0 {
		rpc.SRL = float64(data[offset]) * 0.044
		offset++
	}
	if data[0]&0x40 != 0 {
		rpc.SRR = data[offset]
		offset++
	}
	if data[0]&0x20 != 0 {
		rpc.SAM = int8(data[offset])
		offset++
	}
	if data[0]&0x10 != 0 {
		rpc.PRL = float64(data[offset]) * 0.044
		offset++
	}
	if data[0]&0x08 != 0 {
		rpc.PAM = int8(data[offset])
		offset++
	}
	if data[0]&0x04 != 0 {
		rpc.RPD = float64(int8(data[offset])) / 256
		offset++
	}
	if data[0]&0x02 != 0 {
		rpc.APD = float64(int8(data[offset])) * 0.021972656
	}

	return rpc
}

type ModeSMB struct {
	Rep  uint8
	BDSs []*commbds.Bds
}

func (mb *ModeSMB) Decode(data []byte) (err error) {
	rep := data[0]
	mb.Rep = rep

	for i := 0; i < int(rep*8); i = i + 8 {
		mbData := data[i+1 : i+9]
		var data [8]byte
		copy(data[:], mbData) // convert slice to array of 8 bytes
		bds := new(commbds.Bds)

		err = bds.Decode(data)
		mb.BDSs = append(mb.BDSs, bds)
	}
	return err
}

// modeSMBData returns an array of map.
// Mode S Comm B data as extracted from the aircraft transponder.
func modeSMBData(data []byte) (msb []*commbds.Bds, err error) {
	modeSMBData := new(ModeSMB)
	err = modeSMBData.Decode(data)
	msb = modeSMBData.BDSs
	return msb, err
}

// trackNumber returns an integer.
// An integer value representing a unique reference to a track record within a particular track file.
//func trackNumber(data [2]byte) (tn uint16, err error) {
//	tn = uint16(data[0])<<8 + uint16(data[1])
//	return tn, nil
//}

// cartesianXY returns a slice [X,Y] of float64 NM (1 bit = 1/128 NM) Max range = ±256 NM.
// Calculated position of an aircraft in cartesianXY co-ordinates.
func cartesianXY(data [4]byte) (pos CartesianXYPosition, err error) {
	pos.X = float64(int16(data[0])<<8+int16(data[1])) / 128
	pos.Y = float64(int16(data[2])<<8+int16(data[3])) / 128
	return pos, nil
}

// trackVelocity returns a slice [GroundSpeed,Heading] of float64.
// GroundSpeed returns float64 NM/s (1 bit = 2^-14 NM/s).
// Heading returns a float64 deg (1 bit = approx. 0.0055°).
// Calculated track Velocity expressed in polar co-ordinates.
func trackVelocity(data [4]byte) (v Velocity, err error) {
	v.GroundSpeed = float64(uint16(data[0])<<8+uint16(data[1])) * 0.000061035
	v.Heading = float64(uint16(data[2])<<8+uint16(data[3])) * 0.0055
	return v, nil
}

// trackStatus returns a map of uint8, CNF, RAD, DOU, MAH, CDM id exist: TRE, GHO, SUP, TCC.
// Status of monoradar track (PSR and/or SSR updated).
func trackStatus(data []byte) Status {
	var ts Status

	if data[0]&0x80 != 0 {
		ts.CNF = "tentative_track"
	} else {
		ts.CNF = "confirmed_track"
	}

	tmp := data[0] & 0x60 >> 5
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

	if data[0]&0x10 != 0 {
		ts.DOU = "low_confidence"
	} else {
		ts.DOU = "normal_confidence"
	}

	if data[0]&0x08 != 0 {
		ts.MAH = "horizontal_man_sensed"
	} else {
		ts.MAH = "no_horizontal_man_sensed"
	}

	tmp = data[0] & 0x06 >> 1
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

	if data[0]&0x01 != 0 {
		if data[1]&0x80 != 0 {
			ts.TRE = "end_of_track_lifetime"
		} else {
			ts.TRE = "track_still_alive"
		}

		if data[1]&0x40 != 0 {
			ts.GHO = "ghost_target_track"
		} else {
			ts.GHO = "true_target_track"
		}

		if data[1]&0x20 != 0 {
			ts.SUP = "yes"
		} else {
			ts.SUP = "no"
		}
		if data[1]&0x10 != 0 {
			ts.TCC = "slant_range_correction_used"
		} else {
			ts.TCC = "radar_plane"
		}
	}

	return ts
}

// todo: targetReportDescriptor

//todo: method 5.2.3 Records Item I048/030, Warning/Error Conditions
//todo: method 5.2.6 Records Item I048/050, Mode-2 TransponderRegisterNumber in Octal Representation
//todo: method 5.2.7 Records Item I048/055, Mode-1 TransponderRegisterNumber in Octal Representation
//todo: method 5.2.8 Records Item I048/060, Mode-2 TransponderRegisterNumber Confidence Indicator
//todo: method 5.2.9 Records Item I048/065, Mode-1 TransponderRegisterNumber Confidence Indicator
//todo: method 5.2.11 Records Item I048/080, Mode-3/A TransponderRegisterNumber Confidence Indicator
//todo: method 5.2.13 Records Item I048/100, Mode-C TransponderRegisterNumber and TransponderRegisterNumber Confidence Indicator
//todo: method 5.2.14 Records Item I048/110, Height Measured by a 3D Radar
//todo: method 5.2.15 Records Item I048/120, Radial Doppler Speed
//todo: method 5.2.21 Records Item I048/210, Track Quality

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
func comACASCapabilityFlightStatus(data [2]byte) ACASCapaFlightStatus {
	var a ACASCapaFlightStatus

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

	return a
}
