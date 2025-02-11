package postgres

import (
	"fmt"
	"strings"

	"github.com/livensmi1e/tiny-ide/store"
)

func (p *_postgres) CreateSubmission(create *store.Submission) (*store.Submission, error) {
	fields := []string{"id", "language_id"}
	args := []interface{}{create.ID, create.LanguageID}
	placeholder := []string{"$1", "$2"}
	stmt := "INSERT INTO submission (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholder, ", ") + ") RETURNING id, language_id"
	if err := p.db.QueryRow(stmt, args...).Scan(
		&create.ID,
		&create.LanguageID,
	); err != nil {
		return nil, err
	}
	return create, nil
}

func (p *_postgres) UpdateSubmission(update *store.UpdateSubmission) (*store.Submission, error) {
	set, args := []string{}, []interface{}{}
	if v := update.Status; v != nil {
		set, args = append(set, "status = $1"), append(args, *v)
	}
	if v := update.Stdout; v != nil {
		set, args = append(set, "stdout = $2"), append(args, *v)
	}
	if v := update.Stderr; v != nil {
		set, args = append(set, "stderr = $3"), append(args, *v)
	}
	if v := update.Time; v != nil {
		set, args = append(set, "time = $4"), append(args, *v)
	}
	if v := update.Memory; v != nil {
		set, args = append(set, "memory = $5"), append(args, *v)
	}
	args = append(args, update.ID)
	stmt := `
			UPDATE submission
			SET ` + strings.Join(set, ", ") + `
			WHERE id = $6
			RETURNING id, language_id, status, stdout, stderr, time, memory
			`
	submission := &store.Submission{}
	if err := p.db.QueryRow(stmt, args...).Scan(
		&submission.ID,
		&submission.LanguageID,
		&submission.Status,
		&submission.Stdout,
		&submission.Stderr,
		&submission.Time,
		&submission.Memory,
	); err != nil {
		return nil, err
	}
	return submission, nil
}

func (p *_postgres) ListSubmissions(find *store.FindSubmission) ([]*store.Submission, error) {
	where, args := []string{}, []interface{}{}
	if v := find.ID; v != nil {
		where, args = append(where, "id = $1"), append(args, *v)
	}
	stmt := `SELECT
					id,
					language_id,
					status,
					stdout,
					stderr,
					time,
					memory
			FROM submission`
	if len(where) > 0 {
		stmt += " WHERE " + strings.Join(where, " AND ")
	}
	if v := find.Limit; v != nil {
		stmt += fmt.Sprintf(" LIMIT %d", *v)
	}
	rows, err := p.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]*store.Submission, 0)
	for rows.Next() {
		var submission store.Submission
		if err := rows.Scan(
			&submission.ID,
			&submission.LanguageID,
			&submission.Status,
			&submission.Stdout,
			&submission.Stderr,
			&submission.Time,
			&submission.Memory,
		); err != nil {
			return nil, err
		}
		list = append(list, &submission)
	}
	return list, nil
}
