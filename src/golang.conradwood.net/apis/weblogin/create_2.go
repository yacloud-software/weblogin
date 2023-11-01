// client create: TrackerServiceClient
/*
  Created by /home/cnw/devel/go/yatools/src/golang.yacloud.eu/yatools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.conradwood.net/apis/weblogin/weblogin.proto
   gopackage : golang.conradwood.net/apis/weblogin
   importname: ai_1
   clientfunc: GetTrackerService
   serverfunc: NewTrackerService
   lookupfunc: TrackerServiceLookupID
   varname   : client_TrackerServiceClient_1
   clientname: TrackerServiceClient
   servername: TrackerServiceServer
   gsvcname  : weblogin.TrackerService
   lockname  : lock_TrackerServiceClient_1
   activename: active_TrackerServiceClient_1
*/

package weblogin

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_TrackerServiceClient_1 sync.Mutex
  client_TrackerServiceClient_1 TrackerServiceClient
)

func GetTrackerClient() TrackerServiceClient { 
    if client_TrackerServiceClient_1 != nil {
        return client_TrackerServiceClient_1
    }

    lock_TrackerServiceClient_1.Lock() 
    if client_TrackerServiceClient_1 != nil {
       lock_TrackerServiceClient_1.Unlock()
       return client_TrackerServiceClient_1
    }

    client_TrackerServiceClient_1 = NewTrackerServiceClient(client.Connect(TrackerServiceLookupID()))
    lock_TrackerServiceClient_1.Unlock()
    return client_TrackerServiceClient_1
}

func GetTrackerServiceClient() TrackerServiceClient { 
    if client_TrackerServiceClient_1 != nil {
        return client_TrackerServiceClient_1
    }

    lock_TrackerServiceClient_1.Lock() 
    if client_TrackerServiceClient_1 != nil {
       lock_TrackerServiceClient_1.Unlock()
       return client_TrackerServiceClient_1
    }

    client_TrackerServiceClient_1 = NewTrackerServiceClient(client.Connect(TrackerServiceLookupID()))
    lock_TrackerServiceClient_1.Unlock()
    return client_TrackerServiceClient_1
}

func TrackerServiceLookupID() string { return "weblogin.TrackerService" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.

func init() {
   client.RegisterDependency("weblogin.TrackerService")
   AddService("weblogin.TrackerService")
}
