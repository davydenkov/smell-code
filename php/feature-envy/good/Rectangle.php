<?php

class Rectangle {
    private float $width;
    private float $height;

    public function __construct(float $width, float $height) {
        $this->width = $width;
        $this->height = $height;
    }

    public function getWidth(): float {
        return $this->width;
    }

    public function getHeight(): float {
        return $this->height;
    }

    public function calculateArea(): float {
        return $this->width * $this->height;
    }

    public function calculatePerimeter(): float {
        return 2 * ($this->width + $this->height);
    }

    public function isSquare(): bool {
        return $this->width === $this->height;
    }

    public function calculateDiagonal(): float {
        return sqrt($this->width ** 2 + $this->height ** 2);
    }

    public function getAspectRatio(): float {
        return $this->width / $this->height;
    }
}

class GeometryUtils {
    // Utility methods that don't belong to Rectangle can stay here
    public static function calculateDistanceBetweenPoints(float $x1, float $y1, float $x2, float $y2): float {
        return sqrt(($x2 - $x1) ** 2 + ($y2 - $y1) ** 2);
    }

    public static function calculateAngle(float $opposite, float $adjacent): float {
        return atan($opposite / $adjacent);
    }
}
