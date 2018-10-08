package context

import (
  "net/http"
  "strconv"
  "strings"
  "time"
  "web/db"
  "web/errgo"
  "web/util"

  "github.com/gin-gonic/gin"
  "github.com/gomodule/redigo/redis"
  "github.com/tidwall/gjson"
  "gopkg.in/mgo.v2"
)

type New struct {
  Ctx         *gin.Context
  RawData     []byte
  MgoDB       *mgo.Database
  MgoDBCloser func()
  Redis       redis.Conn
  Errgo       *errgo.Stack
}

func CreateCtx(fn func(*New)) func(*gin.Context) {
  return func(c *gin.Context) {
    util.Println()
    util.Println("------------------------------------------")
    util.Println()

    // 创建上下文
    ctx, err := NewCtx(c)

    // 如果创建过程中有报错，返回错误
    if err != nil {
      ctx.Error(err)
      return
    }

    // defer
    defer ctx.Close()

    // 调用控制器
    fn(ctx)
  }
}

// 创建上下文，连接mgo与redis数据库
func NewCtx(c *gin.Context) (*New, error) {
  bytes, _ := c.GetRawData()

  mg, closer, err := db.CloneMgoDB()
  if err != nil {
    util.Println("[MGO] 😈 Error")
    return nil, err
  }
  if mg != nil {
    util.Println("[MGO] 😄 OK")
  }

  rds := db.GetRedis()
  if rds != nil {
    util.Println("[RDS] 😄 OK")
  }

  return &New{
    c,
    bytes,
    mg,
    closer,
    rds,
    errgo.Create(),
  }, nil
}

// 创建不连接数据库的上下文
func NewBaseCtx(c *gin.Context) *New {
  bytes, _ := c.GetRawData()
  return &New{
    Ctx:     c,
    RawData: bytes,
    Errgo:   errgo.Create(),
  }
}

// 关闭数据库连接
func (c *New) Close() {
  if c.MgoDBCloser != nil {
    c.MgoDBCloser()
    util.Println("[MGO] 👋 CLOSED")
  }
  if c.Redis != nil {
    c.Redis.Close()
    util.Println("[RDS] 👋 CLOSED")
  }
}

// 成功处理
func (c *New) Success(data gin.H) {
  respH := gin.H{
    "msg":  "ok",
    "code": 0,
  }

  if len(data) > 1 { // Almost the length is more than 1, so just check it first.
    respH["data"] = data
  } else if data["data"] != nil {
    respH["data"] = data["data"]
  } else if data != nil && len(data) > 0 {
    respH["data"] = data
  }

  status := http.StatusOK

  if data == nil {
    status = http.StatusNoContent
  }

  c.Ctx.JSON(status, respH)
}

// 处理错误
func (c *New) Error(errNo interface{}) {

  // 根据错误号获取错误内容（错误号是个int或error）
  err := errgo.Get(errNo)

  util.Println()
  util.Println(" >>> ERROR:", err.Message)
  util.Println(" >>> ERROR CODE:", err.Code)
  util.Println(" >>> REQUEST METHOD:", c.Ctx.Request.Method)
  util.Println(" >>> REQUEST URL:", c.Ctx.Request.URL.String())
  util.Println(" >>> USER AGENT:", c.Ctx.Request.UserAgent())
  util.Println()

  // 清除错误栈
  c.Errgo.ClearErrorStack()

  c.Ctx.JSON(err.Status, gin.H{
    "msg":  err.Message,
    "code": err.Code,
    "data": nil,
  })
}

// 从body中取string类型的值
func (c *New) GetRaw(key string) (string, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return strings.TrimSpace(res.Str), res.Exists()
}

