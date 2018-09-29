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
	"fmt"
	"time"
	"strconv"
	
	"github.com/blockchain/imobilechain/fbclient"
	"github.com/blockchain/imobilechain/params"
	"github.com/blockchain/imobilechain/common"

)
const (
	
		channelID      = "mychannel" //注意yaml配置文件也需要设置成这样子
		orgName        = "Org1MSP"   //注意yaml配置文件也需要设置成这样子
		orgAdmin       = "SampleOrg"
		ordererOrgName = "SampleOrg"
		ccID           = "mycc"
		path ="imc.yaml"
	
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
    //fbclient:=fbclient.New(params.ChannelID,params.OrgName,params.OrgAdmin,params.OrdererOrgName,params.CcID,params.Path),

	Fbc=fbclient.New(channelID,orgName,orgAdmin,ordererOrgName,ccID,path)
	go clientMine(Fbc,1,params.DefaulMinerAddr1.String(),params.DefaulMinerAddr2.String()) //模拟终端一
	// go clientMine(Fbc,2,params.DefaulMinerAddr2.String(),params.DefaulMinerAddr3.String()) //模拟终端二
	// go clientMine(Fbc,3,params.DefaulMinerAddr3.String(),params.DefaulMinerAddr1.String()) //模拟终端三
	// go clientMine(Fbc,4,params.DefaulMinerAddr1.String(),params.DefaulMinerAddr1.String()) //模拟终端一
	// go clientMine(Fbc,5,params.DefaulMinerAddr2.String(),params.DefaulMinerAddr2.String()) //模拟终端二
	// go clientMine(Fbc,6,params.DefaulMinerAddr3.String(),params.DefaulMinerAddr3.String()) //模拟终端三
	// go clientMine(Fbc,7,params.DefaulMinerAddr1.String(),params.DefaulMinerAddr3.String()) //模拟终端一
	// go clientMine(Fbc,8,params.DefaulMinerAddr2.String(),params.DefaulMinerAddr1.String()) //模拟终端二
	// go clientMine(Fbc,9,params.DefaulMinerAddr3.String(),params.DefaulMinerAddr2.String()) //模拟终端三

    TestMine(Fbc)
	//CurentMineWiner(Fbc)  //区块生成模拟

}
func clientMine(fbc *fbclient.Fbclient,num int,user string,miner string){
	var i int
	i=1
    var addEventArgs = [][]byte{[]byte(user), []byte(miner)}
	for {
		err:=fbc.AddMobileMiningEvent(addEventArgs)
		if err==nil{
			fmt.Println("--------------------移动矿工挖矿信息---------------------：")
			fmt.Println("模拟终端：",num,"  挖矿次数",i,"  移动矿工",user,"矿池：",miner,"状态: 成功")	
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
	curnum,_:=fbc.CurfbNumber()
	fmt.Println("当前区块号",curnum)
	prevnum,_:=fbc.CurfbNumber()
	var i=uint64(0)
	for {
		fmt.Println("------------------------资产链区块号",i,"-------------------------：")
		curnum,_:=fbc.CurfbNumber()
		Miners,_,_:=fbc.CandidateMiners(prevnum,curnum,i)
		fmt.Println("------------------------候选矿池账号-------------------------：")
		fmt.Println("CandidateMiners:",Miners)
		MobileMiners,_:=fbc.RewardMobileMiners(prevnum,curnum)
		fmt.Println("----------------------待奖励移动矿工账号-------------------------：")
		fmt.Println("MobileMiners:",MobileMiners)
		time.Sleep(15*time.Second)
		prevnum=curnum
		i++  
	}
}




func CurentMineWiner(fbc *fbclient.Fbclient){
	PrevTime,_:=fbc.GetLatestMobileMiningEventTime()
	CurrentTime,_:=fbc.GetLatestMobileMiningEventTime()
	for i:=0;i<1000;i++ {
		var getPoolWinnersArgs = [][]byte{[]byte(PrevTime), []byte(CurrentTime), []byte(strconv.Itoa(i)), []byte("3")}
		poolvalue,_:=fbc.GetPoolWinners(getPoolWinnersArgs)
		value1:=fbc.Str2Adrr(poolvalue)
		var getMobileMiningArgs = [][]byte{[]byte(PrevTime), []byte(CurrentTime),[]byte(strconv.Itoa(i))}
		mobilevalue,_:=fbc.GetMobileMinners(getMobileMiningArgs)
		value2:=fbc.Str2Adrr(mobilevalue)
		fmt.Println("------------------------区块号-------------------------：")
		fmt.Println("当前区块号：",i)
		fmt.Println("------------------------区块时间------------------------：")
        fmt.Println("前一区块时间：",PrevTime,"当前区块时间",CurrentTime)
		fmt.Println("----------------------候选矿工列表-----------------------：")
		for i,_:=range value1 {
			waittime:=fbc.WaitTime(value1,value1[i])
			fmt.Println("账户：",common.HexToAddress(value1[i]).String(),"等待时间：",waittime)
		}
		fmt.Println("--------------------奖励移动矿工列表---------------------：")
		//fmt.Println(value2)
		for i,_:=range value2 {
			fmt.Println(value2[i])
		}
		time.Sleep(15*time.Second)
		var registerBlockSealArgs=[][]byte{[]byte(PrevTime),[]byte(CurrentTime),[]byte(strconv.Itoa(i)),[]byte(poolvalue),[]byte(mobilevalue)}
		err:=fbc.RegisterBlockSeal(registerBlockSealArgs)
		if err==nil{
		    fmt.Println("--------------------区块注册信息---------------------：")
		    fmt.Println("区块号：",i,"  区块时间",CurrentTime," 候选矿工 ",value1," 移动矿工：",value2,"状态: 成功")	
		}
		PrevTime=CurrentTime
		CurrentTime,_=fbc.GetLatestMobileMiningEventTime()

	}
}

