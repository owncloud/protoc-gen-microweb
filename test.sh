#!/bin/bash

# Test for protoc-gen-microweb
# Tests complete generation and usage pipeline

set -e

assert_file_exists() {
    if [ ! -f "$1" ]; then
        echo "❌ File $1 does not exist"
        exit 1
    fi
}

main() {
  echo "=== protoc-gen-microweb Test ==="
  # get git roor dir
  GIT_ROOT=$(git rev-parse --show-toplevel)
  echo "GIT_ROOT: $GIT_ROOT"
  pushd $GIT_ROOT

  echo "0. Assert initial example state"
  # Use greeter example
  cd examples/greeter
  rm -f greeter.pb.go
  rm -f greeter.pb.web.go
  rm -f greeter.pb.micro.go
  assert_file_exists proto/greeter.proto
  assert_file_exists main.go
  assert_file_exists go.mod
  assert_file_exists go.sum

  # Generate code - microweb creates flat structure regardless of module parameter
  echo "1. Generate code..."
  protoc \
    --proto_path=/tmp/googleapis \
    --proto_path=proto/ \
    --go_out=proto/ \
    --go_opt=module=github.com/owncloud/protoc-gen-microweb/examples/greeter/proto \
    --go-grpc_out=proto/ \
    --go-grpc_opt=module=github.com/owncloud/protoc-gen-microweb/examples/greeter/proto \
    --micro_out=proto/ \
    --micro_opt=module=github.com/owncloud/protoc-gen-microweb/examples/greeter/proto \
    --microweb_out=proto/ \
    --microweb_opt=module=github.com/owncloud/protoc-gen-microweb/examples/greeter/proto \
    proto/greeter.proto

  # Assert generated files are in the correct location
  assert_file_exists proto/greeter.proto
  assert_file_exists proto/greeter.pb.go
  assert_file_exists proto/greeter.pb.web.go
  assert_file_exists proto/greeter.pb.micro.go

  echo "2. Run and test server..."
  go run main.go &
  SERVER_PID=$!

  # Wait for server to start
  sleep 2
  
  echo "3. Test POST /api/say..."
  RESPONSE=$(curl -s -w "\n%{http_code}" -X POST http://localhost:8080/api/say \
    -H "Content-Type: application/json" \
    -d '{"name":"test"}')

  # Extract response body and status code
  RESPONSE_BODY=$(echo "$RESPONSE" | head -n 1)
  STATUS_CODE=$(echo "$RESPONSE" | tail -n 1)

  kill $SERVER_PID 2>/dev/null || true

  echo "4. Validate results..."
  echo "Response: $RESPONSE_BODY"
  echo "Status: $STATUS_CODE"

  if [ "$RESPONSE_BODY" = '{"message":"Hello test!"}' ] && [ "$STATUS_CODE" = "201" ]; then
      echo "PASS"
      exit 0
  else
      echo "Expected: {\"message\":\"Hello test!\"} with status 201"
      echo "Got: $RESPONSE_BODY with status $STATUS_CODE"
      echo "FAIL"
      exit 1
  fi
}

main

