// client create: WebloginClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/weblogin/weblogin.proto
   gopackage : golang.conradwood.net/apis/weblogin
   importname: ai_0
   clientfunc: GetWeblogin
   serverfunc: NewWeblogin
   lookupfunc: WebloginLookupID
   varname   : client_WebloginClient_0
   clientname: WebloginClient
   servername: WebloginServer
   gsvcname  : weblogin.Weblogin
   lockname  : lock_WebloginClient_0
   activename: active_WebloginClient_0
*/

package weblogin

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_WebloginClient_0 sync.Mutex
  client_WebloginClient_0 WebloginClient
)

func GetWebloginClient() WebloginClient { 
    if client_WebloginClient_0 != nil {
        return client_WebloginClient_0
    }

    lock_WebloginClient_0.Lock() 
    if client_WebloginClient_0 != nil {
       lock_WebloginClient_0.Unlock()
       return client_WebloginClient_0
    }

    client_WebloginClient_0 = NewWebloginClient(client.Connect(WebloginLookupID()))
    lock_WebloginClient_0.Unlock()
    return client_WebloginClient_0
}

func WebloginLookupID() string { return "weblogin.Weblogin" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("weblogin.Weblogin")
}
