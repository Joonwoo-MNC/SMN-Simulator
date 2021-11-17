package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func UESession(reqNfInstanceId string) { //TODO: Change input data 'data' to appropriate attribute
	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = reqNfInstanceId
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "none"
	jsonStr, _ := json.Marshal(jsonBody)
	//print("reqNFInstanceID: %s", reqNfInstanceId)
	resp, err := http.Post("http://localhost:24245/smn-service/v1/SessionReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}

}

func NWSession(reqNfInstanceId string) { //TODO: Change input data 'data' to appropriate attribute
	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = reqNfInstanceId
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "none"
	jsonStr, _ := json.Marshal(jsonBody)
	//print("reqNFInstanceID: %s", reqNfInstanceId)
	resp, err := http.Post("http://localhost:24249/smn-service/v1/Data", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}

}


func IntraHO(reqNfInstanceId string) { //TODO: Change input data 'data' to appropriate attribute
	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = reqNfInstanceId
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "none"
	jsonStr, _ := json.Marshal(jsonBody)
	//print("reqNFInstanceID: %s", reqNfInstanceId)
	resp, err := http.Post("http://localhost:24246/smn-service/v1/IntraHOReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	
	resp1, err1 := http.Post("http://localhost:24249/smn-service/v1/IntraHOReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	
	if err1 != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp1.Header)
		respBody, _ := ioutil.ReadAll(resp1.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}

}

func InterHO(reqNfInstanceId string) { //TODO: Change input data 'data' to appropriate attribute
	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = reqNfInstanceId
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "none"
	jsonStr, _ := json.Marshal(jsonBody)
	//print("reqNFInstanceID: %s", reqNfInstanceId)
	resp, err := http.Post("http://localhost:24246/smn-service/v1/InterHOReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	
	resp1, err1 := http.Post("http://localhost:24249/smn-service/v1/InterHOReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	
	if err1 != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp1.Header)
		respBody, _ := ioutil.ReadAll(resp1.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}

}

func main() {
	n := 1
	var selection int

	for n < 10 {
		fmt.Println("\n\n####################################")
		fmt.Print("\nChoose the Procedures \n 1) UE-triggered Session Establishment\n 2) Netwrok-triggered Session Establishmnet\n 3) Intra-SC Handover\n 4) Inter-SC Handover\n : ")
		
		fmt.Scanln(&selection)
		if selection == 1 {
			UESession("UE-triggered Session Establishmnet")
		} else if selection == 2 {
			NWSession("Network-triggered Session Establishmnet")
		} else if selection == 3 {
			IntraHO("Intra-SC Handover")
		} else if selection == 4 {
			InterHO("Inter-SC Handover")
		} else {
			fmt.Println("Wrong number")
		}
	}

}
