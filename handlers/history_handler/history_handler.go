package history_handler

import (
	"github.com/gin-gonic/gin"
	"sort"
)

type historyResponse struct {
	Entries []map[string]string `json:"history"`
}

var history map[string]string

func init() {
	history = make(map[string]string)
}

func HandleHistoryRequest(ctx *gin.Context) {

	res := historyResponse{}

	keys := make([]string, 0, len(history))
	for k := range history {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		entry := make(map[string]string)
		entry[k] = history[k]
		res.Entries = append(res.Entries, entry)
	}

	ctx.JSON(200, res)

}

func UpdateHistory(original string, translation string) {
	history[original] = translation
}
