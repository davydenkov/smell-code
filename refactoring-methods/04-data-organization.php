<?php

/**
 * 18. Self-Encapsulate Field
 *
 * BEFORE: Direct field access
 */
class PersonBefore {
    public $name; // Direct access

    public function getName() {
        return $this->name;
    }

    public function setName($name) {
        $this->name = $name;
    }
}

/**
 * AFTER: Self-encapsulate field
 */
class PersonAfter {
    private $name;

    public function getName() {
        return $this->name;
    }

    public function setName($name) {
        $this->name = $name;
    }
}

/**
 * 19. Replacing the data value with an object (Replace Data Value with Object)
 *
 * BEFORE: Primitive data type that should be an object
 */
class OrderBefore {
    private $customer; // Just a string

    public function getCustomerName() {
        return $this->customer;
    }

    public function setCustomer($customer) {
        $this->customer = $customer;
    }
}

/**
 * AFTER: Replace with object
 */
class Customer {
    private $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function getName() {
        return $this->name;
    }
}

class OrderAfter {
    private $customer;

    public function getCustomer() {
        return $this->customer;
    }

    public function setCustomer(Customer $customer) {
        $this->customer = $customer;
    }

    public function getCustomerName() {
        return $this->customer->getName();
    }
}

/**
 * 20. Replacing the value with a reference (Change Value to Reference)
 *
 * BEFORE: Multiple instances of same object
 */
class CustomerValue {
    private $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function getName() {
        return $this->name;
    }
}

class OrderValue {
    private $customer; // New instance for each order

    public function __construct($customerName) {
        $this->customer = new CustomerValue($customerName);
    }
}

/**
 * AFTER: Use reference to single instance
 */
class CustomerReference {
    private $name;
    private static $instances = [];

    private function __construct($name) {
        $this->name = $name;
    }

    public static function create($name) {
        if (!isset(self::$instances[$name])) {
            self::$instances[$name] = new self($name);
        }
        return self::$instances[$name];
    }

    public function getName() {
        return $this->name;
    }
}

class OrderReference {
    private $customer;

    public function __construct($customerName) {
        $this->customer = CustomerReference::create($customerName);
    }
}

/**
 * 21. Replacing a reference with a value (Change Reference to Value)
 *
 * BEFORE: Unnecessary reference when value would suffice
 */
class CurrencyReference {
    private $code;

    public function __construct($code) {
        $this->code = $code;
    }

    public function getCode() {
        return $this->code;
    }
}

class ProductReference {
    private $price;
    private $currency; // Reference object

    public function __construct($price, CurrencyReference $currency) {
        $this->price = $price;
        $this->currency = $currency;
    }
}

/**
 * AFTER: Use value object instead
 */
class CurrencyValue {
    private $code;

    public function __construct($code) {
        $this->code = $code;
    }

    public function getCode() {
        return $this->code;
    }
}

class ProductValue {
    private $price;
    private $currencyCode; // Just the value

    public function __construct($price, $currencyCode) {
        $this->price = $price;
        $this->currencyCode = $currencyCode;
    }

    public function getCurrencyCode() {
        return $this->currencyCode;
    }
}

/**
 * 22. Replacing an array with an object (Replace Array with Object)
 *
 * BEFORE: Using array for structured data
 */
class PerformanceArray {
    public function getPerformanceData() {
        return [
            'goals' => 10,
            'assists' => 5,
            'minutes' => 120
        ];
    }

    public function calculateScore($data) {
        return ($data['goals'] * 2) + ($data['assists'] * 1.5) + ($data['minutes'] / 60);
    }
}

/**
 * AFTER: Replace array with object
 */
class PerformanceData {
    private $goals;
    private $assists;
    private $minutes;

    public function __construct($goals, $assists, $minutes) {
        $this->goals = $goals;
        $this->assists = $assists;
        $this->minutes = $minutes;
    }

    public function getGoals() {
        return $this->goals;
    }

    public function getAssists() {
        return $this->assists;
    }

    public function getMinutes() {
        return $this->minutes;
    }

    public function calculateScore() {
        return ($this->goals * 2) + ($this->assists * 1.5) + ($this->minutes / 60);
    }
}

class PerformanceObject {
    public function getPerformanceData() {
        return new PerformanceData(10, 5, 120);
    }

    public function calculateScore(PerformanceData $data) {
        return $data->calculateScore();
    }
}

/**
 * 23. Duplication of visible data (Duplicate Observed Data)
 *
 * BEFORE: Domain data mixed with presentation
 */
class OrderDomain {
    private $total = 0;

    public function addItem($price) {
        $this->total += $price;
        // Have to update UI here too
        $this->updateDisplay();
    }

    private function updateDisplay() {
        // Update UI elements
    }
}

/**
 * AFTER: Separate domain and presentation data
 */
class OrderDomainSeparated {
    private $total = 0;
    private $observers = [];

