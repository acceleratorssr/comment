package global

import (
	"comment/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// Config 此处需要类型记得是 *config.Config
	// Config应该是指针类型接收 /core/conf.go 中的结构体config
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
	Grpc     *config.Grpc
	GrpcConn *grpc.ClientConn
	SF       *singleflight.Group
)
