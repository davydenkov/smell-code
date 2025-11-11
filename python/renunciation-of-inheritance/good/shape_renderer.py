import math
from abc import ABC, abstractmethod

class ShapeRenderer(ABC):
    def __init__(self, color='black'):
        self.color = color

    @abstractmethod
    def render(self):
        pass

    @abstractmethod
    def get_area(self):
        pass

    def get_color(self):
        return self.color

    def set_color(self, color):
        self.color = color

    def _get_render_prefix(self):
        return f"Rendering in {self.color}"


class CircleRenderer(ShapeRenderer):
    def __init__(self, radius, color='black'):
        super().__init__(color)
        self.radius = radius

    def render(self):
        print(f"{self._get_render_prefix()} circle with radius {self.radius}")

    def get_area(self):
        return math.pi * self.radius * self.radius

    def get_radius(self):
        return self.radius


class RectangleRenderer(ShapeRenderer):
    def __init__(self, width, height, color='black'):
        super().__init__(color)
        self.width = width
        self.height = height

    def render(self):
        print(f"{self._get_render_prefix()} rectangle {self.width}x{self.height}")

    def get_area(self):
        return self.width * self.height

    def get_width(self):
        return self.width

    def get_height(self):
        return self.height


class TriangleRenderer(ShapeRenderer):
    def __init__(self, base, height, color='black'):
        super().__init__(color)
        self.base = base
        self.height = height

    def render(self):
        print(f"{self._get_render_prefix()} triangle with base {self.base} and height {self.height}")

    def get_area(self):
        return 0.5 * self.base * self.height

    def get_base(self):
        return self.base

    def get_height(self):
        return self.height

# Proper inheritance eliminates code duplication and provides consistent interface
