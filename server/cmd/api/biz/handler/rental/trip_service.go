// Code generated by hertz generator.

package rental

import (
	"context"

	rental "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/model/rental"
	"github.com/cloudwego/hertz/pkg/app"
)

// CreateTrip .
// @router /v1/trip [POST]
func CreateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rental.CreateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(rental.TripEntity)

	c.JSON(200, resp)
}

// GetTrip .
// @router /v1/trip/:id [GET]
func GetTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rental.GetTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(rental.Trip)

	c.JSON(200, resp)
}

// GetTrips .
// @router /v1/trips [GET]
func GetTrips(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rental.GetTripsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(rental.GetTripsResponse)

	c.JSON(200, resp)
}

// UpdateTrip .
// @router /v1/trip/:id [PUT]
func UpdateTrip(ctx context.Context, c *app.RequestContext) {
	var err error
	var req rental.UpdateTripRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(rental.Trip)

	c.JSON(200, resp)
}
