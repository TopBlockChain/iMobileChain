package fbclient
import (
	"fmt"
	"strings"
	"math/big"
	
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/blockchain/imobilechain/params"
	"github.com/blockchain/imobilechain/crypto/sha3"
)
//定义fbclient结构
type Fbclient struct {
	//初始化参数
	channelID  string
	orgName  string
	orgAdmin  string
	ordererOrgName string 
	ccID    string
	configPath string
	channelClient *channel.Client
	ledgerClient *ledger.Client
}
//初始化fbclient,进行各项设置
func New(ChannelID string,OrgName string,OrgAdmin string,OrderOrgName string,CCID string,ConfigPath string,sdkOpts ...fabsdk.Option) *Fbclient {
	configOpt:=config.FromFile(ConfigPath)
	sdk, err := fabsdk.New(configOpt,sdkOpts...)
   if err != nil {
				
		fmt.Println("SDK创建失败", err)
	}
	clientChannelContext := sdk.ChannelContext(ChannelID, fabsdk.WithUser(OrgName), fabsdk.WithOrg(OrgName))
	client, err := channel.New(clientChannelContext)
	ledclient, err1 := ledger.New(clientChannelContext)
	if err != nil {
		fmt.Println("Fbclient创建失败", err)
	}
	if err1 != nil {
		fmt.Println("Ledgerclient创建失败", err1)
	}
	return &Fbclient{
		channelID: ChannelID,
		orgName: OrgName,
		orgAdmin: OrgAdmin,
		ordererOrgName: OrderOrgName, 
		ccID: CCID,
		configPath: ConfigPath,
		channelClient: client, 
		ledgerClient: ledclient, 
	}
}
//获取当前最新移动挖矿时间
func (Fbc *Fbclient)GetLatestMobileMiningEventTime() (resp string, err error) {
	response, err := Fbc.channelClient.Query(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "getLatestMobileMiningEventTime"},
	channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
	     fmt.Println("查询最新移动矿工失败！错误码：",err)
	}
	CurrentTime:=string(response.Payload)
	return CurrentTime,err
}
//查询获取候选矿工（矿池）账户地址
func (Fbc *Fbclient)GetPoolWinners(getPoolWinnersArgs [][]byte) (resp string, err error) {
	response, err :=Fbc.channelClient.Query(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "getPoolWinners", Args: getPoolWinnersArgs},
	channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Println("查询候选矿工失败！错误码：",err)
	}
	PoolQueryRep:=string(response.Payload)
	return PoolQueryRep,err
}
//查询获取待奖励移动矿工账户
func (Fbc *Fbclient)GetMobileMinners(getPoolWinnersArgs [][]byte) (resp string, err error) {
	response, err :=Fbc.channelClient.Query(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "getMobileMiningEvents", Args: getPoolWinnersArgs},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Println("查询候待奖励移动矿工失败！错误码：",err)
	}
	MobileMinerQueryRep:=string(response.Payload)
   	return MobileMinerQueryRep,err
}
//移动矿工挖矿invoke方法
func  (Fbc *Fbclient)AddMobileMiningEvent(addMobileMiningArgs [][]byte) error {
	_, err :=Fbc.channelClient.Execute(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "addMobileMiningEvent", Args: addMobileMiningArgs},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return err
	}
	return nil
}
//已并入主链的区块注册invoke方法
func (Fbc *Fbclient)RegisterBlockSeal(registerBlockSealArgs [][]byte) error{
	_, err := Fbc.channelClient.Execute(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "registerBlockSeal", Args: registerBlockSealArgs},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return err
	}
	return nil
}
//将查询所获得的字符串结果生成账户字符串数组
func (Fbc *Fbclient)Str2Adrr(value string)[]string {
	var miners []string
	value1:=strings.Trim(value,"[")
	value2:=strings.Trim(value1,"]")
	value3:=strings.Replace(value2,`"`,"",-1)
	miners=strings.Split(value3,",")
	return miners
} 
//Waiting time query 根据矿工在当前侯选矿工列表中的位置，计算其标准等待时间。
func (Fbc *Fbclient)WaitTime(Miners []string, Coinbase string) (Duration int64) {
	//var Duration = int64(0)
	//根据Coinbase在侯选矿工中的位置决定其等待时间
   for i:=0;i<len(Miners);i++{
	   if Miners[i]==Coinbase{
		   Duration = int64(params.BlockProcessingTime+ params.AveragyWattingTime*i)
		   return Duration  
	   }
   }
   Duration = int64(params.BlockProcessingTime + params.AveragyWattingTime*4)
   return Duration
}
func (Fbc *Fbclient)CandidateMiners(PrevNumber uint64,CurNumber uint64,Seed uint64) (Candidates []string, diff int64, err error) { 
	var Miners []string
	var candidates []string
	for k:=CurNumber;k>PrevNumber;k--{
	    block, err :=Fbc.ledgerClient.QueryBlock(k)
	    if err != nil {
		   return nil, int64(0),err
	    }
	    for _,value:=range block.Data.Data{
		//fmt.Println("key",key,"value:",string(value))
		   AddStr:=string(value)
		   if !strings.Contains(AddStr,"addMobileMiningEvent"){
			  continue
		   }
		   AddIndex:=strings.Index(AddStr,"addMobileMiningEvent")
		//MinersIndex:=strings.Index(AddStr,"0xac9d739C4d83E3501d824c4E308E7812ABA8306d")
		 //  fmt.Println("key:",key,"AddIndex",AddIndex,"addEvent:",AddStr[(AddIndex+66):(AddIndex+108)])
		   Miners=append(Miners,AddStr[(AddIndex+66):(AddIndex+108)])
		}
	}
	//fmt.Println("miners",len(Miners),Miners)
    for j:=uint64(0);j<3&&j<uint64(len(Miners));j++{
		candidates=append(candidates,Miners[HashNumber(int64(len(Miners)),Seed+j)])
	}
 	return candidates,int64(len(Miners)),nil
}
func (Fbc *Fbclient)RewardMobileMiners(PrevNumber uint64,CurNumber uint64) (MobMiners[]string,err error) { 
	for k:=CurNumber;k>PrevNumber;k--{
	    block, err :=Fbc.ledgerClient.QueryBlock(k)
	    if err != nil {
		   return nil, err
	    }
	   for _,value:=range block.Data.Data{
		//fmt.Println("key",key,"value:",string(value))
		   AddStr:=string(value)
		   if !strings.Contains(AddStr,"addMobileMiningEvent"){
			  continue
		   }
		   AddIndex:=strings.Index(AddStr,"addMobileMiningEvent")
		   //MinersIndex:=strings.Index(AddStr,"0xac9d739C4d83E3501d824c4E308E7812ABA8306d")
		   //fmt.Println("key:",key,"AddIndex",AddIndex,"addEvent:",AddStr[(AddIndex+22):(AddIndex+64)])
		   MobMiners=append(MobMiners,AddStr[(AddIndex+22):(AddIndex+64)])
		}
	}	
    return MobMiners,nil
}


