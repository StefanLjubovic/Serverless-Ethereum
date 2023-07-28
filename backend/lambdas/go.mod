module lambdas

go 1.20

require github.com/aws/aws-lambda-go v1.41.0

replace backend => ../../service

require backend v0.0.0-00010101000000-000000000000
