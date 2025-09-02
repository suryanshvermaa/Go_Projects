## Lambda deployment

### For AWS
```bash
aws iam create-ro
le   --role-name lambda-ex   --assume-role-policy-document '{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }'
```

### For localstack

```bash
awslocal iam create-role   --role-name lambda-ex   --assume-role-policy-document '{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }'
```

## Second Command
```bash
aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
```
## Third Command

```bash
aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

### setup
- zip file as
```bash
zip function.zip main
```

## Fourth command
```bash
aws lambda create-function \
  --function-name go-lambda3 \
  --zip-file fileb://function.zip \
  --handler main \
  --runtime go1.x \
  --role arn:aws:iam::000000000000:role/lambda-ex
```

## Invoking Function
```bash
aws lambda invoke   --function-name go-lambda-function --cli-binary-format raw-in-base64-out --payload '{"what is your name?": "suryansh", "what is your age?": 22}'   output.txt
```