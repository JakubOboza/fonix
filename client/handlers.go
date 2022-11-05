package client

import (
	"fmt"
	"net/http"
)

/*
  Example Mo & Dr handlers
*/

const (
	DEFAULT_MO_PATH = "/mock/mosms"
	DEFAULT_DR_PATH = "/mock/drs"
)

// IFVERSION=201001&MONUMBER=447111222333&OPERATOR=o2-uk&DESTINATION=84988&BODY=This%20is%20a%20mobile%20originated%20test%20message&RECEIVETIME=20130202102030&GUID=7CDEB38F-4370-18FD-D7CE-329F21B99209&PRICE=15

func MoHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println("MO REQUEST")
	fmt.Println("Params: ", r.Form.Encode())

	ifversion := r.FormValue("IFVERSION")
	monumber := r.FormValue("MONUMBER")
	operator := r.FormValue("OPERATOR")
	destination := r.FormValue("DESTINATION")
	body := r.FormValue("BODY")
	receiveTime := r.FormValue("RECEIVETIME")
	guid := r.FormValue("GUID")
	price := r.FormValue("PRICE")
	requestID := r.FormValue("REQUESTID")

	fmt.Printf("IFVERSION: %s\nMONUMBER: %s\nOPERATOR: %s\nDESTINATION: %s\nBODY: %s\nRECEIVETIME: %s\nGUID: %s\nPRICE: %s\nREQUESTID: %s\n\n", ifversion, monumber, operator, destination, body, receiveTime, guid, price, requestID)

	w.WriteHeader(http.StatusOK)
}

//  IFVERSION=201001&OPERATOR=o2-uk&MONUMBER=447111222333&DESTINATION=84988
//  &STATUSCODE=DELIVERED&STATUSTEXT=Delivered&STATUSTIME=20130202102030
//  &PRICE=50&GUID=7CDEB38F-4370-18FD-D7CE-329F21B99209

func DrHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Println("DR REQUEST")
	fmt.Println("Params: ", r.Form.Encode())

	ifversion := r.FormValue("IFVERSION")
	monumber := r.FormValue("MONUMBER")
	operator := r.FormValue("OPERATOR")
	destination := r.FormValue("DESTINATION")
	body := r.FormValue("BODY")
	receiveTime := r.FormValue("RECEIVETIME")
	guid := r.FormValue("GUID")
	price := r.FormValue("PRICE")
	requestID := r.FormValue("REQUESTID")

	fmt.Printf("IFVERSION: %s\nMONUMBER: %s\nOPERATOR: %s\nDESTINATION: %s\nBODY: %s\nRECEIVETIME: %s\nGUID: %s\nPRICE: %s\nREQUESTID: %s\n\n", ifversion, monumber, operator, destination, body, receiveTime, guid, price, requestID)

	w.WriteHeader(http.StatusOK)
}

func StartMockDrMoHandler(port int) {
	portString := ":8090"
	http.HandleFunc(DEFAULT_MO_PATH, MoHandler)
	http.HandleFunc(DEFAULT_DR_PATH, DrHandler)
	if port > 1024 && port < 65535 {
		portString = fmt.Sprintf(":%d", port)
	}
	fmt.Printf("Starting mo/dr mock on port: %s\n", portString)
	http.ListenAndServe(portString, nil)
}
