package api

import (
	"fmt"
	"github.com/ccloud/phase2/pkg/model"
	"github.com/ccloud/phase2/pkg/redis"
	"github.com/ccloud/phase2/pkg/retreiver"
	"github.com/gin-gonic/gin"
	redis2 "github.com/go-redis/redis"
	"github.com/pkg/errors"
	"net/http"
)

func PriceHandler(c *gin.Context) {
	var priceReq model.PriceReq
	if err := c.ShouldBindJSON(&priceReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(priceReq.Name)
	var price float64

	cacheVal, err := redis.Cache.CheckPrice(priceReq.Name)
	if err == redis2.Nil {
		fmt.Println("not found in redis.")
		retrievedPrice, err2 := retreiver.RetrievePrice(priceReq.Name)
		if err2 != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errors.Wrap(err2, "unable to retrieve from api").Error(),
			})
			return
		}
		price = retrievedPrice
		fmt.Println(redis.Cache.SetPrice(priceReq.Name, price))

	} else if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errors.Wrap(err, "unable to retrieve from cache").Error(),
		})
		return

	} else {
		fmt.Println("found in cache")
		price = cacheVal
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":  priceReq.Name,
		"Price": price,
	})
}
