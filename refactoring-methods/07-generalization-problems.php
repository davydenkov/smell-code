<?php

/**
 * 57. Lifting the field (Pull Up Field)
 *
 * BEFORE: Duplicate fields in subclasses
 */
class EmployeePullBefore {
    protected $name;
}

class ManagerPullBefore extends EmployeePullBefore {
    protected $name; // Duplicate
    private $budget;
}

class EngineerPullBefore extends EmployeePullBefore {
    protected $name; // Duplicate
    private $skills;
}

/**
 * AFTER: Pull up field to superclass
 */
class EmployeePullAfter {
    protected $name;
}

class ManagerPullAfter extends EmployeePullAfter {
    private $budget;
}

class EngineerPullAfter extends EmployeePullAfter {
    private $skills;
}

/**
 * 58. Lifting the method (Pull Up Method)
 *
 * BEFORE: Duplicate methods in subclasses
 */
class ShapeMethodBefore {
    abstract public function area();
}

class CircleMethodBefore extends ShapeMethodBefore {
    private $radius;

    public function area() {
        return pi() * $this->radius * $this->radius;
    }

    public function circumference() {
        return 2 * pi() * $this->radius;
    }
}

class SquareMethodBefore extends ShapeMethodBefore {
    private $side;

    public function area() {
        return $this->side * $this->side;
    }

    public function circumference() { // Duplicate logic
        return 4 * $this->side;
    }
}

/**
 * AFTER: Pull up method to superclass
 */
abstract class ShapeMethodAfter {
    abstract public function area();
    abstract public function circumference();
}

class CircleMethodAfter extends ShapeMethodAfter {
    private $radius;

    public function area() {
        return pi() * $this->radius * $this->radius;
    }

    public function circumference() {
        return 2 * pi() * $this->radius;
    }
}

class SquareMethodAfter extends ShapeMethodAfter {
    private $side;

    public function area() {
        return $this->side * $this->side;
    }

    public function circumference() {
        return 4 * $this->side;
    }
}

/**
 * 59. Lifting the constructor Body (Pull Up Constructor Body)
 *
 * BEFORE: Duplicate constructor code
 */
class VehicleConstructorBefore {
    protected $make;
    protected $model;
    protected $year;
}

class CarConstructorBefore extends VehicleConstructorBefore {
    private $doors;

    public function __construct($make, $model, $year, $doors) {
        $this->make = $make; // Duplicate
        $this->model = $model; // Duplicate
        $this->year = $year; // Duplicate
        $this->doors = $doors;
    }
}

class TruckConstructorBefore extends VehicleConstructorBefore {
    private $payload;

    public function __construct($make, $model, $year, $payload) {
        $this->make = $make; // Duplicate
        $this->model = $model; // Duplicate
        $this->year = $year; // Duplicate
        $this->payload = $payload;
    }
}

/**
 * AFTER: Pull up constructor body
 */
class VehicleConstructorAfter {
    protected $make;
    protected $model;
    protected $year;

    public function __construct($make, $model, $year) {
        $this->make = $make;
        $this->model = $model;
        $this->year = $year;
    }
}

class CarConstructorAfter extends VehicleConstructorAfter {
    private $doors;

    public function __construct($make, $model, $year, $doors) {
        parent::__construct($make, $model, $year);
        $this->doors = $doors;
    }
}

class TruckConstructorAfter extends VehicleConstructorAfter {
    private $payload;

    public function __construct($make, $model, $year, $payload) {
        parent::__construct($make, $model, $year);
        $this->payload = $payload;
    }
}

/**
 * 60. Method Descent (Push Down Method)
 *
 * BEFORE: Method in wrong class hierarchy level
 */
class AnimalPushBefore {
    public function speak() {
        // Generic implementation
    }
}

class DogPushBefore extends AnimalPushBefore {
    public function speak() {
        return 'Woof';
    }
}

class CatPushBefore extends AnimalPushBefore {
    public function speak() {
        return 'Meow';
    }
}

class FishPushBefore extends AnimalPushBefore {
    // Fish don't speak, but inherits speak method
}

/**
 * AFTER: Push down method to appropriate subclasses
 */
class AnimalPushAfter {
    // No speak method here
}

class DogPushAfter extends AnimalPushAfter {
    public function speak() {
        return 'Woof';
    }
}

