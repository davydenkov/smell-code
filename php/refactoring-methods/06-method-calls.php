<?php

/**
 * 42. Renaming a method (Rename Method)
 *
 * BEFORE: Poorly named method
 */
class CalculatorBefore {
    public function calc($a, $b) { // Unclear name
        return $a + $b;
    }
}

/**
 * AFTER: Rename method to be more descriptive
 */
class CalculatorAfter {
    public function add($a, $b) {
        return $a + $b;
    }
}

/**
 * 43. Adding a parameter (Add Parameter)
 *
 * BEFORE: Method missing required parameter
 */
class EmailSenderBefore {
    public function sendEmail($to, $subject, $body) {
        // Send email with default priority
        $priority = 'normal';
        // Send logic
    }
}

/**
 * AFTER: Add parameter
 */
class EmailSenderAfter {
    public function sendEmail($to, $subject, $body, $priority = 'normal') {
        // Send logic with priority
    }
}

/**
 * 44. Deleting a parameter (Remove Parameter)
 *
 * BEFORE: Unnecessary parameter
 */
class ReportGeneratorBefore {
    public function generateReport($data, $format, $includeHeader = true) {
        if ($format === 'html') {
            // Always include header for HTML
            $includeHeader = true;
        }
        // Generate report
    }
}

/**
 * AFTER: Remove unnecessary parameter
 */
class ReportGeneratorAfter {
    public function generateReport($data, $format) {
        $includeHeader = ($format === 'html');
        // Generate report
    }
}

/**
 * 45. Separation of Query and Modifier (Separate Query from Modifier)
 *
 * BEFORE: Method that both queries and modifies
 */
class BankAccountBefore {
    private $balance = 0;

