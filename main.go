package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func strip_port() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := strings.Split(c.ClientIP(), ":")[0]
		//fmt.Println(strings.Split(ip_port, ":")[0])
		c.String(200, ip)
	}
}

func main() {
	/*
	   FYI
	   gin.context doesn't work onwith ipv6
	*/
	gin.SetMode(gin.ReleaseMode)

	port := "8080"
	iface := os.Getenv("INTERFACE")

	r := gin.Default()

	r.GET("/", strip_port())

	port_env := os.Getenv("PORT")

	if port_env != "" {
		port = port_env
	}

	r.Run(iface + ":" + port)
}
