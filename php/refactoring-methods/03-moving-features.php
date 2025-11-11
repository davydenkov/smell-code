<?php

/**
 * 9. Substitution Algorithm
 *
 * BEFORE: Complex algorithm that can be simplified
 */
class PricingServiceBefore {
    public function calculatePrice($items) {
        $total = 0;
        foreach ($items as $item) {
            if ($item['type'] === 'book') {
                $total += $item['price'] * 0.9; // 10% discount for books
            } elseif ($item['type'] === 'electronics') {
                $total += $item['price'] * 1.1; // 10% markup for electronics
            } else {
                $total += $item['price'];
            }
        }
        return $total;
    }
}

/**
 * AFTER: Substitute with a simpler algorithm
 */
class PricingServiceAfter {
    private $discounts = [
        'book' => 0.9,
        'electronics' => 1.1,
        'default' => 1.0
    ];

    public function calculatePrice($items) {
        $total = 0;
        foreach ($items as $item) {
            $multiplier = $this->discounts[$item['type']] ?? $this->discounts['default'];
            $total += $item['price'] * $multiplier;
        }
        return $total;
    }
}

/**
 * 10. Moving functions between objects (Move Method)
 *
 * BEFORE: Method in wrong class
 */
class AccountBefore {
    private $balance;

    public function __construct($balance) {
        $this->balance = $balance;
    }

    public function getBalance() {
        return $this->balance;
    }

