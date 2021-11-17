package smn_service

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	"strings"
	//"time"

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



func SubscriptionExchange(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute

	log.Print("##############################\n")
	log.Print("Subscription Exchange Process\n")
	/*
	Add Subscription Exchange Code	
	*/
	log.Print("##############################\n")
	
	return
}


func InterLocationUpdate(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	
	log.Print("##############################\n")
	log.Print("8. Location Update Process\n")
	/*
	Add Location Update Code
	*/
	log.Print("##############################\n")

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
		"/SubscriptionExchange",
		SubscriptionExchange,
	},
	
	{
		"smn-service",
		strings.ToUpper("Post"),
		"/InterLocationUpdate",
		InterLocationUpdate,
	},
	
	
	
}
