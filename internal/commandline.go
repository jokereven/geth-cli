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
