package dept

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDeptLogic {
	return &DeleteSysDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDeptLogic) DeleteSysDept(req *types.DeleteSysDeptReq) error {
	deptList, _ := l.svcCtx.SysDeptModel.FindSubDept(l.ctx, req.Id)
	if len(deptList) != 0 {
		return errorx.NewDefaultError(errorx.DeleteDeptErrorCode)
	}

	userList, _ := l.svcCtx.SysUserModel.FindByCondition(l.ctx, "dept_id", req.Id)
	if len(userList) != 0 {
		return errorx.NewDefaultError(errorx.DeptHasUserErrorCode)
	}

	err := l.svcCtx.SysDeptModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}
