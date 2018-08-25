package main

import (
	"net/http"
	"strconv"

	"github.com/RTradeLtd/VaaS/ethereum"
	"github.com/gin-gonic/gin"
)

type API struct {
	Router *gin.Engine
}

// InitializeAPI is used to generate our API
func InitializeAPI() *API {
	api := API{}
	router := gin.Default()
	api.Router = router
	api.Router.POST("/api/v1/ethereum/generator", api.GenerateEthereumKeys)
	return &api
}

// GenerateEthereumKeys is used to generate our ethereum key
func (api *API) GenerateEthereumKeys(c *gin.Context) {
	searchPrefix, exists := c.GetPostForm("search_prefix")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "search_prefix post form does not exist",
		})
		return
	}
	runTimeInSecondsString := c.PostForm("run_time_in_seconds:")
	var runTime int64
	var err error
	if runTimeInSecondsString == "" {
		runTime = 1000000000
	} else {
		runTime, err = strconv.ParseInt(runTimeInSecondsString, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to convert run time to int",
			})
			return
		}
	}
	eg := ethereum.InitializeEthereumGenerator(searchPrefix, runTime)
	eg.Run(c)
}
