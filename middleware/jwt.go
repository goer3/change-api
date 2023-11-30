package middleware

import (
	"change-api/common"
	"change-api/dto"
	"change-api/model"
	"change-api/pkg/gedis"
	"change-api/pkg/response"
	"change-api/pkg/utils"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-module/carbon/v2"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           common.Config.JWT.Realm,                                // JWT 标识
		Key:             []byte(common.Config.JWT.Key),                          // 签名 Key
		Timeout:         time.Duration(common.Config.JWT.Timeout) * time.Second, // Token 有效期
		Authenticator:   authenticator,                                          // 用户登录校验
		PayloadFunc:     payloadFunc,                                            // Token 封装
		LoginResponse:   loginResponse,                                          // 登录成功响应
		Unauthorized:    unauthorized,                                           // 登录，认证失败响应
		IdentityHandler: identityHandler,                                        // 解析 Token
		Authorizator:    authorizator,                                           // 验证 Token
		LogoutResponse:  logoutResponse,                                         // 注销登录
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",     // Token 查找的字段
		TokenHeadName:   "Bearer",                                               // Token 请求头名称
	})
}

// 隶属 Login 中间件，当调用 LoginHandler 就会触发
// 通过从 ctx 中检索出数据，进行用户登录认证
// 返回包含用户信息的 Map 或者 Struct
func authenticator(ctx *gin.Context) (interface{}, error) {
	// 1.获取用户登录提交的数据
	var req dto.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, errors.New("获取用户登录信息失败")
	}

	// 2.获取客户端 IP，确保代理透传客户端真实 IP，如果获取 IP 失败则使用 None 做标识
	ip := ctx.ClientIP()
	if ip == "" {
		ip = "None"
	}

	// 最终存储在 Redis 中用于保存用户登录失败次数的 Key
	key := common.RedisKey.LoginWrongTimes + req.Account + ":" + ip

	// 3.获取 redis 中该用户登录错误次数，通过用户和 IP 进行绑定，避免恶意登录导致用户账户被误锁定
	var conn = gedis.NewOperation()
	// 密码错误次数，默认允许 5 次，如果到达上限，则返回账户锁定信息
	times := conn.GetInt(key).UnwrapWithDefault(0)
	if times >= common.Config.Login.WrongTimes {
		return nil, errors.New("认证次数超过上限，账户已锁定")
	}

	// 4.用户未锁定，则验证用户登录账户类型并查询用户，如果没查到则返回账户密码错误
	db := common.DB
	var user model.User
	var err error

	// 判断用户登录采用的方式，支持使用工号，手机号，Email
	dbt := db.Preload("Role").Preload("Department")
	if utils.IsMobile(req.Account) {
		err = dbt.Where("mobile = ?", req.Account).First(&user).Error
	} else if utils.IsEmail(req.Account) {
		err = dbt.Where("email = ?", req.Account).First(&user).Error
	} else {
		err = dbt.Where("job_id = ?", req.Account).First(&user).Error
	}

	if err != nil {
		common.Log.Error(err)
		return nil, errors.New("用户名或密码错误")
	}

	// 5.查询到用户，则进行密码校验
	if !utils.ComparePassword(user.Password, req.Password) {
		// 密码不对，则在原有的 redis 保存的错误次数上 +1，并设置过期时间
		times += 1
		conn.Set(key, times, gedis.WithExpire(time.Duration(common.Config.Login.LockTime)*time.Second))
		return nil, errors.New("用户名或密码错误")
	}

	// 6.密码正确，则进行用户状态校验
	// 禁用
	if user.Status == common.BrokenStatus {
		return nil, errors.New("用户已禁用，请联系管理员")
	}

	// 未激活
	if user.Status == common.UnActiveStatus {
		// 用户默认未激活，所以第一次登录会提示重置密码然后激活
		// 通过自定义关键字 UnActiveUser 并带上 id，后续方便直接就行判断
		return nil, fmt.Errorf("UnActiveUser:%s", user.Id)
	}

	// 未知状态
	if user.Status != common.NormalStatus {
		return nil, errors.New("用户处于未知状态，请联系管理员")
	}

	// 7.登录成功
	// 删除错误 redis 中的次数
	_, _ = conn.Del(key)

	// 更新数据库中登录信息
	common.DB.Model(&model.User{}).
		Where("id = ?", user.Id).
		Updates(map[string]interface{}{
			"last_login_ip":   ip,
			"last_login_time": carbon.DateTime{Carbon: carbon.Now()},
		})

	// 8.返回登录信息
	// 设置 Context，方便后面使用
	ctx.Set("id", user.Id)

	// 以指针的方式将数据传递给 PayloadFunc 函数继续处理
	return &user, nil
}

