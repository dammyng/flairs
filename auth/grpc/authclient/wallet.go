package authclient


import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"shared/models/appuser"
)



// AddNewUser grpc client to add new user
func CreateNewWallet(in *appuser.Wallet) (*appuser.Wallet, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	result, err := authClient.NewWallet(context.Background(), in)

	if err != nil {
		return nil, err
	}
	return result, nil
}


// AddNewUser grpc client to add new user
func CreateCardRequest(in *appuser.CardRequest) (*appuser.CardRequest, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	result, err := authClient.NewCardRequest(context.Background(), in)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddNewUser grpc client to add new user
func GetUserCardRequest(in *appuser.CardRequest) (*appuser.CardRequests, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	result, err := authClient.FindUserCardRequests(context.Background(), in)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserWallets(in *appuser.WalletArg) (*appuser.Wallets, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	authClient := appuser.NewUserServiceClient(conn)

	result, err := authClient.UserWallets(context.Background(), in)

	if err != nil {
		return nil, err
	}
	return result, nil
}