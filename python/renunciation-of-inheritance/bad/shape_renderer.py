import math

class CircleRenderer:
    def __init__(self, color='black'):
        self.color = color

    def render(self, radius):
        print(f"Rendering circle with radius {radius} in {self.color}")

    def get_area(self, radius):
        return math.pi * radius * radius

    def get_color(self):
        return self.color

    def set_color(self, color):
        self.color = color


class RectangleRenderer:
    def __init__(self, color='black'):
        self.color = color

    def render(self, width, height):
        print(f"Rendering rectangle {width}x{height} in {self.color}")

    def get_area(self, width, height):
        return width * height

    def get_color(self):
        return self.color

    def set_color(self, color):
        self.color = color


class TriangleRenderer:
    def __init__(self, color='black'):
        self.color = color

    def render(self, base, height):
        print(f"Rendering triangle with base {base} and height {height} in {self.color}")

    def get_area(self, base, height):
        return 0.5 * base * height

    def get_color(self):
        return self.color

    def set_color(self, color):
        self.color = color

# Code duplication in color management and structure - should use inheritance!
