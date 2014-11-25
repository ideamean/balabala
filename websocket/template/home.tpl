<!DOCTYPE HTML>
<html>
<link rel="stylesheet" href="/public/css/main.css">
<head>
    <meta charset="utf-8">
    <script src="http://upcdn.b0.upaiyun.com/libs/jquery/jquery-2.0.3.min.js"></script>
    <script>var wsServer = "{{.WebSocket_Addr}}", stop_recv = 0;</script>
    <script src="/public/js/main.js"></script>
</head>

<body style="margin:0;padding:0;">
    <ul id="Msg" class="log-item"></ul>
    <div id="NavBar">
        <button id="StopRecv" class="btn-ctrl btn-stop">停止</button>
        <button id="AddSplit" class="btn-ctrl btn-split">分隔线</button>
        <button id="Empty" class="btn-ctrl btn-empty">清空</button>
    </div>
</body>

</html>