package model

type User struct {
	UID      int64 `gorm:"primary_key" json:"uid"`
	PassType int32 `gorm:"not null;INDEX" json:"pass_type"`
	Point    int32 `gorm:"not null;INDEX" json:"point"`
}

// TableName return table name
func (*User) TableName() string {
	return "np_pass_users"
}

// TODO: LEVEL
func (u *User) Level() int32 {
	return u.Point / 1
}
