package main

import "time"

type ChargingDataRequest struct {
	SubscriberIdentifier     string    `json:"subscriberIdentifier"`
	TenantIdentifier         string    `json:"tenantIdentifier"`
	ChargingID               int       `json:"chargingId"`
	MnSConsumerIdentifier    string    `json:"mnSConsumerIdentifier"`
	NfConsumerIdentification string    `json:"nfConsumerIdentification"`
	InvocationTimeStamp      time.Time `json:"invocationTimeStamp"`
	InvocationSequenceNumber int       `json:"invocationSequenceNumber"`
	RetransmissionIndicator  bool      `json:"retransmissionIndicator"`
	OneTimeEvent             bool      `json:"oneTimeEvent"`
	OneTimeEventType         struct {
		OneTimeEventType string `json:"oneTimeEventType"`
	} `json:"oneTimeEventType"`
	NotifyURI                string `json:"notifyUri"`
	SupportedFeatures        string `json:"supportedFeatures"`
	ServiceSpecificationInfo string `json:"serviceSpecificationInfo"`
	MultipleUnitUsage        []struct {
		RatingGroup   int `json:"ratingGroup"`
		RequestedUnit struct {
			Time                 int `json:"time"`
			TotalVolume          int `json:"totalVolume"`
			UplinkVolume         int `json:"uplinkVolume"`
			DownlinkVolume       int `json:"downlinkVolume"`
			ServiceSpecificUnits int `json:"serviceSpecificUnits"`
		} `json:"requestedUnit"`
		UsedUnitContainer []struct {
			ServiceID                int `json:"serviceId"`
			QuotaManagementIndicator struct {
				QuotaManagementIndicator string `json:"quotaManagementIndicator"`
			} `json:"quotaManagementIndicator"`
			Triggers []struct {
				TriggerType     string `json:"triggerType"`
				TriggerCategory struct {
					TriggerCategory string `json:"triggerCategory"`
				} `json:"triggerCategory"`
				TimeLimit        int       `json:"timeLimit"`
				VolumeLimit      int       `json:"volumeLimit"`
				VolumeLimit64    int       `json:"volumeLimit64"`
				EventLimit       int       `json:"eventLimit"`
				MaxNumberOfccc   int       `json:"maxNumberOfccc"`
				TariffTimeChange time.Time `json:"tariffTimeChange"`
			} `json:"triggers"`
			TriggerTimestamp        time.Time   `json:"triggerTimestamp"`
			Time                    int         `json:"time"`
			TotalVolume             int         `json:"totalVolume"`
			UplinkVolume            int         `json:"uplinkVolume"`
			DownlinkVolume          int         `json:"downlinkVolume"`
			ServiceSpecificUnits    int         `json:"serviceSpecificUnits"`
			EventTimeStamps         []time.Time `json:"eventTimeStamps"`
			LocalSequenceNumber     int         `json:"localSequenceNumber"`
			PDUContainerInformation struct {
				TimeofFirstUsage time.Time `json:"timeofFirstUsage"`
				TimeofLastUsage  time.Time `json:"timeofLastUsage"`
				QoSInformation   struct {
					QosID   string `json:"qosId"`
					FiveQi  int    `json:"5qi"`
					MaxbrUl string `json:"maxbrUl"`
					MaxbrDl string `json:"maxbrDl"`
					GbrUl   string `json:"gbrUl"`
					GbrDl   string `json:"gbrDl"`
					Arp     struct {
						PriorityLevel int `json:"priorityLevel"`
						PreemptCap    struct {
							PreemptCap string `json:"preemptCap"`
						} `json:"preemptCap"`
						PreemptVuln struct {
							PreemptVuln string `json:"preemptVuln"`
						} `json:"preemptVuln"`
					} `json:"arp"`
					Qnc                  bool   `json:"qnc"`
					PriorityLevel        int    `json:"priorityLevel"`
					AverWindow           int    `json:"averWindow"`
					MaxDataBurstVol      int    `json:"maxDataBurstVol"`
					ReflectiveQos        bool   `json:"reflectiveQos"`
					SharingKeyDl         string `json:"sharingKeyDl"`
					SharingKeyUl         string `json:"sharingKeyUl"`
					MaxPacketLossRateDl  int    `json:"maxPacketLossRateDl"`
					MaxPacketLossRateUl  int    `json:"maxPacketLossRateUl"`
					DefQosFlowIndication bool   `json:"defQosFlowIndication"`
				} `json:"qoSInformation"`
				QoSCharacteristics struct {
					FiveQi       int `json:"5qi"`
					ResourceType struct {
						ResourceType string `json:"resourceType"`
					} `json:"resourceType"`
					PriorityLevel     int    `json:"priorityLevel"`
					PacketDelayBudget int    `json:"packetDelayBudget"`
					PacketErrorRate   string `json:"packetErrorRate"`
					AveragingWindow   int    `json:"averagingWindow"`
					MaxDataBurstVol   int    `json:"maxDataBurstVol"`
				} `json:"qoSCharacteristics"`
				AfChargingIdentifier    int `json:"afChargingIdentifier"`
				UserLocationInformation struct {
					EutraLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ecgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalNgenbID            struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalNgenbId"`
					} `json:"eutraLocation"`
					NrLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ncgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalGnbID              struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalGnbId"`
					} `json:"nrLocation"`
					N3GaLocation struct {
						N3GppTai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"n3gppTai"`
						N3IwfID    string `json:"n3IwfId"`
						UeIpv4Addr string `json:"ueIpv4Addr"`
						UeIpv6Addr struct {
							UeIpv6Addr string `json:"ueIpv6Addr"`
						} `json:"ueIpv6Addr"`
						PortNumber int `json:"portNumber"`
					} `json:"n3gaLocation"`
				} `json:"userLocationInformation"`
				UetimeZone string `json:"uetimeZone"`
				RATType    struct {
					RATType string `json:"rATType"`
				} `json:"rATType"`
				ServingNodeID []struct {
					ServingNetworkFunctionInformation string `json:"servingNetworkFunctionInformation"`
					AMFID                             string `json:"aMFId"`
				} `json:"servingNodeID"`
				PresenceReportingAreaInformation struct {
					AdditionalProp1 struct {
						PraID         string `json:"praId"`
						PresenceState struct {
							PresenceState string `json:"presenceState"`
						} `json:"presenceState"`
						TrackingAreaList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"trackingAreaList"`
						EcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgiList"`
						NcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgiList"`
						GlobalRanNodeIDList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalRanNodeIdList"`
					} `json:"additionalProp1"`
					AdditionalProp2 struct {
						PraID         string `json:"praId"`
						PresenceState struct {
							PresenceState string `json:"presenceState"`
						} `json:"presenceState"`
						TrackingAreaList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"trackingAreaList"`
						EcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgiList"`
						NcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgiList"`
						GlobalRanNodeIDList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalRanNodeIdList"`
					} `json:"additionalProp2"`
					AdditionalProp3 struct {
						PraID         string `json:"praId"`
						PresenceState struct {
							PresenceState string `json:"presenceState"`
						} `json:"presenceState"`
						TrackingAreaList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"trackingAreaList"`
						EcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgiList"`
						NcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgiList"`
						GlobalRanNodeIDList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalRanNodeIdList"`
					} `json:"additionalProp3"`
				} `json:"presenceReportingAreaInformation"`
				ThreeGppPSDataOffStatus struct {
					ThreeGppPSDataOffStatus string `json:"3gppPSDataOffStatus"`
				} `json:"3gppPSDataOffStatus"`
				SponsorIdentity                    string `json:"sponsorIdentity"`
				ApplicationserviceProviderIdentity string `json:"applicationserviceProviderIdentity"`
				ChargingRuleBaseName               string `json:"chargingRuleBaseName"`
				TrafficForwardingWay               struct {
					TrafficForwardingWay string `json:"trafficForwardingWay"`
				} `json:"trafficForwardingWay"`
				QosMonitoringReport []struct {
					UlDelays []int `json:"ulDelays"`
					DlDelays []int `json:"dlDelays"`
					RtDelays []int `json:"rtDelays"`
				} `json:"qosMonitoringReport"`
			} `json:"pDUContainerInformation"`
			NSPAContainerInformation struct {
				Latency    int `json:"latency"`
				Throughput struct {
					GuaranteedThpt int `json:"guaranteedThpt"`
					MaximumThpt    int `json:"maximumThpt"`
				} `json:"throughput"`
				MaximumPacketLossRate            string `json:"maximumPacketLossRate"`
				TheNumberOfPDUSessions           int    `json:"theNumberOfPDUSessions"`
				TheNumberOfRegisteredSubscribers int    `json:"theNumberOfRegisteredSubscribers"`
			} `json:"nSPAContainerInformation"`
			PC5ContainerInformation struct {
				CoverageInfoList []struct {
					CoverageStatus bool      `json:"coverageStatus"`
					ChangeTime     time.Time `json:"changeTime"`
					LocationInfo   []struct {
						EutraLocation struct {
							Tai struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"tai"`
							Ecgi struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								EutraCellID string `json:"eutraCellId"`
							} `json:"ecgi"`
							AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
							UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
							GeographicalInformation  string    `json:"geographicalInformation"`
							GeodeticInformation      string    `json:"geodeticInformation"`
							GlobalNgenbID            struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalNgenbId"`
						} `json:"eutraLocation"`
						NrLocation struct {
							Tai struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"tai"`
							Ncgi struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								NrCellID string `json:"nrCellId"`
							} `json:"ncgi"`
							AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
							UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
							GeographicalInformation  string    `json:"geographicalInformation"`
							GeodeticInformation      string    `json:"geodeticInformation"`
							GlobalGnbID              struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalGnbId"`
						} `json:"nrLocation"`
						N3GaLocation struct {
							N3GppTai struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"n3gppTai"`
							N3IwfID    string `json:"n3IwfId"`
							UeIpv4Addr string `json:"ueIpv4Addr"`
							UeIpv6Addr struct {
								UeIpv6Addr string `json:"ueIpv6Addr"`
							} `json:"ueIpv6Addr"`
							PortNumber int `json:"portNumber"`
						} `json:"n3gaLocation"`
					} `json:"locationInfo"`
				} `json:"coverageInfoList"`
				RadioParameterSetInfoList []struct {
					RadioParameterSetValues []string  `json:"radioParameterSetValues"`
					ChangeTimestamp         time.Time `json:"changeTimestamp"`
				} `json:"radioParameterSetInfoList"`
				TransmitterInfoList []struct {
					ProseSourceL2ID string `json:"proseSourceL2Id"`
				} `json:"transmitterInfoList"`
				TimeOfFirstTransmission time.Time `json:"timeOfFirst Transmission"`
				TimeOfFirstReception    time.Time `json:"timeOfFirst Reception"`
			} `json:"pC5ContainerInformation"`
		} `json:"usedUnitContainer"`
		UPFID                string `json:"uPFID"`
		MultihomedPDUAddress struct {
			PduIPv4Address           string `json:"pduIPv4Address"`
			PduIPv6AddresswithPrefix struct {
				PduIPv6AddresswithPrefix string `json:"pduIPv6AddresswithPrefix"`
			} `json:"pduIPv6AddresswithPrefix"`
			PduAddressprefixlength int  `json:"pduAddressprefixlength"`
			IPv4DynamicAddressFlag bool `json:"iPv4dynamicAddressFlag"`
			IPv6DynamicPrefixFlag  bool `json:"iPv6dynamicPrefixFlag"`
			AddIpv6AddrPrefixes    struct {
				AddIpv6AddrPrefixes string `json:"addIpv6AddrPrefixes"`
			} `json:"addIpv6AddrPrefixes"`
			AddIpv6AddrPrefixList any `json:"addIpv6AddrPrefixList"`
		} `json:"multihomedPDUAddress"`
	} `json:"multipleUnitUsage"`
	Triggers []struct {
		TriggerType     string `json:"triggerType"`
		TriggerCategory struct {
			TriggerCategory string `json:"triggerCategory"`
		} `json:"triggerCategory"`
		TimeLimit        int       `json:"timeLimit"`
		VolumeLimit      int       `json:"volumeLimit"`
		VolumeLimit64    int       `json:"volumeLimit64"`
		EventLimit       int       `json:"eventLimit"`
		MaxNumberOfccc   int       `json:"maxNumberOfccc"`
		TariffTimeChange time.Time `json:"tariffTimeChange"`
	} `json:"triggers"`
	Easid                         string `json:"easid"`
	Ednid                         string `json:"ednid"`
	EASProviderIdentifier         string `json:"eASProviderIdentifier"`
	PDUSessionChargingInformation struct {
		ChargingID                int    `json:"chargingId"`
		SMFchargingID             string `json:"sMFchargingId"`
		HomeProvidedChargingID    int    `json:"homeProvidedChargingId"`
		SMFHomeProvidedChargingID string `json:"sMFHomeProvidedChargingId"`
		UserInformation           struct {
			ServedGPSI          string `json:"servedGPSI"`
			ServedPEI           string `json:"servedPEI"`
			UnauthenticatedFlag bool   `json:"unauthenticatedFlag"`
			RoamerInOut         struct {
				RoamerInOut string `json:"roamerInOut"`
			} `json:"roamerInOut"`
		} `json:"userInformation"`
		UserLocationinfo struct {
			EutraLocation struct {
				Tai struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"tai"`
				Ecgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgi"`
				AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
				UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
				GeographicalInformation  string    `json:"geographicalInformation"`
				GeodeticInformation      string    `json:"geodeticInformation"`
				GlobalNgenbID            struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalNgenbId"`
			} `json:"eutraLocation"`
			NrLocation struct {
				Tai struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"tai"`
				Ncgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"ncgi"`
				AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
				UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
				GeographicalInformation  string    `json:"geographicalInformation"`
				GeodeticInformation      string    `json:"geodeticInformation"`
				GlobalGnbID              struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalGnbId"`
			} `json:"nrLocation"`
			N3GaLocation struct {
				N3GppTai struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"n3gppTai"`
				N3IwfID    string `json:"n3IwfId"`
				UeIpv4Addr string `json:"ueIpv4Addr"`
				UeIpv6Addr struct {
					UeIpv6Addr string `json:"ueIpv6Addr"`
				} `json:"ueIpv6Addr"`
				PortNumber int `json:"portNumber"`
			} `json:"n3gaLocation"`
		} `json:"userLocationinfo"`
		MAPDUNon3GPPUserLocationInfo struct {
			EutraLocation struct {
				Tai struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"tai"`
				Ecgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgi"`
				AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
				UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
				GeographicalInformation  string    `json:"geographicalInformation"`
				GeodeticInformation      string    `json:"geodeticInformation"`
				GlobalNgenbID            struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalNgenbId"`
			} `json:"eutraLocation"`
			NrLocation struct {
				Tai struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"tai"`
				Ncgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"ncgi"`
				AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
				UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
				GeographicalInformation  string    `json:"geographicalInformation"`
				GeodeticInformation      string    `json:"geodeticInformation"`
				GlobalGnbID              struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalGnbId"`
			} `json:"nrLocation"`
			N3GaLocation struct {
				N3GppTai struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"n3gppTai"`
				N3IwfID    string `json:"n3IwfId"`
				UeIpv4Addr string `json:"ueIpv4Addr"`
				UeIpv6Addr struct {
					UeIpv6Addr string `json:"ueIpv6Addr"`
				} `json:"ueIpv6Addr"`
				PortNumber int `json:"portNumber"`
			} `json:"n3gaLocation"`
		} `json:"mAPDUNon3GPPUserLocationInfo"`
		Non3GPPUserLocationTime          time.Time `json:"non3GPPUserLocationTime"`
		MAPDUNon3GPPUserLocationTime     time.Time `json:"mAPDUNon3GPPUserLocationTime"`
		PresenceReportingAreaInformation struct {
			AdditionalProp1 struct {
				PraID         string `json:"praId"`
				PresenceState struct {
					PresenceState string `json:"presenceState"`
				} `json:"presenceState"`
				TrackingAreaList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"trackingAreaList"`
				EcgiList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgiList"`
				NcgiList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"ncgiList"`
				GlobalRanNodeIDList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalRanNodeIdList"`
			} `json:"additionalProp1"`
			AdditionalProp2 struct {
				PraID         string `json:"praId"`
				PresenceState struct {
					PresenceState string `json:"presenceState"`
				} `json:"presenceState"`
				TrackingAreaList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"trackingAreaList"`
				EcgiList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgiList"`
				NcgiList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"ncgiList"`
				GlobalRanNodeIDList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalRanNodeIdList"`
			} `json:"additionalProp2"`
			AdditionalProp3 struct {
				PraID         string `json:"praId"`
				PresenceState struct {
					PresenceState string `json:"presenceState"`
				} `json:"presenceState"`
				TrackingAreaList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					Tac string `json:"tac"`
				} `json:"trackingAreaList"`
				EcgiList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgiList"`
				NcgiList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"ncgiList"`
				GlobalRanNodeIDList []struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					N3IwfID string `json:"n3IwfId"`
					GNbID   struct {
						BitLength int    `json:"bitLength"`
						GNBValue  string `json:"gNBValue"`
					} `json:"gNbId"`
					NgeNbID string `json:"ngeNbId"`
				} `json:"globalRanNodeIdList"`
			} `json:"additionalProp3"`
		} `json:"presenceReportingAreaInformation"`
		UetimeZone            string `json:"uetimeZone"`
		PduSessionInformation struct {
			NetworkSlicingInfo struct {
				SNSSAI struct {
					Sst int    `json:"sst"`
					Sd  string `json:"sd"`
				} `json:"sNSSAI"`
			} `json:"networkSlicingInfo"`
			PduSessionID int `json:"pduSessionID"`
			PduType      struct {
				PduType string `json:"pduType"`
			} `json:"pduType"`
			SscMode struct {
				SscMode string `json:"sscMode"`
			} `json:"sscMode"`
			HPlmnID struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"hPlmnId"`
			ServingNetworkFunctionID struct {
				ServingNetworkFunctionInformation string `json:"servingNetworkFunctionInformation"`
				AMFID                             string `json:"aMFId"`
			} `json:"servingNetworkFunctionID"`
			RatType struct {
				RatType string `json:"ratType"`
			} `json:"ratType"`
			MAPDUNon3GPPRATType struct {
				MAPDUNon3GPPRATType string `json:"mAPDUNon3GPPRATType"`
			} `json:"mAPDUNon3GPPRATType"`
			DnnID            string `json:"dnnId"`
			DnnSelectionMode struct {
				DnnSelectionMode string `json:"dnnSelectionMode"`
			} `json:"dnnSelectionMode"`
			ChargingCharacteristics              string `json:"chargingCharacteristics"`
			ChargingCharacteristicsSelectionMode struct {
				ChargingCharacteristicsSelectionMode string `json:"chargingCharacteristicsSelectionMode"`
			} `json:"chargingCharacteristicsSelectionMode"`
			StartTime               time.Time `json:"startTime"`
			StopTime                time.Time `json:"stopTime"`
			ThreeGppPSDataOffStatus struct {
				ThreeGppPSDataOffStatus string `json:"3gppPSDataOffStatus"`
			} `json:"3gppPSDataOffStatus"`
			SessionStopIndicator bool `json:"sessionStopIndicator"`
			PduAddress           struct {
				PduIPv4Address           string `json:"pduIPv4Address"`
				PduIPv6AddresswithPrefix struct {
					PduIPv6AddresswithPrefix string `json:"pduIPv6AddresswithPrefix"`
				} `json:"pduIPv6AddresswithPrefix"`
				PduAddressprefixlength int  `json:"pduAddressprefixlength"`
				IPv4DynamicAddressFlag bool `json:"iPv4dynamicAddressFlag"`
				IPv6DynamicPrefixFlag  bool `json:"iPv6dynamicPrefixFlag"`
				AddIpv6AddrPrefixes    struct {
					AddIpv6AddrPrefixes string `json:"addIpv6AddrPrefixes"`
				} `json:"addIpv6AddrPrefixes"`
				AddIpv6AddrPrefixList any `json:"addIpv6AddrPrefixList"`
			} `json:"pduAddress"`
			Diagnostics              int `json:"diagnostics"`
			AuthorizedQoSInformation struct {
				FiveQi int `json:"5qi"`
				Arp    struct {
					PriorityLevel int `json:"priorityLevel"`
					PreemptCap    struct {
						PreemptCap string `json:"preemptCap"`
					} `json:"preemptCap"`
					PreemptVuln struct {
						PreemptVuln string `json:"preemptVuln"`
					} `json:"preemptVuln"`
				} `json:"arp"`
				PriorityLevel   int `json:"priorityLevel"`
				AverWindow      int `json:"averWindow"`
				MaxDataBurstVol int `json:"maxDataBurstVol"`
			} `json:"authorizedQoSInformation"`
			SubscribedQoSInformation struct {
				FiveQi int `json:"5qi"`
				Arp    struct {
					PriorityLevel int `json:"priorityLevel"`
					PreemptCap    struct {
						PreemptCap string `json:"preemptCap"`
					} `json:"preemptCap"`
					PreemptVuln struct {
						PreemptVuln string `json:"preemptVuln"`
					} `json:"preemptVuln"`
				} `json:"arp"`
				PriorityLevel int `json:"priorityLevel"`
			} `json:"subscribedQoSInformation"`
			AuthorizedSessionAMBR struct {
				Uplink   string `json:"uplink"`
				Downlink string `json:"downlink"`
			} `json:"authorizedSessionAMBR"`
			SubscribedSessionAMBR struct {
				Uplink   string `json:"uplink"`
				Downlink string `json:"downlink"`
			} `json:"subscribedSessionAMBR"`
			ServingCNPlmnID struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"servingCNPlmnId"`
			MAPDUSessionInformation struct {
			} `json:"mAPDUSessionInformation"`
			EnhancedDiagnostics []struct {
				NgApCause struct {
					Group int `json:"group"`
					Value int `json:"value"`
				} `json:"ngApCause"`
				FiveGMmCause int `json:"5gMmCause"`
				FiveGSmCause int `json:"5gSmCause"`
			} `json:"enhancedDiagnostics"`
			RedundantTransmissionType struct {
				RedundantTransmissionType string `json:"redundantTransmissionType"`
			} `json:"redundantTransmissionType"`
			PDUSessionPairID                int  `json:"pDUSessionPairID"`
			CpCIoTOptimisationIndicator     bool `json:"cpCIoTOptimisationIndicator"`
			FiveGSControlPlaneOnlyIndicator bool `json:"5GSControlPlaneOnlyIndicator"`
			SmallDataRateControlIndicator   bool `json:"smallDataRateControlIndicator"`
			FiveGLANTypeService             struct {
				InternalGroupIdentifier string `json:"internalGroupIdentifier"`
			} `json:"5GLANTypeService"`
		} `json:"pduSessionInformation"`
		UnitCountInactivityTimer   int                        `json:"unitCountInactivityTimer"`
		RANSecondaryRATUsageReport RANSecondaryRATUsageReport `json:"pDUSessionChargingInformation"`
		RoamingQBCInformation      struct {
			MultipleQFIcontainer []struct {
				Triggers []struct {
					TriggerType     string `json:"triggerType"`
					TriggerCategory struct {
						TriggerCategory string `json:"triggerCategory"`
					} `json:"triggerCategory"`
					TimeLimit        int       `json:"timeLimit"`
					VolumeLimit      int       `json:"volumeLimit"`
					VolumeLimit64    int       `json:"volumeLimit64"`
					EventLimit       int       `json:"eventLimit"`
					MaxNumberOfccc   int       `json:"maxNumberOfccc"`
					TariffTimeChange time.Time `json:"tariffTimeChange"`
				} `json:"triggers"`
				TriggerTimestamp        time.Time `json:"triggerTimestamp"`
				Time                    int       `json:"time"`
				TotalVolume             int       `json:"totalVolume"`
				UplinkVolume            int       `json:"uplinkVolume"`
				DownlinkVolume          int       `json:"downlinkVolume"`
				LocalSequenceNumber     int       `json:"localSequenceNumber"`
				QFIContainerInformation struct {
					QFI              int       `json:"qFI"`
					ReportTime       time.Time `json:"reportTime"`
					TimeofFirstUsage time.Time `json:"timeofFirstUsage"`
					TimeofLastUsage  time.Time `json:"timeofLastUsage"`
					QoSInformation   struct {
						QosID   string `json:"qosId"`
						FiveQi  int    `json:"5qi"`
						MaxbrUl string `json:"maxbrUl"`
						MaxbrDl string `json:"maxbrDl"`
						GbrUl   string `json:"gbrUl"`
						GbrDl   string `json:"gbrDl"`
						Arp     struct {
							PriorityLevel int `json:"priorityLevel"`
							PreemptCap    struct {
								PreemptCap string `json:"preemptCap"`
							} `json:"preemptCap"`
							PreemptVuln struct {
								PreemptVuln string `json:"preemptVuln"`
							} `json:"preemptVuln"`
						} `json:"arp"`
						Qnc                  bool   `json:"qnc"`
						PriorityLevel        int    `json:"priorityLevel"`
						AverWindow           int    `json:"averWindow"`
						MaxDataBurstVol      int    `json:"maxDataBurstVol"`
						ReflectiveQos        bool   `json:"reflectiveQos"`
						SharingKeyDl         string `json:"sharingKeyDl"`
						SharingKeyUl         string `json:"sharingKeyUl"`
						MaxPacketLossRateDl  int    `json:"maxPacketLossRateDl"`
						MaxPacketLossRateUl  int    `json:"maxPacketLossRateUl"`
						DefQosFlowIndication bool   `json:"defQosFlowIndication"`
					} `json:"qoSInformation"`
					QoSCharacteristics struct {
						FiveQi       int `json:"5qi"`
						ResourceType struct {
							ResourceType string `json:"resourceType"`
						} `json:"resourceType"`
						PriorityLevel     int    `json:"priorityLevel"`
						PacketDelayBudget int    `json:"packetDelayBudget"`
						PacketErrorRate   string `json:"packetErrorRate"`
						AveragingWindow   int    `json:"averagingWindow"`
						MaxDataBurstVol   int    `json:"maxDataBurstVol"`
					} `json:"qoSCharacteristics"`
					UserLocationInformation struct {
						EutraLocation struct {
							Tai struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"tai"`
							Ecgi struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								EutraCellID string `json:"eutraCellId"`
							} `json:"ecgi"`
							AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
							UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
							GeographicalInformation  string    `json:"geographicalInformation"`
							GeodeticInformation      string    `json:"geodeticInformation"`
							GlobalNgenbID            struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalNgenbId"`
						} `json:"eutraLocation"`
						NrLocation struct {
							Tai struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"tai"`
							Ncgi struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								NrCellID string `json:"nrCellId"`
							} `json:"ncgi"`
							AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
							UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
							GeographicalInformation  string    `json:"geographicalInformation"`
							GeodeticInformation      string    `json:"geodeticInformation"`
							GlobalGnbID              struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalGnbId"`
						} `json:"nrLocation"`
						N3GaLocation struct {
							N3GppTai struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"n3gppTai"`
							N3IwfID    string `json:"n3IwfId"`
							UeIpv4Addr string `json:"ueIpv4Addr"`
							UeIpv6Addr struct {
								UeIpv6Addr string `json:"ueIpv6Addr"`
							} `json:"ueIpv6Addr"`
							PortNumber int `json:"portNumber"`
						} `json:"n3gaLocation"`
					} `json:"userLocationInformation"`
					UetimeZone                       string `json:"uetimeZone"`
					PresenceReportingAreaInformation struct {
						AdditionalProp1 struct {
							PraID         string `json:"praId"`
							PresenceState struct {
								PresenceState string `json:"presenceState"`
							} `json:"presenceState"`
							TrackingAreaList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"trackingAreaList"`
							EcgiList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								EutraCellID string `json:"eutraCellId"`
							} `json:"ecgiList"`
							NcgiList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								NrCellID string `json:"nrCellId"`
							} `json:"ncgiList"`
							GlobalRanNodeIDList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalRanNodeIdList"`
						} `json:"additionalProp1"`
						AdditionalProp2 struct {
							PraID         string `json:"praId"`
							PresenceState struct {
								PresenceState string `json:"presenceState"`
							} `json:"presenceState"`
							TrackingAreaList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"trackingAreaList"`
							EcgiList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								EutraCellID string `json:"eutraCellId"`
							} `json:"ecgiList"`
							NcgiList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								NrCellID string `json:"nrCellId"`
							} `json:"ncgiList"`
							GlobalRanNodeIDList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalRanNodeIdList"`
						} `json:"additionalProp2"`
						AdditionalProp3 struct {
							PraID         string `json:"praId"`
							PresenceState struct {
								PresenceState string `json:"presenceState"`
							} `json:"presenceState"`
							TrackingAreaList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								Tac string `json:"tac"`
							} `json:"trackingAreaList"`
							EcgiList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								EutraCellID string `json:"eutraCellId"`
							} `json:"ecgiList"`
							NcgiList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								NrCellID string `json:"nrCellId"`
							} `json:"ncgiList"`
							GlobalRanNodeIDList []struct {
								PlmnID struct {
									Mcc string `json:"mcc"`
									Mnc string `json:"mnc"`
								} `json:"plmnId"`
								N3IwfID string `json:"n3IwfId"`
								GNbID   struct {
									BitLength int    `json:"bitLength"`
									GNBValue  string `json:"gNBValue"`
								} `json:"gNbId"`
								NgeNbID string `json:"ngeNbId"`
							} `json:"globalRanNodeIdList"`
						} `json:"additionalProp3"`
					} `json:"presenceReportingAreaInformation"`
					RATType struct {
						RATType string `json:"rATType"`
					} `json:"rATType"`
					ServingNetworkFunctionID []struct {
						ServingNetworkFunctionInformation string `json:"servingNetworkFunctionInformation"`
						AMFID                             string `json:"aMFId"`
					} `json:"servingNetworkFunctionID"`
					ThreeGppPSDataOffStatus struct {
						ThreeGppPSDataOffStatus string `json:"3gppPSDataOffStatus"`
					} `json:"3gppPSDataOffStatus"`
					ThreeGppChargingID  int      `json:"3gppChargingId"`
					Diagnostics         int      `json:"diagnostics"`
					EnhancedDiagnostics []string `json:"enhancedDiagnostics"`
				} `json:"qFIContainerInformation"`
			} `json:"multipleQFIcontainer"`
			UPFID                  string `json:"uPFID"`
			RoamingChargingProfile struct {
				Triggers []struct {
					TriggerType     string `json:"triggerType"`
					TriggerCategory struct {
						TriggerCategory string `json:"triggerCategory"`
					} `json:"triggerCategory"`
					TimeLimit        int       `json:"timeLimit"`
					VolumeLimit      int       `json:"volumeLimit"`
					VolumeLimit64    int       `json:"volumeLimit64"`
					EventLimit       int       `json:"eventLimit"`
					MaxNumberOfccc   int       `json:"maxNumberOfccc"`
					TariffTimeChange time.Time `json:"tariffTimeChange"`
				} `json:"triggers"`
				PartialRecordMethod struct {
					PartialRecordMethod string `json:"partialRecordMethod"`
				} `json:"partialRecordMethod"`
			} `json:"roamingChargingProfile"`
		} `json:"roamingQBCInformation"`
		SMSChargingInformation struct {
			OriginatorInfo struct {
				OriginatorSUPI         string `json:"originatorSUPI"`
				OriginatorGPSI         string `json:"originatorGPSI"`
				OriginatorOtherAddress struct {
					SMaddressType struct {
						SMaddressType string `json:"sMaddressType"`
					} `json:"sMaddressType"`
					SMaddressData   string `json:"sMaddressData"`
					SMaddressDomain struct {
						DomainName         string `json:"domainName"`
						ThreeGPPIMSIMCCMNC string `json:"3GPPIMSIMCCMNC"`
					} `json:"sMaddressDomain"`
				} `json:"originatorOtherAddress"`
				OriginatorReceivedAddress struct {
					SMaddressType struct {
						SMaddressType string `json:"sMaddressType"`
					} `json:"sMaddressType"`
					SMaddressData   string `json:"sMaddressData"`
					SMaddressDomain struct {
						DomainName         string `json:"domainName"`
						ThreeGPPIMSIMCCMNC string `json:"3GPPIMSIMCCMNC"`
					} `json:"sMaddressDomain"`
				} `json:"originatorReceivedAddress"`
				OriginatorSCCPAddress string `json:"originatorSCCPAddress"`
				SMOriginatorInterface struct {
					InterfaceID   string `json:"interfaceId"`
					InterfaceText string `json:"interfaceText"`
					InterfacePort string `json:"interfacePort"`
					InterfaceType struct {
						InterfaceType string `json:"interfaceType"`
					} `json:"interfaceType"`
				} `json:"sMOriginatorInterface"`
				SMOriginatorProtocolID string `json:"sMOriginatorProtocolId"`
			} `json:"originatorInfo"`
			RecipientInfo []struct {
				RecipientSUPI         string `json:"recipientSUPI"`
				RecipientGPSI         string `json:"recipientGPSI"`
				RecipientOtherAddress struct {
					SMaddressType struct {
						SMaddressType string `json:"sMaddressType"`
					} `json:"sMaddressType"`
					SMaddressData   string `json:"sMaddressData"`
					SMaddressDomain struct {
						DomainName         string `json:"domainName"`
						ThreeGPPIMSIMCCMNC string `json:"3GPPIMSIMCCMNC"`
					} `json:"sMaddressDomain"`
				} `json:"recipientOtherAddress"`
				RecipientReceivedAddress struct {
					SMaddressType struct {
						SMaddressType string `json:"sMaddressType"`
					} `json:"sMaddressType"`
					SMaddressData   string `json:"sMaddressData"`
					SMaddressDomain struct {
						DomainName         string `json:"domainName"`
						ThreeGPPIMSIMCCMNC string `json:"3GPPIMSIMCCMNC"`
					} `json:"sMaddressDomain"`
				} `json:"recipientReceivedAddress"`
				RecipientSCCPAddress   string `json:"recipientSCCPAddress"`
				SMDestinationInterface struct {
					InterfaceID   string `json:"interfaceId"`
					InterfaceText string `json:"interfaceText"`
					InterfacePort string `json:"interfacePort"`
					InterfaceType struct {
						InterfaceType string `json:"interfaceType"`
					} `json:"interfaceType"`
				} `json:"sMDestinationInterface"`
				SMrecipientProtocolID string `json:"sMrecipientProtocolId"`
			} `json:"recipientInfo"`
			UserEquipmentInfo string `json:"userEquipmentInfo"`
			RoamerInOut       struct {
				RoamerInOut string `json:"roamerInOut"`
			} `json:"roamerInOut"`
			UserLocationinfo struct {
				EutraLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ecgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalNgenbID            struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalNgenbId"`
				} `json:"eutraLocation"`
				NrLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ncgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalGnbID              struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalGnbId"`
				} `json:"nrLocation"`
				N3GaLocation struct {
					N3GppTai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"n3gppTai"`
					N3IwfID    string `json:"n3IwfId"`
					UeIpv4Addr string `json:"ueIpv4Addr"`
					UeIpv6Addr struct {
						UeIpv6Addr string `json:"ueIpv6Addr"`
					} `json:"ueIpv6Addr"`
					PortNumber int `json:"portNumber"`
				} `json:"n3gaLocation"`
			} `json:"userLocationinfo"`
			UetimeZone string `json:"uetimeZone"`
			RATType    struct {
				RATType string `json:"rATType"`
			} `json:"rATType"`
			SMSCAddress        string `json:"sMSCAddress"`
			SMDataCodingScheme int    `json:"sMDataCodingScheme"`
			SMMessageType      struct {
				SMMessageType string `json:"sMMessageType"`
			} `json:"sMMessageType"`
			SMReplyPathRequested struct {
				SMReplyPathRequested string `json:"sMReplyPathRequested"`
			} `json:"sMReplyPathRequested"`
			SMUserDataHeader     string    `json:"sMUserDataHeader"`
			SMStatus             string    `json:"sMStatus"`
			SMDischargeTime      time.Time `json:"sMDischargeTime"`
			NumberofMessagesSent int       `json:"numberofMessagesSent"`
			SMServiceType        struct {
				SMServiceType string `json:"sMServiceType"`
			} `json:"sMServiceType"`
			SMSequenceNumber int       `json:"sMSequenceNumber"`
			SMSresult        int       `json:"sMSresult"`
			SubmissionTime   time.Time `json:"submissionTime"`
			SMPriority       struct {
				SMPriority string `json:"sMPriority"`
			} `json:"sMPriority"`
			MessageReference string `json:"messageReference"`
			MessageSize      int    `json:"messageSize"`
			MessageClass     struct {
				ClassIdentifier struct {
					ClassIdentifier string `json:"classIdentifier"`
				} `json:"classIdentifier"`
				TokenText string `json:"tokenText"`
			} `json:"messageClass"`
			DeliveryReportRequested struct {
				DeliveryReportRequested string `json:"deliveryReportRequested"`
			} `json:"deliveryReportRequested"`
		} `json:"sMSChargingInformation"`
		NEFChargingInformation struct {
			ExternalIndividualIdentifier string `json:"externalIndividualIdentifier"`
			GroupIdentifier              string `json:"groupIdentifier"`
			APIDirection                 struct {
				APIDirection string `json:"aPIDirection"`
			} `json:"aPIDirection"`
			APITargetNetworkFunction string `json:"aPITargetNetworkFunction"`
			APIResultCode            int    `json:"aPIResultCode"`
			APIName                  string `json:"aPIName"`
			APIReference             string `json:"aPIReference"`
			APIContent               string `json:"aPIContent"`
		} `json:"nEFChargingInformation"`
		RegistrationChargingInformation struct {
			RegistrationMessagetype struct {
				RegistrationMessagetype string `json:"registrationMessagetype"`
			} `json:"registrationMessagetype"`
			UserInformation struct {
				ServedGPSI          string `json:"servedGPSI"`
				ServedPEI           string `json:"servedPEI"`
				UnauthenticatedFlag bool   `json:"unauthenticatedFlag"`
				RoamerInOut         struct {
					RoamerInOut string `json:"roamerInOut"`
				} `json:"roamerInOut"`
			} `json:"userInformation"`
			UserLocationinfo struct {
				EutraLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ecgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalNgenbID            struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalNgenbId"`
				} `json:"eutraLocation"`
				NrLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ncgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalGnbID              struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalGnbId"`
				} `json:"nrLocation"`
				N3GaLocation struct {
					N3GppTai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"n3gppTai"`
					N3IwfID    string `json:"n3IwfId"`
					UeIpv4Addr string `json:"ueIpv4Addr"`
					UeIpv6Addr struct {
						UeIpv6Addr string `json:"ueIpv6Addr"`
					} `json:"ueIpv6Addr"`
					PortNumber int `json:"portNumber"`
				} `json:"n3gaLocation"`
			} `json:"userLocationinfo"`
			PSCellInformation struct {
				Nrcgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"nrcgi"`
				Ecgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgi"`
			} `json:"pSCellInformation"`
			UetimeZone string `json:"uetimeZone"`
			RATType    struct {
				RATType string `json:"rATType"`
			} `json:"rATType"`
			FiveGMMCapability  any `json:"5GMMCapability"`
			MICOModeIndication struct {
				MICOModeIndication string `json:"mICOModeIndication"`
			} `json:"mICOModeIndication"`
			SmsIndication struct {
				SmsIndication string `json:"smsIndication"`
			} `json:"smsIndication"`
			TaiList []struct {
				PlmnID struct {
					Mcc string `json:"mcc"`
					Mnc string `json:"mnc"`
				} `json:"plmnId"`
				Tac string `json:"tac"`
			} `json:"taiList"`
			ServiceAreaRestriction []struct {
				RestrictionType struct {
					RestrictionType string `json:"restrictionType"`
				} `json:"restrictionType"`
				Areas []struct {
					Tacs      []string `json:"tacs"`
					AreaCodes string   `json:"areaCodes"`
				} `json:"areas"`
				MaxNumOfTAs int `json:"maxNumOfTAs"`
			} `json:"serviceAreaRestriction"`
			RequestedNSSAI []struct {
				Sst int    `json:"sst"`
				Sd  string `json:"sd"`
			} `json:"requestedNSSAI"`
			AllowedNSSAI []struct {
				Sst int    `json:"sst"`
				Sd  string `json:"sd"`
			} `json:"allowedNSSAI"`
			RejectedNSSAI []struct {
				Sst int    `json:"sst"`
				Sd  string `json:"sd"`
			} `json:"rejectedNSSAI"`
			NSSAIMapList []struct {
				ServingSnssai struct {
					Sst int    `json:"sst"`
					Sd  string `json:"sd"`
				} `json:"servingSnssai"`
				HomeSnssai struct {
					Sst int    `json:"sst"`
					Sd  string `json:"sd"`
				} `json:"homeSnssai"`
			} `json:"nSSAIMapList"`
			AmfUeNgapID int `json:"amfUeNgapId"`
			RanUeNgapID int `json:"ranUeNgapId"`
			RanNodeID   struct {
				PlmnID struct {
					Mcc string `json:"mcc"`
					Mnc string `json:"mnc"`
				} `json:"plmnId"`
				N3IwfID string `json:"n3IwfId"`
				GNbID   struct {
					BitLength int    `json:"bitLength"`
					GNBValue  string `json:"gNBValue"`
				} `json:"gNbId"`
				NgeNbID string `json:"ngeNbId"`
			} `json:"ranNodeId"`
		} `json:"registrationChargingInformation"`
		N2ConnectionChargingInformation struct {
			N2ConnectionMessageType int `json:"n2ConnectionMessageType"`
			UserInformation         struct {
				ServedGPSI          string `json:"servedGPSI"`
				ServedPEI           string `json:"servedPEI"`
				UnauthenticatedFlag bool   `json:"unauthenticatedFlag"`
				RoamerInOut         struct {
					RoamerInOut string `json:"roamerInOut"`
				} `json:"roamerInOut"`
			} `json:"userInformation"`
			UserLocationinfo struct {
				EutraLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ecgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalNgenbID            struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalNgenbId"`
				} `json:"eutraLocation"`
				NrLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ncgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalGnbID              struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalGnbId"`
				} `json:"nrLocation"`
				N3GaLocation struct {
					N3GppTai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"n3gppTai"`
					N3IwfID    string `json:"n3IwfId"`
					UeIpv4Addr string `json:"ueIpv4Addr"`
					UeIpv6Addr struct {
						UeIpv6Addr string `json:"ueIpv6Addr"`
					} `json:"ueIpv6Addr"`
					PortNumber int `json:"portNumber"`
				} `json:"n3gaLocation"`
			} `json:"userLocationinfo"`
			PSCellInformation struct {
				Nrcgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"nrcgi"`
				Ecgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgi"`
			} `json:"pSCellInformation"`
			UetimeZone string `json:"uetimeZone"`
			RATType    struct {
				RATType string `json:"rATType"`
			} `json:"rATType"`
			AmfUeNgapID int `json:"amfUeNgapId"`
			RanUeNgapID int `json:"ranUeNgapId"`
			RanNodeID   struct {
				PlmnID struct {
					Mcc string `json:"mcc"`
					Mnc string `json:"mnc"`
				} `json:"plmnId"`
				N3IwfID string `json:"n3IwfId"`
				GNbID   struct {
					BitLength int    `json:"bitLength"`
					GNBValue  string `json:"gNBValue"`
				} `json:"gNbId"`
				NgeNbID string `json:"ngeNbId"`
			} `json:"ranNodeId"`
			RestrictedRatList any `json:"restrictedRatList"`
			ForbiddenAreaList []struct {
				Tacs      []string `json:"tacs"`
				AreaCodes string   `json:"areaCodes"`
			} `json:"forbiddenAreaList"`
			ServiceAreaRestriction []struct {
				RestrictionType struct {
					RestrictionType string `json:"restrictionType"`
				} `json:"restrictionType"`
				Areas []struct {
					Tacs      []string `json:"tacs"`
					AreaCodes string   `json:"areaCodes"`
				} `json:"areas"`
				MaxNumOfTAs int `json:"maxNumOfTAs"`
			} `json:"serviceAreaRestriction"`
			RestrictedCnList any `json:"restrictedCnList"`
			AllowedNSSAI     []struct {
				Sst int    `json:"sst"`
				Sd  string `json:"sd"`
			} `json:"allowedNSSAI"`
			RrcEstCause string `json:"rrcEstCause"`
		} `json:"n2ConnectionChargingInformation"`
		LocationReportingChargingInformation struct {
			LocationReportingMessageType int `json:"locationReportingMessageType"`
			UserInformation              struct {
				ServedGPSI          string `json:"servedGPSI"`
				ServedPEI           string `json:"servedPEI"`
				UnauthenticatedFlag bool   `json:"unauthenticatedFlag"`
				RoamerInOut         struct {
					RoamerInOut string `json:"roamerInOut"`
				} `json:"roamerInOut"`
			} `json:"userInformation"`
			UserLocationinfo struct {
				EutraLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ecgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalNgenbID            struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalNgenbId"`
				} `json:"eutraLocation"`
				NrLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ncgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalGnbID              struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalGnbId"`
				} `json:"nrLocation"`
				N3GaLocation struct {
					N3GppTai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"n3gppTai"`
					N3IwfID    string `json:"n3IwfId"`
					UeIpv4Addr string `json:"ueIpv4Addr"`
					UeIpv6Addr struct {
						UeIpv6Addr string `json:"ueIpv6Addr"`
					} `json:"ueIpv6Addr"`
					PortNumber int `json:"portNumber"`
				} `json:"n3gaLocation"`
			} `json:"userLocationinfo"`
			PSCellInformation struct {
				Nrcgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					NrCellID string `json:"nrCellId"`
				} `json:"nrcgi"`
				Ecgi struct {
					PlmnID struct {
						Mcc string `json:"mcc"`
						Mnc string `json:"mnc"`
					} `json:"plmnId"`
					EutraCellID string `json:"eutraCellId"`
				} `json:"ecgi"`
			} `json:"pSCellInformation"`
			UetimeZone string `json:"uetimeZone"`
			RATType    struct {
				RATType string `json:"rATType"`
			} `json:"rATType"`
			PresenceReportingAreaInformation struct {
				AdditionalProp1 struct {
					PraID         string `json:"praId"`
					PresenceState struct {
						PresenceState string `json:"presenceState"`
					} `json:"presenceState"`
					TrackingAreaList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"trackingAreaList"`
					EcgiList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgiList"`
					NcgiList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgiList"`
					GlobalRanNodeIDList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalRanNodeIdList"`
				} `json:"additionalProp1"`
				AdditionalProp2 struct {
					PraID         string `json:"praId"`
					PresenceState struct {
						PresenceState string `json:"presenceState"`
					} `json:"presenceState"`
					TrackingAreaList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"trackingAreaList"`
					EcgiList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgiList"`
					NcgiList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgiList"`
					GlobalRanNodeIDList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalRanNodeIdList"`
				} `json:"additionalProp2"`
				AdditionalProp3 struct {
					PraID         string `json:"praId"`
					PresenceState struct {
						PresenceState string `json:"presenceState"`
					} `json:"presenceState"`
					TrackingAreaList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"trackingAreaList"`
					EcgiList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgiList"`
					NcgiList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgiList"`
					GlobalRanNodeIDList []struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalRanNodeIdList"`
				} `json:"additionalProp3"`
			} `json:"presenceReportingAreaInformation"`
		} `json:"locationReportingChargingInformation"`
		NSPAChargingInformation struct {
			SingleNSSAI struct {
				Sst int    `json:"sst"`
				Sd  string `json:"sd"`
			} `json:"singleNSSAI"`
		} `json:"nSPAChargingInformation"`
		NSMChargingInformation struct {
			ManagementOperation struct {
				ManagementOperation string `json:"managementOperation"`
			} `json:"managementOperation"`
			IDNetworkSliceInstance                  string `json:"idNetworkSliceInstance"`
			ListOfserviceProfileChargingInformation []struct {
				ServiceProfileIdentifier string `json:"serviceProfileIdentifier"`
				SNSSAIList               []struct {
					Sst int    `json:"sst"`
					Sd  string `json:"sd"`
				} `json:"sNSSAIList"`
				Latency        int    `json:"latency"`
				Availability   int    `json:"availability"`
				Jitter         int    `json:"jitter"`
				Reliability    string `json:"reliability"`
				MaxNumberofUEs int    `json:"maxNumberofUEs"`
				CoverageArea   string `json:"coverageArea"`
				DLThptPerSlice struct {
					GuaranteedThpt int `json:"guaranteedThpt"`
					MaximumThpt    int `json:"maximumThpt"`
				} `json:"dLThptPerSlice"`
				DLThptPerUE struct {
					GuaranteedThpt int `json:"guaranteedThpt"`
					MaximumThpt    int `json:"maximumThpt"`
				} `json:"dLThptPerUE"`
				ULThptPerSlice struct {
					GuaranteedThpt int `json:"guaranteedThpt"`
					MaximumThpt    int `json:"maximumThpt"`
				} `json:"uLThptPerSlice"`
				ULThptPerUE struct {
					GuaranteedThpt int `json:"guaranteedThpt"`
					MaximumThpt    int `json:"maximumThpt"`
				} `json:"uLThptPerUE"`
				MaxNumberofPDUsessions    int    `json:"maxNumberofPDUsessions"`
				KPIMonitoringList         string `json:"kPIMonitoringList"`
				SupportedAccessTechnology int    `json:"supportedAccessTechnology"`
				AddServiceProfileInfo     string `json:"addServiceProfileInfo"`
			} `json:"listOfserviceProfileChargingInformation"`
			ManagementOperationStatus struct {
				ManagementOperationStatus string `json:"managementOperationStatus"`
			} `json:"managementOperationStatus"`
		} `json:"nSMChargingInformation"`
		MMTelChargingInformation struct {
			SupplementaryServices []struct {
				SupplementaryServiceType struct {
					SupplementaryServiceType string `json:"supplementaryServiceType"`
				} `json:"supplementaryServiceType"`
				SupplementaryServiceMode struct {
					SupplementaryServiceMode string `json:"supplementaryServiceMode"`
				} `json:"supplementaryServiceMode"`
				NumberOfDiversions     int    `json:"numberOfDiversions"`
				AssociatedPartyAddress string `json:"associatedPartyAddress"`
				ConferenceID           string `json:"conferenceId"`
				ParticipantActionType  struct {
					ParticipantActionType string `json:"participantActionType"`
				} `json:"participantActionType"`
				ChangeTime           time.Time `json:"changeTime"`
				NumberOfParticipants int       `json:"numberOfParticipants"`
				CUGInformation       string    `json:"cUGInformation"`
			} `json:"supplementaryServices"`
		} `json:"mMTelChargingInformation"`
		IMSChargingInformation struct {
			EventType struct {
				SIPMethod     string `json:"sIPMethod"`
				EventHeader   string `json:"eventHeader"`
				ExpiresHeader int    `json:"expiresHeader"`
			} `json:"eventType"`
			IMSNodeFunctionality struct {
				IMSNodeFunctionality string `json:"iMSNodeFunctionality"`
			} `json:"iMSNodeFunctionality"`
			RoleOfNode struct {
				RoleOfNode string `json:"roleOfNode"`
			} `json:"roleOfNode"`
			UserInformation struct {
				ServedGPSI          string `json:"servedGPSI"`
				ServedPEI           string `json:"servedPEI"`
				UnauthenticatedFlag bool   `json:"unauthenticatedFlag"`
				RoamerInOut         struct {
					RoamerInOut string `json:"roamerInOut"`
				} `json:"roamerInOut"`
			} `json:"userInformation"`
			UserLocationInfo struct {
				EutraLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ecgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						EutraCellID string `json:"eutraCellId"`
					} `json:"ecgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalNgenbID            struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalNgenbId"`
				} `json:"eutraLocation"`
				NrLocation struct {
					Tai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"tai"`
					Ncgi struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						NrCellID string `json:"nrCellId"`
					} `json:"ncgi"`
					AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
					UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
					GeographicalInformation  string    `json:"geographicalInformation"`
					GeodeticInformation      string    `json:"geodeticInformation"`
					GlobalGnbID              struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						N3IwfID string `json:"n3IwfId"`
						GNbID   struct {
							BitLength int    `json:"bitLength"`
							GNBValue  string `json:"gNBValue"`
						} `json:"gNbId"`
						NgeNbID string `json:"ngeNbId"`
					} `json:"globalGnbId"`
				} `json:"nrLocation"`
				N3GaLocation struct {
					N3GppTai struct {
						PlmnID struct {
							Mcc string `json:"mcc"`
							Mnc string `json:"mnc"`
						} `json:"plmnId"`
						Tac string `json:"tac"`
					} `json:"n3gppTai"`
					N3IwfID    string `json:"n3IwfId"`
					UeIpv4Addr string `json:"ueIpv4Addr"`
					UeIpv6Addr struct {
						UeIpv6Addr string `json:"ueIpv6Addr"`
					} `json:"ueIpv6Addr"`
					PortNumber int `json:"portNumber"`
				} `json:"n3gaLocation"`
			} `json:"userLocationInfo"`
			UeTimeZone              string `json:"ueTimeZone"`
			ThreeGppPSDataOffStatus struct {
				ThreeGppPSDataOffStatus string `json:"3gppPSDataOffStatus"`
			} `json:"3gppPSDataOffStatus"`
			IsupCause struct {
				ISUPCauseLocation    int    `json:"iSUPCauseLocation"`
				ISUPCauseValue       int    `json:"iSUPCauseValue"`
				ISUPCauseDiagnostics string `json:"iSUPCauseDiagnostics"`
			} `json:"isupCause"`
			ControlPlaneAddress struct {
				Ipv4Addr string `json:"ipv4Addr"`
				Ipv6Addr struct {
					Ipv6Addr string `json:"ipv6Addr"`
				} `json:"ipv6Addr"`
				E164 string `json:"e164"`
			} `json:"controlPlaneAddress"`
			VlrNumber         string `json:"vlrNumber"`
			MscAddress        string `json:"mscAddress"`
			UserSessionID     string `json:"userSessionID"`
			OutgoingSessionID string `json:"outgoingSessionID"`
			SessionPriority   struct {
				SessionPriority string `json:"sessionPriority"`
			} `json:"sessionPriority"`
			CallingPartyAddresses               []string `json:"callingPartyAddresses"`
			CalledPartyAddress                  string   `json:"calledPartyAddress"`
			NumberPortabilityRoutinginformation string   `json:"numberPortabilityRoutinginformation"`
			CarrierSelectRoutingInformation     string   `json:"carrierSelectRoutingInformation"`
			AlternateChargedPartyAddress        string   `json:"alternateChargedPartyAddress"`
			RequestedPartyAddress               []string `json:"requestedPartyAddress"`
			CalledAssertedIdentities            []string `json:"calledAssertedIdentities"`
			CalledIdentityChanges               []struct {
				CalledIdentity string    `json:"calledIdentity"`
				ChangeTime     time.Time `json:"changeTime"`
			} `json:"calledIdentityChanges"`
			AssociatedURI                []string  `json:"associatedURI"`
			TimeStamps                   time.Time `json:"timeStamps"`
			ApplicationServerInformation []string  `json:"applicationServerInformation"`
			InterOperatorIdentifier      []struct {
				OriginatingIOI string `json:"originatingIOI"`
				TerminatingIOI string `json:"terminatingIOI"`
			} `json:"interOperatorIdentifier"`
			ImsChargingIdentifier     string   `json:"imsChargingIdentifier"`
			RelatedICID               string   `json:"relatedICID"`
			RelatedICIDGenerationNode string   `json:"relatedICIDGenerationNode"`
			TransitIOIList            []string `json:"transitIOIList"`
			EarlyMediaDescription     []struct {
				SDPTimeStamps struct {
					SDPOfferTimestamp  time.Time `json:"sDPOfferTimestamp"`
					SDPAnswerTimestamp time.Time `json:"sDPAnswerTimestamp"`
				} `json:"sDPTimeStamps"`
				SDPMediaComponent []struct {
					SDPMediaName                 string   `json:"sDPMediaName"`
					SDPMediaDescription          []string `json:"SDPMediaDescription"`
					LocalGWInsertedIndication    bool     `json:"localGWInsertedIndication"`
					IPRealmDefaultIndication     bool     `json:"ipRealmDefaultIndication"`
					TranscoderInsertedIndication bool     `json:"transcoderInsertedIndication"`
					MediaInitiatorFlag           struct {
						MediaInitiatorFlag string `json:"mediaInitiatorFlag"`
					} `json:"mediaInitiatorFlag"`
					MediaInitiatorParty                  string `json:"mediaInitiatorParty"`
					ThreeGPPChargingID                   string `json:"threeGPPChargingId"`
					AccessNetworkChargingIdentifierValue string `json:"accessNetworkChargingIdentifierValue"`
					SDPType                              struct {
						SDPType string `json:"sDPType"`
					} `json:"sDPType"`
				} `json:"sDPMediaComponent"`
				SDPSessionDescription []string `json:"sDPSessionDescription"`
			} `json:"earlyMediaDescription"`
			SdpSessionDescription []string `json:"sdpSessionDescription"`
			SdpMediaComponent     []struct {
				SDPMediaName                 string   `json:"sDPMediaName"`
				SDPMediaDescription          []string `json:"SDPMediaDescription"`
				LocalGWInsertedIndication    bool     `json:"localGWInsertedIndication"`
				IPRealmDefaultIndication     bool     `json:"ipRealmDefaultIndication"`
				TranscoderInsertedIndication bool     `json:"transcoderInsertedIndication"`
				MediaInitiatorFlag           struct {
					MediaInitiatorFlag string `json:"mediaInitiatorFlag"`
				} `json:"mediaInitiatorFlag"`
				MediaInitiatorParty                  string `json:"mediaInitiatorParty"`
				ThreeGPPChargingID                   string `json:"threeGPPChargingId"`
				AccessNetworkChargingIdentifierValue string `json:"accessNetworkChargingIdentifierValue"`
				SDPType                              struct {
					SDPType string `json:"sDPType"`
				} `json:"sDPType"`
			} `json:"sdpMediaComponent"`
			ServedPartyIPAddress struct {
				Ipv4Addr string `json:"ipv4Addr"`
				Ipv6Addr struct {
					Ipv6Addr string `json:"ipv6Addr"`
				} `json:"ipv6Addr"`
				E164 string `json:"e164"`
			} `json:"servedPartyIPAddress"`
			ServerCapabilities struct {
				MandatoryCapability []int    `json:"mandatoryCapability"`
				OptionalCapability  []int    `json:"optionalCapability"`
				ServerName          []string `json:"serverName"`
			} `json:"serverCapabilities"`
			TrunkGroupID struct {
				IncomingTrunkGroupID string `json:"incomingTrunkGroupID"`
				OutgoingTrunkGroupID string `json:"outgoingTrunkGroupID"`
			} `json:"trunkGroupID"`
			BearerService string `json:"bearerService"`
			ImsServiceID  string `json:"imsServiceId"`
			MessageBodies []struct {
				ContentType        string `json:"contentType"`
				ContentLength      int    `json:"contentLength"`
				ContentDisposition string `json:"contentDisposition"`
				Originator         struct {
					Originator string `json:"originator"`
				} `json:"originator"`
			} `json:"messageBodies"`
			AccessNetworkInformation           []string `json:"accessNetworkInformation"`
			AdditionalAccessNetworkInformation string   `json:"additionalAccessNetworkInformation"`
			CellularNetworkInformation         string   `json:"cellularNetworkInformation"`
			AccessTransferInformation          []struct {
				AccessTransferType struct {
					AccessTransferType string `json:"accessTransferType"`
				} `json:"accessTransferType"`
				AccessNetworkInformation   []string `json:"accessNetworkInformation"`
				CellularNetworkInformation string   `json:"cellularNetworkInformation"`
				InterUETransfer            struct {
					InterUETransfer string `json:"interUETransfer"`
				} `json:"interUETransfer"`
				UserEquipmentInfo                string `json:"userEquipmentInfo"`
				InstanceID                       string `json:"instanceId"`
				RelatedIMSChargingIdentifier     string `json:"relatedIMSChargingIdentifier"`
				RelatedIMSChargingIdentifierNode struct {
					Ipv4Addr string `json:"ipv4Addr"`
					Ipv6Addr struct {
						Ipv6Addr string `json:"ipv6Addr"`
					} `json:"ipv6Addr"`
					E164 string `json:"e164"`
				} `json:"relatedIMSChargingIdentifierNode"`
				ChangeTime time.Time `json:"changeTime"`
			} `json:"accessTransferInformation"`
			AccessNetworkInfoChange []struct {
				AccessNetworkInformation   []string  `json:"accessNetworkInformation"`
				CellularNetworkInformation string    `json:"cellularNetworkInformation"`
				ChangeTime                 time.Time `json:"changeTime"`
			} `json:"accessNetworkInfoChange"`
			ImsCommunicationServiceID    string   `json:"imsCommunicationServiceID"`
			ImsApplicationReferenceID    string   `json:"imsApplicationReferenceID"`
			CauseCode                    int      `json:"causeCode"`
			ReasonHeader                 []string `json:"reasonHeader"`
			InitialIMSChargingIdentifier string   `json:"initialIMSChargingIdentifier"`
			NniInformation               []struct {
				SessionDirection struct {
					SessionDirection string `json:"sessionDirection"`
				} `json:"sessionDirection"`
				NNIType struct {
					NNIType string `json:"nNIType"`
				} `json:"nNIType"`
				RelationshipMode struct {
					RelationshipMode string `json:"relationshipMode"`
				} `json:"relationshipMode"`
				NeighbourNodeAddress struct {
					Ipv4Addr string `json:"ipv4Addr"`
					Ipv6Addr struct {
						Ipv6Addr string `json:"ipv6Addr"`
					} `json:"ipv6Addr"`
					E164 string `json:"e164"`
				} `json:"neighbourNodeAddress"`
			} `json:"nniInformation"`
			FromAddress                 string `json:"fromAddress"`
			ImsEmergencyIndication      bool   `json:"imsEmergencyIndication"`
			ImsVisitedNetworkIdentifier string `json:"imsVisitedNetworkIdentifier"`
			SipRouteHeaderReceived      string `json:"sipRouteHeaderReceived"`
			SipRouteHeaderTransmitted   string `json:"sipRouteHeaderTransmitted"`
			TadIdentifier               struct {
				TadIdentifier string `json:"tadIdentifier"`
			} `json:"tadIdentifier"`
			FeIdentifierList string `json:"feIdentifierList"`
		} `json:"iMSChargingInformation"`
		EdgeInfrastructureUsageChargingInformation struct {
			MeanVirtualCPUUsage    int       `json:"meanVirtualCPUUsage"`
			MeanVirtualMemoryUsage int       `json:"meanVirtualMemoryUsage"`
			MeanVirtualDiskUsage   int       `json:"meanVirtualDiskUsage"`
			DurationStartTime      time.Time `json:"durationStartTime"`
			DurationEndTime        time.Time `json:"durationEndTime"`
		} `json:"edgeInfrastructureUsageChargingInformation'"`
		EASDeploymentChargingInformation struct {
			LCMStartTime time.Time `json:"lCMStartTime"`
			LCMEndTime   time.Time `json:"lCMEndTime"`
		} `json:"eASDeploymentChargingInformation"`
		DirectEdgeEnablingServiceChargingInformation struct {
			ExternalIndividualIdentifier string `json:"externalIndividualIdentifier"`
			GroupIdentifier              string `json:"groupIdentifier"`
			APIDirection                 struct {
				APIDirection string `json:"aPIDirection"`
			} `json:"aPIDirection"`
			APITargetNetworkFunction string `json:"aPITargetNetworkFunction"`
			APIResultCode            int    `json:"aPIResultCode"`
			APIName                  string `json:"aPIName"`
			APIReference             string `json:"aPIReference"`
			APIContent               string `json:"aPIContent"`
		} `json:"directEdgeEnablingServiceChargingInformation"`
		ExposedEdgeEnablingServiceChargingInformation struct {
			ExternalIndividualIdentifier string `json:"externalIndividualIdentifier"`
			GroupIdentifier              string `json:"groupIdentifier"`
			APIDirection                 struct {
				APIDirection string `json:"aPIDirection"`
			} `json:"aPIDirection"`
			APITargetNetworkFunction string `json:"aPITargetNetworkFunction"`
			APIResultCode            int    `json:"aPIResultCode"`
			APIName                  string `json:"aPIName"`
			APIReference             string `json:"aPIReference"`
			APIContent               string `json:"aPIContent"`
		} `json:"exposedEdgeEnablingServiceChargingInformation"`
		ProSeChargingInformation struct {
			AnnouncingPlmnID struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"announcingPlmnID"`
			AnnouncingUeHplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"announcingUeHplmnIdentifier"`
			AnnouncingUeVplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"announcingUeVplmnIdentifier"`
			MonitoringUeHplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"monitoringUeHplmnIdentifier"`
			MonitoringUeVplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"monitoringUeVplmnIdentifier"`
			DiscovererUeHplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"discovererUeHplmnIdentifier"`
			DiscovererUeVplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"discovererUeVplmnIdentifier"`
			DiscovereeUeHplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"discovereeUeHplmnIdentifier"`
			DiscovereeUeVplmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"discovereeUeVplmnIdentifier"`
			MonitoredPlmnIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"monitoredPlmnIdentifier"`
			ProseApplicationID          string   `json:"proseApplicationID"`
			ApplicationID               string   `json:"ApplicationId"`
			ApplicationSpecificDataList []string `json:"applicationSpecificDataList"`
			ProseFunctionality          struct {
				ProseFunctionality string `json:"proseFunctionality"`
			} `json:"proseFunctionality"`
			ProseEventType struct {
				ProseEventType string `json:"proseEventType"`
			} `json:"proseEventType"`
			DirectDiscoveryModel struct {
				DirectDiscoveryModel string `json:"directDiscoveryModel"`
			} `json:"directDiscoveryModel"`
			ValidityPeriod int `json:"validityPeriod"`
			RoleOfUE       struct {
				RoleOfUE string `json:"roleOfUE"`
			} `json:"roleOfUE"`
			ProseRequestTimestamp   time.Time `json:"proseRequestTimestamp"`
			PC3ProtocolCause        int       `json:"pC3ProtocolCause"`
			MonitoringUEIdentifier  string    `json:"monitoringUEIdentifier"`
			RequestedPLMNIdentifier struct {
				Mcc string `json:"mcc"`
				Mnc string `json:"mnc"`
			} `json:"requestedPLMNIdentifier"`
			TimeWindow int `json:"timeWindow"`
			RangeClass struct {
				RangeClass string `json:"rangeClass"`
			} `json:"rangeClass"`
			ProximityAlertIndication       bool      `json:"proximityAlertIndication"`
			ProximityAlertTimestamp        time.Time `json:"proximityAlertTimestamp"`
			ProximityCancellationTimestamp time.Time `json:"proximityCancellationTimestamp"`
			ProseUEToNetworkRelayUEID      string    `json:"proseUEToNetworkRelayUEID"`
			ProseDestinationLayer2ID       string    `json:"proseDestinationLayer2ID"`
			PFIContainerInformation        []struct {
				PFI              string    `json:"pFI"`
				ReportTime       time.Time `json:"reportTime"`
				TimeofFirstUsage time.Time `json:"timeofFirstUsage"`
				TimeofLastUsage  time.Time `json:"timeofLastUsage"`
				QoSInformation   struct {
					QosID   string `json:"qosId"`
					FiveQi  int    `json:"5qi"`
					MaxbrUl string `json:"maxbrUl"`
					MaxbrDl string `json:"maxbrDl"`
					GbrUl   string `json:"gbrUl"`
					GbrDl   string `json:"gbrDl"`
					Arp     struct {
						PriorityLevel int `json:"priorityLevel"`
						PreemptCap    struct {
							PreemptCap string `json:"preemptCap"`
						} `json:"preemptCap"`
						PreemptVuln struct {
							PreemptVuln string `json:"preemptVuln"`
						} `json:"preemptVuln"`
					} `json:"arp"`
					Qnc                  bool   `json:"qnc"`
					PriorityLevel        int    `json:"priorityLevel"`
					AverWindow           int    `json:"averWindow"`
					MaxDataBurstVol      int    `json:"maxDataBurstVol"`
					ReflectiveQos        bool   `json:"reflectiveQos"`
					SharingKeyDl         string `json:"sharingKeyDl"`
					SharingKeyUl         string `json:"sharingKeyUl"`
					MaxPacketLossRateDl  int    `json:"maxPacketLossRateDl"`
					MaxPacketLossRateUl  int    `json:"maxPacketLossRateUl"`
					DefQosFlowIndication bool   `json:"defQosFlowIndication"`
				} `json:"qoSInformation"`
				QoSCharacteristics struct {
					FiveQi       int `json:"5qi"`
					ResourceType struct {
						ResourceType string `json:"resourceType"`
					} `json:"resourceType"`
					PriorityLevel     int    `json:"priorityLevel"`
					PacketDelayBudget int    `json:"packetDelayBudget"`
					PacketErrorRate   string `json:"packetErrorRate"`
					AveragingWindow   int    `json:"averagingWindow"`
					MaxDataBurstVol   int    `json:"maxDataBurstVol"`
				} `json:"qoSCharacteristics"`
				UserLocationInformation struct {
					EutraLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ecgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalNgenbID            struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalNgenbId"`
					} `json:"eutraLocation"`
					NrLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ncgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalGnbID              struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalGnbId"`
					} `json:"nrLocation"`
					N3GaLocation struct {
						N3GppTai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"n3gppTai"`
						N3IwfID    string `json:"n3IwfId"`
						UeIpv4Addr string `json:"ueIpv4Addr"`
						UeIpv6Addr struct {
							UeIpv6Addr string `json:"ueIpv6Addr"`
						} `json:"ueIpv6Addr"`
						PortNumber int `json:"portNumber"`
					} `json:"n3gaLocation"`
				} `json:"userLocationInformation"`
				UetimeZone                       string `json:"uetimeZone"`
				PresenceReportingAreaInformation struct {
					AdditionalProp1 struct {
						PraID         string `json:"praId"`
						PresenceState struct {
							PresenceState string `json:"presenceState"`
						} `json:"presenceState"`
						TrackingAreaList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"trackingAreaList"`
						EcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgiList"`
						NcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgiList"`
						GlobalRanNodeIDList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalRanNodeIdList"`
					} `json:"additionalProp1"`
					AdditionalProp2 struct {
						PraID         string `json:"praId"`
						PresenceState struct {
							PresenceState string `json:"presenceState"`
						} `json:"presenceState"`
						TrackingAreaList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"trackingAreaList"`
						EcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgiList"`
						NcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgiList"`
						GlobalRanNodeIDList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalRanNodeIdList"`
					} `json:"additionalProp2"`
					AdditionalProp3 struct {
						PraID         string `json:"praId"`
						PresenceState struct {
							PresenceState string `json:"presenceState"`
						} `json:"presenceState"`
						TrackingAreaList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"trackingAreaList"`
						EcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgiList"`
						NcgiList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgiList"`
						GlobalRanNodeIDList []struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalRanNodeIdList"`
					} `json:"additionalProp3"`
				} `json:"presenceReportingAreaInformation"`
			} `json:"pFIContainerInformation"`
			TransmissionDataContainer []struct {
				LocalSequenceNumber     string    `json:"localSequenceNumber"`
				ChangeTime              time.Time `json:"changeTime"`
				CoverageStatus          bool      `json:"coverageStatus"`
				UserLocationInformation struct {
					EutraLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ecgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalNgenbID            struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalNgenbId"`
					} `json:"eutraLocation"`
					NrLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ncgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalGnbID              struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalGnbId"`
					} `json:"nrLocation"`
					N3GaLocation struct {
						N3GppTai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"n3gppTai"`
						N3IwfID    string `json:"n3IwfId"`
						UeIpv4Addr string `json:"ueIpv4Addr"`
						UeIpv6Addr struct {
							UeIpv6Addr string `json:"ueIpv6Addr"`
						} `json:"ueIpv6Addr"`
						PortNumber int `json:"portNumber"`
					} `json:"n3gaLocation"`
				} `json:"userLocationInformation"`
				DataVolume       int    `json:"dataVolume"`
				ChangeCondition  string `json:"changeCondition"`
				RadioResourcesID struct {
					RadioResourcesID string `json:"radioResourcesId"`
				} `json:"radioResourcesId"`
				RadioFrequency     string `json:"radioFrequency"`
				PC5RadioTechnology string `json:"pC5RadioTechnology"`
			} `json:"transmissionDataContainer"`
			ReceptionDataContainer []struct {
				LocalSequenceNumber     string    `json:"localSequenceNumber"`
				ChangeTime              time.Time `json:"changeTime"`
				CoverageStatus          bool      `json:"coverageStatus"`
				UserLocationInformation struct {
					EutraLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ecgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							EutraCellID string `json:"eutraCellId"`
						} `json:"ecgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalNgenbID            struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalNgenbId"`
					} `json:"eutraLocation"`
					NrLocation struct {
						Tai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"tai"`
						Ncgi struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							NrCellID string `json:"nrCellId"`
						} `json:"ncgi"`
						AgeOfLocationInformation int       `json:"ageOfLocationInformation"`
						UeLocationTimestamp      time.Time `json:"ueLocationTimestamp"`
						GeographicalInformation  string    `json:"geographicalInformation"`
						GeodeticInformation      string    `json:"geodeticInformation"`
						GlobalGnbID              struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							N3IwfID string `json:"n3IwfId"`
							GNbID   struct {
								BitLength int    `json:"bitLength"`
								GNBValue  string `json:"gNBValue"`
							} `json:"gNbId"`
							NgeNbID string `json:"ngeNbId"`
						} `json:"globalGnbId"`
					} `json:"nrLocation"`
					N3GaLocation struct {
						N3GppTai struct {
							PlmnID struct {
								Mcc string `json:"mcc"`
								Mnc string `json:"mnc"`
							} `json:"plmnId"`
							Tac string `json:"tac"`
						} `json:"n3gppTai"`
						N3IwfID    string `json:"n3IwfId"`
						UeIpv4Addr string `json:"ueIpv4Addr"`
						UeIpv6Addr struct {
							UeIpv6Addr string `json:"ueIpv6Addr"`
						} `json:"ueIpv6Addr"`
						PortNumber int `json:"portNumber"`
					} `json:"n3gaLocation"`
				} `json:"userLocationInformation"`
				DataVolume       int    `json:"dataVolume"`
				ChangeCondition  string `json:"changeCondition"`
				RadioResourcesID struct {
					RadioResourcesID string `json:"radioResourcesId"`
				} `json:"radioResourcesId"`
				RadioFrequency     string `json:"radioFrequency"`
				PC5RadioTechnology string `json:"pC5RadioTechnology"`
			} `json:"receptionDataContainer"`
		} `json:"proSeChargingInformation"`
	}
}
type ResourceId string

