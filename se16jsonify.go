package main

import (
	c "./conn"
	s "./structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"simonwaldherr.de/go/golibs/as"
	"simonwaldherr.de/go/golibs/structs"
	"simonwaldherr.de/go/gwv"
	"simonwaldherr.de/go/saprfc"
	"strings"
	"time"
)

var SAPconnection *saprfc.Connection

func abapSystem() saprfc.ConnectionParameter {
	return saprfc.ConnectionParameter{
		Dest:      c.Dest,
		Client:    c.Client,
		User:      c.User,
		Passwd:    c.Passwd,
		Lang:      c.Lang,
		Ashost:    c.Ashost,
		Sysnr:     c.Sysnr,
		Saprouter: c.Saprouter,
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

func ReadTable(searchtable, searchtype, searchvalue string, v reflect.Value, t reflect.Type) ([][]string, error) {
	fields := []FieldStruct{}
	
	structs.ReflectHelper(v, t, 0, func(name string, vtype string, value interface{}, depth int) {
		fields = append(fields, FieldStruct{name, 0, 0, "", ""})
	})

	query := fmt.Sprintf("%s LIKE '%s'", searchtype, searchvalue)
	data, err := RfcReadTable(searchtable, query, fields)
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
