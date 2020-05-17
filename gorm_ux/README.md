# gorm 的高级使用

## callbacks 
[参见 https://gorm.io/zh_CN/docs/write_plugins.html](https://gorm.io/zh_CN/docs/write_plugins.html)

函数签名 `func (scope*gorm.Scope)`

使用 callback 定制一个高级的 Model

>   Gorm 内置的 Model 默认的时间是 time.Time 

``` go
type Model struct {
 	Id int64 `json:"id"gorm:"primary_key"`
	DeletedAt *int64 `sql:"index" json:"deleted_at,omitempty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Name      string
}
```



这个时候就需要替换 gorm 默认的已经注册的回调函数了

[create](https://github.com/jinzhu/gorm/blob/master/callback_create.go), [update](https://github.com/jinzhu/gorm/blob/master/callback_update.go), [delete](https://github.com/jinzhu/gorm/blob/master/callback_delete.go)

重新编写回调函数如下：

```go
// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {

	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", gorm.NowFunc().Unix())
	}
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {

	if !scope.HasError() {
		now := gorm.NowFunc().Unix()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			updatedAtField.Set(now)

		}
	}
}

// deleteCallback used to delete data from database or set deleted_at to current time (when using with soft delete)
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedAtField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
```



替换现有的注册函数：

```go
	g.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	g.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	g.Callback().Delete().Replace("gorm:delete", deleteCallback)
```

[完整的例子](./gorm_demo.go)

## 其他的高级用法

```go
// SetMaxIdleCons 设置连接池中的最大闲置连接数。
db.DB().SetMaxIdleConns(10)

// SetMaxOpenCons 设置数据库的最大连接数量。
db.DB().SetMaxOpenConns(100)

// SetConnMaxLifetiment 设置连接的最大可复用时间。
db.DB().SetConnMaxLifetime(time.Hour)
```

