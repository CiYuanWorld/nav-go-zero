package logic

import (
	"context"

	"nav-go-zero/app/nav-web/api/internal/svc"
	"nav-go-zero/app/nav-web/api/internal/types"

	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.PingReq) (resp *types.PingRes, err error) {
	return &types.PingRes{
		Message: "pong",
		Time:    carbon.Now().ToDateTimeString(),
	}, nil
}
