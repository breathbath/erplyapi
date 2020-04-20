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

/**
@api {get} /customers/:ids Customers list by ids
@apiDescription Lists customers by comma separated ids list
@apiName Customers list by ids
@apiGroup Customers

@apiExample {String} With many ids
		/customers/1,2,3
@apiExample {String} With one id
		/customers/1

@apiSuccessExample Success-Response
HTTP/1.1 200 OK
[
    {
        "id": 6,
        "customerID": 6,
        "type_id": "",
        "fullName": "Comp INC",
        "companyName": "Comp INC",
        "firstName": "",
        "lastName": "",
        "groupID": 14,
        "EDI": "",
        "phone": "",
        "eInvoiceEmail": "",
        "email": "",
        "fax": "",
        "code": "3333",
        "referenceNumber": "",
        "vatNumber": "",
        "bankName": "",
        "bankAccountNumber": "",
        "bankIBAN": "",
        "bankSWIFT": "",
        "paymentDays": 0,
        "notes": "",
        "lastModified": 1587311074,
        "customerType": "COMPANY",
        "address": "",
        "addresses": null,
        "street": "",
        "address2": "",
        "city": "",
        "postalCode": "",
        "country": "",
        "state": "",
        "contactPersons": []
    },
    {
        "id": 3,
        "customerID": 3,
        "type_id": "",
        "fullName": "mustermann, max",
        "companyName": "",
        "firstName": "max",
        "lastName": "mustermann",
        "groupID": 14,
        "EDI": "",
        "phone": "",
        "eInvoiceEmail": "",
        "email": "",
        "fax": "",
        "code": "",
        "referenceNumber": "",
        "vatNumber": "",
        "bankName": "",
        "bankAccountNumber": "",
        "bankIBAN": "",
        "bankSWIFT": "",
        "paymentDays": 0,
        "notes": "",
        "lastModified": 1587298463,
        "customerType": "PERSON",
        "address": "",
        "addresses": null,
        "street": "",
        "address2": "",
        "city": "",
        "postalCode": "",
        "country": "",
        "state": "",
        "contactPersons": []
    }
]

@apiErrorExample Not found(404)
HTTP/1.1 404 Not found
[]

@apiPermission registered user
*/
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

/**
@api {post} /customers Customers create
@apiDescription Creates a new customer
@apiName Customers create
@apiGroup Customers

@apiParamExample {json} Body:
{
	"CompanyName": "My Personal Inc",
	"Address":            "Elm str",
	"PostalCode":         "100234",
	"Country":            "USA",
	"FullName":           "Big Boss",
	"RegistryCode":      "1234",
	"VatNumber":          "23456",
	"Email":              "no@mail.me",
	"Phone":              "+13434134233134",
	"BankName":           "Best Bank",
	"BankAccountNumber":  "3434937493749813"
}

@apiSuccessExample Success-Response
HTTP/1.1 200 OK
{
    "clientID": 11,
    "customerID": 11
}

@apiErrorExample Bad request(400)
HTTP/1.1 400 Bad request
{
    "message": "ERPLY API: Can not save customer with empty name or registry number status: Error"
}

@apiPermission registered user
*/
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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, resp)
}
