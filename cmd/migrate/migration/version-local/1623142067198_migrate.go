/**
* @Description: 婚姻状况
* @Author: jinyidong
* @Date: 2021/6/10
* @Version V1.0
 */
package version_local

import (
	"go-admin/cmd/migrate/migration"
	"go-admin/cmd/migrate/migration/models"
	common "go-admin/common/models"
	"runtime"
	"time"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), marrySituationInit)
}
func marrySituationInit(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		list7 := []models.DictType{
			{DictId: 59, DictName: "婚姻情况", DictType: "sys_marry_situation", Status: 2, ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "婚姻情况列表", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		list9 := []models.DictData{
			{DictCode: 1850, DictSort: 0, DictLabel: "未婚", DictValue: "1", DictType: "sys_marry_situation", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1851, DictSort: 1, DictLabel: "离异", DictValue: "2", DictType: "sys_marry_situation", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1852, DictSort: 2, DictLabel: "丧偶", DictValue: "3", DictType: "sys_marry_situation", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1853, DictSort: 3, DictLabel: "短婚未育", DictValue: "4", DictType: "sys_marry_situation", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
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
