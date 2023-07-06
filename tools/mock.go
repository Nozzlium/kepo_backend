package tools

import (
	"context"
	"nozzlium/kepo_backend/constants"
)

var ClaimContext = JwtClaims{
	UserId: 1,
}

func GetMockClaimContext(parent context.Context) context.Context {
	return context.WithValue(parent, constants.USER_ID_CLAIMS, ClaimContext)
}
