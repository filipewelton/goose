package hooks

import (
	"net/http"
	httpLib "retail_flow/internal/drivers/http/lib"
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"
	"strings"
)

func AuthenticateAPIAccess(next http.Handler) http.Handler {
	var guardian lib.APIGuardian

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			httpLib.RenderErrorResponse(
				exceptions.UnauthenticatedExceptions.ErrUnauthenticatedFrontend,
				w,
				r,
			)

			return
		}

		token := strings.ReplaceAll(authorization, "Bearer ", "")

		if err := guardian.Validate(token); err != nil {
			httpLib.RenderErrorResponse(err, w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
