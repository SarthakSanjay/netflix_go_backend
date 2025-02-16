package utils

import (
	"context"
	"fmt"

	"github.com/sarthaksanjay/netflix-go/services"
	"github.com/sarthaksanjay/netflix-go/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type contextKey string

const UserContextKey contextKey = "user"

func ExtractUserIdFromContext(ctx context.Context) (primitive.ObjectID, error) {
	claims, ok := ctx.Value(types.UserContextKey).(*services.Claims)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("could not extract user from context")
	}
	return claims.UserId, nil
}
