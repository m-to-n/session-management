dapr run --log-level debug --app-id session-management --app-protocol grpc --dapr-http-port 3400 --dapr-grpc-port 34000 --app-port 34001 --components-path ./dapr-components go run main.go