# HTTP Mocker [![Build Status](https://api.travis-ci.org/kulshekhar/http-mocker.svg?branch=master)](https://travis-ci.org/kulshekhar/http-mocker)

HTTP Mocker can be used to simulate HTTP responses from an API service. It returns a response that is identical to the payload of the request. It can also be helpful in mocking AJAX requests made by third party Javascript libraries.

## Usage

The release page currently has a pre-built binary for Ubuntu x64. If there is enough demand, additional pre-built binary may be added in the future. For now, you have to manually build the binary for other platforms.

To build HTTP Mocker, you need `Go` installed. Once you have `Go` installed, clone this repository and build it:

```
git clone https://github.com/kulshekhar/http-mocker
cd http-mocker
go build
```

To run it:

```
./http-mocker
```

> Note: This starts HTTP Mocker on port 18000. You can customize the port by setting the `INTERCEPTOR_PORT` environment variable.

Now HTTP Mocker will respond to an HTTP POST request with a response identical to the value of the `response` key from the request payload. For example, if you make the following request:

```
curl -X POST -H 'Content-Type:application/json' -d '{"response":1}' -i http://localhost:18000
```

HTTP Mocker will respond with `1`. If you make the following request:

```
curl -X POST -H 'Content-Type:application/json' -d '{"response":{"f1":"aa"}}' -i http://localhost:18000
```

HTTP Mocker will respond with:

```json
{"f1":"aa"}
```

## AJAX Interceptor

When HTTP Mocker is running, accessing `http://localhost:18000/_script.js` will serve a script that will allow you to intercept and redirect AJAX requests. You can include it like so:

```html
<script src="http://localhost:18000/_script.js"></script>
```

and use it in this manner:

```javascript
let xi = new XHRInterceptor();

xi.matchAll([
  { expression: /\/post\/[0-9]+\/[a-zA-Z]+/, response: { message: 'nice post' }, statusCode: 200 },
  { expression: '/get', response: { message: 'nice get' } },
]);

xi.intercept();
```
The above snippet will intercept two types of AJAX requests (to any domain):

- all `POST` requests to any URL whose pathname match the form `/post/<SOME_NUMBER>/<SOME_TEXT>` (for example, `http://example.com/post/123/abc`)
- all `GET` requests to `/get`

You can also set the `statusCode` for a particular matcher. If set, HTTP Mocker will respond with this status code. The default response status code is `200`.

these requests will be replaced by `POST` requests to HTTP Mocker. The content of the `response` key in the matcher will be the payload sent to HTTP Mocker.

For example, the equivalent `curl` command for the first matcher will be:

```
curl -X POST -H 'Content-Type:application/json' -d '{"response":{"message":"nice post"}}' -i http://localhost:18000
```

and the equivalent `curl` command for the second matcher will be:

```
curl -X POST -H 'Content-Type:application/json' -d '{"response":{"message":"nice get"}}' -i http://localhost:18000
```

To be effective, the value of the `response` key should be identical to the actual response of the API call that is being mocked.
