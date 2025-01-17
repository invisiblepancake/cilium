package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// PhysicalConnectionType is a nested struct in vpc response
type PhysicalConnectionType struct {
	Type                           string                            `json:"Type" xml:"Type"`
	Status                         string                            `json:"Status" xml:"Status"`
	CreationTime                   string                            `json:"CreationTime" xml:"CreationTime"`
	AdLocation                     string                            `json:"AdLocation" xml:"AdLocation"`
	ReservationActiveTime          string                            `json:"ReservationActiveTime" xml:"ReservationActiveTime"`
	ReservationOrderType           string                            `json:"ReservationOrderType" xml:"ReservationOrderType"`
	PortNumber                     string                            `json:"PortNumber" xml:"PortNumber"`
	Spec                           string                            `json:"Spec" xml:"Spec"`
	ChargeType                     string                            `json:"ChargeType" xml:"ChargeType"`
	ReservationInternetChargeType  string                            `json:"ReservationInternetChargeType" xml:"ReservationInternetChargeType"`
	Description                    string                            `json:"Description" xml:"Description"`
	Bandwidth                      int64                             `json:"Bandwidth" xml:"Bandwidth"`
	EnabledTime                    string                            `json:"EnabledTime" xml:"EnabledTime"`
	LineOperator                   string                            `json:"LineOperator" xml:"LineOperator"`
	PeerLocation                   string                            `json:"PeerLocation" xml:"PeerLocation"`
	RedundantPhysicalConnectionId  string                            `json:"RedundantPhysicalConnectionId" xml:"RedundantPhysicalConnectionId"`
	Name                           string                            `json:"Name" xml:"Name"`
	CircuitCode                    string                            `json:"CircuitCode" xml:"CircuitCode"`
	EndTime                        string                            `json:"EndTime" xml:"EndTime"`
	PortType                       string                            `json:"PortType" xml:"PortType"`
	BusinessStatus                 string                            `json:"BusinessStatus" xml:"BusinessStatus"`
	LoaStatus                      string                            `json:"LoaStatus" xml:"LoaStatus"`
	AccessPointId                  string                            `json:"AccessPointId" xml:"AccessPointId"`
	AccessPointType                string                            `json:"AccessPointType" xml:"AccessPointType"`
	HasReservationData             string                            `json:"HasReservationData" xml:"HasReservationData"`
	PhysicalConnectionId           string                            `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
	ProductType                    string                            `json:"ProductType" xml:"ProductType"`
	VirtualPhysicalConnectionCount int                               `json:"VirtualPhysicalConnectionCount" xml:"VirtualPhysicalConnectionCount"`
	ParentPhysicalConnectionId     string                            `json:"ParentPhysicalConnectionId" xml:"ParentPhysicalConnectionId"`
	ParentPhysicalConnectionAliUid int64                             `json:"ParentPhysicalConnectionAliUid" xml:"ParentPhysicalConnectionAliUid"`
	VlanId                         string                            `json:"VlanId" xml:"VlanId"`
	OrderMode                      string                            `json:"OrderMode" xml:"OrderMode"`
	VpconnStatus                   string                            `json:"VpconnStatus" xml:"VpconnStatus"`
	ExpectSpec                     string                            `json:"ExpectSpec" xml:"ExpectSpec"`
	ResourceGroupId                string                            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	AdDetailLocation               string                            `json:"AdDetailLocation" xml:"AdDetailLocation"`
	QosId                          string                            `json:"QosId" xml:"QosId"`
	Tags                           TagsInDescribePhysicalConnections `json:"Tags" xml:"Tags"`
}
