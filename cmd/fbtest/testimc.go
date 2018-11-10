package main

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/mocks"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	channelID      = "mychannel" //注意yaml配置文件也需要设置成这样子
	orgName        = "Org1MSP"   //注意yaml配置文件也需要设置成这样子
	orgAdmin       = "SampleOrg"
	ordererOrgName = "SampleOrg"
	ccID           = "mycc"
)

/*
linux下注意把配置文件里面路径的"\"换成"/""
注意证书文件的命名格式： [USERNAME]@.-cert.pem
同时，[USERNAME]与[USERNAME]@.-cert.pem会生成一个对应命名的keystore
EntityMatchers
channelID：配置文件里面的channels配置命名也需要跟channelID保持一致,peers也是一样
*/
// ExampleCC query and transaction arguments
var queryArgs = [][]byte{[]byte("query"), []byte("b")}
var txArgs = [][]byte{[]byte("move"), []byte("a"), []byte("0x0001"), []byte("0xc000")}
var addMobileMiningEventArgs = [][]byte{[]byte("0x5731e5b5a1202e47cddc24eec9d4c08fc8816747"), []byte("0xeD867421dabc9dC2785E54411497ae2327f28dfe")}
var getPoolWinnersArgs = [][]byte{[]byte("0"), []byte("2536310975299272624"), []byte("0")} //, []byte("25")}

func mockChannelProvider(channelID string) context.ChannelProvider {

	channelProvider := func() (context.Channel, error) {
		return mocks.NewMockChannel(channelID)
	}

	return channelProvider
}

/*
peer => -o 13.230.40.147:7050
channel ID string => -C myc
chaincode名 => -n mycc
JSON字串的链代码构造函数消息（默认”{}”） => -c '{"Args":["addMobileMiningEvent", "0x0001", "0xa000"]}'

Organization  MSP
./hyperleger/fabric/msp
*/
func setupAndRun(configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option) {
	//Init the sdk config
	sdk, err := fabsdk.New(configOpt, sdkOpts...)
	if err != nil {
		fmt.Println("Failed to create new SDK: %s", err)
		return
	}
	defer sdk.Close()
	configBackend, err := sdk.Config()
	// ************ setup complete ************** //
	//fmt.Println(configBackend.Lookup("client.logging.level"))
	channelid, _ := configBackend.Lookup("client.channelID")
	fmt.Println(channelid.(string))
	chainCodeID, _ := configBackend.Lookup("client.chainCodeID")
	fmt.Println(chainCodeID.(string))
	orgname, _ := configBackend.Lookup("client.organization")
	fmt.Println(orgname.(string))
	username, _ := configBackend.Lookup("client.userName")
	fmt.Println(username.(string))
	currpeers, _ := configBackend.Lookup("client.currpeers")
	str_currpeers := currpeers.(string)
	fmt.Println(str_currpeers)
	peers := strings.Split(str_currpeers, ",")
	fmt.Printf("peers: %v \n", peers)
	// fmt.Println(strings.Index(str_currpeers, "peer0.org1.staging.imobilechain.org:7051"))
	// fmt.Println(strings.Index(str_currpeers, "peer0.org2.staging.imobilechain.org:7051"))
	// fmt.Println(strings.Index(str_currpeers, "peer0.org3.staging.imobilechain.org:7051"))
	// fmt.Println(len(str_currpeers))

	ctx, err := sdk.Context()()
	if err != nil {
		// return errors.WithMessage(err, "Error creating anonymous provider")
		fmt.Printf("Error creating anonymous provider: %v", err)
		return
	}

	// prepare channel client context using client context
	clientChannelContext := sdk.ChannelContext(channelid.(string), fabsdk.WithUser(username.(string)), fabsdk.WithOrg(orgname.(string)))
	//clientChannelContext := mockChannelProvider(channelID)
	// fmt.Println("clientChannelContext=>", clientChannelContext)

	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)

	if err != nil {
		fmt.Println("Failed to create new channel client: %s", err)
		return
	}
	var endpointConfig = ctx.EndpointConfig()
	var networkConfig = endpointConfig.NetworkConfig()
	var allPeers []fab.Peer
	var selectPeer fab.Peer
	var selectPeers []fab.Peer
	allPeersByOrg := make(map[string][]fab.Peer)
	for orgID := range networkConfig.Organizations {
		fmt.Printf("orgID: %s \n", orgID)
		// if orgID != strings.ToLower(orgname.(string)) {
		// 	continue
		// }
		peersConfig, ok := endpointConfig.PeersConfig(orgID)
		if !ok {
			fmt.Printf("failed to get peer configs for org: %v", orgID)
			return
			// return errors.Errorf("failed to get peer configs for org [%s]", orgID)
		}

		// cliconfig.Config().Logger().Debugf("Peers for org [%s]: %v\n", orgID, peersConfig)

		var peers []fab.Peer
		for _, p := range peersConfig {
			endorser, err := ctx.InfraProvider().CreatePeerFromConfig(&fab.NetworkPeer{PeerConfig: p})
			if err != nil {
				fmt.Printf("failed to create peer from config: %v", err)
				return
				// return errors.Wrapf(err, "failed to create peer from config")
			}
			selectPeer = endorser
			peers = append(peers, endorser)
			url := endorser.URL()
			fmt.Println(url)
			index := strings.Index(str_currpeers, url)
			if index >= 0 {
				urllen := len(url)
				selectPeers = append(selectPeers, endorser)
				fmt.Println(index + urllen)
				fmt.Println(len(str_currpeers))
				str_currpeers = strings.Replace(str_currpeers, url, "", -1) //[index+urllen : len(str_currpeers)-1]
				fmt.Println(str_currpeers)
			}
			// action.orgIDByPeer[endorser.URL()] = orgID
			// if selectPeer.URL() == "peer1.org1.example.com:7051" {
			// 	break
			// }
			// break
		}
		allPeersByOrg[orgID] = peers
		allPeers = append(allPeers, peers...)
		// break
		// if selectPeer.URL() == "peer1.org1.example.com:7051" {
		// 	break
		// }
	}
	fmt.Printf("allPeers: %v \n", allPeers)
	fmt.Printf("allPeersByOrg[orgID]: %v \n", allPeersByOrg[strings.ToLower(orgname.(string))])
	fmt.Printf("selectPeer: %v \n", selectPeer)
	fmt.Printf("selectPeers: %v \n", selectPeers)
	// value, err := queryCC(client)

	value, err := getPoolWinners(client, selectPeers)
	if err != nil {
		return
	}
	fmt.Println("value is ", string(value))

	// eventID := "test([a-zA-Z]+)"

	// Register chaincode event (pass in channel which receives event details when the event is complete)
	// reg, notifier, err := client.RegisterChaincodeEvent(ccID, eventID)
	// if err != nil {
	// 	fmt.Println("Failed to register cc event: %s", err)
	// 	return
	// }
	// defer client.UnregisterChaincodeEvent(reg)

	// Move funds
	//executeCC(client)

	addMobileMiningEvent(client, selectPeers)

	// select {
	// case ccEvent := <-notifier:
	// 	fmt.Println("Received CC event: %#v\n", ccEvent)
	// case <-time.After(time.Second * 20):
	// 	fmt.Println("Did NOT receive CC event for eventId(%s)\n", eventID)
	// }
	ledgerClient, err := ledger.New(clientChannelContext)
	if err != nil {
		fmt.Println("Failed to create new ledger client: %s", err)
	}
	//ledgerInfoBefore :=
	getBlockchainInfo(ledgerClient, selectPeers)
}

