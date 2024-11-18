package user

import (
	"context"
	"sort"
	"testing"

	"encore.dev/beta/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		id           string
		expectedUser *UserResponse
		expectErr    bool
		expectedErr  string
		expectedCode errs.ErrCode
	}{
		{
			name:         "valid user - Alice",
			id:           "1",
			expectedUser: &UserResponse{User: User{ID: "1", Name: "Alice"}},
			expectErr:    false,
		},
		{
			name:         "valid user - Bob",
			id:           "2",
			expectedUser: &UserResponse{User: User{ID: "2", Name: "Bob"}},
			expectErr:    false,
		},
		{
			name:         "invalid user - not found",
			id:           "999",
			expectedUser: nil,
			expectErr:    true,
			expectedErr:  "not_found: user with id 999 not found",
			expectedCode: errs.NotFound,
		},
		{
			name:         "empty ID - invalid",
			id:           "",
			expectedUser: nil,
			expectErr:    true,
			expectedErr:  "not_found: user with id  not found",
			expectedCode: errs.NotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := GetUser(ctx, tt.id)

			if tt.expectErr {
				require.Error(t, err)
				assert.Nil(t, resp)
				assert.Equal(t, tt.expectedErr, err.Error())
				assert.Equal(t, tt.expectedCode, errs.Code(err))
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedUser, resp)
			}
		})
	}
}

// Helper function to sort users by ID
func sortUsersByID(users []User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()

	expectedUsers := []User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Caroline"},
		{ID: "4", Name: "Dave"},
	}

	resp, err := ListUsers(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Sort both slices by ID to ensure consistent comparison
	sortUsersByID(resp.Users)
	sortUsersByID(expectedUsers)

	assert.Equal(t, expectedUsers, resp.Users)
}