package services

import (
	"context"
	"stocks/models"
	"stocks/models/request"
	"stocks/models/response"
	"stocks/repositories"
	"stocks/util"
	"stocks/util/errors"

	"log"

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// StockMovServiceContract is the abstraction for service layer on roles resource
type StockMovServiceContract interface {
	CreateOne(*mrequest.StockMovCreate) (*mresponse.StockMovCreate, *mresponse.ErrorResponse)
	ReadOne(*mrequest.StockMovRead) (*mresponse.StockMovRead, *mresponse.ErrorResponse)
	UpdateOne(*mrequest.StockMovUpdate) (*mresponse.StockMovUpdate, *mresponse.ErrorResponse)
	DeleteOne(*mrequest.StockMovDelete) (*mresponse.StockMovDelete, *mresponse.ErrorResponse)
	CreateMany(*[]*mrequest.StockMovCreate) (*[]*mresponse.StockMovCreate, *mresponse.ErrorResponse)
	List(request *mrequest.ListRequest) (*mresponse.StockMovList, *mresponse.ErrorResponse)
	CreateStockMovementsFromInvoice(invoice *models.Invoice) (*[]*mresponse.StockMovCreate, *mresponse.ErrorResponse)
	ListStockMovCount(request *mrequest.ListRequest) (*mresponse.StockMovCountList, *mresponse.ErrorResponse) 
}

// StockMovService is the layer between http client and repository for Stock Mov resource
type StockMovService struct {
	StockMovRepository *repositories.StockMovRepository
}

// NewStockMovService is the constructor of StockMovService
func NewStockMovService(smr *repositories.StockMovRepository) *StockMovService {
	return &StockMovService{
		StockMovRepository: smr,
	}
}

// CreateOne saves provided model instance to database
func (this *StockMovService) CreateOne(request *mrequest.StockMovCreate) (*mresponse.StockMovCreate, *mresponse.ErrorResponse) {

	// validate request
	e := errors.ValidateRequest(request)
	if e != nil {
		return nil, e
	}

	res, err := this.StockMovRepository.CreateOne(request)

	if err != nil {
		errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, errR
	}

	id := res.InsertedID.(objectid.ObjectID)
	sm := mresponse.StockMovCreate{
		ID: id.Hex(),
	}

	return &sm, nil
}

// TODO: implement
func (this *StockMovService) ReadOne(sm *mrequest.StockMovRead) (*mresponse.StockMovRead, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (this *StockMovService) UpdateOne(sm *mrequest.StockMovUpdate) (*mresponse.StockMovUpdate, *mresponse.ErrorResponse) {
	return nil, nil
}

// TODO: implement
func (this *StockMovService) DeleteOne(sm *mrequest.StockMovDelete) (*mresponse.StockMovDelete, *mresponse.ErrorResponse) {
	return nil, nil
}

// CreateMany saves many Stocks in one bulk operation
func (this *StockMovService) CreateMany(request *[]*mrequest.StockMovCreate) (*[]*mresponse.StockMovCreate, *mresponse.ErrorResponse) {

	if len(*request) == 0 {
		resp := make([]*mresponse.StockMovCreate, 0)
		return &resp, nil
	}

	res, err := this.StockMovRepository.InsertMany(request)

	if err != nil {
		mngBulkError := err.(mongo.BulkWriteError)
		writeErrors := mngBulkError.WriteErrors
		for _, err := range writeErrors {
			log.Println(err)
		}
	}

	result := make([]*mresponse.StockMovCreate, len(res.InsertedIDs))
	for i, insertedID := range res.InsertedIDs {
		id := insertedID.(objectid.ObjectID)
		result[i] = &mresponse.StockMovCreate{
			ID: id.Hex(),
		}
	}

	return &result, nil
}

// List returns a list of Stock Mov with pagination and filtering options
func (this *StockMovService) List(request *mrequest.ListRequest) (*mresponse.StockMovList, *mresponse.ErrorResponse) {

	total, perPage, page, cursor, err := this.StockMovRepository.List(request)

	if err != nil {
		e := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, e
	}

	docs := []*mresponse.StockMovRead{}

	for cursor.Next(context.Background()) {
		doc := mresponse.StockMovRead{}
		if err := cursor.Decode(&doc); err != nil {
			errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
			return nil, errR
		}

		doc.ID = doc.IDdb.Hex()

		docs = append(docs, &doc)
	}

	resp := mresponse.StockMovList{
		Total:   total,
		PerPage: perPage,
		Page:    page,
		Items:   &docs,
	}
	return &resp, nil
}

// ListStockMovCount returns a list with the stock count for all products
func (this *StockMovService) ListStockMovCount(request *mrequest.ListRequest) (*mresponse.StockMovCountList, *mresponse.ErrorResponse) {

	total, perPage, page, cursor, err := this.StockMovRepository.ListStockMovCount(request)

	if err != nil {
		e := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
		return nil, e
	}

	docs := []*mresponse.StockMovCount{}

	for cursor.Next(context.Background()) {
		doc := mresponse.StockMovCount{}
		if err := cursor.Decode(&doc); err != nil {
			errR := errors.HandleErrorResponse(errors.SERVICE_UNAVAILABLE, nil, err.Error())
			return nil, errR
		}

		docs = append(docs, &doc)
	}

	resp := mresponse.StockMovCountList{
		Total:   total,
		PerPage: perPage,
		Page:    page,
		Items:   &docs,
	}
	return &resp, nil
}

// CreateStockMovementFromInvoice creates stock movements of type SALE from an invoices
func (this *StockMovService) CreateStockMovementsFromInvoice(invoice *models.Invoice) (*[]*mresponse.StockMovCreate, *mresponse.ErrorResponse) {

	stockMovements := make([]*mrequest.StockMovCreate, 0)

	docID := invoice.InvoiceNo
	time, _ := util.ParseDateTime(invoice.SystemEntryDate)
	// ignore error on parsing time
	for i, line := range invoice.Lines {
		mov := mrequest.StockMovCreate{}
		mov.DocumentID = docID
		mov.MovementType = models.SALE
		mov.Line = int32(i)
		mov.Quantity = line.Quantity
		mov.ProductCode = line.ProductCode
		mov.UnitOfMeasure = line.UnitOfMeasure
		mov.Time = time
		mov.Dir = "OUT"
		mov.WharehouseID = "1" // for now, no logic of wharehouse implemented. default is "1" TODO: implement wharehouse logic

		stockMovements = append(stockMovements, &mov)
	}

	return this.CreateMany(&stockMovements)
}
