package wicc_wallet_utils_go

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"
)


func TestEthCreateSimpleRawTransaction(t *testing.T) {

	privateKeyStr := "6B93D965D9981F9066CCC44B9DBF32B50F411C0DCEDF4A41CA4E7424ABDB617F"
	from := "0x232D23C22543144B988F738C701Df6dfd6eAcA4c"
	to := "81FD1F7aE91041aAc5fCF7d8Ed3e1dd88Cc1359a"
	amount := "100000000000000000"
	//amount := "1000000000000000000000000"
	amountBig,_ := ConvertToBigInt(amount,10)

	txcount ,err:= tw.WalletClient.ethGetTransactionCount(from,LEATEST)
	if err != nil {
		t.Errorf("Failed to ethGetTransactionCount: %v",err)
	}
	t.Log("txcount=",txcount)

	txFeeInfo, err := tw.WalletClient.GetTransactionFeeEstimated(from,to,amountBig,"")
	if err != nil {
		t.Errorf("Failed to GetTransactionFeeEstimated: %v",err)
	}
	t.Logf("txFeeInfo.GasLimit=%v, txFeeInfo.GasPrice=%v\n",txFeeInfo.GasLimit,txFeeInfo.GasPrice)


	tx,err := NewETHSimpleTransaction(to,amount,strconv.FormatInt(int64(txcount),10),txFeeInfo.GasLimit.String(),txFeeInfo.GasPrice.String())
	if err != nil {
		t.Errorf("Failed to NewETHSimpleTransaction: %v",err)
	}
	rawtx ,err :=tx.CreateRawTx(privateKeyStr,tw.Config.ChainID)
	if err != nil{
		fmt.Println("err=",err)
	}
	fmt.Println("rawtx=",rawtx)

}


func TestERC20CreateTransferRawTransaction(t *testing.T) {

/*	privateKeyStr := "6B93D965D9981F9066CCC44B9DBF32B50F411C0DCEDF4A41CA4E7424ABDB617F"
	from := "0x232D23C22543144B988F738C701Df6dfd6eAcA4c"
	to := "81FD1F7aE91041aAc5fCF7d8Ed3e1dd88Cc1359a"*/

/*	privateKeyStr := "2F0ED572654E79084565F19A8C8C9ECD36E67C2DE42017A0CF4D951CA9EC1FC4"
	from := "0x96b4213eD85031b02A1bE101FfA3F82ee929285E"
	to := "E63600aDf44c5B3562Ad2273CBaf15433d339193"*/

	privateKeyStr := "798b289e6115daed769b21fe62e43a7eb81efb98c43c759ae265fe72992f3bfd"
	from := "0x60792E51423AC671AA5b4EBA4A976349EBf0f541"
	to := "0cb24762b2698ab48a8fe8cefafea2a468980456"


	contractAddr := "0x8E1dA42EbC22F91d528ceB9865f241167Ebb8A0f"  //WICC合约
	amount := big.NewInt(10000000000)

	txcount ,err:= tw.WalletClient.ethGetTransactionCount(from,LEATEST)
	if err != nil {
		t.Errorf("Failed to ethGetTransactionCount: %v",err)
	}
	t.Log("txcount=",txcount)

	data ,err := makeERC20TokenTransferData(contractAddr,to,amount)
	//0xa9059cbb00000000000000000000000081fd1f7ae91041aac5fcf7d8ed3e1dd88cc1359a00000000000000000000000000000000000000000000000000000002540be400
	//0xa9059cbb000000000000000000000000e63600adf44c5b3562ad2273cbaf15433d33919300000000000000000000000000000000000000000000000000000002540be400

	if err != nil {
		t.Errorf("makeERC20TokenTransData, err=%v", err)
	}

/*	txFeeInfo,err := tw.WalletClient.GetTransactionFeeEstimated(from,contractAddr,big.NewInt(0),data)
	if err != nil {
		t.Errorf("Failed to GetTransactionFeeEstimated: %v",err)
	}
	t.Logf("txFeeInfo.GasLimit=%v, txFeeInfo.GasPrice=%v\n",txFeeInfo.GasLimit,txFeeInfo.GasPrice)*/

	//tx,err := NewERC20TransferTransaction(contractAddr,"0",strconv.FormatInt(int64(txcount),10),txFeeInfo.GasLimit.String(),txFeeInfo.GasPrice.String(),data)
	tx,err := NewERC20TransferTransaction(contractAddr,"0",strconv.FormatInt(int64(txcount),10),"200000","2000000000",data)
	if err != nil {
		t.Errorf("Failed to NewERC20TransferTransaction: %v",err)
	}
	rawtx ,err :=tx.CreateRawTx(privateKeyStr,tw.Config.ChainID)
	if err != nil{
		fmt.Println("err=",err)
	}
	fmt.Println("rawtx=",rawtx)
}