// 从body中取string类型的值，没有则使用默认值
func (c *New) GetRawDefault(key string, def string) string {
  val, ok := c.GetRaw(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取array类型的值
func (c *New) GetRawArray(key string) ([]gjson.Result, bool) {
  res := gjson.GetBytes(c.RawData, key)
  arr := res.Array()
  return arr, res.Exists()
}

// 从body中取array类型的值，没有则使用默认值
func (c *New) GetRawArrayDefault(key string, def []gjson.Result) []gjson.Result {
  val, ok := c.GetRawArray(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取time类型的值
func (c *New) GetRawTime(key string) (time.Time, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return res.Time(), res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawTimeDefault(key string, def time.Time) time.Time {
  val, ok := c.GetRawTime(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取int类型的值
func (c *New) GetRawInt(key string) (int, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return int(res.Int()), res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawIntDefault(key string, def int) int {
  val, ok := c.GetRawInt(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取bool类型的值
func (c *New) GetRawBool(key string) (bool, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return res.Bool(), res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawBoolDefault(key string, def bool) bool {
  val, ok := c.GetRawBool(key)
  if !ok {
    return def
  }
  return val
}

// 从body中取json类型的值
func (c *New) GetRawJSON(key string) (gjson.Result, bool) {
  res := gjson.GetBytes(c.RawData, key)
  return res, res.Exists()
}

// 从body中取time类型的值，没有则使用默认值
func (c *New) GetRawJSONDefault(key string, def gjson.Result) gjson.Result {
  val, ok := c.GetRawJSON(key)
  if !ok {
    return def
  }
  return val
}

// 从params中取string类型的值
func (c *New) GetParam(key string) (string, bool) {
  res, ok := c.Ctx.Params.Get(key)
  return res, ok
}

// 从params中取string类型的值，没有则使用默认值
func (c *New) GetParamDefault(key string, def string) string {
  val, ok := c.GetParam(key)
  if !ok {
    return def
  }
  return val
}

// 从params中取int类型的值
func (c *New) GetParamInt(key string) (int, bool) {
  res, ok := c.Ctx.Params.Get(key)
  intRes, _ := strconv.Atoi(res)
  return intRes, ok
}

// 从params中取int类型的值，没有则使用默认值
func (c *New) GetParamIntDefault(key string, def int) int {
  val, ok := c.GetParamInt(key)
  if !ok {
    return def
  }
  return val
}

// 从params中取bool类型的值
func (c *New) GetParamBool(key string) (bool, bool) {
  res, ok := c.Ctx.Params.Get(key)
  return res == "true", ok
}

// 从params中取bool类型的值，没有则使用默认值
func (c *New) GetParamBoolDefault(key string, def bool) bool {
  val, ok := c.GetParamBool(key)
  if !ok {
    return def
  }
  return val
}

// 从query中取string类型的值
func (c *New) GetQuery(key string) (string, bool) {
  res, ok := c.Ctx.GetQuery(key)
  return res, ok
}

// 从query中取string类型的值，没有则使用默认值
func (c *New) GetQueryDefault(key string, def string) string {
  val, ok := c.GetQuery(key)
  if !ok {
    return def
  }
  return val
}

// 从query中取int类型的值
func (c *New) GetQueryInt(key string) (int, bool) {
  res, ok := c.Ctx.GetQuery(key)
  if !ok {
    return 0, false
  }
  intRes, err := strconv.Atoi(res)
  if err != nil {
    return 0, false
  }
  return intRes, true
}

// 从query中取int类型的值，没有则使用默认值
func (c *New) GetQueryIntDefault(key string, def int) int {
  val, ok := c.GetQueryInt(key)
  if !ok {
    return def
  }
  return val
}

// 从query中取bool类型的值
func (c *New) GetQueryBool(key string) (bool, bool) {
  res, ok := c.Ctx.GetQuery(key)
  if !ok {
    return false, false
  }
  return res == "true", true
}

// 从query中取bool类型的值，没有则使用默认值
func (c *New) GetQueryBoolDefault(key string, def bool) bool {
  val, ok := c.GetQueryBool(key)
  if !ok {
    return def
  }
  return val
}

// 从context中取字符串类型的值
func (c *New) Get(key string) (string, bool) {
  res, ok := c.Ctx.Get(key)
  if !ok {
    return "", false
  }
  return res.(string), true
}

// 从context中取bool类型的值
func (c *New) GetInt(key string) (int, bool) {
  res, ok := c.Ctx.Get(key)
  if !ok {
    return 0, false
  }
  return res.(int), true
}

// get value by string
func (c *New) GetBool(key string) (bool, bool) {
  res, ok := c.Ctx.Get(key)
  if !ok {
    return false, false
  }
  return res.(bool), true
}
