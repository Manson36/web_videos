package dbops

import (
	"database/sql"
	"github.com/web_videos/api/defs"
	"strconv"
)

func InserSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare(
		"INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(sid, ttlStr, uname)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare(
		"SELECT TTL login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		return nil, err
	}

	defer stmtOut.Close()

	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	ttlInt, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return nil, err
	}
	ss.TTL = ttlInt
	ss.UserName = uname

	return ss, nil
}
