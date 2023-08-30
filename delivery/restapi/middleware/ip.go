package middleware

import (
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi/response"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/infrastructures/config"
	"net/http"
)

func IPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RemoteAddr != "ip" && config.AppStatus == "production" {
			response.NewError(w, r, model.ErrForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
