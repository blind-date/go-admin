/**
* @Description: 性格字典
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
	migration.Migrate.SetVersion(migration.GetFilename(fileName), natureInit)
}

func natureInit(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		list7 := []models.DictType{
			{DictId: 30, DictName: "性格", DictType: "sys_nature", Status: 2, ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "性格列表", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}

		list9 := []models.DictData{
			{DictCode: 50, DictSort: 0, DictLabel: "动静皆宜", DictValue: "0", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 51, DictSort: 1, DictLabel: "温柔善良", DictValue: "1", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 52, DictSort: 2, DictLabel: "外向", DictValue: "2", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 53, DictSort: 3, DictLabel: "内向", DictValue: "3", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 54, DictSort: 4, DictLabel: "活泼开朗", DictValue: "4", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 55, DictSort: 5, DictLabel: "沉默寡言", DictValue: "5", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 56, DictSort: 6, DictLabel: "幽默", DictValue: "6", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 57, DictSort: 7, DictLabel: "稳重", DictValue: "7", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 58, DictSort: 8, DictLabel: "坚强", DictValue: "8", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 59, DictSort: 9, DictLabel: "成熟", DictValue: "9", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 60, DictSort: 10, DictLabel: "能说会道", DictValue: "10", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 61, DictSort: 11, DictLabel: "真诚", DictValue: "11", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 62, DictSort: 12, DictLabel: "独立", DictValue: "12", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 63, DictSort: 13, DictLabel: "依赖", DictValue: "13", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 64, DictSort: 14, DictLabel: "温柔体贴", DictValue: "14", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 65, DictSort: 15, DictLabel: "不拘小节", DictValue: "15", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 66, DictSort: 16, DictLabel: "交际广泛", DictValue: "16", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 67, DictSort: 17, DictLabel: "倔强", DictValue: "17", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 68, DictSort: 18, DictLabel: "豪爽", DictValue: "18", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 69, DictSort: 19, DictLabel: "害羞", DictValue: "19", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 70, DictSort: 20, DictLabel: "耿直", DictValue: "20", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 71, DictSort: 21, DictLabel: "乐观向上", DictValue: "21", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 72, DictSort: 22, DictLabel: "敢作敢当", DictValue: "22", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 73, DictSort: 23, DictLabel: "助人为乐", DictValue: "23", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 74, DictSort: 24, DictLabel: "老实", DictValue: "24", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 75, DictSort: 25, DictLabel: "敏感", DictValue: "25", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 76, DictSort: 26, DictLabel: "果断", DictValue: "26", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 77, DictSort: 27, DictLabel: "快言快语", DictValue: "27", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 78, DictSort: 28, DictLabel: "豪放不羁", DictValue: "28", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 79, DictSort: 29, DictLabel: "多愁善感", DictValue: "29", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 80, DictSort: 30, DictLabel: "循规蹈矩", DictValue: "30", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 81, DictSort: 31, DictLabel: "其他", DictValue: "31", DictType: "sys_nature", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
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
