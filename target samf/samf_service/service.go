package samf_service

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

func InterHOReport(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute

	log.Print("##############################\n")
	log.Print("2. Report RAN Usage Data\n")
	log.Print("##############################\n")
}

func InterReReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	log.Print("##############################\n")
	log.Print("4. Admission Control Process\n")
	/*
		Add Admission Control Code
	*/
	log.Print("##############################\n")
	log.Print("##############################\n")
	log.Print("5. State Exchange Process\n")
	/*
		Add State Exchange Code
	*/
	log.Print("##############################\n")

	c.JSON(http.StatusOK, gin.H{
		"nfService":     "smn",
		"reqNFInstance": "test cp",
		"reqTime":       reqBody["reqTime"],
		"data":          "finished",
	})

	// c.BindJSON(&replyto)

	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "test"
	jsonBody["nfService"] = "smn"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "None"
	jsonStr, _ := json.Marshal(jsonBody)
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}

	resp, err := http.Post("http://localhost:24249/smn-service/v1/InterHOStateExchageReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		//fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		//fmt.Println(jsonData)
	}
	return
}

func InterHOStateExchageResp(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	log.Print("##############################\n")
	log.Print("6. Registration Response\n")
	log.Print("##############################\n")

	c.JSON(http.StatusOK, gin.H{
		"nfService":     "smn",
		"reqNFInstance": "test cp",
		"reqTime":       reqBody["reqTime"],
		"data":          "finished",
	})

	// c.BindJSON(&replyto)

	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "test"
	jsonBody["nfService"] = "smn"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "None"
	jsonStr, _ := json.Marshal(jsonBody)
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}

	resp, err := http.Post("http://localhost:24249/smn-service/v1/InterReResp", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		//fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		//fmt.Println(jsonData)
	}

	log.Print("##############################\n")
	log.Print("8. Location Update Request\n")
	log.Print("##############################\n")

	resp1, err1 := http.Post("http://localhost:24247/smn-service/v1/InterLocationUpdate", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err1 != nil {
		fmt.Println("error: %v", err1)
	} else {
		//fmt.Println(resp1.Header)
		respBody, _ := ioutil.ReadAll(resp1.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		//fmt.Println(jsonData)
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
		"/InterHOReport",
		InterHOReport,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/InterReReq",
		InterReReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/InterHOStateExchageResp",
		InterHOStateExchageResp,
	},
}
