package request

// RegisterByUsernamePwd 用户名密码注册
type RegisterByUsernamePwd struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// LoginByUsernamePwd 用户名密码登录
type LoginByUsernamePwd struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ChangePassword 新旧密码普通修改密码
type ChangePassword struct {
	OldPassword string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword string `json:"new_password" form:"new_password" binding:"required"`
}

// UpdateUserProfile 修改账户信息: {性别,头像}
type UpdateUserProfile struct {
	Gender int    `json:"gender"`
	Avatar string `json:"avatar"`
}

// UpdateUserEmailByVerificationCode  修改账户邮箱(验证码)
type UpdateUserEmailByVerificationCode struct {
	Email            string `json:"email" form:"email" binding:"required"`
	VerificationCode string `json:"verification_code" form:"verification_code" binding:"required"`
}

// UpdateUserPhoneByVerificationCode  修改账户手机号(验证码)
type UpdateUserPhoneByVerificationCode struct {
	Phone            string `json:"phone" form:"phone" binding:"required"`
	VerificationCode string `json:"verification_code" form:"verification_code" binding:"required"`
}

// ClearUserPhoneByVerificationCode  解绑账户手机号(验证码)
type ClearUserPhoneByVerificationCode struct {
	Phone            string `json:"phone" form:"phone" binding:"required"`
	VerificationCode string `json:"verification_code" form:"verification_code" binding:"required"`
}
