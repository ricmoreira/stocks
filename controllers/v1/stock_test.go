package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"stocks/models"
	"stocks/models/request"
	"stocks/models/response"
	"stocks/util/errors"
	"testing"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// stub StockService behaviour
type MockStockMovService struct{}

// mocked behaviour for CreateOne
func (ps *MockStockMovService) CreateOne(pReq *mrequest.StockMovCreate) (*mresponse.StockMovCreate, *mresponse.ErrorResponse) {
	// validate request
	err := errors.ValidateRequest(pReq)
	if err != nil {
		return nil, err
	}

	pRes := mresponse.StockMovCreate{}
	pRes.ID = "some-unique-id"

	return &pRes, nil
}

// mocked behaviour for ReadOne
func (ps *MockStockMovService) ReadOne(p *mrequest.StockMovRead) (*mresponse.StockMovRead, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for UpdateOne
func (ps *MockStockMovService) UpdateOne(p *mrequest.StockMovUpdate) (*mresponse.StockMovUpdate, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

// mocked behaviour for DeleteOne
func (ps *MockStockMovService) DeleteOne(p *mrequest.StockMovDelete) (*mresponse.StockMovDelete, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func (ps *MockStockMovService) CreateMany(*[]*mrequest.StockMovCreate) (*[]*mresponse.StockMovCreate, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func (ps *MockStockMovService) List(*mrequest.ListRequest) (*mresponse.StockMovList, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func (ps *MockStockMovService) CreateStockMovementsFromInvoice(invoice *models.Invoice) (*[]*mresponse.StockMovCreate, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}
func (ps *MockStockMovService) ListStockMovCount(request *mrequest.ListRequest) (*mresponse.StockMovCountList, *mresponse.ErrorResponse) {
	// TODO: implement in the future
	return nil, nil
}

func TestCreateStockAction(t *testing.T) {

	// Mock the server

	// Switch to test mode in order to don't get such noisy output
	gin.SetMode(gin.TestMode)

	sms := &MockStockMovService{}

	pc := StockMovController{
		StockMovService: sms,
	}

	r := gin.Default()

	r.POST("/api/v1/stock/movement", pc.CreateAction)

	// TEST SUCCESS

	// Mock a request
	body := mrequest.StockMovCreate{}
	body.MovementType = "SALE"
	body.DocumentID = "FS1235"
	body.Quantity = 1.0
	body.ProductCode = "123456789"
	body.UnitOfMeasure = "UNI"

	jsonValue, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/stock/movement", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder in order to inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Do asssertions
	if w.Code != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		bodyString := string(bodyBytes)

		t.Fatalf("Expected to get status %d but instead got %d\nResponse body:\n%s", http.StatusOK, w.Code, bodyString)
	}
}
