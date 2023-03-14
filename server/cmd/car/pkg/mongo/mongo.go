package mongo

import (
	"context"
	"fmt"

	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/id"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/car"
	mgutil "github.com/CyanAsterisk/FreeCar/server/shared/mongo"
	"github.com/CyanAsterisk/FreeCar/server/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	carField      = "car"
	statusField   = carField + ".status"
	driverField   = carField + ".driver"
	positionField = carField + ".position"
	tripIDField   = carField + ".tripid"
	powerField    = carField + ".power"
	initLatitude  = 29.5
	initLongitude = 106.6
)

// CarRecord defines a car record in mongo db.
type CarRecord struct {
	mgutil.IDField `bson:"inline"`
	Car            *car.Car `bson:"car"`
}

type Manager struct {
	col *mongo.Collection
}

// NewManager creates a mongo manager.
func NewManager(db *mongo.Database) *Manager {
	return &Manager{col: db.Collection(consts.CarCollection)}
}

// CreateCar creates a car.
func (m *Manager) CreateCar(c context.Context, plateNum string) (*CarRecord, error) {
	cr := &CarRecord{
		Car: &car.Car{
			Position: &car.Location{
				Latitude:  initLatitude,
				Longitude: initLongitude,
			},
			Status:   car.CarStatus_LOCKED,
			PlateNum: plateNum,
			Power:    100,
		},
	}
	cr.ID = mgutil.NewObjID()
	_, err := m.col.InsertOne(c, cr)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

// GetCar gets a car.
func (m *Manager) GetCar(c context.Context, id id.CarID) (*CarRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	return convertSingleResult(m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objID,
	}))
}

// GetCars gets cars.
func (m *Manager) GetCars(c context.Context) ([]*CarRecord, error) {
	filter := bson.M{}
	res, err := m.col.Find(c, filter, options.Find())
	if err != nil {
		return nil, err
	}

	var cars []*CarRecord
	for res.Next(c) {
		var cr CarRecord
		err := res.Decode(&cr)
		if err != nil {
			return nil, err
		}
		cars = append(cars, &cr)
	}
	return cars, nil
}

// CarUpdate defines updates to a car.
// Only specified fields will be updated.
type CarUpdate struct {
	Status       car.CarStatus
	Position     *car.Location
	Driver       *car.Driver
	Power        float64
	UpdateTripID bool
	TripID       id.TripID
}

// UpdateCar updates a car. If status is specified,
// it updates the car only when existing record matches the
// status specified.
func (m *Manager) UpdateCar(c context.Context, id id.CarID, status car.CarStatus, update *CarUpdate) (*CarRecord, error) {
	objID, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	filter := bson.M{
		mgutil.IDFieldName: objID,
	}
	if status != car.CarStatus_CS_NOT_SPECIFIED {
		filter[statusField] = status
	}

	u := bson.M{}
	if update.Status != car.CarStatus_CS_NOT_SPECIFIED {
		u[statusField] = update.Status
	}
	if update.Driver != nil {
		u[driverField] = update.Driver
	}
	if update.Position != nil {
		u[positionField] = update.Position
	}
	if update.UpdateTripID {
		u[tripIDField] = update.TripID.String()
	}
	if update.Power > 0 {
		u[powerField] = update.Power
	}
	res := m.col.FindOneAndUpdate(c, filter, mgutil.Set(u),
		options.FindOneAndUpdate().SetReturnDocument(options.After))

	return convertSingleResult(res)
}

func convertSingleResult(res *mongo.SingleResult) (*CarRecord, error) {
	if err := res.Err(); err != nil {
		return nil, err
	}
	var cr CarRecord
	err := res.Decode(&cr)
	if err != nil {
		return nil, fmt.Errorf("cannot decode: %v", err)
	}
	return &cr, nil
}