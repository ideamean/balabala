var ws;
$(function(){
    ws = new WebSocket(wsServer);
    ws.onopen = function(evt) {
        appendItem('Connected to WebSocket server.');
    };
    ws.onclose = function(evt) {
        appendItem('Disconnected');
    };
    ws.onmessage = function(evt) {
        if(stop_recv){
            return;
        } 
        var str = evt.data;
        var arr = str.match(/(".*?")/g);
        if(arr != null){
            var request = arr[0].replace(/(^")|("$)/, '');
            var method = request.split(' ')[0], url = request.split(' ')[1], http = request.split(' ')[2];
            str = str.replace(method, '<span class="method">' + method + '</span>')
                     .replace(url, '<span class="url">' + url + '</span>')
                     .replace(http, '<span class="http">' + http + '</span>');
        }
        appendItem(str);
    };
    ws.onerror = function(evt) {
        console.log('Error occured: ' + evt.data);
    };
    
    $('#StopRecv').bind('click',function(){
        var text = $(this).text();
        if(text == '停止'){
            $(this).text('开始');
            $(this).removeClass('btn-stop').addClass('btn-start');
            stop_recv = 1;
        }else{
            $(this).text('停止');
            $(this).removeClass('btn-start').addClass('btn-stop');
            stop_recv = 0;
        }
    });
    
    $('#AddSplit').bind('click',function(){
        addSplit();
    });
    
    $('#Empty').bind('click',function(){
        $('#Msg').empty();
    });
});

function addSplit(){
    $('#Msg').prepend('<li class="log-split">&nbsp;</li>');
}
function appendItem(msg){
    $('#Msg').prepend('<li class="log-item">' + msg + '</li>');
    var n = $('#Msg').find('li').size();
    if( n > 10){
        $('#Msg').find('li:gt(200)').remove();
    }
}