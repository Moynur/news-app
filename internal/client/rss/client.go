//go:generate mockgen -package=accountusers -self_package=git.curve.tools/go/account-curve-cards/internal/adapter/accountusers -destination=client_mock.go git.curve.tools/go/account-curve-cards/internal/adapter/accountusers Client

package rss

import (
	"context"
	"fmt"
)

const (
	// Used to perform search queries based on the account id
	accountIDKey = "account_id"
	// Used to perform search queries based on the user ID
	userIDKey = "id"
)

type Client interface {
	GetUserAccounts(ctx context.Context, serviceToken string, userID int64) ([]string, error)
}

// Service communicates with account users
type accountUsers struct {
	client       Client
	serviceToken string
}

// NewClient produces a new account users client
func NewClient(client Client, serviceToken string) *accountUsers {
	return &accountUsers{
		client:       client,
		serviceToken: serviceToken,
	}
}

// Get User AccountID from account users
func (s *accountUsers) GetUserAccount(ctx context.Context, userID int64) (string, error) {
	accounts, err := s.client.GetUserAccounts(ctx, s.serviceToken, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get accountID from userID %v from account-users: %w", userID, err)
	}
	if len(accounts) < 1 {
		return "", fmt.Errorf("no accounts found for userID %v", userID)
	} else if len(accounts) > 1 {
		return "", fmt.Errorf("more than one account found for userID %v", userID)
	}

	return accounts[0], nil
}
