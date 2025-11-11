<?php

class Rectangle {
    public $width;
    public $height;

    public function __construct($width, $height) {
        $this->width = $width;
        $this->height = $height;
    }
}

class GeometryUtils {
    // This method has feature envy - it accesses many fields of Rectangle
    public function calculateArea(Rectangle $rectangle) {
        return $rectangle->width * $rectangle->height;
    }

    public function calculatePerimeter(Rectangle $rectangle) {
        return 2 * ($rectangle->width + $rectangle->height);
    }

    public function isSquare(Rectangle $rectangle) {
        return $rectangle->width === $rectangle->height;
    }

    public function calculateDiagonal(Rectangle $rectangle) {
        return sqrt($rectangle->width ** 2 + $rectangle->height ** 2);
    }

    public function getAspectRatio(Rectangle $rectangle) {
        return $rectangle->width / $rectangle->height;
    }
}
