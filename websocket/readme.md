websocket 应用
================

功能：调用/msg_send接口将提交的数据实时展示到页面中
------------------------------------------------------

应用：页面中实时显示nginx access 日志

测试命令： curl -d 'test message !' '127.0.0.1:8003/msg_send'

* 1: 运行服务
* 2: 浏览器打开http://127.0.0.1:8003/
* 3: 修改script/nginx_websocket_msg_post.sh文件中 nginx.access.log 的路径为真实路径
* 4: 执行
   nohup ./nginx_websocket_msg_post.sh &

![screen](https://github.com/ideamean/balabala/blob/master/websocket/snapshot/screen.jpg)
