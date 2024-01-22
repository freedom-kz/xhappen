package server

import (
	v1 "xhappen/api/portal/v1"
	"xhappen/app/portal/internal/conf"
	"xhappen/app/portal/internal/service"
	"xhappen/pkg/filter/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, user *service.UserService, logger log.Logger) *http.Server {

	keyFunc := func(token *jwtv4.Token) (interface{}, error) {
		return []byte(c.Auth.Jwt.Secret), nil
	}

	jwtOption := jwt.WithClaims(func() jwtv4.Claims { return &jwtv4.RegisteredClaims{} })

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			logging.Server(logger),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
			),
			jwt.Server(keyFunc, user.VerifyToken, jwtOption),
		),
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUserHTTPServer(srv, user)

	route := srv.Route("/")
	route.POST("/auth/upload", service.UploadFile)

	return srv
}
