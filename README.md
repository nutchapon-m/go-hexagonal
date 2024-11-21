# Go Hexagonal Architecture Template
### Go Hexagonal Architecture Template สำหรับผู้ที่สนใจอยากจะนำไปใช้เป็น template สำหรับขึ้นโปรเจคในการทำงาน API หลังบ้านต่าง ๆ ของ Web application
ความรู้เกี่ยวกับภาษา go อื่นๆ สามารถศึกษาเพิ่มเติมได้ที่ [codebangkok](https://github.com/codebangkok/golang) ซึ่งถือเป็นแหล่งความรู้ที่ทำให้ repository นี้เกิดขึ้น
#### คำสั่งสำหรับการ clone repository นี้
```bash
git clone https://github.com/NutchaponMet/go-hexagonal-architecture-template.git
go mod init <your-project-name>
```
> [!NOTE]
> ภายในโปรเจคนี้จะประกอบไปด้วย library ที่สำคัญอยู่หลายตัวที่ใช้สำหรับสร้าง Backend Web application
> ซึ่งผู้ที่นำ template นี้ไปใช้งานจะต้องมีความเข้าใจใน library พื้นฐานดังกล่าวก่อน ทั้งนี้ สามารถเข้าไปอ่านทำความเข้าใจเพิ่มเติมได้ตามหัวข้อด้านล่างนี้ได้เลย

### Project Structure
```bash
├── Dockerfile
├── README.md
├── cmd
│   └── main.go
├── config.yml
├── external
│   ├── cache
│   │   └── redis.go
│   ├── db
│   │   ├── mysql.go
│   │   └── postgres.go
│   └── queues
│       └── rmqp.go
├── go.mod
├── go.sum
├── internal
│   ├── configs
│   │   └── configs.go
│   ├── core
│   │   ├── domains
│   │   │   └── user.go
│   │   ├── ports
│   │   │   ├── auth_ports.go
│   │   │   └── user_ports.go
│   │   └── services
│   │       ├── auth_services.go
│   │       └── utils.go
│   ├── handlers
│   │   ├── auth_handlers.go
│   │   └── utils.go
│   ├── pkgs
│   │   ├── errs
│   │   │   └── errs.go
│   │   ├── logs
│   │   │   └── logs.go
│   │   └── utils
│   │       └── utils.go
│   └── repositories
│       └── user_repositories.go
└── server
    ├── middleware.go
    ├── routes
    │   └── v1
    │       └── auth.go
    └── server.go
```

---
### 1. [Go fiber](https://docs.gofiber.io/) - web framework
```bash
go get github.com/gofiber/fiber/v2
```
-----------
### 2. [Gorm](https://gorm.io/index.html) - ตัวจัดการงานเกี่ยวกับฐานข้อมูล ติดกับฐานข้อมูล
```bash
go get -u gorm.io/gorm
```
> ### Driver
> sqlite driver
> ```bash
> go get -u gorm.io/driver/sqlite
> ```
> MySql driver
> ```bash
> go get -u gorm.io/driver/mysql
> ```
> Postgres driver
> ```bash
> go get -u gorm.io/driver/postgres
> ```
------------
### 3. [zap](https://pkg.go.dev/go.uber.org/zap) - ตัวจัดการ log
```bash
go get go.uber.org/zap

## ซึ่ง fiber ก็มี package ที่ลองรับการทำ log middleware ร่วมกับ zap ด้วย

go get github.com/gofiber/contrib/fiberzap/v2
```
------------
### 4. [viper](https://pkg.go.dev/github.com/spf13/viper) - ตัวจัดการ configuration
```bash
go get github.com/spf13/viper
```
------------
## Configuration File
##### config.yml --> dev configuration
```yaml
# dev config file
app:
  port: 5555
db: 
  username: "admin0001"
  password: "admin1234"
  host: "localhost"
  port: 4444
  dbname: "pgdb"
mongodb:
  driver: "mongodb"
  host: "localhost"
  port: 8080
  username: "user"
  password: "password"
  dbname: "you_db"
infisical:
  mode: dev
  url: "http://example.com"
  clientID: "xxxxxxxx"
  clientSecret: "xxxxxxxxx"
  projectID: "xxxxxxxxxx"
```
##### .env --> production configuration environment file
```.env
APP_PORT=5555
DB_USERNAME="admin0001"
DB_PASSWORD="admin1234"
DB_HOST="db"
DB_PORT=5432
DB_DBNAME="appdb"
MONGODB_DRIVER="mongodb"
MONGODB_USERNAME="you_username"
MONGODB_PASSWORD="your_password"
MONGODB_HOST="localhost"
MONGODB_PORT=8080
MONGODB_DBNAME="your_dbname"
INFISICAL_MODE="prod"
INFISICAL_URL="http://example.com"
INFISICAL_CLIENTID="xxxxxxxx"
INFISICAL_CLIENTSECRET="xxxxxxxx"
INFISICAL_PROJECTID="xxxxxxxx"
```

