#！ /bin/bash
//在mac下是这里，其他不一定

# Build web UI
cd ~/work/src/github.com/avenssi/video_server/web
go install
cp ~/work/bin/web ~/work/bin/video_server_web_ui/web
cp -R ~/woker/src/github.com/avenssi/video_server/templates ~/work/bin/video_server_web_ui/web
