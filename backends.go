package apollo

import (
	"fmt"

	"github.com/Salvionied/apollo/txBuilding/Backend/BlockFrostChainContext"
	"github.com/Salvionied/apollo/txBuilding/Backend/FixedChainContext"
)

/**
	NewEmptyBackend creates and returns an empty FixedChainContext instance,
	which is iused for cases where no specific backend context is required.

	Returns:
		FixedChainContext.FixedChainContext: An empty FixedChainContext instance.
*/
func NewEmptyBackend() FixedChainContext.FixedChainContext {
	return FixedChainContext.InitFixedChainContext()
}

/**
	NewBlockfrostBackend creates a BlockFrostChainContext instance based
	on the specified network and project ID.

	Params:
		projectId (string): The project ID to authenticate with BlockFrost.
		network (Network): The network to configure the BlockFrost context for.

	Returns:
		BlockFrostChainContext.BlockFrostChainContext: A BlockFrostChainContext instance configured for the specified network.
*/
func NewBlockfrostBackend(
	projectId string,
	network Network,

) (BlockFrostChainContext.BlockFrostChainContext, error) {
	switch network {
	case MAINNET:
		return BlockFrostChainContext.NewBlockfrostChainContext(
			BLOCKFROST_BASE_URL_MAINNET,
			int(MAINNET),
			projectId,
		), nil
	case TESTNET:
		return BlockFrostChainContext.NewBlockfrostChainContext(
			BLOCKFROST_BASE_URL_TESTNET,
			int(TESTNET),
			projectId,
		), nil
	case PREVIEW:
		return BlockFrostChainContext.NewBlockfrostChainContext(
			BLOCKFROST_BASE_URL_PREVIEW,
			int(TESTNET),
			projectId,
		), nil
	case PREPROD:
		return BlockFrostChainContext.NewBlockfrostChainContext(
			BLOCKFROST_BASE_URL_PREPROD,
			int(TESTNET),
			projectId,
		), nil
	default:
		return BlockFrostChainContext.BlockFrostChainContext{}, fmt.Errorf("Invalid network")
	}
}
