/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * API version: 2018-11-01T03:57:11Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type Zone struct {

	// 존(Zone)번호
ZoneNo *string `json:"zoneNo,omitempty"`

	// 존(Zone)코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 존(Zone)명
ZoneName *string `json:"zoneName,omitempty"`

	// 존(Zone)설명
ZoneDescription *string `json:"zoneDescription,omitempty"`

	// 리전(Region)번호
RegionNo *string `json:"regionNo,omitempty"`
}
