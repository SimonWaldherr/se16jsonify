package main

import (
	c "./conn"
	"fmt"
	"reflect"
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

func main() {
	HTTPD := gwv.NewWebServer(8088, 60)

	HTTPD.URLhandler(
		gwv.URL("^/kna1/.*$", knahandler, gwv.MANUAL),
		gwv.URL("^/mara/.*$", marahandler, gwv.MANUAL),
		gwv.URL("^/lips/.*$", lipshandler, gwv.MANUAL),
		gwv.URL("^/likp/.*$", likphandler, gwv.MANUAL),
		gwv.URL("^/knmt/.*$", knmthandler, gwv.MANUAL),
		gwv.URL("^/mbew/.*$", mbewhandler, gwv.MANUAL),
		gwv.URL("^/vbfa/.*$", vbfahandler, gwv.MANUAL),
	)

	HTTPD.Start()
	HTTPD.WG.Wait()
}
