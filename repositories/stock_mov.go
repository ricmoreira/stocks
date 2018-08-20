package repositories

import (
	"context"
	"fmt"
	"stocks/models/request"
	"stocks/models/response"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/mongodb/mongo-go-driver/mongo/insertopt"
)

// StockMovRepository performs CRUD operations on users resource
type StockMovRepository struct {
	stockMov MongoCollection
}

// NewStockRepository is the constructor for StockRepository
func NewStockMovRepository(db *DBCollections) *StockMovRepository {
	return &StockMovRepository{stockMov: db.StockMov}
}

// CreateOne saves provided model instance to database
func (this *StockMovRepository) CreateOne(request *mrequest.StockMovCreate) (*mongo.InsertOneResult, error) {

	return this.stockMov.InsertOne(context.Background(), request)
}

// ReadOne returns a invoice based on StockMov ID sent in request
func (this *StockMovRepository) ReadOne(sm *mrequest.StockMovRead) (*mresponse.StockMovRead, error) {
	
	oID, err := objectid.FromHex(sm.ID)
	if err != nil {
		return nil, err
	}

	result := this.stockMov.FindOne(
		context.Background(),
		bson.NewDocument(bson.EC.ObjectID("_id", oID)),
	)

	res := mresponse.StockMovRead{}
	
	err = result.Decode(res)
	if err != nil {
		return nil, err
	}

	res.ID = res.IDdb.Hex()

	return &res, nil
}

// TODO: implement
func (this *StockMovRepository) UpdateOne(sm *mrequest.StockMovUpdate) (*mresponse.StockMovUpdate, error) {
	return nil, nil
}

// TODO: implement
func (this *StockMovRepository) DeleteOne(sm *mrequest.StockMovDelete) (*mresponse.StockMovDelete, error) {
	return nil, nil
}

func (this *StockMovRepository) InsertMany(request *[]*mrequest.StockMovCreate) (*mongo.InsertManyResult, error) {
	// transform to []interface{} (https://golang.org/doc/faq#convert_slice_of_interface)
	s := make([]interface{}, len(*request))
	for i, v := range *request {
		s[i] = v
	}

	// { ordered: false } ordered is false in order to don't stop execution because an error ocurred on one of the inserts
	opt := insertopt.Ordered(false)
	return this.stockMov.InsertMany(context.Background(), s, opt)
}

func (this *StockMovRepository) List(req *mrequest.ListRequest) (int64, int64, int64, mongo.Cursor, error) {

	args := []*bson.Element{}

	for i, v := range req.Filters {
		args = append(args, bson.EC.String(i, fmt.Sprintf("%v", v)))
	}

	total, e := this.stockMov.Count(
		context.Background(),
		bson.NewDocument(args...),
	)

	sorting := map[string]int{}
	var sortingValue int
	if req.Order == "reverse" {
		sortingValue = -1
	} else {
		sortingValue = 1
	}
	sorting[req.Sort] = sortingValue

	perPage := int64(req.PerPage)
	page := int64(req.Page)
	cursor, e := this.stockMov.Find(
		context.Background(),
		bson.NewDocument(args...),
		findopt.Sort(sorting),
		findopt.Skip(int64(req.PerPage*(req.Page-1))),
		findopt.Limit(perPage),
	)

	return total, perPage, page, cursor, e
}
