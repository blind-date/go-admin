/**
* @Description: 自定义系统配置
* @Author: jinyidong
* @Date: 2021/6/8
* @Version V1.0
 */
package version_local

import (
	"runtime"

	"gorm.io/gorm"

	"go-admin/cmd/migrate/migration"
	common "go-admin/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), sysConfig)
}

// 自定义系统配置
func sysConfig(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		err := tx.Exec("INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, '阿里云短信签名', 'ali_sms_code_sign', '', 'N', 1, '', 1, 0, '2021-07-05 08:52:06.067', '2021-07-05 08:52:06.067', NULL);").Error
		if err != nil {
			return err
		}

		err = tx.Exec("INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, '阿里云短信AccessKey', 'ali_sms_code_access_key', '', 'N', 1, '', 1, 0, '2021-07-05 08:52:06.067', '2021-07-05 08:52:06.067', NULL);").Error
		if err != nil {
			return err
		}

		err = tx.Exec("INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, '阿里云短信AccessSecret', 'ali_sms_code_access_secret', '', 'N', 1, '', 1, 0, '2021-07-05 08:52:06.067', '2021-07-05 08:52:06.067', NULL);").Error
		if err != nil {
			return err
		}

		err = tx.Exec("INSERT INTO `sys_config` (`id`, `config_name`, `config_key`, `config_value`, `config_type`, `is_frontend`, `remark`, `create_by`, `update_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, '阿里云短信模板ID', 'ali_sms_code_template', '', 'N', 1, '', 1, 0, '2021-07-05 08:52:06.067', '2021-07-05 08:52:06.067', NULL);").Error
		if err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error
	})
}
