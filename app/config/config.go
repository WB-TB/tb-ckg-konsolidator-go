package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env"
)

type Configurations struct {
	HTTPPort              string `env:"HTTP_PORT"`
	HTTPServerTimeOut     int    `env:"HTTP_SERVER_TIMEOUT"`
	MongoConnectionString string `env:"MONGO_CONNECTION_STRING"`
	MongoDatabaseName     string `env:"MONGO_DBNAME"`
	MongoCollectionName   string `env:"MONGO_COLLECTION_NAME"`
	UnderMaintenance      bool   `env:"UNDER_MAINTENANCE"`
	RateLimit             int    `env:"RATE_LIMIT"`
	EntrypointMessage     string `env:"ENTRYPOINT_MSG"`
	APIKeySecret          string `env:"API_KEY_SECRET"`
	Environment           string `env:"environment"`

	// CKG TB
	CkgTbUseCache                   bool   `env:"CKG_TB_USE_CACHE"`
	CkgTbMongoDatabaseName          string `env:"CKG_TB_MONGO_DBNAME"`
	CkgTbMongoCollectionTransaction string `env:"CKG_TB_MONGO_COLLECTION_TRANSACTION"`
	CkgTbGetDataPageSize            int    `env:"CKG_TB_GET_DATA_PAGE_SIZE" envDefault:"1000"`
	CkgTbGetDataTimeout             int    `env:"CKG_TB_GET_DATA_TIMEOUT" envDefault:"60"`
	CkgTbPostDataTimeout            int    `env:"CKG_TB_POST_DATA_TIMEOUT" envDefault:"180"`
}

var (
	configuration Configurations
	mutex         sync.Once
)

func GetConfig() Configurations {
	mutex.Do(func() {
		configuration = newConfig()
	})

	return configuration
}

func newConfig() Configurations {
	var cfg = Configurations{}
	if err := env.Parse(&cfg); err != nil {
		log.Panic(err.Error())
	}

	return cfg
}
