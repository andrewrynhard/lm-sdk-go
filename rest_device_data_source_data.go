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

type RestDeviceDataSourceData struct {

	Instances map[string]RestRawDataValues `json:"instances,omitempty"`

	DataPoints []string `json:"dataPoints,omitempty"`

	DataSourceName string `json:"dataSourceName,omitempty"`
}
