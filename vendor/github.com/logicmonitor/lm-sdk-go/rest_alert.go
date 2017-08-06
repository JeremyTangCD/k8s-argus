/* 
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package logicmonitor

type RestAlert struct {

	ResourceId int32 `json:"resourceId,omitempty"`

	InstanceName string `json:"instanceName,omitempty"`

	MonitorObjectId int32 `json:"monitorObjectId,omitempty"`

	EndEpoch int64 `json:"endEpoch,omitempty"`

	Rule string `json:"rule,omitempty"`

	Threshold string `json:"threshold,omitempty"`

	Type_ string `json:"type,omitempty"`

	StartEpoch int64 `json:"startEpoch,omitempty"`

	InternalId string `json:"internalId,omitempty"`

	AckComment string `json:"ackComment,omitempty"`

	MonitorObjectName string `json:"monitorObjectName,omitempty"`

	DataPointName string `json:"dataPointName,omitempty"`

	InstanceId int32 `json:"instanceId,omitempty"`

	DataPointId int32 `json:"dataPointId,omitempty"`

	NextRecipient int32 `json:"nextRecipient,omitempty"`

	Id string `json:"id,omitempty"`

	DetailMessage JsonObject `json:"detailMessage,omitempty"`

	RuleId int32 `json:"ruleId,omitempty"`

	AlertValue string `json:"alertValue,omitempty"`

	AckedBy string `json:"ackedBy,omitempty"`

	Severity int32 `json:"severity,omitempty"`

	Sdted bool `json:"sdted,omitempty"`

	AckedEpoch int64 `json:"ackedEpoch,omitempty"`

	Chain string `json:"chain,omitempty"`

	SDT JsonObject `json:"SDT,omitempty"`

	SubChainId int32 `json:"subChainId,omitempty"`

	ReceivedList string `json:"receivedList,omitempty"`

	CustomColumns JsonObject `json:"customColumns,omitempty"`

	Acked bool `json:"acked,omitempty"`

	ResourceTemplateType string `json:"resourceTemplateType,omitempty"`

	ClearValue string `json:"clearValue,omitempty"`

	InstanceDescription string `json:"instanceDescription,omitempty"`

	MonitorObjectGroups JsonArray `json:"monitorObjectGroups,omitempty"`

	ChainId int32 `json:"chainId,omitempty"`

	ResourceTemplateId int32 `json:"resourceTemplateId,omitempty"`

	Cleared bool `json:"cleared,omitempty"`

	ResourceTemplateName string `json:"resourceTemplateName,omitempty"`
}