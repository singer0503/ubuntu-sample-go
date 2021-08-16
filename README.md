# ubuntu-sample-go

```
sudo su -

sudo setcap CAP_NET_BIND_SERVICE=+eip /home/singer0503/ubuntu-sample-go/06-https-sample/main.go

nohup sudo go run main.go &

sudo lsof -i -P -n | grep "LISTEN"
```
