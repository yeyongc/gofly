package api

import (
	"gogofly/dao"
	"gogofly/global/constants"
	"gogofly/service"
	"gogofly/service/dto"

	"github.com/gin-gonic/gin"
)

// user action types
const (
	ACTION_DELETE_USER      = 0x000001
	ACTION_UPDATE_USER      = 0x000002
	ACTION_GET_USER_BY_ID   = 0x000003
	ACTION_GET_USER_BY_NAME = 0x000004
	ACTION_ADD_USER         = 0x000005
	ACTION_GET_USER_LIST    = 0x000006
	ACTION_LOGIN_USER       = 0x000007
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

// @Summary 用户登录
// @Tag 用户管理
// @Param name formData string true "用户名"
// @Param password formData string true "用户密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/public/user/login [post]
func Login(ctx *gin.Context) {
	handleUserAction(ctx, UserAction{
		Type: ACTION_LOGIN_USER,
	})
}

// @Summary 用户添加
// @Tag 用户管理
// @Param name formData string true "用户名"
// @Param password formData string true "用户密码"
// @Param real_name fromData string false "姓名"
// @Param mobile formData string false "电话"
// @Param email formData string false "邮箱"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登陆失败"
// @Router /api/v1/user/add [post]
func AddUser(ctx *gin.Context) {
	handleUserAction(ctx, UserAction{
		Type: ACTION_ADD_USER,
	})

}

func GetUserById(ctx *gin.Context) {
	handleUserAction(ctx, UserAction{
		BindURI: true,
		Type:    ACTION_GET_USER_BY_ID,
	})
}

func GetUserList(ctx *gin.Context) {
	handleUserAction(ctx, UserAction{
		Type: ACTION_GET_USER_LIST,
	})
}

func UpdateUser(ctx *gin.Context) {
	handleUserAction(ctx, UserAction{
		BindURI: true,
		Type:    ACTION_UPDATE_USER,
	})
}

func DeleteUserById(ctx *gin.Context) {

	handleUserAction(ctx, UserAction{
		BindURI: true,
		Type:    ACTION_DELETE_USER,
	})

}

type UserAction struct {
	BindURI bool // 是否通过uri获取参数
	Type    int  // action 类型
}

// 处理用户请求行为
func handleUserAction(ctx *gin.Context, config UserAction) {
	m := NewUserApi(ctx)

	switch config.Type {

	// 用户删除
	case ACTION_DELETE_USER:
		var userIDDTO dto.UserIDDTO

		if !canBuildRequest(&m, &BuildRequestOption{DTO: &userIDDTO, BindParamsFromUri: config.BindURI}) {
			return
		}
		err := m.Service.DeleteUser(&userIDDTO)
		sendResponse(&m, err, constants.ERR_CODE_SERVICE_USER_DELETE, nil)

	// 用户更新
	case ACTION_UPDATE_USER:
		var userUpdateDTO dto.UserUpdatetDTO
		if !canBuildRequest(&m, &BuildRequestOption{DTO: &userUpdateDTO, BindParamsFromUri: config.BindURI}) {
			return
		}
		err := m.Service.UpdateUser(&userUpdateDTO)
		sendResponse(&m, err, constants.ERR_CODE_SERVICE_USER_UPDATE, nil)

	// 用户添加或者注册
	case ACTION_ADD_USER:
		var userAddDTO dto.UserAddDTO

		if !canBuildRequest(&m, &BuildRequestOption{DTO: &userAddDTO, BindParamsFromUri: config.BindURI}) {
			return
		}
		err := m.Service.AddUser(&userAddDTO)
		sendResponse(&m, err, constants.ERR_CODE_SERVICE_USER_ADD, map[string]any{
			"Msg": "注册成功",
		})

	// 根据id获取用户信息
	case ACTION_GET_USER_BY_ID:
		var userIDDTO dto.UserIDDTO

		if !canBuildRequest(&m, &BuildRequestOption{DTO: &userIDDTO, BindParamsFromUri: config.BindURI}) {
			return
		}
		user, err := m.Service.GetUserById(&userIDDTO)
		sendResponse(&m, err, constants.ERR_CODE_SERVICE_USER_GET, user)

	// 获取用户列表
	case ACTION_GET_USER_LIST:
		var userListDTO dto.UserListDTO

		if !canBuildRequest(&m, &BuildRequestOption{DTO: &userListDTO, BindParamsFromUri: config.BindURI}) {
			return
		}
		userList, userNum, err := m.Service.GetUserList(&userListDTO)
		sendResponse(&m, err, constants.ERR_CODE_SERVICE_USER_LIST_GET, map[string]any{
			"list": userList,
			"size": userNum,
		})

	// 用户登录
	case ACTION_LOGIN_USER:
		var userLoginDTO dto.UserLoginDTO
		if !canBuildRequest(&m, &BuildRequestOption{&userLoginDTO, config.BindURI}) {
			return
		}

		user, token, err := m.Service.Login(&userLoginDTO)
		sendResponse(&m, err, constants.ERR_CODE_SERVICE_USER_LOGIN, map[string]any{
			"user":  user,
			"token": token,
		})
	}

}

// 判断用户请求request内容是否满足api期待的内容
func canBuildRequest(m *UserApi, option *BuildRequestOption) bool {
	err := m.BuildRequest(option).GetErrors()
	return err == nil
}

// 返回响应
func sendResponse(m *UserApi, err error, statusCode int, data any) {
	if err != nil {
		m.ServerFail(ResponseJson{
			Msg:  err.Error(),
			Code: statusCode,
		})

	} else {
		m.Ok(ResponseJson{
			Data: data,
		})
	}
}

func NewUserApi(ctx *gin.Context) UserApi {
	return UserApi{
		BaseApi: NewBaseApi(ctx),
		Service: service.NewUserService(dao.NewUserDao()),
	}
}
