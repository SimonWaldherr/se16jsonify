package main

import (
	s "./structs"
	"encoding/json"
	"io"
	"net/http"
	"simonwaldherr.de/go/golibs/as"
	"strings"
	"fmt"
)

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

func likphandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/likp/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "VBELN", as.FixedLengthBefore(url, "0", 10)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "VBELN" && searchtype != "VKBUR" && searchtype != "KUNAG" && searchtype != "KUNNR" && searchtype != "ZUKRL" {
			searchtype = "VBELN"
		}

		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")

		if searchtype == "VBELN" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		} else if searchtype == "KUNNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		} else if searchtype == "KUNAG" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		} else {
			searchvalue = "%" + searchvalue
		}
	}

	jstr, _ := json.Marshal(LoadLIKPStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func knmthandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/knmt/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "MATNR", as.FixedLengthBefore(url, "0", 18)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "MATNR" && searchtype != "KUNNR" && searchtype != "KDMAT" {
			searchtype = "MATNR"
		}

		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")

		if searchtype == "MATNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 18)
		} else if searchtype == "KUNNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		} else {
			searchvalue = "%" + searchvalue
		}
	}

	jstr, _ := json.Marshal(LoadKNMTStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func mbewhandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/mbew/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		if url == "*" {
			searchtype, searchvalue = "*", "LBKUM > '0'"
		} else {
			searchtype, searchvalue = "MATNR", as.FixedLengthBefore(url, "0", 18)
		}
	} else {
		searchvalue = x[1]
		searchtype = "MATNR"
		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")
		searchvalue = as.FixedLengthBefore(searchvalue, "0", 18)
	}

	mbew := LoadMBEWStruct(searchtype, searchvalue)
	var jstr []byte

	if searchtype == "*" {
		var gld1000, gld1290, gldrest, gldline float64
		for key := range mbew {
			if mbew[key].VPRSV == "V" {
				gldline = as.Float(mbew[key].SALK3)
			} else {
				gldline = as.Float(mbew[key].LBKUM) * as.Float(mbew[key].VERPR) / as.Float(mbew[key].PEINH)
			}
			if mbew[key].BWKEY == "1000" {
				gld1000 += gldline
			} else if mbew[key].BWKEY == "1290" {
				gld1290 += gldline
			} else {
				gldrest += gldline
			}

			var mbewdata []*s.Mbew

			gld1000mbew := &s.Mbew{
				MATNR: "*",
				LBKUM: "",
				SALK3: fmt.Sprintf("%f", gld1000),
				VPRSV: "",
				VERPR: "",
				STPRS: "",
				PEINH: "",
				BWKEY: "1000",
			}
			mbewdata = append(mbewdata, gld1000mbew)
			gld1290mbew := &s.Mbew{
				MATNR: "*",
				LBKUM: "",
				SALK3: fmt.Sprintf("%f", gld1290),
				VPRSV: "",
				VERPR: "",
				STPRS: "",
				PEINH: "",
				BWKEY: "1290",
			}
			mbewdata = append(mbewdata, gld1290mbew)

			jstr, _ = json.Marshal(mbewdata)
		}
	} else {
		jstr, _ = json.Marshal(mbew)
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func vbfahandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/vbfa/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "VBELN", as.FixedLengthBefore(url, "0", 18)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "VBELN" && searchtype != "VBELV" {
			searchtype = "VBELN"
		}

		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")
		searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
	}

	jstr, _ := json.Marshal(LoadVBFAStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}

func vbbehandler(rw http.ResponseWriter, req *http.Request) (string, int) {
	var searchtype, searchvalue string
	url := strings.Replace(req.RequestURI, "/vbbe/", "", 1)
	x := strings.Split(url, "/")

	if len(x) < 2 {
		searchtype, searchvalue = "MATNR", as.FixedLengthBefore(url, "0", 18)
	} else {
		searchtype, searchvalue = x[0], x[1]
		searchtype = strings.ToUpper(searchtype)

		if searchtype != "MATNR" && searchtype != "WERKS" && searchtype != "LGORT" && searchtype != "KUNNR" && searchtype != "VBELN" {
			searchtype = "MATNR"
		}
		
		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")
		
		if searchtype == "MATNR" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 18)
		} else if searchtype == "VBELN" {
			searchvalue = as.FixedLengthBefore(searchvalue, "0", 10)
		}
	}
	
	searchvalue = fmt.Sprintf(" %s = '%s' ", searchtype, searchvalue)
	searchtype = "*"

	jstr, _ := json.Marshal(LoadVBBEStruct(searchtype, searchvalue))
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, string(jstr))
	return "", http.StatusOK
}
