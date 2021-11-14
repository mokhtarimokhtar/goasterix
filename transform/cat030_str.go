package transform

import (
	"encoding/hex"
	"errors"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"strconv"
	"strings"
)

var (
	ErrPistTypeUnknown   = errors.New("[ASTERIX Error] Piste TYPE Unknown")
	ErrPistSlrUnknown    = errors.New("[ASTERIX Error] Piste SLR Unknown")
	ErrPistCorUnknown    = errors.New("[ASTERIX Error] Piste COR Unknown")
	ErrPistDs1ds2Unknown = errors.New("[ASTERIX Error] Piste Ds1Ds2 Unknown")
)

type Mov struct {
	Trans string `json:"trans,omitempty"`
	Longi string `json:"longi,omitempty"`
	Verti string `json:"verti,omitempty"`
}

type Pist struct {
	Liv    string `json:"liv,omitempty"`
	Cnf    string `json:"cnf,omitempty"`
	Man    string `json:"man,omitempty"`
	Tva    string `json:"tva,omitempty"`
	Type   string `json:"type,omitempty"`
	Mort   string `json:"mort,omitempty"`
	Cre    string `json:"cre,omitempty"`
	Slr    string `json:"slr,omitempty"`
	Cor    string `json:"cor,omitempty"`
	Ds1ds2 string `json:"ds1ds2,omitempty"`
	For    string `json:"for,omitempty"`
	Ama    string `json:"ama,omitempty"`
	Spi    string `json:"spi,omitempty"`
	Me     string `json:"me,omitempty"`
}

type Spe struct {
	SY  uint8 `json:"sy"`
	M   uint8 `json:"m"`
	S   uint8 `json:"s"`
	O1  uint8 `json:"o1"`
	O2  uint8 `json:"o2"`
	O3  uint8 `json:"o3"`
	O4  uint8 `json:"o4"`
	O5  uint8 `json:"o5"`
	O6  uint8 `json:"o6"`
	O7  uint8 `json:"o7"`
	O8  uint8 `json:"o8"`
	O9  uint8 `json:"o9"`
	O10 uint8 `json:"o10"`
	O11 uint8 `json:"o11"`
	O12 uint8 `json:"o12"`
	O13 uint8 `json:"o13"`
	O14 uint8 `json:"o14"`
	O15 uint8 `json:"o15"`
	O16 uint8 `json:"o16"`
	O17 uint8 `json:"o17"`
	O18 uint8 `json:"o18"`
	O19 uint8 `json:"o19"`
	R   uint8 `json:"r"`
	C   uint8 `json:"c"`
}

type Flstr struct {
	Vc        string  `json:"vc"`
	Gc        string  `json:"gc"`
	NiveauVol float64 `json:"niveauVol"`
}

type Vit struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ModeA struct {
	V    string `json:"v"`
	G    string `json:"g"`
	C    string `json:"c"`
	Code uint16 `json:"code"`
}

type Altic struct {
	QNC int16 `json:"qnc,omitempty"`
	Alt int16 `json:"alt,omitempty"`
}

type NumPiste struct {
	Version uint8  `json:"version"`
	Nap     uint8  `json:"nap"`
	ST      string `json:"st"`
	NS      string `json:"ns"`
	Numero  uint16 `json:"numero"`
}

type Cat030STRModel struct {
	SacSic    *SourceIdentifier    `json:"sourceIdentifier,omitempty"`
	Num       *NumPiste            `json:"num,omitempty"`
	Hptu      float64              `json:"hptu,omitempty"`
	Pist      *Pist                `json:"pist,omitempty"`
	Alis      *ModeA               `json:"alis,omitempty"`
	Pos       *CartesianXYPosition `json:"pos,omitempty"`
	Qual      uint8                `json:"qual,omitempty"`
	Flpc      *Flstr               `json:"flpc,omitempty"`
	Flpm      *Flstr               `json:"flpm,omitempty"`
	Vit       *Vit                 `json:"vit,omitempty"`
	Mov       *Mov                 `json:"mov,omitempty"`
	Taux      float64              `json:"taux,omitempty"`
	Spe       *Spe                 `json:"spe,omitempty"`
	RadSacSic *SourceIdentifier    `json:"radSacSic,omitempty"`
	Ivol      string               `json:"ivol,omitempty"`
	Pln       uint16               `json:"pln,omitempty"`
	Av        string               `json:"av,omitempty"`
	Turb      string               `json:"turb,omitempty"`
	Terd      string               `json:"terd,omitempty"`
	Tera      string               `json:"tera,omitempty"`
	Altic     *Altic               `json:"altic,omitempty"`
	Adrs      string               `json:"adrs,omitempty"`
	Ids       string               `json:"ids,omitempty"`
}

