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

type RestDashboardGroup struct {

	FullPath string `json:"fullPath,omitempty"`

	UserPermission string `json:"userPermission,omitempty"`

	Name string `json:"name"`

	NumOfDirectSubGroups int64 `json:"numOfDirectSubGroups,omitempty"`

	NumOfDashboards int64 `json:"numOfDashboards,omitempty"`

	Description string `json:"description,omitempty"`

	Id int32 `json:"id,omitempty"`

	Dashboards []DashboardData `json:"dashboards,omitempty"`

	ParentId int32 `json:"parentId,omitempty"`

	NumOfDirectDashboards int64 `json:"numOfDirectDashboards,omitempty"`
}
