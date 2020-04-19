package erply

import (
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

type apiClientMock struct {
	PostCustomerIn             *api.CustomerConstructor
	PostCustomerReport         *api.CustomerImportReport
	GetCustomersByIDsCustomers api.Customers
	GetCustomersByIDsIDs       []string
}

func (acm apiClientMock) GetConfParameters() (*api.ConfParameter, error) {
	return nil, nil
}

func (acm apiClientMock) GetWarehouses() (w api.Warehouses, e error) {
	return
}

func (acm apiClientMock) GetUserName() (s string, e error) {
	return
}

func (acm apiClientMock) GetSalesDocumentByID(id string) (sd []api.SaleDocument, e error) {
	return
}

func (acm apiClientMock) GetSalesDocumentsByIDs(id []string) (sd []api.SaleDocument, e error) {
	return
}

func (acm *apiClientMock) GetCustomersByIDs(customerID []string) (c api.Customers, e error) {
	acm.GetCustomersByIDsIDs = customerID
	return acm.GetCustomersByIDsCustomers, nil
}

func (acm apiClientMock) GetCustomerByRegNumber(regNumber string) (c *api.Customer, e error) {
	return
}

func (acm apiClientMock) GetCustomerByGLN(gln string) (c *api.Customer, e error) {
	return
}

func (acm apiClientMock) GetSupplierByName(name string) (c *api.Customer, e error) {
	return
}

func (acm apiClientMock) GetVatRatesByID(vatRateID string) (vr api.VatRates, e error) {
	return
}

func (acm apiClientMock) GetCompanyInfo() (ci *api.CompanyInfo, e error) {
	return
}

func (acm apiClientMock) GetProductUnits() (pu []api.ProductUnit, e error) {
	return
}

func (acm apiClientMock) GetProductsByIDs(ids []string) (p []api.Product, e error) {
	return
}

func (acm apiClientMock) GetProductsByCode3(code3 string) (p *api.Product, e error) {
	return
}

func (acm apiClientMock) GetAddresses() (a *api.Address, e error) {
	return
}

func (acm apiClientMock) PostPurchaseDocument(in *api.PurchaseDocumentConstructor, provider string) (pdir api.PurchaseDocImportReports, e error) {
	return
}

func (acm apiClientMock) PostSalesDocumentFromWoocomm(in *api.SaleDocumentConstructor, shopOrderID string) (sdir api.SaleDocImportReports, e error) {
	return
}

func (acm apiClientMock) PostSalesDocument(in *api.SaleDocumentConstructor, provider string) (sdir api.SaleDocImportReports, e error) {
	return
}

func (acm *apiClientMock) PostCustomer(in *api.CustomerConstructor) (cir *api.CustomerImportReport, e error) {
	acm.PostCustomerIn = in
	return acm.PostCustomerReport, nil
}

func (acm apiClientMock) PostSupplier(in *api.CustomerConstructor) (cir *api.CustomerImportReport, e error) {
	return
}

func (acm apiClientMock) DeleteDocumentsByID(id string) error {
	return nil
}

func (acm apiClientMock) GetPointsOfSaleByID(posID string) (ps *api.PointOfSale, e error) {
	return
}

func (acm apiClientMock) VerifyIdentityToken(jwt string) (si *api.SessionInfo, e error) {
	return
}

func (acm apiClientMock) GetIdentityToken() (it *api.IdentityToken, e error) {
	return
}

func (acm apiClientMock) Close() {

}

type clientFactoryMock struct {
	apiClientMock *apiClientMock
}

func (cfm clientFactoryMock) CreateClient() (cl api.IClient, err error) {
	return cfm.apiClientMock, nil
}

func TestGetCustomersById(t *testing.T) {
	cacheMock := &cacheMock{
		ReadFound:        false,
		ReadTargetToGive: []byte(`{}`),
	}
	apiClientMock := &apiClientMock{
		GetCustomersByIDsCustomers: api.Customers{
			api.Customer{
				ID:         123,
				CustomerID: 123,
				FullName:   "Max Mustermann",
			},
		},
	}

	factoryMock := clientFactoryMock{
		apiClientMock: apiClientMock,
	}
	customersAPI := Customers{
		ClientFactory: factoryMock,
		Cache:         cacheMock,
	}

	customers, err := customersAPI.GetById([]string{"123", "345"})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	assert.Len(t, customers, 1)
	if len(customers) != 1 {
		return
	}

	customer := customers[0]
	assert.Equal(t, 123, customer.ID)
	assert.Equal(t, 123, customer.CustomerID)
	assert.Equal(t, "Max Mustermann", customer.FullName)
	assert.Equal(t, []string{"123", "345"}, apiClientMock.GetCustomersByIDsIDs)
}

func TestSaveCustomer(t *testing.T) {
	cacheMock := &cacheMock{
		ReadFound:        false,
		ReadTargetToGive: []byte(`{}`),
	}
	apiClientMock := &apiClientMock{
		PostCustomerReport: &api.CustomerImportReport{
			ClientID: 12,
			CustomerID: 12,
		},
	}

	factoryMock := clientFactoryMock{
		apiClientMock: apiClientMock,
	}
	customersAPI := Customers{
		ClientFactory: factoryMock,
		Cache:         cacheMock,
	}

	resp, err := customersAPI.Save(&api.CustomerConstructor{
		CompanyName: "Some Company",
		Address: "Some address",
	})
	assert.NoError(t, err)
	if err != nil {
		return
	}

	customerGiven := apiClientMock.PostCustomerIn
	assert.NotNil(t, customerGiven)
	if customerGiven == nil {
		return
	}

	assert.Equal(t, "Some Company", customerGiven.CompanyName)
	assert.Equal(t, "Some address", customerGiven.Address)
	assert.Equal(t, 12, resp.CustomerID)
	assert.Equal(t, 12, resp.ClientID)
}
