// Code generated by Kitex v0.4.4. DO NOT EDIT.

package tripservice

import (
	trip "github.com/CyanAsterisk/FreeCar/server/cmd/api/kitex_gen/trip"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler trip.TripService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}