import math

class Rectangle:
    def __init__(self, width, height):
        self.width = width
        self.height = height


class GeometryUtils:
    # This method has feature envy - it accesses many fields of Rectangle
    def calculate_area(self, rectangle):
        return rectangle.width * rectangle.height

    def calculate_perimeter(self, rectangle):
        return 2 * (rectangle.width + rectangle.height)

    def is_square(self, rectangle):
        return rectangle.width == rectangle.height

    def calculate_diagonal(self, rectangle):
        return math.sqrt(rectangle.width ** 2 + rectangle.height ** 2)

    def get_aspect_ratio(self, rectangle):
        return rectangle.width / rectangle.height
