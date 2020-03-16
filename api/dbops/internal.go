package dbops

import (
	"database/sql"
	"github.com/web_videos/api/defs"
	"log"
	"strconv"
	"sync"
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

func RetrieveAllSession() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlStr string
		var login_name string

		if err := rows.Scan(&id, ttlStr, login_name); err != nil {
			log.Println("Retrieve sessions error, errmsg:", err.Error())
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlStr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{UserName: login_name, TTL: ttl}
			m.Store(id, ss)
		}
	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = sid")
	if err != nil {
		log.Println("Delete session error, errmsg:", err.Error())
		return err
	}
	defer stmtDel.Close()

	_, err = stmtDel.Exec(sid)
	if err != nil {
		log.Println("Delete session error, errmsg:", err.Error())
		return err
	}

	return nil
}
