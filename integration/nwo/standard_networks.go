
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456079077543936>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

func BasicSolo() *Config {
	return &Config{
		Organizations: []*Organization{{
			Name:          "OrdererOrg",
			MSPID:         "OrdererMSP",
			Domain:        "example.com",
			EnableNodeOUs: false,
			Users:         0,
			CA:            &CA{Hostname: "ca"},
		}, {
			Name:          "Org1",
			MSPID:         "Org1MSP",
			Domain:        "org1.example.com",
			EnableNodeOUs: true,
			Users:         2,
			CA:            &CA{Hostname: "ca"},
		}, {
			Name:          "Org2",
			MSPID:         "Org2MSP",
			Domain:        "org2.example.com",
			EnableNodeOUs: true,
			Users:         2,
			CA:            &CA{Hostname: "ca"},
		}},
		Consortiums: []*Consortium{{
			Name: "SampleConsortium",
			Organizations: []string{
				"Org1",
				"Org2",
			},
		}},
		Consensus: &Consensus{
			Type: "solo",
		},
		SystemChannel: &SystemChannel{
			Name:    "systemchannel",
			Profile: "TwoOrgsOrdererGenesis",
		},
		Orderers: []*Orderer{
			{Name: "orderer", Organization: "OrdererOrg"},
		},
		Channels: []*Channel{
			{Name: "testchannel", Profile: "TwoOrgsChannel"},
		},
		Peers: []*Peer{{
			Name:         "peer0",
			Organization: "Org1",
			Channels: []*PeerChannel{
				{Name: "testchannel", Anchor: true},
			},
		}, {
			Name:         "peer1",
			Organization: "Org1",
			Channels: []*PeerChannel{
				{Name: "testchannel", Anchor: false},
			},
		}, {
			Name:         "peer0",
			Organization: "Org2",
			Channels: []*PeerChannel{
				{Name: "testchannel", Anchor: true},
			},
		}, {
			Name:         "peer1",
			Organization: "Org2",
			Channels: []*PeerChannel{
				{Name: "testchannel", Anchor: false},
			},
		}},
		Profiles: []*Profile{{
			Name:     "TwoOrgsOrdererGenesis",
			Orderers: []string{"orderer"},
		}, {
			Name:          "TwoOrgsChannel",
			Consortium:    "SampleConsortium",
			Organizations: []string{"Org1", "Org2"},
		}},
	}
}

func BasicKafka() *Config {
	config := BasicSolo()
	config.Consensus.Type = "kafka"
	config.Consensus.ZooKeepers = 1
	config.Consensus.Brokers = 1
	return config
}

func BasicEtcdRaft() *Config {
	config := BasicSolo()
	config.Consensus.Type = "etcdraft"
	config.Profiles = []*Profile{{
		Name:     "SampleDevModeEtcdRaft",
		Orderers: []string{"orderer"},
	}, {
		Name:          "TwoOrgsChannel",
		Consortium:    "SampleConsortium",
		Organizations: []string{"Org1", "Org2"},
	}}
	config.SystemChannel.Profile = "SampleDevModeEtcdRaft"
	return config
}

func MultiChannelEtcdRaft() *Config {
	config := BasicSolo()
	config.Consensus.Type = "etcdraft"
	config.Profiles = []*Profile{{
		Name:     "SampleDevModeEtcdRaft",
		Orderers: []string{"orderer"},
	}, {
		Name:          "TwoOrgsChannel",
		Consortium:    "SampleConsortium",
		Organizations: []string{"Org1", "Org2"},
	}}
	config.SystemChannel.Profile = "SampleDevModeEtcdRaft"
	config.Channels = []*Channel{
		{Name: "testchannel1", Profile: "TwoOrgsChannel"},
		{Name: "testchannel2", Profile: "TwoOrgsChannel"}}

	for _, peer := range config.Peers {
		peer.Channels = []*PeerChannel{
			{Name: "testchannel1", Anchor: true},
			{Name: "testchannel2", Anchor: true},
		}
	}

	return config
}

func MultiNodeEtcdRaft() *Config {
	config := BasicEtcdRaft()
	config.Orderers = []*Orderer{
		{Name: "orderer1", Organization: "OrdererOrg"},
		{Name: "orderer2", Organization: "OrdererOrg"},
		{Name: "orderer3", Organization: "OrdererOrg"},
	}
	config.Profiles = []*Profile{{
		Name:     "SampleDevModeEtcdRaft",
		Orderers: []string{"orderer1", "orderer2", "orderer3"},
	}, {
		Name:          "TwoOrgsChannel",
		Consortium:    "SampleConsortium",
		Organizations: []string{"Org1", "Org2"},
	}}
	return config
}

