/**
* @Description: 择偶条件权重
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
	migration.Migrate.SetVersion(migration.GetFilename(fileName), memberSpouseIntegrateWeight)
}
func memberSpouseIntegrateWeight(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		list7 := []models.DictType{
			{DictId: 92, DictName: "会员择偶条件字段权重", DictType: "sys_member_spouse_column_weight", Status: 2, ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "会员字段权重", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}
		list9 := []models.DictData{
			{DictCode: 1400, DictSort: 0, DictLabel: "refuse_industries", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1401, DictSort: 1, DictLabel: "height_from", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1402, DictSort: 2, DictLabel: "weight_from", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1403, DictSort: 3, DictLabel: "nature", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1404, DictSort: 4, DictLabel: "post", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1405, DictSort: 5, DictLabel: "appearance", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1406, DictSort: 6, DictLabel: "age_from", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1407, DictSort: 7, DictLabel: "focus_requirement", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1408, DictSort: 8, DictLabel: "company", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1409, DictSort: 9, DictLabel: "focus_shortcoming", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1410, DictSort: 10, DictLabel: "house_requirement", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1411, DictSort: 11, DictLabel: "education", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1412, DictSort: 12, DictLabel: "salary", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1413, DictSort: 13, DictLabel: "marry_situation", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1414, DictSort: 14, DictLabel: "car_requirement", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1415, DictSort: 15, DictLabel: "has_child_requirement", DictValue: "1", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
			{DictCode: 1416, DictSort: 16, DictLabel: "area_requirement", DictValue: "2", DictType: "sys_member_spouse_column_weight", CssClass: "", ListClass: "", IsDefault: "", Status: 2, Default: "", ControlBy: models.ControlBy{CreateBy: 1, UpdateBy: 1}, Remark: "", ModelTime: models.ModelTime{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
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
