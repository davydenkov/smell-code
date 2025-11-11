import json
import urllib.request

class HttpClient:
    def __init__(self, base_url):
        self.base_url = base_url

    def get(self, endpoint):
        url = self.base_url + endpoint
        req = urllib.request.Request(url)
        req.add_header('Content-Type', 'application/json')

        with urllib.request.urlopen(req) as response:
            return json.loads(response.read().decode())

    # Missing POST, PUT, DELETE methods - incomplete library!
