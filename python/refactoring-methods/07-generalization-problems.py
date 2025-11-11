import math
from abc import ABC, abstractmethod

# 57. Lifting the field (Pull Up Field)
#
# BEFORE: Duplicate fields in subclasses
class EmployeePullBefore(ABC):
    def __init__(self, name):
        self.name = name

class ManagerPullBefore(EmployeePullBefore):
    def __init__(self, name, budget):
        super().__init__(name)
        self.name = name  # Duplicate
        self._budget = budget

class EngineerPullBefore(EmployeePullBefore):
    def __init__(self, name, skills):
        super().__init__(name)
        self.name = name  # Duplicate
        self._skills = skills

# AFTER: Pull up field to superclass
class EmployeePullAfter(ABC):
    def __init__(self, name):
        self.name = name

class ManagerPullAfter(EmployeePullAfter):
    def __init__(self, name, budget):
        super().__init__(name)
        self._budget = budget

class EngineerPullAfter(EmployeePullAfter):
    def __init__(self, name, skills):
        super().__init__(name)
        self._skills = skills

# 58. Lifting the method (Pull Up Method)
#
# BEFORE: Duplicate methods in subclasses
class ShapeMethodBefore(ABC):
    @abstractmethod
    def area(self):
        pass

class CircleMethodBefore(ShapeMethodBefore):
    def __init__(self, radius):
        self._radius = radius

    def area(self):
        return math.pi * self._radius * self._radius

    def circumference(self):
        return 2 * math.pi * self._radius

class SquareMethodBefore(ShapeMethodBefore):
    def __init__(self, side):
        self._side = side

    def area(self):
        return self._side * self._side

    def circumference(self):  # Similar logic could be generalized
        return 4 * self._side

# AFTER: Pull up method to superclass
class ShapeMethodAfter(ABC):
    @abstractmethod
    def area(self):
        pass

    @abstractmethod
    def circumference(self):
        pass

class CircleMethodAfter(ShapeMethodAfter):
    def __init__(self, radius):
        self._radius = radius

    def area(self):
        return math.pi * self._radius * self._radius

    def circumference(self):
        return 2 * math.pi * self._radius

class SquareMethodAfter(ShapeMethodAfter):
    def __init__(self, side):
        self._side = side

    def area(self):
        return self._side * self._side

    def circumference(self):
        return 4 * self._side

# 59. Lifting the constructor Body (Pull Up Constructor Body)
#
# BEFORE: Duplicate constructor code
class VehicleConstructorBefore(ABC):
    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year

class CarConstructorBefore(VehicleConstructorBefore):
    def __init__(self, make, model, year, doors):
        self.make = make  # Duplicate
        self.model = model  # Duplicate
        self.year = year  # Duplicate
        self._doors = doors

class TruckConstructorBefore(VehicleConstructorBefore):
    def __init__(self, make, model, year, payload):
        self.make = make  # Duplicate
        self.model = model  # Duplicate
        self.year = year  # Duplicate
        self._payload = payload

# AFTER: Pull up constructor body
class VehicleConstructorAfter(ABC):
    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year

class CarConstructorAfter(VehicleConstructorAfter):
    def __init__(self, make, model, year, doors):
        super().__init__(make, model, year)
        self._doors = doors

class TruckConstructorAfter(VehicleConstructorAfter):
    def __init__(self, make, model, year, payload):
        super().__init__(make, model, year)
        self._payload = payload

# 60. Method Descent (Push Down Method)
#
# BEFORE: Method in wrong class hierarchy level
class AnimalPushBefore(ABC):
    def speak(self):
        return "Generic sound"

class DogPushBefore(AnimalPushBefore):
    def speak(self):
        return 'Woof'

class CatPushBefore(AnimalPushBefore):
    def speak(self):
        return 'Meow'

class FishPushBefore(AnimalPushBefore):
    pass  # Fish don't speak, but inherits speak method

# AFTER: Push down method to appropriate subclasses
class AnimalPushAfter(ABC):
    pass  # No speak method here

class DogPushAfter(AnimalPushAfter):
    def speak(self):
        return 'Woof'

class CatPushAfter(AnimalPushAfter):
    def speak(self):
        return 'Meow'

class FishPushAfter(AnimalPushAfter):
    pass  # Fish don't speak

# 61. Field Descent (Push Down Field)
# Similar concept - moving fields down to subclasses that actually use them

# 62. Subclass extraction (Extract Subclass)
class EmployeeExtractBefore:
    def __init__(self, name, type_code):
        self._name = name
        self._type_code = type_code

    def get_name(self):
        return self._name

    def get_type_code(self):
        return self._type_code

# AFTER: Extract subclasses
class EmployeeExtractAfter(ABC):
    def __init__(self, name):
        self._name = name

    def get_name(self):
        return self._name

    @abstractmethod
    def get_type(self):
        pass

class ManagerExtract(EmployeeExtractAfter):
    def get_type(self):
        return 'manager'

class EngineerExtract(EmployeeExtractAfter):
    def get_type(self):
        return 'engineer'

# 63. Allocation of the parent class (Extract Superclass)
# Similar to the examples above - creating common superclass

# 64. Interface extraction (Extract Interface)
class WorkerInterface(ABC):
    @abstractmethod
    def work(self):
        pass

    @abstractmethod
    def get_salary(self):
        pass

# 65. Collapse Hierarchy
# Removing unnecessary inheritance levels

# 66. Formation of the method template (Form Template Method)
class DocumentTemplate(ABC):
    def process_document(self):
        self.open_document()
        self.read_content()
        self.process_content()
        self.save_document()
        self.close_document()

    def open_document(self):
        print("Opening document")

    def save_document(self):
        print("Saving document")

    def close_document(self):
        print("Closing document")

    @abstractmethod
    def read_content(self):
        pass

    @abstractmethod
    def process_content(self):
        pass

class PDFDocument(DocumentTemplate):
    def read_content(self):
        print("Reading PDF content")

    def process_content(self):
        print("Processing PDF content")

class WordDocument(DocumentTemplate):
    def read_content(self):
        print("Reading Word content")

    def process_content(self):
        print("Processing Word content")

# 67. Replacement of inheritance by delegation
class StackInheritance:
    def __init__(self):
        self._items = []

    def push(self, item):
        self._items.append(item)

    def pop(self):
        return self._items.pop()

# AFTER: Replace with delegation
class StackDelegation:
    def __init__(self):
        self._list = []

    def push(self, item):
        self._list.append(item)

    def pop(self):
        return self._list.pop()

# 68. Replacement of delegation by inheritance
# Opposite - using inheritance instead of delegation
