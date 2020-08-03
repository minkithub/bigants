// Package web 은 http 핸들러를 구성해 서버를 구현한다.
package web

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/allscape/bigants/grasshopper/app/auth"
	"github.com/allscape/bigants/grasshopper/app/gql"
	"github.com/allscape/bigants/grasshopper/model/iam"
)

type Options struct {
	Port               string
	CorsAllowedOrigins []string
}

func New(
	gqls *gql.Service,
	iam *iam.Service,
	auths *auth.Service,
	opt Options,
) *Service {
	return &Service{gqls, iam, auths, opt}
}

type Service struct {
	gqls  *gql.Service
	iam   *iam.Service
	auths *auth.Service
	opt   Options
}

func (s *Service) Start(ctx context.Context) error {
	mux := mux.NewRouter()

	// 핸들러함수가 s의 메서드여야 할까, 아니면 스태틱 생성자를 가져야 할까?
	// 둘 다 가능하긴 한데 어떤 게 최선일지 모르겠다.
	// 스태틱 생성자에서는 의존성이 명시적으로 드러난다.
	mux.HandleFunc("/playground", handlePlayground)
	mux.Handle("/graphql", newGraphQLHandler(s.gqls))
	mux.Handle("/authenticate", newAuthenticateHandler(s.auths))
	mux.Path("/users").Methods("POST").HandlerFunc(s.PostUsers)

	var h http.Handler = mux
	// h := http.TimeoutHandler(mux, 7*time.Second, "Timeout")
	h = withAuth(s.auths, h)
	h = withCors(s.opt.CorsAllowedOrigins, h)

	server := &http.Server{Addr: ":" + string(s.opt.Port), Handler: h}

	go func() {
		<-ctx.Done()
		server.Close()
	}()
	return server.ListenAndServe()
}
