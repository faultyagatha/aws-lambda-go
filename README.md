# aws-lambda-go

API:

GET: /books?isbn=xxx
POST: /books

JSON record: {"isbn":"978-1420931693","title":"The Republic","author":"Plato"}

How to setup:

1. build an executable with `env GOOS=linux GOARCH=amd64` to make sure lambda will be properly invoked:

`env GOOS=linux GOARCH=amd64 go build cmd/main.go`

zip it: `zip -j main.zip main`

2. setup aws cli: 

`install awscli`
`aws configure --profile`

You'll need to configure credentials of the IAM user:

```yaml
AWS Access Key ID [None]: access-key-ID
AWS Secret Access Key [None]: secret-access-key
Default region name [None]: us-east-1
Default output format [None]: json
```

3. get the lambda package for go:

`go get github.com/aws/aws-lambda-go/lambda`

4. set up an IAM role for lambda:

`aws iam create-role --role-name lambda-books-executor \
--assume-role-policy-document file://trust-policy.json`

note the Arn field:
`"Arn": "arn:aws:iam::[account-id]:role/lambda-books-executor",`

5. specify the permissions that the role has:

`aws iam attach-role-policy --role-name lambda-books-executor \
--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole`

6. deploy the lambda to aws:

`aws lambda create-function --function-name books --runtime go1.x \
--role arn:aws:iam::[account-id]:role/lambda-books-executor \
--handler main --zip-file fileb://main.zip`

where [account-id] is the id in the arn role of the lambda

7. test deployed lambda:
`aws lambda invoke --function-name books output.json`

8. create a new table in DynamoDB:

`aws dynamodb create-table --table-name Books \
--attribute-definitions AttributeName=ISBN,AttributeType=S \
--key-schema AttributeName=ISBN,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5`

note `[TableArn]`

9. add items using put-item:

`aws dynamodb put-item --table-name Books --item '{"ISBN": {"S": "978-1420931693"}, "Title": {"S": "The Republic"}, "Author":  {"S": "Plato"}}'`

10. add go package:
`go get github.com/aws/aws-sdk-go`

11. add database code and rebuild + re-zip the file

12. update lambda:

`aws lambda update-function-code --function-name books \
--zip-file fileb://main.zip`

13. add permissions to run GetItem on a DynamoDB instance and attach it to the lambda-books-executor role:

`aws iam put-role-policy --role-name lambda-books-executor \
--policy-name dynamodb-item-crud-role \
--policy-document file://privilege-policy.json`


