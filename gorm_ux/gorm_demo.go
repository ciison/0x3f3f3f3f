package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	g *gorm.DB
)

type DemoX struct {
	Id        int64  `json:"id"gorm:"primary_key"`
	DeletedAt *int64 `sql:"index" json:"deleted_at,omitempty"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	Name      string
}

func init() {
	var err error
	g, err = gorm.Open("mysql", "root:12345@tcp/demo")
	if err != nil {
		panic(err)
	}
	g.LogMode(true)

	g.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	g.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	g.Callback().Delete().Replace("gorm:delete", deleteCallback)

	g.AutoMigrate(&DemoX{})

}

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

func (c *DemoX) BeforeUpdate() (err error) {
	//if c.Id == 0 {
	//	return errors.New("id could not be null")
	//}
	return
}
func main() {
	var err error
	g.Model(&DemoX{}).Create(&DemoX{Name: time.Now().String()})

	var dx = DemoX{
		Id: 2,
		//DeletedAt: nil,
		CreatedAt: 0,
		Name:      "cao",
	}
	d, err := json.Marshal(&dx)
	fmt.Println(string(d))
	err = g.Model(&dx).Update("name", "草拟吗").Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("dx:%#v\n", dx)
	err = g.Delete(&dx).Error
	fmt.Println(err)
	dx.Id = 4
	err = g.Model(&dx).Take(&dx).Error
	fmt.Println(err)
	fmt.Printf("%#v\n", dx)
	err = g.Model(&dx).Delete(&dx).Error
	fmt.Println(err)
	fmt.Println(dx)
	var dxs []DemoX
	g.Model(&DemoX{}).Find(&dxs)
	fmt.Println(dxs)
}
