package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TLinkModel = (*customTLinkModel)(nil)

type (
	// TLinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTLinkModel.
	TLinkModel interface {
		tLinkModel
	}

	customTLinkModel struct {
		*defaultTLinkModel
	}
)

// NewTLinkModel returns a model for the database table.
func NewTLinkModel(conn sqlx.SqlConn, c cache.CacheConf) TLinkModel {
	return &customTLinkModel{
		defaultTLinkModel: newTLinkModel(conn, c),
	}
}
