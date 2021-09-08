/**
* @Description: 是否喝酒
* @Author: jinyidong
* @Date: 2021/6/8
* @Version V1.0
 */
package version_local

import (
	"go-admin/cmd/migrate/migration"
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"
	"gorm.io/gorm"
	"runtime"
	"time"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), drinkInit)
}

func drinkInit(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		list7 := []models.DictType{
			{DictId: 47, DictName: "喝酒", DictType: "sys_drink", Status: 2, ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "是否喝酒列表", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		list9 := []models.DictData{
			{DictCode: 660, DictSort: 0, DictLabel: "不限", DictValue: "1", DictType: "sys_drink", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 661, DictSort: 1, DictLabel: "不喝酒", DictValue: "2", DictType: "sys_drink", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 662, DictSort: 2, DictLabel: "社交需要喝", DictValue: "3", DictType: "sys_drink", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 663, DictSort: 3, DictLabel: "兴致时小酌", DictValue: "4", DictType: "sys_drink", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 664, DictSort: 4, DictLabel: "酒不离口", DictValue: "5", DictType: "sys_drink", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		err := tx.Create(list7).Error
		if err != nil {
			return err
		}

		err = tx.Create(list9).Error
		if err != nil {
			return err
		}

		if err := models.InitDb(tx); err != nil {

		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
