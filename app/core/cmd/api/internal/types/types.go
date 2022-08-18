// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type UserInfoResp struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type UserProfileInfoResp struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int64  `json:"gender"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
}

type UpdateProfileReq struct {
	Username string `json:"username" validate:"required,min=2,max=12"`
	Nickname string `json:"nickname"`
	Gender   int64  `json:"gender" validate:"gte=0,lte=2"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
}

type Menu struct {
	Id           int64  `json:"id"`
	ParentId     int64  `json:"parentId"`
	Name         string `json:"name"`
	Router       string `json:"router"`
	Type         int64  `json:"type"`
	Icon         string `json:"icon"`
	OrderNum     int64  `json:"orderNum"`
	ViewPath     string `json:"viewPath"`
	IsShow       int64  `json:"isShow"`
	ActiveRouter string `json:"activeRouter"`
}

type UserPermMenuResp struct {
	Menus []Menu   `json:"menus"`
	Perms []string `json:"perms"`
}

type UpdatePasswordReq struct {
	OldPassword string `json:"oldPassword" validate:"min=6,max=12"`
	NewPassword string `json:"newPassword" validate:"min=6,max=12"`
}

type LoginCaptchaResp struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
}

type PermMenu struct {
	Id           int64    `json:"id"`
	ParentId     int64    `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         int64    `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     int64    `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       int64    `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
}

type SysPermMenuListResp struct {
	PermMenuList []PermMenu `json:"list"`
}

type AddSysPermMenuReq struct {
	ParentId     int64    `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         int64    `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     int64    `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       int64    `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
}

type DeleteSysPermMenuReq struct {
	Id int64 `json:"id"`
}

type UpdateSysPermMenuReq struct {
	Id           int64    `json:"id"`
	ParentId     int64    `json:"parentId"`
	Name         string   `json:"name"`
	Router       string   `json:"router"`
	Perms        []string `json:"perms"`
	Type         int64    `json:"type"`
	Icon         string   `json:"icon"`
	OrderNum     int64    `json:"orderNum"`
	ViewPath     string   `json:"viewPath"`
	IsShow       int64    `json:"isShow"`
	ActiveRouter string   `json:"activeRouter"`
}

type Role struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	UniqueKey string `json:"uniqueKey"`
	Remark    string `json:"remark"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
}

type SysRoleListResp struct {
	RoleList []Role `json:"list"`
}

type AddSysRoleReq struct {
	ParentId    int64   `json:"parentId"`
	Name        string  `json:"name"`
	UniqueKey   string  `json:"uniqueKey"`
	PermMenuIds []int64 `json:"permMenuIds"`
	Remark      string  `json:"remark"`
	Status      int64   `json:"status"`
	OrderNum    int64   `json:"orderNum"`
}

type DeleteSysRoleReq struct {
	Id int64 `json:"id"`
}

type UpdateSysRoleReq struct {
	Id          int64   `json:"id"`
	ParentId    int64   `json:"parentId"`
	Name        string  `json:"name"`
	UniqueKey   string  `json:"uniqueKey"`
	PermMenuIds []int64 `json:"permMenuIds"`
	Remark      string  `json:"remark"`
	Status      int64   `json:"status"`
	OrderNum    int64   `json:"orderNum"`
}

type UpdateSysRolePermMenuReq struct {
	Id          int64   `json:"id"`
	PermMenuIds []int64 `json:"permMenuIds"`
}

type Dept struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
	Remark    string `json:"remark"`
}

type SysDeptListResp struct {
	DeptList []Dept `json:"list"`
}

type AddSysDeptReq struct {
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
	Remark    string `json:"remark"`
}

type DeleteSysDeptReq struct {
	Id int64 `json:"id"`
}

type UpdateSysDeptReq struct {
	Id        int64  `json:"id"`
	ParentId  int64  `json:"parentId"`
	Name      string `json:"name"`
	FullName  string `json:"fullName"`
	UniqueKey string `json:"uniqueKey"`
	Status    int64  `json:"status"`
	OrderNum  int64  `json:"orderNum"`
	Remark    string `json:"remark"`
}

type TransferSysDeptReq struct {
	Id       int64 `json:"id"`
	ParentId int64 `json:"parentId"`
}

type Job struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
}

type SysJobListResp struct {
	JobList []Job `json:"list"`
}

type AddSysJobReq struct {
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
}

type DeleteSysJobReq struct {
	Id int64 `json:"id"`
}

type UpdateSysJobReq struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
}

type Profession struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
}

type SysProfessionListResp struct {
	ProfessionList []Profession `json:"list"`
}

type AddSysProfessionReq struct {
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
}

type DeleteSysProfessionReq struct {
	Id int64 `json:"id"`
}

type UpdateSysProfessionReq struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	OrderNum int64  `json:"orderNum"`
}

type User struct {
	Id         int64    `json:"id"`
	Account    string   `json:"account"`
	Username   string   `json:"username"`
	Nickname   string   `json:"nickname"`
	Avatar     string   `json:"avatar"`
	Gender     int64    `json:"gender"`
	Birthday   string   `json:"birthday"`
	Email      string   `json:"email"`
	Mobile     string   `json:"mobile"`
	Profession string   `json:"profession"`
	Job        string   `json:"job"`
	Dept       string   `json:"dept"`
	Roles      []string `json:"roles"`
	Status     int64    `json:"status"`
	OrderNum   int64    `json:"orderNum"`
	Remark     string   `json:"remark"`
}

type SysUserListReq struct {
	Page    int64   `json:"page"`
	Limit   int64   `json:"limit"`
	DeptIds []int64 `json:"deptIds"`
}

type Pagination struct {
	Page  int64 `json:"page"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}

type SysUserListResp struct {
	UserList   []User     `json:"list"`
	Pagination Pagination `json:"pagination"`
}

type AddSysUserReq struct {
	Account      string  `json:"account"`
	Username     string  `json:"username"`
	Nickname     string  `json:"nickname"`
	Avatar       string  `json:"avatar"`
	Gender       int64   `json:"gender"`
	Birthday     string  `json:"birthday"`
	Email        string  `json:"email"`
	Mobile       string  `json:"mobile"`
	ProfessionId int64   `json:"profession"`
	JobId        int64   `json:"jobId"`
	DeptId       int64   `json:"deptId"`
	RoleIds      []int64 `json:"roleIds"`
	Status       int64   `json:"status"`
	OrderNum     int64   `json:"orderNum"`
	Remark       string  `json:"remark"`
}

type DeleteSysUserReq struct {
	Id int64 `json:"id"`
}

type UpdateSysUserReq struct {
	Id           int64   `json:"id"`
	Username     string  `json:"username"`
	Nickname     string  `json:"nickname"`
	Avatar       string  `json:"avatar"`
	Gender       int64   `json:"gender"`
	Birthday     string  `json:"birthday"`
	Email        string  `json:"email"`
	Mobile       string  `json:"mobile"`
	ProfessionId int64   `json:"profession"`
	JobId        int64   `json:"jobId"`
	DeptId       int64   `json:"deptId"`
	RoleIds      []int64 `json:"roleIds"`
	Status       int64   `json:"status"`
	OrderNum     int64   `json:"orderNum"`
	Remark       string  `json:"remark"`
}

type UpdateSysUserPasswordReq struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

type TransferSysUserReq struct {
	Id     int64 `json:"id"`
	DeptId int64 `json:"deptId"`
}
