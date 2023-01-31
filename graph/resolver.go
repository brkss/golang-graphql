package graph

import db "github.com/brkss/golang-graphql/db/sqlc"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Store db.Store
}
