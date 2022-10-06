package keeper

import (
	"fmt"

	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

func (suite *KeeperTestSuite) TestDidDocumentKeeperSetAndGet() {
	testCases := []struct {
		msg     string
		didFn   func() types.DidDocument
		expPass bool
	}{
		{
			"iid stored successfully",
			func() types.DidDocument {
				dd, _ := types.NewDidDocument("did:cash:subject")
				return dd
			},
			true,
		},
	}
	for _, tc := range testCases {
		dd := tc.didFn()

		suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id), dd)
		suite.keeper.SetDidDocument(suite.ctx, []byte(dd.Id+"1"), dd)
		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			if tc.expPass {
				_, found := suite.keeper.GetDidDocument(
					suite.ctx,
					[]byte(dd.Id),
				)
				suite.Require().True(found)

				allEntities := suite.keeper.GetAllDidDocuments(
					suite.ctx,
				)
				suite.Require().Equal(2, len(allEntities))
			} else {
				// TODO write failure cases
				suite.Require().False(tc.expPass)
			}
		})
	}
}
