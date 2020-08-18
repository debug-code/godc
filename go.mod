module godc

go 1.13

require (
	debug-code.com/dclog v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.15

)

replace debug-code.com/dclog => ./vendor/debug-code.com/dclog
