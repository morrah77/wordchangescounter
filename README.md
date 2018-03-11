#Word changes counter
Calculates sum of minimal necessary changes for each passed word to make them acceptable

###About implementation and structure
Logic is placed into `matcher` package, namely `matcher/levenshteindistance` implementing rather common solution way for offered problem based on Levenshtein metric.

Except this one `matcher` package contains, just for curiosity, auxiliary subpackage `matcher/straightforward` implementing really brute force solution, so, if you want to drive your machine hanging out, import `matcher/straightforward` instead of `matcher/levenshteindistance` in `main.go` :). 

Some unit tests && benchmarks also present.

##Build
`./control.sh build`

or

`rm -rf ./build/* && go build -o ./build/wccounter && chmod u+x ./build/wccounter && cp -r ./resources/ ./build/ && cp ./docs/example_input ./build && cp ./docs/187 ./build`

or even with docker (meaningless for this task, just for example)

`docker build -t wordchangescounter .`
##Run
`./control.sh run`

or

`cd ./build && ./wccounter`

or

`./build/wccounter --source-path='./build/187'`

or

`./build/wccounter --source-path='<path/to/input/file>'`

or even with docker (meaningless for this task, just for example)

`docker run --rm -it --name=wordchangescounter wordchangescounter`

##Test

###Automatically (unit tests just for matcher package)

`./control.sh test`

or

`./control.sh bench .` (to run tests with benchmarks)

or

`go test ./matcher/core`

`go test ./matcher/levenshtaindistance/`

`go test ./matcher/straightforward/`

or

`go test ./matcher/<levenshtaindistance|straightforward> -bench .`

(there're several benchmarks for important packages functions)

###Manually
`./build/wccounter --source-path='./build/187'`

must print `187
`

or even with docker (meaningless for this task, just for example)

`docker build -t wordchangescounter . && docker run --rm -it --name=wordchangescounter wordchangescounter`
