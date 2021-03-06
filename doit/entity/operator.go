package entity

import "time"

const TableOperator = "operator"

type OperatorState int8

const (
	OperatorStateEnabled  OperatorState = 1 + iota
	OperatorStateDisabled               = 99
)

type Operator struct {
	ID           string                 `json:"id" gorm:"pk"`
	Name         string                 `json:"name"`
	PasswordHash []byte                 `json:"-"`
	RealName     string                 `json:"real_name"`
	State        OperatorState          `json:"state"`
	CreateTime   time.Time              `json:"create_time" gorm:"created"`
	UpdateTime   time.Time             `json:"update_time" gorm:"updated"`
}

type OperatorSession struct {
	Operator
	SignInTime     string `json:"sign_in_time"`
	LastSignInTime string `json:"last_sign_in_time"`
}

func (Operator) TableName() string {
	return TableOperator
}
