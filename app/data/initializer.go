package data

import (
	channelinterface "clientapp/channel_interface"
	"clientapp/models"
)

func GetInitialUsers() map[string]*models.UserInfo {

	return map[string]*models.UserInfo{
		"jj1": {UserID: "jj1", Organization: "org1", Role: models.ADMIN, ChannelInterfaces: make(map[string]*channelinterface.ChannelInterace, 0)},
		"ik1": {UserID: "ik1", Organization: "org2", Role: models.USER, ChannelInterfaces: make(map[string]*channelinterface.ChannelInterace, 0)},
		"sg1": {UserID: "sg1", Organization: "org3", Role: models.USER, ChannelInterfaces: make(map[string]*channelinterface.ChannelInterace, 0)},
	}

}

func GetInitialChainCode() map[string]string {
	return map[string]string{
		"tradechannel1": "traderchaincode1",
		"tradechannel2": "traderchaincode2",
	}
}
