package leaderboards

import (
	"coderangers/cookies"
	"coderangers/dao"
	"coderangers/templating"
	"net/http"
)

func LeaderboardsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := GetTopUsers(100, 0)
		if err != nil {
			templating.ErrorPage(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := struct {
			Leaderboards []User
			IsAdmin      bool
			IsLoggedIn   bool
		}{
			users,
			dao.IsAdmin(r),
			cookies.IsLoggedIn(r),
		}
		templating.RenderPageWithBase(w, "leaderboards", data)
	default:
		templating.ErrorPage(w, "", http.StatusMethodNotAllowed)
	}
}
