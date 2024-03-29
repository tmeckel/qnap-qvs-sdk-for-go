package usersapi

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/qnap/qvs-sdk-for-go/services/users"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	ClearLog(ctx context.Context, ID int32) (result users.UpdateClearLogTimeResponse, err error)
	Create(ctx context.Context, body users.WriteUser) (result users.CreateResponse, err error)
	CreatePermissions(ctx context.Context, userID int32, vmid int32, body users.VMPermissionBase) (result users.CreateVMPermissionResponse, err error)
	DeletePermissions(ctx context.Context, userID int32, vmid int32) (result users.DeleteVMPermissionResponse, err error)
	Destroy(ctx context.Context, ID int32) (result users.DestroyResponse, err error)
	GetCurrentUserInfo(ctx context.Context) (result users.GetCurrentUserInfoResponse, err error)
	GetPermissions(ctx context.Context, userID int32, vmid int32) (result users.VMPermissionResponse, err error)
	List(ctx context.Context) (result users.ListResponse, err error)
	Retrieve(ctx context.Context, ID int32) (result users.RetrieveResponse, err error)
	Update(ctx context.Context, ID int32, body *users.WriteUser) (result users.UpdateResponse, err error)
	UpdatePermissions(ctx context.Context, userID int32, vmid int32, body *users.VMPermissionBase) (result users.UpdateVMPermissionResponse, err error)
}

var _ ClientAPI = (*users.Client)(nil)
