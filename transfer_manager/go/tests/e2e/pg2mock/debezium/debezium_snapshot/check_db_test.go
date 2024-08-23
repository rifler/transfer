package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/doublecloud/transfer/library/go/test/yatest"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract"
	server "github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract/model"
	debeziumcommon "github.com/doublecloud/transfer/transfer_manager/go/pkg/debezium/common"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/debezium/testutil"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/providers/postgres/pgrecipe"
	"github.com/doublecloud/transfer/transfer_manager/go/tests/helpers"
	"github.com/stretchr/testify/require"
)

var (
	Source = *pgrecipe.RecipeSource(pgrecipe.WithPrefix(""), pgrecipe.WithInitDir("init_source"))
)

func init() {
	_ = os.Setenv("YC", "1") // to not go to vanga
	Source.WithDefaults()
}

//---------------------------------------------------------------------------------------------------------------------

func TestSnapshot(t *testing.T) {
	defer require.NoError(t, helpers.CheckConnections(
		helpers.LabeledPort{Label: "PG source", Port: Source.Port},
	))

	canonizedDebeziumKeyBytes, err := ioutil.ReadFile(yatest.SourcePath("transfer_manager/go/tests/e2e/pg2mock/debezium/debezium_snapshot/testdata/change_item_key.txt"))
	require.NoError(t, err)
	canonizedDebeziumValBytes, err := ioutil.ReadFile(yatest.SourcePath("transfer_manager/go/tests/e2e/pg2mock/debezium/debezium_snapshot/testdata/change_item_val.txt"))
	require.NoError(t, err)
	canonizedDebeziumVal := string(canonizedDebeziumValBytes)

	//------------------------------------------------------------------------------

	sinker := &helpers.MockSink{}
	target := server.MockDestination{
		SinkerFactory: func() abstract.Sinker { return sinker },
		Cleanup:       server.DisabledCleanup,
	}
	transfer := helpers.MakeTransfer("fake", &Source, &target, abstract.TransferTypeSnapshotOnly)

	var changeItems []abstract.ChangeItem
	sinker.PushCallback = func(input []abstract.ChangeItem) {
		changeItems = append(changeItems, input...)
	}

	helpers.Activate(t, transfer)

	require.Equal(t, 5, len(changeItems))
	require.Equal(t, changeItems[0].Kind, abstract.InitShardedTableLoad)
	require.Equal(t, changeItems[1].Kind, abstract.InitTableLoad)
	require.Equal(t, changeItems[2].Kind, abstract.InsertKind)
	require.Equal(t, changeItems[3].Kind, abstract.DoneTableLoad)
	require.Equal(t, changeItems[4].Kind, abstract.DoneShardedTableLoad)

	fmt.Printf("changeItem dump: %s\n", changeItems[2].ToJSONString())

	testutil.CheckCanonizedDebeziumEvent(t, &changeItems[2], "fullfillment", "pguser", "pg", true, []debeziumcommon.KeyValue{{DebeziumKey: string(canonizedDebeziumKeyBytes), DebeziumVal: &canonizedDebeziumVal}})

	changeItemBuf, err := json.Marshal(changeItems[2])
	require.NoError(t, err)
	changeItemDeserialized := helpers.UnmarshalChangeItem(t, changeItemBuf)
	testutil.CheckCanonizedDebeziumEvent(t, changeItemDeserialized, "fullfillment", "pguser", "pg", true, []debeziumcommon.KeyValue{{DebeziumKey: string(canonizedDebeziumKeyBytes), DebeziumVal: &canonizedDebeziumVal}})
}