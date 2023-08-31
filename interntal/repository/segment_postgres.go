package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(segment ent.Segment) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", segmentsTable)
	row := r.db.QueryRow(query, segment.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SegmentPostgres) AddUserSegment(userId int, segments SegmentRequest) ([]int, error) {
	ids := make([]int, 0, 4)
	errArr := make([]int, 0)
	for _, seg := range segments.Segments {
		var id int
		segmentName := seg
		query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", segmentsTable)
		row := r.db.QueryRow(query, segmentName)
		if err := row.Scan(&id); err != nil {
			return errArr, err
		}
		query = fmt.Sprintf("INSERT INTO %s (user_id, segment_id) VALUES ($1, $2) RETURNING id", userSegmentsTable)
		row = r.db.QueryRow(query, userId, id)

		if err := row.Scan(&id); err != nil {
			return errArr, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *SegmentPostgres) DeleteUserSegment(userId int, segments SegmentRequest) error {
	for _, seg := range segments.Segments {
		segmentName := seg
		var id int
		query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", segmentsTable)
		row := r.db.QueryRow(query, segmentName)
		if err := row.Scan(&id); err != nil {
			return err
		}
		query = fmt.Sprintf("DELETE FROM %s us WHERE user_id = $1 AND segment_id = $2", userSegmentsTable)
		_, err := r.db.Exec(query, userId, id)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *SegmentPostgres) DeleteSegment(segment ent.Segment) error {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", segmentsTable)
	row := r.db.QueryRow(query, segment.Name)
	if err := row.Scan(&id); err != nil {
		return err
	}
	query = fmt.Sprintf(`DELETE FROM %s s USING %s us WHERE s.id = us.segment_id 
	AND us.segment_id = $1`, segmentsTable, userSegmentsTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf(`DELETE FROM %s us USING %s s WHERE us.segment_id = s.id AND us.segment_id = $1`, userSegmentsTable, segmentsTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}
