package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RTradeLtd/VaaS/ethereum"
	"github.com/gin-gonic/gin"
	"github.com/lytics/grid"
)

type API struct {
	client *grid.Client
	Router *gin.Engine
}

// InitializeAPI is used to generate our API
func InitializeAPI(gc *grid.Client) *API {
	api := API{}
	router := gin.Default()
	api.Router = router
	api.Router.POST("/api/v1/ethereum/generate/locally", api.GenerateEthereumKeysLocally)
	if gc != nil {
		api.client = gc
		api.Router.POST("/api/v1/ethereum/generate/distributed/:worker", api.GenerateEthereumKeysDistributedly)
	}
	return &api
}

// GenerateEthereumKeysLocally is used to generate our ethereum key locally, on the API node
func (api *API) GenerateEthereumKeysLocally(c *gin.Context) {
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
	eg.RunAPI(c)
}

func (api *API) GenerateEthereumKeysDistributedly(c *gin.Context) {
	worker := c.Param("worker")
	searchPrefix, exists := c.GetPostForm("search_prefix")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "search_prefix post form does not exist",
		})
		return
	}

	genReq := &GenerationRequest{
		SearchPrefix: searchPrefix,
	}

	resp, err := api.client.Request(time.Second*2, worker, genReq)
	fmt.Printf("response %#v\nerr %v\n", resp, err)
	if gr, ok := resp.(*GenerationResponse); ok {
		c.JSON(http.StatusOK, gin.H{
			"key":     fmt.Sprintf("Resposne %s", gr.Key),
			"address": fmt.Sprintf("Address %s", gr.Address),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "wrong resposne type",
	})
	return
}
