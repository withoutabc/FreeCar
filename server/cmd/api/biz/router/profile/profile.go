// Code generated by hertz generator. DO NOT EDIT.

package Profile

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	profile "github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/handler/profile"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_v1 := root.Group("/v1", _v1Mw()...)
		_v1.GET("/profile", append(_profileMw(), profile.GetProfile)...)
		_profile := _v1.Group("/profile", _profileMw()...)
		_profile.POST("/photo", append(_createprofilephotoMw(), profile.CreateProfilePhoto)...)
		_profile.DELETE("/photo", append(_clearprofilephotoMw(), profile.ClearProfilePhoto)...)
		_profile.GET("/photo", append(_photoMw(), profile.GetProfilePhoto)...)
		_photo := _profile.Group("/photo", _photoMw()...)
		_photo.POST("/complete", append(_completeprofilephotoMw(), profile.CompleteProfilePhoto)...)
		_v1.POST("/profile", append(_submitprofileMw(), profile.SubmitProfile)...)
		_v1.DELETE("/profile", append(_clearprofileMw(), profile.ClearProfile)...)
	}
}