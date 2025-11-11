public class Rectangle {
    public double width;
    public double height;

    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }
}

class GeometryUtils {
    // This method has feature envy - it accesses many fields of Rectangle
    public double calculateArea(Rectangle rectangle) {
        return rectangle.width * rectangle.height;
    }

    public double calculatePerimeter(Rectangle rectangle) {
        return 2 * (rectangle.width + rectangle.height);
    }

    public boolean isSquare(Rectangle rectangle) {
        return rectangle.width == rectangle.height;
    }

    public double calculateDiagonal(Rectangle rectangle) {
        return Math.sqrt(rectangle.width * rectangle.width + rectangle.height * rectangle.height);
    }

    public double getAspectRatio(Rectangle rectangle) {
        return rectangle.width / rectangle.height;
    }
}