type Percentage float64

type Url string

type IpPacketFilterSet struct {
	SrcIp     string `json:"srcIp,omitempty"`
	DstIp     string `json:"dstIp,omitempty"`
	Protocol  int    `json:"protocol,omitempty"`
	SrcPort   int    `json:"srcPort,omitempty"`
	DstPort   int    `json:"dstPort,omitempty"`
	ToSTc     string `json:"toSTc,omitempty"`
	FlowLabel int    `json:"flowLabel,omitempty"`
	Spi       int    `json:"spi,omitempty"`
	Direction string `json:"direction,omitempty"`
}

type ServiceDataFlowDescription struct {
	FlowDescription IpPacketFilterSet `json:"flowDescription,omitempty"`
	DomainName      string            `json:"domainName,omitempty"`
}

type M5QoSSpecification struct {
	MarBwDlBitRate    BitRate `json:"marBwDlBitRate"`
	MarBwUlBitRate    BitRate `json:"marBwUlBitRate"`
	MinDesBwDlBitRate BitRate `json:"minDesBwDlBitRate,omitempty"`
	MinDesBwUlBitRate BitRate `json:"minDesBwUlBitRate,omitempty"`
	MirBwDlBitRate    BitRate `json:"mirBwDlBitRate"`
	MirBwUlBitRate    BitRate `json:"mirBwUlBitRate"`
	DesLatency        int     `json:"desLatency,omitempty"`
	DesLoss           int     `json:"desLoss,omitempty"`
}
type Unit string
type BitRate struct {
	Unit     Unit   `json:"unit,omitempty"`
	Value    uint32 `json:"value,omitempty"`
	Exponent int32  `json:"exponent,omitempty"`
}
type M1QoSSpecification struct {
	QoSReference        string  `json:"qosReference,omitempty"`
	MaxBtrUl            BitRate `json:"maxBtrUl,omitempty"`
	MaxBtrDl            BitRate `json:"maxBtrDl,omitempty"`
	MaxAuthBtrUl        BitRate `json:"maxAuthBtrUl,omitempty"`
	MaxAuthBtrDl        BitRate `json:"maxAuthBtrDl,omitempty"`
	DefPacketLossRateDl int     `json:"defPacketLossRateDl,omitempty"`
	DefPacketLossRateUl int     `json:"defPacketLossRateUl,omitempty"`
}

