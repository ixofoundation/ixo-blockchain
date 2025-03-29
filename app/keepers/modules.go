package keepers

// UNFORKING v2 TODO: Eventually should get rid of this in favor of NewBasicManagerFromManager
// Right now is strictly used for default genesis creation and registering codecs prior to app init
// Unclear to me how to use NewBasicManagerFromManager for this purpose though prior to app init
import (
	"cosmossdk.io/x/evidence"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	packetforward "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v8"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	"github.com/cosmos/ibc-go/modules/capability"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	transfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	tendermint "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	"github.com/ixofoundation/ixo-blockchain/v5/x/bonds"
	"github.com/ixofoundation/ixo-blockchain/v5/x/claims"
	"github.com/ixofoundation/ixo-blockchain/v5/x/entity"
	entityclient "github.com/ixofoundation/ixo-blockchain/v5/x/entity/client"
	"github.com/ixofoundation/ixo-blockchain/v5/x/epochs"
	"github.com/ixofoundation/ixo-blockchain/v5/x/iid"
	"github.com/ixofoundation/ixo-blockchain/v5/x/liquidstake"
	"github.com/ixofoundation/ixo-blockchain/v5/x/mint"
	smartaccount "github.com/ixofoundation/ixo-blockchain/v5/x/smart-account"
	"github.com/ixofoundation/ixo-blockchain/v5/x/token"
	tokenclient "github.com/ixofoundation/ixo-blockchain/v5/x/token/client"
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = module.NewBasicManager(
	auth.AppModuleBasic{},
	genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(
		[]govclient.ProposalHandler{
			paramsclient.ProposalHandler,
			entityclient.ProposalHandler,
			tokenclient.ProposalHandler,
		},
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	consensus.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	wasm.AppModuleBasic{},
	icq.AppModuleBasic{},
	ica.AppModuleBasic{},
	packetforward.AppModuleBasic{},
	tendermint.AppModuleBasic{},
	feegrantmodule.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	ibcfee.AppModuleBasic{},
	ibchooks.AppModuleBasic{},

	iid.AppModuleBasic{},
	bonds.AppModuleBasic{},
	entity.AppModuleBasic{},
	token.AppModuleBasic{},
	claims.AppModuleBasic{},
	smartaccount.AppModuleBasic{},
	epochs.AppModuleBasic{},
	liquidstake.AppModuleBasic{},
)
