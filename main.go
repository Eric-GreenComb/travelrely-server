package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/travelrely-server/config"
	"github.com/Eric-GreenComb/travelrely-server/handler"
	"github.com/Eric-GreenComb/travelrely-server/persist"
)

// Cors Cors
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "1728000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

var (
	g errgroup.Group
)

var router *gin.Engine

func main() {
	if config.ServerConfig.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDatabase()

	router = gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api */
	router.GET("/", handler.Index)
	router.GET("/health", handler.Health)

	// database api
	r1 := router.Group("api/v1")
	{
		/* user */
		r1.POST("/user/register", handler.RegisterUser)
		r1.POST("/user/unregister", handler.UnRegisterUser)

		// msisdn
		r1.POST("/msisdn/subscribe", handler.SubscribeMsisdn)
		r1.POST("/msisdn/unsubscribe", handler.UnsubscribeMsisdn)
		r1.POST("/msisdn/state", handler.GetMsisdnState)
		r1.POST("/msisdn/history", handler.GetMsisdnHistory)
		r1.POST("/asset/info", handler.GetAssetInfo)

		// transfers
		r1.POST("/transfer/create", handler.CreateTransfer)
		r1.POST("/transfer/query", handler.QueryTransfer)

		// block
		r1.GET("/channels/:channel/height", handler.QueryChainHeight)
	}

	for _, _port := range config.ServerConfig.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  5 * time.Minute,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return server.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
