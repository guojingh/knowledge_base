# Gorm学习笔记

官方网址：https://gorm.io/zh_CN/docs/index.html

## 模型定义

### 参考模型

```go
// user结构模型
type User struct {
	ID           uint           `gorm:"primaryKey;comment:主键ID"` // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // 时间指针可以为空
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}
```

### 字段级权限控制

```go
type User struct {
	Name string `gorm:"<-:create"` // 允许读和创建
	Name string `gorm:"<-:update"` // 允许读和更新
	Name string `gorm:"<-"`        // 允许读和写（创建和更新）
	Name string `gorm:"<-:false"`  // 允许读，禁止写
	Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
	Name string `gorm:"->;<-:create"` // 允许读和写
	Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
	Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段
	Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段
}
```

### 时间字段使用

```go
type User struct {
    CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
    UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
    Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳纳秒数填充更新时间
    Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
    Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}
```

### 字段标签使用说明

https://gorm.io/zh_CN/docs/models.html

## 连接数据库

`GORM` 官方支持的数据库类型有：`MySQL`, `PostgreSQL`, `SQLite`, `SQL Server` 和 `TiDB`

```go
var DB *gorm.DB

// 初始化Gorm连接
func InitUser() {
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
       panic(err)
    }
    DB = db
}

// 初始化数据库表
func MysqlTables(db *gorm.DB) {
    if err := db.AutoMigrate(User{}); err != nil {
       fmt.Printf("gorm auto migrate failed, err:%v\n", err)
       os.Exit(0)
    }
    fmt.Printf("gorm auto migrate success...")
}
```











































