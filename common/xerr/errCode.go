package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const SERVER_COMMON_ERROR uint32 = 100001

//请求参数错误
const REQUEST_PARAM_ERROR uint32 = 100002

//Token过期错误
const TOKEN_EXPIRE_ERROR uint32 = 100003

//Token生成错误
const TOKEN_GENERATE_ERROR uint32 = 100004

//数据库相关错误
const DB_ERROR uint32 = 100005

//用户模块