func TestAAACreateTransferRawTransaction(t *testing.T) {

	/*	privateKeyStr := "6B93D965D9981F9066CCC44B9DBF32B50F411C0DCEDF4A41CA4E7424ABDB617F"
		from := "0x232D23C22543144B988F738C701Df6dfd6eAcA4c"
		to := "81FD1F7aE91041aAc5fCF7d8Ed3e1dd88Cc1359a"*/

	/*	privateKeyStr := "2F0ED572654E79084565F19A8C8C9ECD36E67C2DE42017A0CF4D951CA9EC1FC4"
		from := "0x96b4213eD85031b02A1bE101FfA3F82ee929285E"
		to := "E63600aDf44c5B3562Ad2273CBaf15433d339193"*/

	privateKeyStr := "798b289e6115daed769b21fe62e43a7eb81efb98c43c759ae265fe72992f3bfd"
	from := "0x60792E51423AC671AA5b4EBA4A976349EBf0f541"
	to := "0cb24762b2698ab48a8fe8cefafea2a468980456"


	contractAddr := "0xdAC17F958D2ee523a2206206994597C13D831ec7"  //USDT合约
	amount := big.NewInt(710000000000)

	txcount ,err:= tw.WalletClient.ethGetTransactionCount(from,LEATEST)
	if err != nil {
		t.Errorf("Failed to ethGetTransactionCount: %v",err)
	}
	t.Log("txcount=",txcount)

	data ,err := makeERC20TokenTransferData(contractAddr,to,amount)
	//0xa9059cbb00000000000000000000000081fd1f7ae91041aac5fcf7d8ed3e1dd88cc1359a00000000000000000000000000000000000000000000000000000002540be400
	//0xa9059cbb000000000000000000000000e63600adf44c5b3562ad2273cbaf15433d33919300000000000000000000000000000000000000000000000000000002540be400

	if err != nil {
		t.Errorf("makeERC20TokenTransData, err=%v", err)
	}

	/*	txFeeInfo,err := tw.WalletClient.GetTransactionFeeEstimated(from,contractAddr,big.NewInt(0),data)
		if err != nil {
			t.Errorf("Failed to GetTransactionFeeEstimated: %v",err)
		}
		t.Logf("txFeeInfo.GasLimit=%v, txFeeInfo.GasPrice=%v\n",txFeeInfo.GasLimit,txFeeInfo.GasPrice)*/

	//tx,err := NewERC20TransferTransaction(contractAddr,"0",strconv.FormatInt(int64(txcount),10),txFeeInfo.GasLimit.String(),txFeeInfo.GasPrice.String(),data)
	tx,err := NewERC20TransferTransaction(contractAddr,"0",strconv.FormatInt(int64(txcount),10),"200000","42000000000",data)
	if err != nil {
		t.Errorf("Failed to NewERC20TransferTransaction: %v",err)
	}
	rawtx ,err :=tx.CreateRawTx(privateKeyStr,tw.Config.ChainID)
	if err != nil{
		fmt.Println("err=",err)
	}
	fmt.Println("rawtx=",rawtx)
}

