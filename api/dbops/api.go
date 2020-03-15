package dbops

import (
	"database/sql"
	"github.com/web_videos/api/defs"
	"github.com/web_videos/api/utils"
	"log"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetuserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Println(err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create UUID
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	//在VideoInfo中会有一个createTime是写入库的时间，在这里我们是生成displayCtime是在写入库之前调用函数生成一个时间，用于排序
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	//如果在""之间要进行换行，就要将""换成``才不会出错
	stmtIns, err := dbConn.Prepare(
		"INSERT INTO video_info (id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, err
}

func GetVideo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE vid = ?|")
	if err != nil {
		return nil, err
	}

	var aid int
	var name string
	var dct string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	defer stmtOut.Close()
	return res, nil
}

func DeleteVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmtDel.Close()

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	return nil
}

func AddNewComment(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?,?)")
	if err != nil {
		return err
	}

	defer stmtIns.Close()

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	//这里需要：users join comments ->获取users表中login_name
	stmtOut, err := dbConn.Prepare(
		`SELETE comments.id, users.login_name, comments.content FROM comments INNER JOIN users  
	ON comments.author_id = users.id WHERE comments.video_id = ? AND comment.time > FROM_UNIXTIME(?) 
	AND comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		return nil, err
	}

	defer stmtOut.Close()
	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err = rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, AuthorName: name, Content: content}
		res = append(res, c)
	}

	return res, nil
}
