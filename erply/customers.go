package erply

import (
	"fmt"
	"github.com/breathbath/erplyapi/cache"
	"github.com/breathbath/erplyapi/cache/redis"
	"github.com/breathbath/go_utils/utils/env"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

//BuildCustomers wraps boilerplate code to build Customers API
func BuildCustomers() Customers {
	cacheClient := redis.Client{}
	authProvider := AuthProvider{
		CacheClient: cacheClient,
		AuthFunc:    api.VerifyUser,
	}
	clientFactory := DefaultClientFactory{
		AuthProvider: authProvider,
	}
	ca := Customers{
		ClientFactory: clientFactory,
		Cache:         cacheClient,
	}

	return ca
}

//Customers wraps customers API of ERPLY
type Customers struct {
	ClientFactory ClientFactory
	Cache         cache.Client
}

//Get wraps getCustomers method
func (c Customers) GetById(ids []string) (customers api.Customers, err error) {
	customers = api.Customers{}

	log.Debugf("Will get customers by IDs %v from Erply API", ids)
	if len(ids) == 0 {
		return customers, fmt.Errorf("no ids provided in input")
	}

	log.Debug("Will try to find customer IDs in cache")
	idsToFetchFromAPI := make([]string, 0, len(ids))
	for _, id := range ids {
		var customer api.Customer
		found, err := c.Cache.Read(customersCacheKey+"_"+id, &customer)
		if err != nil {
			return customers, err
		}

		if found {
			customers = append(customers, customer)
			log.Debugf("Found customer with ID %s in cache", id)
		} else {
			idsToFetchFromAPI = append(idsToFetchFromAPI, id)
			log.Debugf("Didn't find customer with ID %s in cache", id)
		}
	}
	if len(idsToFetchFromAPI) == 0 {
		log.Debug("All customer ids are found in cache")
		return
	}

	log.Debugf("Will fetch customers with ids %v from Erply API", idsToFetchFromAPI)
	cl, err := c.ClientFactory.CreateClient()
	if err != nil {
		return customers, err
	}

	customersFromApi, err := cl.GetCustomersByIDs(idsToFetchFromAPI)
	if err != nil {
		return customers, err
	}

	log.Debugf("Got %d customers from Erply API", len(idsToFetchFromAPI))

	cacheTtl := time.Duration(env.ReadEnvInt("ERPLY_DATA_CACHE_TTL_SECONDS", 3600))
	for _, ctmr := range customersFromApi {
		customers = append(customers, ctmr)
		err = c.Cache.Store(fmt.Sprintf("%s_%d", customersCacheKey, ctmr.ID), ctmr, time.Second*cacheTtl)
		if err != nil {
			return
		}
	}

	return
}

//Save wraps saveCustomer method
func (c Customers) Save(input *api.CustomerConstructor) (report *api.CustomerImportReport, err error) {
	cl, err := c.ClientFactory.CreateClient()
	if err != nil {
		return
	}
	log.Debugf("Will try to store customer data in API: %+v", input)

	report, err = cl.PostCustomer(input)
	if err != nil {
		return
	}

	log.Debug("Successfully stored customer data in API")

	return
}

//GetCustomersByIdHandler adapts the corresponding logic to gin router
func GetCustomersByIdHandler(c *gin.Context) {
	var queryHolder struct {
		Ids string `uri:"ids" binding:"required"`
	}

	err := c.ShouldBindUri(&queryHolder)
	if err != nil {
		log.Fatal(err)
	}

	ids := strings.Split(queryHolder.Ids, ",")

	customersAPI := BuildCustomers()
	items, err := customersAPI.GetById(ids)
	if err != nil {
		log.Fatal(err)
	}
	if len(items) == 0 {
		c.JSONP(http.StatusNotFound, items)
		return
	}
	c.JSONP(http.StatusOK, items)
}

//CreateCustomerHandler adapts the corresponding logic to gin router
func CreateCustomerHandler(c *gin.Context) {
	var customerInput api.CustomerConstructor
	if err := c.ShouldBindJSON(&customerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customersAPI := BuildCustomers()
	resp, err := customersAPI.Save(&customerInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSONP(http.StatusOK, resp)
}
