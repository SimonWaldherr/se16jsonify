package main

import (
	"./conn"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"simonwaldherr.de/go/golibs/as"
	"simonwaldherr.de/go/gwv"
	"simonwaldherr.de/go/saprfc"
	"strings"
	"time"
)

var SAPconnection *saprfc.Connection

type kna1 struct {
	KUNNR string
	LAND1 string
	NAME1 string
	NAME2 string
	NAME3 string
	ORT01 string
	PSTLZ string
	STRAS string
	TELF1 string
	TELF2 string
}

type mara struct {
	MATNR string
	MATKL string
	MEINS string
	BRGEW string
	GEWEI string
	EAN11 string
	MFRPN string
	PRDHA string
	MAKTX string
	MAKTG string
}

type lips struct {
	VBELN string
	POSNR string
	MATNR string
	MATKL string
	ARKTX string
	EANNR string
	LGORT string
	LFIMG string
	VRKME string
	VKBUR string
}

func abapSystem() saprfc.ConnectionParameter {
	return saprfc.ConnectionParameter{
		Dest:      conn.Dest,
		Client:    conn.Client,
		User:      conn.User,
		Passwd:    conn.Passwd,
		Lang:      conn.Lang,
		Ashost:    conn.Ashost,
		Sysnr:     conn.Sysnr,
		Saprouter: conn.Saprouter,
	}
}

func connect() error {
	var err error
	SAPconnection, err = saprfc.ConnectionFromParams(abapSystem())
	return err
}

func close() {
	SAPconnection.Close()
}

type OpStruct struct {
	TEXT string
}

type FieldStruct struct {
	FIELDNAME string
	OFFSET    int
	LENGTH    int
	T         string
	FIELDTEXT string
}
type DatStruct struct {
	WA string
}

type parameter struct {
	QUERY_TABLE string
	DELIMITER   string
	NO_DATA     string
	ROWSKIPS    int
	ROWCOUNT    int
	OPTIONS     []OpStruct
	FIELDS      []FieldStruct
	DATA        []DatStruct
}

func RfcReadTable(table, query string, fields []FieldStruct) ([][]string, error) {
	opstruct := []OpStruct{OpStruct{query}}
	dat := []DatStruct{DatStruct{""}}
	params := parameter{table, ";", "", 0, 2000000, opstruct, fields, dat}

	r, err := SAPconnection.Call("RFC_READ_TABLE", params)
	if err != nil {
		return [][]string{}, err
	}

	var app bool
	var ret [][]string

	echoStruct := r["DATA"].([]interface{})
	for _, value := range echoStruct {
		values := value.(map[string]interface{})
		for _, val := range values {
			if strings.Trim(val.(string), "\t\n\r ;") != "" {
				valstr := strings.Split(val.(string), ";")
				if app {
					ret = append(ret, valstr)
				} else {
					ret = [][]string{valstr}
					app = true
				}
			}
		}
	}
	return ret, nil
}

