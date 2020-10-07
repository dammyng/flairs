package authserver

import (
	"context"
	"fmt"
	"log"
	"shared/models/appuser"
)

// GetUser - get a single user
func (a *AuthServer) FindCardRequest(ctx context.Context, in *appuser.CardRequest) (*appuser.CardRequest, error) {
	result, err := a.DbHandler.FindCardRequestById(in.ID)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUser - get a single user
func (a *AuthServer) FindUserCardRequests(ctx context.Context, in *appuser.CardRequest) (*appuser.CardRequests, error) {

	var u = []*appuser.CardRequest{}
	data, err := a.DbHandler.FindUserCardRequests(in.UserId)
	if err != nil {
		log.Print("Error querying Db")
		return nil, err
	}
	for _, t := range data {
		u = append(u, &t)
	}
	return &appuser.CardRequests{
		Results: u,
	}, nil

}

// GetUser - get a single user
func (a *AuthServer) FindWallet(ctx context.Context, in *appuser.Wallet) (*appuser.Wallet, error) {
	result, err := a.DbHandler.FindWalletById(fmt.Sprintf("%v", in.WalletSig))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUser - get a single user
func (a *AuthServer) NewCardRequest(ctx context.Context, in *appuser.CardRequest) (*appuser.CardRequest, error) {
	err := a.DbHandler.AddCardRequest(*in)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// GetUser - get a single user
func (a *AuthServer) NewWallet(ctx context.Context, in *appuser.Wallet) (*appuser.Wallet, error) {
	err := a.DbHandler.AddWallet(*in)
	if err != nil {
		return nil, err
	}
	return new(appuser.Wallet), err
}


// GetUser - get a single user
func (a *AuthServer) UserWallets(ctx context.Context, in *appuser.WalletArg) (*appuser.Wallets, error) {

	var u = []*appuser.Wallet{}
	data, err := a.DbHandler.FindUserWallets(in.UserId)
	if err != nil {
		log.Print("Error querying Db")
		return nil, err
	}
	for _, t := range data {
		u = append(u, &t)
	}
	return &appuser.Wallets{
		Results: u,
	}, nil

}