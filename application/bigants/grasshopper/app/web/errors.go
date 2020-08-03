package web

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/allscape/bigants/grasshopper/app/auth"
)

// HandleError 함수는 오류를 사용자에게 HTTP 응답으로 반환할 때 쓰인다.
func handleError(w http.ResponseWriter, err error) {
	appErr := appErrors.Get(err)
	if appErr == nil {
		appErr = &appError{
			internal: err,
			status:   http.StatusInternalServerError,
			message:  "예상하지 못한 문제가 발생했습니다.",
		}
	}
	http.Error(w, appErr.message, appErr.status)
	log.Printf("[err:server] %d, %s: %s", appErr.status, appErr.message, appErr.internal)
}

type appError struct {
	internal error
	status   int
	message  string
}

type errRules map[error]*struct {
	status  int
	message string
}

var (
	errMethodNotAllowed = fmt.Errorf("Method Not Allowed")
)

// appErrors 맵은 앱이 다룰 수 있는 오류와 그 처리 결과 메시지를 나타낸다.
// 우리 앱 코드 내에서 정의한 오류만 처리하는 것을 원칙으로 한다.
// 어떤 외부 오류가 처리되어야 한다면 앱 내 오류로 재정의하고 감싼다.
var appErrors = errRules{
	errMethodNotAllowed: {
		status:  http.StatusMethodNotAllowed,
		message: "혀용되지 않은 메서드입니다.",
	},
	auth.ErrBasicAuth: {
		status:  http.StatusUnauthorized,
		message: "아이디 또는 비밀번호가 잘못되었습니다.",
	},
	auth.ErrBearerAuth: {
		status:  http.StatusUnauthorized,
		message: "로그인이 유효하지 않습니다. 다시 로그인 해 주세요.",
	},
	auth.ErrCredential: {
		status:  http.StatusBadRequest,
		message: "유효하지 않은 인증 방법입니다.",
	},
}

func (m *errRules) Get(err error) (appErr *appError) {
	for expected, item := range appErrors {
		if errors.Is(err, expected) {
			appErr = &appError{
				internal: err,
				status:   item.status,
				message:  item.message,
			}
			break
		}
	}
	return appErr
}
