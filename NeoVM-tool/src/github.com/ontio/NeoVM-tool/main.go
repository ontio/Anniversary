/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"time"

	"github.com/ontio/ontology/account"
	"github.com/ontio/ontology/cmd"
	"github.com/ontio/ontology/cmd/utils"
	"github.com/ontio/ontology/common/config"
	"github.com/ontio/ontology/common/log"
	"github.com/ontio/ontology/core/payload"
	"github.com/ontio/ontology/core/store/ledgerstore"
	"github.com/ontio/ontology/core/store/overlaydb"
	"github.com/ontio/ontology/core/types"
	common2 "github.com/ontio/ontology/http/base/common"
	"github.com/ontio/ontology/smartcontract"
	"github.com/ontio/ontology/smartcontract/common"
	"github.com/ontio/ontology/smartcontract/event"
	"github.com/ontio/ontology/smartcontract/storage"
	"github.com/urfave/cli"
	"math"
	"reflect"
)

const (
	DEFAULT_BYTECODE = "./test.avm.str"
	DEFAULT_TESTCASE = "./testcases.txt"
	DEFAULT_WALLET = "./wallet.dat"
	DEFAULT_LEDGER_PATH = "./Chain"
)

var (
	//Ontology setting
	NvmByteCodeFlag = cli.StringFlag{
		Name:  "bytecode,b",
		Usage: "smart contract bytecode.",
		Value: DEFAULT_BYTECODE,
	}
	TestCasesFlag = cli.StringFlag{
		Name:  "testcases,t",
		Usage: "test cases",
		Value: DEFAULT_TESTCASE,
	}
	//WalletFlag = cli.StringFlag{
	//	Name:  "wallet,w",
	//	Usage: "wallet file",
	//	Value: DEFAULT_WALLET,
	//}
	LedgerPathFlag = cli.StringFlag{
		Name:  "ledger,l",
		Usage: "ledger path",
		Value: DEFAULT_LEDGER_PATH,
	}
)

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "NeoVM CLI"
	app.Action = neovmCLI
	app.Version = config.Version
	app.Copyright = "Copyright in 2018 The Ontology Authors"
	app.Flags = []cli.Flag{
		NvmByteCodeFlag,
		TestCasesFlag,
		LedgerPathFlag,
		//WalletFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		cmd.PrintErrorMsg(err.Error())
		os.Exit(1)
	}
}

func neovmCLI(ctx *cli.Context) {
	// create account
	owner := account.NewAccount("")

	// read nvm bytecode
	codeFile := ctx.String(utils.GetFlagName(NvmByteCodeFlag))
	code, err := ioutil.ReadFile(codeFile)
	if err != nil {
		log.Errorf("open nvm code file failed: %s", err)
		return
	}

	gaslimit := uint64(100000000)
	mtx := utils.NewDeployCodeTransaction(0, gaslimit, code, true, "test", "test", "test", "test", "test")
	d := mtx.Payload.(*payload.DeployCode)
	if d == nil {
		log.Errorf("failed to get smart contract deploy address")
		return
	}
	contractAddr := d.Address()

	// init ledger
	ledgerDir := ctx.String(utils.GetFlagName(LedgerPathFlag))
	dbPath := fmt.Sprintf("%s%s%s", ledgerDir, string(os.PathSeparator), ledgerstore.DBDirState)
	merklePath := fmt.Sprintf("%s%s%s", ledgerDir, string(os.PathSeparator), ledgerstore.MerkleTreeStorePath)
	stateStore, err := ledgerstore.NewStateStore(dbPath, merklePath, 0)
	if err != nil {
		log.Errorf("failed to create state store: %s", err)
		return
	}
	overlay := stateStore.NewOverlayDB()

	// deploy nvm byte code
	if err := executeDeployTx(stateStore, overlay, owner, mtx); err != nil {
		log.Errorf("failed to deploy smart contract: %s", err)
		return
	}
	log.Infof("deploy done, address = %s", contractAddr.ToHexString())

	// load testcases
	testcaseFile := ctx.String(utils.GetFlagName(TestCasesFlag))
	f, err := os.Open(testcaseFile)
	if err != nil {
		log.Errorf("failed to read testcase file %s: %s", testcaseFile, err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(bufio.NewReader(f))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Errorf("read testcase file: %s", err)
			return
		}

		if len(line) != 4 {
			log.Errorf("bad testcase: %v", line)
			continue
		}

		testcase_id := line[0]
		result := line[3]
		//params := []interface{}{line[1], []interface{}{line[2]}}
		params := []interface{}{"init"}
		mtx, err := common2.NewNeovmInvokeTransaction(0, gaslimit, contractAddr, params)
		if err != nil {
			log.Errorf("create tx for testcase %s failed: %s", testcase_id, err)
			break
		}
		testResult, err := executeInvokeTx(stateStore, overlay, owner, mtx)
		if err != nil {
			log.Errorf("process testcase %s failed: %s", testcase_id, err)
		}
		if testResult != result {
			log.Errorf("testcase %s failed: %s vs %s", testcase_id, testResult, result)
		}
	}
}

func executeDeployTx(store *ledgerstore.StateStore, overlay *overlaydb.OverlayDB, user *account.Account, mtx *types.MutableTransaction) error {
	cache := storage.NewCacheDB(overlay)

	if err := utils.SignTransaction(user, mtx); err != nil {
		return fmt.Errorf("sign deploy: %s", err)
	}
	tx, err := mtx.IntoImmutable()
	if err != nil {
		return fmt.Errorf("deploy tx immu: %s", err)
	}

	notify := &event.ExecuteNotify{TxHash: tx.Hash(), State: event.CONTRACT_STATE_FAIL}
	if err := store.HandleDeployTransaction(nil, overlay, cache, tx, nil, notify); err != nil {
		return fmt.Errorf("handle deploy tx: %s", err)
	}
	return nil
}

func executeInvokeTx(store *ledgerstore.StateStore, overlay *overlaydb.OverlayDB, user *account.Account, mtx *types.MutableTransaction) (string, error) {
	if err := utils.SignTransaction(user, mtx); err != nil {
		return "", fmt.Errorf("failed to sign tx: %s", err)
	}
	tx, err := mtx.IntoImmutable()
	if err != nil {
		return "", fmt.Errorf("failed to invoke tx immu: %s", err)
	}

	cache := storage.NewCacheDB(overlay)
	config := &smartcontract.Config{
		Time:   uint32(time.Now().Unix()),
		Height: 1000,
		Tx:     tx,
	}
	if tx.TxType != types.Invoke {
		return "", fmt.Errorf("preexec Tx is not Invoke")
	}
	invoke := tx.Payload.(*payload.InvokeCode)

	sc := smartcontract.SmartContract{
		Config:  config,
		Store:   nil,
		CacheDB: cache,
		Gas:     math.MaxUint64,
		PreExec: true,
	}

	//start the smart contract executive function
	engine, _ := sc.NewExecuteEngine(invoke.Code)
	result, err := engine.Invoke()
	if err != nil {
		return "", fmt.Errorf("preexec invoke failed: %s", err)
	}
	log.Infof("tx result: %v, %v, gas %v \n", result, reflect.TypeOf(result), math.MaxUint64 - sc.Gas)
	for _, n := range sc.Notifications {
		log.Infof(" %v : %v", n.ContractAddress, n.States)
	}
	cv, err := common.ConvertNeoVmTypeHexString(result)
	if err != nil {
		return "", fmt.Errorf("preexec invoke failed to convert result")
	}
	return cv.(string), nil
}
