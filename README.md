# News App

App for retrieving news feeds

### Areas for improvement
- logging using a log library like zap or logrus to create a lot more meaningful logs and make it easy to write to file most the logs act as debug statements for now, with a logging library you can have logging fields and log for example the transaction ID along with a message, and using a log aggregator tool see the lifespan of a transaction
- Idempotency on all requests
- Client Authentication
- Database integration testing
- The code to decide to accept based on pan could have en re-framed as a fraud engine or something similar to give more context on why it was declined
- A table/db for card information storage the pan is only stored for the purposes of fulfilling the requirements of the task, to take this further a separate table or db which owns all PCI information would be nice in reality a different service
- Metrics

### Setup 
Setup is quite simple assuming you have docker if you want to run this code simply do 
- docker-compose up --build \
This api will be available on localhost:8080 for any requests

### Testing
Unit tests: go test ./... \
Integration tests: go test -tags=integration ./...
