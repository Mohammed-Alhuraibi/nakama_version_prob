package server

import (
	"context"

	frostlineapi "github.com/FrostlineGamesRepos/FrostlineGames-Nakama/api"
	"github.com/heroiclabs/nakama-common/runtime"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetUserData fetches the user data from the database.
func GetUserInfo(ctx context.Context, emp *emptypb.Empty) (*frostlineapi.UserInfoResponse, error) {
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "no user ID in context")
	}

	// userID := ctx.Value(ctxUserIDKey{}).(uuid.UUID)

	userInfo := &frostlineapi.UserInfoResponse{
		Id:          userID,
		Username:    "test",
		DisplayName: "test",
		AvatarUrl:   "test",
		Email:       "test",
		Lang:        "test",
		Location:    "test",
		Timezone:    "test",
		Coins:       "0",
		Level:       "0",
		Xp:          "0",
		Gems:        "0",
		LastLogin:   "test",
	}

	return userInfo, nil
}

func GetInventory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*frostlineapi.GetUserInventory, error) {
	// userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	// if !ok {
	// 	return nil, status.Error(codes.InvalidArgument, "no user ID in context")
	// }

	// userID := ctx.Value(ctxUserIDKey{}).(uuid.UUID)

	inventory := &frostlineapi.GetUserInventory{
		Items: []*frostlineapi.InventoryItem{
			{
				Id:          "1",
				Name:        "test",
				Description: "test",
				Price:       "0",
				ImageUrl:    "test",
				Currency:    "0",
			},
		},
	}

	return inventory, nil
}

// GetShop fetches the shop data from the database.
func GetShop(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*frostlineapi.GetShopData, error) {
	shop := &frostlineapi.GetShopData{
		Items: []*frostlineapi.ShopItem{
			{
				Id:          "1",
				Name:        "test",
				Description: "test",
				Price:       "0",
				ImageUrl:    "test",
				Currency:    "0",
			},
		},
	}

	return shop, nil
}

// Get purchased transactions
func GetTransactions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*frostlineapi.GetTransactionsList, error) {
	transactions := &frostlineapi.GetTransactionsList{
		Transactions: []*frostlineapi.Transaction{
			{
				Id:          "1",
				Name:        "test",
				Description: "test",
			},
		},
	}

	return transactions, nil
}

// get spesific transactions for UserId
func GetTransactionInfo(ctx context.Context, in *frostlineapi.TransactionId, opts ...grpc.CallOption) (*frostlineapi.GetTransactionResponse, error) {
	transaction := &frostlineapi.GetTransactionResponse{
		Id:          "1",
		Name:        "test",
		Description: "test",
	}

	return transaction, nil
}
