package interceptor

import (
	"context"

	"github.com/soner3/evently/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthenticationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if ShouldNotFilter(info.FullMethod) {
			return handler(ctx, req)
		}
		err = checkAuth(&ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "User Unauthorized: %v", err)
		}
		return handler(ctx, req)
	}
}

func checkAuth(ctx *context.Context) error {
	token, err := util.ExtractTokenFromHeader(ctx)
	if err != nil {
		return err
	}
	_, _, err = util.ValidateToken(*token)
	if err != nil {
		return err
	}

	return nil
}

func ShouldNotFilter(method string) bool {
	switch {
	case method == "/auth.v1.AuthService/Login":
		return true
	case method == "/user.v1.UserService/CreateUser":
		return true
	default:
		return false
	}
}

func GetSubFromToken(ctx *context.Context) (string, error) {
	token, err := util.ExtractTokenFromHeader(ctx)
	if err != nil {
		return "", err
	}
	_, claims, err := util.ValidateToken(*token)
	if err != nil {
		return "", err
	}

	return claims.Subject, nil

}
