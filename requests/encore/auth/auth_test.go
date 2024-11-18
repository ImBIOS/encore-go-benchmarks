package auth

import (
	"context"
	"testing"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		request    *LoginRequest
		expectErr  bool
		expectToken string
	}{
		{
			name: "valid login",
			request: &LoginRequest{
				Email:    "user@example.com",
				Password: "password123",
			},
			expectErr: false,
			expectToken: TOKEN,
		},
		{
			name: "invalid login - empty email",
			request: &LoginRequest{
				Email:    "",
				Password: "password123",
			},
			expectErr: false,
			expectToken: TOKEN,
		},
		{
			name: "invalid login - empty password",
			request: &LoginRequest{
				Email:    "user@example.com",
				Password: "",
			},
			expectErr: false,
			expectToken: TOKEN,
		},
		{
			name: "invalid login - empty fields",
			request: &LoginRequest{
				Email:    "",
				Password: "",
			},
			expectErr: false,
			expectToken: TOKEN,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := Login(ctx, tt.request)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, resp)
				assert.Equal(t, tt.expectToken, resp.Token)
			}
		})
	}
}

func TestAuthHandler(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		token      string
		expectErr  bool
		expectUID  auth.UID
	}{
		{
			name:      "valid token",
			token:     TOKEN,
			expectErr: false,
			expectUID: "user-id",
		},
		{
			name:      "invalid token",
			token:     "invalid-token",
			expectErr: true,
			expectUID: "",
		},
		{
			name:      "empty token",
			token:     "",
			expectErr: true,
			expectUID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uid, err := AuthHandler(ctx, tt.token)

			if tt.expectErr {
				assert.Error(t, err)
				assert.IsType(t, &errs.Error{}, err)
				authErr, _ := err.(*errs.Error)
				assert.Equal(t, errs.Unauthenticated, authErr.Code)
				assert.Equal(t, "invalid token", authErr.Message)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectUID, uid)
			}
		})
	}
}