    // This method belongs in Bank class, not Account
    public function transferTo(Account $target, $amount) {
        if ($this->balance >= $amount) {
            $this->balance -= $amount;
            $target->balance += $amount;
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Move method to appropriate class
 */
class AccountAfter {
    private $balance;

    public function __construct($balance) {
        $this->balance = $balance;
    }

    public function getBalance() {
        return $this->balance;
    }

    public function decreaseBalance($amount) {
        $this->balance -= $amount;
    }

    public function increaseBalance($amount) {
        $this->balance += $amount;
    }
}

class Bank {
    public function transfer(Account $from, Account $to, $amount) {
        if ($from->getBalance() >= $amount) {
            $from->decreaseBalance($amount);
            $to->increaseBalance($amount);
            return true;
        }
        return false;
    }
}

/**
 * 11. Moving the field (Move Field)
 *
 * BEFORE: Field in wrong class
 */
class CustomerBefore {
    private $name;
    private $address; // This should be in Address class

    public function __construct($name, $street, $city, $zipCode) {
        $this->name = $name;
        $this->address = ['street' => $street, 'city' => $city, 'zipCode' => $zipCode];
    }

    public function getAddress() {
        return $this->address['street'] . ', ' . $this->address['city'] . ' ' . $this->address['zipCode'];
    }
}

/**
 * AFTER: Move field to dedicated class
 */
class Address {
    private $street;
    private $city;
    private $zipCode;

    public function __construct($street, $city, $zipCode) {
        $this->street = $street;
        $this->city = $city;
        $this->zipCode = $zipCode;
    }

    public function getFullAddress() {
        return $this->street . ', ' . $this->city . ' ' . $this->zipCode;
    }
}

class CustomerAfter {
    private $name;
    private $address;

    public function __construct($name, Address $address) {
        $this->name = $name;
        $this->address = $address;
    }

    public function getAddress() {
        return $this->address->getFullAddress();
    }
}

/**
 * 12. Class Allocation (Extract Class)
 *
 * BEFORE: Class has too many responsibilities
 */
class PersonBefore {
    private $name;
    private $phoneNumber;
    private $officeAreaCode;
    private $officeNumber;

    public function getTelephoneNumber() {
        return '(' . $this->officeAreaCode . ') ' . $this->officeNumber;
    }
}

/**
 * AFTER: Extract telephone number to separate class
 */
class TelephoneNumber {
    private $areaCode;
    private $number;

    public function __construct($areaCode, $number) {
        $this->areaCode = $areaCode;
        $this->number = $number;
    }

    public function getTelephoneNumber() {
        return '(' . $this->areaCode . ') ' . $this->number;
    }
}

class PersonAfter {
    private $name;
    private $phoneNumber;
    private $officeTelephone;

    public function __construct($name) {
        $this->name = $name;
    }

    public function getOfficeTelephone() {
        return $this->officeTelephone->getTelephoneNumber();
    }

    public function setOfficeTelephone(TelephoneNumber $telephone) {
        $this->officeTelephone = $telephone;
    }
}

/**
 * 13. Embedding a class (Inline Class)
 *
 * BEFORE: Unnecessary class with single responsibility
 */
class OrderProcessorBefore {
    private $validator;

    public function __construct() {
        $this->validator = new OrderValidator();
    }

    public function process($order) {
        if ($this->validator->isValid($order)) {
            // Process order
        }
    }
}

class OrderValidator {
    public function isValid($order) {
        return $order['total'] > 0;
    }
}

/**
 * AFTER: Inline the class
 */
class OrderProcessorAfter {
    public function process($order) {
        if ($this->isValidOrder($order)) {
            // Process order
        }
    }

    private function isValidOrder($order) {
        return $order['total'] > 0;
    }
}

/**
 * 14. Hiding delegation (Hide Delegate)
 *
 * BEFORE: Client has to know about delegation
 */
class DepartmentBefore {
    private $manager;

    public function __construct(Person $manager) {
        $this->manager = $manager;
    }

    public function getManager() {
        return $this->manager;
    }
}

class Person {
    private $department;

    public function getDepartment() {
        return $this->department;
    }
}

// Client code
$manager = $person->getDepartment()->getManager();

/**
 * AFTER: Hide the delegation
 */
class DepartmentAfter {
    private $manager;

    public function __construct(PersonAfter $manager) {
        $this->manager = $manager;
    }

    public function getManager() {
        return $this->manager;
    }
}

class PersonAfter {
    private $department;

    public function getDepartment() {
        return $this->department;
    }

    public function getManager() {
        return $this->department->getManager();
    }
}

// Client code - much cleaner
$manager = $person->getManager();

/**
 * 15. Removing the intermediary (Remove Middle Man)
 *
 * BEFORE: Too much delegation
 */
class PersonWithMiddleMan {
    private $department;

    public function getDepartment() {
        return $this->department;
    }

    public function getManager() {
        return $this->department->getManager();
    }

    public function getDepartmentName() {
        return $this->department->getName();
    }
}

/**
 * AFTER: Remove middle man if delegation is too heavy
 */
class PersonDirect {
    private $department;
    private $manager; // Direct reference

    public function getManager() {
        return $this->manager;
    }

    public function getDepartment() {
        return $this->department;
    }
}

/**
 * 16. Introduction of an external method (Introduce Foreign Method)
 *
 * BEFORE: Using external class method in wrong place
 */
class ReportGeneratorBefore {
    public function generateReport() {
        $date = new DateTime();
        $nextMonth = $date->modify('+1 month'); // Foreign method usage

        // Generate report for next month
    }
}

/**
 * AFTER: Introduce foreign method
 */
class ReportGeneratorAfter {
    public function generateReport() {
        $date = new DateTime();
        $nextMonth = $this->nextMonth($date);

        // Generate report for next month
    }

    private function nextMonth(DateTime $date) {
        return clone $date->modify('+1 month');
    }
}

/**
 * 17. The introduction of local extension (Introduce Local Extension)
 *
 * BEFORE: Adding methods to external class (not possible)
 */
class DateUtil {
    public static function nextMonth(DateTime $date) {
        return clone $date->modify('+1 month');
    }

    public static function previousMonth(DateTime $date) {
        return clone $date->modify('-1 month');
    }
}

/**
 * AFTER: Create local extension class
 */
class DateTimeExtension extends DateTime {
    public function nextMonth() {
        return clone $this->modify('+1 month');
    }

    public function previousMonth() {
        return clone $this->modify('-1 month');
    }
}
