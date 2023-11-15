package helpers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

func SmartQueryString(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, contractAddr, queryMsg string, res interface{}) error {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(queryMsg), &jsonMap); err != nil {
		t.Fatal(err)
	}
	err := chain.QueryContract(ctx, contractAddr, jsonMap, &res)
	return err
}

func SetupContract(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, keyname string, fileLoc string, message string) (codeId, contract string) {
	codeId, err := chain.StoreContract(ctx, keyname, fileLoc)
	if err != nil {
		t.Fatal(err)
	}

	contractAddr, err := chain.InstantiateContract(ctx, keyname, codeId, message, true)
	if err != nil {
		t.Fatal(err)
	}

	return codeId, contractAddr
}

func ExecuteMsgWithAmount(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, contractAddr, amount, message string) {
	// amount is #utoken

	// There has to be a way to do this in ictest?
	cmd := []string{
		"junod", "tx", "wasm", "execute", contractAddr, message,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", user.KeyName(),
		"--gas", "500000",
		"--amount", amount,
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}
	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	debugOutput(t, string(stdout))

	if err := testutil.WaitForBlocks(ctx, 2, chain); err != nil {
		t.Fatal(err)
	}
}

func ExecuteMsgWithFee(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, user ibc.Wallet, contractAddr, amount, feeCoin, message string) {
	// amount is #utoken

	// There has to be a way to do this in ictest?
	cmd := []string{
		"junod", "tx", "wasm", "execute", contractAddr, message,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", user.KeyName(),
		"--gas", "500000",
		"--fees", feeCoin,
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}

	if amount != "" {
		cmd = append(cmd, "--amount", amount)
	}

	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	debugOutput(t, string(stdout))

	if err := testutil.WaitForBlocks(ctx, 2, chain); err != nil {
		t.Fatal(err)
	}
}

func ExecuteAuthzGrantMsgWithFee(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, granter ibc.Wallet, grantee ibc.Wallet, contractAddr, amount, feeCoin, message string) {
	cmd := []string{
		"junod", "tx", "authz", "grant", grantee.FormattedAddress(), "generic",
		"--msg-type", "/cosmos.authz.v1beta1.MsgExec",
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", granter.KeyName(),
		"--gas", "500000",
		"--fees", feeCoin,
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}

	if amount != "" {
		cmd = append(cmd, "--amount", amount)
	}

	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	debugOutput(t, string(stdout))

	if err := testutil.WaitForBlocks(ctx, 2, chain); err != nil {
		t.Fatal(err)
	}
}

func ExecuteAuthzExecMsgWithFee(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, grantee ibc.Wallet, contractAddr, amount, feeCoin, message string) {
	// Get the node to execute the command & write output to file
	node := chain.Nodes()[0]
	filePath := "authz.json"
	generateMsg := []string{
		"junod", "tx", "wasm", "execute", contractAddr, message,
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", grantee.KeyName(),
		"--gas", "500000",
		"--fees", feeCoin,
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"--generate-only",
	}

	// Generate msg output
	res, resErr, err := node.Exec(ctx, generateMsg, nil)
	if resErr != nil {
		t.Fatal(resErr)
	}
	if err != nil {
		t.Fatal(err)
	}

	// Write output to file
	err = node.WriteFile(ctx, res, filePath)
	if err != nil {
		t.Fatal(err)
	}

	// Execute the command
	cmd := []string{
		"junod", "tx", "authz", "exec", node.HomeDir() + "/" + filePath,
		"--node", chain.GetRPCAddress(),
		"--home", chain.HomeDir(),
		"--chain-id", chain.Config().ChainID,
		"--from", grantee.KeyName(),
		"--gas", "500000",
		"--fees", feeCoin,
		"--keyring-dir", chain.HomeDir(),
		"--keyring-backend", keyring.BackendTest,
		"-y",
	}

	if amount != "" {
		cmd = append(cmd, "--amount", amount)
	}

	stdout, _, err := chain.Exec(ctx, cmd, nil)
	require.NoError(t, err)

	debugOutput(t, string(stdout))

	if err := testutil.WaitForBlocks(ctx, 2, chain); err != nil {
		t.Fatal(err)
	}
}
