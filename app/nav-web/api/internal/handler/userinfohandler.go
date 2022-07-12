package handler

import (
	"net/http"

	"nav-go-zero/app/nav-web/api/internal/logic"
	"nav-go-zero/app/nav-web/api/internal/svc"
	"nav-go-zero/app/nav-web/api/internal/types"
	"nav-go-zero/app/pkg/errorx"
	"nav-go-zero/app/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewDefaultError("解析错误"))
			return
		}

		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		response.Response(w, resp, err)
	}
}
