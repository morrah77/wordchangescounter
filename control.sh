#!/bin/sh

if [ -z "$1" ]; then
 echo 'Please specify a command [build|test|bench|run]'
fi

case "$1" in
  build)
    echo "Building..."
    rm -rf ./build/* && go build -o ./build/wccounter && chmod u+x ./build/wccounter && cp -r ./resources/ ./build/ && cp ./docs/example_input ./build && cp ./docs/187 ./build
     if [ $? = 0 ]; then
       echo "Built successfully"
     else
       echo 'Build failed'
     fi ;;
  test)
    echo "Running tests..."
    go test ./matcher/core/ && \
    go test ./matcher/levenshtaindistance/ && \
    go test ./matcher/straightforward/ ;;
  bench)
    echo "Running tests with benchmarks..."
    go test ./matcher/core/ && \
    go test ./matcher/levenshtaindistance/ -bench . && \
    go test ./matcher/straightforward/ -bench . ;;
  run)
    echo "Running..."
    START_TIME=`date +%s`
    ./build/wccounter --source-path='./build/187'
    FINISH_TIME=`date +%s`
    DURATION=$((FINISH_TIME-START_TIME))
    if [ $? = 0 ]; then
      echo "Finished successfully in "$DURATION"s"
    else
      echo 'Failed'
    fi ;;
esac
