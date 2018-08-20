package models

import (
	db "activitypro/database"
)

type VxUsers struct {
	Id                int    `form:"id" xorm:"not null pk autoincr index(idx) INT(11)"`
	Name              string `form:"name" xorm:"not null index(ids) VARCHAR(191)"`
	Email             string `form:"email" xorm:"unique VARCHAR(191)"`
	Password          string `form:"password" xorm:"VARCHAR(191)"`
	RememberToken     string `form:"remember_token" xorm:"VARCHAR(100)"`
	CreatedAt         int    `form:"created_at" xorm:"not null INT(11)"`
	UpdatedAt         int    `form:"updated_at" xorm:"INT(11)"`
	Tel               string `form:"tel" xorm:"not null comment('手机号') VARCHAR(15)"`
	QqLogin           string `form:"qq_login" xorm:"default '0' comment('qq登录') VARCHAR(255)"`
	WeixinLogin       string `form:"weixin_login" xorm:"default '0' comment('微信登录') VARCHAR(255)"`
	WeiboLogin        string `form:"weibo_login" xorm:"default '0' comment('微博登录') VARCHAR(255)"`
	Sex               int    `form:"sex" xorm:"not null default 0 comment('0未知1男2女') TINYINT(4)"`
	PersonalSignature string `form:"personal_signature" xorm:"default '懒懒的路人' comment('个性签名') VARCHAR(255)"`
	Profession        string `form:"profession" xorm:"comment('职业') VARCHAR(255)"`
	Residence         string `form:"residence" xorm:"comment('常居地') VARCHAR(255)"`
	Avatar            string `form:"avatar" xorm:"comment('头像') VARCHAR(255)"`
	LastIp            string `form:"last_ip" xorm:"comment('最后登录ip') VARCHAR(255)"`
	LastLogin         int    `form:"last_login" xorm:"comment('最后登录时间') INT(11)"`
	LoginCount        int    `form:"login_count" xorm:"not null default 0 comment('登录次数') INT(11)"`
	InviteAdminid     int    `form:"invite_adminid" xorm:"not null default 0 comment('管理人员邀请id') index(ids) INT(11)"`
	Status            int    `form:"status" xorm:"not null default 1 comment('1正常用户2禁用用户') TINYINT(3)"`
	BirthDay          string `form:"birth_day" xorm:"comment('出生日期') VARCHAR(100)"`
	Type              int    `form:"type" xorm:"not null default 1 comment('1普通用户 2大咖 3运营账号') index TINYINT(3)"`
	KolCreateAt       int    `form:"kol_create_at" xorm:"comment('大咖注册时间') INT(11)"`
	Salt              string `form:"salt" xorm:"not null comment('密码盐') CHAR(6)"`
	Truename          string `form:"truename" xorm:"comment('真实姓名') VARCHAR(100)"`
	Source            int    `form:"source" xorm:"default 1 comment('来源 1 菜谱APP 来源 2 问答APP') index(ids) index(idx) TINYINT(4)"`
	IsExistence       int    `form:"is_existence" xorm:"not null default 1 comment('1存在 2删除') index(ids) index(idx) TINYINT(4)"`
	Account           string `form:"account" xorm:"not null comment('手机号和账号') index(ids) CHAR(32)"`
	DeletedAt         int    `form:"deleted_at" xorm:"comment('删除时间') INT(11)"`
	RegistrationId    string `form:"registration_id" xorm:"comment('极光标识') CHAR(255)"`
	OperateBack       string `form:"operate_back" xorm:"comment('运营备注') CHAR(255)"`
	PhoneVersion      string `form:"phone_version" xorm:"not null default '未知设备' comment('登录手机设备号') CHAR(32)"`
	AppVersion        string `form:"app_version" xorm:"not null default '未知版本' comment('当前登录手机版本号') CHAR(32)"`
	Controller        int    `form:"controller" xorm:"default 3 comment('1ios 2安卓 3未知') TINYINT(4)"`
	OriginalAvatar    string `form:"original_avatar" xorm:"comment('原始头像') VARCHAR(255)"`
	IsExamine         int    `form:"is_examine" xorm:"not null default 1 comment('1待审核 2通过 3不通过') TINYINT(4)"`
	Pigeon            string `form:"pigeon" xorm:"comment('信鸽标识') CHAR(64)"`
	IsFollowRecommend int    `form:"is_follow_recommend" xorm:"not null default 1 comment('是否存在推荐关注 1未推荐 2推荐') TINYINT(4)"`
}

func (user *VxUsers) AddUser() (id int64, err error) {
	return db.Orm.Insert(&user)
}