func getBlockchainInfo(ledgerClient *ledger.Client, peer []fab.Peer) *fab.BlockchainInfoResponse {
	channelCfg, err := ledgerClient.QueryConfig(ledger.WithTargets(peer[0]), ledger.WithMinTargets(1))
	if err != nil {
		fmt.Printf("QueryConfig return error: %s", err)
	}
	if len(channelCfg.Orderers()) == 0 {
		fmt.Printf("Failed to retrieve channel orderers")
	}
	expectedOrderer := "orderer0.orderer.t1.imobilechain.org"
	if !strings.Contains(channelCfg.Orderers()[0], expectedOrderer) {
		fmt.Printf("Expecting %s, got %s", expectedOrderer, channelCfg.Orderers()[0])
	}
	ledgerInfoBefore, err := ledgerClient.QueryInfo(ledger.WithTargets(peer[0]), ledger.WithMinTargets(1), ledger.WithMaxTargets(3))
	if err != nil {
		fmt.Printf("QueryInfo return error: %s", err)
	}
	// Test Query Block by Hash - retrieve current block by hash
	block, err := ledgerClient.QueryBlockByHash(ledgerInfoBefore.BCI.CurrentBlockHash, ledger.WithTargets(peer[0].(fab.Peer)), ledger.WithMinTargets(1))
	if err != nil {
		fmt.Printf("QueryBlockByHash return error: %s", err)
	}
	if block == nil {
		fmt.Printf("Block info not available")
	}
	fmt.Printf("%v", ledgerInfoBefore)
	return ledgerInfoBefore
}

func addMobileMiningEvent(client *channel.Client, peer []fab.Peer) {
	var opts []channel.RequestOption
	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
	var peers []fab.Peer
	peers = append(peers, peer...)
	opts = append(opts, channel.WithTargets(peers...))
	Response, err := client.Execute(channel.Request{ChaincodeID: ccID, Fcn: "addMobileMiningEventToArray", Args: addMobileMiningEventArgs},
		opts...) //channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Println("addMobileMiningEventToArray Failed ", err)
	}
	fmt.Println("Response is ", Response)
}
func getPoolWinners(client *channel.Client, peer []fab.Peer) ([]byte, error) {
	var opts []channel.RequestOption
	opts = append(opts, channel.WithRetry(retry.DefaultChannelOpts))
	var peers []fab.Peer
	peers = append(peers, peer...)
	opts = append(opts, channel.WithTargets(peers...))
	response, err := client.Query(channel.Request{ChaincodeID: ccID, Fcn: "getMobileMiningEvents", Args: getPoolWinnersArgs},
		opts...) //channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Println("getMobileMiningEvents Failed ", err)
	}
	fmt.Println(response)

	return response.Payload, err
}

func executeCC(client *channel.Client) {
	_, err := client.Execute(channel.Request{ChaincodeID: ccID, Fcn: "invoke", Args: txArgs},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Println("Failed to move funds: %s", err)
	}
}

func queryCC(client *channel.Client) ([]byte, error) {
	response, err := client.Query(channel.Request{ChaincodeID: ccID, Fcn: "invoke", Args: queryArgs},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Println("Failed to query funds: %s", err)
	}
	fmt.Println(response)

	return response.Payload, err
}

func main() {
	//configPath := ".\\hyperledger\\fabric\\coresdk.yaml"
	// configPath := "imc.yaml"
	configPath := "fabric-sdk-config.1.1.3.yaml"
	//End to End testing
	setupAndRun(config.FromFile(configPath))
}
