// client create: SCModCommsServiceClient
/*
  Created by /srv/home/cnw/devel/go/go-tools/src/golang.conradwood.net/gotools/protoc-gen-cnw/protoc-gen-cnw.go
*/

/* geninfo:
   filename  : protos/golang.singingcat.net/apis/scmodcomms/scmodcomms.proto
   gopackage : golang.singingcat.net/apis/scmodcomms
   importname: ai_0
   clientfunc: GetSCModCommsService
   serverfunc: NewSCModCommsService
   lookupfunc: SCModCommsServiceLookupID
   varname   : client_SCModCommsServiceClient_0
   clientname: SCModCommsServiceClient
   servername: SCModCommsServiceServer
   gscvname  : scmodcomms.SCModCommsService
   lockname  : lock_SCModCommsServiceClient_0
   activename: active_SCModCommsServiceClient_0
*/

package scmodcomms

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_SCModCommsServiceClient_0 sync.Mutex
  client_SCModCommsServiceClient_0 SCModCommsServiceClient
)

func GetSCModCommsClient() SCModCommsServiceClient { 
    if client_SCModCommsServiceClient_0 != nil {
        return client_SCModCommsServiceClient_0
    }

    lock_SCModCommsServiceClient_0.Lock() 
    if client_SCModCommsServiceClient_0 != nil {
       lock_SCModCommsServiceClient_0.Unlock()
       return client_SCModCommsServiceClient_0
    }

    client_SCModCommsServiceClient_0 = NewSCModCommsServiceClient(client.Connect(SCModCommsServiceLookupID()))
    lock_SCModCommsServiceClient_0.Unlock()
    return client_SCModCommsServiceClient_0
}

func GetSCModCommsServiceClient() SCModCommsServiceClient { 
    if client_SCModCommsServiceClient_0 != nil {
        return client_SCModCommsServiceClient_0
    }

    lock_SCModCommsServiceClient_0.Lock() 
    if client_SCModCommsServiceClient_0 != nil {
       lock_SCModCommsServiceClient_0.Unlock()
       return client_SCModCommsServiceClient_0
    }

    client_SCModCommsServiceClient_0 = NewSCModCommsServiceClient(client.Connect(SCModCommsServiceLookupID()))
    lock_SCModCommsServiceClient_0.Unlock()
    return client_SCModCommsServiceClient_0
}

func SCModCommsServiceLookupID() string { return "scmodcomms.SCModCommsService" } // returns the ID suitable for lookup in the registry. treat as opaque, subject to change.
