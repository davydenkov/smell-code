<?php

/**
 * 69. Separation of inheritance (Tease Apart Inheritance)
 *
 * BEFORE: Class hierarchy mixing two different responsibilities
 */
class EmployeeTeaseBefore {
    protected $name;
    protected $rate;

    public function __construct($name, $rate) {
        $this->name = $name;
        $this->rate = $rate;
    }

    public function getName() {
        return $this->name;
    }
}

class SalariedEmployeeTeaseBefore extends EmployeeTeaseBefore {
    public function getPay() {
        return $this->rate;
    }
}

class CommissionedEmployeeTeaseBefore extends EmployeeTeaseBefore {
    private $commission;

    public function __construct($name, $rate, $commission) {
        parent::__construct($name, $rate);
        $this->commission = $commission;
    }

    public function getPay() {
        return $this->rate + $this->commission;
    }
}

/**
 * AFTER: Tease apart inheritance into two separate hierarchies
 */
interface Payable {
    public function getPay();
}

class EmployeeTeaseAfter {
    protected $name;

    public function __construct($name) {
        $this->name = $name;
    }

    public function getName() {
        return $this->name;
    }
}

class SalariedEmployeeTeaseAfter extends EmployeeTeaseAfter implements Payable {
    private $salary;

    public function __construct($name, $salary) {
        parent::__construct($name);
        $this->salary = $salary;
    }

    public function getPay() {
        return $this->salary;
    }
}

class CommissionedEmployeeTeaseAfter extends EmployeeTeaseAfter implements Payable {
    private $baseSalary;
    private $commission;

    public function __construct($name, $baseSalary, $commission) {
        parent::__construct($name);
        $this->baseSalary = $baseSalary;
        $this->commission = $commission;
    }

    public function getPay() {
        return $this->baseSalary + $this->commission;
    }
}

/**
 * 70. Converting a procedural project into objects (Convert Procedural Design to Objects)
 *
 * BEFORE: Procedural code with global functions and data
 */
class ProceduralDesignBefore {
    private static $accounts = [];

    public static function createAccount($id, $balance) {
        self::$accounts[$id] = $balance;
    }

    public static function getBalance($id) {
        return self::$accounts[$id] ?? 0;
    }

    public static function deposit($id, $amount) {
        if (isset(self::$accounts[$id])) {
            self::$accounts[$id] += $amount;
        }
    }

    public static function withdraw($id, $amount) {
        if (isset(self::$accounts[$id]) && self::$accounts[$id] >= $amount) {
            self::$accounts[$id] -= $amount;
            return true;
        }
        return false;
    }
}

/**
 * AFTER: Convert to object-oriented design
 */
class Account {
    private $id;
    private $balance;

    public function __construct($id, $balance = 0) {
        $this->id = $id;
        $this->balance = $balance;
    }

    public function getId() {
        return $this->id;
    }

    public function getBalance() {
        return $this->balance;
    }

    public function deposit($amount) {
        $this->balance += $amount;
    }

    public function withdraw($amount) {
        if ($this->balance >= $amount) {
            $this->balance -= $amount;
            return true;
        }
        return false;
    }
}

class Bank {
    private $accounts = [];

    public function createAccount($id, $balance = 0) {
        $account = new Account($id, $balance);
        $this->accounts[$id] = $account;
        return $account;
    }

    public function getAccount($id) {
        return $this->accounts[$id] ?? null;
    }

    public function getBalance($id) {
        $account = $this->getAccount($id);
        return $account ? $account->getBalance() : 0;
    }

    public function deposit($id, $amount) {
        $account = $this->getAccount($id);
        if ($account) {
            $account->deposit($amount);
        }
    }

    public function withdraw($id, $amount) {
        $account = $this->getAccount($id);
        return $account ? $account->withdraw($amount) : false;
    }
}

/**
 * 71. Separating the domain from the representation (Separate Domain from Presentation)
 *
 * BEFORE: Domain logic mixed with presentation
 */
class OrderPresentationBefore {
    private $items = [];
    private $total = 0;

    public function addItem($name, $price, $quantity) {
        $this->items[] = ['name' => $name, 'price' => $price, 'quantity' => $quantity];
        $this->total += $price * $quantity;

        // Presentation logic mixed in
        echo "Added $quantity x $name to order\n";
        echo "Current total: $" . number_format($this->total, 2) . "\n";
    }

    public function getTotal() {
        return $this->total;
    }

    public function displayOrder() {
        echo "Order Summary:\n";
        foreach ($this->items as $item) {
            echo "- {$item['quantity']} x {$item['name']} @ $" . number_format($item['price'], 2) . "\n";
        }
        echo "Total: $" . number_format($this->total, 2) . "\n";
    }
}

