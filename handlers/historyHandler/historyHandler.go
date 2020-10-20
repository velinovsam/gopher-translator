package historyHandler

import (
	"github.com/gin-gonic/gin"
)

type serviceHistory struct {
	Entries []map[string]string `json:"history"`
}

var h serviceHistory

func HandleHistoryRequest(ctx *gin.Context) {
	ctx.JSON(200, h)
}

func AddToHistory(original string, translation string) {

	entry := make(map[string]string)
	entry[original] = translation

	h.Entries = append(h.Entries, entry)

}