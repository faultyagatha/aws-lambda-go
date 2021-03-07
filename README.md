# aws-lambda-go

API:

GET: /books?isbn=xxx
POST: /books

JSON record: {"isbn":"978-1420931693","title":"The Republic","author":"Plato"}

How to setup:

1. setup aws cli: 

`install awscli`
`aws configure --profile`

You'll need to configure credentials of the IAM user:

```yaml
AWS Access Key ID [None]: access-key-ID
AWS Secret Access Key [None]: secret-access-key
Default region name [None]: us-east-1
Default output format [None]: json
```

2. get the lambda package for go:

`go get github.com/aws/aws-lambda-go/lambda`

3. set up an IAM role for lambda:

`aws iam create-role --role-name lambda-books-executor \
--assume-role-policy-document file://trust-policy.json`

note the Arn field:
`"Arn": "arn:aws:iam::[account-id]:role/lambda-books-executor",`

4. specify the permissions that the role has:

`aws iam attach-role-policy --role-name lambda-books-executor \
--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole`

5. deploy the lambda to aws:

`aws lambda create-function --function-name books --runtime go1.x \
--role arn:aws:iam::[account-id]:role/lambda-books-executor \
--handler main --zip-file fileb://main.zip`

where [account-id] is the id in the arn role of the lambda

6. test deployed lambda:
`aws lambda invoke --function-name books output.json`