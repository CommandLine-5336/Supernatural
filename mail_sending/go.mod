module mail_sending

go 1.25.0

require gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df

require (
    github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/joho/godotenv v1.5.1
	github.com/robfig/cron/v3 v3.0.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.2
)
