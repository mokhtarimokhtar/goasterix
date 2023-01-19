package transform

import (
	"encoding/hex"
	"math"
	"strings"

	"github.com/mokhtarimokhtar/goasterix"
)

const (
	BYTESIZE = 8
)

type WGS84Coordinates struct {
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}

type GeometricHeight struct {
	Height      float64 `json:"height,omitempty"`
	GreaterThan bool    `json:"greaterthan,omitempty"`
}

type AirSpeed struct {
	IM       string  `json:"im,omitempty"`
	AirSpeed float64 `json:"airspeed,omitempty"`
}
type TrueAirSpeed struct {
	RE    int   `json:"re,omitempty"`
	Speed int16 `json:"speed,omitempty"`
}

// TODO: Write the potential messages/states in const
// TODO: Look into next extensions ( not clearly defined by spec)
type SecondExtensionTRD struct {
	LLC            string `json:"llc,omitempty"`
	IPC            string `json:"ipc,omitempty"`
	NOGO           string `json:"nogo,omitempty"`
	CPR            string `json:"cpr,omitempty"`
	LDPJ           string `json:"ldpj,omitempty"`
	RCF            string `json:"rcf,omitempty"`
	ThirdExtension byte   `json:"fx,omitempty"`
}
type FirstExtensionTRD struct {
	DCR             string              `json:"dcr,omitempty"`
	GBS             string              `json:"gbs,omitempty"`
	SIM             string              `json:"sim,omitempty"`
	TST             string              `json:"tst,omitempty"`
	SAA             string              `json:"saa,omitempty"`
	CL              string              `json:"cl,omitempty"`
	SecondExtension *SecondExtensionTRD `json:"fx,omitempty"`
}
type TargetReportDescriptor struct {
	ATP string             `json:"atp,omitempty"`
	ARC string             `json:"arc,omitempty"`
	RC  string             `json:"rc,omitempty"`
	RAB string             `json:"rab,omitempty"`
	FX  *FirstExtensionTRD `json:"fx,omitempty"`
}

type NIC_Version2OrHigher struct {
	NIC int    `json:"nic,omitempty"`
	AB  string `json:"ab,omitempty"`
	AC  string `json:"ac,omitempty"`
}
type PIC struct {
	PIC                       int                   `json:"pic,omitempty"`
	IntegrityContainmentBound float64               `json:"integritycontainmentbound,omitempty"`
	NUCp                      int                   `json:"nucp,omitempty"`
	NIC_DO260A                string                `json:"nic_do260a,omitempty"`
	NIC_Version2OrHigher      *NIC_Version2OrHigher `json:"nic_version2orhigher,omitempty"`
}
type ThirdExtensionQI struct {
	PIC *PIC `json:"pic,omitempty"`
	FX  byte `json:"fx,omitempty"`
}
type SecondExtensionQI struct {
	SILS string            `json:"sils,omitempty"`
	SDA  int               `json:"sda,omitempty"`
	GVA  int               `json:"gva,omitempty"`
	FX   *ThirdExtensionQI `json:"thirdextension,omitempty"`
}
type FirstExtensionQI struct {
	NICBaro int                `json:"nicbaro,omitempty"`
	SIL     int                `json:"sil,omitempty"`
	NACp    int                `json:"nacp,omitempty"`
	FX      *SecondExtensionQI `json:"fx,omitempty"`
}
type QualityIndicators struct {
	NUCrOrNACv int               `json:"nucrornacv,omitempty"`
	NUCpOrNIC  int               `json:"nucpornic,omitempty"`
	FX         *FirstExtensionQI `json:"fx,omitempty"`
}

type MOPSVersion struct {
	VNS string `json:"vns,omitempty"`
	VN  string `json:"vn,omitempty"`
	LTT string `json:"ltt,omitempty"`
}

type Mode3ACodeInOctal struct {
	A4 int `json:"a4,omitempty"`
	A2 int `json:"a2,omitempty"`
	A1 int `json:"a1,omitempty"`
	B4 int `json:"b4,omitempty"`
	B2 int `json:"b2,omitempty"`
	B1 int `json:"b1,omitempty"`
	C4 int `json:"c4,omitempty"`
	C2 int `json:"c2,omitempty"`
	C1 int `json:"c1,omitempty"`
	D4 int `json:"d4,omitempty"`
	D2 int `json:"d2,omitempty"`
	D1 int `json:"d1,omitempty"`
}

type TargetStatus struct {
	ICF  string `json:"icf,omitempty"`
	LNAV string `json:"lnav,omitempty"`
	ME   bool   `json:"me,omitempty"`
	PS   string `json:"ps,omitempty"`
	SS   string `json:"ss,omitempty"`
}

type VerticalRate struct {
	RE           string  `json:"re,omitempty"`
	VerticalRate float32 `json:"verticalrate,omitempty"`
}

type AirborneGroundVector struct {
	RE          string  `json:"re,omitempty"`
	GroundSpeed float32 `json:"groundspeed,omitempty"`
	TrackAngle  float32 `json:"trackangle,omitempty"`
}

