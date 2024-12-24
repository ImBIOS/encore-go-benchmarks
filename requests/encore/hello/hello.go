// hello.go
package hello

import (
	"context"
	"fmt"
	"strconv"

	"encore.dev/beta/errs"
)

// Response represents the response format for all endpoints in this service.
type Response struct {
	Message string `json:"message"`
}

// HelloHandler handles GET requests to "/hello" and returns a static "Hello World!" message.
//encore:api public method=GET path=/hello
func HelloHandler(ctx context.Context) (*Response, error) {
	return &Response{Message: "Hello World!"}, nil
}

// Params is a generic type with an ID field, similar to Foo<T> in TypeScript.
type Params struct {
	ID int `json:"id"`
}

// HelloIDHandler handles GET requests to "/hello/:id" and returns a message with the provided ID.
//encore:api public method=GET path=/hello/:id
func HelloIDHandler(ctx context.Context, id string) (*Response, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, &errs.Error{Code: errs.InvalidArgument, Message: "ID must be a number"}
	}
	msg := fmt.Sprintf("hello %d %s", parsedID, "int")
	return &Response{Message: msg}, nil
}

// Schema represents the structure for the POST /schema request data, replicating your Schema interface.
type Schema struct {
	Foo                      string  `json:"foo" header:"x-foo"`
	Name                     *string `json:"name"`
	Excitement               *int    `json:"excitement"`
	SomeKey                  *string `json:"someKey"`
	SomeOtherKey             *int    `json:"someOtherKey"`
	RequiredKey              []int   `json:"requiredKey"`
	NullableKey              *int    `json:"nullableKey"`
	MultipleTypesKey         string  `json:"multipleTypesKey"` // Replace with specific type as needed
	MultipleRestrictedTypesKey int   `json:"multipleRestrictedTypesKey"` // Replace with specific type as needed
	EnumKey                  *string `json:"enumKey"`
}


// SchemaHandler handles POST requests to "/schema" and returns a "Hello, World" message.
//encore:api public method=POST path=/schema
func SchemaHandler(ctx context.Context, data *Schema) (*Response, error) {
	// Custom processing or validation for fields like EnumKey, MultipleTypesKey, etc. can go here.
	if data.EnumKey != nil && *data.EnumKey != "John" && *data.EnumKey != "Foo" {
		return nil, &errs.Error{
			Code:    errs.InvalidArgument,
			Message: fmt.Sprintf("enumKey must be 'John' or 'Foo', got '%s'", *data.EnumKey),
		}
	}

	return &Response{Message: "Hello, World"}, nil
}
