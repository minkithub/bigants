// Package finder 는 키를 통해 엔티티를 찾거나 조건을 통해 엔티티 키를 페이지네이션한다.
package finder

import "github.com/allscape/bigants/grasshopper/app/db/tables"

// Adapter 구조체는 도메인 모델에 필요한 Finder를 구현한다.
type Adapter struct {
	DB tables.SQLHandle
}
