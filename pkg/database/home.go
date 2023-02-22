package database

import (
	"context"
	"ispick-project-21022023/pkg/model"
)

func (db *Db) AppVersion(ctx context.Context) (model.AppVersion, error) {
	var appVersion model.AppVersion
	row_1 := db.dbconn.QueryRow("SELECT id, COALESCE(app_release_no) as app_release_no, COALESCE(app_release_description) as app_release_description, COALESCE(ref_id) as ref_id, COALESCE(ref_name) as ref_name, COALESCE(created_at) as created_at, COALESCE(updated_at) as updated_at, COALESCE(deleted_at, '') as deleted_at FROM app_releases ORDER BY created_at DESC LIMIT 1")

	err := row_1.Scan(
		&appVersion.Id,
		&appVersion.App_release_no,
		&appVersion.App_release_description,
		&appVersion.Ref_id,
		&appVersion.Ref_name,
		&appVersion.Created_at,
		&appVersion.Updated_at,
		&appVersion.Deleted_at,
	)
	if err != nil {
		return appVersion, err
	}
	return appVersion, nil
}
