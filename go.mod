module backend

go 1.14

require (
	github.com/NewPage-Community/go-steam v0.0.0-20190902233638-f1c4971ee6b7
	github.com/alicebob/miniredis/v2 v2.11.4
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/json-iterator/go v1.1.9
	github.com/opentracing/opentracing-go v1.1.0
	github.com/robfig/cron/v3 v3.0.0
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/smartwalle/alipay/v3 v3.1.5
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.uber.org/zap v1.15.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	google.golang.org/genproto v0.0.0-20200608115520-7c474a2e3482
	google.golang.org/grpc v1.29.1
	gorm.io/datatypes v0.0.0-20200709131824-976937c55e2d
	gorm.io/driver/mysql v0.3.0
	gorm.io/gorm v1.20.0
)
