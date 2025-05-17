Top Notch: https://systemdesign.one/url-shortening-system-design/#summary

### Requirements:

1.URL shortening: given a long URL => return a much shorter URL
2.URL redirecting: given a shorter URL => redirect to the original URL
3.High availability, scalability, and fault tolerance considerations

### Capacity Estimation:
1. Write operation: 100 million URLs are generated per day.
2. Write operation per second: 100 million / 24 /3600 = 1160
3. Read operation: Assuming ratio of read operation to write operation is 10:1, read operation per second: 1160 * 10 = 11,600
4. Assuming the URL shortener service will run for 10 years, this means we must support 100 million * 365 * 10 = 365 billion records.
5. Assume average URL length is 100.
6. Storage requirement over 10 years: 365 billion * 100 bytes * 10 years = 365 TB

### Designing

#### Apis

POST api/v1/data/shorten
- request parameter: {longUrl: longURLString}
- return shortURL

GET api/v1/shortUrl
- Return longURL for HTTP redirection

#### Url redirection
![alt_text](./images/img_1.png)

#### Url shortening 
![alt_text](./images/img_2.png)

#### Data model
![alt_text](./images/img_3.png)

#### Hash optimisation
- Since in short url according to our estimations , We don't require very long hashed string and here we can save our storage.
- This is made up of alphanumeric, total -> 62 [0-9, a-z , A-Z].
- So 7 length of short url would be enough
  ![alt_text](./images/img_4.png)


1. Hash + collision resolution:
   Inshort get the long hashed value from function and take only first 7 letter. This might cause collision. So start adding letter one by one more and check if its not already exist. Cons: Call db every time or cache . Not much efficient. Optimisation use bloom filter.
2. Base 64 conversion:
   ![alt_text](./images/img_5.png)
   ![alt_text](./images/img_6.png)

#### Final Design
![alt_text](./images/img.png)



### Token based zookeeper approach

```
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/samuel/go-zookeeper/zk"
)

type IDState struct {
	RangeID    string `json:"range_id"`
	CurrentID  int64  `json:"current_id"`
	ServerID   string `json:"server_id"`
	RangeStart int64  `json:"range_start"`
	RangeEnd   int64  `json:"range_end"`
	UpdatedAt  string `json:"updated_at"`
}

type IDServer struct {
	db         *dynamodb.DynamoDB
	table      string
	rangeID    string
	rangeStart int64
	rangeEnd   int64
	nextID     int64
	zkConn     *zk.Conn
}

func NewIDServer(serverID, table string, zkServers []string) *IDServer {
	zkConn, _, err := zk.Connect(zkServers, time.Second*5)
	if err != nil {
		log.Fatalf("Failed to connect to Zookeeper: %v", err)
	}

	rangeStart := allocateRange(zkConn)
	rangeEnd := rangeStart + 999999
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)
	rangeID := fmt.Sprintf("range-%d-%d", rangeStart, rangeEnd)

	s := &IDServer{
		db:         db,
		table:      table,
		rangeID:    rangeID,
		rangeStart: rangeStart,
		rangeEnd:   rangeEnd,
		nextID:     rangeStart,
		zkConn:     zkConn,
	}
	s.saveState(serverID)
	return s
}

func allocateRange(zkConn *zk.Conn) int64 {
	path := "/id_ranges"
	children, _, err := zkConn.Children(path)
	if err != nil && err != zk.ErrNoNode {
		log.Fatalf("Failed to list znodes: %v", err)
	}

	max := int64(0)
	for _, child := range children {
		parts := strings.Split(child, "-")
		if len(parts) == 3 {
			end, _ := strconv.ParseInt(parts[2], 10, 64)
			if end > max {
				max = end
			}
		}
	}
	newStart := max + 1
	newEnd := newStart + 999999
	newRange := fmt.Sprintf("%s/range-%d-%d", path, newStart, newEnd)
	_, err = zkConn.Create(newRange, []byte(""), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if err != nil {
		log.Fatalf("Failed to create znode: %v", err)
	}
	return newStart
}

func (s *IDServer) saveState(serverID string) {
	item := IDState{
		RangeID:    s.rangeID,
		CurrentID:  s.nextID,
		ServerID:   serverID,
		RangeStart: s.rangeStart,
		RangeEnd:   s.rangeEnd,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Failed to marshal state: %v", err)
	}
	_, err = s.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})
	if err != nil {
		log.Fatalf("Failed to save state to DynamoDB: %v", err)
	}
}

func (s *IDServer) GenerateID(serverID string) string {
	id := atomic.AddInt64(&s.nextID, 1)
	if id > s.rangeEnd {
		log.Fatalf("ID range exhausted, restart required")
	}
	s.saveState(serverID)
	return fmt.Sprintf("%010d", id)
}

func main() {
	serverID := os.Getenv("SERVER_ID")
	zkHosts := []string{"127.0.0.1:2181"}
	dynamoTable := "id_state"

	srv := NewIDServer(serverID, dynamoTable, zkHosts)

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		id := srv.GenerateID(serverID)
		w.Write([]byte(id))
	})

	log.Println("ID server running on :8080")
	http.ListenAndServe(":8080", nil)
}
```
