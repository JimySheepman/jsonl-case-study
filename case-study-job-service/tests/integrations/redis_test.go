//go:build integration

package integrations

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type testStruct struct {
	Test string `json:"test"`
}

func Test_Integration_Redis_Get(t *testing.T) {
	ctx := context.Background()
	key := "test-key"
	value := "test-value"
	expiration := time.Duration(0)

	setKey := redisRepo.Set(ctx, key, value, expiration)
	require.Nil(t, setKey.Err())

	getKey := redisRepo.Get(ctx, key)
	require.Nil(t, getKey.Err())

	actual, err := getKey.Result()
	require.Nil(t, err)
	require.Equal(t, value, actual)
}

func Test_Integration_Redis_Set(t *testing.T) {
	ctx := context.Background()
	key := "test-key-2"
	value := "test-value-2"
	expiration := time.Duration(0)

	setKey := redisRepo.Set(ctx, key, value, expiration)
	require.Nil(t, setKey.Err())

	getKey := redisRepo.Get(ctx, key)
	require.Nil(t, getKey.Err())

	actual, err := getKey.Result()
	require.Nil(t, err)
	require.Equal(t, value, actual)
}

func Test_Integration_Redis_GetBody(t *testing.T) {
	ctx := context.Background()
	key := "test-key-3"
	value := `{"test":"test"}`
	expiration := time.Duration(0)

	setKey := redisRepo.Set(ctx, key, value, expiration)
	require.Nil(t, setKey.Err())

	ts := &testStruct{}
	ok, err := redisRepo.GetBody(ctx, key, ts)
	require.Nil(t, err)
	require.Equal(t, true, ok)
}
