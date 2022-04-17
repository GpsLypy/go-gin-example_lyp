module github.com/EDDYCJY/go-gin-example

go 1.17

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

require (
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => /home/go-application/go-gin-example/pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => /home/go-application/go-gin-example/middleware
	github.com/EDDYCJY/go-gin-example/models => /home/go-application/go-gin-example/models
	github.com/EDDYCJY/go-gin-example/pkg/e => /home/go-application/go-gin-example/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => /home/go-application/go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => /home/go-application/go-gin-example/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => /home/go-application/go-gin-example/routers
	github.com/EDDYCJY/go-gin-example/routers/api => /home/go-application/go-gin-example/routers/api
)
