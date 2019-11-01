package workplace

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/peter-mueller/sit-o-mat/user"
)

type Controller struct {
	Service *Service
}

func (c *Controller) CreateWorkplace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username, password, ok := r.BasicAuth()
	if !ok || !user.LoginAdmin(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var workplace Workplace
	err := json.NewDecoder(r.Body).Decode(&workplace)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	workplace, err = c.Service.CreateWorkplace(ctx, workplace)
	handle(err)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workplace)
}

func (c *Controller) DeleteWorkplace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	ctx := context.Background()
	err := c.Service.DeleteWorkplaceByName(ctx, name)
	handle(err)
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) GetAllWorkplaces(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	owner := r.URL.Query().Get("owner")

	ctx := context.Background()
	workplaces, err := c.Service.FindAllWorkplacesSortByRating(ctx)
	handle(err)

	filteredWorkplaces := make([]Workplace, 0)

	for _, workplace := range workplaces {
		if owner == "" || workplace.CurrentOwner == owner {
			filteredWorkplaces = append(filteredWorkplaces, workplace)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filteredWorkplaces)
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