    public function withdraw($amount) {
        if ($this->balance >= $amount) {
            $this->balance -= $amount;
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Separate query from modifier
 */
class BankAccountAfter {
    private $balance = 0;

    public function canWithdraw($amount) {
        return $this->balance >= $amount;
    }

    public function withdraw($amount) {
        if ($this->canWithdraw($amount)) {
            $this->balance -= $amount;
            return true;
        }
        return false;
    }
}

/**
 * 46. Parameterization of the method (Parameterize Method)
 *
 * BEFORE: Similar methods with different values
 */
class ReportGeneratorParamBefore {
    public function generateWeeklyReport() {
        return $this->generateReport(7);
    }

    public function generateMonthlyReport() {
        return $this->generateReport(30);
    }

    public function generateQuarterlyReport() {
        return $this->generateReport(90);
    }

    private function generateReport($days) {
        // Generate report for specified days
    }
}

/**
 * AFTER: Parameterize method
 */
class ReportGeneratorParamAfter {
    public function generateReport($days) {
        // Generate report for specified days
    }

    public function generateWeeklyReport() {
        return $this->generateReport(7);
    }

    public function generateMonthlyReport() {
        return $this->generateReport(30);
    }

    public function generateQuarterlyReport() {
        return $this->generateReport(90);
    }
}

/**
 * 47. Replacing a parameter with explicit methods (Replace Parameter with Explicit Methods)
 *
 * BEFORE: Parameter determines behavior
 */
class EmployeeExplicitBefore {
    const ENGINEER = 0;
    const SALESMAN = 1;
    const MANAGER = 2;

    private $type;

    public function __construct($type) {
        $this->type = $type;
    }

    public function getSalary($baseSalary) {
        switch ($this->type) {
            case self::ENGINEER:
                return $baseSalary * 1.0;
            case self::SALESMAN:
                return $baseSalary * 1.1;
            case self::MANAGER:
                return $baseSalary * 1.2;
            default:
                return $baseSalary;
        }
    }
}

/**
 * AFTER: Replace parameter with explicit methods
 */
class EmployeeExplicitAfter {
    public function getEngineerSalary($baseSalary) {
        return $baseSalary * 1.0;
    }

    public function getSalesmanSalary($baseSalary) {
        return $baseSalary * 1.1;
    }

    public function getManagerSalary($baseSalary) {
        return $baseSalary * 1.2;
    }
}

/**
 * 48. Save the Whole Object
 *
 * BEFORE: Passing individual fields
 */
class OrderWholeBefore {
    private $customer;

    public function __construct($customerName, $customerAddress) {
        $this->customer = ['name' => $customerName, 'address' => $customerAddress];
    }

    public function calculateShipping() {
        return $this->getShippingCost($this->customer['name'], $this->customer['address']);
    }

    private function getShippingCost($name, $address) {
        // Calculate based on name and address
        return 10.0;
    }
}

/**
 * AFTER: Pass whole object
 */
class Customer {
    private $name;
    private $address;

    public function __construct($name, $address) {
        $this->name = $name;
        $this->address = $address;
    }

    public function getName() {
        return $this->name;
    }

    public function getAddress() {
        return $this->address;
    }
}

class OrderWholeAfter {
    private $customer;

    public function __construct(Customer $customer) {
        $this->customer = $customer;
    }

    public function calculateShipping() {
        return $this->getShippingCost($this->customer);
    }

    private function getShippingCost(Customer $customer) {
        // Calculate based on customer object
        return 10.0;
    }
}

/**
 * 49. Replacing a parameter with a method call (Replace Parameter with Method)
 *
 * BEFORE: Parameter calculated outside method
 */
class DiscountCalculatorParamBefore {
    public function calculateDiscount($price, $customerType) {
        // customerType passed in
        return $price * $this->getDiscountRate($customerType);
    }

    private function getDiscountRate($customerType) {
        switch ($customerType) {
            case 'premium':
                return 0.1;
            case 'regular':
                return 0.05;
            default:
                return 0.0;
        }
    }
}

class OrderParamBefore {
    private $customer;

    public function getDiscountedPrice($price) {
        $calculator = new DiscountCalculatorParamBefore();
        return $calculator->calculateDiscount($price, $this->customer->getType());
    }
}

/**
 * AFTER: Replace parameter with method call
 */
class DiscountCalculatorParamAfter {
    public function calculateDiscount($price, $customer) {
        return $price * $this->getDiscountRate($customer->getType());
    }

    private function getDiscountRate($customerType) {
        switch ($customerType) {
            case 'premium':
                return 0.1;
            case 'regular':
                return 0.05;
            default:
                return 0.0;
        }
    }
}

class OrderParamAfter {
    private $customer;

    public function getDiscountedPrice($price) {
        $calculator = new DiscountCalculatorParamAfter();
        return $calculator->calculateDiscount($price, $this->customer);
    }
}

/**
 * 50. Introduction of the boundary object (Introduce Parameter Object)
 *
 * BEFORE: Multiple parameters
 */
class TemperatureRangeBefore {
    public function withinRange($minTemp, $maxTemp, $currentTemp) {
        return $currentTemp >= $minTemp && $currentTemp <= $maxTemp;
    }

    public function getAverageTemp($minTemp, $maxTemp) {
        return ($minTemp + $maxTemp) / 2;
    }
}

/**
 * AFTER: Introduce parameter object
 */
class TemperatureRange {
    private $minTemp;
    private $maxTemp;

    public function __construct($minTemp, $maxTemp) {
        $this->minTemp = $minTemp;
        $this->maxTemp = $maxTemp;
    }

    public function getMinTemp() {
        return $this->minTemp;
    }

    public function getMaxTemp() {
        return $this->maxTemp;
    }

    public function withinRange($currentTemp) {
        return $currentTemp >= $this->minTemp && $currentTemp <= $this->maxTemp;
    }

    public function getAverageTemp() {
        return ($this->minTemp + $this->maxTemp) / 2;
    }
}

/**
 * 51. Removing the Value Setting Method
 *
 * BEFORE: Setter that's not needed
 */
class SensorBefore {
    private $temperature;

    public function __construct($temperature) {
        $this->temperature = $temperature;
    }

    public function getTemperature() {
        return $this->temperature;
    }

    public function setTemperature($temperature) { // Not needed if immutable
        $this->temperature = $temperature;
    }
}

/**
 * AFTER: Remove setter for immutable object
 */
class SensorAfter {
    private $temperature;

    public function __construct($temperature) {
        $this->temperature = $temperature;
    }

    public function getTemperature() {
        return $this->temperature;
    }

    // setTemperature removed
}

/**
 * 52. Hiding a method (Hide Method)
 *
 * BEFORE: Public method that should be private
 */
class DataProcessorHideBefore {
    public function validateData($data) { // Should be private
        return !empty($data) && is_array($data);
    }

    public function processData($data) {
        if ($this->validateData($data)) {
            // Process data
        }
    }
}

/**
 * AFTER: Hide method
 */
class DataProcessorHideAfter {
    private function validateData($data) {
        return !empty($data) && is_array($data);
    }

    public function processData($data) {
        if ($this->validateData($data)) {
            // Process data
        }
    }
}

/**
 * 53. Replacing the constructor with the factory method (Replace Constructor with Factory Method)
 *
 * BEFORE: Complex constructor
 */
class ComplexObjectBefore {
    private $type;
    private $config;

    public function __construct($type, $config = []) {
        $this->type = $type;
        $this->config = $config;

        if ($type === 'database') {
            $this->config = array_merge(['host' => 'localhost', 'port' => 3306], $config);
        } elseif ($type === 'file') {
            $this->config = array_merge(['path' => '/tmp', 'format' => 'json'], $config);
        }
    }
}

/**
 * AFTER: Replace constructor with factory method
 */
class ComplexObjectAfter {
    private $type;
    private $config;

    private function __construct($type, $config) {
        $this->type = $type;
        $this->config = $config;
    }

    public static function createDatabaseConnection($config = []) {
        $config = array_merge(['host' => 'localhost', 'port' => 3306], $config);
        return new self('database', $config);
    }

    public static function createFileHandler($config = []) {
        $config = array_merge(['path' => '/tmp', 'format' => 'json'], $config);
        return new self('file', $config);
    }
}

/**
 * 54. Encapsulation of top-down type conversion (Encapsulate Downcast)
 *
 * BEFORE: Downcast in client code
 */
class ShapeCollectionBefore {
    private $shapes = [];

    public function addShape($shape) {
        $this->shapes[] = $shape;
    }

    public function getShapes() {
        return $this->shapes;
    }
}

// Client code
$collection = new ShapeCollectionBefore();
// ... add shapes
$circles = array_filter($collection->getShapes(), function($shape) {
    return $shape instanceof Circle; // Downcast check
});

/**
 * AFTER: Encapsulate downcast
 */
class ShapeCollectionAfter {
    private $shapes = [];

    public function addShape($shape) {
        $this->shapes[] = $shape;
    }

    public function getCircles() {
        return array_filter($this->shapes, function($shape) {
            return $shape instanceof Circle;
        });
    }

    public function getSquares() {
        return array_filter($this->shapes, function($shape) {
            return $shape instanceof Square;
        });
    }
}

/**
 * 55. Replacing the error code with an exceptional situation (Replace Error Code with Exception)
 *
 * BEFORE: Error codes
 */
class FileReaderErrorBefore {
    const FILE_NOT_FOUND = 1;
    const PERMISSION_DENIED = 2;

    public function readFile($filename) {
        if (!file_exists($filename)) {
            return self::FILE_NOT_FOUND;
        }

        if (!is_readable($filename)) {
            return self::PERMISSION_DENIED;
        }

        return file_get_contents($filename);
    }
}

// Client code
$reader = new FileReaderErrorBefore();
$result = $reader->readFile('test.txt');
if ($result === FileReaderErrorBefore::FILE_NOT_FOUND) {
    // Handle error
} elseif ($result === FileReaderErrorBefore::PERMISSION_DENIED) {
    // Handle error
} else {
    // Use content
}

/**
 * AFTER: Replace error codes with exceptions
 */
class FileReaderExceptionAfter {
    public function readFile($filename) {
        if (!file_exists($filename)) {
            throw new FileNotFoundException("File not found: $filename");
        }

        if (!is_readable($filename)) {
            throw new PermissionDeniedException("Permission denied: $filename");
        }

        return file_get_contents($filename);
    }
}

class FileNotFoundException extends Exception {}
class PermissionDeniedException extends Exception {}

// Client code
try {
    $reader = new FileReaderExceptionAfter();
    $content = $reader->readFile('test.txt');
    // Use content
} catch (FileNotFoundException $e) {
    // Handle file not found
} catch (PermissionDeniedException $e) {
    // Handle permission denied
}

/**
 * 56. Replacing an exceptional situation with a check (Replace Exception with Test)
 *
 * BEFORE: Using exception for control flow
 */
class StackExceptionBefore {
    private $items = [];

    public function pop() {
        if (empty($this->items)) {
            throw new EmptyStackException();
        }
        return array_pop($this->items);
    }
}

class EmptyStackException extends Exception {}

// Client code
$stack = new StackExceptionBefore();
try {
    $item = $stack->pop();
} catch (EmptyStackException $e) {
    $item = null; // Default value
}

/**
 * AFTER: Replace exception with test
 */
class StackTestAfter {
    private $items = [];

    public function isEmpty() {
        return empty($this->items);
    }

    public function pop() {
        return array_pop($this->items);
    }
}

// Client code
$stack = new StackTestAfter();
$item = $stack->isEmpty() ? null : $stack->pop();
