package items

import "github.com/breathbath/erplyapi/migrations"

type Migration20200427092907 struct {
	migrations.BasicItem
}

func (mm Migration20200427092907) GetId() string {
	return "Migration20200427092907"
}

//todo add query items
func (mm Migration20200427092907) GetQueries() []migrations.Query {
	return []migrations.Query{
		migrations.NewQuery(`
CREATE TABLE visit_metrics (
  id int unsigned NOT NULL AUTO_INCREMENT,
  created_at datetime NOT NULL,
  location varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  device_hash varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  erply_id varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (id),
  UNIQUE KEY uid (created_at,location,device_hash,erply_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`),
	}
}
