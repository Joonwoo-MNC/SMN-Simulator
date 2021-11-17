package smn_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

type ResponseInfo struct {
	nfService       string `json:"nfService"`
	reqNFInstanceID string `json:"reqNFInstance"`
	reqTime         string `json:"reqTime"`
	data            string `json:"data"`
}



func NoticeData(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}
	log.Print("##############################\n")
	log.Print("2. Data Notification\n")
	log.Print("##############################\n")
		log.Print("##############################\n")
	log.Print("2. Data Notification Response\n")
	log.Print("##############################\n")
	// log.Println(reqBody)
	// replyto, _ := json.Marshal(reqBody)
	// reqData := []byte(`{ "nfService":"training-requseted"}`)
	c.JSON(http.StatusOK, gin.H{
		"nfService":     "test-nwdaf",
		"reqNFInstance": "test-mtlf",
		"reqTime":       reqBody["reqTime"],
		"data":          "finished",
	})

	// c.BindJSON(&replyto)

	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "test"
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "None"
	jsonStr, _ := json.Marshal(jsonBody)
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}
	resp, err := http.Post("http://localhost:24249/smn-service/v1/NoticeDataResp", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	
	
	log.Print("##############################\n")
	log.Print("3. Subscription Exchange Process\n")
	/*
	Add Subscription Exchange Code
	*/
	
	log.Print("##############################\n")
	resp1, err1 := http.Post("http://127.0.0.3:8000/smn-service/v1/SubscriptionExchange", "application/json", bytes.NewBuffer([]byte(jsonStr)))

	if err1 != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp1.Header)
		respBody, _ := ioutil.ReadAll(resp1.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}

	log.Print("##############################\n")
	log.Print("4. Data Transfer Process\n")
	/*
	Add Data Transfer Code
	*/
	
	log.Print("##############################\n")
	resp2, err2 := http.Post("http://localhost:24249/smn-service/v1/DataTransfer", "application/json", bytes.NewBuffer([]byte(jsonStr)))


	if err2 != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp2.Header)
		respBody, _ := ioutil.ReadAll(resp2.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	
	log.Print("##############################\n")
	log.Print("5. Notification with Data Transfer\n")
	/*
	Add Notification with Data Transfer Code
	*/
	log.Print("##############################\n")
	resp3, err3 := http.Post("http://localhost:24248/smn-service/v1/NotificationData", "application/json", bytes.NewBuffer([]byte(jsonStr)))


	if err3 != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp3.Header)
		respBody, _ := ioutil.ReadAll(resp3.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}


	return
}


func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/smn-service/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"POST",
		"/",
		Index,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/NoticeData",
		NoticeData,
	},
}