/**
 * AFTER: Separate domain from presentation
 */
class OrderItem {
    private $name;
    private $price;
    private $quantity;

    public function __construct($name, $price, $quantity) {
        $this->name = $name;
        $this->price = $price;
        $this->quantity = $quantity;
    }

    public function getName() {
        return $this->name;
    }

    public function getPrice() {
        return $this->price;
    }

    public function getQuantity() {
        return $this->quantity;
    }

    public function getTotal() {
        return $this->price * $this->quantity;
    }
}

class OrderDomainAfter {
    private $items = [];

    public function addItem($name, $price, $quantity) {
        $item = new OrderItem($name, $price, $quantity);
        $this->items[] = $item;
    }

    public function getItems() {
        return $this->items;
    }

    public function getTotal() {
        $total = 0;
        foreach ($this->items as $item) {
            $total += $item->getTotal();
        }
        return $total;
    }
}

class OrderPresenter {
    public function displayItemAdded($item) {
        echo "Added {$item->getQuantity()} x {$item->getName()} to order\n";
    }

    public function displayOrderSummary(OrderDomainAfter $order) {
        echo "Order Summary:\n";
        foreach ($order->getItems() as $item) {
            echo "- {$item->getQuantity()} x {$item->getName()} @ $" . number_format($item->getPrice(), 2) . "\n";
        }
        echo "Total: $" . number_format($order->getTotal(), 2) . "\n";
    }
}

class OrderService {
    private $order;
    private $presenter;

    public function __construct() {
        $this->order = new OrderDomainAfter();
        $this->presenter = new OrderPresenter();
    }

    public function addItem($name, $price, $quantity) {
        $item = new OrderItem($name, $price, $quantity);
        $this->order->addItem($name, $price, $quantity);
        $this->presenter->displayItemAdded($item);
        $this->presenter->displayOrderSummary($this->order);
    }

    public function getOrder() {
        return $this->order;
    }
}

/**
 * 72. Hierarchy Extraction (Extract Hierarchy)
 *
 * BEFORE: Single class handling multiple responsibilities
 */
class ComputerExtractBefore {
    private $type;
    private $cpu;
    private $ram;
    private $storage;

    public function __construct($type, $cpu, $ram, $storage) {
        $this->type = $type;
        $this->cpu = $cpu;
        $this->ram = $ram;
        $this->storage = $storage;
    }

    public function getSpecs() {
        $specs = "CPU: {$this->cpu}\n";
        $specs .= "RAM: {$this->ram}GB\n";
        $specs .= "Storage: {$this->storage}GB\n";

        if ($this->type === 'desktop') {
            $specs .= "Form Factor: Desktop\n";
        } elseif ($this->type === 'laptop') {
            $specs .= "Form Factor: Laptop\n";
            $specs .= "Battery Life: 8 hours\n";
        } elseif ($this->type === 'server') {
            $specs .= "Form Factor: Server Rack\n";
            $specs .= "Redundancy: RAID 10\n";
        }

        return $specs;
    }
}

/**
 * AFTER: Extract hierarchy
 */
abstract class ComputerExtractAfter {
    protected $cpu;
    protected $ram;
    protected $storage;

    public function __construct($cpu, $ram, $storage) {
        $this->cpu = $cpu;
        $this->ram = $ram;
        $this->storage = $storage;
    }

    abstract public function getFormFactor();
    abstract public function getSpecialFeatures();

    public function getBasicSpecs() {
        return "CPU: {$this->cpu}\n" .
               "RAM: {$this->ram}GB\n" .
               "Storage: {$this->storage}GB\n";
    }

    public function getSpecs() {
        return $this->getBasicSpecs() .
               "Form Factor: " . $this->getFormFactor() . "\n" .
               $this->getSpecialFeatures();
    }
}

class DesktopComputer extends ComputerExtractAfter {
    public function getFormFactor() {
        return "Desktop";
    }

    public function getSpecialFeatures() {
        return "Expansion Slots: Multiple PCI\n";
    }
}

class LaptopComputer extends ComputerExtractAfter {
    public function getFormFactor() {
        return "Laptop";
    }

    public function getSpecialFeatures() {
        return "Battery Life: 8 hours\nWeight: 2.5 lbs\n";
    }
}

class ServerComputer extends ComputerExtractAfter {
    public function getFormFactor() {
        return "Server Rack";
    }

    public function getSpecialFeatures() {
        return "Redundancy: RAID 10\nHot Swap Drives: Yes\n";
    }
}
