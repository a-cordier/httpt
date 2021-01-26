# httpt

## Example 

```yaml
version: '3.4'

services:
  httpt:
    image: ghcr.io/a-cordier/httpt:latest
    ports:
      - 8888:8080
```

HTTPT is a simple test server used by the developer-portal team for integration
testing and debugging purpose.

## Available endpoints

### `/status/{statusCode}`

For each status code defined by the HTTP spec, httpt exposes a handler that
will send back this status, along with a text response indicating the status
message. The handlers are designed to respond on any method.

- [*] /status/101 (Returns HTTP status 101: Switching Protocols)
- [*] /status/102 (Returns HTTP status 102: Processing)
- [*] /status/103 (Returns HTTP status 103: Early Hints)
- [*] /status/200 (Returns HTTP status 200: OK)
- [*] /status/201 (Returns HTTP status 201: Created)
- [*] /status/202 (Returns HTTP status 202: Accepted)
- [*] /status/203 (Returns HTTP status 203: Non-Authoritative Information)
- [*] /status/204 (Returns HTTP status 204: No Content)
- [*] /status/205 (Returns HTTP status 205: Reset Content)
- [*] /status/206 (Returns HTTP status 206: Partial Content)
- [*] /status/207 (Returns HTTP status 207: Multi-Status)
- [*] /status/208 (Returns HTTP status 208: Already Reported)
- [*] /status/226 (Returns HTTP status 226: IM Used)
- [*] /status/300 (Returns HTTP status 300: Multiple Choices)
- [*] /status/302 (Returns HTTP status 302: Found)
- [*] /status/303 (Returns HTTP status 303: See Other)
- [*] /status/304 (Returns HTTP status 304: Not Modified)
- [*] /status/305 (Returns HTTP status 305: Use Proxy)
- [*] /status/307 (Returns HTTP status 307: Temporary Redirect)
- [*] /status/308 (Returns HTTP status 308: Permanent Redirect)
- [*] /status/400 (Returns HTTP status 400: Bad Request)
- [*] /status/401 (Returns HTTP status 401: Unauthorized)
- [*] /status/402 (Returns HTTP status 402: Payment Required)
- [*] /status/403 (Returns HTTP status 403: Forbidden)
- [*] /status/404 (Returns HTTP status 404: Not Found)
- [*] /status/405 (Returns HTTP status 405: Method Not Allowed)
- [*] /status/406 (Returns HTTP status 406: Not Acceptable)
- [*] /status/407 (Returns HTTP status 407: Proxy Authentication Required)
- [*] /status/408 (Returns HTTP status 408: Request Timeout)
- [*] /status/409 (Returns HTTP status 409: Conflict)
- [*] /status/410 (Returns HTTP status 410: Gone)
- [*] /status/411 (Returns HTTP status 411: Length Required)
- [*] /status/412 (Returns HTTP status 412: Precondition Failed)
- [*] /status/413 (Returns HTTP status 413: Request Entity Too Large)
- [*] /status/414 (Returns HTTP status 414: Request URI Too Long)
- [*] /status/415 (Returns HTTP status 415: Unsupported Media Type)
- [*] /status/416 (Returns HTTP status 416: Requested Range Not Satisfiable)
- [*] /status/417 (Returns HTTP status 417: Expectation Failed)
- [*] /status/418 (Returns HTTP status 418: I'm a teapot)
- [*] /status/421 (Returns HTTP status 421: Misdirected Request)
- [*] /status/422 (Returns HTTP status 422: Unprocessable Entity)
- [*] /status/423 (Returns HTTP status 423: Locked)
- [*] /status/424 (Returns HTTP status 424: Failed Dependency)
- [*] /status/425 (Returns HTTP status 425: Too Early)
- [*] /status/426 (Returns HTTP status 426: Upgrade Required)
- [*] /status/428 (Returns HTTP status 428: Precondition Required)
- [*] /status/429 (Returns HTTP status 429: Too Many Requests)
- [*] /status/431 (Returns HTTP status 431: Request Header Fields Too Large)
- [*] /status/451 (Returns HTTP status 451: Unavailable For Legal Reasons)
- [*] /status/500 (Returns HTTP status 500: Internal Server Error)
- [*] /status/501 (Returns HTTP status 501: Not Implemented)
- [*] /status/502 (Returns HTTP status 502: Bad Gateway)
- [*] /status/503 (Returns HTTP status 503: Service Unavailable)
- [*] /status/504 (Returns HTTP status 504: Gateway Timeout)
- [*] /status/505 (Returns HTTP status 505: HTTP Version Not Supported)
- [*] /status/506 (Returns HTTP status 506: Variant Also Negotiates)
- [*] /status/507 (Returns HTTP status 507: Insufficient Storage)
- [*] /status/508 (Returns HTTP status 508: Loop Detected)
- [*] /status/510 (Returns HTTP status 510: Not Extended)
- [*] /status/511 (Returns HTTP status 511: Network Authentication Required)

### `/headers`

This handler will print the request headers received by httpt. This can be handy
if you are testing how headers are forwarded by an intermediate proxy. Again,
the handler works with any HTTP method.

- [*] /headers

### `/wait?seconds={count}`

This handler waits for `{count}` seconds before returning a `2OO OK` response.
This can prove usefull if you are looking after an intermediate proxy timeout.
The endpoint catches all methods as well;

- [*] /wait?seconds={count}

### `/john-doe?times={count}`

The john doe endpoint will answer an array of simple payloads that are designed
to weight each 100 bytes. That way, one can test a proxy/client behavior for
different kind of payload. Might prove usefull if you are trying to understand
chunk behavior for an intermediate proxy or to benchmark latencies / debug bandwith
issues. Again, the endpoint handles all HTTP methods.

- [*] /john-doe?times={count}
