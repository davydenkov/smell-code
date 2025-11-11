<?php

class HttpClient
{
    private $baseUrl;

    public function __construct($baseUrl)
    {
        $this->baseUrl = $baseUrl;
    }

    public function get($endpoint)
    {
        $url = $this->baseUrl . $endpoint;
        $context = stream_context_create([
            'http' => [
                'method' => 'GET',
                'header' => 'Content-Type: application/json'
            ]
        ]);

        $response = file_get_contents($url, false, $context);
        return json_decode($response, true);
    }

    // Missing POST, PUT, DELETE methods - incomplete library!
}
