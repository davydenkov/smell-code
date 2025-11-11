import math

class Rectangle:
    def __init__(self, width: float, height: float):
        self._width = width
        self._height = height

    def get_width(self) -> float:
        return self._width

    def get_height(self) -> float:
        return self._height

    def calculate_area(self) -> float:
        return self._width * self._height

    def calculate_perimeter(self) -> float:
        return 2 * (self._width + self._height)

    def is_square(self) -> bool:
        return self._width == self._height

    def calculate_diagonal(self) -> float:
        return math.sqrt(self._width ** 2 + self._height ** 2)

    def get_aspect_ratio(self) -> float:
        return self._width / self._height


class GeometryUtils:
    # Utility methods that don't belong to Rectangle can stay here
    @staticmethod
    def calculate_distance_between_points(x1: float, y1: float, x2: float, y2: float) -> float:
        return math.sqrt((x2 - x1) ** 2 + (y2 - y1) ** 2)

    @staticmethod
    def calculate_angle(opposite: float, adjacent: float) -> float:
        return math.atan(opposite / adjacent)
