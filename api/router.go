package api

import (
	"github.com/julienschmidt/httprouter"
	activatevoting "github.com/programmierigel/voting/api/activateVoting"
	checkpassword "github.com/programmierigel/voting/api/checkPassword"
	deactivatevoting "github.com/programmierigel/voting/api/deactivateVoting"
	deleteall "github.com/programmierigel/voting/api/deleteAll"
	getcandidates "github.com/programmierigel/voting/api/getCandidates"
	insertnewvotable "github.com/programmierigel/voting/api/insertNewVotable"
	"github.com/programmierigel/voting/api/makevote"
	"github.com/programmierigel/voting/api/ping"
	showvoting "github.com/programmierigel/voting/api/showVoting"
	votingactive "github.com/programmierigel/voting/api/votingActive"
	"github.com/programmierigel/voting/storage"
)

func GetRouter(store storage.Store, password string) *httprouter.Router {
	router := httprouter.New()

	// QUERY
	router.GET("/ping", ping.Handle())
	router.GET("/showVoting", showvoting.Handle(store))
	router.GET("/getCandidates", getcandidates.Handle(store))
	router.GET("/votingActive", votingactive.Handle(store))
	// COMMAND
	router.POST("/activateVoting", activatevoting.Handle(store))
	router.POST("/deactivateVoting", deactivatevoting.Handle(store))
	router.POST("/deleteAll", deleteall.Handle(store))
	router.POST("/makeVote", makevote.Handle(store))
	router.POST("/insertNewVotable", insertnewvotable.Handle(store))
	router.POST("/checkPassword", checkpassword.Handle(store))

	return router
}
