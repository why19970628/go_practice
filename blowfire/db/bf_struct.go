package db

import "database/sql"


type BfSlotList struct {
	Id           int64         `json:"id" xorm:"pk autoincr BIGINT(11) 'id'"`
	MachineId    int64         `json:"machine_id" xorm:"not null comment('机器序号') BIGINT(20) 'machine_id'"`
	MachineName  string        `json:"machine_name" xorm:"not null comment('机器名称') VARCHAR(128) 'machine_name'"`
	RequireLevel int64         `json:"require_level" xorm:"not null BIGINT(20) 'require_level'"`
	CardType     string        `json:"card_type" xorm:"not null comment('类型') VARCHAR(64) 'card_type'"`
	IsTempUnlock sql.NullInt64 `json:"is_temp_unlock" xorm:"TINYINT(1) 'is_temp_unlock'"`
	AddTime      sql.NullInt64 `json:"add_time" xorm:"default NULL BIGINT(20) 'add_time'"`
	UpdateTime   sql.NullInt64 `json:"update_time" xorm:"default NULL BIGINT(20) 'update_time'"`
	DelTime      sql.NullInt64 `json:"del_time" xorm:"default NULL BIGINT(20) 'del_time'"`
}
