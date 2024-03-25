package jwt

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
)

const (
	// bearerWord the bearer key word for authorization
	BearerWord string = "Bearer"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	AuthorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"

	VERIFY_PATH string = "/auth"

	AUTHEDKEY = "UID"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod jwt.SigningMethod
	claims        func() jwt.Claims
}

// WithSigningMethod with signing method option.
func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

// WithClaims with customer claim
// If you use it in Server, f needs to return a new jwt.Claims object each time to avoid concurrent write problems
// If you use it in Client, f only needs to return a single object to provide performance
func WithClaims(f func() jwt.Claims) Option {
	return func(o *options) {
		o.claims = f
	}
}

func Server(keyFunc jwt.Keyfunc, verifyToken func(ctx context.Context, token string) (string, error), opts ...Option) kratosHttp.FilterFunc {
	o := &options{
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, opt := range opts {
		opt(o)
	}

	jwtTokenConvert := func(ctx context.Context, jwtToken string) (uid string, err error) {
		var (
			tokenInfo *jwt.Token
		)
		if o.claims != nil {
			tokenInfo, err = jwt.ParseWithClaims(jwtToken, o.claims(), keyFunc)
		} else {
			tokenInfo, err = jwt.Parse(jwtToken, keyFunc)
		}
		if err != nil {
			ve, ok := err.(*jwt.ValidationError)
			if !ok {
				err = errors.Unauthorized(reason, err.Error())
				return
			}
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = ErrTokenInvalid
				return
			}
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				err = ErrTokenExpired
				return
			}
			err = ErrTokenParseFail
			return

		}
		if !tokenInfo.Valid {
			err = ErrTokenInvalid
			return
		}
		if tokenInfo.Method != o.signingMethod {
			err = ErrUnSupportSigningMethod
			return
		}

		uid, err = verifyToken(ctx, jwtToken)
		if err != nil {
			err = ErrTokenExpired
			return
		}
		if claims, ok := tokenInfo.Claims.(*jwt.RegisteredClaims); !ok || uid != claims.ID {
			err = ErrTokenExpired
			return
		}
		return
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if keyFunc == nil {
				kratosHttp.DefaultErrorEncoder(w, r, ErrMissingKeyFunc)
				return
			}

			if strings.HasPrefix(r.URL.Path, VERIFY_PATH) {
				//需要验证
				auths := strings.SplitN(r.Header.Get(AuthorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], BearerWord) {
					kratosHttp.DefaultErrorEncoder(w, r, ErrMissingJwtToken)
					return
				}
				jwtToken := auths[1]
				uid, err := jwtTokenConvert(r.Context(), jwtToken)
				if err != nil {
					kratosHttp.DefaultErrorEncoder(w, r, err)
					return
				}
				r.Header.Set(AUTHEDKEY, uid)
				next.ServeHTTP(w, r)
			} else {
				//尝试验证放入数据
				auths := strings.SplitN(r.Header.Get(AuthorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], BearerWord) {
					next.ServeHTTP(w, r)
					return
				}
				jwtToken := auths[1]
				uid, err := jwtTokenConvert(r.Context(), jwtToken)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				r.Header.Set(AUTHEDKEY, uid)
				next.ServeHTTP(w, r)
			}
		})
	}
}
