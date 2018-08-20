package models

import db "activitypro/database"

type VxActivity struct {
	Id               int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Title            string `json:"title" xorm:"not null comment('标题') VARCHAR(255)"`
	Content          string `json:"content" xorm:"comment('正文') TEXT"`
	ClassificationId string `json:"classification_id" xorm:"comment('官方分类id') VARCHAR(255)"`
	Uid              int    `json:"uid" xorm:"not null comment('提问者id') index(idx) INT(11)"`
	CreatedId        int    `json:"created_id" xorm:"not null comment('创建者id') INT(11)"`
	Type             int    `json:"type" xorm:"not null default 1 comment('1启用 2禁用') index(idx) TINYINT(4)"`
	CreatedAt        int    `json:"created_at" xorm:"not null comment('创建时间') INT(11)"`
	Follow           int    `json:"follow" xorm:"not null default 0 comment('关注数') INT(11)"`
	Attend           int    `json:"attend" xorm:"not null default 0 comment('参加数') INT(11)"`
	DeletedAt        int    `json:"deleted_at" xorm:"not null comment('删除时间') INT(11)"`
	IsExistence      int    `json:"is_existence" xorm:"not null default 1 comment('1存在 2删除') index index(idx) TINYINT(4)"`
	Status           int    `json:"status" xorm:"not null default 1 comment('1非精选 非推荐 2推荐池 3精选 4移除精选') TINYINT(4)"`
	Views            int    `json:"views" xorm:"not null default 0 comment('查看次数') INT(11)"`
	UpdatedAt        int    `json:"updated_at" xorm:"comment('修改时间') INT(11)"`
	Controller       int    `json:"controller" xorm:"default 4 comment('1ios 2安卓 3PC 4未知（发布手机类型）') TINYINT(4)"`
	PhoneVersion     string `json:"phone_version" xorm:"comment('手机版本号') CHAR(32)"`
	AppVersion       string `json:"app_version" xorm:"comment('当前APP的版本号') CHAR(32)"`
	IsExamine        int    `json:"is_examine" xorm:"not null default 1 comment('1待审核 2审核通过 3审核 不通过 4进入审核列表') TINYINT(4)"`
	SendAt           int    `json:"send_at" xorm:"not null default 0 comment('最后一次答案发布时间') INT(11)"`
}

func (act *VxActivity) AddActivity() (id int64, err error) {
	return db.Orm.Insert(&act)
}
