package graph

import (
	db "github.com/brkss/golang-graphql/db/sqlc"
	"github.com/brkss/golang-graphql/token"
	"github.com/brkss/golang-graphql/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Store db.Store
	Config *utils.Config
	Maker token.Maker
}
