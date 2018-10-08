package conf

const (
  MongodbDisabled = true // 禁用mongodb
  RedisDisabled   = true  // 禁用redis
)

const (
  MongodbName     = "test"
  MongodbUrl      = "mongodb://localhost:27017/" + MongodbName
  RedisdbUrl      = "localhost:6379"
  RedisExpireTime = 0 // redis缓存的过期时间，0表示不设置过期时间
)

// jwt密钥
var JwtSecret = []byte("test!@#")
