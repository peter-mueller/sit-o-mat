package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"context"

	"github.com/julienschmidt/httprouter"

	"github.com/peter-mueller/sit-o-mat/httperror"
	"github.com/peter-mueller/sit-o-mat/sitomat"
	"github.com/peter-mueller/sit-o-mat/user"
	"github.com/peter-mueller/sit-o-mat/workplace"

	"fmt"

	"gocloud.dev/docstore"

	_ "gocloud.dev/docstore/memdocstore"

	"errors"
)

func userCollection() *docstore.Collection {
	ctx := context.Background()
	coll, err := docstore.OpenCollection(ctx, "mem://user/name")
	if err != nil {
		panic(err)
	}
	return coll
}
func workplaceCollection() *docstore.Collection {
	ctx := context.Background()
	coll, err := docstore.OpenCollection(ctx, "mem://workplace/name")
	if err != nil {
		panic(err)
	}
	return coll
}

func main() {
	coll := userCollection()
	defer coll.Close()
	userService := user.Service{Collection: coll}
	userController := user.Controller{Service: &userService}

	workplaceColl := workplaceCollection()
	defer workplaceColl.Close()
	workplaceService := workplace.Service{Collection: workplaceColl}
	workplaceController := workplace.Controller{Service: &workplaceService}

	sitomatService := sitomat.Service{
		UserService:      &userService,
		WorkplaceService: &workplaceService,
	}
	sitomatController := sitomat.Controller{Service: &sitomatService}

	r := httprouter.New()
	r.POST("/user", userController.RegisterUser)
	r.GET("/user/:name", userController.GetUser)
	r.DELETE("/user/:name", userController.DeleteUser)
	r.PATCH("/user/:name", userController.PatchUser)

	r.POST("/workplace", workplaceController.CreateWorkplace)
	r.DELETE("/workplace/:name", workplaceController.DeleteWorkplace)
	r.GET("/workplace", workplaceController.GetAllWorkplaces)

	r.GET("/sitomat", sitomatController.ManualAssign)
	r.PanicHandler = panicHandler

	fmt.Println("Starting Server")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

func panicHandler(w http.ResponseWriter, r *http.Request, data interface{}) {
	err, ok := data.(error)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var httperr httperror.Error
	if errors.As(err, &httperr) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httperr.Status)
		json.NewEncoder(w).Encode(httperr)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
}
