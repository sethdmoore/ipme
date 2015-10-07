package main

import (
	"fmt"
	"github.com/sethdmoore/ipme/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"os"
	"strings"
)

func strip_port() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := strings.Split(c.ClientIP(), ":")[0]
		//fmt.Println(strings.Split(ip_port, ":")[0])
		fmt.Println(ip)
		c.String(200, ip)
	}
}

func main() {
	/*
	   FYI
	   gin.context doesn't work onwith ipv6
	*/

	port := "8080"
	iface := os.Getenv("INTERFACE")

	r := gin.Default()

	r.GET("/", strip_port())

	port_env := os.Getenv("PORT")

	if port_env != "" {
		port = port_env
	}

	fmt.Println(port)
	r.Run(iface + ":" + port)
}
