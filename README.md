# Cleaning Service Go (Echo)
Pembuatan aplikasi Monitoring Cleaning Service dengan golang.

**Development by:** 
- Findryankp

## Init Project
- go mod init cleaningService

## Install Packages
* go get github.com/labstack/echo/v4
* go get github.com/joho/godotenv
* go get -u gorm.io/gorm
* go get -u gorm.io/driver/mysql

**Swaggo documentation API:** 
* go get -u github.com/swaggo/swag/cmd/swag
* go get -u github.com/swaggo/files
* go get -u github.com/swaggo/gin-swagger

## Step By step (MAC OS)
- export PATH=$(go env GOPATH)/bin:$PATH
- swag init