// Write writes a single ASTERIX Record to Cat030STRModel.
// Items is a slice of Items DataField.
func (data *Cat030STRModel) write(items []uap.DataField) {
	for _, item := range items {
		switch item.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		// case 2 N∕A
		case 3:
			//Numéro de Piste STR
			var payload [3]byte
			copy(payload[:], item.Payload[:])
			tmp := num(payload)
			data.Num = &tmp
		case 4:
			// HPTU returns a float64 in second (1 bit = 1/128 s)
			// Absolute time stamping of the information provided in the track message, in the form
			// of elapsed time since last midnight.
			// The time of day value is reset to 0 each day at midnight.
			// Ref: 7.3.4 HPTU : Heure TU de la piste
			var payload [3]byte
			copy(payload[:], item.Payload[:])
			data.Hptu, _ = timeOfDay(payload)
		case 5:
			// Etat piste
			tmp, _ := pist(item.Payload)
			data.Pist = &tmp
		case 6:
			// alis : Mode A lissé piste
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := alis(payload)
			data.Alis = &tmp
		case 7:
			// Position cartésienne calculée
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp := pos(payload)
			data.Pos = &tmp
		case 8:
			// QUAL returns an integer of track quality range = 0 to 7(best).
			// Ref: 7.3.8 QUAL : Qualité piste
			data.Qual = item.Payload[0] & 0xFE >> 1
		case 9:
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := flp(payload)
			data.Flpc = &tmp
		case 10:
			// Niveau de vol mesuré de la piste
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := flp(payload)
			data.Flpm = &tmp
		case 11:
			// VIT : Vitesse calculée dans le plan (coordonnées cartésiennes)
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp := vitCal(payload)
			data.Vit = &tmp
		case 12:
			// mov : Mode de vol, tendance verticale
			var payload [1]byte
			copy(payload[:], item.Payload[:])
			tmp := mov(payload)
			data.Mov = &tmp
		case 13:
			// Taux returns Rate of ascent / descent in float64 FL/min
			// TAUX : Taux de montée/descente
			data.Taux = float64(int16(item.Payload[0])<<8+int16(item.Payload[1])) * 5.859375
		case 14:
			// spe : Marquage spécial (Special purpose)
			tmp, _ := spe(item.Payload)
			data.Spe = &tmp
		case 15:
			// RAD : Numéro de radar
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := sacSic(payload)
			data.RadSacSic = &tmp
		case 16:
			// IVOL : Indicatif de vol complet
			data.Ivol = string(item.Payload)
		case 17:
			// PLN : Numéro de plan de vol CAUTRA (number flight plan)
			data.Pln = uint16(item.Payload[0])<<8 + uint16(item.Payload[1])
		case 18:
			// AV : Type d’avion (type aircraft)
			data.Av = string(item.Payload)
		case 19:
			data.Turb = string(item.Payload[:])
		case 20:
			// Terd Terrain de départ (departure)
			data.Terd = string(item.Payload)
		case 21:
			// Tera Terrain d’arrivée (arrival)
			data.Tera = string(item.Payload)
		//case 22:
		// obsolete for this version
		// altic : Altitude calculée de la piste
		//var payload [2]byte
		//copy(payload[:], item.Payload[:])
		//tmp := altic(payload)
		//data.Altic = &tmp
		case 23:
			// ADRS : Adresse mode S
			data.Adrs = strings.ToUpper(hex.EncodeToString(item.Payload[:]))
		case 24:
			// IDS : Identification mode S
			var payload [6]byte
			copy(payload[:], item.Payload[:])
			data.Ids, _ = modeSIdentification(payload)
		}
	}
}

// vitCal returns a slice [X,Y] of float64 NM/s.
// Calculated track Velocity expressed in Cartesian coordinates.
// Ref: 7.3.11 VIT : Vitesse calculée dans le plan (coordonnées cartésiennes)
func vitCal(data [4]byte) Vit {
	var vit Vit
	vit.X = float64(int16(data[0])<<8+int16(data[1])) * 0.000061035
	vit.Y = float64(int16(data[2])<<8+int16(data[3])) * 0.000061035
	return vit
}

// flp returns a integer (1 bit = 1/4 FL), range = -15 to 1500 FL
// Flight Level into binary representation converted in a integer (16bits).
// Ref: 7.3.10 FLPM : Niveau de vol mesuré de la piste
func flp(data [2]byte) Flstr {
	var flpm Flstr
	if data[0]&0x80 != 0 {
		flpm.Vc = "code_not_validated"
	} else {
		flpm.Vc = "code_validated"
	}

	if data[0]&0x40 != 0 {
		flpm.Gc = "garbled_code"
	} else {
		flpm.Gc = "default"
	}

	tmp := uint16(data[0])<<8 + uint16(data[1])
	tmp = tmp & 0x3FFF
	niveauVol := goasterix.TwoComplement16(13, tmp)
	flpm.NiveauVol = float64(niveauVol) / 4 // divide by 4 is in 100's feet

	return flpm
}

