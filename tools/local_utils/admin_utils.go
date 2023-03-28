package local_utils

import (
	"context"
)

const (
	RoleIdKey             = "roleId"
	DeptIdKey             = "deptId"
	DataAccessTypeKey     = "dataAccessType"
	AdminLoginUsernameKey = "adminLoginUserName"
	OrgNoKey              = "orgNo"
)

//TODO 找一个util包放
func GetRoleId(ctx context.Context) int64 {
	roleId, ok := ctx.Value(RoleIdKey).(int64)
	if !ok {
		return 0
	}
	return roleId
}

func GetDeptId(ctx context.Context) int64 {
	deptId, ok := ctx.Value(DeptIdKey).(int64)
	if !ok {
		return 0
	}
	return deptId
}

func GetDataAccessType(ctx context.Context) int64 {
	dataAccessType, ok := ctx.Value(DataAccessTypeKey).(int64)
	if !ok {
		return 0
	}
	return dataAccessType
}
func GetOrgNo(ctx context.Context) string {
	orgNo, ok := ctx.Value(OrgNoKey).(string)
	if !ok {
		return ""
	}
	return orgNo
}
