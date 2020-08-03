// Package auth 는 토큰 기반의 인증 로직을 구현한다.
package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/allscape/bigants/grasshopper/model/iam"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrCredential 오류는 토큰이 잘못된 형식으로 오는 경우 발생한다
	ErrCredential = errors.New("Credential Error")
	// ErrBasicAuth 오류는 아이디와 비밀번호가 매치하다가 실패했을 때 발생한다. 데이터베이스 통신 오류일수도 있다.
	ErrBasicAuth = errors.New("Basic Authentication Error")
	// ErrBearerAuth 오류는 유효하지 않은 JWT에서 발생한다.
	ErrBearerAuth = errors.New("Bearer Authentication Error")
)

type Options struct {
	AppName     string
	JWTSecret   string
	JWTDuration time.Duration
}

const idClaimKey = "id"

// Service 는 기본적인 인증 관련 메서드들을 제공한다.
type Service struct {
	opt  Options
	iams *iam.Service
}

// New 함수는 새로운 서비스를 만든다.
func New(opt Options, iams *iam.Service) *Service {
	return &Service{opt, iams}
}

func (sv *Service) SignJWT(id string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		idClaimKey: id,
		"exp":      time.Now().Add(time.Second * sv.opt.JWTDuration).Unix(),
		"iss":      sv.opt.AppName,
	})
	tokenString, err := token.SignedString([]byte(sv.opt.JWTSecret))
	return &tokenString, err
}

func (sv *Service) EncodeBasic(id string, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", id, password)))
}

func (sv *Service) DecodeToken(ctx context.Context, token string) (user *iam.User, err error) {
	if len(token) == 0 {
		return nil, ErrCredential
	}
	parts := strings.SplitN(token, " ", 2)
	if len(parts) != 2 {
		return nil, ErrCredential
	}
	scheme := parts[0]
	credential := parts[1]

	var id string
	if scheme == "Basic" {
		id, err = sv.doBasicAuth(ctx, credential)
	} else if scheme == "Bearer" {
		id, err = sv.doBearerAuth(ctx, credential)
	} else {
		return nil, ErrCredential
	}
	if err != nil {
		return nil, err
	}
	return sv.iams.FindUserOrFail(ctx, id)
}

func (sv *Service) doBasicAuth(ctx context.Context, credential string) (string, error) {
	payload, _ := base64.StdEncoding.DecodeString(credential)
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return "", ErrCredential
	}
	id := pair[0]
	password := pair[1]

	user, err := sv.iams.FindUserOrFail(ctx, id)
	if err != nil {
		return "", fmt.Errorf("auth: %w: %v", ErrBasicAuth, err)
	}
	if err := user.MatchPassword(password); err != nil {
		return "", err
	}
	return user.ID, nil
}

func (sv *Service) doBearerAuth(ctx context.Context, credential string) (string, error) {
	token, err := sv.validateJWT(&credential)
	if err != nil {
		return "", fmt.Errorf("auth: invalid JWT: %w", ErrBearerAuth)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return "", fmt.Errorf("auth: claim is not ok: %w", ErrBearerAuth)
	}
	id, ok := claims[idClaimKey].(string)
	if !ok {
		return "", fmt.Errorf("auth: username should be string: %w", ErrBearerAuth)
	}
	return id, nil
}

func (sv *Service) validateJWT(tokenString *string) (*jwt.Token, error) {
	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("auth: unexpected signing method(%v): %w", token.Header["alg"], ErrCredential)
		}
		return []byte(sv.opt.JWTSecret), nil
	})
	return token, err
}
