package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lightcast/happiness/constants"
	"lightcast/happiness/model"
	"net/http"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "service is healthy"})
	}
}

func HappinessByFacetIdHandler(data model.HappinessIndexMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		facet := c.Param(constants.FACET) // need to validate
		facet_id := c.Param(constants.FACET_ID)

		index_data := model.NewHappinessByFacetIdData(facet)
		value, found := data[facet_id]
		if !found {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("county '%s' not found", facet_id)})
			return
		}

		index_data.Indices = append(index_data.Indices, model.IndexMapping{Id: facet_id, Value: value})
		c.IndentedJSON(http.StatusOK, model.NewHappinessByFacetIdResponse(*index_data))
	}
}

func HappinessByFacetIdsHandler(data model.HappinessIndexMap) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request model.HappinessByFacetIdsRequest
		if err := c.BindJSON(&request); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			return
		}

		facet := c.Param(constants.FACET)
		index_data := model.NewHappinessByFacetIdsData(facet)
		for _, county := range request.Counties {
			if value, found := data[county]; found {
				index_mapping := model.IndexMapping{Id: county, Value: value}
				index_data.Indices = append(index_data.Indices, index_mapping)
			}
		}
		metrics, err := ComputeMetrics(index_data.Indices, request.Metrics)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		index_data.Metrics = metrics

		c.IndentedJSON(http.StatusOK, model.NewHappinessByFacetIdsResponse(*index_data))
	}
}
