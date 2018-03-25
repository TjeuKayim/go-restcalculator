# restcalculator

A REST server for the [newton-api](https://github.com/aunyks/newton-api) that uses Go's standard library.

*Currently, it only supports the `abs` operation.*

## Newton-API

### How does it work?
1. Send a GET request to newton with a url-encoded math expression and your preferred operation.
2. Get back a JSON response with your problem solved.

### Requests
1. Send a GET request to newton.
```
https://newton.now.sh/:operation/:expression
```
Note: `:operation` is the math operation that you want to perform. 
`:expression` is the *url-encoded* math expression on which you want to operate.

2. You'll be returned a JSON object with the operation you requested, 
the expression you provided, and the result of the operation performed on the expression.

```
{
  "operation":"derive",
  "expression":"x^2",
  "result":"2 x"
}
```

### Endpoints
| Operation |    API Endpoint   |       Result      |
|:---------:|:-----------------:|:-----------------:|
| Simplify  | /simplify/2^2+2(2)| 8                 |
| Factor    | /factor/x^2 + 2x  | x (x + 2)         |
| Derive    | /derive/x^2+2x    | 2 x + 2           |
| Integrate | /integrate/x^2+2x | 1/3 x^3 + x^2 + C |
| Find 0's  | /zeroes/x^2+2x    | [-2, 0]           |
| Find Tangent| /tangent/2lx^3  | 12 x + -16        |
| Area Under Curve| /area/2:4lx^3| 60               |
| Cosine    | /cos/pi            | -1                 |
| Sine      | /sin/0            | 0                 |
| Tangent   | /tan/0            | 0                 |
| Inverse Cosine    | /arccos/1            | 0                 |
| Inverse Sine    | /arcsin/0            | 0                 |
| Inverse Tangent    | /arctan/0            | 0                 |
| Absolute Value    | /abs/-1            | 1                 |  
| Logarithm | /log/2l8           | 3               |
