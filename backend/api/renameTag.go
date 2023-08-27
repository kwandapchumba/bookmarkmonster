package api

import (
	"encoding/json"
	"log"
	"net/http"

	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type renameTagRequest struct {
	Tag sqlc.Tag `json:"tag"`
}

func (h *BaseHandler) RenameTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *renameTagRequest

	if err := jsonDecoder.Decode(&req); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		} else {
			log.Printf("error decoding request body to struct: %v", err)
			utils.Response(w, "something went wrong", http.StatusBadRequest)
			return
		}
	}

	q := sqlc.New(h.pool)

	tag, err := q.RenameTag(ctx, sqlc.RenameTagParams{Name: req.Tag.Name, ID: req.Tag.ID})
	if err != nil {
		log.Printf("could not rename tag: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, tag)
}
