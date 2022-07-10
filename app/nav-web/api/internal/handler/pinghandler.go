package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"navapp/app/nav-web/api/internal/logic"
	"navapp/app/nav-web/api/internal/svc"
	"navapp/app/nav-web/api/internal/types"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
