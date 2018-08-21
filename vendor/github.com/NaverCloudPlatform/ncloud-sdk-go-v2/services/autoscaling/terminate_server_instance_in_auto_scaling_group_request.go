/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-08-07T06:47:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type TerminateServerInstanceInAutoScalingGroupRequest struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo"`

	// 기대용량치감소여부
ShouldDecrementDesiredCapacity *bool `json:"shouldDecrementDesiredCapacity"`
}