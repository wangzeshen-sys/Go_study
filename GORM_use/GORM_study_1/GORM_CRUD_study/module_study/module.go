package module_study

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// CreditCard 信用卡
type CreditCard struct {
	gorm.Model
	UserID uint
	Mumber string
}
type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键(外键)，tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` 设置sql类型, `unique_index`为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"`
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string 
	Code string 
}
type User struct {
	gorm.Model
	Birthday          time.Time
	Age               int
	Name              string     `gorm:"size:255"`       // string默认长度为255，使用这种tag重设
	Num               int        `gorm:"AUTO_INCREMENT"` // 自增
	CreditCard        CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作为外键)
	Email             []Email    // One-To-Many (拥有多个 - Email表的UserID作外键)
	BillingAddress    Address    // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID  sql.NullInt64
	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int
	IgnoreMe          int        `gorm:"-"` // 忽略这个字段
	Languages         []Language `gorm:"many2many:user_lanhuages"`
}



// 基本模型的定义
// type Model struct {
// 	ID        uint `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time
//   }


// 添加字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
// type User struct {
// 	gorm.Model
// 	Name string
//   }


// 只需要字段 `ID`, `CreatedAt`
// type User struct {
// 	ID        uint
// 	CreatedAt time.Time
// 	Name      string
//  }


// 表明是结构体名称的负数形式
// type User struct {} // 默认表名是`users`

// 设置User的表名为`profiles`
// func (User) TableName() string {
//   return "profiles"
// }

// func (u User) TableName() string {
//     if u.Role == "admin" {
//         return "admin_users"
//     } else {
//         return "users"
//     }
// }

// 全局禁用表名复数
// db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响


