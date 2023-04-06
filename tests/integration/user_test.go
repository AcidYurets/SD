package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService(t *testing.T) {
	uuid := "123e4567-e89b-12d3-a456-426655440000"

	user, err := service.GetByUuid(context.Background(), uuid)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
