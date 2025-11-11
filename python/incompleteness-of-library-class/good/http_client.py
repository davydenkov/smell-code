import json
import urllib.request

class HttpClient:
    def __init__(self, base_url, default_headers=None):
        self.base_url = base_url
        self.default_headers = default_headers or {}
        self.default_headers.setdefault('Content-Type', 'application/json')

    def get(self, endpoint, headers=None):
        return self._request('GET', endpoint, None, headers)

    def post(self, endpoint, data=None, headers=None):
        return self._request('POST', endpoint, data, headers)

    def put(self, endpoint, data=None, headers=None):
        return self._request('PUT', endpoint, data, headers)

    def delete(self, endpoint, headers=None):
        return self._request('DELETE', endpoint, None, headers)

    def patch(self, endpoint, data=None, headers=None):
        return self._request('PATCH', endpoint, data, headers)

    def _request(self, method, endpoint, data=None, additional_headers=None):
        url = self.base_url + endpoint
        headers = {**self.default_headers}
        if additional_headers:
            headers.update(additional_headers)

        req = urllib.request.Request(url, method=method)
        for key, value in headers.items():
            req.add_header(key, value)

        if data is not None:
            req.data = json.dumps(data).encode()

        try:
            with urllib.request.urlopen(req) as response:
                return json.loads(response.read().decode())
        except Exception as e:
            # Handle errors appropriately
            raise e
