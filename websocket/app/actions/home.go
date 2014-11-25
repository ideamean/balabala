package actions

import (
	"html/template"
	"net/http"
	"websocket/app"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/home.tpl")
	if err != nil {
		panic(err)
	}
	data := app.Tpl_Model_Home{
		WebSocket_Addr: app.WEBSOCKET_ADDR,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