func TestEthSendRawTransaction(t *testing.T) {

	//rawtx := "f86b06843b9aca008252089481fd1f7ae91041aac5fcf7d8ed3e1dd88cc1359a88016345785d8a00008029a018adef9b4ec654de5ecb7976b52faf6e5ddfc22902dffa8ab5f5eabd07f4c1eca020d35fc7f7b58ee5e53a7e3ea4849eba1b26a7748db373fbb4f48b7f30e5e4ea"
	//rawtx := "f86308843b9aca0082d0ed948e1da42ebc22f91d528ceb9865f241167ebb8a0f808029a034fce598f3abe285c12d53f1bae5b790b598e168b33b7f1ca0e55afb0652837fa06c4a00544fbe49be8bcc321a7bd271c775113d79a1550f583d0887df899b20bd"
	//rawtx := "f8a809843b9aca0082d0ed948ne1da42ebc22f91d528ceb9865f241167ebb8a0f80b844a9059cbb00000000000000000000000081fd1f7ae91041aac5fcf7d8ed3e1dd88cc1359a0000000000000000000000000000000000000000000000000000000005f5e10029a01bef8dd99141dd4cb6461f901e4c03b1df28c6422dc2ee5478ae2b73cb2610d2a032ed60e4137b9619d41930fef0193205fa4d57a174bf0014bb1a246e36cb4480"
	//rawtx := "f86b0b843b9aca008252089481fd1f7ae91041aac5fcf7d8ed3e1dd88cc1359a88016345785d8a00008029a0fc16ac404d694561f44e12e9f1ab4916244c352add8836ecc05ae80ef9f7e3e9a05a2c6e6b6fc65373700c6fbf003d055a38c49d626b370b71eec1809978f027fe"
	//rawtx := "f86c37830186a08252289481fd1f7ae91041aac5fcf7d8ed3e1dd88cc1359a88016345785d8a00008212342aa029c94192b56a5a194bbe9575bd14d385546943b934d5db73b848a8bcd65860f3a017c9c625e79b8c1d585ebfc64d3e5b48bbf34a6d8aa901c7edd462440339e2f1"
 	rawtx := "f8a90b847735940083030d40948e1da42ebc22f91d528ceb9865f241167ebb8a0f80b844a9059cbb000000000000000000000000e63600adf44c5b3562ad2273cbaf15433d33919300000000000000000000000000000000000000000000000000000002540be4002aa0c2b1c7f94018250f23376f2f0830daf6d4067f552ef7d7548f2bbe39832aac29a036e8b6d5aa46d7dc73443a178fdc2d8b19edb5f562237238c7f754b1f271ff44"
	txid ,err  := tw.WalletClient.EthSendRawTransaction(rawtx)
	if err != nil{
		fmt.Println("err=",err)
	}
	fmt.Println("txid=",txid)
}


func TestConvertToBigInt(t *testing.T) {
	value := "10000000000000000000000000000000"
	base := 10
	r,_ := ConvertToBigInt(value,base)

	fmt.Printf("r=%+v\n",r)
}



func TestDecodeETHRawTx(t *testing.T) {

	rawTXstring := "f8d48201b485239f82ba008301c2639465b49f7aee40347f5a90b714be4ef086f3fe5e2c884109caeb7bde4000b864c0f4ed310000000000000000000000000000000000000000000000003f24d8e4a0070000000000000000000000000000b932a70a57673d89f4acffbe830e8ed7f75fb9e000000000000000000000000000000000000000000000000000000000000047f625a0c634001385ac4ba20389eb2f3efb414bfdbf2cd23d3f1912ad0a11c4d848b321a04606a1aafa0619f4e8befd0582524bc52fc03d16e48423ce649da1b8e6ce4052"

	jsonIndent, err := DecodeETHRawTx(rawTXstring)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonIndent))
}

