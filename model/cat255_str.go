package model

import (
	"errors"
	"goasterix/uap"
)
var (
	ErrCartOrdUnknown = errors.New("[ASTERIX Error] CART ORD Unknown")
)

type BiaisRadar struct {
	SacSic       SourceIdentifier `json:"SourceIdentifier"`
	GainDistance float64          `json:"gainDistance"`
	BiaisDistance float64          `json:"biaisDistance"`
	BiaisAzimut   float64          `json:"biaisAzimut"`
	BiaisDatation float64          `json:"biaisDatation"`
}
type CarteActive struct {
	Nom string `json:"nom"`
	Ord string `json:"ord"`
}
type NivC struct {
	NivInf int16 `json:"nivinf"`
	NivSup int16 `json:"nivsup"`
}

type PresenceSTPV struct {
	Version uint8  `json:"version"`
	Nap     uint8  `json:"nap"`
	NS      string `json:"ns"`
	ST      string `json:"st,omitempty"`
	PS      string `json:"ps,omitempty"`
}

type Cat255STRModel struct {
	SacSic *SourceIdentifier `json:"SourceIdentifier,omitempty"`
	Hem    float64           `json:"hem,omitempty"`
	Spe    *PresenceSTPV     `json:"spe,omitempty"`
	Nivc   *NivC             `json:"nivc,omitempty"`
	Txtc   string            `json:"txtc,omitempty"`
	Cart   *CarteActive      `json:"cart,omitempty"`
	Biais  []BiaisRadar      `json:"biais,omitempty"`
}

func (data *Cat255STRModel) write(items []uap.DataField) {
	for _, item := range items {
		switch item.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		case 2:
			// HEM : Heure d’émission du message d’alerte
			var payload [3]byte
			copy(payload[:], item.Payload[:])
			data.Hem, _ = timeOfDay(payload)
		case 3:
			// SPE : Présence STR-STPV
			tmp, _ := speStpv(item.Payload)
			data.Spe = &tmp
		case 4:
			// NIVC : Niveaux optionnels assignés à la carte dynamique
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := nivCarte(payload)
			data.Nivc = &tmp
		case 5:
			// TXTC : Texte optionnel de la carte dynamique
			data.Txtc = string(item.Payload[1:])
		case 6:
			// CART : activation de cartes dynamiques
			var payload [9]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := carte(payload)
			data.Cart = &tmp
		case 7:
			// BIAIS : Valeurs des biais courants radars
			data.Biais, _ = biaisExtract(item.Payload)
		}
	}
}

func speStpv(data []byte) (spe PresenceSTPV, err error) {
	spe.Version = data[0] & 0xE0 >> 5
	spe.Nap = data[0] & 0x18 >> 3

	tmpNs := data[0] & 0x06
	switch tmpNs {
	case 0:
		spe.NS = "principal"
	case 1:
		spe.NS = "secours"
	case 2:
		spe.NS = "test"
	}
	if data[0] & 0x01 != 0 {
		if data[1] & 0x80 != 0 {
			spe.ST = "evaluation"
		} else {
			spe.ST = "operational"
		}
		if data[1] & 0x40 != 0 {
			spe.PS = "stpv_deconnecte_str"
		} else {
			spe.PS = "stpv_connecte_str"
		}
	}

	return spe, nil
}

func nivCarte(data [4]byte) (nivc NivC, err error) {
	nivc.NivInf = int16(data[0])<<8 + int16(data[1])
	nivc.NivSup = int16(data[2])<<8 + int16(data[3])
	return nivc, nil
}

func carte(data [9]byte) (cart CarteActive, err error) {
	cart.Nom = string(data[:8])
	tmpOrd := data[8] & 0xE0 >> 5
	switch tmpOrd {
	case 0:
		cart.Ord = "activation_carte"
	case 1:
		cart.Ord = "annulation_carte"
	default:
		cart.Ord = "unknowm"
		err = ErrCartOrdUnknown
	}
	return cart, nil
}

func biaisExtract(data []byte) (biais []BiaisRadar, err error) {
	n := int(data[0])
	for i:=0; i < n; i++ {
		b := BiaisRadar{}
		var sacsic [2]byte
		copy(sacsic[:], data[i+1:i+3])
		b.SacSic,_ = sacSic(sacsic)
		b.GainDistance = float64(uint16(data[i+3])<<8 + uint16(data[i+4]))/6384
		b.BiaisDistance = float64(int16(data[i+5])<<8 + int16(data[i+6]))
		b.BiaisAzimut = float64(int16(data[i+7])<<8 + int16(data[i+8]))*0.0055
		b.BiaisDatation = float64(int16(data[i+8])<<8 + int16(data[i+9]))/1024
		biais = append(biais, b)
	}
	return biais, nil
}