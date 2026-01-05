package repository

import (
	"log/slog"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
)

func (r *Repository) FindRefreshTokenById(id gocql.UUID) (string /*refreshToken*/, error) {
	var rt string
	err := r.session.Query(
		"SELECT refresh_token from member_by_id WHERE id = ?",
		id,
	).Scan(&rt)
	if err != nil {
		slog.Info("fail to get refresh token",
			"err", err,
		)
		return "", err
	}
	return rt, nil
}

func (r *Repository) SaveRefreshTokenById(id gocql.UUID, rt string) error {
	err := r.session.Query(
		"UPDATE member_by_id USING TTL ? SET refresh_token = ? WHERE id = ?",
		r.rtExp, rt, id,
	).Exec()
	if err != nil {
		slog.Error("fail to save refresh token",
			"err", err,
		)
		return err
	}
	return nil
}
