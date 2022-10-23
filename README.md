# geth-cli

## reference
https://github.com/dukedaily/go-bitcoin-demo

https://github.com/go-programming-tour-book/tour

#### 帮助

```
./geth-cli -h
```

[![asciicast](https://asciinema.org/a/gRegeEiD9h2ueZevenIx83udx.svg)](https://asciinema.org/a/gRegeEiD9h2ueZevenIx83udx)

#### 生成钱包

```
./geth-cli createWallet new
```

[![asciicast](https://asciinema.org/a/Wf2sUMNZYJ5mGa7GvhzJyYfyY.svg)](https://asciinema.org/a/Wf2sUMNZYJ5mGa7GvhzJyYfyY)

#### 创建区块链

```
 ./geth-cli createWallet new
 ./geth-cli create address -a 1JNDCBtxwUc4aHAZgc7Ld2Qaq7CWmzUwsG
```

[![asciicast](https://asciinema.org/a/B8ewzjvBd4HDUrsTYsEPTkTMx.svg)](https://asciinema.org/a/B8ewzjvBd4HDUrsTYsEPTkTMx)

#### 打印区块链

```
./geth-cli print
```

[![asciicast](https://asciinema.org/a/e8PbKenxzqOo5NqoCQoUNyjaU.svg)](https://asciinema.org/a/e8PbKenxzqOo5NqoCQoUNyjaU)

#### 打印区块链所有交易

```
./geth-cli printTx
```

[![asciicast](https://asciinema.org/a/MSAA6xvKHmdi3kzjbFx71Z70h.svg)](https://asciinema.org/a/MSAA6xvKHmdi3kzjbFx71Z70h)

#### 列出所有钱包地址

```
./geth-cli listAddress
```

[![asciicast](https://asciinema.org/a/257bygs808t2FEb24OOlTBaC0.svg)](https://asciinema.org/a/257bygs808t2FEb24OOlTBaC0)

#### 获取余额

```
./geth-cli createWallet new
./geth-cli getBalance -a 19VKMWZxvVVsMUCNZhksBFQEa8poaKRJ8N
```

[![asciicast](https://asciinema.org/a/AQrWL2cFqPvvnWZoiecjjZMGD.svg)](https://asciinema.org/a/AQrWL2cFqPvvnWZoiecjjZMGD)

#### 添加区块

```
./geth-cli addBlock
```

[![asciicast](https://asciinema.org/a/toX6DG3lA2WC4IrFFifF12At6.svg)](https://asciinema.org/a/toX6DG3lA2WC4IrFFifF12At6)

#### 转账

```
geth-cli git:main # ./geth-cli send -f 19VKMWZxvVVsMUCNZhksBFQEa8poaKRJ8N -t 1NvMYVVTLD72Tz4Kn6vM51KtU5Q3DX1EB3 -a 5 -m 1B7hh26KJ9Rshi1L44yCSdD1yAzKHepja9 -s "send 5 btc"
找到付款人的私钥和公钥，准备创建交易...
outputIndex: 0
发现挖矿交易，无需遍历inputs
金额不足，创建交易失败!
注意，找到一笔无效的转账交易, 不添加到区块!
添加区块前，对交易进行校验...
verifyTransaction开始...
发现挖矿交易，无需校验!
当前交易校验成功:688890352f2904381b70e7241223c3afcbf652cc9a903965ebefbaa6dad5cf4b
merkleRoot:8f6b7fd0901e5fc3c4497b92ff4e198315b0a95451be6b35be9ab8d7882a61a3
开始挖矿...
挖矿成功,hash :0000523c46d6b94b23400c8777edb4cdf5df18561a5cf04fa069bd6b6661d6fd, nonce :23748
添加区块成功，转账成功!
```
