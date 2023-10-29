package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/lomsa-dev/gonull"
)

type Button struct {
	ID    primitive.ObjectID `bson:"_id"`
	Label string
	Image gonull.Nullable[string]
	Audio gonull.Nullable[string]
}

