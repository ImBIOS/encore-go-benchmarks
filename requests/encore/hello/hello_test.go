// hello_test.go
package hello

import (
	"context"
	"testing"

	"encore.dev/beta/errs"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	ctx := context.Background()
	resp, err := HelloHandler(ctx)

	assert.NoError(t, err, "expected no error from HelloHandler")
	assert.NotNil(t, resp, "expected a non-nil response from HelloHandler")
	assert.Equal(t, "Hello World!", resp.Message, "expected Hello World! message")
}

func TestHelloIDHandler(t *testing.T) {
	ctx := context.Background()

	t.Run("valid ID", func(t *testing.T) {
		resp, err := HelloIDHandler(ctx, "42")

		assert.NoError(t, err, "expected no error for valid ID")
		assert.NotNil(t, resp, "expected a non-nil response for valid ID")
		assert.Equal(t, "hello 42 int", resp.Message, "expected correct message with ID")
	})

	t.Run("invalid ID", func(t *testing.T) {
		resp, err := HelloIDHandler(ctx, "abc")

		assert.Error(t, err, "expected error for invalid ID")
		assert.Nil(t, resp, "expected nil response for invalid ID")

		var encoreErr *errs.Error
		if assert.ErrorAs(t, err, &encoreErr) {
			assert.Equal(t, errs.InvalidArgument, encoreErr.Code, "expected InvalidArgument error code")
			assert.Equal(t, "ID must be a number", encoreErr.Message, "expected specific error message")
		}
	})
}

func TestSchemaHandler(t *testing.T) {
	ctx := context.Background()

	t.Run("valid Schema data", func(t *testing.T) {
		data := &Schema{
			Foo:         "bar",
			RequiredKey: []int{1, 2, 3},
		}
		resp, err := SchemaHandler(ctx, data)

		assert.NoError(t, err, "expected no error with valid Schema data")
		assert.NotNil(t, resp, "expected a non-nil response")
		assert.Equal(t, "Hello, World", resp.Message, "expected Hello, World message")
	})

	t.Run("invalid EnumKey", func(t *testing.T) {
		enumKey := "InvalidValue"
		data := &Schema{
			Foo:      "bar",
			EnumKey:  &enumKey,
			RequiredKey: []int{1, 2, 3},
		}
		resp, err := SchemaHandler(ctx, data)

		assert.Error(t, err, "expected error for invalid EnumKey")
		assert.Nil(t, resp, "expected nil response for invalid EnumKey")

		var encoreErr *errs.Error
		if assert.ErrorAs(t, err, &encoreErr) {
			assert.Equal(t, errs.InvalidArgument, encoreErr.Code, "expected InvalidArgument error code")
			assert.Contains(t, encoreErr.Message, "enumKey must be 'John' or 'Foo'", "expected specific error message")
		}
	})

	t.Run("missing RequiredKey", func(t *testing.T) {
		data := &Schema{
			Foo: "bar",
		}
		resp, err := SchemaHandler(ctx, data)

		assert.NoError(t, err, "expected no error even if RequiredKey is missing") // Adjust based on SchemaHandler's behavior for RequiredKey
		assert.NotNil(t, resp, "expected a non-nil response even if RequiredKey is missing")
		assert.Equal(t, "Hello, World", resp.Message, "expected Hello, World message")
	})
}
