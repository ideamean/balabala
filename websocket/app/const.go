package app

const (
	WEBSERVER_PORT string = "8003"
	WEBSERVER_ADDR string = "http://127.0.0.1:" + WEBSERVER_PORT
	WEBSOCKET_ADDR string = "ws://127.0.0.1:" + WEBSERVER_PORT + "/ws"
)

const (
	Error_No_Online_User = 10000
	Error_PostData_Empty = 10001
)
