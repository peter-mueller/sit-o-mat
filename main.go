package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/peter-mueller/sit-o-mat/httperror"
	"github.com/peter-mueller/sit-o-mat/sitomat"
	"github.com/peter-mueller/sit-o-mat/user"
	"github.com/peter-mueller/sit-o-mat/workplace"

	"fmt"

	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"

	"github.com/robfig/cron/v3"

	"errors"
)

func main() {

	userService := user.Service{}
	userController := user.Controller{Service: &userService}

	workplaceService := workplace.Service{}
	workplaceController := workplace.Controller{Service: &workplaceService}

	sitomatService := sitomat.Service{
		UserService:      &userService,
		WorkplaceService: &workplaceService,
	}
	sitomatController := sitomat.Controller{Service: &sitomatService}

	c := cron.New()
	_, err := c.AddFunc("@midnight", func() {
		ctx := context.Background()
		err := sitomatController.Service.AssignWorkplaces(ctx)
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	r := httprouter.New()
	r.HandleOPTIONS = true
	r.HandleMethodNotAllowed = true
	r.GlobalOPTIONS = http.HandlerFunc(corsHandler)

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
		Handler: corsDecorator{r},
		Addr:    lookupEnv("SITOMAT_ADDR", "127.0.0.1:8080"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("  Addr: ", srv.Addr)
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

func corsHandler(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers
	header := w.Header()
	header.Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, PUT")
	header.Set("Access-Control-Allow-Headers", "authorization")
	header.Set("Access-Control-Allow-Origin", "*")

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}

type corsDecorator struct {
	router *httprouter.Router
}

func (c corsDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	c.router.ServeHTTP(w, r)
}

func lookupEnv(env string, alternative string) string {
	value, ok := os.LookupEnv(env)
	if !ok {
		log.Printf("Using default for %v: %v", env, alternative)
		return alternative
	}
	return value
}
