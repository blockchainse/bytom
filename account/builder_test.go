package account

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/bytom/blockchain/txbuilder"
	"github.com/bytom/consensus"
	"github.com/bytom/crypto/ed25519/chainkd"
	"github.com/bytom/protocol/bc"
	"github.com/bytom/testutil"
)

func TestMergeSpendAction(t *testing.T) {
	testBTM := &bc.AssetID{}
	if err := testBTM.UnmarshalText([]byte("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")); err != nil {
		t.Fatal(err)
	}

	testAssetID1 := &bc.AssetID{}
	if err := testAssetID1.UnmarshalText([]byte("50ec80b6bc48073f6aa8fa045131a71213c33f3681203b15ddc2e4b81f1f4730")); err != nil {
		t.Fatal(err)
	}

	testAssetID2 := &bc.AssetID{}
	if err := testAssetID2.UnmarshalText([]byte("43c6946d092b2959c1a82e90b282c68fca63e66de289048f6acd6cea9383c79c")); err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		testActions     []txbuilder.Action
		wantActions     []txbuilder.Action
		testActionCount int
		wantActionCount int
	}{
		{
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  200,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  300,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  300,
					},
					AccountID: "test_account",
				}),
			},
			wantActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  600,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  300,
					},
					AccountID: "test_account",
				}),
			},
			testActionCount: 4,
			wantActionCount: 2,
		},
		{
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  200,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  500,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  300,
					},
					AccountID: "test_account",
				}),
			},
			wantActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  600,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  500,
					},
					AccountID: "test_account",
				}),
			},
			testActionCount: 4,
			wantActionCount: 2,
		},
		{
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  200,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID2,
						Amount:  300,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  300,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID2,
						Amount:  500,
					},
					AccountID: "test_account",
				}),
			},
			wantActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  500,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID2,
						Amount:  800,
					},
					AccountID: "test_account",
				}),
			},
			testActionCount: 5,
			wantActionCount: 3,
		},
		{
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  200,
					},
					AccountID: "test_account1",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  500,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  300,
					},
					AccountID: "test_account1",
				}),
			},
			wantActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  600,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  200,
					},
					AccountID: "test_account1",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  300,
					},
					AccountID: "test_account1",
				}),
			},
			testActionCount: 4,
			wantActionCount: 3,
		},
		{
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendUTXOAction{
					OutputID: &bc.Hash{V0: 128},
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  200,
					},
					AccountID: "test_account1",
				}),
				txbuilder.Action(&spendUTXOAction{
					OutputID: &bc.Hash{V0: 256},
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID2,
						Amount:  300,
					},
					AccountID: "test_account2",
				}),
			},
			wantActions: []txbuilder.Action{
				txbuilder.Action(&spendUTXOAction{
					OutputID: &bc.Hash{V0: 128},
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testBTM,
						Amount:  100,
					},
					AccountID: "test_account",
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID1,
						Amount:  200,
					},
					AccountID: "test_account1",
				}),
				txbuilder.Action(&spendUTXOAction{
					OutputID: &bc.Hash{V0: 256},
				}),
				txbuilder.Action(&spendAction{
					AssetAmount: bc.AssetAmount{
						AssetId: testAssetID2,
						Amount:  300,
					},
					AccountID: "test_account2",
				}),
			},
			testActionCount: 5,
			wantActionCount: 5,
		},
	}

	for _, c := range cases {
		gotActions := MergeSpendAction(c.testActions)

		gotMap := make(map[string]uint64)
		wantMap := make(map[string]uint64)
		for _, got := range gotActions {
			switch got := got.(type) {
			case *spendAction:
				gotKey := got.AssetId.String() + got.AccountID
				gotMap[gotKey] = got.Amount
			default:
				continue
			}
		}

		for _, want := range c.wantActions {
			switch want := want.(type) {
			case *spendAction:
				wantKey := want.AssetId.String() + want.AccountID
				wantMap[wantKey] = want.Amount
			default:
				continue
			}
		}

		for key := range gotMap {
			if gotMap[key] != wantMap[key] {
				t.Fatalf("gotMap[%s]=%v, wantMap[%s]=%v", key, gotMap[key], key, wantMap[key])
			}
		}

		if len(gotActions) != c.wantActionCount {
			t.Fatalf("number of gotActions=%d, wantActions=%d", len(gotActions), c.wantActionCount)
		}
	}
}
func getBlockHeight() uint64 {
	return 100
}

func mockUTXO(controlProg *CtrlProgram, assetID *bc.AssetID, outputID uint64, amount uint64) *UTXO {
	utxo := &UTXO{}
	utxo.OutputID = bc.Hash{V0: outputID}
	utxo.SourceID = bc.Hash{V0: 2}
	utxo.AssetID = *assetID
	utxo.Amount = amount
	utxo.SourcePos = 0
	utxo.ControlProgram = controlProg.ControlProgram
	utxo.AccountID = controlProg.AccountID
	utxo.Address = controlProg.Address
	utxo.ControlProgramIndex = controlProg.KeyIndex
	return utxo
}

