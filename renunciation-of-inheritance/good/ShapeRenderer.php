<?php

abstract class ShapeRenderer
{
    protected $color;

    public function __construct($color = 'black')
    {
        $this->color = $color;
    }

    abstract public function render();
    abstract public function getArea();

    public function getColor()
    {
        return $this->color;
    }

    public function setColor($color)
    {
        $this->color = $color;
    }

    protected function getRenderPrefix()
    {
        return "Rendering in {$this->color}";
    }
}

class CircleRenderer extends ShapeRenderer
{
    private $radius;

    public function __construct($radius, $color = 'black')
    {
        parent::__construct($color);
        $this->radius = $radius;
    }

    public function render()
    {
        echo $this->getRenderPrefix() . " circle with radius {$this->radius}\n";
    }

    public function getArea()
    {
        return pi() * $this->radius * $this->radius;
    }

    public function getRadius()
    {
        return $this->radius;
    }
}

class RectangleRenderer extends ShapeRenderer
{
    private $width;
    private $height;

    public function __construct($width, $height, $color = 'black')
    {
        parent::__construct($color);
        $this->width = $width;
        $this->height = $height;
    }

    public function render()
    {
        echo $this->getRenderPrefix() . " rectangle {$this->width}x{$this->height}\n";
    }

    public function getArea()
    {
        return $this->width * $this->height;
    }

    public function getWidth()
    {
        return $this->width;
    }

    public function getHeight()
    {
        return $this->height;
    }
}

class TriangleRenderer extends ShapeRenderer
{
    private $base;
    private $height;

    public function __construct($base, $height, $color = 'black')
    {
        parent::__construct($color);
        $this->base = $base;
        $this->height = $height;
    }

    public function render()
    {
        echo $this->getRenderPrefix() . " triangle with base {$this->base} and height {$this->height}\n";
    }

    public function getArea()
    {
        return 0.5 * $this->base * $this->height;
    }

    public function getBase()
    {
        return $this->base;
    }

    public function getHeight()
    {
        return $this->height;
    }
}

// Proper inheritance eliminates code duplication and provides consistent interface
