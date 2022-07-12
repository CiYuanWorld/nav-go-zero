package logic

import (
	"context"
	"nav-go-zero/app/services/user/rpc/user"

	"nav-go-zero/app/nav-web/api/internal/svc"
	"nav-go-zero/app/nav-web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoRes, err error) {
	uuid := req.UUID
	var userRes *user.UserResponse
	if uuid != "" {
		userRes, err = l.svcCtx.UserRpc.GetUserByUUID(l.ctx, &user.UUIDRequest{
			Uuid: uuid,
		})
		if err != nil {
			return nil, err
		}
	} else {
		userRes, err = l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
			Id: "1",
		})
		if err != nil {
			return nil, err
		}
	}

	resp = &types.UserInfoRes{
		UUID:     userRes.GetUsername(),
		Username: userRes.Username,
		Nickname: "",
	}
	return
}
