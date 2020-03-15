package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
流程：
init(dblogin,truncate tables) ->run tests ->clear data(truncate tables)
为了保证每次test出现代码的干扰，我们会先初始化数据 再run test 再清理数据的过程
*/

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessioins")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("awen", "132456")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetuserCredential("awen")
	if pwd != "123456" || err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}

}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("awen", "123456")
	if err != nil {
		t.Errorf("Error of DELETEUser:%v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetuserCredential("awen")
	if pwd != "123456" || err != nil {
		t.Errorf("Error of RegetUser:%v", err)
	}
}

var tempvid string

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddUser)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(8888, "谷米")
	if err != nil {
		t.Errorf("Error of AddVideoInfo:%v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo:%v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo:%v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideo(tempvid)
	if err != nil && vi != nil {
		t.Errorf("Error of RegetVideoInfo:%v", err)
	}
}

func TestCommentFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComment", testAddComment)
	t.Run("ListComments", testListComments)
}

func testAddComment(t *testing.T) {
	vid := "12456"
	aid := 1
	content := "I like this video"

	err := AddNewComment(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddNewComment:%v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12456"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/100000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments:%v", err)
	}

	for i, v := range res {
		fmt.Printf("comment:%d:%v \n", i, v)
	}
}
