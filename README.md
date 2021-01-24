# sam-sample

```bash
# setup
aws configure set aws_access_key_id dummy     --profile local
aws configure set aws_secret_access_key dummy --profile local
aws configure set region ap-northeast-1       --profile local

# compile
docker-compose run --rm go bash
`GOOS=linux GOARCH=amd64 go build main.go`

# start
docker-compose up
sam local start-api --docker-network samapp

# test
curl http://localhost:3000/hello

# dynamo
# create test data
# host
aws dynamodb --profile local --endpoint-url http://localhost:4569 create-table --cli-input-json file://./doc/db_local.json
# scan
aws dynamodb --profile local --endpoint-url http://localhost:4569 scan --table-name local_company_table 
# delte
aws dynamodb --profile local --endpoint-url http://localhost:4569 delete-table --table-name local_company_table


```