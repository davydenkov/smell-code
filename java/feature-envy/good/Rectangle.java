public class Rectangle {
    private double width;
    private double height;

    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }

    public double getWidth() {
        return width;
    }

    public double getHeight() {
        return height;
    }

    public double calculateArea() {
        return width * height;
    }

    public double calculatePerimeter() {
        return 2 * (width + height);
    }

    public boolean isSquare() {
        return width == height;
    }

    public double calculateDiagonal() {
        return Math.sqrt(width * width + height * height);
    }

    public double getAspectRatio() {
        return width / height;
    }
}

class GeometryUtils {
    // Utility methods that don't belong to Rectangle can stay here
    public static double calculateDistanceBetweenPoints(double x1, double y1, double x2, double y2) {
        return Math.sqrt((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1));
    }

    public static double calculateAngle(double opposite, double adjacent) {
        return Math.atan(opposite / adjacent);
    }
}
