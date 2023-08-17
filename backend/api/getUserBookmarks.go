package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func (h *BaseHandler) GetUserBookmarks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	userid := payload.UserID

	q := sqlc.New(h.pool)

	bookmarks, err := q.GetUserBookmarks(ctx, userid)
	if err != nil {
		var pgErr *pgconn.PgError

		switch {
		case errors.As(err, &pgErr):
			log.Println(pgErr)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return

		case errors.Is(err, pgx.ErrNoRows):
			log.Println(err)
			utils.Response(w, "user bookmarks not found", http.StatusNoContent)
			return

		default:
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	utils.JsonResponse(w, bookmarks)
}
