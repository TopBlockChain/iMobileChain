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
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"

)
// const (
// 	channelID      = "mychannel" //注意yaml配置文件也需要设置成这样子
// 	orgName        = "Org1MSP"   //注意yaml配置文件也需要设置成这样子
// 	orgAdmin       = "SampleOrg"
// 	ordererOrgName = "SampleOrg"
// 	ccID           = "mycc"
// )

//定义fbclient结构
type Fbclient struct {
	//初始化参数
	ccID string
	configPath string
	channelClient *channel.Client
	LedgerClient *ledger.Client
	SelectPeers []fab.Peer
	AllPeers  []fab.Peer
	CurPeers []fab.Peer
}
//初始化fbclient,进行各项设置
func New(ConfigPath string,sdkOpts ...fabsdk.Option) *Fbclient {
	configOpt:=config.FromFile(ConfigPath)
	sdk, err := fabsdk.New(configOpt,sdkOpts...)
    if err != nil {
		fmt.Println("SDK创建失败", err)
	}
	//defer sdk.Close()
	configBackend, err := sdk.Config()    //获得配置表数据
	channelid, _ := configBackend.Lookup("client.channelID")   //查询配置表channelID
	chainCodeID, _ := configBackend.Lookup("client.chainCodeID")  //查询配置表chanelCodeID
	orgname, _ := configBackend.Lookup("client.organization")    //查询配置表orgname
	username, _ := configBackend.Lookup("client.userName")   //查询配置表username
	currpeers, _ := configBackend.Lookup("client.currpeers")  //查询配置表当前节点
	str_currpeers := currpeers.(string)   //当前节点的字符串格式
    //peers := strings.Split(str_currpeers, ",")   //将当前节点转换为数组类型
    // fmt.Println("orgname",orgname)
	//prepare channel client context using client context
	//准备终端创建环境所需的 channelid、username、orgname
	clientChannelContext := sdk.ChannelContext(channelid.(string), fabsdk.WithUser(username.(string)), fabsdk.WithOrg(orgname.(string)))
	//fmt.Println("prent client",channelid.(string), username.(string), orgname.(string))
	//分别创建通道及账本客户端
	client, err := channel.New(clientChannelContext)
	ledclient, err1 := ledger.New(clientChannelContext)
	if err != nil {
		fmt.Println("Fbclient创建失败", err)
	}
	if err1 != nil {
		fmt.Println("Ledgerclient创建失败", err1)
	}

	ctx, err := sdk.Context()()     //获得环境数据
	if err != nil {
		// return errors.WithMessage(err, "Error creating anonymous provider")
		fmt.Printf("Error creating anonymous provider: %v", err)
		return nil
	}

	var endpointConfig = ctx.EndpointConfig()   //获得终端配置
	var networkConfig = endpointConfig.NetworkConfig()   //获得网络配置
	var allPeers []fab.Peer   //定义所有节点数组
	var presentPeers []fab.Peer //定义当前组织的节点数组
	var selectPeers []fab.Peer  //定义签名节点数组
	allPeersByOrg := make(map[string][]fab.Peer)  //按组织定义节点数组
	for orgID := range networkConfig.Organizations {
		//获得当前组织终端的节点配置
		peersConfig, ok := endpointConfig.PeersConfig(orgID)
		if !ok {
			fmt.Printf("failed to get peer configs for org: %v", orgID)
			return nil
		}
        //对当前组织的所有终端进行遍历和处理
		var peers []fab.Peer
		for _, p := range peersConfig {
			//获得签名节点
			endorser, err := ctx.InfraProvider().CreatePeerFromConfig(&fab.NetworkPeer{PeerConfig: p})
			if err != nil {
				fmt.Printf("failed to create peer from config: %v", err)
				return nil
			}
		//	selectPeer = endorser
		    //将签名节点加入数组
			peers = append(peers, endorser)
			url := endorser.URL()
			
			//查询签名节点是否在当前节点数组中
			index := strings.Index(str_currpeers, url)
			if index >= 0 {
			//	urllen := len(url)
			    //将签名节点加入
				selectPeers = append(selectPeers, endorser)
			    //从当前节点表中把签名节点清除
				str_currpeers = strings.Replace(str_currpeers, url, "", -1) //[index+urllen : len(str_currpeers)-1]
			}
		}
		//把当前组织节点放入该数组的节点数组中
		allPeersByOrg[orgID] = peers
		//fmt.Println("orgid",orgID,"peers",peers,"orgname",strings.ToLower(orgname.(string)))
		//将当前组织节点数组加入所有节点数组中
		allPeers = append(allPeers, peers...)
		if strings.ToLower(orgID)==strings.ToLower(orgname.(string)){
			presentPeers=peers
		}
	}
	return &Fbclient{
		ccID: chainCodeID.(string),
		configPath: ConfigPath,
		channelClient: client, 
		LedgerClient: ledclient, 
		SelectPeers: selectPeers,
		AllPeers: allPeers,
		CurPeers: presentPeers,
	}
}
// //获取当前最新移动挖矿时间
// func (Fbc *Fbclient)GetLatestMobileMiningEventTime() (resp string, err error) {
// 	var opts []channel.RequestOption
// 	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
// 	var peers []fab.Peer
// 	peers = append(peers, Fbc.selectPeers...)
// 	opts = append(opts, channel.WithTargets(peers...))
// 	response, err := Fbc.channelClient.Query(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "getLatestMobileMiningEventTime"},
// 	opts...)
// 	if err != nil {
// 	     fmt.Println("查询最新移动矿工失败！错误码：",err)
// 	}
// 	CurrentTime:=string(response.Payload)
// 	return CurrentTime,err
// }
// //查询获取候选矿工（矿池）账户地址
// func (Fbc *Fbclient)GetPoolWinners(getPoolWinnersArgs [][]byte) (resp string, err error) {
// 	var opts []channel.RequestOption
// 	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
// 	var peers []fab.Peer
// 	peers = append(peers, Fbc.selectPeers...)
// 	opts = append(opts, channel.WithTargets(peers...))