    public function addItem($price) {
        $this->total += $price;
        $this->notifyObservers();
    }

    public function getTotal() {
        return $this->total;
    }

    public function addObserver($observer) {
        $this->observers[] = $observer;
    }

    private function notifyObservers() {
        foreach ($this->observers as $observer) {
            $observer->update($this->total);
        }
    }
}

class OrderDisplay {
    private $order;

    public function __construct(OrderDomainSeparated $order) {
        $this->order = $order;
        $this->order->addObserver($this);
    }

    public function update($total) {
        // Update display with new total
    }
}

/**
 * 24. Replacing Unidirectional communication with Bidirectional
 * communication (Change Unidirectional Association to Bidirectional)
 *
 * BEFORE: One-way association
 */
class CustomerUni {
    private $orders = [];

    public function addOrder($order) {
        $this->orders[] = $order;
        // Order doesn't know about customer
    }
}

class OrderUni {
    private $items = [];
}

/**
 * AFTER: Bidirectional association
 */
class CustomerBi {
    private $orders = [];

    public function addOrder(OrderBi $order) {
        $this->orders[] = $order;
        $order->setCustomer($this);
    }
}

class OrderBi {
    private $customer;
    private $items = [];

    public function setCustomer(CustomerBi $customer) {
        $this->customer = $customer;
    }

    public function getCustomer() {
        return $this->customer;
    }
}

/**
 * 25. Replacing Bidirectional communication with Unidirectional
 * communication (Change Bidirectional Association to Unidirectional)
 *
 * BEFORE: Unnecessary bidirectional association
 */
class CustomerBidirectional {
    private $orders = [];

    public function addOrder(OrderBidirectional $order) {
        $this->orders[] = $order;
        $order->setCustomer($this);
    }
}

class OrderBidirectional {
    private $customer;

    public function setCustomer(CustomerBidirectional $customer) {
        $this->customer = $customer;
    }

    public function getCustomer() {
        return $this->customer;
    }
}

/**
 * AFTER: Remove bidirectional link
 */
class CustomerUnidirectional {
    private $orders = [];

    public function addOrder(OrderUnidirectional $order) {
        $this->orders[] = $order;
    }
}

class OrderUnidirectional {
    private $customerId;

    public function __construct($customerId) {
        $this->customerId = $customerId;
    }

    public function getCustomerId() {
        return $this->customerId;
    }
}

/**
 * 26. Replacing the magic number with a symbolic constant
 * (Replace Magic Number with Symbolic Constant)
 *
 * BEFORE: Magic numbers
 */
class GeometryBefore {
    public function calculateCircleArea($radius) {
        return 3.14159 * $radius * $radius; // Magic number
    }

    public function calculateCircleCircumference($radius) {
        return 2 * 3.14159 * $radius; // Same magic number
    }
}

/**
 * AFTER: Use symbolic constant
 */
class GeometryAfter {
    const PI = 3.14159;

    public function calculateCircleArea($radius) {
        return self::PI * $radius * $radius;
    }

    public function calculateCircleCircumference($radius) {
        return 2 * self::PI * $radius;
    }
}

/**
 * 27. Encapsulate Field
 *
 * BEFORE: Public field
 */
class PersonPublic {
    public $name;
}

/**
 * AFTER: Encapsulated field
 */
class PersonEncapsulated {
    private $name;

    public function getName() {
        return $this->name;
    }

    public function setName($name) {
        $this->name = $name;
    }
}

/**
 * 28. Encapsulate Collection
 *
 * BEFORE: Direct access to collection
 */
class TeamBefore {
    public $players = []; // Direct access

    public function addPlayer($player) {
        $this->players[] = $player;
    }
}

/**
 * AFTER: Encapsulated collection
 */
class TeamAfter {
    private $players = [];

    public function addPlayer($player) {
        $this->players[] = $player;
    }

    public function removePlayer($player) {
        $key = array_search($player, $this->players);
        if ($key !== false) {
            unset($this->players[$key]);
        }
    }

    public function getPlayers() {
        return array_values($this->players); // Return copy
    }

    public function getPlayerCount() {
        return count($this->players);
    }
}

/**
 * 29. Replacing a record with a Data Class
 *
 * BEFORE: Using array as data structure
 */
class EmployeeArray {
    public function createEmployee($data) {
        return [
            'name' => $data['name'],
            'salary' => $data['salary'],
            'department' => $data['department']
        ];
    }

    public function getSalary($employee) {
        return $employee['salary'];
    }
}

/**
 * AFTER: Use data class
 */
class Employee {
    private $name;
    private $salary;
    private $department;

    public function __construct($name, $salary, $department) {
        $this->name = $name;
        $this->salary = $salary;
        $this->department = $department;
    }

    public function getName() {
        return $this->name;
    }

    public function getSalary() {
        return $this->salary;
    }

    public function getDepartment() {
        return $this->department;
    }
}

