# MOZ Covid19 API Golang SDK

### Example usage:

```go
package main

import (
	"fmt"
	"github.com/DimaSorokin/medzakupivli-client/client"
	"log"
	"time"
)

func main() {
	c := client.New()

	hospitalSamples, err := c.GetProtective(client.CovidArgs{
		Edrpou:     "<number edrpou>",
		ReportData: time.Now(),
	})
	if err != nil {
		log.Fatalf("error getting protective list: %v", err)
	}
	
	log.Println("We got the response:", hospitalSamples)

}
```