// 隶属 Login 中间件，接收 Authenticator 验证成功后传递过来的数据，进行封装成 Token
// MapClaims 必须包含 IdentityKey
// MapClaims 会被嵌入 Token 中，后续可以通过 ExtractClaims 对 Token 进行解析获取到
func payloadFunc(data interface{}) jwt.MapClaims {
	// 断言判断获取传递过来数据是不是用户数据
	if v, ok := data.(*model.User); ok {
		// 封装一些常用的字段，方便直接使用
		return jwt.MapClaims{
			jwt.IdentityKey: v.Id,        // id
			"UserName":      v.Name,      // 用户名字
			"RoleId":        v.Role.Id,   // 角色 Id
			"RoleName":      v.Role.Name, // 角色名称
		}
	}
	return jwt.MapClaims{}
}

// 隶属 Login 中间件，响应用户请求
// 接收 PayloadFunc 传递过来的 Token 信息，返回登录成功
func loginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	// 用户响应数据
	var res dto.LoginResponse
	res.Token = token
	res.Expire = expire.Format(common.SecTimeFormat)

	// 不允许多设备登录配置
	if !common.Config.Login.MultiDevices {
		// 获取前面 Context 设置的值，并验证是否合法
		id, _ := ctx.Get("id")
		if v, ok := id.(model.Nanoid); !ok || v == "" {
			response.FailedWithMessage("用户登录状态异常")
		}

		// 将新的 Token 存到 Redis 中，用户下一次请求的时候就去验证该 Token
		key := common.RedisKey.UserToken + string(id.(model.Nanoid))
		cache := gedis.NewOperation()
		cache.Set(key, token, gedis.WithExpire(time.Duration(common.Config.JWT.Timeout)*time.Second))
	}

	// 响应请求
	response.SuccessWithData(res)
}

// 登录失败，验证失败的响应
func unauthorized(ctx *gin.Context, code int, message string) {
	// 判断是否是用户未激活报错
	if strings.HasPrefix(message, "UnActiveUser") {
		// 获取 id
		id := strings.Split(message, ":")[1]

		// 生成随机字符串，该字符串作为重置密码的 Token
		token := utils.RandString(16)

		// 将数据保存到 Redis，后续用户可以根据该 Token 就行密码重置
		key := common.RedisKey.ResetPasswordToken + token
		cache := gedis.NewOperation()
		cache.Set(key, id, gedis.WithExpire(time.Duration(common.Config.Login.ResetTokenTime)*time.Second))

		// 响应客户端
		response.FailedWithCodeAndMessage(response.UnActived, token)
		return
	}
	response.FailedWithCodeAndMessage(response.Unauthorized, message)
}

// 用户登录后的中间件，用于解析 Token
func identityHandler(ctx *gin.Context) interface{} {
	// 从 Context 中获取用户的 id
	id, _ := utils.GetStringFromContext(ctx, "identity")
	return &model.User{
		Id: model.Nanoid(id),
	}
}

// 用户登录后的中间件，用于验证 Token
func authorizator(data interface{}, ctx *gin.Context) bool {
	user, ok := data.(*model.User)
	if ok && user.Id != "" {
		// 不允许多设备登录配置
		if !common.Config.Login.MultiDevices {
			// Key
			token := jwt.GetToken(ctx)
			key := common.RedisKey.UserToken + string(user.Id)

			// 验证该用户的 Token 和 Redis 中的是否一致
			cache := gedis.NewOperation()
			if cache.GetString(key).Unwrap() != token {
				return false
			}
		}
		return true
	}
	return false
}

// 注销登录
func logoutResponse(ctx *gin.Context, code int) {
	// 清理 Redis 保存的数据
	// Todo
	response.Success()
}

// 双因素认证 2FA
// 开启双因素认证之后，用户登录，先验证用户账户密码状态没问题后，判断用户是否绑定 2FA
// 如果未绑定，则让用户下载 2FA 软件并绑定
// 如果已经绑定，则弹出让用户输入验证码的框框
// 验证通过之后才返回对应的 Token