// pos returns a slice [X,Y] of float64 NM (1 bit = 1/64 NM),
// range = - 512 NM .. 511.984 NM.
// Calculated position of an aircraft expressed in Cartesian coordinates.
// Ref: 7.3.7 POS : Position cartésienne calculée
func pos(data [4]byte) CartesianXYPosition {
	var pos CartesianXYPosition
	pos.X = float64(int16(data[0])<<8+int16(data[1])) / 64
	pos.Y = float64(int16(data[2])<<8+int16(data[3])) / 64
	return pos
}

// num returns a map
// Numéro de Piste STR:
// N° de version: Numéro de la version logicielle en service
// NAP: Numéro du calculateur
// ST: Statut du serveur
// N/S: Mode du serveur.
func num(data [3]byte) NumPiste {
	var num NumPiste
	num.Version = data[0] & 0xE0 >> 5
	num.Nap = data[0] & 0x18 >> 3

	if data[0]&0x04 != 0 {
		num.ST = "evaluation"
	} else {
		num.ST = "operational"
	}

	tmpNs := data[0] & 0x03
	switch tmpNs {
	case 0:
		num.NS = "principal"
	case 1:
		num.NS = "secours"
	case 2:
		num.NS = "test"
	}

	tmpPiste := uint16(data[1])<<8 + uint16(data[2])
	tmpPiste = tmpPiste & 0x1FFE >> 1
	num.Numero = tmpPiste

	return num
}

// pist return a map
// Etat piste
func pist(data []byte) (piste Pist, err error) {

	if data[0]&0x80 != 0 {
		piste.Liv = "simule_ou_plot_test"
	} else {
		piste.Liv = "trafic_reel"
	}
	if data[0]&0x40 != 0 {
		piste.Cnf = "piste_initialisation"
	} else {
		piste.Cnf = "piste_confirmee"
	}
	if data[0]&0x20 != 0 {
		piste.Man = "piste_virage"
	} else {
		piste.Man = "defaut"
	}
	if data[0]&0x10 != 0 {
		piste.Tva = "piste_pas_niveau_vol_valide"
	} else {
		piste.Tva = "defaut"
	}

	typePiste := data[0] & 0x0E >> 1
	switch typePiste {
	case 0:
		piste.Type = "piste_association_multiple_primaire_secondaire"
	case 1:
		piste.Type = "piste_association_primaire_pure"
	case 2:
		piste.Type = "piste_association_multiple_secondaire_pure"
	case 3:
		piste.Type = "piste_monoradar_p_plus_s"
	case 4:
		piste.Type = "piste_monoradar_secondaire_pure"
	case 5:
		piste.Type = "piste_monoradar_primaire_pure"
	case 6:
		piste.Type = "undefined"
	case 7:
		piste.Type = "piste_en_manque"
	default:
		piste.Type = "unknowm"
		err = ErrPistTypeUnknown
	}

	if data[0]&0x01 != 0 {
		if data[1]&0x80 != 0 {
			piste.Mort = "mort_de_piste"
		} else {
			piste.Mort = "defaut"
		}
		if data[1]&0x40 != 0 {
			piste.Cre = "creation_de_piste"
		} else {
			piste.Cre = "defaut"
		}

		slr := data[1] & 0x30 >> 4
		switch slr {
		case 0:
			piste.Slr = "coordonnees_projetees_niveau_calcule"
		case 1:
			piste.Slr = "coordonnees_projetees_niveau_mesure"
		case 2:
			piste.Slr = "coordonnees_projetees_niveau_forfaitaire"
		case 3:
			piste.Slr = "coordonnees_rabattues"
		default:
			piste.Slr = "unknowm"
			err = ErrPistSlrUnknown
		}

		cor := data[1] & 0x0E >> 1
		switch cor {
		case 0:
			piste.Cor = "piste_correlation_plan_vol_confirmee"
		case 1:
			piste.Cor = "piste_correlation_plan_vol_associee"
		case 2:
			piste.Cor = "piste_correlation_plan_vol_gelee"
		case 3:
			piste.Cor = "piste_post_correlation"
		case 4:
			piste.Cor = "undefined"
		case 5:
			piste.Cor = "undefined"
		case 6:
			piste.Cor = "undefined"
		case 7:
			piste.Cor = "piste_non_correlee_plan_vol"
		default:
			piste.Cor = "unknowm"
			err = ErrPistCorUnknown
		}

		if data[1]&0x01 != 0 {
			ds1ds2 := data[2] & 0xC0 >> 6
			switch ds1ds2 {
			case 0:
				piste.Ds1ds2 = "defaut"
			case 1:
				piste.Ds1ds2 = "detournement_code_75000"
			case 2:
				piste.Ds1ds2 = "panne_radio_code_7600"
			case 3:
				piste.Ds1ds2 = "detresse_code_7700"
			default:
				piste.Ds1ds2 = "unknowm"
				err = ErrPistDs1ds2Unknown
			}

			if data[2]&0x20 != 0 {
				piste.For = "vol_en_formation"
			} else {
				piste.For = "defaut"
			}
			if data[2]&0x10 != 0 {
				piste.Ama = "piste_non_amalgamee"
			} else {
				piste.Ama = "piste_amalgamee"
			}
			if data[2]&0x08 != 0 {
				piste.Spi = "special_pulse_ident"
			} else {
				piste.Spi = "defaut"
			}
			if data[2]&0x04 != 0 {
				piste.Me = "detresse_militaire"
			} else {
				piste.Me = "defaut"
			}
		}
	}

	return piste, err
}

