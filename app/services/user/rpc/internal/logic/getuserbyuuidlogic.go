package logic

import (
	"context"
	"errors"

	"nav-go-zero/app/services/user/model"
	"nav-go-zero/app/services/user/rpc/internal/svc"
	"nav-go-zero/app/services/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByUUIDLogic {
	return &GetUserByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByUUIDLogic) GetUserByUUID(in *user.UUIDRequest) (*user.UserResponse, error) {
	userInfo, err := l.svcCtx.TUserModel.FindOneByUuid(l.ctx, in.Uuid)
	logx.Info(userInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.New("用户名不存在")
		}
		return nil, err
	}

	return &user.UserResponse{
		Id:       userInfo.Uuid,
		Username: userInfo.Username,
		Gender:   "male",
	}, nil
}
