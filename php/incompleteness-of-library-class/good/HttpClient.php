<?php

class HttpClient
{
    private $baseUrl;
    private $defaultHeaders;

    public function __construct($baseUrl, $defaultHeaders = [])
    {
        $this->baseUrl = $baseUrl;
        $this->defaultHeaders = array_merge([
            'Content-Type: application/json'
        ], $defaultHeaders);
    }

    public function get($endpoint, $headers = [])
    {
        return $this->request('GET', $endpoint, null, $headers);
    }

    public function post($endpoint, $data = null, $headers = [])
    {
        return $this->request('POST', $endpoint, $data, $headers);
    }

    public function put($endpoint, $data = null, $headers = [])
    {
        return $this->request('PUT', $endpoint, $data, $headers);
    }

    public function delete($endpoint, $headers = [])
    {
        return $this->request('DELETE', $endpoint, null, $headers);
    }

    public function patch($endpoint, $data = null, $headers = [])
    {
        return $this->request('PATCH', $endpoint, $data, $headers);
    }

    private function request($method, $endpoint, $data = null, $additionalHeaders = [])
    {
        $url = $this->baseUrl . $endpoint;
        $headers = array_merge($this->defaultHeaders, $additionalHeaders);

        $options = [
            'http' => [
                'method' => $method,
                'header' => implode("\r\n", $headers)
            ]
        ];

        if ($data !== null) {
            $options['http']['content'] = json_encode($data);
        }

        $context = stream_context_create($options);
        $response = file_get_contents($url, false, $context);

        return json_decode($response, true);
    }
}
