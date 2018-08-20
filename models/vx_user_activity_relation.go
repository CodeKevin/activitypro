package models

type VxUserActivityRelation struct {
	Id         int `json:"id" xorm:"not null pk autoincr INT(11)"`
	Uid        int `json:"uid" xorm:"not null comment('用户id') index(idx) INT(11)"`
	ActivityId int `json:"activity_id" xorm:"not null comment('活动id') index(idx) INT(11)"`
	Status     int `json:"status" xorm:"default 1 comment('1参加 2取消参加') index(idx) TINYINT(4)"`
	CreatedAt  int `json:"created_at" xorm:"not null comment('报名时间') INT(11)"`
	UpdatedAt  int `json:"updated_at" xorm:"comment('取消报名时间') INT(11)"`
	Type       int `json:"type" xorm:"default 1 comment('1问题正常状态 2问题禁用状态') TINYINT(4)"`
}
