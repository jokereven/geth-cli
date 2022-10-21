package internal

import "fmt"

func CreateWallet() {
	wm := NewWalletManager()
	if wm == nil {
		fmt.Println("createWallet失败!")
		return
	}
	address := wm.createWallet()

	if len(address) == 0 {
		fmt.Println("创建钱包失败！")
		return
	}

	fmt.Println("新钱包地址为:", address)
}

func TocreateBlockChain(address string) {
	if !isValidAddress(address) {
		fmt.Println("传入地址无效，无效地址为:", address)
		return
	}

	err := CreateBlockChain(address)
	if err != nil {
		fmt.Println("CreateBlockChain failed:", err)
		return
	}
	fmt.Println("执行完毕!")
}

func Print() {
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("print err:", err)
		return
	}

	defer bc.db.Close()

	//调用迭代器，输出blockChain
	it := bc.NewIterator()
	for {
		//调用Next方法，获取区块，游标左移
		block := it.Next()

		fmt.Printf("\n++++++++++++++++++++++\n")
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PrevHash : %x\n", block.PrevHash)
		fmt.Printf("MerkleRoot : %x\n", block.MerkleRoot)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Bits : %d\n", block.Bits)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Hash : %x\n", block.Hash)
		// fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Data : %s\n", block.Transactions[0].TXInputs[0].ScriptSig) //矿工写入的数据

		pow := NewProofOfWork(block)
		fmt.Printf("IsValid: %v\n", pow.IsValid())

		//退出条件
		if block.PrevHash == nil {
			fmt.Println("区块链遍历结束!")
			break
		}
	}

}

func PrintTx() {
	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("getBalance err:", err)
		return
	}

	defer bc.db.Close()

	it := bc.NewIterator()
	for {
		block := it.Next()
		fmt.Println("\n+++++++++++++++++ 区块分割 +++++++++++++++")

		for _, tx := range block.Transactions {
			//直接打印交易
			fmt.Println(tx)
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func ListAddress() {
	wm := NewWalletManager()
	if wm == nil {
		fmt.Println(" NewWalletManager 失败!")
		return
	}

	addresses := wm.listAddresses()
	for _, address := range addresses {
		fmt.Printf("%s\n", address)
	}
}

func GetBalance(address string) {
	if !isValidAddress(address) {
		fmt.Println("传入地址无效，无效地址为:", address)
		return
	}

	bc, err := GetBlockChainInstance()
	if err != nil {
		fmt.Println("getBalance err:", err)
		return
	}

	defer bc.db.Close()

	//得到查询地址的公钥哈希
	pubKeyHash := getPubKeyHashFromAddress(address)

	//获取所有相关的utxo集合
	utxoinfos := bc.FindMyUTXO(pubKeyHash)
	total := 0.0

	for _, utxo := range utxoinfos {

		//这两种方式都可以使用
		// total += utxo.Value
		total += utxo.TXOutput.Value
	}

	fmt.Printf("'%s'的金额为:%f\n", address, total)
}

func AddBlock(data string) {
	// TODO fix AddBlock
	// fmt.Println("添加区块被调用!")
	// bc, _ := GetBlockChainInstance()
	// err := bc.AddBlock(data)
	// if err != nil {
	// 	fmt.Println("AddBlock failed:", err)
	// 	return
	// }
	// fmt.Println("添加区块成功!")
}
