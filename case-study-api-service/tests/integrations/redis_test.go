//go:build integration

package integrations

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

type testStruct struct {
	Test string `json:"test"`
}

func Test_Integration_Redis_GetBody(t *testing.T) {
	ctx := context.Background()
	key := "test-key-3"

	ts := &testStruct{}
	ok, err := redisRepo.GetProduct(ctx, key, ts)
	require.Nil(t, err)
	require.Equal(t, false, ok)
}