type ChargingSpecification struct {
	SponId     string   `json:"sponId,omitempty"`
	SponStatus string   `json:"sponStatus,omitempty"`
	Gpsi       []string `json:"gpsi,omitempty"`
}

type TypedLocation struct {
	LocationIdentifierType CellIdentifierType `json:"locationIdentifierType"`
	Location               string             `json:"location"`
}

type OperationSuccessResponse struct {
	Success bool   `json:"success"`
	Reason  string `json:"reason,omitempty"`
}

type CellIdentifierType string

const (
	Cgi  CellIdentifierType = "CGI"
	Ecgi CellIdentifierType = "ECGI"
	Ncgi CellIdentifierType = "NCGI"
)

type SdfMethod string

const (
	FiveTuple SdfMethod = "5_TUPLE"
	TwoTuple  SdfMethod = "2_TUPLE"
	// ... other SdfMethods skipped for brevity
)

type DurationSec time.Duration

type DateTime time.Time

type Uri string
type RANSecondaryRATUsageReport struct {
	RANSecondaryRATType struct {
		RANSecondaryRATType string `json:"rANSecondaryRATType"`
	} `json:"rANSecondaryRATType"`
	QosFlowsUsageReports []struct {
		QFI            int       `json:"qFI"`
		StartTimestamp time.Time `json:"startTimestamp"`
		EndTimestamp   time.Time `json:"endTimestamp"`
		UplinkVolume   int       `json:"uplinkVolume"`
		DownlinkVolume int       `json:"downlinkVolume"`
	} `json:"qosFlowsUsageReports"`
}
