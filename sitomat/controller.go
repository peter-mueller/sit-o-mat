package sitomat

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/peter-mueller/sit-o-mat/user"
)

type Controller struct {
	Service *Service
}

func (c *Controller) ManualAssign(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username, password, ok := r.BasicAuth()
	if !ok || !user.LoginAdmin(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	err := c.Service.AssignWorkplaces(ctx)
	if err != nil {
		panic(err)
	}
}
