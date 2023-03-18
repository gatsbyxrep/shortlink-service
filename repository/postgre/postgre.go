package postgre

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Implements operations with Postgres
// Probably statements should be prepared, but https://github.com/jackc/pgx/issues/791 says not
type Postgre struct {
	conn *pgx.Conn
}

func Init(hostName, port, dbName, username, password string) (Postgre, error) {
	// Probably need to be changed if runs in docker
	//url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", hostName, port, username, password, dbName)
	// postgres://postgres:postgrespw/postgres@postgreshost.docker.internal:32771"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, hostName, port, dbName)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return Postgre{}, err
	}
	return Postgre{
		conn: conn,
	}, nil
}
func (p Postgre) PushOriginalAndShort(original, short string) error {
	_, err := p.conn.Exec(context.Background(), "insert into links(original_link, short_link) values ($1, $2)", original, short)
	return err
}

func (p Postgre) GetByShortLink(short string) (string, error) {
	originalLink := ""
	err := p.conn.QueryRow(context.Background(), "select original_link from links where short_link=$1", short).Scan(&originalLink)
	if err != nil {
		return "", err
	}
	return originalLink, nil
}
func (p Postgre) Close() {
	p.conn.Close(context.Background())
}
