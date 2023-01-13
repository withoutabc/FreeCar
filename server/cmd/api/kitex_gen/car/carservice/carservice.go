// Code generated by Kitex v0.4.4. DO NOT EDIT.

package carservice

import (
	"context"
	car "github.com/CyanAsterisk/FreeCar/server/cmd/api/kitex_gen/car"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return carServiceServiceInfo
}

var carServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CarService"
	handlerType := (*car.CarService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateCar": kitex.NewMethodInfo(createCarHandler, newCarServiceCreateCarArgs, newCarServiceCreateCarResult, false),
		"GetCar":    kitex.NewMethodInfo(getCarHandler, newCarServiceGetCarArgs, newCarServiceGetCarResult, false),
		"GetCars":   kitex.NewMethodInfo(getCarsHandler, newCarServiceGetCarsArgs, newCarServiceGetCarsResult, false),
		"LockCar":   kitex.NewMethodInfo(lockCarHandler, newCarServiceLockCarArgs, newCarServiceLockCarResult, false),
		"UnlockCar": kitex.NewMethodInfo(unlockCarHandler, newCarServiceUnlockCarArgs, newCarServiceUnlockCarResult, false),
		"UpdateCar": kitex.NewMethodInfo(updateCarHandler, newCarServiceUpdateCarArgs, newCarServiceUpdateCarResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "car",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func createCarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*car.CarServiceCreateCarArgs)
	realResult := result.(*car.CarServiceCreateCarResult)
	success, err := handler.(car.CarService).CreateCar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCarServiceCreateCarArgs() interface{} {
	return car.NewCarServiceCreateCarArgs()
}

func newCarServiceCreateCarResult() interface{} {
	return car.NewCarServiceCreateCarResult()
}

func getCarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*car.CarServiceGetCarArgs)
	realResult := result.(*car.CarServiceGetCarResult)
	success, err := handler.(car.CarService).GetCar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCarServiceGetCarArgs() interface{} {
	return car.NewCarServiceGetCarArgs()
}

func newCarServiceGetCarResult() interface{} {
	return car.NewCarServiceGetCarResult()
}

func getCarsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*car.CarServiceGetCarsArgs)
	realResult := result.(*car.CarServiceGetCarsResult)
	success, err := handler.(car.CarService).GetCars(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCarServiceGetCarsArgs() interface{} {
	return car.NewCarServiceGetCarsArgs()
}

func newCarServiceGetCarsResult() interface{} {
	return car.NewCarServiceGetCarsResult()
}

func lockCarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*car.CarServiceLockCarArgs)
	realResult := result.(*car.CarServiceLockCarResult)
	success, err := handler.(car.CarService).LockCar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCarServiceLockCarArgs() interface{} {
	return car.NewCarServiceLockCarArgs()
}

func newCarServiceLockCarResult() interface{} {
	return car.NewCarServiceLockCarResult()
}

func unlockCarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*car.CarServiceUnlockCarArgs)
	realResult := result.(*car.CarServiceUnlockCarResult)
	success, err := handler.(car.CarService).UnlockCar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCarServiceUnlockCarArgs() interface{} {
	return car.NewCarServiceUnlockCarArgs()
}

func newCarServiceUnlockCarResult() interface{} {
	return car.NewCarServiceUnlockCarResult()
}

func updateCarHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*car.CarServiceUpdateCarArgs)
	realResult := result.(*car.CarServiceUpdateCarResult)
	success, err := handler.(car.CarService).UpdateCar(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCarServiceUpdateCarArgs() interface{} {
	return car.NewCarServiceUpdateCarArgs()
}

func newCarServiceUpdateCarResult() interface{} {
	return car.NewCarServiceUpdateCarResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateCar(ctx context.Context, req *car.CreateCarRequest) (r *car.CarEntity, err error) {
	var _args car.CarServiceCreateCarArgs
	_args.Req = req
	var _result car.CarServiceCreateCarResult
	if err = p.c.Call(ctx, "CreateCar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCar(ctx context.Context, req *car.GetCarRequest) (r *car.Car, err error) {
	var _args car.CarServiceGetCarArgs
	_args.Req = req
	var _result car.CarServiceGetCarResult
	if err = p.c.Call(ctx, "GetCar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCars(ctx context.Context, req *car.GetCarsRequest) (r *car.GetCarsResponse, err error) {
	var _args car.CarServiceGetCarsArgs
	_args.Req = req
	var _result car.CarServiceGetCarsResult
	if err = p.c.Call(ctx, "GetCars", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) LockCar(ctx context.Context, req *car.LockCarRequest) (r *car.LockCarResponse, err error) {
	var _args car.CarServiceLockCarArgs
	_args.Req = req
	var _result car.CarServiceLockCarResult
	if err = p.c.Call(ctx, "LockCar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UnlockCar(ctx context.Context, req *car.UnlockCarRequest) (r *car.UnlockCarResponse, err error) {
	var _args car.CarServiceUnlockCarArgs
	_args.Req = req
	var _result car.CarServiceUnlockCarResult
	if err = p.c.Call(ctx, "UnlockCar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateCar(ctx context.Context, req *car.UpdateCarRequest) (r *car.UpdateCarResponse, err error) {
	var _args car.CarServiceUpdateCarArgs
	_args.Req = req
	var _result car.CarServiceUpdateCarResult
	if err = p.c.Call(ctx, "UpdateCar", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
