package service

import (
	"context"
	"strconv"
	"strings"
	v1 "xhappen/api/basic/v1"
	"xhappen/pkg/filter/jwt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport"
)

func GetUserID(ctx context.Context) (int64, error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return 0, v1.ErrorUnknown("context info loss")
	}
	uid := strings.TrimSpace(tr.RequestHeader().Get(jwt.AUTHEDKEY))

	if uid != "" {
		id, err := strconv.Atoi(uid)
		if err != nil {
			return 0, v1.ErrorUnknown("can not get userid from kratos context")
		}
		return int64(id), nil
	} else {
		return 0, errors.Unauthorized("UNAUTHORIZED", "can not get uid from header")
	}
}

func GetToken(ctx context.Context) (string, error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return "", v1.ErrorUnknown("is not kratos context")
	}

	auths := strings.SplitN(tr.RequestHeader().Get(jwt.AuthorizationKey), " ", 2)
	if len(auths) != 2 || !strings.EqualFold(auths[0], jwt.BearerWord) {
		return "", errors.Unauthorized("UNAUTHORIZED", "JWT token is missing")
	}

	jwtToken := auths[1]

	return jwtToken, nil
}