type TrajectoryIntentData struct {
	REP       int8    `json:"rep,omitempty"`
	TCA       int     `json:"tca,omitempty"`
	NC        int     `json:"nc,omitempty"`
	TCPNumber int     `json:"tcpnumber,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	PointType int     `json:"pointtype,omitempty"`
	TD        int     `json:"td,omitempty"`
	TRA       int     `json:"tra,omitempty"`
	TOA       int     `json:"toa,omitempty"`
	TOV       string  `json:"tov,omitempty"`
	TTR       string  `json:"ttr,omitempty"`
}
type TrajectoryIntentStatus struct {
	NAV string `json:"nav,omitempty"`
	NVB string `json:"nvb,omitempty"`
	FX  int    `json:"fx,omitempty"`
}
type TrajectoryIntent struct {
	TIS     string                  `json:"tis,omitempty"`
	TISBody *TrajectoryIntentStatus `json:"tisbody,omitempty"`
	TID     string                  `json:"tid,omitempty"`
	TIDBody *TrajectoryIntentData   `json:"tidbody,omitempty"`
	FX      string                  `json:"fx,omitempty"`
}

type AircraftOperationStatus struct {
	RA      string `json:"ra,omitempty"`
	TC      string `json:"tc,omitempty"`
	TS      string `json:"ts,omitempty"`
	ARV     string `json:"arv,omitempty"`
	CDTIA   string `json:"cdtia,omitempty"`
	NotTCAS string `json:"nottcas,omitempty"`
	SA      string `json:"sa,omitempty"`
}

type FirstExtensionSCC struct {
	LW LengthWidth        `json:"lw,omitempty"`
	FX *FirstExtensionSCC `json:"firstextension,omitempty"`
}
type SurfaceCapabilitiesAndCharacteristics struct {
	POA   string             `json:"poa,omitempty"`
	CDTIS string             `json:"cdtis,omitempty"`
	B2Low string             `json:"b2low,omitempty"`
	RAS   string             `json:"ras,omitempty"`
	IDENT string             `json:"ident,omitempty"`
	FX    *FirstExtensionSCC `json:"firstextension,omitempty"`
}

type ModeSMBData struct {
	REP  int8   `json:"rep,omitempty"`
	MB   string `json:"mb,omitempty"`
	BDS1 int16  `json:"bds1,omitempty"`
	BDS2 int8   `json:"bds2,omitempty"`
}

type ACASResolutionAdvisoryReport struct {
	TYP  int8  `json:"typ,omitempty"`
	STYP int8  `json:"styp,omitempty"`
	ARA  int16 `json:"ara,omitempty"`
	RAC  int16 `json:"rac,omitempty"`
	RAT  int16 `json:"rat,omitempty"`
	MTE  int16 `json:"mte,omitempty"`
	TTI  int8  `json:"tti,omitempty"`
	TID  int32 `json:"tid,omitempty"`
}

type Cat021Model struct {
	AircraftOperationStatus                        *AircraftOperationStatus               `json:"aircraftOperationStatus,omitempty"`
	DataSourceIdentification                       *SourceIdentifier                      `json:"DataSourceIdentification,omitempty"`
	ServiceIdentification                          byte                                   `json:"ServiceIdentification,omitempty"`
	ServiceManagement                              float32                                `json:"ServiceManagement,omitempty"`
	EmitterCategory                                string                                 `json:"EmitterCategory,omitempty"`
	TargetReportDescriptor                         *TargetReportDescriptor                `json:"TargetReportDescriptor,omitempty"`
	Mode3ACode                                     *Mode3ACodeInOctal                     `json:"Mode3ACode,omitempty"`
	TimeOfApplicabilityForPosition                 float64                                `json:"timeOfApplicabilityForPosition,omitempty"`
	TimeOfApplicabilityForVelocity                 float64                                `json:"timeOfApplicabilityForVelocity,omitempty"`
	TimeOfMessageReceptionForPosition              float64                                `json:"TimeOfMessageReceptionForPosition,omitempty"`
	TimeOfMessageReceptionForPositionHighPrecision *TimeOfDayHighPrecision                `json:"TimeOfMessageReceptionForPositionHighPrecision,omitempty"`
	TimeOfMessageReceptionForVelocity              float64                                `json:"TimeOfMessageReceptionForVelocity,omitempty"`
	TimeOfMessageReceptionForVelocityHighPrecision *TimeOfDayHighPrecision                `json:"TimeOfMessageReceptionForVelocityHighPrecision,omitempty"`
	TimeOfReportTransmission                       float64                                `json:"TimeOfReportTransmission,omitempty"`
	TargetAddress                                  string                                 `json:"TargetAddress,omitempty"`
	QualityIndicators                              *QualityIndicators                     `json:"QualityIndicators,omitempty"`
	TrajectoryIntent                               string                                 `json:"TrajectoryIntent,omitempty"`
	PositionWGS84                                  *WGS84Coordinates                      `json:"PositionWGS84,omitempty"`
	PositionWGS84HighRes                           *WGS84Coordinates                      `json:"PositionWGS84HighRes,omitempty"`
	MessageAmplitude                               int16                                  `json:"MessageAmplitude,omitempty"`
	GeometricHeight                                *GeometricHeight                       `json:"GeometricHeight,omitempty"`
	FlightLevel                                    float32                                `json:"FlightLevel,omitempty"`
	SelectedAltitude                               *SelectedAltitude                      `json:"SelectedAltitude,omitempty"`
	FinalStateSelectedAltitude                     *StateSelectedAltitude                 `json:"FinalStateSelectedAltitude,omitempty"`
	AirSpeed                                       *AirSpeed                              `json:"AirSpeed,omitempty"`
	TrueAirSpeed                                   *TrueAirSpeed                          `json:"TrueAirSpeed,omitempty"`
	MagneticHeading                                float64                                `json:"MagneticHeading,omitempty"`
	BarometricVerticalRate                         *VerticalRate                          `json:"BarometricVerticalRate,omitempty"`
	GeometricVerticalRate                          float64                                `json:"GeometricVerticalRate,omitempty"`
	AirborneGroundVector                           *AirborneGroundVector                  `json:"AirborneGroundVector,omitempty"`
	TrackNumber                                    uint16                                 `json:"TrackNumber,omitempty"`
	TrackAngleRate                                 float32                                `json:"TrackAngleRate,omitempty"`
	TargetIdentification                           string                                 `json:"TargetIdentification,omitempty"`
	TargetStatus                                   *TargetStatus                          `json:"TargetStatus,omitempty"`
	MOPSVersion                                    *MOPSVersion                           `json:"MPOSVersion,omitempty"`
	MetInformation                                 string                                 `json:"MetInformation,omitempty"`
	RollAngle                                      float64                                `json:"RollAngle,omitempty"`
	ModeSMBData                                    *ModeSMBData                           `json:"ModeSMBData,omitempty"`
	ACASResolutionAdvisoryReport                   *ACASResolutionAdvisoryReport          `json:"ACASResolutionAdvisoryReport,omitempty"`
	SurfaceCapabilitiesAndCharacteristic           *SurfaceCapabilitiesAndCharacteristics `json:"surfacecapabilitiesAndCharacteristics,omitempty"`
	ReceiverID                                     uint8                                  `json:"ReceiverID,omitempty"`
}

func (data *Cat021Model) write(rec goasterix.Record) {
	for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := sacSic(payload)
			data.DataSourceIdentification = &tmp
		case 2:
			tmp := targetReportDescriptor(*item.Compound)
			data.TargetReportDescriptor = &tmp
		case 3:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := trackNumber(payload)
			data.TrackNumber = tmp
		case 4:
			data.ServiceIdentification = item.Fixed.Data[0] // TODO: Double check?
		case 5:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfApplicabilityForPosition, _ = timeOfDay(payload)
		case 6:
			var payload []byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := wgs84Coordinates(payload)
			data.PositionWGS84 = &tmp
		case 7:
			var payload []byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := wgs84Coordinates(payload)
			data.PositionWGS84HighRes = &tmp
		case 8:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfApplicabilityForVelocity, _ = timeOfDay(payload)
		case 9:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := airSpeed(payload)
			data.AirSpeed = &tmp
		case 10:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := trueAirSpeed(payload)
			data.TrueAirSpeed = &tmp
		case 11:
			data.TargetAddress = strings.ToUpper(hex.EncodeToString(item.Fixed.Data))
		case 12:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessageReceptionForPosition, _ = timeOfDay(payload)
		case 13:
			// TODO: Check correctness
			var payload [4]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := timeOfDayHighPrecision(payload)
			data.TimeOfMessageReceptionForPositionHighPrecision = &tmp
		case 14:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfMessageReceptionForVelocity, _ = timeOfDay(payload)
		case 15:
			// TODO: Check correctness
			var payload [4]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := timeOfDayHighPrecision(payload)
			data.TimeOfMessageReceptionForVelocityHighPrecision = &tmp
		case 16:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := geometricHeight(payload)
			data.GeometricHeight = &tmp
		case 17:
			tmp := qualityIndicators(*item.Compound)
			data.QualityIndicators = &tmp
		case 18:
			tmp := mOPS(*item.Compound)
			data.MOPSVersion = &tmp
		case 19:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := mode3ACodeCAT021(payload)
			data.Mode3ACode = tmp
		case 20:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := rollAngle(payload)
			data.RollAngle = tmp
		case 21:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			// Method is from Cat062
			data.FlightLevel = measuredFlightLevel(payload)
		case 22:
			data.MagneticHeading = float64(uint16(item.Fixed.Data[0])<<8+uint16(item.Fixed.Data[1])) * 0.0055
		case 23:
			var payload []byte
			copy(payload, item.Fixed.Data[:])
			tmp := targetStatus(payload)
			data.TargetStatus = tmp
		case 24:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.BarometricVerticalRate = verticalRate(payload)
		case 25:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.BarometricVerticalRate = verticalRate(payload)
		case 26:
			var payload [4]byte
			copy(payload[:], item.Fixed.Data[:])
			data.AirborneGroundVector = airborneGroundVector(payload)
		case 27:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TrackAngleRate = trackAngleRate(payload)
		case 28:
			var payload [3]byte
			copy(payload[:], item.Fixed.Data[:])
			data.TimeOfReportTransmission, _ = timeOfDay(payload)
		case 29:
			var payload [6]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp, _ := modeSIdentification(payload)
			data.TargetIdentification = tmp
		case 30:
			var payload [1]byte
			copy(payload[:], item.Fixed.Data[:])
			data.EmitterCategory = emitterCategory(payload)
		case 31:
			// TODO: Implement Met Weather after clarrification with Eurocontrol
		case 32:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.SelectedAltitude = selectedAltitude(payload)
		case 33:
			var payload [2]byte
			copy(payload[:], item.Fixed.Data[:])
			data.FinalStateSelectedAltitude = finalSelectedAltitude(payload)
		case 34:
			// TODO: Implement Trajectory Intent
			// tmp := getTrajectoryIntent(*item.Compound)
			// data.TrajectoryIntent = &tmp
		case 35:
			data.ServiceManagement = float32(uint16(item.Fixed.Data[0])) * 0.5
		case 36:
			var payload [1]byte
			copy(payload[:], item.Fixed.Data[:])
			data.AircraftOperationStatus = aircraftOperationalStatus(payload)
		case 37:
			var payload []byte
			copy(payload[:], item.Fixed.Data[:])
			data.SurfaceCapabilitiesAndCharacteristic = surfaceCapabilitiesAndCharacteristics(payload)
		case 38:
			data.MessageAmplitude = goasterix.TwoComplement16(8, uint16(item.Fixed.Data[0]))
		case 39:
			var payload []byte
			copy(payload[:], item.Repetitive.Payload()[:])
			tmp := modeSMBDataCAT021(payload)
			data.ModeSMBData = tmp
		case 40:
			var payload [7]byte
			copy(payload[:], item.Fixed.Data[:])
			tmp := aCASResolutionAdvisoryReport(payload)
			data.ACASResolutionAdvisoryReport = tmp
		case 41:
			data.ReceiverID = uint8(item.Fixed.Data[0])
		case 42:
			// Do stuff
		}
	}
}

// TODO: Refactor to cover for arbitrary number of extensions (currently only covers
//       two as that's explicitly in the spec)
func targetReportDescriptor(cp goasterix.Compound) TargetReportDescriptor {
	trd := new(TargetReportDescriptor)

	tmpList := cp.Payload()
	tmp := tmpList[0]

	switch tmp & 0xE0 >> 5 {
	case 0:
		trd.ATP = "24-Bit ICAO address"
	case 1:
		trd.ATP = "Duplicate Address"
	case 2:
		trd.ATP = "Surface vehicle address"
	case 3:
		trd.ATP = "Anonymous address"
	default: // 4-7
		trd.ATP = "Reserved for future use"
	}

	switch tmp & 0x18 >> 3 {
	case 0:
		trd.ARC = "25ft"
	case 1:
		trd.ARC = "100ft"
	case 2:
		trd.ARC = "Unknown"
	case 3:
		trd.ARC = "Invalid"
	}

	if tmp&0x4 == 0 {
		trd.RC = "Default"
	} else {
		trd.RC = "Range Check passed, CPR Validation pending"
	}

	if tmp&0x2 == 0 {
		trd.RAB = "Report from target transponder"
	} else {
		trd.RAB = "Report from field monitor (fixed transponder)"
	}

	if isFieldExtention(tmp) {
		fx1 := new(FirstExtensionTRD)

		fstItem := 0
		tmp = tmpList[fstItem] //?

		if tmp&0x80 == 0 {
			fx1.DCR = "No differential correction"
		} else {
			fx1.DCR = "Differential correction"
		}

		if tmp&0x40 == 0 {
			fx1.GBS = "Not set"
		} else {
			fx1.GBS = "Set"
		}

		if tmp&0x20 == 0 {
			fx1.SIM = "Actual"
		} else {
			fx1.SIM = "Simulated"
		}

		if tmp&0x10 == 0 {
			fx1.TST = "Default"
		} else {
			fx1.TST = "Test target"
		}

		if tmp&0x8 == 0 {
			fx1.SAA = "Capable"
		} else {
			fx1.SAA = "Not capable"
		}

		switch tmp & 0x6 >> 1 {
		case 0:
			fx1.CL = "Report valid"
		case 1:
			fx1.CL = "Report suspect"
		case 2:
			fx1.CL = "No info"
		case 3:
			fx1.CL = "Reserved for future use"
		}

		if isFieldExtention(tmp) {
			fx2 := new(SecondExtensionTRD)

			sndItem := 0
			tmp = tmpList[sndItem] //?

			if tmp&0x40 == 0 {
				fx2.LLC = "default"
			} else {
				fx2.LLC = "Target is suspect"
			}

			if tmp&0x20 == 0 {
				fx2.IPC = "default"
			} else {
				fx2.IPC = " Independent Position Check failed "
			}

			if tmp&0x10 == 0 {
				fx2.NOGO = "Not set"
			} else {
				fx2.NOGO = "Set"
			}

			if tmp&0x8 == 0 {
				fx2.CPR = "CPR validation correct"
			} else {
				fx2.CPR = "CPR vallidation failed"
			}

			if tmp&0x4 == 0 {
				fx2.LDPJ = "Not detected"
			} else {
				fx2.LDPJ = "Detected"
			}

			if tmp&0x2 == 0 {
				fx2.RCF = "Default"
			} else {
				fx2.RCF = "Range check failed"
			}

			// TODO: Investigate and implement sequential field extensions
			fx2.ThirdExtension = tmp & 0x1
		}

		trd.FX = fx1

	}

	return *trd
}

func qualityIndicators(cp goasterix.Compound) QualityIndicators {
	qi := new(QualityIndicators)

	tmp := cp.Primary[0]

	qi.NUCrOrNACv = int(tmp & 0xE0 >> 5)

	qi.NUCpOrNIC = int(tmp & 0x1E >> 1)

	if isFieldExtention(tmp) {
		fx1 := new(FirstExtensionQI)

		fstItem := 0
		fstByte := 0
		tmp = cp.Secondary[fstItem].Payload()[fstByte]

		fx1.NICBaro = int(tmp & 0x80 >> (BYTESIZE - 1))

		fx1.SIL = int(tmp & 0x60 >> (BYTESIZE - 3))

		fx1.NACp = int(tmp & 0x1E >> 1)

		if isFieldExtention(tmp) {
			fx2 := new(SecondExtensionQI)
			sndItem := 1
			tmp = cp.Secondary[sndItem].Payload()[fstByte]

			if tmp&0x20 == 0 {
				fx2.SILS = "flight-hour"
			} else {
				fx2.SILS = "sample"
			}

			fx2.SDA = int(tmp & 0x18 >> 3)

			fx2.GVA = int(tmp & 0x06 >> 1)

			if isFieldExtention(tmp) {
				fx3 := new(ThirdExtensionQI)
				thirdItem := 2
				tmp = cp.Secondary[thirdItem].Payload()[fstByte]

				fx3.PIC = getPIC(int(tmp & 0xF0 >> 4))

				// TODO: Confirm if further extensions are needed and how they'll be formatted
				//		 (not clear in spec)
				fx3.FX = tmp & 0x01

				fx2.FX = fx3

			}

			fx1.FX = fx2
		}

		qi.FX = fx1
	}

	return *qi
}

func wgs84Coordinates(data []byte) WGS84Coordinates {
	var pos WGS84Coordinates

	if len(data) == 6 {
		tmpLatitude := uint32(data[0])<<(2*BYTESIZE) + uint32(data[1])<<BYTESIZE + uint32(data[2])
		pos.Latitude = float32(goasterix.TwoComplement32(24, tmpLatitude)) * 0.00002145767

		tmpLongitude := uint32(data[3])<<(2*BYTESIZE) + uint32(data[4])<<BYTESIZE + uint32(data[5])
		pos.Longitude = float32(goasterix.TwoComplement32(32, tmpLongitude)) * 0.00002145767
	} else { // high precision data
		tmpLatitude := uint32(data[0])<<23 + uint32(data[1])<<15 + uint32(data[2])<<7 + uint32(data[3])
		pos.Latitude = float32(goasterix.TwoComplement32(32, tmpLatitude)) * 0.00000016764

		tmpLongitude := uint32(data[4])<<23 + uint32(data[5])<<15 + uint32(data[6])<<7 + uint32(data[7])
		pos.Longitude = float32(goasterix.TwoComplement32(32, tmpLongitude)) * 0.00000016764
	}

	return pos
}

func airSpeed(data [2]byte) AirSpeed {
	var speed AirSpeed

	tmp := data[0]
	speedValue := float64(uint32(data[0]&0x7F)<<8 + uint32(data[1]&0xFF))
	if tmp&0x80 == 0 {
		speed.IM = "IAS"
		speed.AirSpeed = speedValue * math.Pow(2, -14)

	} else {
		speed.IM = "Mach"
		speed.AirSpeed = speedValue * 0.001
	}

	return speed
}

func trueAirSpeed(data [2]byte) TrueAirSpeed {
	return TrueAirSpeed{
		RE:    int(data[0] & 0x80),
		Speed: int16(uint32(data[0]&0x7F)<<BYTESIZE + uint32(data[1]&0xFF)),
	}
}

func geometricHeight(data [2]byte) GeometricHeight {
	tmpHeight := goasterix.TwoComplement16(16, uint16(data[0])<<BYTESIZE+uint16(data[1]))

	greaterThan := false
	int16Max := int16(32767)
	if tmpHeight == int16Max {
		greaterThan = true
	}

	LSB := 6.25
	return GeometricHeight{
		Height:      float64(tmpHeight) * LSB,
		GreaterThan: greaterThan,
	}
}

func mOPS(cp goasterix.Compound) MOPSVersion {
	mops := new(MOPSVersion)

	tmp := cp.Primary[0]

	if tmp&0x40>>6 == 0 {
		mops.VNS = "supported"
	} else {
		mops.VNS = "not supported"
	}

	switch tmp & 0x38 >> 3 {
	case 0:
		mops.VN = "ED102/DO-260"
	case 1:
		mops.VN = "DO-260A"
	case 2:
		mops.VN = "ED102A/DO-260B"
	case 3:
		mops.VN = "ED102B-DO-260C"
	}

	switch tmp & 0x07 {
	case 0:
		mops.LTT = "other"
	case 1:
		mops.LTT = "uat"
	case 2:
		mops.LTT = "1090 es"
	case 3:
		mops.LTT = "vdl 4"
	default:
		mops.LTT = "undefined"
	}

	return *mops
}

func mode3ACodeCAT021(data [2]byte) *Mode3ACodeInOctal {
	tmpMode3ACode := new(Mode3ACodeInOctal)
	tmpData := data

	tmpMode3ACode.A1 = int(tmpData[0] & 0x8)
	tmpMode3ACode.A2 = int(tmpData[0] & 0x4)
	tmpMode3ACode.A4 = int(tmpData[0] & 0x2)
	tmpMode3ACode.B1 = int(tmpData[0] & 0x1)
	tmpMode3ACode.B2 = int(tmpData[1] & 0x80)
	tmpMode3ACode.B4 = int(tmpData[1] & 0x40)
	tmpMode3ACode.C1 = int(tmpData[1] & 0x20)
	tmpMode3ACode.C2 = int(tmpData[1] & 0x10)
	tmpMode3ACode.C4 = int(tmpData[1] & 0x8)
	tmpMode3ACode.C1 = int(tmpData[1] & 0x4)
	tmpMode3ACode.C2 = int(tmpData[1] & 0x2)
	tmpMode3ACode.C4 = int(tmpData[1] & 0x1)

	return tmpMode3ACode

}

func targetStatus(data []byte) *TargetStatus {
	ts := new(TargetStatus)
	tmp := data[0]

	if tmp&0x80 == 0 {
		ts.ICF = "No intent change active"
	} else {
		ts.ICF = "Intent change flag raised"
	}

	if tmp&0x40 == 0 {
		ts.LNAV = "engaged"
	} else {
		ts.LNAV = "not engaged"
	}

	if tmp&0x20 == 0 {
		ts.PS = "No military emergency"
	} else {
		ts.PS = "Military emergency"
	}

	switch tmp & 0x1C >> 2 {
	case 0:
		ts.PS = "No emergency/not reported"
	case 1:
		ts.PS = "General emergency"
	case 2:
		ts.PS = "Lifeguard/medical emergency"
	case 3:
		ts.PS = "Minimum fuel"
	case 4:
		ts.PS = "No communications"
	case 5:
		ts.PS = "Unlawful interference"
	case 6:
		ts.PS = "\"Downed\" Aircraft"
	}

	switch tmp & 0x03 {
	case 0:
		ts.SS = "No condition reported"
	case 1:
		ts.SS = "Permanent Alert (Emergency condition)"
	case 2:
		ts.SS = "Temporary Alert (change in Mode 3/A Code other than emergency)"
	case 3:
		ts.SS = "SPI set"
	}

	return ts
}

func rollAngle(data [2]byte) float64 {
	lsbResolution := 0.01
	sum := uint32(data[0])<<BYTESIZE + uint32(data[1])
	tmpRoll := goasterix.TwoComplement32(16, sum)
	return float64(tmpRoll) * float64(lsbResolution)
}

func verticalRate(data [2]byte) *VerticalRate {
	baroRate := new(VerticalRate)

	if int16(data[0])&0xF0>>BYTESIZE-1 == 0 {
		baroRate.RE = "Value in defined range"
	} else {
		baroRate.RE = "Value exceeds defined range "
	}

	baroRate.VerticalRate = float32(int16(data[0])&0x7F<<BYTESIZE+int16(data[1])) * 6.25

	return baroRate
}

func airborneGroundVector(data [4]byte) *AirborneGroundVector {
	agv := new(AirborneGroundVector)

	if data[0]&0x80>>3 == 0 {
		agv.RE = "Value in defined range"
	} else {
		agv.RE = "Value exceeds defined range"
	}

	agv.GroundSpeed = float32(int16(data[0])&0x7F<<BYTESIZE+int16(data[1])&0xFF) * float32(math.Pow(2, -14))

	tmpLSBTrackAngle := float32(360 / math.Pow(2, 16))
	agv.TrackAngle = float32(int16(data[2])&0xFF<<BYTESIZE+int16(data[3])&0xFF) * tmpLSBTrackAngle

	return agv
}

func trackAngleRate(data [2]byte) float32 {
	return float32(int16(data[0])&0x03<<BYTESIZE+int16(data[1])&0xFF) * float32(1.0/32)
}

func emitterCategory(data [1]byte) string {
	emitterCategoryStr := ""

	switch int8(data[0]) {
	case 0:
		emitterCategoryStr = "No ADS-B Emitter Category Information"
	case 1:
		emitterCategoryStr = "light aircraft <= 15500 lbs"
	case 2:
		emitterCategoryStr = "15500 lbs < small aircraft <75000 lbs"
	case 3:
		emitterCategoryStr = "75000 lbs < medium a/c < 300000 lbs"
	case 4:
		emitterCategoryStr = "High Vortex Large"
	case 5:
		emitterCategoryStr = "300000 lbs <= heavy aircraft"
	case 6:
		emitterCategoryStr = "highly manoeuvrable (5g acceleration capability) and high speed (>400 knots cruise)"
	case 7:
		emitterCategoryStr = "reserved"
	case 8:
		emitterCategoryStr = "reserved"
	case 9:
		emitterCategoryStr = "reserved"
	case 10:
		emitterCategoryStr = "rotocraft"
	case 11:
		emitterCategoryStr = "glider / sailplane"
	case 12:
		emitterCategoryStr = "lighter-than-air"
	case 13:
		emitterCategoryStr = "unmanned aerial vehicle"
	case 14:
		emitterCategoryStr = "space / transatmospheric vehicle"
	case 15:
		emitterCategoryStr = "ultralight / handglider / paraglider"
	case 16:
		emitterCategoryStr = "parachutist / skydiver"
	case 17:
		emitterCategoryStr = "reserved"
	case 18:
		emitterCategoryStr = "reserved"
	case 19:
		emitterCategoryStr = "reserved"
	case 20:
		emitterCategoryStr = "surface emergency vehicle"
	case 21:
		emitterCategoryStr = "surface service vehicle"
	case 22:
		emitterCategoryStr = "fixed ground or tethered obstruction"
	case 23:
		emitterCategoryStr = "cluster obstacle"
	case 24:
		emitterCategoryStr = "line obstacle"
	}

	return emitterCategoryStr
}

func selectedAltitude(data [2]byte) *SelectedAltitude {
	tmp := new(SelectedAltitude)
	if data[0]&0x80 != 0 {
		tmp.SAS = "source_information_provided"
	} else {
		tmp.SAS = "no_source_information_provided"
	}
	source := data[0] & 0x60 >> 5
	switch source {
	case 0:
		tmp.Source = "unknown"
	case 1:
		tmp.Source = "aircraft_altitude"
	case 2:
		tmp.Source = "fcu_mcp_selected_altitude"
	case 3:
		tmp.Source = "fms_selected_altitude"
	}
	altitude := goasterix.TwoComplement16(13, uint16(data[0]&0x1f)<<8+uint16(data[1]))
	tmp.Altitude = float64(altitude) * 25
	return tmp
}

func finalSelectedAltitude(data [2]byte) *StateSelectedAltitude {
	tmp := new(StateSelectedAltitude)
	if data[0]&0x80 != 0 {
		tmp.MV = "manage_vertical_mode_active"
	} else {
		tmp.MV = "manage_vertical_mode_not_active"
	}
	if data[0]&0x40 != 0 {
		tmp.AH = "altitude_hold_active"
	} else {
		tmp.AH = "altitude_hold_not_active"
	}
	if data[0]&0x20 != 0 {
		tmp.AM = "approach_mode_active"
	} else {
		tmp.AM = "approach_mode_not_active"
	}

	altitude := goasterix.TwoComplement16(13, uint16(data[0]&0x1f)<<8+uint16(data[1]))
	tmp.Altitude = float64(altitude) * 25
	return tmp
}

/*
func trajectoryIntent(cp goasterix.Compound) *TrajectoryIntent {

	tmpTI := new(TrajectoryIntent)
	subfieldPositionInCompound := 0

	if cp.Primary[0]&0x80>>BYTESIZE == 1 {
		tmpData := cp.Secondary[subfieldPositionInCompound].Extended

		tmpTIS := new(TrajectoryIntentStatus)
		if tmpData.Primary[0]&0x80>>BYTESIZE == 0 {
			tmpTIS.NAV = "Trajectory Intent Data is available for this aircraft"
		} else {
			tmpTIS.NAV = "Trajectory Intent Data is not available for this aircraft"
		}
		if tmpData.Secondary[0]&0x40>>BYTESIZE == 0 {
			tmpTIS.NVB = "Trajectory Intent Data is valid"
		} else {
			tmpTIS.NVB = "Trajectory Intent Data is not valid"
		}
		tmpTIS.FX = tmpData&0x01 // TODO: Implement this

		subfieldPositionInCompound += 1
	}

	if cp.Primary[0]&0x40>>BYTESIZE == 1 {
		tmpData := cp.Secondary[subfieldPositionInCompound].Extended

		tmpTID := new(TrajectoryIntentData)
		// TODO: Fill out the extraction logic
	}

	return tmpTI

}
*/

func aircraftOperationalStatus(data [1]byte) *AircraftOperationStatus {
	tmp := data[0]
	tmpAOS := new(AircraftOperationStatus)

	if uint16(tmp&0x80)>>BYTESIZE == 0 {
		tmpAOS.RA = "TCAS II or ACAS RA not active"
	} else {
		tmpAOS.RA = "TCAS RA active"
	}

	switch uint16(tmp&0x60)>>BYTESIZE - 3 {
	case 0:
		tmpAOS.TC = "no capability for Trajectory Change Reports"
	case 1:
		tmpAOS.TC = "support for TC+0 reports only"
	case 2:
		tmpAOS.TC = "support for multiple TC reports"
	case 3:
		tmpAOS.TC = "reserved"
	}

	if uint16(tmp&0x10)>>BYTESIZE-4 == 0 {
		tmpAOS.TS = "no capability to support Target State Reports"
	} else {
		tmpAOS.TS = "capable of supporting target State Reports"
	}

	if uint16(tmp&0x08)>>BYTESIZE-3 == 0 {
		tmpAOS.ARV = "no capability to generate ARV-reports"
	} else {
		tmpAOS.ARV = "capable of generate ARV-reports"
	}

	if uint16(tmp&0x04)>>BYTESIZE-2 == 0 {
		tmpAOS.CDTIA = "CDTI not operational"
	} else {
		tmpAOS.CDTIA = "CDTI operational"
	}

	if uint16(tmp&0x02)>>BYTESIZE-1 == 0 {
		tmpAOS.NotTCAS = "TCAS operational"
	} else {
		tmpAOS.NotTCAS = "TCAS not operational"
	}

	if tmp&0x01 == 0 {
		tmpAOS.SA = "Antenna Diversity"
	} else {
		tmpAOS.SA = "Single Antenna only"
	}

	return tmpAOS
}

func surfaceCapabilitiesAndCharacteristics(data []byte) *SurfaceCapabilitiesAndCharacteristics {
	tmpSCAC := new(SurfaceCapabilitiesAndCharacteristics)

	if uint16(data[0]&0x20)>>5 == 0 {
		tmpSCAC.POA = "Position transmitted is not ADS-B position reference point"
	} else {
		tmpSCAC.POA = "Position transmitted is the ADS-B position reference point"
	}

	if uint16(data[0]&0x20)>>5 == 0 {
		tmpSCAC.CDTIS = "CDTI not operationa"
	} else {
		tmpSCAC.CDTIS = "CDTI operational"
	}

	if uint16(data[0]&0x20)>>5 == 0 {
		tmpSCAC.B2Low = ">= 70 Watts"
	} else {
		tmpSCAC.B2Low = "< 70 Watts"
	}

	if uint16(data[0]&0x20)>>5 == 0 {
		tmpSCAC.RAS = "Aircraft not receiving ATC-services"
	} else {
		tmpSCAC.RAS = "Aircraft receiving ATC services"
	}

	if uint16(data[0]&0x20)>>5 == 0 {
		tmpSCAC.IDENT = "IDENT switch not active"
	} else {
		tmpSCAC.IDENT = "IDENT switch active"
	}

	addAdditionalFieldExtensions := (uint16(data[0]&0x01) == 1)
	i := 1
	for addAdditionalFieldExtensions && i < len(data) {
		tmpFX := new(FirstExtensionSCC)

		tmpFX.LW = TableLWV1[uint32(data[i]&0xF0)>>BYTESIZE-4]

		addAdditionalFieldExtensions = (uint16(data[i]&0x01) == 1)
		i += 1
	}

	return tmpSCAC
}

func modeSMBDataCAT021(data []byte) *ModeSMBData {
	modesmbdata := new(ModeSMBData)

	modesmbdata.REP = int8(data[0])

	modesmbdata.MB = hex.EncodeToString(data[1:8])

	modesmbdata.BDS1 = int16(data[8]&0xF0)>>BYTESIZE - 4

	modesmbdata.BDS2 = int8(data[8] & 0x0F)

	return modesmbdata
}

func aCASResolutionAdvisoryReport(data [7]byte) *ACASResolutionAdvisoryReport {
	tmpAcas := new(ACASResolutionAdvisoryReport)
	tmpAcas.TYP = int8(data[0]&0xF8) >> 3
	tmpAcas.STYP = int8(data[0] & 0x07)
	tmpAcas.ARA = int16(data[1])<<BYTESIZE - 2 + int16(data[2]&0xFA)>>2
	tmpAcas.RAC = int16(data[2]&0x03)<<2 + int16(data[3]&0xC0)>>BYTESIZE - 2
	tmpAcas.RAT = int16(data[3]&0x02)>>BYTESIZE - 3
	tmpAcas.MTE = int16(data[3]&0x10)>>BYTESIZE - 4
	tmpAcas.TTI = int8(data[3]&0x0C) >> 2
	tmpAcas.TID = int32(data[3]&0x03)<<3*BYTESIZE + int32(data[4])<<2*BYTESIZE + int32(data[5])<<BYTESIZE + int32(data[6])

	return tmpAcas
}

func isFieldExtention(data byte) bool {
	return data&0x01 == 1
}

func getPIC(data int) *PIC {
	tmpPIC := new(PIC)
	tmpNICV2 := new(NIC_Version2OrHigher)
	switch data {
	case 0:
		tmpPIC.NUCp = 0
		tmpPIC.NIC_DO260A = "0"

		tmpNICV2.NIC = 0
	case 1:
		tmpPIC.IntegrityContainmentBound = 20.0
		tmpPIC.NUCp = 1
		tmpPIC.NIC_DO260A = "1"

		tmpNICV2.NIC = 1
	case 2:
		tmpPIC.IntegrityContainmentBound = 10.0
		tmpPIC.NUCp = 2
	case 3:
		tmpPIC.IntegrityContainmentBound = 8.0
		tmpPIC.NIC_DO260A = "2"

		tmpNICV2.NIC = 2
	case 4:
		tmpPIC.IntegrityContainmentBound = 4.0
		tmpPIC.NIC_DO260A = "3"

		tmpNICV2.NIC = 3
	case 5:
		tmpPIC.IntegrityContainmentBound = 2.0
		tmpPIC.NUCp = 3
		tmpPIC.NIC_DO260A = "4"

		tmpNICV2.NIC = 4
	case 6:
		tmpPIC.IntegrityContainmentBound = 1.0
		tmpPIC.NUCp = 4
		tmpPIC.NIC_DO260A = "5"

		tmpNICV2.NIC = 5
	case 7:
		tmpPIC.IntegrityContainmentBound = 0.6
		tmpPIC.NIC_DO260A = "6 (+ 1)"

		tmpNICV2.NIC = 6
		tmpNICV2.AB = "1/1"
		tmpNICV2.AC = "0/1"
	case 8:
		tmpPIC.IntegrityContainmentBound = 0.5
		tmpPIC.NIC_DO260A = "6 (+ 0)"

		tmpNICV2.NIC = 6
		tmpNICV2.AB = "0/0"
	case 9:
		tmpPIC.IntegrityContainmentBound = 0.3

		tmpNICV2.NIC = 6
		tmpNICV2.AB = "0/1"
		tmpNICV2.AC = "1/0"
	case 10:
		tmpPIC.IntegrityContainmentBound = 0.2
		tmpPIC.NUCp = 6
		tmpPIC.NIC_DO260A = "7"

		tmpNICV2.NIC = 7
	case 11:
		tmpPIC.IntegrityContainmentBound = 0.1
		tmpPIC.NUCp = 7
		tmpPIC.NIC_DO260A = "8"

		tmpNICV2.NIC = 8
	case 12:
		tmpPIC.IntegrityContainmentBound = 0.04
		tmpPIC.NIC_DO260A = "9"

		tmpNICV2.NIC = 9
	case 13:
		tmpPIC.IntegrityContainmentBound = 0.013
		tmpPIC.NUCp = 8
		tmpPIC.NIC_DO260A = "10"

		tmpNICV2.NIC = 10
	case 14:
		tmpPIC.IntegrityContainmentBound = 0.004
		tmpPIC.NUCp = 9
		tmpPIC.NIC_DO260A = "11"

		tmpNICV2.NIC = 11
	}
	tmpPIC.NIC_Version2OrHigher = tmpNICV2

	return tmpPIC
}
