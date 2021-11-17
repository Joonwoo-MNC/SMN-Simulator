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

func ResourceReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	/*
		Add Propagation Delay
	*/
	log.Print("##############################\n")
	log.Print("6. Resource Allocation Process\n")
	/*
		Add Resource Allocation Code
	*/
	log.Print("##############################\n")
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
	resp, err := http.Post("http://localhost:24245/smn-service/v1/ResourceReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func ResourceResp(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	/*
		Add Propagation Delay
	*/
	log.Print("##############################\n")
	log.Print("8. Session Response\n")
	log.Print("##############################\n")
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
	resp, err := http.Post("http://localhost:24243/smn-service/v1/ResourceResp", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func SessionReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}
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

	log.Print("##############################\n")
	log.Print("1. Session Request\n")
	/*
		Add Propatation Delay
	*/
	log.Print("##############################\n")

	resp, err := http.Post("http://localhost:24243/smn-service/v1/SessionReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func PagingReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	/*
		Add Propatation Delay
	*/
	log.Print("##############################\n")
	log.Print("7. Paging Process\n")
	/*
		Add Paging Code
	*/
	log.Print("##############################\n")

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
	resp, err := http.Post("http://localhost:24245/smn-service/v1/PagingReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func IntraHOReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	log.Print("##############################\n")
	log.Print("1. Mornitoring\n")
	log.Print("##############################\n")
	/*
		Add Propatation Delay
	*/
	log.Print("##############################\n")
	log.Print("3. Intra-SC Handover Request\n")
	log.Print("##############################\n")

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
	resp, err := http.Post("http://localhost:24243/smn-service/v1/IntraHOReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func InterHOReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	log.Print("##############################\n")
	log.Print("1. Mornitoring\n")
	log.Print("##############################\n")
	/*
		Add Propatation Delay
	*/
	log.Print("##############################\n")
	log.Print("3. Registration Request\n")
	log.Print("##############################\n")

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
	resp, err := http.Post("http://localhost:24250/smn-service/v1/InterReReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func IntraHOStateExchangeReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}
	/*
		Add Propagation Delay
	*/
	log.Print("##############################\n")
	log.Print("4. State Exchange Process\n")
	/*
		Add State Exchange Code
	*/
	log.Print("##############################\n")

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

	resp, err := http.Post("http://localhost:24246/smn-service/v1/IntraHOStateExchangeReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func IntraHOResp(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	/*
		Add Propagation Delay
	*/
	log.Print("##############################\n")
	log.Print("6. Release Resoruce Request\n")
	log.Print("##############################\n")

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

	resp, err := http.Post("http://localhost:24246/smn-service/v1/IntraReleaseReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func InterHOStateExchageReq(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	/*
		Add Propagation Delay
	*/
	log.Print("##############################\n")
	log.Print("5. State Exchange Process\n")
	/*
		Add State Exchange Code
	*/
	log.Print("##############################\n")

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

	resp, err := http.Post("http://localhost:24246/smn-service/v1/InterHOStateExchageReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)
	}
	return
}

func InterReResp(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}

	/*
		Add Propagation Delay
	*/
	log.Print("##############################\n")
	log.Print("7. Release Resoruce Request\n")
	log.Print("##############################\n")

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

	resp, err := http.Post("http://localhost:24246/smn-service/v1/InterReleaseReq", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		respBody, _ := ioutil.ReadAll(resp.Body)
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
		"/ResourceReq",
		ResourceReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/ResourceResp",
		ResourceResp,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/SessionReq",
		SessionReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/PagingReq",
		PagingReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/IntraHOReq",
		IntraHOReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/IntraHOStateExchangeReq",
		IntraHOStateExchangeReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/IntraHOResp",
		IntraHOResp,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/InterHOReq",
		InterHOReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/InterHOStateExchageReq",
		InterHOStateExchageReq,
	},

	{
		"smn-service",
		strings.ToUpper("Post"),
		"/InterReResp",
		InterReResp,
	},
}