class CatPushAfter extends AnimalPushAfter {
    public function speak() {
        return 'Meow';
    }
}

class FishPushAfter extends AnimalPushAfter {
    // No speak method - appropriate for Fish
}

/**
 * 61. Field Descent (Push Down Field)
 *
 * BEFORE: Field in wrong hierarchy level
 */
class EmployeeFieldBefore {
    protected $salary; // Not all employees have salary
}

class SalariedEmployeeFieldBefore extends EmployeeFieldBefore {
    // Uses salary
}

class ContractorFieldBefore extends EmployeeFieldBefore {
    // Doesn't use salary, but inherits it
}

/**
 * AFTER: Push down field
 */
class EmployeeFieldAfter {
    // No salary field
}

class SalariedEmployeeFieldAfter extends EmployeeFieldAfter {
    protected $salary;
}

class ContractorFieldAfter extends EmployeeFieldAfter {
    protected $hourlyRate;
}

/**
 * 62. Subclass extraction (Extract Subclass)
 *
 * BEFORE: Class with conditional behavior
 */
class JobExtractBefore {
    private $type;
    private $rate;
    private $commission;

    public function __construct($type, $rate, $commission = null) {
        $this->type = $type;
        $this->rate = $rate;
        $this->commission = $commission;
    }

    public function getPay() {
        if ($this->type === 'salaried') {
            return $this->rate;
        } else {
            return $this->rate + $this->commission;
        }
    }
}

/**
 * AFTER: Extract subclass
 */
abstract class JobExtractAfter {
    protected $rate;

    public function __construct($rate) {
        $this->rate = $rate;
    }

    abstract public function getPay();
}

class SalariedJob extends JobExtractAfter {
    public function getPay() {
        return $this->rate;
    }
}

class CommissionedJob extends JobExtractAfter {
    private $commission;

    public function __construct($rate, $commission) {
        parent::__construct($rate);
        $this->commission = $commission;
    }

    public function getPay() {
        return $this->rate + $this->commission;
    }
}

/**
 * 63. Allocation of the parent class (Extract Superclass)
 *
 * BEFORE: Duplicate code in classes
 */
class DepartmentSuperBefore {
    private $name;
    private $head;

    public function __construct($name, $head) {
        $this->name = $name;
        $this->head = $head;
    }

    public function getName() {
        return $this->name;
    }

    public function getHead() {
        return $this->head;
    }
}

class CompanySuperBefore {
    private $name;
    private $head;

    public function __construct($name, $head) {
        $this->name = $name;
        $this->head = $head;
    }

    public function getName() {
        return $this->name;
    }

    public function getHead() {
        return $this->head;
    }
}

/**
 * AFTER: Extract superclass
 */
abstract class Party {
    private $name;
    private $head;

    public function __construct($name, $head) {
        $this->name = $name;
        $this->head = $head;
    }

    public function getName() {
        return $this->name;
    }

    public function getHead() {
        return $this->head;
    }
}

class DepartmentSuperAfter extends Party {
}

class CompanySuperAfter extends Party {
}

/**
 * 64. Interface extraction (Extract Interface)
 *
 * BEFORE: Clients depend on concrete class
 */
class PrinterInterfaceBefore {
    public function print($document) {
        // Print logic
    }

    public function getStatus() {
        // Status logic
    }

    public function cancelJob($jobId) {
        // Cancel logic
    }
}

/**
 * AFTER: Extract interface
 */
interface Printer {
    public function print($document);
    public function getStatus();
}

class LaserPrinter implements Printer {
    public function print($document) {
        // Print logic
    }

    public function getStatus() {
        // Status logic
    }

    public function cancelJob($jobId) {
        // Cancel logic - not part of interface
    }
}

class InkjetPrinter implements Printer {
    public function print($document) {
        // Print logic
    }

    public function getStatus() {
        // Status logic
    }
}

/**
 * 65. Collapse Hierarchy
 *
 * BEFORE: Unnecessary class hierarchy
 */
class EmployeeCollapseBefore {
}

class ManagerCollapseBefore extends EmployeeCollapseBefore {
    private $department;
}

/**
 * AFTER: Collapse hierarchy if only one subclass
 */
class EmployeeCollapseAfter {
    private $department; // Moved up
}

/**
 * 66. Formation of the method template (Form Template Method)
 *
 * BEFORE: Duplicate algorithm structure
 */
