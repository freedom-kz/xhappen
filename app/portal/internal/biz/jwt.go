package biz

import (
	"context"
	"strconv"
	"time"
	"xhappen/app/portal/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
)

type JWTRepo interface {
	SaveToken(ctx context.Context, token string, id string) error
	VerifyToken(ctx context.Context, token string) (string, error)
}

type jwtTokenOption struct {
	signingMethod jwt.SigningMethod
	signingKey    []byte
	expire        time.Duration
}

type JwtUseCase struct {
	option *jwtTokenOption
	repo   JWTRepo
	log    *log.Helper
}

func NewJwtUseCase(bc *conf.Bootstrap, repo JWTRepo, logger log.Logger) *JwtUseCase {
	option := &jwtTokenOption{
		signingMethod: jwt.SigningMethodHS256,
		signingKey:    []byte(bc.Auth.Jwt.Secret),
		expire:        time.Duration(bc.Auth.Jwt.Expiration),
	}

	return &JwtUseCase{
		option: option,
		repo:   repo,
		log:    log.NewHelper(log.With(logger, "module", "usecase/jwt")),
	}
}

func (useCase *JwtUseCase) GenerateToken(ctx context.Context, id int64) (string, error) {
	idStr := strconv.FormatInt(id, 10)
	claims := &jwt.RegisteredClaims{
		Issuer:    "xhappen",
		Subject:   "xhappen",
		ID:        idStr,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(useCase.option.expire)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(useCase.option.signingMethod, claims)
	tokenStr, err := token.SignedString(useCase.option.signingKey)
	if err != nil {
		return "", err
	}
	err = useCase.repo.SaveToken(ctx, tokenStr, idStr)
	return tokenStr, err
}

func (useCase *JwtUseCase) VerifyToken(ctx context.Context, token string) (string, error) {
	return useCase.repo.VerifyToken(ctx, token)
}