func ReadKNA1(searchtype, searchvalue string) ([][]string, error) {
	fields := []FieldStruct{
		FieldStruct{"KUNNR", 0, 0, "", ""},
		FieldStruct{"LAND1", 0, 0, "", ""},
		FieldStruct{"NAME1", 0, 0, "", ""},
		FieldStruct{"NAME2", 0, 0, "", ""},
		FieldStruct{"NAME3", 0, 0, "", ""},
		FieldStruct{"ORT01", 0, 0, "", ""},
		FieldStruct{"PSTLZ", 0, 0, "", ""},
		FieldStruct{"STRAS", 0, 0, "", ""},
		FieldStruct{"TELF1", 0, 0, "", ""},
		FieldStruct{"TELF2", 0, 0, "", ""},
	}
	query := fmt.Sprintf("%s LIKE '%s'", searchtype, searchvalue)
	data, err := RfcReadTable("KNA1", query, fields)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func ReadMARA(searchtype, searchvalue string) ([][]string, error) {
	fields := []FieldStruct{
		FieldStruct{"MATNR", 0, 0, "", ""},
		FieldStruct{"MATKL", 0, 0, "", ""},
		FieldStruct{"MEINS", 0, 0, "", ""},
		FieldStruct{"BRGEW", 0, 0, "", ""},
		FieldStruct{"GEWEI", 0, 0, "", ""},
		FieldStruct{"EAN11", 0, 0, "", ""},
		FieldStruct{"MFRPN", 0, 0, "", ""},
		FieldStruct{"PRDHA", 0, 0, "", ""},
		FieldStruct{"MAKTX", 0, 0, "", ""},
		FieldStruct{"MAKTG", 0, 0, "", ""},
	}
	query := fmt.Sprintf("%s LIKE '%s'", searchtype, searchvalue)
	data, err := RfcReadTable("MARAV", query, fields)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func ReadLIPS(searchtype, searchvalue string) ([][]string, error) {
	fields := []FieldStruct{
		FieldStruct{"VBELN", 0, 0, "", ""},
		FieldStruct{"POSNR", 0, 0, "", ""},
		FieldStruct{"MATNR", 0, 0, "", ""},
		FieldStruct{"MATKL", 0, 0, "", ""},
		FieldStruct{"ARKTX", 0, 0, "", ""},
		FieldStruct{"EANNR", 0, 0, "", ""},
		FieldStruct{"LGORT", 0, 0, "", ""},
		FieldStruct{"LFIMG", 0, 0, "", ""},
		FieldStruct{"VRKME", 0, 0, "", ""},
		FieldStruct{"VKBUR", 0, 0, "", ""},
	}
	query := fmt.Sprintf("%s LIKE '%s'", searchtype, searchvalue)
	data, err := RfcReadTable("LIPS", query, fields)
	if err != nil {
		fmt.Println(err)
	}
	return data, err
}

func SAPBoolConv(str string) bool {
	if str == "X" || str == "x" {
		return true
	}
	return false
}

func SAPTimeConv(itime, idate string) time.Time {
	t, err := time.Parse("20060102 150405", idate+" "+itime)
	if err != nil {
		fmt.Printf("in: %v\nout: %v\nerr: %v\n\n", idate+" "+itime, t, err)
	}
	return t
}

func SAPStringClean(str string) string {
	return strings.Trim(str, " \t")
}

func stripCtlAndExtFromUTF8(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= 32 && r < 127 {
			return r
		}
		return -1
	}, str)
}

func LoadKNAStruct(searchtype, searchvalue string) []*kna1 {
	connect()
	data, _ := ReadKNA1(searchtype, searchvalue)
	close()

	var knadata []*kna1

	for key := range data {
		row := &kna1{
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

func LoadMARAStruct(searchtype, searchvalue string) []*mara {
	connect()
	data, _ := ReadMARA(searchtype, searchvalue)
	close()

	var maradata []*mara

	for key := range data {
		row := &mara{
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
		}
		maradata = append(maradata, row)
	}
	return maradata
}

func LoadLIPSStruct(searchtype, searchvalue string) []*lips {
	connect()
	data, _ := ReadLIPS(searchtype, searchvalue)
	close()

	var lipsdata []*lips

	for key := range data {
		row := &lips{
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

func knahandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/kna1/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "KUNNR", as.FixedLengthBefore(url, "0", 10)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "KUNNR" && searchtype != "NAME1" && searchtype != "SORTL" && searchtype != "PSTLZ" && searchtype != "BRSCH" {
			searchtype = "KUNNR"
		}

		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")

		if searchtype == "KUNNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		} else {
			searchvalue = "%" + searchvalue + "%"
		}
	}

	jstr, _ := json.Marshal(LoadKNAStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func marahandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/mara/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "MATNR", as.FixedLengthBefore(url, "0", 18)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "MATNR" && searchtype != "MAKTX" && searchtype != "MAKTG" && searchtype != "EAN11" {
			searchtype = "MATNR"
		}

		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")

		if searchtype == "MATNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 18)
		} else {
			searchvalue = "%" + searchvalue + "%"
		}
	}

	jstr, _ := json.Marshal(LoadMARAStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func lipshandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/lips/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "VBELN", as.FixedLengthBefore(url, "0", 10)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "VBELN" && searchtype != "MATNR" && searchtype != "MATKL" && searchtype != "EANNR" {
			searchtype = "VBELN"
		}

		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")

		if searchtype == "VBELN" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		} else if searchtype == "MATNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 18)
		} else {
			searchvalue = "%" + searchvalue + "%"
		}
	}

	jstr, _ := json.Marshal(LoadLIPSStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func main() {
	HTTPD := gwv.NewWebServer(8082, 60)

	HTTPD.URLhandler(
		gwv.URL("^/kna1/.*$", knahandler, gwv.MANUAL),
		gwv.URL("^/mara/.*$", marahandler, gwv.MANUAL),
		gwv.URL("^/lips/.*$", lipshandler, gwv.MANUAL),
	)

	HTTPD.Start()
	HTTPD.WG.Wait()
}