class EmployeeDataClass {
    public function createEmployee($name, $salary, $department) {
        return new Employee($name, $salary, $department);
    }

    public function getSalary(Employee $employee) {
        return $employee->getSalary();
    }
}

/**
 * 30. Replacing Type Code with Class
 *
 * BEFORE: Type code as constants
 */
class EmployeeTypeCode {
    const ENGINEER = 0;
    const SALESMAN = 1;
    const MANAGER = 2;

    private $type;

    public function __construct($type) {
        $this->type = $type;
    }

    public function getTypeCode() {
        return $this->type;
    }

    public function getMonthlySalary() {
        switch ($this->type) {
            case self::ENGINEER:
                return 5000;
            case self::SALESMAN:
                return 4000;
            case self::MANAGER:
                return 6000;
            default:
                return 0;
        }
    }
}

/**
 * AFTER: Replace type code with class
 */
abstract class EmployeeType {
    abstract public function getMonthlySalary();

    public static function createEngineer() {
        return new EngineerType();
    }

    public static function createSalesman() {
        return new SalesmanType();
    }

    public static function createManager() {
        return new ManagerType();
    }
}

class EngineerType extends EmployeeType {
    public function getMonthlySalary() {
        return 5000;
    }
}

class SalesmanType extends EmployeeType {
    public function getMonthlySalary() {
        return 4000;
    }
}

class ManagerType extends EmployeeType {
    public function getMonthlySalary() {
        return 6000;
    }
}

class EmployeeTypeClass {
    private $type;

    public function __construct(EmployeeType $type) {
        $this->type = $type;
    }

    public function getMonthlySalary() {
        return $this->type->getMonthlySalary();
    }
}

/**
 * 31. Replacing Type Code with Subclasses
 *
 * BEFORE: Type code in base class
 */
class EmployeeSubBefore {
    const ENGINEER = 0;
    const SALESMAN = 1;
    const MANAGER = 2;

    private $type;
    private $salary;

    public function __construct($type, $salary) {
        $this->type = $type;
        $this->salary = $salary;
    }

    public function getSalary() {
        return $this->salary;
    }

    public function getType() {
        return $this->type;
    }
}

/**
 * AFTER: Replace type code with subclasses
 */
abstract class EmployeeSubAfter {
    protected $salary;

    public function __construct($salary) {
        $this->salary = $salary;
    }

    public function getSalary() {
        return $this->salary;
    }

    abstract public function getType();
}

class Engineer extends EmployeeSubAfter {
    public function getType() {
        return 'engineer';
    }
}

class Salesman extends EmployeeSubAfter {
    public function getType() {
        return 'salesman';
    }
}

class Manager extends EmployeeSubAfter {
    public function getType() {
        return 'manager';
    }
}

/**
 * 32. Replacing Type Code with State/Strategy
 *
 * BEFORE: Type code with behavior
 */
class EmployeeStateBefore {
    const JUNIOR = 0;
    const SENIOR = 1;
    const LEAD = 2;

    private $level;

    public function __construct($level) {
        $this->level = $level;
    }

    public function getSalaryMultiplier() {
        switch ($this->level) {
            case self::JUNIOR:
                return 1.0;
            case self::SENIOR:
                return 1.5;
            case self::LEAD:
                return 2.0;
            default:
                return 1.0;
        }
    }
}

/**
 * AFTER: Use state/strategy pattern
 */
interface EmployeeLevel {
    public function getSalaryMultiplier();
}

class JuniorLevel implements EmployeeLevel {
    public function getSalaryMultiplier() {
        return 1.0;
    }
}

class SeniorLevel implements EmployeeLevel {
    public function getSalaryMultiplier() {
        return 1.5;
    }
}

class LeadLevel implements EmployeeLevel {
    public function getSalaryMultiplier() {
        return 2.0;
    }
}

class EmployeeStateAfter {
    private $level;

    public function __construct(EmployeeLevel $level) {
        $this->level = $level;
    }

    public function getSalaryMultiplier() {
        return $this->level->getSalaryMultiplier();
    }
}

/**
 * 33. Replacing Subclass with Fields
 *
 * BEFORE: Unnecessary subclasses
 */
abstract class PersonSub {
    protected $name;
    protected $gender;

    public function __construct($name, $gender) {
        $this->name = $name;
        $this->gender = $gender;
    }

    public function getName() {
        return $this->name;
    }

    abstract public function isMale();
}

class Male extends PersonSub {
    public function isMale() {
        return true;
    }
}

class Female extends PersonSub {
    public function isMale() {
        return false;
    }
}

/**
 * AFTER: Replace subclass with field
 */
class PersonField {
    private $name;
    private $gender; // 'male' or 'female'

    public function __construct($name, $gender) {
        $this->name = $name;
        $this->gender = $gender;
    }

    public function getName() {
        return $this->name;
    }

    public function isMale() {
        return $this->gender === 'male';
    }

    public function getGender() {
        return $this->gender;
    }
}
