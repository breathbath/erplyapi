package db

import (
	"github.com/breathbath/erplyapi/graph"
	"github.com/breathbath/erplyapi/metrics"
	baseDb "github.com/breathbath/go_utils/utils/sqlDb"
	log "github.com/sirupsen/logrus"
	"time"
)

type Visits struct {
	Db *baseDb.DbGateway
}

func (v Visits) Add(metric metrics.VisitMetric) error {
	log.Infof("Will add new metric: %v+", metric)
	res, err := v.Db.Exec(
		`INSERT INTO visit_metrics
(created_at, location, device_hash, erply_id)
VALUES (?,?,?,?)`,
		time.Now().UTC(),
		metric.Location,
		metric.DeviceHash,
		metric.ErplyID,
	)

	if err != nil {
		return err
	}

	insertId, err := res.LastInsertId()
	log.Infof("Added a new metric to visit_metrics with id: %d", insertId)

	return err
}

func (v Visits) VisitsByHour(fromTo graph.FromTo, erplyID string) (kv []graph.KeyValue, err error) {
	from := fromTo.From.Time
	if fromTo.From.IsNull {
		from = time.Now().UTC().Add(-10 * time.Hour)
	}
	to := fromTo.To.Time
	if fromTo.To.IsNull {
		to = time.Now().UTC()
	}

	err = v.Db.FindByQueryFlex(
		&kv,
		`
select count(*) key, DATE_FORMAT(created_at, '%d-%m-%Y %H:00') val
from visit_metrics
where erply_id = ?
  And created_at BETWEEN ? and ?
group by hour(created_at), day(created_at);
`,
		erplyID,
		from,
		to,
	)

	return
}
