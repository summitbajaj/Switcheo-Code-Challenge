package match

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"match/testutil/sample"
	matchsimulation "match/x/match/simulation"
	"match/x/match/types"
)

// avoid unused import issue
var (
	_ = matchsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateMatchInfo = "op_weight_msg_match_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMatchInfo int = 100

	opWeightMsgUpdateMatchInfo = "op_weight_msg_match_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMatchInfo int = 100

	opWeightMsgDeleteMatchInfo = "op_weight_msg_match_info"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMatchInfo int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	matchGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MatchInfoList: []types.MatchInfo{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MatchInfoCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&matchGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMatchInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMatchInfo, &weightMsgCreateMatchInfo, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMatchInfo = defaultWeightMsgCreateMatchInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMatchInfo,
		matchsimulation.SimulateMsgCreateMatchInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMatchInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMatchInfo, &weightMsgUpdateMatchInfo, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMatchInfo = defaultWeightMsgUpdateMatchInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMatchInfo,
		matchsimulation.SimulateMsgUpdateMatchInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMatchInfo int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMatchInfo, &weightMsgDeleteMatchInfo, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMatchInfo = defaultWeightMsgDeleteMatchInfo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMatchInfo,
		matchsimulation.SimulateMsgDeleteMatchInfo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMatchInfo,
			defaultWeightMsgCreateMatchInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				matchsimulation.SimulateMsgCreateMatchInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMatchInfo,
			defaultWeightMsgUpdateMatchInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				matchsimulation.SimulateMsgUpdateMatchInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMatchInfo,
			defaultWeightMsgDeleteMatchInfo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				matchsimulation.SimulateMsgDeleteMatchInfo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
