package metrics

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VisitMetric struct {
	Location   string `json:"location" binding:"required"`
	DeviceHash string `json:"device_hash" binding:"required"`
	ErplyID    string `json:"erply_id" binding:"required"`
}

type VisitStore interface {
	Add(metric VisitMetric) error
}

type Endpoint struct {
	VisitStore VisitStore
}

/**
@api {post} /visits Visit metric create
@apiDescription Adds a new visit metric
@apiName Visit create
@apiGroup Visit
@apiUse JsonHeader
@apiUse AuthFrontHeader

@apiParamExample {json} Body:
{
	"location": 	"Rome", #required
	"device_hash": 	"djfasdlfjlfjasdlkfas", #required
	"erply_id":		"100234" #required
}

@apiSuccessExample Success-Response
HTTP/1.1 200 OK
{
	"message":"Success"
}

@apiErrorExample Bad request(400)
HTTP/1.1 400 Bad request
{
    "error": "Key: 'VisitMetric.Location' Error:Field validation for 'Location' failed on the 'required' tag"
}

@apiPermission registered user
*/
//CreateVisitsHandler registers API to add visit metrics
func (e Endpoint) CreateVisitsHandler(c *gin.Context) {
	var visit VisitMetric
	if err := c.ShouldBindJSON(&visit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := e.VisitStore.Add(visit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Success"})
}