// 	response, err :=Fbc.channelClient.Query(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "getPoolWinners", Args: getPoolWinnersArgs},
// 	opts...)
// 	if err != nil {
// 		fmt.Println("查询候选矿工失败！错误码：",err)
// 	}
// 	PoolQueryRep:=string(response.Payload)
// 	return PoolQueryRep,err
// }
// //查询获取待奖励移动矿工账户
// func (Fbc *Fbclient)GetMobileMinners(getPoolWinnersArgs [][]byte) (resp string, err error) {
// 	var opts []channel.RequestOption
// 	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
// 	var peers []fab.Peer
// 	peers = append(peers, Fbc.selectPeers...)
// 	opts = append(opts, channel.WithTargets(peers...))

// 	response, err :=Fbc.channelClient.Query(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "getMobileMiningEvents", Args: getPoolWinnersArgs},
// 		opts...)
// 	if err != nil {
// 		fmt.Println("查询候待奖励移动矿工失败！错误码：",err)
// 	}
// 	MobileMinerQueryRep:=string(response.Payload)
//    	return MobileMinerQueryRep,err
// }
//移动矿工挖矿invoke方法addMobileMiningEvent
func  (Fbc *Fbclient)AddMobileMiningEvent(addMobileMiningArgs [][]byte) error {
	var opts []channel.RequestOption
	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
	var peers []fab.Peer
	peers = append(peers, Fbc.SelectPeers...)
	opts = append(opts, channel.WithTargets(peers...))
	//fmt.Println(Fbc.ccID,peers)
	_, err :=Fbc.channelClient.Execute(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "addMobileMiningEventToArray", Args: addMobileMiningArgs},
		opts...)
	if err != nil {
	    fmt.Println(err)
		return err
	}
	return nil
}
//已并入主链的区块注册invoke方法
// func (Fbc *Fbclient)RegisterBlockSeal(registerBlockSealArgs [][]byte) error{
// 	var opts []channel.RequestOption
// 	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
// 	var peers []fab.Peer
// 	peers = append(peers, Fbc.selectPeers...)
// 	opts = append(opts, channel.WithTargets(peers...))
// 	_, err := Fbc.channelClient.Execute(channel.Request{ChaincodeID: Fbc.ccID, Fcn: "registerBlockSeal", Args: registerBlockSealArgs},
// 		opts...)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
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
func (Fbc *Fbclient)WaitTime(Miners []string, Coinbase string,blocknumber uint64) (Duration int64) {
	//var Duration = int64(0)
	//根据Coinbase在侯选矿工中的位置决定其等待时间
   for i:=0;i<len(Miners);i++{
	   if Miners[i]==Coinbase{
		   Duration = int64(params.BlockProcessingTime+ params.AveragyWattingTime*i)
		   return Duration  
	   }
   }
   Duration = int64(params.BlockProcessingTime + params.AveragyWattingTime*3)+int64(HashNumber(int64(params.AveragyWattingTime),Coinbase+string(blocknumber)))
   return Duration
}
func (Fbc *Fbclient)CandidateMiners(PrevNumber uint64,CurNumber uint64,Seed uint64) (Candidates []string, diff int64, err error) { 
	var Miners []string
	var candidates []string
	var opts []ledger.RequestOption
	opts = append(opts, ledger.WithTargets(Fbc.AllPeers...))
    for k:=CurNumber;k>PrevNumber;k--{
	    block, err :=Fbc.LedgerClient.QueryBlock(k,opts...)
		if err != nil {
			return nil, int64(0),err
		 }
		for _,value:=range block.Data.Data{
		   AddStr:=string(value)
		   if !strings.Contains(AddStr,"addMobileMiningEventToArray"){
			  continue
		   }
		   AddIndex:=strings.Index(AddStr,"addMobileMiningEventToArray")
		   Miners=append(Miners,AddStr[(AddIndex+73):(AddIndex+115)])
		}
	}
	for j:=uint64(0);j<3&&j<uint64(len(Miners));j++{
		candidates=append(candidates,Miners[HashNumber(int64(len(Miners)),string(Seed+j))])
	}
 	return candidates,int64(len(Miners)+1),nil
}
func (Fbc *Fbclient)RewardMobileMiners(PrevNumber uint64,CurNumber uint64) (MobMiners[]string,err error) { 
	var opts []ledger.RequestOption
	opts = append(opts, ledger.WithTargets(Fbc.AllPeers...))
	for k:=CurNumber;k>PrevNumber;k--{
	    block, err :=Fbc.LedgerClient.QueryBlock(k,opts...)
	    if err != nil {
			  return nil, err
	    }
	   for _,value:=range block.Data.Data{
		   AddStr:=string(value)
		   if !strings.Contains(AddStr,"addMobileMiningEventToArray"){
			  continue
		   }
		   AddIndex:=strings.Index(AddStr,"addMobileMiningEventToArray")
		   MobMiners=append(MobMiners,AddStr[(AddIndex+29):(AddIndex+71)])
		}
	}
    return MobMiners,nil
}

func (Fbc *Fbclient)CurfbNumber() (CurNum uint64,err error) { 
	var opts []ledger.RequestOption
	opts = append(opts, ledger.WithTargets(Fbc.AllPeers...))
	ledgerInfo, err := Fbc.LedgerClient.QueryInfo(opts...)
	if err != nil {
			return uint64(0),err
     }
	//  fmt.Println("lederInfo",ledgerInfo.BCI.Height)
	//  fmt.Println("lederInfo",ledgerInfo.BCI.String())
	//  fmt.Println("lederInfo",ledgerInfo.BCI.CurrentBlockHash)
	return ledgerInfo.BCI.Height-1,nil
}
func HashNumber(Number int64, seed string) uint64{
	s:=seed   //以当前区块作为哈希计算的内容
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
