package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/web_videos/scheduler/dbops"
	"net/http"
)

func videoDelRecRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	if len(vid) == 0 {
		sendResponse(w, 400, "video should not be empty")
		return
	}

	err := dbops.AddVideoDeletionRecord(vid)
	if err != nil {
		sendResponse(w, 500, "Internal error")
		return
	}

	sendResponse(w, 200, "")
}
