package ident

import (
	"fmt"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

var (
	// ErrUnexpectedTypeName 오류는 타입 이름이 기대와 다를 때의 오류이다.
	ErrUnexpectedTypeName = fmt.Errorf("ident: UnexpectedTypeName")
	// ErrInvalidGlobalID 오류는 주어진 아이디의 형식이 올바르지 않을 때의 오류이다.
	ErrInvalidGlobalID = fmt.Errorf("ident: InvalidGlobalID")
)

// EncodeID 함수는 타입이름과 아이디를 통해 globalID를 포매팅한다.
func EncodeID(typeName string, id string) graphql.ID {
	return graphql.ID(fmt.Sprintf("%s:%s", typeName, id))
}

// DecodeID 함수는 주어진 globalID에서 타입명과 아이디를 찾아 반환한다.
func DecodeID(globalID graphql.ID) (typeName string, id string, err error) {
	tokens := strings.SplitN(string(globalID), ":", 2)
	if len(tokens) != 2 {
		return "", "", fmt.Errorf("%w(%s)", ErrInvalidGlobalID, globalID)
	}
	return tokens[0], tokens[1], nil
}

// DecodeIDExpecting 함수는 주어진 globalID가 expectedTypeName 타입을 가질 것으로 기대하며 디코딩 후 아이디를 반환한다.
func DecodeIDExpecting(expectedTypeName string, globalID graphql.ID) (string, error) {
	typeName, id, err := DecodeID(globalID)
	if err != nil {
		return "", err
	}
	if typeName != expectedTypeName {
		return "", fmt.Errorf("%w(expected %s but got: %s)", ErrUnexpectedTypeName, expectedTypeName, typeName)
	}
	return id, nil
}

// DecodeIDPointerExpecting 함수는 포인터에 대해 DecodeIDExpecting 함수를 실행한다.
func DecodeIDPointerExpecting(expectedTypeName string, globalIDp *graphql.ID) (*string, error) {
	if globalIDp == nil {
		return nil, nil
	}
	typeName, id, err := DecodeID(*globalIDp)
	if err != nil {
		return nil, err
	}
	if typeName != expectedTypeName {
		return nil, fmt.Errorf("%w(expected %s but got: %s)", ErrUnexpectedTypeName, expectedTypeName, typeName)
	}
	return &id, nil
}

// DecodeIDSliceExpecting 함수는 graphql.ID 슬라이스의 각 요소에 대해 DecodeIDSliceExpecting 함수를 실행한다.
func DecodeIDSliceExpecting(expectedTypeName string, globalIDs []graphql.ID) (res []string, err error) {
	res = make([]string, len(globalIDs))
	for i := range globalIDs {
		res[i], err = DecodeIDExpecting(expectedTypeName, globalIDs[i])
		if err != nil {
			break
		}
	}
	return res, err
}
