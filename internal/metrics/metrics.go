package metrics

import "sync/atomic"

var (
	HS256Count int64
	RS256Count int64
)

func IncHS256() { atomic.AddInt64(&HS256Count, 1) }
func IncRS256() { atomic.AddInt64(&RS256Count, 1) }
