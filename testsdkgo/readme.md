# API Client of SDKClient

## Installation

Download the package using:
```
go get github.com/brianrugwizam/petstore2/testsdkgo
```

## Usage

An example for a client that connects to the Software Defined Networking API (SDN) behind the Arch (with an Arch
Workspace as the BasePath), ratelimiting, metrics and CCP to retrieve the NPA password:

```go
package main

import (
	"github.com/brianrugwizam/petstore2/testsdkgo"
	"github.com/prometheus/client_golang/"
	"github.com/prometheus/client_golang/prometheus"
	"dev.azure.com/INGCDaaS/IngOne/_git/P19673-go-modules.git/cyberarkaimccp"
	"dev.azure.com/INGCDaaS/IngOne/_git/P19673-go-modules.git/cyberarkaimccp/types"
)

func main() {
	// create the ccp client to use for making requests to CyberArk CCP
	ccp, err := cyberarkaimccp.NewClient(
		cyberarkaimccp.SetHost("https://ccp.pwv.iam.ing.net"),
		cyberarkaimccp.SetHttpClient([]byte("publiccert"), []byte("publickey"), nil),
    )
	if err != nil {
		panic(err)
	}

	client, err := testsdkgo.NewClient("sdn.ing.net",
		testsdkgo.WithBasePath("/sdn-cdt"),
		testsdkgo.WithRateLimiting(10, 30),
		testsdkgo.WithMetrics(prometheus.DefaultRegisterer),
		testsdkgo.WithCCP(ccp, &types.PasswordParams{
			AppID:  "someapp",
			Object: "somenpa@ad.ing.net",
			Safe:   "safename",
			Folder: "Root",
		}, testsdkgo.BasicAuth))
	if err != nil {
		panic(err)
	}

	// Make calls
	success, failure, response, err := client.PetsAPI.CreatePet(context.Background(), CreatePetBody{})
}
```


