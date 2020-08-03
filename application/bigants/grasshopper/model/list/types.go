// Package list 는 순서 있는 목록의 페이지네이션 기본형을 정의한다.
package list

import "encoding/json"

// PageOption 구조체는 커서 기반 페이지네이션 옵션을 정의한다.
type PageOption struct {
	First  *int32
	After  *string
	Last   *int32
	Before *string
}

// Connection is result of connection query
type Connection struct {
	Cursors  []string
	PageInfo PageInfo
}

// PageInfo implements Relay Connection PageInfo fields
type PageInfo struct {
	hasNextPage     bool
	hasPreviousPage bool
}

// HasNextPage returns if the connection has next page
func (i PageInfo) HasNextPage() bool { return i.hasNextPage }

// HasPreviousPage returns if the connection has previous page
func (i PageInfo) HasPreviousPage() bool { return i.hasPreviousPage }

// StartCursor returns the first cursor pointer if exists
func (c *Connection) StartCursor() *string {
	if len(c.Cursors) > 0 {
		return &c.Cursors[0]
	}
	return nil
}

// EndCursor returns the last cursor pointer if exists
func (c *Connection) EndCursor() *string {
	if len(c.Cursors) > 0 {
		return &c.Cursors[len(c.Cursors)-1]
	}
	return nil
}

// Receive is used for sql scanning
func (c *Connection) Receive(int) []interface{} {
	c.Cursors = append(c.Cursors, "")
	return []interface{}{&c.PageInfo.hasPreviousPage, &c.PageInfo.hasNextPage, &c.Cursors[len(c.Cursors)-1]}
}

func (c Connection) MarshalJSON() ([]byte, error) {
	type pageInfo struct {
		HasNextPage     bool    `json:"hasNextPage"`
		HasPreviousPage bool    `json:"hasPreviousPage"`
		StartCursor     *string `json:"startCursor"`
		EndCursor       *string `json:"endCursor"`
	}
	type connection struct {
		Cursors  []string `json:"cursors"`
		PageInfo pageInfo `json:"pageInfo"`
	}
	return json.Marshal(connection{
		Cursors: c.Cursors,
		PageInfo: pageInfo{
			HasNextPage:     c.PageInfo.HasNextPage(),
			HasPreviousPage: c.PageInfo.HasPreviousPage(),
			StartCursor:     c.StartCursor(),
			EndCursor:       c.EndCursor(),
		},
	})
}
