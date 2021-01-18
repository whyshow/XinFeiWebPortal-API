package models

// 配置文件结构体
type HonorConf struct {
	Array []ClassifyArray `json:"array"`
}
type ClassifyArray struct {
	Id       string `json:"id"`
	Classify string `json:"classify"`
}
