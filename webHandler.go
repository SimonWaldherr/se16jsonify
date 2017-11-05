package main

import (
	"encoding/json"
	"io"
	"net/http"
	"simonwaldherr.de/go/golibs/as"
	"strings"
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
		searchtype, searchvalue = "MATNR", as.FixedLengthBefore(url, "0", 18)
	} else {
		searchvalue = x[1]
		searchtype = "MATNR"
		searchvalue = strings.Trim(searchvalue, "%\t\n\r ")
		searchvalue = as.FixedLengthBefore(searchvalue, "0", 18)
	}

	jstr, _ := json.Marshal(LoadMBEWStruct(searchtype, searchvalue))
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
