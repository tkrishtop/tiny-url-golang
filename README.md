## Application usage

The application schema:

Encode long into short: send POST request with an URL body

```
curl -X POST http://localhost:3003/url/v1/long2short \
   -H 'Content-Type: text/plain' \
   -d 'https://example.com/api/v1/users/12345'
```

Decode short URL: send GET request with the hash as a parameter

```
curl -X GET http://localhost:3003/url/v1/short2long/RGSfGb
```

## Run it out

Console 1:
```
$ go run pkg/codec/codec.go 

Request to encode long URL: https://example.com/api/v1/users/12345
Encoded: https://tiny.com/RGSfGb
Request to decode short URL: RGSfGb
Encoded: https://example.com/api/v1/users/12345
Request to encode long URL: https://google.com
Encoded: https://tiny.com/OtIJsc
Request to decode short URL: OtIJsc
Encoded: https://google.com
Request to decode short URL: RGSfGb
Encoded: https://example.com/api/v1/users/12345
Request to decode short URL: XXX
Encoded: not found
```

Console 2:
```
$ curl -X POST http://localhost:3003/url/v1/long2short -H 'Content-Type: text/plain' -d 'https://example.com/api/v1/users/12345'
https://tiny.com/RGSfGb

$ curl -X GET http://localhost:3003/url/v1/short2long/RGSfGb
https://example.com/api/v1/users/12345

$ curl -X POST http://localhost:3003/url/v1/long2short -H 'Content-Type: text/plain' -d 'https://google.com'
OtIJsc

$ curl -X GET http://localhost:3003/url/v1/short2long/OtIJsc
https://google.com

$ curl -X GET http://localhost:3003/url/v1/short2long/RGSfGb
https://example.com/api/v1/users/12345

$ curl -X GET http://localhost:3003/url/v1/short2long/XXX
not found
```