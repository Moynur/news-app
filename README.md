# News App

App for retrieving news feeds

### Areas for improvement
- logging using a log library like zap or logrus to create a lot more meaningful logs and make it easy to write to file most the logs act as debug statements for now, with a logging library you can have logging fields and log for example the transaction ID along with a message, and using a log aggregator tool see the lifespan of a transaction
- Since everything is a GET requests caching would prevent extra db reads
- A nicer way to show that there are no more articles to load 404 is fairly generic and can imply the wrong things
- Cursor implementation could be linked to the client so cursors are unique to the entity requesting articles and cached this way instead
- Client Authentication
- Database integration testing
- Service integration tests
- 3rd party API tests
- adding context to requests -- maybe controversial
- Metrics
- Audit table/log of client requests
- database credentials could be config
- relying on duplicate key error could be handled more gracefully
- this single endpoint could technically service all needs, but it would put a lot more burden on the client (mobile app) to do some heavier lifting providing less info and then having another endpoint for a specific article would be nicer
- Service could be split in 2, one handling sourcing data and another to handle client requests for this information by requesting from sourcing service.

### Setup 
Setup is quite simple assuming you have docker if you want to run this code simply do 
- docker-compose up --build \
This api will be available on localhost:8080 for any requests

### Testing
Unit tests: go test ./... \
Integration tests: go test -tags=integration ./...

### Local Testing 
You can send a CURL request like 
```
curl --location --request GET 'http://localhost:8080/loadArticles' \
--header 'Content-Type: application/json' \
--data-raw '
{
// "cursor": 6, Implemented should return the next cursor value you can use to load more information
// "title": "Shock contraction of 0.3% for UK economy in April as CBI demands '\''vital actions'\'' to prevent recession", Implemented will do a like string match
// "provider": "some provider", Will Accept req but not return anything
// "category": "some category" Will Accept req but return 404 as its not implemented
}
'
```