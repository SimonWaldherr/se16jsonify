package main

import (
	s "./structs"
	"reflect"
)

func LoadKNAStruct(searchtype, searchvalue string) []*s.Kna1 {
	connect()
	data, _ := ReadTable("KNA1", searchtype, searchvalue, reflect.ValueOf(s.Kna1{}), reflect.TypeOf(s.Kna1{}))
	close()

	var knadata []*s.Kna1

	for key := range data {
		row := &s.Kna1{
			KUNNR: SAPStringClean(data[key][0]),
			LAND1: SAPStringClean(data[key][1]),
			NAME1: SAPStringClean(data[key][2]),
			NAME2: SAPStringClean(data[key][3]),
			NAME3: SAPStringClean(data[key][4]),
			ORT01: SAPStringClean(data[key][5]),
			PSTLZ: SAPStringClean(data[key][6]),
			STRAS: SAPStringClean(data[key][7]),
			TELF1: SAPStringClean(data[key][8]),
			TELF2: SAPStringClean(data[key][9]),
		}
		knadata = append(knadata, row)
	}
	return knadata
}

func LoadMARAStruct(searchtype, searchvalue string) []*s.Mara {
	connect()
	data, _ := ReadTable("MARAV", searchtype, searchvalue, reflect.ValueOf(s.Mara{}), reflect.TypeOf(s.Mara{}))
	close()

	var maradata []*s.Mara

	for key := range data {
		row := &s.Mara{
			MATNR: SAPStringClean(data[key][0]),
			MATKL: SAPStringClean(data[key][1]),
			MEINS: SAPStringClean(data[key][2]),
			BRGEW: SAPStringClean(data[key][3]),
			GEWEI: SAPStringClean(data[key][4]),
			EAN11: SAPStringClean(data[key][5]),
			MFRPN: SAPStringClean(data[key][6]),
			PRDHA: SAPStringClean(data[key][7]),
			MAKTX: SAPStringClean(data[key][8]),
			MAKTG: SAPStringClean(data[key][9]),
			MFRNR: SAPStringClean(data[key][10]),
		}
		maradata = append(maradata, row)
	}
	return maradata
}

func LoadLIPSStruct(searchtype, searchvalue string) []*s.Lips {
	connect()
	data, _ := ReadTable("LIPS", searchtype, searchvalue, reflect.ValueOf(s.Lips{}), reflect.TypeOf(s.Lips{}))
	close()

	var lipsdata []*s.Lips

	for key := range data {
		row := &s.Lips{
			VBELN: SAPStringClean(data[key][0]),
			POSNR: SAPStringClean(data[key][1]),
			MATNR: SAPStringClean(data[key][2]),
			MATKL: SAPStringClean(data[key][3]),
			ARKTX: SAPStringClean(data[key][4]),
			EANNR: SAPStringClean(data[key][5]),
			LGORT: SAPStringClean(data[key][6]),
			LFIMG: SAPStringClean(data[key][7]),
			VRKME: SAPStringClean(data[key][8]),
			VKBUR: SAPStringClean(data[key][9]),
		}
		lipsdata = append(lipsdata, row)
	}
	return lipsdata
}

func LoadLIKPStruct(searchtype, searchvalue string) []*s.Likp {
	connect()
	data, _ := ReadTable("LIKP", searchtype, searchvalue, reflect.ValueOf(s.Likp{}), reflect.TypeOf(s.Likp{}))
	close()

	var likpdata []*s.Likp

	for key := range data {
		row := &s.Likp{
			VBELN:         SAPStringClean(data[key][0]),
			VBTYP:         SAPStringClean(data[key][1]),
			LFART:         SAPStringClean(data[key][2]),
			VKBUR:         SAPStringClean(data[key][3]),
			ROUTE:         SAPStringClean(data[key][4]),
			KUNAG:         SAPStringClean(data[key][5]),
			KUNNR:         SAPStringClean(data[key][6]),
			VLSTK:         SAPStringClean(data[key][7]),
			WADAT_IST:     SAPStringClean(data[key][8]),
			SPE_WAUHR_IST: SAPStringClean(data[key][9]),
			ZUKRL:         SAPStringClean(data[key][10]),
			LGNUM:         SAPStringClean(data[key][11]),
		}
		likpdata = append(likpdata, row)
	}
	return likpdata
}

func LoadKNMTStruct(searchtype, searchvalue string) []*s.Knmt {
	connect()
	data, _ := ReadTable("KNMT", searchtype, searchvalue, reflect.ValueOf(s.Knmt{}), reflect.TypeOf(s.Knmt{}))
	close()

	var knmtdata []*s.Knmt

	for key := range data {
		row := &s.Knmt{
			MATNR: SAPStringClean(data[key][0]),
			KUNNR: SAPStringClean(data[key][1]),
			VKORG: SAPStringClean(data[key][2]),
			KDMAT: SAPStringClean(data[key][3]),
			MEINS: SAPStringClean(data[key][4]),
			ERNAM: SAPStringClean(data[key][5]),
			ERDAT: SAPStringClean(data[key][6]),
		}
		knmtdata = append(knmtdata, row)
	}
	return knmtdata
}

func LoadMBEWStruct(searchtype, searchvalue string) []*s.Mbew {
	connect()
	data, _ := ReadTable("MBEW", searchtype, searchvalue, reflect.ValueOf(s.Mbew{}), reflect.TypeOf(s.Mbew{}))
	close()

	var mbewdata []*s.Mbew

	for key := range data {
		row := &s.Mbew{
			MATNR: SAPStringClean(data[key][0]),
			LBKUM: SAPStringClean(data[key][1]),
			SALK3: SAPStringClean(data[key][2]),
			VPRSV: SAPStringClean(data[key][3]),
			VERPR: SAPStringClean(data[key][4]),
			STPRS: SAPStringClean(data[key][5]),
			PEINH: SAPStringClean(data[key][6]),
			BWKEY: SAPStringClean(data[key][7]),
		}
		mbewdata = append(mbewdata, row)
	}
	return mbewdata
}
