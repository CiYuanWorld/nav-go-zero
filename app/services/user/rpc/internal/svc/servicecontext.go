package svc

import (
	"nav-go-zero/app/services/user/model"
	"nav-go-zero/app/services/user/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	TUserModel model.TUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		TUserModel: model.NewTUserModel(conn, c.CacheRedis),
	}
}
