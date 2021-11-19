.PHONY:	server client
server:
	cd server && go install ${LDFLAGS} weblogin-server.go `ls -1 *.go|grep -v weblogin-server.go`
client:
	echo no client yet
	cd client ; go install ${LDFLAGS} weblogin-client.go
all:	client server


