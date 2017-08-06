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

type RestServiceGroup struct {

	AlertStatus string `json:"alertStatus,omitempty"`

	NumOfServices int32 `json:"numOfServices,omitempty"`

	FullPath string `json:"fullPath,omitempty"`

	StopMonitoring bool `json:"stopMonitoring,omitempty"`

	HasServicesDisabled bool `json:"hasServicesDisabled,omitempty"`

	UserPermission string `json:"userPermission,omitempty"`

	TestLocation string `json:"testLocation,omitempty"`

	Description string `json:"description,omitempty"`

	DisableAlerting bool `json:"disableAlerting,omitempty"`

	GroupStatus string `json:"groupStatus,omitempty"`

	SdtStatus string `json:"sdtStatus,omitempty"`

	ParentId int32 `json:"parentId,omitempty"`

	Name string `json:"name"`

	NumOfDirectSubGroups int32 `json:"numOfDirectSubGroups,omitempty"`

	SubGroups []GroupData `json:"subGroups,omitempty"`

	Id int32 `json:"id,omitempty"`

	AlertStatusPriority int32 `json:"alertStatusPriority,omitempty"`

	AlertDisableStatus string `json:"alertDisableStatus,omitempty"`

	ServiceProperties []NameAndValue `json:"serviceProperties,omitempty"`

	NumOfDirectServices int32 `json:"numOfDirectServices,omitempty"`
}