/**
* @Description: 对配偶不能接受的缺点
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
	migration.Migrate.SetVersion(migration.GetFilename(fileName), spouseFocusShortcomingInit)
}
func spouseFocusShortcomingInit(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		list7 := []models.DictType{
			{DictId: 77, DictName: "对配偶不能接受的缺点", DictType: "sys_spouse_focus_shortcoming", Status: 2, ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "对配偶不能接受的缺点列表", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		list9 := []models.DictData{
			{DictCode: 1140, DictSort: 1, DictLabel: " 浓妆艳抹，穿着暴露", DictValue: "1", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1141, DictSort: 2, DictLabel: " 思想极端，偏激", DictValue: "2", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1142, DictSort: 3, DictLabel: " 大男子主义，大女人主义", DictValue: "3", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1143, DictSort: 4, DictLabel: " 有家族遗传病史", DictValue: "4", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1144, DictSort: 5, DictLabel: " 有体味", DictValue: "5", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1145, DictSort: 6, DictLabel: " 行为不端", DictValue: "6", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1146, DictSort: 7, DictLabel: " 沉迷网游", DictValue: "7", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1147, DictSort: 8, DictLabel: " 不修边幅，邋遢不注重个人卫生", DictValue: "8", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1148, DictSort: 9, DictLabel: " 情史丰富", DictValue: "9", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1149, DictSort: 10, DictLabel: " 花心，劈腿", DictValue: "10", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1150, DictSort: 11, DictLabel: " 吸毒，赌博", DictValue: "11", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1151, DictSort: 12, DictLabel: " 恶意撒谎", DictValue: "12", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1152, DictSort: 13, DictLabel: " 拜金，过分物质", DictValue: "13", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1153, DictSort: 14, DictLabel: " 势利，斤斤计较", DictValue: "14", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1154, DictSort: 15, DictLabel: " 夜生活丰富", DictValue: "15", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1155, DictSort: 16, DictLabel: " 酗酒，酗烟", DictValue: "16", DictType: "sys_spouse_focus_shortcoming", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
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
