class Rectangle {
    constructor(width, height) {
        this.width = width;
        this.height = height;
    }
}

class GeometryUtils {
    // This method has feature envy - it accesses many fields of Rectangle
    calculateArea(rectangle) {
        return rectangle.width * rectangle.height;
    }

    calculatePerimeter(rectangle) {
        return 2 * (rectangle.width + rectangle.height);
    }

    isSquare(rectangle) {
        return rectangle.width === rectangle.height;
    }

    calculateDiagonal(rectangle) {
        return Math.sqrt(rectangle.width ** 2 + rectangle.height ** 2);
    }

    getAspectRatio(rectangle) {
        return rectangle.width / rectangle.height;
    }
}

module.exports = {
    Rectangle,
    GeometryUtils
};
