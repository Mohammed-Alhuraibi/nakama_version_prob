// Copyright 2020 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	moduleName = "frostline"
)

var (
	errInternalError  = runtime.NewError("internal server error", 13) // INTERNAL
	errMarshal        = runtime.NewError("cannot marshal type", 13)   // INTERNAL
	errUnmarshal      = runtime.NewError("cannot unmarshal type", 13) // INTERNAL
	errNoInputAllowed = runtime.NewError("no input allowed", 3)       // INVALID_ARGUMENT
	errNoUserIdFound  = runtime.NewError("no user ID in context", 3)  // INVALID_ARGUMENT
)

const (
	rpcName_GetUserInfo        = "GetUserInfo"
	rpcName_GetInventory       = "GetInventory"
	rpcName_GetShop            = "GetShop"
	rpcName_GetTransactions    = "GetTransactions"
	rpcName_GetTransactionInfo = "GetTransactionInfo"
)

// func SetSessionVars(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
// 	logger.Info("User session contains key-value pairs set the client: %v", in.GetAccount().Vars)

// 	if in.GetAccount().Vars == nil {
// 		in.GetAccount().Vars = map[string]string{}
// 	}
// 	in.GetAccount().Vars["firebase_uid"] = "firebase_uid"

// 	return in, nil
// }

// noinspection GoUnusedExportedFunction
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	marshaler := &protojson.MarshalOptions{
		UseEnumNumbers: true,
	}
	unmarshaler := &protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}

	if err := initializer.RegisterRpc(rpcName_GetUserInfo, rpcGetUserInfo(marshaler, unmarshaler)); err != nil {
		return err
	}

	if err := initializer.RegisterRpc(rpcName_GetInventory, rpcGetInventory(marshaler, unmarshaler)); err != nil {
		return err
	}

	if err := initializer.RegisterRpc(rpcName_GetShop, rpcGetShop(marshaler, unmarshaler)); err != nil {
		return err
	}

	if err := initializer.RegisterRpc(rpcName_GetTransactions, rpcGetTransactions(marshaler, unmarshaler)); err != nil {
		return err
	}

	if err := initializer.RegisterRpc(rpcName_GetTransactionInfo, rpcGetTransactionInfo(marshaler, unmarshaler)); err != nil {
		return err
	}

	logger.Info("Plugin loaded in '%d' msec.", time.Since(initStart).Milliseconds())
	return nil
}
