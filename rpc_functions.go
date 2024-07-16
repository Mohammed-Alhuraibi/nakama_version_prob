package main

import (
	"context"
	"database/sql"

	"github.com/FrostlineGamesRepos/FrostlineGames-Nakama/server"
	"github.com/heroiclabs/nakama-common/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

type nakamaRpcFunc func(context.Context, runtime.Logger, *sql.DB, runtime.NakamaModule, string) (string, error)

func rpcGetUserInfo(marshaler *protojson.MarshalOptions, _ *protojson.UnmarshalOptions) nakamaRpcFunc {
	return func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

		resp, err := server.GetUserInfo(ctx, nil)

		if err != nil {
			logger.Error("error GetUserInfo: %v", err.Error())
			return "", errMarshal
		}

		response, err := marshaler.Marshal(resp)
		if err != nil {
			logger.Error("error marshaling response payload: %v", err.Error())
			return "", errMarshal
		}

		return string(response), nil

	}
}

func rpcGetInventory(marshaler *protojson.MarshalOptions, _ *protojson.UnmarshalOptions) nakamaRpcFunc {
	return func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

		resp, err := server.GetInventory(ctx, nil)

		if err != nil {
			logger.Error("error GetInventory: %v", err.Error())
			return "", errMarshal
		}

		response, err := marshaler.Marshal(resp)
		if err != nil {
			logger.Error("error marshaling response payload: %v", err.Error())
			return "", errMarshal
		}

		return string(response), nil

	}
}

func rpcGetShop(marshaler *protojson.MarshalOptions, _ *protojson.UnmarshalOptions) nakamaRpcFunc {
	return func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

		resp, err := server.GetShop(ctx, nil)

		if err != nil {
			logger.Error("error GetShop: %v", err.Error())
			return "", errMarshal
		}

		response, err := marshaler.Marshal(resp)
		if err != nil {
			logger.Error("error marshaling response payload: %v", err.Error())
			return "", errMarshal
		}

		return string(response), nil

	}
}

func rpcGetTransactions(marshaler *protojson.MarshalOptions, _ *protojson.UnmarshalOptions) nakamaRpcFunc {
	return func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

		resp, err := server.GetTransactions(ctx, nil)

		if err != nil {
			logger.Error("error GetTransaction: %v", err.Error())
			return "", errMarshal
		}

		response, err := marshaler.Marshal(resp)
		if err != nil {
			logger.Error("error marshaling response payload: %v", err.Error())
			return "", errMarshal
		}

		return string(response), nil

	}
}

func rpcGetTransactionInfo(marshaler *protojson.MarshalOptions, _ *protojson.UnmarshalOptions) nakamaRpcFunc {
	return func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

		resp, err := server.GetTransactionInfo(ctx, nil)

		if err != nil {
			logger.Error("error GetTransactionInfo: %v", err.Error())
			return "", errMarshal
		}

		response, err := marshaler.Marshal(resp)
		if err != nil {
			logger.Error("error marshaling response payload: %v", err.Error())
			return "", errMarshal
		}

		return string(response), nil

	}
}
