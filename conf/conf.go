package conf

const (
  MgoDBName       = "test"
  MgoDBUrl        = "mongodb://localhost:27017/test"
  RedisDBUrl      = "localhost:6379"
  RedisExpireTime = 0
)

// jwt密钥
var JwtSecret = []byte("test!@#")
