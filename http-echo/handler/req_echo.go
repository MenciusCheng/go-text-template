package handler

import "time"

type ReqEcho struct {
	CreatedAt time.Time
	Request   string
}

var ReqEchoList = make([]ReqEcho, 0)
