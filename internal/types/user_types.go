package types

import (
	"time"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
)

var _ time.Time

// Tip: suggested filling in the binding rules https://github.com/go-playground/validator in request struct fields tag.

// CreateUserRequest request params
type CreateUserRequest struct {
	TwId          string    `json:"twId" binding:""`      // 推特id
	TwName        string    `json:"twName" binding:""`    // 推特名
	AvatarUrl     string    `json:"avatarUrl" binding:""` // 头像
	Token         string    `json:"token" binding:""`
	TokenSecret   string    `json:"tokenSecret" binding:""`
	WalletAddress string    `json:"walletAddress" binding:""` // 钱包地址
	UpdatedAt     time.Time `json:"updatedAt" binding:""`
	CreatedAt     time.Time `json:"createdAt" binding:""`
	DeletedAt     time.Time `json:"deletedAt" binding:""`
	LikeCount     int       `json:"likeCount" binding:""` // 点赞数量
}

// UpdateUserByIDRequest request params
type UpdateUserByIDRequest struct {
	ID uint64 `json:"id" binding:""` // uint64 id

	TwId          string    `json:"twId" binding:""`      // 推特id
	TwName        string    `json:"twName" binding:""`    // 推特名
	AvatarUrl     string    `json:"avatarUrl" binding:""` // 头像
	Token         string    `json:"token" binding:""`
	TokenSecret   string    `json:"tokenSecret" binding:""`
	WalletAddress string    `json:"walletAddress" binding:""` // 钱包地址
	UpdatedAt     time.Time `json:"updatedAt" binding:""`
	CreatedAt     time.Time `json:"createdAt" binding:""`
	DeletedAt     time.Time `json:"deletedAt" binding:""`
	LikeCount     int       `json:"likeCount" binding:""` // 点赞数量
}

// UserObjDetail detail
type UserObjDetail struct {
	ID string `json:"id"` // convert to string id

	TwId          string    `json:"twId"`      // 推特id
	TwName        string    `json:"twName"`    // 推特名
	AvatarUrl     string    `json:"avatarUrl"` // 头像
	Token         string    `json:"token"`
	TokenSecret   string    `json:"tokenSecret"`
	WalletAddress string    `json:"walletAddress"` // 钱包地址
	UpdatedAt     time.Time `json:"updatedAt"`
	CreatedAt     time.Time `json:"createdAt"`
	DeletedAt     time.Time `json:"deletedAt"`
	LikeCount     int       `json:"likeCount"` // 点赞数量
}

// CreateUserRespond only for api docs
type CreateUserRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		ID uint64 `json:"id"` // id
	} `json:"data"` // return data
}

// UpdateUserByIDRespond only for api docs
type UpdateUserByIDRespond struct {
	Result
}

// GetUserByIDRespond only for api docs
type GetUserByIDRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		User UserObjDetail `json:"user"`
	} `json:"data"` // return data
}

// DeleteUserByIDRespond only for api docs
type DeleteUserByIDRespond struct {
	Result
}

// DeleteUsersByIDsRequest request params
type DeleteUsersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// DeleteUsersByIDsRespond only for api docs
type DeleteUsersByIDsRespond struct {
	Result
}

// GetUserByConditionRequest request params
type GetUserByConditionRequest struct {
	query.Conditions
}

// GetUserByConditionRespond only for api docs
type GetUserByConditionRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		User UserObjDetail `json:"user"`
	} `json:"data"` // return data
}

// ListUsersByIDsRequest request params
type ListUsersByIDsRequest struct {
	IDs []uint64 `json:"ids" binding:"min=1"` // id list
}

// ListUsersByIDsRespond only for api docs
type ListUsersByIDsRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users []UserObjDetail `json:"users"`
	} `json:"data"` // return data
}

// ListUsersRequest request params
type ListUsersRequest struct {
	query.Params
}

// ListUsersRespond only for api docs
type ListUsersRespond struct {
	Code int    `json:"code"` // return code
	Msg  string `json:"msg"`  // return information description
	Data struct {
		Users []UserObjDetail `json:"users"`
	} `json:"data"` // return data
}
