module github.com/EDDYCJY/go-gin-example

go 1.17

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/astaxie/beego v1.12.3
	github.com/boombuler/barcode v1.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ini/ini v1.66.4
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/swaggo/gin-swagger v1.4.2
	github.com/swaggo/swag v1.8.1
	github.com/tealeg/xlsx v1.0.5
	github.com/unknwon/com v1.0.1
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/image v0.0.0-20220413100746-70e8d0d3baa9 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => /home/go-application/go-gin-example/pkg/conf
	github.com/EDDYCJY/go-gin-example/docs => /home/go-application/go-gin-example/docs
	github.com/EDDYCJY/go-gin-example/middleware => /home/go-application/go-gin-example/middleware
	github.com/EDDYCJY/go-gin-example/models => /home/go-application/go-gin-example/models
	github.com/EDDYCJY/go-gin-example/pkg/e => /home/go-application/go-gin-example/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/file => /home/go-application/go-gin-example/pkg/file
	github.com/EDDYCJY/go-gin-example/pkg/logging => /home/go-application/go-gin-example/pkg/logging
	github.com/EDDYCJY/go-gin-example/pkg/setting => /home/go-application/go-gin-example/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => /home/go-application/go-gin-example/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => /home/go-application/go-gin-example/routers
	github.com/EDDYCJY/go-gin-example/routers/api => /home/go-application/go-gin-example/routers/api
)
