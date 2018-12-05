// Copyright 2014 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// geth is the official command-line client for Ethereum.
package main

import (
	"math/big"
	"fmt"
	"time"
	
	"github.com/blockchain/imobilechain/fbclient"
	"github.com/blockchain/imobilechain/params"
	"github.com/blockchain/imobilechain/common"

)
const (
	
		// channelID      = "mychannel" //注意yaml配置文件也需要设置成这样子
		// orgName        = "Org1MSP"   //注意yaml配置文件也需要设置成这样子
		// orgAdmin       = "SampleOrg"
		// ordererOrgName = "SampleOrg"
		// ccID           = "mycc"
		path ="fabric-sdk-config.yaml"
	
)

// ExampleCC query and transaction arguments
// var addMobileMiningEventArgs1 = [][]byte{[]byte(params.DefaulMinerAddr1.String()), []byte(params.DefaulMinerAddr2.String())}
// var addMobileMiningEventArgs2 = [][]byte{[]byte(params.DefaulMinerAddr2.String()), []byte(params.DefaulMinerAddr3.String())}
// var addMobileMiningEventArgs3 = [][]byte{[]byte(params.DefaulMinerAddr3.String()), []byte(params.DefaulMinerAddr1.String())}
//var getPoolWinnersArgs = [][]byte{[]byte("0"), []byte("2536310975299272628"), []byte("3"), []byte("3")}
//var getMobileMiningArgs = [][]byte{[]byte("1535607601000000000"), []byte("1536307801000000000"),[]byte("2")}
//var registerBlockSealArgs=[][]byte{[]byte("1535607601000000000"), []byte("1536307801000000000"),[]byte("2"),[]byte("0xa000,0xb000,0xc000"),[]byte("0x0001,0x0002,0x0003"),}

func main() {
	var Fbc *fbclient.Fbclient
  	Fbc=fbclient.New(path)
	//ledgclient:=fbclient.CreateLedgerClient(fbclient.Fbclient.Fsdk)
//	curnum,_:=Fbc.CurfbNumber()
	fmt.Println("allpeers",Fbc.AllPeers,"selectpeers",Fbc.SelectPeers,"currentpeers",Fbc.CurPeers)

	// go clientMine(Fbc,1,big.NewInt(1000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(2000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(3000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(4000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(5000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(6000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(7000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(8000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(9000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(10000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(1100000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(1200000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(1400000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(1500000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(1600000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(1700000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(1800000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(1900000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(10000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(20000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(30000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(40000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(50000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(60000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(70000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(80000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(90000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(100000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(11000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(12000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(14000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(15000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(16000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(17000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(18000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(19000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(11000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(12000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(13000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(14000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(15000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(16000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(17000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(18000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(19000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(110000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(11100000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(11200000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(11400000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(11500000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(11600000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(11700000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(11800000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(11900000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(110000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(120000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(130000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(1140000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(150000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(160000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(170000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(180000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(190000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(1100000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(111000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(112000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(114000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(115000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(116000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(117000000000000000),params.MinerPool1.String()) //模拟终端一
	// go clientMine(Fbc,2,big.NewInt(118000000000000000),params.MinerPool2.String()) //模拟终端一
	// go clientMine(Fbc,1,big.NewInt(119000000000000000),params.MinerPool1.String()) //模拟终端一
	// // go clientMine(Fbc,3,params.Miner3.String(),params.MinerPool3.String()) //模拟终端一
	// // go clientMine(Fbc,4,params.Miner4.String(),params.MinerPool4.String()) //模拟终端一
	// // go clientMine(Fbc,5,params.Miner5.String(),params.MinerPool5.String()) //模拟终端一
    TestMine(Fbc)
	//CurentMineWiner(Fbc)  //区块生成模拟

}
func clientMine(fbc *fbclient.Fbclient,num int,user *big.Int,miner string){
	var i int
	i=1
	for {
		user.Add(user,big.NewInt(1))
		var addEventArgs = [][]byte{[]byte(common.BigToAddress(user).String()), []byte(miner)}
		err:=fbc.AddMobileMiningEvent(addEventArgs)
		if err==nil{
			fmt.Println("--------------------移动矿工挖矿信息---------------------：")
			fmt.Println("模拟终端：",num,"  挖矿次数",i,"  移动矿工",common.BigToAddress(user).String(),"矿池：",miner,"状态: 成功")	
		}else{
			fmt.Println("挖矿未能成功：",err)
		}
		i++	
	}
}
// func queryblock(fbc *fbclient.Fbclient){
// 	for {
// 	   blocknum:=fbc.QueryBlockInfo()
// 	   fmt.Println("blocknum",blocknum)
//        err:=fbc.QueryBlockNumber(blocknum-1)
// 	   if err!=nil{
// 		   fmt.Println(err)
// 		}
// 		time.Sleep(15*time.Second)
// 	}
// }
func TestMine(fbc *fbclient.Fbclient){
	//curnum,_:=fbclient.CurfbNumber(fbclient.Fbclient.LedgerClient,fbclient.Fbclient.AllPeers)
	//fmt.Println("当前区块号",curnum)
	prevnum,_:=fbc.CurfbNumber()
	var i=uint64(0)
	for {
		fmt.Println("------------------------资产链区块号",i,"-------------------------：")
		curnum,_:=fbc.CurfbNumber()
		Miners,_,_:=fbc.CandidateMiners(prevnum,curnum,i)
		duration:=fbc.WaitTime(Miners,params.MinerPool2.String(),uint64(i))   //计算等待时
		fmt.Println("------------------------候选矿池账号-------------------------：",curnum)
		fmt.Println("------------------------等待时间-------------------------：",duration)
		fmt.Println("CandidateMiners:",Miners)
		MobileMiners,_:=fbc.RewardMobileMiners(prevnum,curnum)
		fmt.Println("----------------------待奖励移动矿工账号-------------------------：",curnum)
		fmt.Println("MobileMiners:",MobileMiners)
		time.Sleep(15*time.Second)
		prevnum=curnum
		i++  
	}
}
