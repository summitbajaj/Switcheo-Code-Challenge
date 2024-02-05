package match

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "match/api/match/match"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "MatchInfoAll",
					Use:       "list-match-info",
					Short:     "List all match_info",
				},
				{
					RpcMethod:      "MatchInfo",
					Use:            "show-match-info [id]",
					Short:          "Shows a match_info by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateMatchInfo",
					Use:            "create-match-info [matchdate] [home] [away]",
					Short:          "Create match_info",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "matchdate"}, {ProtoField: "home"}, {ProtoField: "away"}},
				},
				{
					RpcMethod:      "UpdateMatchInfo",
					Use:            "update-match-info [id] [matchdate] [home] [away]",
					Short:          "Update match_info",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "matchdate"}, {ProtoField: "home"}, {ProtoField: "away"}},
				},
				{
					RpcMethod:      "DeleteMatchInfo",
					Use:            "delete-match-info [id]",
					Short:          "Delete match_info",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
