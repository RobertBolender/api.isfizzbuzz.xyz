GET http://localhost:8080/api/1
HTTP 200
[Asserts]
jsonpath "$.message" == "1"

GET http://localhost:8080/api/3
HTTP 200
[Asserts]
jsonpath "$.message" == "Fizz"

GET http://localhost:8080/api/5
HTTP 200
[Asserts]
jsonpath "$.message" == "Buzz"

GET http://localhost:8080/api/15
HTTP 200
[Asserts]
jsonpath "$.message" == "FizzBuzz"

GET http://localhost:8080/api/151515
HTTP 200
[Asserts]
jsonpath "$.message" == "FizzBuzz"

GET http://localhost:8080/api/15151515
HTTP 402
[Asserts]
body contains "Please upgrade to a paid plan"

GET http://localhost:8080/api/e=mc2
HTTP 400
[Asserts]
body contains "Invalid number"