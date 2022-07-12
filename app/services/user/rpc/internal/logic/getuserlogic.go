package logic

import (
	"context"
	"errors"

	"nav-go-zero/app/services/user/model"
	"nav-go-zero/app/services/user/rpc/internal/svc"
	"nav-go-zero/app/services/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	// userInfo, err := l.svcCtx.TUserModel.FindOne().FindOneByUsername(l.ctx, "wu2kong")
	userInfo, err := l.svcCtx.TUserModel.FindOne(l.ctx, 1)
	logx.Info(userInfo)
	if err != nil {
		logx.Info(err)
		if err == model.ErrNotFound {
			return nil, errors.New("用户名不存在")
		}
		return nil, err
	}

	logx.Info("wu2kong: err")

	return &user.UserResponse{
		Id:       userInfo.Uuid,
		Username: userInfo.Username,
		Gender:   "male",
	}, nil
}
