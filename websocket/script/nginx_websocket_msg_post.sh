#!/bin/sh
tail -f nginx.access.log|
while read msg
do
    result=$(echo $msg|sed -f urlencode.sed)
    echo $result
    curl -s -d "$result" "http://127.0.0.1:8003/msg_send" >> /dev/null
done
