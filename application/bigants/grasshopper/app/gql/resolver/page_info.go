package resolver

import (
	"github.com/allscape/bigants/grasshopper/model/list"
)

type PageInfo struct {
	HasNextPage     bool    `json:"has_next_page"`
	HasPreviousPage bool    `json:"has_previous_page"`
	StartCursor     *string `json:"start_cursor"`
	EndCursor       *string `json:"end_cursor"`
}

func (rsv *Resolver) ResolvePageInfo(cursors []string, info list.PageInfo) (next PageInfo) {
	if len(cursors) > 0 {
		next.StartCursor = &cursors[0]
		next.EndCursor = &cursors[len(cursors)-1]
	}
	next.HasNextPage = info.HasNextPage()
	next.HasPreviousPage = info.HasPreviousPage()
	return next
}
