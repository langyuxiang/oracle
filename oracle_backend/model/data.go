package model

import (
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	DataType        string  `json:"data_type,omitempty"`
	DataSourceId    int     `json:"data_source_id,omitempty" gorm:"foreignKey:FDataSourceId;references:DataSourceId"`
	NeedWorkTime    string  `json:"need_work_time"`
	IsBelievable    bool    `json:"is_believable,omitempty"`
	BelievableLevel int     `json:"believable_level,omitempty"`
	DoEvilNum       int     `json:"do_evil_num,omitempty"`
	DataImportance  float32 `json:"data_importance,omitempty"`
	NodeFailureRate float32 `json:"node_failure_rate,omitempty"`
	SourceChannel   float32 `json:"source_channel,omitempty"`
	Integrity       float32 `json:"integrity,omitempty"`
	UsageRate       float32 `json:"usage_rate,omitempty"`
}
