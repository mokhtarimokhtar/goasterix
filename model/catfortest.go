package model

import (
	"fmt"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/item"
)

type CatForTestModel struct {
	SacSic *SourceIdentifier `json:"sourceIdentifier,omitempty"`
}

// sacSic returns a SourceIdentifier with:
// Sac: an integer of System Area TransponderRegisterNumber.
// Sic: an integer of System Identification TransponderRegisterNumber.
func sacSic(data [2]byte) (src SourceIdentifier, err error) {
	src.Sac = data[0]
	src.Sic = data[1]
	return src, nil
}

func (data *CatForTestModel) write(rec goasterix.IRecord) {
	dataItems := rec.GetItems()
	for _, dataItem := range dataItems {
		switch dataItem.GetFrn() {
		case 1:
			fmt.Println("dataItem", dataItem.(*item.Fixed).String())
			// decode sac sic
			//var payload [2]byte
			//copy(payload[:], item.Fixed.Data)
			tmp := new(SourceIdentifier)
			//tmp.Sac = dataItem.(*item.Fixed).SubItems[0].Data[0]
			//tmp.Sic = dataItem.(*item.Fixed).SubItems[1].Data[0]

			sub := dataItem.GetSubItems()
			tmp.Sac = sub[0].Data[0]
			tmp.Sic = sub[1].Data[0]

			data.SacSic = tmp
		}
	}

	/*for _, item := range rec.Items {
		switch item.Meta.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Fixed.Data)
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		}
	}*/
}

// trackStatus returns a map of uint8, CNF, RAD, DOU, MAH, CDM id exist: TRE, GHO, SUP, TCC.
// Status of monoradar track (PSR and/or SSR updated).
/*
func trackStatus(item goasterix.Extended) Status {
	var ts Status

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
	return ts
}
*/