// Test the normal build chain transaction
// Test build failed if the number of test assets is insufficient
func TestMergeSpendActionUTXO(t *testing.T) {
	m := mockAccountManager(t)
	alias := "UPPER"
	testAccount, err := m.Create([]chainkd.XPub{testutil.TestXPub}, 1, alias)

	if err != nil {
		t.Fatal(err)
	}

	testBTM := &bc.AssetID{}
	if err := testBTM.UnmarshalText([]byte("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		utxoAmount  []uint64
		testActions []txbuilder.Action
		wantAmount  [][10]uint64
		wantError   error
	}{
		{utxoAmount: []uint64{50000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 1000000000, 1000000000, 1000000000, 1000000000, 2000000000, 2000000000, 2000000000, 3000000000, 4000000000, 4000000000, 5000000000, 6000000000, 6000000000, 7000000000, 8000000000, 9000000000},
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					accounts: m,
					AssetAmount: bc.AssetAmount{
						AssetId: consensus.BTMAssetID,
						Amount:  62500000000,
					},
					AccountID: testAccount.ID,
				})},
			wantAmount: [][10]uint64{{9000000000, 8000000000, 7000000000, 6000000000, 6000000000, 5000000000, 4000000000, 4000000000, 3000000000, 2000000000}, {53990000000, 2000000000, 2000000000, 1000000000, 1000000000, 1000000000, 1000000000, 100000000, 100000000, 100000000}},
			wantError:  nil,
		},
		{utxoAmount: []uint64{50000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 1000000000, 1000000000, 1000000000, 1000000000, 2000000000, 2000000000, 2000000000, 3000000000, 4000000000, 4000000000, 5000000000, 6000000000, 6000000000, 7000000000, 8000000000, 9000000000},
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					accounts: m,
					AssetAmount: bc.AssetAmount{
						AssetId: consensus.BTMAssetID,
						Amount:  63000000000,
					},
					AccountID: testAccount.ID,
				})},
			wantError: ErrInsufficient,
		},
	}

	for _, test := range cases {
		for i, amount := range test.utxoAmount {
			controlProg, err := m.CreateAddress(testAccount.ID, false)
			if err != nil {
				t.Fatal(err)
			}
			utxo := mockUTXO(controlProg, consensus.BTMAssetID, uint64(i), amount)
			data, err := json.Marshal(utxo)
			if err != nil {
				t.Fatal(err)
			}
			m.db.Set(StandardUTXOKey(utxo.OutputID), data)
		}
		maxTime := time.Now().Add(1000000)
		tpls, err := MergeSpendActionUTXO(nil, test.testActions, maxTime, 0)
		if err != test.wantError {
			t.Fatal(err)
		}
		if err != nil {
			continue
		}
		key := actTemplatesKey(testAccount.ID, consensus.BTMAssetID)
		tpl, ok := tpls[key]
		if !ok {
			t.Fatal("tpl err")
		}

		for i, v := range tpl {
			for j, input := range v.Transaction.Inputs {
				if test.wantAmount[i][j] != input.Amount() {
					t.Fatal("tpl err")
				}
			}
		}
	}
}

//TestMergeSpendActionUTXOFailRollback Test build chained transaction failure rollback
func TestMergeSpendActionUTXOFailRollback(t *testing.T) {
	m := mockAccountManager(t)
	alias := "UPPER"
	testAccount, err := m.Create([]chainkd.XPub{testutil.TestXPub}, 1, alias)

	if err != nil {
		t.Fatal(err)
	}

	testBTM := &bc.AssetID{}
	if err := testBTM.UnmarshalText([]byte("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		utxoAmount  []uint64
		testActions []txbuilder.Action
		wantAmount  [][10]uint64
		wantError   error
	}{
		{utxoAmount: []uint64{50000000, 100000000, 100000000, 100000000, 100000000, 100000000, 100000000, 1000000000, 1000000000, 1000000000, 1000000000, 2000000000, 2000000000, 2000000000, 3000000000, 4000000000, 4000000000, 5000000000, 6000000000, 6000000000, 7000000000, 8000000000, 9000000000},
			testActions: []txbuilder.Action{
				txbuilder.Action(&spendAction{
					accounts: m,
					AssetAmount: bc.AssetAmount{
						AssetId: consensus.BTMAssetID,
						Amount:  62640000000,
					},
					AccountID: testAccount.ID,
				})},
			wantError: ErrReserved,
		},
	}

	for _, test := range cases {
		for i, amount := range test.utxoAmount {
			controlProg, err := m.CreateAddress(testAccount.ID, false)
			if err != nil {
				t.Fatal(err)
			}
			utxo := mockUTXO(controlProg, consensus.BTMAssetID, uint64(i), amount)
			data, err := json.Marshal(utxo)
			if err != nil {
				t.Fatal(err)
			}
			m.db.Set(StandardUTXOKey(utxo.OutputID), data)
		}
		maxTime := time.Now().Add(1000000000)
		_, err := MergeSpendActionUTXO(nil, test.testActions, maxTime, 0)
		if err != test.wantError {
			t.Fatal(err)
		}
		if len(m.utxoKeeper.reserved) != 0 || len(m.utxoKeeper.reservations) != 0 {
			t.Fatal("Chain transaction rollback failed")
		}
	}

}
