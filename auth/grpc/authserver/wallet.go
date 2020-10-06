package authserver

import (
	"context"
	"shared/models/appuser"
)

// GetUser - get a single user
func (a *AuthServer) FindCardRequest(ctx context.Context, in *appuser.CardRequest) (*appuser.CardRequest, error) {
	result, err :=a.DbHandler.FindCardRequestById(in.ID)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUser - get a single user
// func (a *AuthServer) FindUserCardRequests(ctx context.Context, in string) (*appuser.CardRequests, error) {
//	result, err :=a.DbHandler.FindUserCardRequests(in)
//	if err != nil {
//		return nil, err
//	}
//	return &result, nil
//}

// GetUser - get a single user
func (a *AuthServer) FindWallet(ctx context.Context, in *appuser.Wallet) (*appuser.Wallet, error) {

	result, err := a.DbHandler.Fin
	return &appuser.Wallet{}, nil
}

// GetUser - get a single user
func (a *AuthServer) NewCardRequest(ctx context.Context, in *appuser.CardRequest) (*appuser.CardRequest, error) {
	return &appuser.CardRequest{}, nil
}

// GetUser - get a single user
func (a *AuthServer) NewWallet(ctx context.Context, in *appuser.Wallet) (*appuser.Wallet, error) {
	return &appuser.Wallet{}, nil
}
