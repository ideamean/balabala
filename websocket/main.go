package main

import (
	"./ws"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"websocket/app"
	"websocket/app/actions"
)

var server ws.Server

func WSHandle(conn *websocket.Conn) {
	c := ws.NewConnection(conn)
	server.Register(c)
	defer func() { server.Unregister(c) }()
	go c.Writer()
	c.Reader(server)
}

func msgsend(w http.ResponseWriter, r *http.Request) {
	if server.GetOnlineUserNum() <= 0 {
		app.Resp_Error(w, app.Error_No_Online_User, "no online user")
		return
	}

	str, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.Resp_Error(w, app.Error_PostData_Empty, "read post body error")
		return
	}

	go func(str string) {
		server.AddBroadcast(str)
	}(string(str))

	var msg = &app.RetMsg{OnlineUserNum: server.GetOnlineUserNum(), State: 1}
	app.Resp_Json(w, msg)
}

func main() {
	server = ws.NewServer()
	go server.Run()
	http.Handle("/ws", websocket.Handler(WSHandle))
	http.HandleFunc("/msg_send", msgsend)
	http.HandleFunc("/", actions.Home)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Printf("Notice: Webserver Started, Open %s \n", app.WEBSERVER_ADDR)
	if err := http.ListenAndServe(":"+app.WEBSERVER_PORT, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