func (Fbc *Fbclient)CurfbNumber() (CurNum uint64,err error) { 
	// Test Query Block by Number
	ledgerInfo, err := Fbc.ledgerClient.QueryInfo()
	if err != nil {
		fmt.Println("QuerInfo return error", err)
		return uint64(0),err
	}
	// fmt.Println("lederInfo",ledgerInfo.BCI.Height)
	// fmt.Println("lederInfo",ledgerInfo.BCI.String())
	// fmt.Println("lederInfo",ledgerInfo.BCI.CurrentBlockHash)
	
	return ledgerInfo.BCI.Height-1,nil
}
func HashNumber(Number int64, seed uint64,) uint64{
	s:=string(seed)   //以当前区块作为哈希计算的内容
 	h :=  sha3.New256()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	mp:=new(big.Int).Div(new(big.Int).SetBytes(bs),new(big.Int).SetInt64(1e18))
	mp=mp.Div(mp,big.NewInt(1e18))
	mp=mp.Div(mp,big.NewInt(1e18))
	mp=mp.Div(mp,big.NewInt(1e14))
	mp=mp.Div(mp,big.NewInt(1.2e1))
	floatNumber:=float64(mp.Int64())/float64(100000000)
	randNum:=uint64(float64(Number)*floatNumber)
	return randNum
}