// alis : Mode A lissé piste
// Mode A : Ce champ correspond au code mode A lissé de la piste.
func alis(data [2]byte) ModeA {
	var alis ModeA
	if data[0]&0x80 != 0 {
		alis.V = "code_invalide"
	} else {
		alis.V = "code_valide"
	}
	if data[0]&0x40 != 0 {
		alis.G = "code_garbling"
	} else {
		alis.G = "defaut"
	}
	if data[0]&0x20 != 0 {
		alis.C = "code_changement"
	} else {
		alis.C = "code_pas_changement"
	}

	tmp := uint16(data[0])&0x000F<<8 + uint16(data[1])&0x00FF
	tmpModeA := strconv.FormatUint(uint64(tmp), 8)
	modeA, _ := strconv.ParseInt(tmpModeA, 10, 16)
	alis.Code = uint16(modeA)

	return alis
}

// mov : Mode de vol, tendance verticale
func mov(data [1]byte) Mov {
	var mov Mov
	tmpTrans := data[0] & 0xC0 >> 6
	switch tmpTrans {
	case 0:
		mov.Trans = "ligne_droite"
	case 1:
		mov.Trans = "virage_droite"
	case 2:
		mov.Trans = "virage_gauche"
	case 3:
		mov.Trans = "tendance_indeterminee"
	}

	tmpLongi := data[0] & 0x30 >> 4
	switch tmpLongi {
	case 0:
		mov.Longi = "vitesse_sol_constante"
	case 1:
		mov.Longi = "vitesse_sol_augmentation"
	case 2:
		mov.Longi = "vitesse_sol_diminution"
	case 3:
		mov.Longi = "tendance_indeterminee"
	}

	tmpVerti := data[0] & 0x0C >> 2
	switch tmpVerti {
	case 0:
		mov.Verti = "vol_palier"
	case 1:
		mov.Verti = "vol_montee"
	case 2:
		mov.Verti = "vol_descente"
	case 3:
		mov.Verti = "tendance_indeterminee"
	}

	return mov
}

// spe : Marquage spécial (Special purpose)
func spe(data []byte) (spe Spe, err error) {
	spe.SY = data[0] & 0xF8 >> 3
	spe.M = data[0] & 0x04 >> 2
	spe.S = data[0] & 0x02 >> 1

	if data[0]&0x01 != 0 {
		spe.O19 = data[1] & 0x80 >> 7
		spe.O18 = data[1] & 0x40 >> 6
		spe.O17 = data[1] & 0x20 >> 5
		spe.O16 = data[1] & 0x10 >> 4
		spe.O15 = data[1] & 0x08 >> 3
		spe.O14 = data[1] & 0x04 >> 2
		spe.O13 = data[1] & 0x02 >> 1

		if data[1]&0x01 != 0 {
			spe.O12 = data[2] & 0x80 >> 7
			spe.O11 = data[2] & 0x40 >> 6
			spe.O10 = data[2] & 0x20 >> 5
			spe.O9 = data[2] & 0x10 >> 4
			spe.O8 = data[2] & 0x08 >> 3
			spe.O7 = data[2] & 0x04 >> 2
			spe.O6 = data[2] & 0x02 >> 1

			if data[2]&0x01 != 0 {
				spe.O5 = data[3] & 0x80 >> 7
				spe.O4 = data[3] & 0x40 >> 6
				spe.O3 = data[3] & 0x20 >> 5
				spe.O2 = data[3] & 0x10 >> 4
				spe.O1 = data[3] & 0x08 >> 3
				spe.R = data[3] & 0x04 >> 2
				spe.C = data[3] & 0x02 >> 1
			}
		}
	}
	return spe, nil
}

// altic : Altitude calculée de la piste
// jamais transmis (obsolete).
func altic(data [2]byte) Altic {
	var altic Altic
	altic.QNC = int16(data[0] & 0x80 >> 7)
	altic.Alt = int16(data[0])&0x007F<<8 + int16(data[1])&0x00FF
	return altic
}
