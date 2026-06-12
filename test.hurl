# Tests for /api/{number} route
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

GET http://localhost:8080/api/8
HTTP 200
[Asserts]
jsonpath "$.message" == "8"

GET http://localhost:8080/api/9
HTTP 200
[Asserts]
jsonpath "$.message" == "Fizz"

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
jsonpath "$.error" contains "Please upgrade to a paid plan"

GET http://localhost:8080/api/e%3Dmc2
HTTP 400
[Asserts]
jsonpath "$.error" contains "Invalid number"

# Tests for /api/fizz/{number} route
GET http://localhost:8080/api/fizz/3
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/fizz/5
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/fizz/9
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/fizz/10
HTTP 200
[Asserts]
jsonpath "$" == false

# Tests for /api/buzz/{number} route
GET http://localhost:8080/api/buzz/3
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/buzz/5
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/buzz/9
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/buzz/10
HTTP 200
[Asserts]
jsonpath "$" == true

# Tests for /api/fizzbuzz/{number} route
GET http://localhost:8080/api/fizzbuzz/3
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/fizzbuzz/5
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/fizzbuzz/15
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/fizzbuzz/30
HTTP 200
[Asserts]
jsonpath "$" == true

# Tests for /api/prime/{number} route
GET http://localhost:8080/api/prime/1
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/2
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/3
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/4
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/5
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/9
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/15
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/17
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/19
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/23
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/24
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/25
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/29
HTTP 200
[Asserts]
jsonpath "$" == true

GET http://localhost:8080/api/prime/151515
HTTP 200
[Asserts]
jsonpath "$" == false

GET http://localhost:8080/api/prime/15151515
HTTP 402
[Asserts]
jsonpath "$.error" contains "Please upgrade to a paid plan"

GET http://localhost:8080/api/prime/e%3Dmc2
HTTP 400
[Asserts]
jsonpath "$.error" contains "Invalid number"