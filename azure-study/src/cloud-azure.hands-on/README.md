
brew tap aws/tap
brew install aws-sam-cli
sam --version
sam init --runtime go1.x --name step4


go get "github.com/aws/aws-lambda-go/events"
go get "github.com/aws/aws-lambda-go/lambda"

sam local start-api