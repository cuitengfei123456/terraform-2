package entity

type CreateLogtankRequestBody struct {
	Logtank *CreateLogTankOption `json:"logtank"`
}
type CreateLogTankOption struct{
	LoadBalancerId string `json:"loadbalance_id"`

	LogGroupId string `json:"log_group_id"`

	LogTopicId string `json:"log_topic_id"`
}