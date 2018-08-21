/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-08-20T06:35:09Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetNasVolumeInstanceListRequest struct {

	// 볼륨할당프로토콜유형코드
VolumeAllotmentProtocolTypeCode *string `json:"volumeAllotmentProtocolTypeCode,omitempty"`

	// 이벤트설정여부
IsEventConfiguration *bool `json:"isEventConfiguration,omitempty"`

	// 스냅샷볼륨설정여부
IsSnapshotConfiguration *bool `json:"isSnapshotConfiguration,omitempty"`

	// NAS볼륨인스턴스번호리스트
NasVolumeInstanceNoList []*string `json:"nasVolumeInstanceNoList,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`
}