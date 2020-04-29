package db

import (
	"github.com/breathbath/erplyapi/graph"
	"github.com/breathbath/erplyapi/metrics"
	"github.com/breathbath/erplyapi/reports"
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

func (v Visits) VisitsByHour(fromTo reports.FromTo, erplyID string) (kv []graph.KeyValue, err error) {
	q := `
select count(*) 'val', DATE_FORMAT(created_at, '%d-%m-%Y %H:00') 'key'
from visit_metrics
where erply_id = ?
  And created_at BETWEEN ? and ?
group by hour(created_at), dayofyear(created_at), year(created_at)
order by created_at
`
	return v.queryReport(q, fromTo, erplyID, time.Now().UTC().Add(-24*time.Hour))
}

func (v Visits) VisitsByDay(fromTo reports.FromTo, erplyID string) (kv []graph.KeyValue, err error) {
	q := `
select count(*) 'val', DATE_FORMAT(created_at, '%d-%m-%Y') 'key'
from visit_metrics
where erply_id = ?
  And created_at BETWEEN ? and ?
group by dayofyear(created_at), year(created_at)
order by created_at
`
	return v.queryReport(q, fromTo, erplyID, time.Now().UTC().AddDate(0, 0, -7))
}

func (v Visits) VisitsByMonth(fromTo reports.FromTo, erplyID string) (kv []graph.KeyValue, err error) {
	q := `
select count(*) 'val', DATE_FORMAT(created_at, '%m-%Y') 'key'
from visit_metrics
where erply_id = ?
  And created_at BETWEEN ? and ?
group by month(created_at), year(created_at)
order by created_at
`
	return v.queryReport(q, fromTo, erplyID, time.Now().UTC().AddDate(0, -1, 0))
}

func (v Visits) VisitsByLocation(fromTo reports.FromTo, erplyID string) (kv []graph.KeyValue, err error) {
	q := `select count(*) 'val', location 'key'
from visit_metrics
where erply_id = ?
  And created_at BETWEEN ? and ?
group by location
order by created_at
`
	return v.queryReport(q, fromTo, erplyID, time.Now().UTC().AddDate(0, 0, -1))
}

func (v Visits) queryReport(sql string, fromTo reports.FromTo, erplyID string, defaultFrom time.Time) (kv []graph.KeyValue, err error) {
	from := fromTo.From.Time
	if fromTo.From.IsNull {
		from = defaultFrom
	}
	to := fromTo.To.Time
	if fromTo.To.IsNull {
		to = time.Now().UTC()
	}

	err = ScanByQuery(
		v.Db,
		&kv,
		sql,
		erplyID,
		from.Format(time.RFC3339),
		to.Format(time.RFC3339),
	)

	return
}
