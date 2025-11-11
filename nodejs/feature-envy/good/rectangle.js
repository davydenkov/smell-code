class Rectangle {
    constructor(width, height) {
        this.width = width;
        this.height = height;
    }

    getWidth() {
        return this.width;
    }

    getHeight() {
        return this.height;
    }

    calculateArea() {
        return this.width * this.height;
    }

    calculatePerimeter() {
        return 2 * (this.width + this.height);
    }

    isSquare() {
        return this.width === this.height;
    }

    calculateDiagonal() {
        return Math.sqrt(this.width ** 2 + this.height ** 2);
    }

    getAspectRatio() {
        return this.width / this.height;
    }
}

class GeometryUtils {
    // Utility methods that don't belong to Rectangle can stay here
    static calculateDistanceBetweenPoints(x1, y1, x2, y2) {
        return Math.sqrt((x2 - x1) ** 2 + (y2 - y1) ** 2);
    }

    static calculateAngle(opposite, adjacent) {
        return Math.atan(opposite / adjacent);
    }
}

module.exports = {
    Rectangle,
    GeometryUtils
};
