/**
* @Description: 宗教信仰
* @Author: jinyidong
* @Date: 2021/6/8
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
	migration.Migrate.SetVersion(migration.GetFilename(fileName), faithInit)
}

func faithInit(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		list7 := []models.DictType{
			{DictId: 37, DictName: "宗教信仰", DictType: "sys_faith", Status: 2, ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "宗教信仰列表", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		list9 := []models.DictData{
			{DictCode: 240, DictSort: 0, DictLabel: "不限", DictValue: "1", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 241, DictSort: 1, DictLabel: "无神论", DictValue: "2", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 242, DictSort: 2, DictLabel: "佛教", DictValue: "3", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 243, DictSort: 3, DictLabel: "道教", DictValue: "4", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 244, DictSort: 4, DictLabel: "基督教", DictValue: "5", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 245, DictSort: 5, DictLabel: "天主教", DictValue: "6", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 246, DictSort: 6, DictLabel: "信神但未定教", DictValue: "7", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 247, DictSort: 7, DictLabel: "儒教", DictValue: "8", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 248, DictSort: 8, DictLabel: "犹太教", DictValue: "9", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 249, DictSort: 9, DictLabel: "回教", DictValue: "10", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 250, DictSort: 10, DictLabel: "伊斯兰教", DictValue: "11", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 251, DictSort: 11, DictLabel: "印度教", DictValue: "12", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 252, DictSort: 12, DictLabel: "其它", DictValue: "13", DictType: "sys_faith", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
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
