package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	Service *Service
}

func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username, password, ok := r.BasicAuth()
	if !ok || !LoginAdmin(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	user, err = c.Service.RegisterUser(ctx, user.Name)
	handle(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&struct {
		User     User
		Password string
	}{
		User:     user,
		Password: user.Password,
	})
	handle(err)
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := User{
		Name: ps.ByName("name"),
	}
	ctx := context.Background()
	user, err := c.Service.FindUserByName(ctx, user.Name)
	handle(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username, password, ok := r.BasicAuth()
	if !ok || !LoginAdmin(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := User{
		Name: ps.ByName("name"),
	}
	ctx := context.Background()
	err := c.Service.DeleteUser(ctx, user)
	handle(err)

	w.WriteHeader(http.StatusOK)
}

var (
	RequestDayOperation = "RequestDay"
)

type patchRequest struct {
	Operation string          `json:"op"`
	Path      string          `json:"path"`
	Value     json.RawMessage `json:"value"`
}

func (c *Controller) PatchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username, password, ok := r.BasicAuth()
	if !ok || username != ps.ByName("name") {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user := User{
		Name:     ps.ByName("name"),
		Password: password,
	}

	ctx := context.Background()
	ok, err := c.Service.LoginUser(ctx, user)
	handle(err)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err = c.Service.FindUserByName(ctx, user.Name)
	handle(err)

	var patch patchRequest
	err = json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if patch.Operation == "replace" && patch.Path == "/WeeklyRequests" {
		var weeklyRequests WeeklyRequests
		err = json.Unmarshal(patch.Value, &weeklyRequests)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.WeeklyRequests = weeklyRequests
		user, err = c.Service.UpdateUser(ctx, user)
		handle(err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(&user)
		handle(err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
