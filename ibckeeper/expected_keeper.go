package ibckeeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	ibcexported "github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
)

type (
	// ChannelKeeper defines the expected IBC channel keeper
	ChannelKeeper interface {
		GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool)
		GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
		SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet ibcexported.PacketI) error
		ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
	}

	// PortKeeper defines the expected IBC port keeper
	PortKeeper interface {
		BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
	}

	// ScopedKeeper defines the expected IBC scoped keeper
	ScopedKeeper interface {
		GetCapability(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool)
		AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool
		ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error
	}
)