class ReportGeneratorTemplateBefore {
    public function generateHTMLReport() {
        $data = $this->getData();
        $header = $this->formatHeader();
        $body = $this->formatBody($data);
        $footer = $this->formatFooter();
        return $header . $body . $footer;
    }

    public function generatePDFReport() {
        $data = $this->getData(); // Duplicate
        $header = $this->formatPDFHeader(); // Different
        $body = $this->formatPDFBody($data); // Different
        $footer = $this->formatPDFFooter(); // Different
        return $header . $body . $footer;
    }

    protected function getData() {
        return ['item1', 'item2'];
    }

    protected function formatHeader() {
        return '<h1>Report</h1>';
    }

    protected function formatBody($data) {
        return '<body>' . implode('', $data) . '</body>';
    }

    protected function formatFooter() {
        return '<footer>End</footer>';
    }

    protected function formatPDFHeader() {
        return 'PDF Report Header';
    }

    protected function formatPDFBody($data) {
        return 'PDF Body: ' . implode('', $data);
    }

    protected function formatPDFFooter() {
        return 'PDF Footer';
    }
}

/**
 * AFTER: Form template method
 */
abstract class ReportGeneratorTemplateAfter {
    public final function generateReport() {
        $data = $this->getData();
        $header = $this->formatHeader();
        $body = $this->formatBody($data);
        $footer = $this->formatFooter();
        return $this->assembleReport($header, $body, $footer);
    }

    protected function getData() {
        return ['item1', 'item2'];
    }

    abstract protected function formatHeader();
    abstract protected function formatBody($data);
    abstract protected function formatFooter();
    abstract protected function assembleReport($header, $body, $footer);
}

class HTMLReportGenerator extends ReportGeneratorTemplateAfter {
    protected function formatHeader() {
        return '<h1>Report</h1>';
    }

    protected function formatBody($data) {
        return '<body>' . implode('', $data) . '</body>';
    }

    protected function formatFooter() {
        return '<footer>End</footer>';
    }

    protected function assembleReport($header, $body, $footer) {
        return $header . $body . $footer;
    }
}

class PDFReportGenerator extends ReportGeneratorTemplateAfter {
    protected function formatHeader() {
        return 'PDF Report Header';
    }

    protected function formatBody($data) {
        return 'PDF Body: ' . implode('', $data);
    }

    protected function formatFooter() {
        return 'PDF Footer';
    }

    protected function assembleReport($header, $body, $footer) {
        return $header . $body . $footer;
    }
}

/**
 * 67. Replacement of inheritance by delegation (Replace Inheritance with Delegation)
 *
 * BEFORE: Inheritance where delegation would be better
 */
class StackInheritanceBefore extends ArrayObject {
    public function push($item) {
        $this->append($item);
    }

    public function pop() {
        if ($this->count() === 0) {
            throw new Exception('Stack is empty');
        }
        $lastIndex = $this->count() - 1;
        $item = $this[$lastIndex];
        $this->offsetUnset($lastIndex);
        return $item;
    }
}

/**
 * AFTER: Replace inheritance with delegation
 */
class StackDelegationAfter {
    private $items;

    public function __construct() {
        $this->items = new ArrayObject();
    }

    public function push($item) {
        $this->items->append($item);
    }

    public function pop() {
        if ($this->items->count() === 0) {
            throw new Exception('Stack is empty');
        }
        $lastIndex = $this->items->count() - 1;
        $item = $this->items[$lastIndex];
        $this->items->offsetUnset($lastIndex);
        return $item;
    }

    public function count() {
        return $this->items->count();
    }
}

/**
 * 68. Replacement of delegation by inheritance (Replace Delegation with Inheritance)
 *
 * BEFORE: Delegation where inheritance would be simpler
 */
class MyStringDelegateBefore {
    private $string;

    public function __construct($string) {
        $this->string = $string;
    }

    public function length() {
        return strlen($this->string);
    }

    public function substr($start, $length = null) {
        return substr($this->string, $start, $length);
    }

    public function strpos($needle) {
        return strpos($this->string, $needle);
    }
}

/**
 * AFTER: Replace delegation with inheritance
 */
class MyStringInheritAfter extends ArrayObject {
    public function __construct($string) {
        parent::__construct(str_split($string));
    }

    public function toString() {
        return implode('', $this->getArrayCopy());
    }
}
