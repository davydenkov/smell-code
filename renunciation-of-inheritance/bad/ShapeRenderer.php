<?php

class CircleRenderer
{
    private $color;

    public function __construct($color = 'black')
    {
        $this->color = $color;
    }

    public function render($radius)
    {
        echo "Rendering circle with radius {$radius} in {$this->color}\n";
    }

    public function getArea($radius)
    {
        return pi() * $radius * $radius;
    }

    public function getColor()
    {
        return $this->color;
    }

    public function setColor($color)
    {
        $this->color = $color;
    }
}

class RectangleRenderer
{
    private $color;

    public function __construct($color = 'black')
    {
        $this->color = $color;
    }

    public function render($width, $height)
    {
        echo "Rendering rectangle {$width}x{$height} in {$this->color}\n";
    }

    public function getArea($width, $height)
    {
        return $width * $height;
    }

    public function getColor()
    {
        return $this->color;
    }

    public function setColor($color)
    {
        $this->color = $color;
    }
}

class TriangleRenderer
{
    private $color;

    public function __construct($color = 'black')
    {
        $this->color = $color;
    }

    public function render($base, $height)
    {
        echo "Rendering triangle with base {$base} and height {$height} in {$this->color}\n";
    }

    public function getArea($base, $height)
    {
        return 0.5 * $base * $height;
    }

    public function getColor()
    {
        return $this->color;
    }

    public function setColor($color)
    {
        $this->color = $color;
    }
}

// Code duplication in color management and structure - should use inheritance!
