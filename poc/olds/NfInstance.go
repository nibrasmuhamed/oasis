package main

type NfInstance struct {
	NfInstanceID          string `json:"nfInstanceId"`
	NfInstanceName        string `json:"nfInstanceName"`
	NfType                string `json:"nfType"`
	NfStatus              string `json:"nfStatus"`
	CollocatedNfInstances []struct {
		NfInstanceID string `json:"nfInstanceId"`
		NfType       string `json:"nfType"`
	} `json:"collocatedNfInstances"`
	HeartBeatTimer    int      `json:"heartBeatTimer"`
	PlmnList          []string `json:"plmnList"`
	SnpnList          []string `json:"snpnList"`
	SNssais           []string `json:"sNssais"`
	PerPlmnSnssaiList []struct {
		PlmnID     string   `json:"plmnId"`
		SNssaiList []string `json:"sNssaiList"`
		Nid        string   `json:"nid"`
	} `json:"perPlmnSnssaiList"`
	NsiList          []string `json:"nsiList"`
	Fqdn             string   `json:"fqdn"`
	InterPlmnFqdn    string   `json:"interPlmnFqdn"`
	Ipv4Addresses    []string `json:"ipv4Addresses"`
	Ipv6Addresses    []string `json:"ipv6Addresses"`
	AllowedPlmns     []string `json:"allowedPlmns"`
	AllowedSnpns     []string `json:"allowedSnpns"`
	AllowedNfTypes   []string `json:"allowedNfTypes"`
	AllowedNfDomains []string `json:"allowedNfDomains"`
	AllowedNssais    []string `json:"allowedNssais"`
	Priority         int      `json:"priority"`
	Capacity         int      `json:"capacity"`
	Load             int      `json:"load"`
	LoadTimeStamp    string   `json:"loadTimeStamp"`
	Locality         string   `json:"locality"`
	ExtLocality      struct {
		AdditionalProp1 string `json:"additionalProp1"`
	} `json:"extLocality"`
	UdrInfo struct {
		GroupID    string `json:"groupId"`
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		GpsiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"gpsiRanges"`
		ExternalGroupIdentifiersRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"externalGroupIdentifiersRanges"`
		SupportedDataSets  []string `json:"supportedDataSets"`
		SharedDataIDRanges []struct {
			Pattern string `json:"pattern"`
		} `json:"sharedDataIdRanges"`
	} `json:"udrInfo"`
	UdrInfoList struct {
		AdditionalProp1 struct {
			GroupID    string `json:"groupId"`
			SupiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			GpsiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"gpsiRanges"`
			ExternalGroupIdentifiersRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"externalGroupIdentifiersRanges"`
			SupportedDataSets  []string `json:"supportedDataSets"`
			SharedDataIDRanges []struct {
				Pattern string `json:"pattern"`
			} `json:"sharedDataIdRanges"`
		} `json:"additionalProp1"`
	} `json:"udrInfoList"`
	UdmInfo struct {
		GroupID    string `json:"groupId"`
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		GpsiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"gpsiRanges"`
		ExternalGroupIdentifiersRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"externalGroupIdentifiersRanges"`
		RoutingIndicators              []string `json:"routingIndicators"`
		InternalGroupIdentifiersRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"internalGroupIdentifiersRanges"`
		SuciInfos []struct {
			RoutingInds  []string `json:"routingInds"`
			HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
		} `json:"suciInfos"`
	} `json:"udmInfo"`
	UdmInfoList struct {
		AdditionalProp1 struct {
			GroupID    string `json:"groupId"`
			SupiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			GpsiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"gpsiRanges"`
			ExternalGroupIdentifiersRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"externalGroupIdentifiersRanges"`
			RoutingIndicators              []string `json:"routingIndicators"`
			InternalGroupIdentifiersRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"internalGroupIdentifiersRanges"`
			SuciInfos []struct {
				RoutingInds  []string `json:"routingInds"`
				HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
			} `json:"suciInfos"`
		} `json:"additionalProp1"`
	} `json:"udmInfoList"`
	AusfInfo struct {
		GroupID    string `json:"groupId"`
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		RoutingIndicators []string `json:"routingIndicators"`
		SuciInfos         []struct {
			RoutingInds  []string `json:"routingInds"`
			HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
		} `json:"suciInfos"`
	} `json:"ausfInfo"`
	AusfInfoList struct {
		AdditionalProp1 struct {
			GroupID    string `json:"groupId"`
			SupiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			RoutingIndicators []string `json:"routingIndicators"`
			SuciInfos         []struct {
				RoutingInds  []string `json:"routingInds"`
				HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
			} `json:"suciInfos"`
		} `json:"additionalProp1"`
	} `json:"ausfInfoList"`
	AmfInfo struct {
		AmfSetID     string   `json:"amfSetId"`
		AmfRegionID  string   `json:"amfRegionId"`
		GuamiList    []string `json:"guamiList"`
		TaiList      []string `json:"taiList"`
		TaiRangeList []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		BackupInfoAmfFailure []string `json:"backupInfoAmfFailure"`
		BackupInfoAmfRemoval []string `json:"backupInfoAmfRemoval"`
		N2InterfaceAmfInfo   struct {
			Ipv4EndpointAddress []string `json:"ipv4EndpointAddress"`
			Ipv6EndpointAddress []string `json:"ipv6EndpointAddress"`
			AmfName             string   `json:"amfName"`
		} `json:"n2InterfaceAmfInfo"`
		AmfOnboardingCapability bool `json:"amfOnboardingCapability"`
		HighLatencyCom          bool `json:"highLatencyCom"`
	} `json:"amfInfo"`
	AmfInfoList struct {
		AdditionalProp1 struct {
			AmfSetID     string   `json:"amfSetId"`
			AmfRegionID  string   `json:"amfRegionId"`
			GuamiList    []string `json:"guamiList"`
			TaiList      []string `json:"taiList"`
			TaiRangeList []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
			BackupInfoAmfFailure []string `json:"backupInfoAmfFailure"`
			BackupInfoAmfRemoval []string `json:"backupInfoAmfRemoval"`
			N2InterfaceAmfInfo   struct {
				Ipv4EndpointAddress []string `json:"ipv4EndpointAddress"`
				Ipv6EndpointAddress []string `json:"ipv6EndpointAddress"`
				AmfName             string   `json:"amfName"`
			} `json:"n2InterfaceAmfInfo"`
			AmfOnboardingCapability bool `json:"amfOnboardingCapability"`
			HighLatencyCom          bool `json:"highLatencyCom"`
		} `json:"additionalProp1"`
	} `json:"amfInfoList"`
	SmfInfo struct {
		SNssaiSmfInfoList []struct {
			SNssai         string `json:"sNssai"`
			DnnSmfInfoList []struct {
				Dnn      string   `json:"dnn"`
				DnaiList []string `json:"dnaiList"`
			} `json:"dnnSmfInfoList"`
		} `json:"sNssaiSmfInfoList"`
		TaiList      []string `json:"taiList"`
		TaiRangeList []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		PgwFqdn           string   `json:"pgwFqdn"`
		PgwIPAddrList     []string `json:"pgwIpAddrList"`
		AccessType        []string `json:"accessType"`
		Priority          int      `json:"priority"`
		VsmfSupportInd    bool     `json:"vsmfSupportInd"`
		PgwFqdnList       []string `json:"pgwFqdnList"`
		IsmfSupportInd    bool     `json:"ismfSupportInd"`
		SmfUPRPCapability bool     `json:"smfUPRPCapability"`
	} `json:"smfInfo"`
	SmfInfoList struct {
		AdditionalProp1 struct {
			SNssaiSmfInfoList []struct {
				SNssai         string `json:"sNssai"`
				DnnSmfInfoList []struct {
					Dnn      string   `json:"dnn"`
					DnaiList []string `json:"dnaiList"`
				} `json:"dnnSmfInfoList"`
			} `json:"sNssaiSmfInfoList"`
			TaiList      []string `json:"taiList"`
			TaiRangeList []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
			PgwFqdn           string   `json:"pgwFqdn"`
			PgwIPAddrList     []string `json:"pgwIpAddrList"`
			AccessType        []string `json:"accessType"`
			Priority          int      `json:"priority"`
			VsmfSupportInd    bool     `json:"vsmfSupportInd"`
			PgwFqdnList       []string `json:"pgwFqdnList"`
			IsmfSupportInd    bool     `json:"ismfSupportInd"`
			SmfUPRPCapability bool     `json:"smfUPRPCapability"`
		} `json:"additionalProp1"`
	} `json:"smfInfoList"`
	UpfInfo struct {
		SNssaiUpfInfoList []struct {
			SNssai         string `json:"sNssai"`
			DnnUpfInfoList []struct {
				Dnn               string   `json:"dnn"`
				DnaiList          []string `json:"dnaiList"`
				PduSessionTypes   []string `json:"pduSessionTypes"`
				Ipv4AddressRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"ipv4AddressRanges"`
				Ipv6PrefixRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"ipv6PrefixRanges"`
				NatedIpv4AddressRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"natedIpv4AddressRanges"`
				NatedIpv6PrefixRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"natedIpv6PrefixRanges"`
				Ipv4IndexList      []string `json:"ipv4IndexList"`
				Ipv6IndexList      []string `json:"ipv6IndexList"`
				NetworkInstance    string   `json:"networkInstance"`
				DnaiNwInstanceList struct {
					AdditionalProp1 string `json:"additionalProp1"`
				} `json:"dnaiNwInstanceList"`
			} `json:"dnnUpfInfoList"`
			RedundantTransport bool `json:"redundantTransport"`
		} `json:"sNssaiUpfInfoList"`
		SmfServingArea       []string `json:"smfServingArea"`
		InterfaceUpfInfoList []struct {
			InterfaceType         string   `json:"interfaceType"`
			Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
			Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
			EndpointFqdn          string   `json:"endpointFqdn"`
			NetworkInstance       string   `json:"networkInstance"`
		} `json:"interfaceUpfInfoList"`
		IwkEpsInd       bool     `json:"iwkEpsInd"`
		SxaInd          bool     `json:"sxaInd"`
		PduSessionTypes []string `json:"pduSessionTypes"`
		AtsssCapability string   `json:"atsssCapability"`
		UeIPAddrInd     bool     `json:"ueIpAddrInd"`
		TaiList         []string `json:"taiList"`
		TaiRangeList    []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		WAgfInfo struct {
			Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
			Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
			EndpointFqdn          string   `json:"endpointFqdn"`
		} `json:"wAgfInfo"`
		TngfInfo struct {
			Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
			Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
			EndpointFqdn          string   `json:"endpointFqdn"`
		} `json:"tngfInfo"`
		TwifInfo struct {
			Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
			Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
			EndpointFqdn          string   `json:"endpointFqdn"`
		} `json:"twifInfo"`
		Priority              int      `json:"priority"`
		RedundantGtpu         bool     `json:"redundantGtpu"`
		Ipups                 bool     `json:"ipups"`
		DataForwarding        bool     `json:"dataForwarding"`
		SupportedPfcpFeatures string   `json:"supportedPfcpFeatures"`
		UpfEvents             []string `json:"upfEvents"`
	} `json:"upfInfo"`
	UpfInfoList struct {
		AdditionalProp1 struct {
			SNssaiUpfInfoList []struct {
				SNssai         string `json:"sNssai"`
				DnnUpfInfoList []struct {
					Dnn               string   `json:"dnn"`
					DnaiList          []string `json:"dnaiList"`
					PduSessionTypes   []string `json:"pduSessionTypes"`
					Ipv4AddressRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"ipv4AddressRanges"`
					Ipv6PrefixRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"ipv6PrefixRanges"`
					NatedIpv4AddressRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"natedIpv4AddressRanges"`
					NatedIpv6PrefixRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"natedIpv6PrefixRanges"`
					Ipv4IndexList      []string `json:"ipv4IndexList"`
					Ipv6IndexList      []string `json:"ipv6IndexList"`
					NetworkInstance    string   `json:"networkInstance"`
					DnaiNwInstanceList struct {
						AdditionalProp1 string `json:"additionalProp1"`
					} `json:"dnaiNwInstanceList"`
				} `json:"dnnUpfInfoList"`
				RedundantTransport bool `json:"redundantTransport"`
			} `json:"sNssaiUpfInfoList"`
			SmfServingArea       []string `json:"smfServingArea"`
			InterfaceUpfInfoList []struct {
				InterfaceType         string   `json:"interfaceType"`
				Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
				Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
				EndpointFqdn          string   `json:"endpointFqdn"`
				NetworkInstance       string   `json:"networkInstance"`
			} `json:"interfaceUpfInfoList"`
			IwkEpsInd       bool     `json:"iwkEpsInd"`
			SxaInd          bool     `json:"sxaInd"`
			PduSessionTypes []string `json:"pduSessionTypes"`
			AtsssCapability string   `json:"atsssCapability"`
			UeIPAddrInd     bool     `json:"ueIpAddrInd"`
			TaiList         []string `json:"taiList"`
			TaiRangeList    []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
			WAgfInfo struct {
				Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
				Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
				EndpointFqdn          string   `json:"endpointFqdn"`
			} `json:"wAgfInfo"`
			TngfInfo struct {
				Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
				Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
				EndpointFqdn          string   `json:"endpointFqdn"`
			} `json:"tngfInfo"`
			TwifInfo struct {
				Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
				Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
				EndpointFqdn          string   `json:"endpointFqdn"`
			} `json:"twifInfo"`
			Priority              int      `json:"priority"`
			RedundantGtpu         bool     `json:"redundantGtpu"`
			Ipups                 bool     `json:"ipups"`
			DataForwarding        bool     `json:"dataForwarding"`
			SupportedPfcpFeatures string   `json:"supportedPfcpFeatures"`
			UpfEvents             []string `json:"upfEvents"`
		} `json:"additionalProp1"`
	} `json:"upfInfoList"`
	PcfInfo struct {
		GroupID    string   `json:"groupId"`
		DnnList    []string `json:"dnnList"`
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		GpsiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"gpsiRanges"`
		RxDiamHost      string `json:"rxDiamHost"`
		RxDiamRealm     string `json:"rxDiamRealm"`
		V2XSupportInd   bool   `json:"v2xSupportInd"`
		ProseSupportInd bool   `json:"proseSupportInd"`
		ProseCapability struct {
			ProseDirectDiscovey      bool `json:"proseDirectDiscovey"`
			ProseDirectCommunication bool `json:"proseDirectCommunication"`
			ProseL2UetoNetworkRelay  bool `json:"proseL2UetoNetworkRelay"`
			ProseL3UetoNetworkRelay  bool `json:"proseL3UetoNetworkRelay"`
			ProseL2RemoteUe          bool `json:"proseL2RemoteUe"`
			ProseL3RemoteUe          bool `json:"proseL3RemoteUe"`
		} `json:"proseCapability"`
		V2XCapability struct {
			LteV2X bool `json:"lteV2x"`
			NrV2X  bool `json:"nrV2x"`
		} `json:"v2xCapability"`
	} `json:"pcfInfo"`
	PcfInfoList struct {
		AdditionalProp1 struct {
			GroupID    string   `json:"groupId"`
			DnnList    []string `json:"dnnList"`
			SupiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			GpsiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"gpsiRanges"`
			RxDiamHost      string `json:"rxDiamHost"`
			RxDiamRealm     string `json:"rxDiamRealm"`
			V2XSupportInd   bool   `json:"v2xSupportInd"`
			ProseSupportInd bool   `json:"proseSupportInd"`
			ProseCapability struct {
				ProseDirectDiscovey      bool `json:"proseDirectDiscovey"`
				ProseDirectCommunication bool `json:"proseDirectCommunication"`
				ProseL2UetoNetworkRelay  bool `json:"proseL2UetoNetworkRelay"`
				ProseL3UetoNetworkRelay  bool `json:"proseL3UetoNetworkRelay"`
				ProseL2RemoteUe          bool `json:"proseL2RemoteUe"`
				ProseL3RemoteUe          bool `json:"proseL3RemoteUe"`
			} `json:"proseCapability"`
			V2XCapability struct {
				LteV2X bool `json:"lteV2x"`
				NrV2X  bool `json:"nrV2x"`
			} `json:"v2xCapability"`
		} `json:"additionalProp1"`
	} `json:"pcfInfoList"`
	BsfInfo struct {
		DnnList           []string `json:"dnnList"`
		IPDomainList      []string `json:"ipDomainList"`
		Ipv4AddressRanges []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"ipv4AddressRanges"`
		Ipv6PrefixRanges []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"ipv6PrefixRanges"`
		RxDiamHost  string `json:"rxDiamHost"`
		RxDiamRealm string `json:"rxDiamRealm"`
		GroupID     string `json:"groupId"`
		SupiRanges  []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		GpsiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"gpsiRanges"`
	} `json:"bsfInfo"`
	BsfInfoList struct {
		AdditionalProp1 struct {
			DnnList           []string `json:"dnnList"`
			IPDomainList      []string `json:"ipDomainList"`
			Ipv4AddressRanges []struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"ipv4AddressRanges"`
			Ipv6PrefixRanges []struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"ipv6PrefixRanges"`
			RxDiamHost  string `json:"rxDiamHost"`
			RxDiamRealm string `json:"rxDiamRealm"`
			GroupID     string `json:"groupId"`
			SupiRanges  []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			GpsiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"gpsiRanges"`
		} `json:"additionalProp1"`
	} `json:"bsfInfoList"`
	ChfInfo struct {
		SupiRangeList []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRangeList"`
		GpsiRangeList []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"gpsiRangeList"`
		PlmnRangeList []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"plmnRangeList"`
		GroupID              string `json:"groupId"`
		PrimaryChfInstance   string `json:"primaryChfInstance"`
		SecondaryChfInstance string `json:"secondaryChfInstance"`
	} `json:"chfInfo"`
	ChfInfoList struct {
		AdditionalProp1 struct {
			SupiRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRangeList"`
			GpsiRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"gpsiRangeList"`
			PlmnRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"plmnRangeList"`
			GroupID              string `json:"groupId"`
			PrimaryChfInstance   string `json:"primaryChfInstance"`
			SecondaryChfInstance string `json:"secondaryChfInstance"`
		} `json:"additionalProp1"`
	} `json:"chfInfoList"`
	NefInfo struct {
		NefID   string `json:"nefId"`
		PfdData struct {
			AppIds []string `json:"appIds"`
			AfIds  []string `json:"afIds"`
		} `json:"pfdData"`
		AfEeData struct {
			AfEvents []string `json:"afEvents"`
			AfIds    []string `json:"afIds"`
			AppIds   []string `json:"appIds"`
		} `json:"afEeData"`
		GpsiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"gpsiRanges"`
		ExternalGroupIdentifiersRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"externalGroupIdentifiersRanges"`
		ServedFqdnList []string `json:"servedFqdnList"`
		TaiList        []string `json:"taiList"`
		TaiRangeList   []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		DnaiList          []string `json:"dnaiList"`
		UnTrustAfInfoList []struct {
			AfID           string `json:"afId"`
			SNssaiInfoList []struct {
				SNssai      string `json:"sNssai"`
				DnnInfoList []struct {
					Dnn string `json:"dnn"`
				} `json:"dnnInfoList"`
			} `json:"sNssaiInfoList"`
			MappingInd bool `json:"mappingInd"`
		} `json:"unTrustAfInfoList"`
		UasNfFunctionalityInd bool `json:"uasNfFunctionalityInd"`
	} `json:"nefInfo"`
	NrfInfo struct {
		ServedUdrInfo struct {
			AdditionalProp1 struct {
				GroupID    string `json:"groupId"`
				SupiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				GpsiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"gpsiRanges"`
				ExternalGroupIdentifiersRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"externalGroupIdentifiersRanges"`
				SupportedDataSets  []string `json:"supportedDataSets"`
				SharedDataIDRanges []struct {
					Pattern string `json:"pattern"`
				} `json:"sharedDataIdRanges"`
			} `json:"additionalProp1"`
		} `json:"servedUdrInfo"`
		ServedUdrInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					GroupID    string `json:"groupId"`
					SupiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					GpsiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"gpsiRanges"`
					ExternalGroupIdentifiersRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"externalGroupIdentifiersRanges"`
					SupportedDataSets  []string `json:"supportedDataSets"`
					SharedDataIDRanges []struct {
						Pattern string `json:"pattern"`
					} `json:"sharedDataIdRanges"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedUdrInfoList"`
		ServedUdmInfo struct {
			AdditionalProp1 struct {
				GroupID    string `json:"groupId"`
				SupiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				GpsiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"gpsiRanges"`
				ExternalGroupIdentifiersRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"externalGroupIdentifiersRanges"`
				RoutingIndicators              []string `json:"routingIndicators"`
				InternalGroupIdentifiersRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"internalGroupIdentifiersRanges"`
				SuciInfos []struct {
					RoutingInds  []string `json:"routingInds"`
					HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
				} `json:"suciInfos"`
			} `json:"additionalProp1"`
		} `json:"servedUdmInfo"`
		ServedUdmInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					GroupID    string `json:"groupId"`
					SupiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					GpsiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"gpsiRanges"`
					ExternalGroupIdentifiersRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"externalGroupIdentifiersRanges"`
					RoutingIndicators              []string `json:"routingIndicators"`
					InternalGroupIdentifiersRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"internalGroupIdentifiersRanges"`
					SuciInfos []struct {
						RoutingInds  []string `json:"routingInds"`
						HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
					} `json:"suciInfos"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedUdmInfoList"`
		ServedAusfInfo struct {
			AdditionalProp1 struct {
				GroupID    string `json:"groupId"`
				SupiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				RoutingIndicators []string `json:"routingIndicators"`
				SuciInfos         []struct {
					RoutingInds  []string `json:"routingInds"`
					HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
				} `json:"suciInfos"`
			} `json:"additionalProp1"`
		} `json:"servedAusfInfo"`
		ServedAusfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					GroupID    string `json:"groupId"`
					SupiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					RoutingIndicators []string `json:"routingIndicators"`
					SuciInfos         []struct {
						RoutingInds  []string `json:"routingInds"`
						HNwPubKeyIds []int    `json:"hNwPubKeyIds"`
					} `json:"suciInfos"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedAusfInfoList"`
		ServedAmfInfo struct {
			AdditionalProp1 struct {
				AmfSetID     string   `json:"amfSetId"`
				AmfRegionID  string   `json:"amfRegionId"`
				GuamiList    []string `json:"guamiList"`
				TaiList      []string `json:"taiList"`
				TaiRangeList []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
				BackupInfoAmfFailure []string `json:"backupInfoAmfFailure"`
				BackupInfoAmfRemoval []string `json:"backupInfoAmfRemoval"`
				N2InterfaceAmfInfo   struct {
					Ipv4EndpointAddress []string `json:"ipv4EndpointAddress"`
					Ipv6EndpointAddress []string `json:"ipv6EndpointAddress"`
					AmfName             string   `json:"amfName"`
				} `json:"n2InterfaceAmfInfo"`
				AmfOnboardingCapability bool `json:"amfOnboardingCapability"`
				HighLatencyCom          bool `json:"highLatencyCom"`
			} `json:"additionalProp1"`
		} `json:"servedAmfInfo"`
		ServedAmfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					AmfSetID     string   `json:"amfSetId"`
					AmfRegionID  string   `json:"amfRegionId"`
					GuamiList    []string `json:"guamiList"`
					TaiList      []string `json:"taiList"`
					TaiRangeList []struct {
						PlmnID       string `json:"plmnId"`
						TacRangeList []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"tacRangeList"`
						Nid string `json:"nid"`
					} `json:"taiRangeList"`
					BackupInfoAmfFailure []string `json:"backupInfoAmfFailure"`
					BackupInfoAmfRemoval []string `json:"backupInfoAmfRemoval"`
					N2InterfaceAmfInfo   struct {
						Ipv4EndpointAddress []string `json:"ipv4EndpointAddress"`
						Ipv6EndpointAddress []string `json:"ipv6EndpointAddress"`
						AmfName             string   `json:"amfName"`
					} `json:"n2InterfaceAmfInfo"`
					AmfOnboardingCapability bool `json:"amfOnboardingCapability"`
					HighLatencyCom          bool `json:"highLatencyCom"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedAmfInfoList"`
		ServedSmfInfo struct {
			AdditionalProp1 struct {
				SNssaiSmfInfoList []struct {
					SNssai         string `json:"sNssai"`
					DnnSmfInfoList []struct {
						Dnn      string   `json:"dnn"`
						DnaiList []string `json:"dnaiList"`
					} `json:"dnnSmfInfoList"`
				} `json:"sNssaiSmfInfoList"`
				TaiList      []string `json:"taiList"`
				TaiRangeList []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
				PgwFqdn           string   `json:"pgwFqdn"`
				PgwIPAddrList     []string `json:"pgwIpAddrList"`
				AccessType        []string `json:"accessType"`
				Priority          int      `json:"priority"`
				VsmfSupportInd    bool     `json:"vsmfSupportInd"`
				PgwFqdnList       []string `json:"pgwFqdnList"`
				IsmfSupportInd    bool     `json:"ismfSupportInd"`
				SmfUPRPCapability bool     `json:"smfUPRPCapability"`
			} `json:"additionalProp1"`
		} `json:"servedSmfInfo"`
		ServedSmfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SNssaiSmfInfoList []struct {
						SNssai         string `json:"sNssai"`
						DnnSmfInfoList []struct {
							Dnn      string   `json:"dnn"`
							DnaiList []string `json:"dnaiList"`
						} `json:"dnnSmfInfoList"`
					} `json:"sNssaiSmfInfoList"`
					TaiList      []string `json:"taiList"`
					TaiRangeList []struct {
						PlmnID       string `json:"plmnId"`
						TacRangeList []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"tacRangeList"`
						Nid string `json:"nid"`
					} `json:"taiRangeList"`
					PgwFqdn           string   `json:"pgwFqdn"`
					PgwIPAddrList     []string `json:"pgwIpAddrList"`
					AccessType        []string `json:"accessType"`
					Priority          int      `json:"priority"`
					VsmfSupportInd    bool     `json:"vsmfSupportInd"`
					PgwFqdnList       []string `json:"pgwFqdnList"`
					IsmfSupportInd    bool     `json:"ismfSupportInd"`
					SmfUPRPCapability bool     `json:"smfUPRPCapability"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedSmfInfoList"`
		ServedUpfInfo struct {
			AdditionalProp1 struct {
				SNssaiUpfInfoList []struct {
					SNssai         string `json:"sNssai"`
					DnnUpfInfoList []struct {
						Dnn               string   `json:"dnn"`
						DnaiList          []string `json:"dnaiList"`
						PduSessionTypes   []string `json:"pduSessionTypes"`
						Ipv4AddressRanges []struct {
							Start string `json:"start"`
							End   string `json:"end"`
						} `json:"ipv4AddressRanges"`
						Ipv6PrefixRanges []struct {
							Start string `json:"start"`
							End   string `json:"end"`
						} `json:"ipv6PrefixRanges"`
						NatedIpv4AddressRanges []struct {
							Start string `json:"start"`
							End   string `json:"end"`
						} `json:"natedIpv4AddressRanges"`
						NatedIpv6PrefixRanges []struct {
							Start string `json:"start"`
							End   string `json:"end"`
						} `json:"natedIpv6PrefixRanges"`
						Ipv4IndexList      []string `json:"ipv4IndexList"`
						Ipv6IndexList      []string `json:"ipv6IndexList"`
						NetworkInstance    string   `json:"networkInstance"`
						DnaiNwInstanceList struct {
							AdditionalProp1 string `json:"additionalProp1"`
						} `json:"dnaiNwInstanceList"`
					} `json:"dnnUpfInfoList"`
					RedundantTransport bool `json:"redundantTransport"`
				} `json:"sNssaiUpfInfoList"`
				SmfServingArea       []string `json:"smfServingArea"`
				InterfaceUpfInfoList []struct {
					InterfaceType         string   `json:"interfaceType"`
					Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
					Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
					EndpointFqdn          string   `json:"endpointFqdn"`
					NetworkInstance       string   `json:"networkInstance"`
				} `json:"interfaceUpfInfoList"`
				IwkEpsInd       bool     `json:"iwkEpsInd"`
				SxaInd          bool     `json:"sxaInd"`
				PduSessionTypes []string `json:"pduSessionTypes"`
				AtsssCapability string   `json:"atsssCapability"`
				UeIPAddrInd     bool     `json:"ueIpAddrInd"`
				TaiList         []string `json:"taiList"`
				TaiRangeList    []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
				WAgfInfo struct {
					Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
					Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
					EndpointFqdn          string   `json:"endpointFqdn"`
				} `json:"wAgfInfo"`
				TngfInfo struct {
					Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
					Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
					EndpointFqdn          string   `json:"endpointFqdn"`
				} `json:"tngfInfo"`
				TwifInfo struct {
					Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
					Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
					EndpointFqdn          string   `json:"endpointFqdn"`
				} `json:"twifInfo"`
				Priority              int      `json:"priority"`
				RedundantGtpu         bool     `json:"redundantGtpu"`
				Ipups                 bool     `json:"ipups"`
				DataForwarding        bool     `json:"dataForwarding"`
				SupportedPfcpFeatures string   `json:"supportedPfcpFeatures"`
				UpfEvents             []string `json:"upfEvents"`
			} `json:"additionalProp1"`
		} `json:"servedUpfInfo"`
		ServedUpfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SNssaiUpfInfoList []struct {
						SNssai         string `json:"sNssai"`
						DnnUpfInfoList []struct {
							Dnn               string   `json:"dnn"`
							DnaiList          []string `json:"dnaiList"`
							PduSessionTypes   []string `json:"pduSessionTypes"`
							Ipv4AddressRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"ipv4AddressRanges"`
							Ipv6PrefixRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"ipv6PrefixRanges"`
							NatedIpv4AddressRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"natedIpv4AddressRanges"`
							NatedIpv6PrefixRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"natedIpv6PrefixRanges"`
							Ipv4IndexList      []string `json:"ipv4IndexList"`
							Ipv6IndexList      []string `json:"ipv6IndexList"`
							NetworkInstance    string   `json:"networkInstance"`
							DnaiNwInstanceList struct {
								AdditionalProp1 string `json:"additionalProp1"`
							} `json:"dnaiNwInstanceList"`
						} `json:"dnnUpfInfoList"`
						RedundantTransport bool `json:"redundantTransport"`
					} `json:"sNssaiUpfInfoList"`
					SmfServingArea       []string `json:"smfServingArea"`
					InterfaceUpfInfoList []struct {
						InterfaceType         string   `json:"interfaceType"`
						Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
						Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
						EndpointFqdn          string   `json:"endpointFqdn"`
						NetworkInstance       string   `json:"networkInstance"`
					} `json:"interfaceUpfInfoList"`
					IwkEpsInd       bool     `json:"iwkEpsInd"`
					SxaInd          bool     `json:"sxaInd"`
					PduSessionTypes []string `json:"pduSessionTypes"`
					AtsssCapability string   `json:"atsssCapability"`
					UeIPAddrInd     bool     `json:"ueIpAddrInd"`
					TaiList         []string `json:"taiList"`
					TaiRangeList    []struct {
						PlmnID       string `json:"plmnId"`
						TacRangeList []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"tacRangeList"`
						Nid string `json:"nid"`
					} `json:"taiRangeList"`
					WAgfInfo struct {
						Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
						Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
						EndpointFqdn          string   `json:"endpointFqdn"`
					} `json:"wAgfInfo"`
					TngfInfo struct {
						Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
						Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
						EndpointFqdn          string   `json:"endpointFqdn"`
					} `json:"tngfInfo"`
					TwifInfo struct {
						Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
						Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
						EndpointFqdn          string   `json:"endpointFqdn"`
					} `json:"twifInfo"`
					Priority              int      `json:"priority"`
					RedundantGtpu         bool     `json:"redundantGtpu"`
					Ipups                 bool     `json:"ipups"`
					DataForwarding        bool     `json:"dataForwarding"`
					SupportedPfcpFeatures string   `json:"supportedPfcpFeatures"`
					UpfEvents             []string `json:"upfEvents"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedUpfInfoList"`
		ServedPcfInfo struct {
			AdditionalProp1 struct {
				GroupID    string   `json:"groupId"`
				DnnList    []string `json:"dnnList"`
				SupiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				GpsiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"gpsiRanges"`
				RxDiamHost      string `json:"rxDiamHost"`
				RxDiamRealm     string `json:"rxDiamRealm"`
				V2XSupportInd   bool   `json:"v2xSupportInd"`
				ProseSupportInd bool   `json:"proseSupportInd"`
				ProseCapability struct {
					ProseDirectDiscovey      bool `json:"proseDirectDiscovey"`
					ProseDirectCommunication bool `json:"proseDirectCommunication"`
					ProseL2UetoNetworkRelay  bool `json:"proseL2UetoNetworkRelay"`
					ProseL3UetoNetworkRelay  bool `json:"proseL3UetoNetworkRelay"`
					ProseL2RemoteUe          bool `json:"proseL2RemoteUe"`
					ProseL3RemoteUe          bool `json:"proseL3RemoteUe"`
				} `json:"proseCapability"`
				V2XCapability struct {
					LteV2X bool `json:"lteV2x"`
					NrV2X  bool `json:"nrV2x"`
				} `json:"v2xCapability"`
			} `json:"additionalProp1"`
		} `json:"servedPcfInfo"`
		ServedPcfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					GroupID    string   `json:"groupId"`
					DnnList    []string `json:"dnnList"`
					SupiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					GpsiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"gpsiRanges"`
					RxDiamHost      string `json:"rxDiamHost"`
					RxDiamRealm     string `json:"rxDiamRealm"`
					V2XSupportInd   bool   `json:"v2xSupportInd"`
					ProseSupportInd bool   `json:"proseSupportInd"`
					ProseCapability struct {
						ProseDirectDiscovey      bool `json:"proseDirectDiscovey"`
						ProseDirectCommunication bool `json:"proseDirectCommunication"`
						ProseL2UetoNetworkRelay  bool `json:"proseL2UetoNetworkRelay"`
						ProseL3UetoNetworkRelay  bool `json:"proseL3UetoNetworkRelay"`
						ProseL2RemoteUe          bool `json:"proseL2RemoteUe"`
						ProseL3RemoteUe          bool `json:"proseL3RemoteUe"`
					} `json:"proseCapability"`
					V2XCapability struct {
						LteV2X bool `json:"lteV2x"`
						NrV2X  bool `json:"nrV2x"`
					} `json:"v2xCapability"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedPcfInfoList"`
		ServedBsfInfo struct {
			AdditionalProp1 struct {
				DnnList           []string `json:"dnnList"`
				IPDomainList      []string `json:"ipDomainList"`
				Ipv4AddressRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"ipv4AddressRanges"`
				Ipv6PrefixRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"ipv6PrefixRanges"`
				RxDiamHost  string `json:"rxDiamHost"`
				RxDiamRealm string `json:"rxDiamRealm"`
				GroupID     string `json:"groupId"`
				SupiRanges  []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				GpsiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"gpsiRanges"`
			} `json:"additionalProp1"`
		} `json:"servedBsfInfo"`
		ServedBsfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					DnnList           []string `json:"dnnList"`
					IPDomainList      []string `json:"ipDomainList"`
					Ipv4AddressRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"ipv4AddressRanges"`
					Ipv6PrefixRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"ipv6PrefixRanges"`
					RxDiamHost  string `json:"rxDiamHost"`
					RxDiamRealm string `json:"rxDiamRealm"`
					GroupID     string `json:"groupId"`
					SupiRanges  []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					GpsiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"gpsiRanges"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedBsfInfoList"`
		ServedChfInfo struct {
			AdditionalProp1 struct {
				SupiRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRangeList"`
				GpsiRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"gpsiRangeList"`
				PlmnRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"plmnRangeList"`
				GroupID              string `json:"groupId"`
				PrimaryChfInstance   string `json:"primaryChfInstance"`
				SecondaryChfInstance string `json:"secondaryChfInstance"`
			} `json:"additionalProp1"`
		} `json:"servedChfInfo"`
		ServedChfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SupiRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRangeList"`
					GpsiRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"gpsiRangeList"`
					PlmnRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"plmnRangeList"`
					GroupID              string `json:"groupId"`
					PrimaryChfInstance   string `json:"primaryChfInstance"`
					SecondaryChfInstance string `json:"secondaryChfInstance"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedChfInfoList"`
		ServedNefInfo struct {
			AdditionalProp1 struct {
				NefID   string `json:"nefId"`
				PfdData struct {
					AppIds []string `json:"appIds"`
					AfIds  []string `json:"afIds"`
				} `json:"pfdData"`
				AfEeData struct {
					AfEvents []string `json:"afEvents"`
					AfIds    []string `json:"afIds"`
					AppIds   []string `json:"appIds"`
				} `json:"afEeData"`
				GpsiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"gpsiRanges"`
				ExternalGroupIdentifiersRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"externalGroupIdentifiersRanges"`
				ServedFqdnList []string `json:"servedFqdnList"`
				TaiList        []string `json:"taiList"`
				TaiRangeList   []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
				DnaiList          []string `json:"dnaiList"`
				UnTrustAfInfoList []struct {
					AfID           string `json:"afId"`
					SNssaiInfoList []struct {
						SNssai      string `json:"sNssai"`
						DnnInfoList []struct {
							Dnn string `json:"dnn"`
						} `json:"dnnInfoList"`
					} `json:"sNssaiInfoList"`
					MappingInd bool `json:"mappingInd"`
				} `json:"unTrustAfInfoList"`
				UasNfFunctionalityInd bool `json:"uasNfFunctionalityInd"`
			} `json:"additionalProp1"`
		} `json:"servedNefInfo"`
		ServedNwdafInfo struct {
			AdditionalProp1 struct {
				EventIds     []string `json:"eventIds"`
				NwdafEvents  []string `json:"nwdafEvents"`
				TaiList      []string `json:"taiList"`
				TaiRangeList []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
				NwdafCapability struct {
					AnalyticsAggregation          bool `json:"analyticsAggregation"`
					AnalyticsMetadataProvisioning bool `json:"analyticsMetadataProvisioning"`
				} `json:"nwdafCapability"`
				AnalyticsDelay     string   `json:"analyticsDelay"`
				ServingNfSetIDList []string `json:"servingNfSetIdList"`
				ServingNfTypeList  []string `json:"servingNfTypeList"`
				MlAnalyticsList    []struct {
					MlAnalyticsIds   []string `json:"mlAnalyticsIds"`
					SnssaiList       []string `json:"snssaiList"`
					TrackingAreaList []string `json:"trackingAreaList"`
				} `json:"mlAnalyticsList"`
			} `json:"additionalProp1"`
		} `json:"servedNwdafInfo"`
		ServedNwdafInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					EventIds     []string `json:"eventIds"`
					NwdafEvents  []string `json:"nwdafEvents"`
					TaiList      []string `json:"taiList"`
					TaiRangeList []struct {
						PlmnID       string `json:"plmnId"`
						TacRangeList []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"tacRangeList"`
						Nid string `json:"nid"`
					} `json:"taiRangeList"`
					NwdafCapability struct {
						AnalyticsAggregation          bool `json:"analyticsAggregation"`
						AnalyticsMetadataProvisioning bool `json:"analyticsMetadataProvisioning"`
					} `json:"nwdafCapability"`
					AnalyticsDelay     string   `json:"analyticsDelay"`
					ServingNfSetIDList []string `json:"servingNfSetIdList"`
					ServingNfTypeList  []string `json:"servingNfTypeList"`
					MlAnalyticsList    []struct {
						MlAnalyticsIds   []string `json:"mlAnalyticsIds"`
						SnssaiList       []string `json:"snssaiList"`
						TrackingAreaList []string `json:"trackingAreaList"`
					} `json:"mlAnalyticsList"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedNwdafInfoList"`
		ServedPcscfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					AccessType              []string `json:"accessType"`
					DnnList                 []string `json:"dnnList"`
					GmFqdn                  string   `json:"gmFqdn"`
					GmIpv4Addresses         []string `json:"gmIpv4Addresses"`
					GmIpv6Addresses         []string `json:"gmIpv6Addresses"`
					MwFqdn                  string   `json:"mwFqdn"`
					MwIpv4Addresses         []string `json:"mwIpv4Addresses"`
					MwIpv6Addresses         []string `json:"mwIpv6Addresses"`
					ServedIpv4AddressRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"servedIpv4AddressRanges"`
					ServedIpv6PrefixRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"servedIpv6PrefixRanges"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedPcscfInfoList"`
		ServedGmlcInfo struct {
			AdditionalProp1 struct {
				ServingClientTypes []string `json:"servingClientTypes"`
				GmlcNumbers        []string `json:"gmlcNumbers"`
			} `json:"additionalProp1"`
		} `json:"servedGmlcInfo"`
		ServedLmfInfo struct {
			AdditionalProp1 struct {
				ServingClientTypes []string `json:"servingClientTypes"`
				LmfID              string   `json:"lmfId"`
				ServingAccessTypes []string `json:"servingAccessTypes"`
				ServingAnNodeTypes []string `json:"servingAnNodeTypes"`
				ServingRatTypes    []string `json:"servingRatTypes"`
				TaiList            []string `json:"taiList"`
				TaiRangeList       []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
				SupportedGADShapes []string `json:"supportedGADShapes"`
			} `json:"additionalProp1"`
		} `json:"servedLmfInfo"`
		ServedNfInfo struct {
			AdditionalProp1 struct {
				NfType string `json:"nfType"`
			} `json:"additionalProp1"`
		} `json:"servedNfInfo"`
		ServedHssInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					GroupID    string `json:"groupId"`
					ImsiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"imsiRanges"`
					ImsPrivateIdentityRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"imsPrivateIdentityRanges"`
					ImsPublicIdentityRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"imsPublicIdentityRanges"`
					MsisdnRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"msisdnRanges"`
					ExternalGroupIdentifiersRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"externalGroupIdentifiersRanges"`
					HssDiameterAddress      string   `json:"hssDiameterAddress"`
					AdditionalDiamAddresses []string `json:"additionalDiamAddresses"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedHssInfoList"`
		ServedUdsfInfo struct {
			AdditionalProp1 struct {
				GroupID    string `json:"groupId"`
				SupiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				StorageIDRanges struct {
					AdditionalProp1 []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"additionalProp1"`
				} `json:"storageIdRanges"`
			} `json:"additionalProp1"`
		} `json:"servedUdsfInfo"`
		ServedUdsfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					GroupID    string `json:"groupId"`
					SupiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					StorageIDRanges struct {
						AdditionalProp1 []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"additionalProp1"`
					} `json:"storageIdRanges"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedUdsfInfoList"`
		ServedScpInfoList struct {
			AdditionalProp1 struct {
				ScpDomainInfoList struct {
					AdditionalProp1 struct {
						ScpFqdn        string `json:"scpFqdn"`
						ScpIPEndPoints []struct {
							Ipv4Address string `json:"ipv4Address"`
							Ipv6Address string `json:"ipv6Address"`
							Transport   string `json:"transport"`
							Port        int    `json:"port"`
						} `json:"scpIpEndPoints"`
						ScpPrefix string `json:"scpPrefix"`
						ScpPorts  struct {
							AdditionalProp1 int `json:"additionalProp1"`
						} `json:"scpPorts"`
					} `json:"additionalProp1"`
				} `json:"scpDomainInfoList"`
				ScpPrefix string `json:"scpPrefix"`
				ScpPorts  struct {
					AdditionalProp1 int `json:"additionalProp1"`
				} `json:"scpPorts"`
				AddressDomains []string `json:"addressDomains"`
				Ipv4Addresses  []string `json:"ipv4Addresses"`
				Ipv6Prefixes   []string `json:"ipv6Prefixes"`
				Ipv4AddrRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"ipv4AddrRanges"`
				Ipv6PrefixRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"ipv6PrefixRanges"`
				ServedNfSetIDList []string `json:"servedNfSetIdList"`
				RemotePlmnList    []string `json:"remotePlmnList"`
				RemoteSnpnList    []string `json:"remoteSnpnList"`
				IPReachability    string   `json:"ipReachability"`
				ScpCapabilities   []string `json:"scpCapabilities"`
			} `json:"additionalProp1"`
		} `json:"servedScpInfoList"`
		ServedSeppInfoList struct {
			AdditionalProp1 struct {
				SeppPrefix string `json:"seppPrefix"`
				SeppPorts  struct {
					AdditionalProp1 int `json:"additionalProp1"`
				} `json:"seppPorts"`
				RemotePlmnList []string `json:"remotePlmnList"`
				RemoteSnpnList []string `json:"remoteSnpnList"`
				N32Purposes    []string `json:"n32Purposes"`
			} `json:"additionalProp1"`
		} `json:"servedSeppInfoList"`
		ServedAanfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					RoutingIndicators []string `json:"routingIndicators"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
			AdditionalProp2 struct {
				AdditionalProp1 struct {
					RoutingIndicators []string `json:"routingIndicators"`
				} `json:"additionalProp1"`
			} `json:"additionalProp2"`
			AdditionalProp3 struct {
				AdditionalProp1 struct {
					RoutingIndicators []string `json:"routingIndicators"`
				} `json:"additionalProp1"`
			} `json:"additionalProp3"`
		} `json:"servedAanfInfoList"`
		Served5GDdnmfInfo struct {
			AdditionalProp1 struct {
				PlmnID string `json:"plmnId"`
			} `json:"additionalProp1"`
		} `json:"served5gDdnmfInfo"`
		ServedMfafInfoList struct {
			AdditionalProp1 struct {
				ServingNfTypeList  []string `json:"servingNfTypeList"`
				ServingNfSetIDList []string `json:"servingNfSetIdList"`
				TaiList            []string `json:"taiList"`
				TaiRangeList       []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
			} `json:"additionalProp1"`
		} `json:"servedMfafInfoList"`
		ServedEasdfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SNssaiEasdfInfoList []struct {
						SNssai           string `json:"sNssai"`
						DnnEasdfInfoList []struct {
							Dnn      string   `json:"dnn"`
							DnaiList []string `json:"dnaiList"`
						} `json:"dnnEasdfInfoList"`
					} `json:"sNssaiEasdfInfoList"`
					EasdfN6IPAddressList []string `json:"easdfN6IpAddressList"`
					UpfN6IPAddressList   []string `json:"upfN6IpAddressList"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
			AdditionalProp2 struct {
				AdditionalProp1 struct {
					SNssaiEasdfInfoList []struct {
						SNssai           string `json:"sNssai"`
						DnnEasdfInfoList []struct {
							Dnn      string   `json:"dnn"`
							DnaiList []string `json:"dnaiList"`
						} `json:"dnnEasdfInfoList"`
					} `json:"sNssaiEasdfInfoList"`
					EasdfN6IPAddressList []string `json:"easdfN6IpAddressList"`
					UpfN6IPAddressList   []string `json:"upfN6IpAddressList"`
				} `json:"additionalProp1"`
			} `json:"additionalProp2"`
			AdditionalProp3 struct {
				AdditionalProp1 struct {
					SNssaiEasdfInfoList []struct {
						SNssai           string `json:"sNssai"`
						DnnEasdfInfoList []struct {
							Dnn      string   `json:"dnn"`
							DnaiList []string `json:"dnaiList"`
						} `json:"dnnEasdfInfoList"`
					} `json:"sNssaiEasdfInfoList"`
					EasdfN6IPAddressList []string `json:"easdfN6IpAddressList"`
					UpfN6IPAddressList   []string `json:"upfN6IpAddressList"`
				} `json:"additionalProp1"`
			} `json:"additionalProp3"`
		} `json:"servedEasdfInfoList"`
		ServedDccfInfoList struct {
			AdditionalProp1 struct {
				ServingNfTypeList  []string `json:"servingNfTypeList"`
				ServingNfSetIDList []string `json:"servingNfSetIdList"`
				TaiList            []string `json:"taiList"`
				TaiRangeList       []struct {
					PlmnID       string `json:"plmnId"`
					TacRangeList []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"tacRangeList"`
					Nid string `json:"nid"`
				} `json:"taiRangeList"`
			} `json:"additionalProp1"`
		} `json:"servedDccfInfoList"`
		ServedMbSmfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SNssaiInfoList struct {
						AdditionalProp1 struct {
							SNssai      string `json:"sNssai"`
							DnnInfoList []struct {
								Dnn string `json:"dnn"`
							} `json:"dnnInfoList"`
						} `json:"additionalProp1"`
					} `json:"sNssaiInfoList"`
					TmgiRangeList struct {
						AdditionalProp1 struct {
							MbsServiceIDStart string `json:"mbsServiceIdStart"`
							MbsServiceIDEnd   string `json:"mbsServiceIdEnd"`
							PlmnID            string `json:"plmnId"`
							Nid               string `json:"nid"`
						} `json:"additionalProp1"`
					} `json:"tmgiRangeList"`
					TaiList      []string `json:"taiList"`
					TaiRangeList []struct {
						PlmnID       string `json:"plmnId"`
						TacRangeList []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"tacRangeList"`
						Nid string `json:"nid"`
					} `json:"taiRangeList"`
					MbsSessionList struct {
						AdditionalProp1 struct {
							MbsSessionID    string `json:"mbsSessionId"`
							MbsAreaSessions struct {
								AdditionalProp1 string `json:"additionalProp1"`
							} `json:"mbsAreaSessions"`
						} `json:"additionalProp1"`
					} `json:"mbsSessionList"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedMbSmfInfoList"`
		ServedTsctsfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SNssaiInfoList struct {
						AdditionalProp1 struct {
							SNssai      string `json:"sNssai"`
							DnnInfoList []struct {
								Dnn string `json:"dnn"`
							} `json:"dnnInfoList"`
						} `json:"additionalProp1"`
					} `json:"sNssaiInfoList"`
					ExternalGroupIdentifiersRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"externalGroupIdentifiersRanges"`
					SupiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"supiRanges"`
					GpsiRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"gpsiRanges"`
					InternalGroupIdentifiersRanges []struct {
						Start   string `json:"start"`
						End     string `json:"end"`
						Pattern string `json:"pattern"`
					} `json:"internalGroupIdentifiersRanges"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedTsctsfInfoList"`
		ServedMbUpfInfoList struct {
			AdditionalProp1 struct {
				AdditionalProp1 struct {
					SNssaiMbUpfInfoList []struct {
						SNssai         string `json:"sNssai"`
						DnnUpfInfoList []struct {
							Dnn               string   `json:"dnn"`
							DnaiList          []string `json:"dnaiList"`
							PduSessionTypes   []string `json:"pduSessionTypes"`
							Ipv4AddressRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"ipv4AddressRanges"`
							Ipv6PrefixRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"ipv6PrefixRanges"`
							NatedIpv4AddressRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"natedIpv4AddressRanges"`
							NatedIpv6PrefixRanges []struct {
								Start string `json:"start"`
								End   string `json:"end"`
							} `json:"natedIpv6PrefixRanges"`
							Ipv4IndexList      []string `json:"ipv4IndexList"`
							Ipv6IndexList      []string `json:"ipv6IndexList"`
							NetworkInstance    string   `json:"networkInstance"`
							DnaiNwInstanceList struct {
								AdditionalProp1 string `json:"additionalProp1"`
							} `json:"dnaiNwInstanceList"`
						} `json:"dnnUpfInfoList"`
						RedundantTransport bool `json:"redundantTransport"`
					} `json:"sNssaiMbUpfInfoList"`
					MbSmfServingArea       []string `json:"mbSmfServingArea"`
					InterfaceMbUpfInfoList []struct {
						InterfaceType         string   `json:"interfaceType"`
						Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
						Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
						EndpointFqdn          string   `json:"endpointFqdn"`
						NetworkInstance       string   `json:"networkInstance"`
					} `json:"interfaceMbUpfInfoList"`
					TaiList      []string `json:"taiList"`
					TaiRangeList []struct {
						PlmnID       string `json:"plmnId"`
						TacRangeList []struct {
							Start   string `json:"start"`
							End     string `json:"end"`
							Pattern string `json:"pattern"`
						} `json:"tacRangeList"`
						Nid string `json:"nid"`
					} `json:"taiRangeList"`
					Priority              int    `json:"priority"`
					SupportedPfcpFeatures string `json:"supportedPfcpFeatures"`
				} `json:"additionalProp1"`
			} `json:"additionalProp1"`
		} `json:"servedMbUpfInfoList"`
		ServedTrustAfInfo struct {
			AdditionalProp1 struct {
				SNssaiInfoList []struct {
					SNssai      string `json:"sNssai"`
					DnnInfoList []struct {
						Dnn string `json:"dnn"`
					} `json:"dnnInfoList"`
				} `json:"sNssaiInfoList"`
				AfEvents        []string `json:"afEvents"`
				AppIds          []string `json:"appIds"`
				InternalGroupID []string `json:"internalGroupId"`
				MappingInd      bool     `json:"mappingInd"`
			} `json:"additionalProp1"`
		} `json:"servedTrustAfInfo"`
		ServedNssaafInfo struct {
			AdditionalProp1 struct {
				SupiRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"supiRanges"`
				InternalGroupIdentifiersRanges []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"internalGroupIdentifiersRanges"`
			} `json:"additionalProp1"`
		} `json:"servedNssaafInfo"`
	} `json:"nrfInfo"`
	UdsfInfo struct {
		GroupID    string `json:"groupId"`
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		StorageIDRanges struct {
			AdditionalProp1 []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"additionalProp1"`
		} `json:"storageIdRanges"`
	} `json:"udsfInfo"`
	UdsfInfoList struct {
		AdditionalProp1 struct {
			GroupID    string `json:"groupId"`
			SupiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			StorageIDRanges struct {
				AdditionalProp1 []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"additionalProp1"`
			} `json:"storageIdRanges"`
		} `json:"additionalProp1"`
	} `json:"udsfInfoList"`
	NwdafInfo struct {
		EventIds     []string `json:"eventIds"`
		NwdafEvents  []string `json:"nwdafEvents"`
		TaiList      []string `json:"taiList"`
		TaiRangeList []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		NwdafCapability struct {
			AnalyticsAggregation          bool `json:"analyticsAggregation"`
			AnalyticsMetadataProvisioning bool `json:"analyticsMetadataProvisioning"`
		} `json:"nwdafCapability"`
		AnalyticsDelay     string   `json:"analyticsDelay"`
		ServingNfSetIDList []string `json:"servingNfSetIdList"`
		ServingNfTypeList  []string `json:"servingNfTypeList"`
		MlAnalyticsList    []struct {
			MlAnalyticsIds   []string `json:"mlAnalyticsIds"`
			SnssaiList       []string `json:"snssaiList"`
			TrackingAreaList []string `json:"trackingAreaList"`
		} `json:"mlAnalyticsList"`
	} `json:"nwdafInfo"`
	NwdafInfoList struct {
		AdditionalProp1 struct {
			EventIds     []string `json:"eventIds"`
			NwdafEvents  []string `json:"nwdafEvents"`
			TaiList      []string `json:"taiList"`
			TaiRangeList []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
			NwdafCapability struct {
				AnalyticsAggregation          bool `json:"analyticsAggregation"`
				AnalyticsMetadataProvisioning bool `json:"analyticsMetadataProvisioning"`
			} `json:"nwdafCapability"`
			AnalyticsDelay     string   `json:"analyticsDelay"`
			ServingNfSetIDList []string `json:"servingNfSetIdList"`
			ServingNfTypeList  []string `json:"servingNfTypeList"`
			MlAnalyticsList    []struct {
				MlAnalyticsIds   []string `json:"mlAnalyticsIds"`
				SnssaiList       []string `json:"snssaiList"`
				TrackingAreaList []string `json:"trackingAreaList"`
			} `json:"mlAnalyticsList"`
		} `json:"additionalProp1"`
	} `json:"nwdafInfoList"`
	PcscfInfoList struct {
		AdditionalProp1 struct {
			AccessType              []string `json:"accessType"`
			DnnList                 []string `json:"dnnList"`
			GmFqdn                  string   `json:"gmFqdn"`
			GmIpv4Addresses         []string `json:"gmIpv4Addresses"`
			GmIpv6Addresses         []string `json:"gmIpv6Addresses"`
			MwFqdn                  string   `json:"mwFqdn"`
			MwIpv4Addresses         []string `json:"mwIpv4Addresses"`
			MwIpv6Addresses         []string `json:"mwIpv6Addresses"`
			ServedIpv4AddressRanges []struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"servedIpv4AddressRanges"`
			ServedIpv6PrefixRanges []struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"servedIpv6PrefixRanges"`
		} `json:"additionalProp1"`
	} `json:"pcscfInfoList"`
	HssInfoList struct {
		AdditionalProp1 struct {
			GroupID    string `json:"groupId"`
			ImsiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"imsiRanges"`
			ImsPrivateIdentityRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"imsPrivateIdentityRanges"`
			ImsPublicIdentityRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"imsPublicIdentityRanges"`
			MsisdnRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"msisdnRanges"`
			ExternalGroupIdentifiersRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"externalGroupIdentifiersRanges"`
			HssDiameterAddress      string   `json:"hssDiameterAddress"`
			AdditionalDiamAddresses []string `json:"additionalDiamAddresses"`
		} `json:"additionalProp1"`
	} `json:"hssInfoList"`
	CustomInfo struct {
	} `json:"customInfo"`
	RecoveryTime         string `json:"recoveryTime"`
	NfServicePersistence bool   `json:"nfServicePersistence"`
	NfServiceList        struct {
		AdditionalProp1 struct {
			ServiceInstanceID string `json:"serviceInstanceId"`
			ServiceName       string `json:"serviceName"`
			Versions          []struct {
				APIVersionInURI string `json:"apiVersionInUri"`
				APIFullVersion  string `json:"apiFullVersion"`
				Expiry          string `json:"expiry"`
			} `json:"versions"`
			Scheme          string `json:"scheme"`
			NfServiceStatus string `json:"nfServiceStatus"`
			Fqdn            string `json:"fqdn"`
			InterPlmnFqdn   string `json:"interPlmnFqdn"`
			IPEndPoints     []struct {
				Ipv4Address string `json:"ipv4Address"`
				Ipv6Address string `json:"ipv6Address"`
				Transport   string `json:"transport"`
				Port        int    `json:"port"`
			} `json:"ipEndPoints"`
			APIPrefix                        string `json:"apiPrefix"`
			DefaultNotificationSubscriptions []struct {
				NotificationType   string   `json:"notificationType"`
				CallbackURI        string   `json:"callbackUri"`
				N1MessageClass     string   `json:"n1MessageClass"`
				N2InformationClass string   `json:"n2InformationClass"`
				Versions           []string `json:"versions"`
				Binding            string   `json:"binding"`
				AcceptedEncoding   string   `json:"acceptedEncoding"`
				SupportedFeatures  string   `json:"supportedFeatures"`
				ServiceInfoList    struct {
					AdditionalProp1 struct {
						Versions          []string `json:"versions"`
						SupportedFeatures string   `json:"supportedFeatures"`
					} `json:"additionalProp1"`
				} `json:"serviceInfoList"`
			} `json:"defaultNotificationSubscriptions"`
			AllowedPlmns               []string `json:"allowedPlmns"`
			AllowedSnpns               []string `json:"allowedSnpns"`
			AllowedNfTypes             []string `json:"allowedNfTypes"`
			AllowedNfDomains           []string `json:"allowedNfDomains"`
			AllowedNssais              []string `json:"allowedNssais"`
			AllowedOperationsPerNfType struct {
				AdditionalProp1 []string `json:"additionalProp1"`
			} `json:"allowedOperationsPerNfType"`
			AllowedOperationsPerNfInstance struct {
				AdditionalProp1 []string `json:"additionalProp1"`
			} `json:"allowedOperationsPerNfInstance"`
			AllowedOperationsPerNfInstanceOverrides bool     `json:"allowedOperationsPerNfInstanceOverrides"`
			Priority                                int      `json:"priority"`
			Capacity                                int      `json:"capacity"`
			Load                                    int      `json:"load"`
			LoadTimeStamp                           string   `json:"loadTimeStamp"`
			RecoveryTime                            string   `json:"recoveryTime"`
			SupportedFeatures                       string   `json:"supportedFeatures"`
			NfServiceSetIDList                      []string `json:"nfServiceSetIdList"`
			SNssais                                 []string `json:"sNssais"`
			PerPlmnSnssaiList                       []struct {
				PlmnID     string   `json:"plmnId"`
				SNssaiList []string `json:"sNssaiList"`
				Nid        string   `json:"nid"`
			} `json:"perPlmnSnssaiList"`
			VendorID                        string `json:"vendorId"`
			SupportedVendorSpecificFeatures struct {
				AdditionalProp1 []struct {
					FeatureName    string `json:"featureName"`
					FeatureVersion string `json:"featureVersion"`
				} `json:"additionalProp1"`
			} `json:"supportedVendorSpecificFeatures"`
			Oauth2Required       bool `json:"oauth2Required"`
			PerPlmnOauth2ReqList struct {
				Oauth2RequiredPlmnIDList    []string `json:"oauth2RequiredPlmnIdList"`
				Oauth2NotRequiredPlmnIDList []string `json:"oauth2NotRequiredPlmnIdList"`
			} `json:"perPlmnOauth2ReqList"`
		} `json:"additionalProp1"`
	} `json:"nfServiceList"`
	NfProfileChangesSupportInd       bool `json:"nfProfileChangesSupportInd"`
	DefaultNotificationSubscriptions []struct {
		NotificationType   string   `json:"notificationType"`
		CallbackURI        string   `json:"callbackUri"`
		N1MessageClass     string   `json:"n1MessageClass"`
		N2InformationClass string   `json:"n2InformationClass"`
		Versions           []string `json:"versions"`
		Binding            string   `json:"binding"`
		AcceptedEncoding   string   `json:"acceptedEncoding"`
		SupportedFeatures  string   `json:"supportedFeatures"`
		ServiceInfoList    struct {
			AdditionalProp1 struct {
				Versions          []string `json:"versions"`
				SupportedFeatures string   `json:"supportedFeatures"`
			} `json:"additionalProp1"`
		} `json:"serviceInfoList"`
	} `json:"defaultNotificationSubscriptions"`
	LmfInfo struct {
		ServingClientTypes []string `json:"servingClientTypes"`
		LmfID              string   `json:"lmfId"`
		ServingAccessTypes []string `json:"servingAccessTypes"`
		ServingAnNodeTypes []string `json:"servingAnNodeTypes"`
		ServingRatTypes    []string `json:"servingRatTypes"`
		TaiList            []string `json:"taiList"`
		TaiRangeList       []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		SupportedGADShapes []string `json:"supportedGADShapes"`
	} `json:"lmfInfo"`
	GmlcInfo struct {
		ServingClientTypes []string `json:"servingClientTypes"`
		GmlcNumbers        []string `json:"gmlcNumbers"`
	} `json:"gmlcInfo"`
	NfSetIDList           []string `json:"nfSetIdList"`
	ServingScope          []string `json:"servingScope"`
	LcHSupportInd         bool     `json:"lcHSupportInd"`
	OlcHSupportInd        bool     `json:"olcHSupportInd"`
	NfSetRecoveryTimeList struct {
		AdditionalProp1 string `json:"additionalProp1"`
	} `json:"nfSetRecoveryTimeList"`
	ServiceSetRecoveryTimeList struct {
		AdditionalProp1 string `json:"additionalProp1"`
	} `json:"serviceSetRecoveryTimeList"`
	ScpDomains []string `json:"scpDomains"`
	ScpInfo    struct {
		ScpDomainInfoList struct {
			AdditionalProp1 struct {
				ScpFqdn        string `json:"scpFqdn"`
				ScpIPEndPoints []struct {
					Ipv4Address string `json:"ipv4Address"`
					Ipv6Address string `json:"ipv6Address"`
					Transport   string `json:"transport"`
					Port        int    `json:"port"`
				} `json:"scpIpEndPoints"`
				ScpPrefix string `json:"scpPrefix"`
				ScpPorts  struct {
					AdditionalProp1 int `json:"additionalProp1"`
				} `json:"scpPorts"`
			} `json:"additionalProp1"`
		} `json:"scpDomainInfoList"`
		ScpPrefix string `json:"scpPrefix"`
		ScpPorts  struct {
			AdditionalProp1 int `json:"additionalProp1"`
		} `json:"scpPorts"`
		AddressDomains []string `json:"addressDomains"`
		Ipv4Addresses  []string `json:"ipv4Addresses"`
		Ipv6Prefixes   []string `json:"ipv6Prefixes"`
		Ipv4AddrRanges []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"ipv4AddrRanges"`
		Ipv6PrefixRanges []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"ipv6PrefixRanges"`
		ServedNfSetIDList []string `json:"servedNfSetIdList"`
		RemotePlmnList    []string `json:"remotePlmnList"`
		RemoteSnpnList    []string `json:"remoteSnpnList"`
		IPReachability    string   `json:"ipReachability"`
		ScpCapabilities   []string `json:"scpCapabilities"`
	} `json:"scpInfo"`
	SeppInfo struct {
		SeppPrefix string `json:"seppPrefix"`
		SeppPorts  struct {
			AdditionalProp1 int `json:"additionalProp1"`
		} `json:"seppPorts"`
		RemotePlmnList []string `json:"remotePlmnList"`
		RemoteSnpnList []string `json:"remoteSnpnList"`
		N32Purposes    []string `json:"n32Purposes"`
	} `json:"seppInfo"`
	VendorID                        string `json:"vendorId"`
	SupportedVendorSpecificFeatures struct {
		AdditionalProp1 []struct {
			FeatureName    string `json:"featureName"`
			FeatureVersion string `json:"featureVersion"`
		} `json:"additionalProp1"`
	} `json:"supportedVendorSpecificFeatures"`
	AanfInfoList struct {
		AdditionalProp1 struct {
			RoutingIndicators []string `json:"routingIndicators"`
		} `json:"additionalProp1"`
	} `json:"aanfInfoList"`
	FiveGDdnmfInfo struct {
		PlmnID string `json:"plmnId"`
	} `json:"5gDdnmfInfo"`
	MfafInfo struct {
		ServingNfTypeList  []string `json:"servingNfTypeList"`
		ServingNfSetIDList []string `json:"servingNfSetIdList"`
		TaiList            []string `json:"taiList"`
		TaiRangeList       []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
	} `json:"mfafInfo"`
	EasdfInfoList struct {
		AdditionalProp1 struct {
			SNssaiEasdfInfoList []struct {
				SNssai           string `json:"sNssai"`
				DnnEasdfInfoList []struct {
					Dnn      string   `json:"dnn"`
					DnaiList []string `json:"dnaiList"`
				} `json:"dnnEasdfInfoList"`
			} `json:"sNssaiEasdfInfoList"`
			EasdfN6IPAddressList []string `json:"easdfN6IpAddressList"`
			UpfN6IPAddressList   []string `json:"upfN6IpAddressList"`
		} `json:"additionalProp1"`
	} `json:"easdfInfoList"`
	DccfInfo struct {
		ServingNfTypeList  []string `json:"servingNfTypeList"`
		ServingNfSetIDList []string `json:"servingNfSetIdList"`
		TaiList            []string `json:"taiList"`
		TaiRangeList       []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
	} `json:"dccfInfo"`
	NsacfInfoList struct {
		AdditionalProp1 struct {
			NsacfCapability struct {
				SupportUeSAC  bool `json:"supportUeSAC"`
				SupportPduSAC bool `json:"supportPduSAC"`
			} `json:"nsacfCapability"`
			TaiList      []string `json:"taiList"`
			TaiRangeList []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
		} `json:"additionalProp1"`
	} `json:"nsacfInfoList"`
	MbSmfInfoList struct {
		AdditionalProp1 struct {
			SNssaiInfoList struct {
				AdditionalProp1 struct {
					SNssai      string `json:"sNssai"`
					DnnInfoList []struct {
						Dnn string `json:"dnn"`
					} `json:"dnnInfoList"`
				} `json:"additionalProp1"`
			} `json:"sNssaiInfoList"`
			TmgiRangeList struct {
				AdditionalProp1 struct {
					MbsServiceIDStart string `json:"mbsServiceIdStart"`
					MbsServiceIDEnd   string `json:"mbsServiceIdEnd"`
					PlmnID            string `json:"plmnId"`
					Nid               string `json:"nid"`
				} `json:"additionalProp1"`
			} `json:"tmgiRangeList"`
			TaiList      []string `json:"taiList"`
			TaiRangeList []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
			MbsSessionList struct {
				AdditionalProp1 struct {
					MbsSessionID    string `json:"mbsSessionId"`
					MbsAreaSessions struct {
						AdditionalProp1 string `json:"additionalProp1"`
					} `json:"mbsAreaSessions"`
				} `json:"additionalProp1"`
			} `json:"mbsSessionList"`
		} `json:"additionalProp1"`
	} `json:"mbSmfInfoList"`
	TsctsfInfoList struct {
		AdditionalProp1 struct {
			SNssaiInfoList struct {
				AdditionalProp1 struct {
					SNssai      string `json:"sNssai"`
					DnnInfoList []struct {
						Dnn string `json:"dnn"`
					} `json:"dnnInfoList"`
				} `json:"additionalProp1"`
			} `json:"sNssaiInfoList"`
			ExternalGroupIdentifiersRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"externalGroupIdentifiersRanges"`
			SupiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"supiRanges"`
			GpsiRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"gpsiRanges"`
			InternalGroupIdentifiersRanges []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"internalGroupIdentifiersRanges"`
		} `json:"additionalProp1"`
	} `json:"tsctsfInfoList"`
	MbUpfInfoList struct {
		AdditionalProp1 struct {
			SNssaiMbUpfInfoList []struct {
				SNssai         string `json:"sNssai"`
				DnnUpfInfoList []struct {
					Dnn               string   `json:"dnn"`
					DnaiList          []string `json:"dnaiList"`
					PduSessionTypes   []string `json:"pduSessionTypes"`
					Ipv4AddressRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"ipv4AddressRanges"`
					Ipv6PrefixRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"ipv6PrefixRanges"`
					NatedIpv4AddressRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"natedIpv4AddressRanges"`
					NatedIpv6PrefixRanges []struct {
						Start string `json:"start"`
						End   string `json:"end"`
					} `json:"natedIpv6PrefixRanges"`
					Ipv4IndexList      []string `json:"ipv4IndexList"`
					Ipv6IndexList      []string `json:"ipv6IndexList"`
					NetworkInstance    string   `json:"networkInstance"`
					DnaiNwInstanceList struct {
						AdditionalProp1 string `json:"additionalProp1"`
					} `json:"dnaiNwInstanceList"`
				} `json:"dnnUpfInfoList"`
				RedundantTransport bool `json:"redundantTransport"`
			} `json:"sNssaiMbUpfInfoList"`
			MbSmfServingArea       []string `json:"mbSmfServingArea"`
			InterfaceMbUpfInfoList []struct {
				InterfaceType         string   `json:"interfaceType"`
				Ipv4EndpointAddresses []string `json:"ipv4EndpointAddresses"`
				Ipv6EndpointAddresses []string `json:"ipv6EndpointAddresses"`
				EndpointFqdn          string   `json:"endpointFqdn"`
				NetworkInstance       string   `json:"networkInstance"`
			} `json:"interfaceMbUpfInfoList"`
			TaiList      []string `json:"taiList"`
			TaiRangeList []struct {
				PlmnID       string `json:"plmnId"`
				TacRangeList []struct {
					Start   string `json:"start"`
					End     string `json:"end"`
					Pattern string `json:"pattern"`
				} `json:"tacRangeList"`
				Nid string `json:"nid"`
			} `json:"taiRangeList"`
			Priority              int    `json:"priority"`
			SupportedPfcpFeatures string `json:"supportedPfcpFeatures"`
		} `json:"additionalProp1"`
	} `json:"mbUpfInfoList"`
	TrustAfInfo struct {
		SNssaiInfoList []struct {
			SNssai      string `json:"sNssai"`
			DnnInfoList []struct {
				Dnn string `json:"dnn"`
			} `json:"dnnInfoList"`
		} `json:"sNssaiInfoList"`
		AfEvents        []string `json:"afEvents"`
		AppIds          []string `json:"appIds"`
		InternalGroupID []string `json:"internalGroupId"`
		MappingInd      bool     `json:"mappingInd"`
	} `json:"trustAfInfo"`
	NssaafInfo struct {
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		InternalGroupIdentifiersRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"internalGroupIdentifiersRanges"`
	} `json:"nssaafInfo"`
	HniList   []string `json:"hniList"`
	IwmscInfo struct {
		MsisdnRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"msisdnRanges"`
		SupiRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"supiRanges"`
		TaiRangeList []struct {
			PlmnID       string `json:"plmnId"`
			TacRangeList []struct {
				Start   string `json:"start"`
				End     string `json:"end"`
				Pattern string `json:"pattern"`
			} `json:"tacRangeList"`
			Nid string `json:"nid"`
		} `json:"taiRangeList"`
		ScNumber string `json:"scNumber"`
	} `json:"iwmscInfo"`
	MnpfInfo struct {
		MsisdnRanges []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"msisdnRanges"`
	} `json:"mnpfInfo"`
	SmsfInfo struct {
		RoamingUeInd        bool `json:"roamingUeInd"`
		RemotePlmnRangeList []struct {
			Start   string `json:"start"`
			End     string `json:"end"`
			Pattern string `json:"pattern"`
		} `json:"remotePlmnRangeList"`
	} `json:"smsfInfo"`
}
