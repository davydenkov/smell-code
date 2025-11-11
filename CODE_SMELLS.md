# Code Smells Guide

This guide provides detailed explanations of the **10 code smells** covered in this project, along with their refactoring solutions.

## Table of Contents

1. [Code Duplication](#1-code-duplication)
2. [Data Classes](#2-data-classes)
3. [Data Clumps](#3-data-clumps)
4. [Divergent Modifications](#4-divergent-modifications)
5. [Feature Envy](#5-feature-envy)
6. [Incompleteness of Library Class](#6-incompleteness-of-library-class)
7. [Large Class](#7-large-class)
8. [Long Method](#8-long-method)
9. [Long Parameters](#9-long-parameters)
10. [Renunciation of Inheritance](#10-renunciation-of-inheritance)

---

## 1. Code Duplication

### What is it?
Code duplication occurs when the same code appears in multiple places within a codebase.

### Why is it bad?
- **Maintenance nightmare**: Changes need to be made in multiple places
- **Bug risk**: Fixing a bug in one place but forgetting others
- **Code bloat**: Increases codebase size unnecessarily
- **Readability**: Makes code harder to understand

### Example Scenario
```php
// BAD: Duplicated tax calculation logic
class OrderProcessor {
    public function calculateTotal($subtotal) {
        $tax = $subtotal * 0.08;  // Tax calculation duplicated
        return $subtotal + $tax;
    }
}

class InvoiceProcessor {
    public function calculateTotal($subtotal) {
        $tax = $subtotal * 0.08;  // Same logic duplicated
        return $subtotal + $tax;
    }
}
```

### Refactoring Solution
Extract the duplicated logic into a separate, reusable class or method.

```php
// GOOD: Extracted to reusable TaxCalculator
class TaxCalculator {
    public function calculateTax($amount) {
        return $amount * 0.08;
    }
}

class OrderProcessor {
    private $taxCalculator;

    public function __construct(TaxCalculator $taxCalculator) {
        $this->taxCalculator = $taxCalculator;
    }

    public function calculateTotal($subtotal) {
        $tax = $this->taxCalculator->calculateTax($subtotal);
        return $subtotal + $tax;
    }
}
```

### Files to Check
- `code-duplication/bad/OrderProcessor.php`
- `code-duplication/good/OrderProcessor.php`
- `code-duplication/good/TaxCalculator.php`

---

## 2. Data Classes

### What is it?
Data classes are classes that contain only fields and simple getters/setters, with no business logic.

### Why is it bad?
- **Anemic domain model**: No behavior where it belongs
- **Logic scattered**: Business rules spread across the application
- **Encapsulation violation**: Data exposed without protection
- **Maintenance issues**: Changes require finding all usage points

### Example Scenario
```php
// BAD: Data class with no behavior
class User {
    private $name;
    private $email;
    private $password;

    public function getName() { return $this->name; }
    public function setName($name) { $this->name = $name; }
    public function getEmail() { return $this->email; }
    public function setEmail($email) { $this->email = $email; }
    // ... more getters/setters
}

// Logic scattered elsewhere
class UserService {
    public function isValidEmail($email) {
        return filter_var($email, FILTER_VALIDATE_EMAIL);
    }
}
```

### Refactoring Solution
Move behavior into the data classes where it belongs.

```php
// GOOD: Data class with behavior
class User {
    private $name;
    private $email;
    private $password;

    public function __construct($name, $email, $password) {
        $this->name = $name;
        $this->setEmail($email);
        $this->setPassword($password);
    }

    public function getName() { return $this->name; }

    public function getEmail() { return $this->email; }
    public function setEmail($email) {
        if (!$this->isValidEmail($email)) {
            throw new InvalidArgumentException('Invalid email');
        }
        $this->email = $email;
    }

    public function getPassword() { return $this->password; }
    public function setPassword($password) {
        $this->password = $this->hashPassword($password);
    }

    private function isValidEmail($email) {
        return filter_var($email, FILTER_VALIDATE_EMAIL);
    }

    private function hashPassword($password) {
        return password_hash($password, PASSWORD_DEFAULT);
    }
}
```

### Files to Check
- `data-classes/bad/User.php`
- `data-classes/good/User.php`

---

## 3. Data Clumps

### What is it?
Data clumps are groups of data that always appear together and are passed around as a group.

### Why is it bad?
- **Parameter explosion**: Methods end up with many parameters
- **Inconsistency**: Data can get out of sync
- **Missing encapsulation**: Related data not grouped logically
- **Maintenance burden**: Changes affect many methods

### Example Scenario
```php
// BAD: Data clump parameters
class CustomerService {
    public function createCustomer($firstName, $lastName, $street, $city, $state, $zip) {
        // All address parameters always passed together
    }

    public function updateAddress($customerId, $street, $city, $state, $zip) {
        // Same address parameters again
    }
}
```

### Refactoring Solution
Create a class to group the related data together.

```php
// GOOD: Extract Address class
class Address {
    private $street;
    private $city;
    private $state;
    private $zip;

    public function __construct($street, $city, $state, $zip) {
        $this->street = $street;
        $this->city = $city;
        $this->state = $state;
        $this->zip = $zip;
    }

    // getters and setters...
}

class CustomerService {
    public function createCustomer($firstName, $lastName, Address $address) {
        // Much cleaner
    }

    public function updateAddress($customerId, Address $address) {
        // Cleaner still
    }
}
```

### Files to Check
- `data-clumps/bad/CustomerService.php`
- `data-clumps/good/Address.php`
- `data-clumps/good/CustomerService.php`

---

## 4. Divergent Modifications

### What is it?
A class that is modified for many different reasons indicates it has multiple responsibilities.

### Why is it bad?
- **Single Responsibility Principle violation**: Class does too many things
- **Fragile code**: Changes in one area affect others
- **Testing difficulty**: Hard to test individual concerns
- **Maintenance complexity**: Understanding and modifying becomes harder

### Example Scenario
```php
// BAD: Single class handling multiple concerns
class FinancialService {
    public function processPayment($amount, $method) {
        // Payment processing logic
    }

    public function generateReport($type) {
        // Report generation logic
    }

    public function sendEmail($to, $subject, $body) {
        // Email sending logic
    }

    public function saveTransaction($transaction) {
        // Data persistence logic
    }
}
```

### Refactoring Solution
Split the class into separate classes, each handling one responsibility.

```php
// GOOD: Separate classes for separate concerns
class PaymentService {
    public function processPayment($amount, $method) {
        // Only payment logic
    }
}

class ReportGenerator {
    public function generateReport($type) {
        // Only report logic
    }
}

class EmailService {
    public function sendEmail($to, $subject, $body) {
        // Only email logic
    }
}

class TransactionRepository {
    public function saveTransaction($transaction) {
        // Only persistence logic
    }
}
```

### Files to Check
- `divergent-modifications/bad/FinancialService.php`
- `divergent-modifications/good/` (multiple files)

---

## 5. Feature Envy

### What is it?
A method that seems more interested in another class than the one it's defined in.

### Why is it bad?
- **Poor encapsulation**: Method knows too much about another class
- **Tight coupling**: Classes become dependent on each other's internals
- **Maintenance issues**: Changes in one class affect methods in another
- **Logic placement**: Behavior should be with the data it operates on

### Example Scenario
```php
// BAD: Method uses Rectangle data more than its own
class Graphics {
    public function drawRectangle($x1, $y1, $x2, $y2) {
        $width = $x2 - $x1;    // Using Rectangle data
        $height = $y2 - $y1;   // Using Rectangle data
        $area = $width * $height; // Calculating Rectangle properties
        // Drawing logic...
    }
}
```

### Refactoring Solution
Move the method to the class whose data it uses most.

```php
// GOOD: Method moved to Rectangle class
class Rectangle {
    private $x1, $y1, $x2, $y2;

    public function __construct($x1, $y1, $x2, $y2) {
        $this->x1 = $x1; $this->y1 = $y1;
        $this->x2 = $x2; $this->y2 = $y2;
    }

    public function getWidth() {
        return $this->x2 - $this->x1;
    }

    public function getHeight() {
        return $this->y2 - $this->y1;
    }

    public function getArea() {
        return $this->getWidth() * $this->getHeight();
    }
}

class Graphics {
    public function drawRectangle(Rectangle $rect) {
        $area = $rect->getArea(); // Much cleaner
        // Drawing logic...
    }
}
```

### Files to Check
- `feature-envy/bad/Rectangle.php`
- `feature-envy/good/Rectangle.php`

---

## 6. Incompleteness of Library Class

### What is it?
A library or utility class that doesn't provide all the functionality you need.

### Why is it bad?
- **Repeated code**: Similar extension methods across the codebase
- **Inconsistency**: Different approaches to extending functionality
- **Maintenance burden**: Multiple places to update when library changes
- **Code duplication**: Same extension logic repeated

### Example Scenario
```php
// BAD: Extending library class functionality inline
class ApiClient {
    private $httpClient;

    public function makeRequest($url) {
        $response = $this->httpClient->get($url);
        $data = json_decode($response->getBody(), true);

        // Custom logic repeated everywhere
        if (!$response->isSuccess()) {
            throw new ApiException('API request failed');
        }

        return $data;
    }
}
```

### Refactoring Solution
Create extension methods or wrapper classes to add the missing functionality.

```php
// GOOD: Extend the library class
class HttpClientExtension extends HttpClient {
    public function getJson($url) {
        $response = $this->get($url);

        if (!$response->isSuccess()) {
            throw new ApiException('API request failed');
        }

        return json_decode($response->getBody(), true);
    }

    public function postJson($url, $data) {
        $response = $this->post($url, json_encode($data));

        if (!$response->isSuccess()) {
            throw new ApiException('API request failed');
        }

        return json_decode($response->getBody(), true);
    }
}

class ApiClient {
    private $httpClient;

    public function makeRequest($url) {
        return $this->httpClient->getJson($url); // Much cleaner
    }
}
```

### Files to Check
- `incompleteness-of-library-class/bad/HttpClient.php`
- `incompleteness-of-library-class/good/HttpClient.php`

---

## 7. Large Class

### What is it?
A class that has grown too large and is trying to do too many things.

### Why is it bad?
- **Single Responsibility Principle violation**: Too many responsibilities
- **Hard to understand**: Overwhelming amount of code
- **Hard to test**: Many different testing scenarios
- **Fragile**: Changes in one area affect the whole class
- **Duplication risk**: Similar logic may exist within the class

### Example Scenario
```php
// BAD: One class doing everything
class UserService {
    // User management
    public function createUser($data) { /* ... */ }
    public function updateUser($id, $data) { /* ... */ }
    public function deleteUser($id) { /* ... */ }

    // Email functionality
    public function sendWelcomeEmail($user) { /* ... */ }
    public function sendPasswordReset($user) { /* ... */ }

    // Payment processing
    public function processPayment($user, $amount) { /* ... */ }

    // Reporting
    public function generateUserReport() { /* ... */ }

    // Data validation
    public function validateUserData($data) { /* ... */ }
}
```

### Refactoring Solution
Split the large class into smaller, focused classes.

```php
// GOOD: Separate classes for separate responsibilities
class UserRepository {
    public function createUser($data) { /* ... */ }
    public function updateUser($id, $data) { /* ... */ }
    public function deleteUser($id) { /* ... */ }
}

class EmailService {
    public function sendWelcomeEmail($user) { /* ... */ }
    public function sendPasswordReset($user) { /* ... */ }
}

class PaymentService {
    public function processPayment($user, $amount) { /* ... */ }
}

class ReportService {
    public function generateUserReport() { /* ... */ }
}

class User {
    // User-specific behavior
    public function validateData($data) { /* ... */ }
}
```

### Files to Check
- `large-class/bad/UserService.php`
- `large-class/good/` (multiple files)

---

## 8. Long Method

### What is it?
A method that has grown too long and is doing too many things.

### Why is it bad?
- **Hard to understand**: Too much logic in one place
- **Hard to test**: Complex testing scenarios
- **Duplication**: Similar logic may be repeated
- **Maintenance**: Changes affect large portions of code
- **Reusability**: Hard to reuse parts of the method

### Example Scenario
```php
// BAD: One long method doing everything
class UserManager {
    public function processUserRegistration($userData) {
        // Validate input
        if (empty($userData['name'])) {
            throw new Exception('Name is required');
        }

        // Create user object
        $user = new User($userData['name'], $userData['email']);

        // Hash password
        $user->setPassword(password_hash($userData['password'], PASSWORD_DEFAULT));

        // Save to database
        $this->saveUser($user);

        // Send welcome email
        $this->sendWelcomeEmail($user);

        // Log the registration
        $this->logRegistration($user);

        // Return success
        return $user;
    }
}
```

### Refactoring Solution
Break down the long method into smaller, focused methods.

```php
// GOOD: Extracted into smaller methods
class UserManager {
    public function processUserRegistration($userData) {
        $user = $this->createUser($userData);
        $this->saveUser($user);
        $this->sendWelcomeEmail($user);
        $this->logRegistration($user);
        return $user;
    }

    private function createUser($userData) {
        $this->validateUserData($userData);
        $user = new User($userData['name'], $userData['email']);
        $user->setPassword($this->hashPassword($userData['password']));
        return $user;
    }

    private function validateUserData($userData) {
        if (empty($userData['name'])) {
            throw new Exception('Name is required');
        }
        // More validation...
    }

    private function hashPassword($password) {
        return password_hash($password, PASSWORD_DEFAULT);
    }

    private function saveUser($user) { /* ... */ }
    private function sendWelcomeEmail($user) { /* ... */ }
    private function logRegistration($user) { /* ... */ }
}
```

### Files to Check
- `long-method/bad/UserManager.php`
- `long-method/good/` (multiple files)

---

## 9. Long Parameters

### What is it?
Methods that accept too many parameters.

### Why is it bad?
- **Hard to understand**: Method calls are confusing
- **Error prone**: Easy to pass parameters in wrong order
- **Maintenance**: Adding/removing parameters affects many call sites
- **Testing**: Many combinations to test
- **Readability**: Method signatures become unwieldy

### Example Scenario
```php
// BAD: Too many parameters
class OrderService {
    public function createOrder($customerId, $productId, $quantity, $shippingAddress, $billingAddress, $paymentMethod, $couponCode, $notes) {
        // Method with 8 parameters - hard to use and understand
    }
}
```

### Refactoring Solution
Group related parameters into objects or use method chaining.

```php
// GOOD: Group parameters into objects
class Address {
    // Address fields...
}

class OrderData {
    public $customerId;
    public $productId;
    public $quantity;
    public $shippingAddress;
    public $billingAddress;
    public $paymentMethod;
    public $couponCode;
    public $notes;
}

class OrderService {
    public function createOrder(OrderData $orderData) {
        // Much cleaner method signature
    }
}
```

### Files to Check
- `long-parameters/bad/OrderService.php`
- `long-parameters/good/` (multiple files)

---

## 10. Renunciation of Inheritance

### What is it?
Inheritance used incorrectly, or inheritance not used when it would be appropriate.

### Why is it bad?
- **Tight coupling**: Subclasses depend heavily on parent implementation
- **Fragile base class**: Changes to parent break subclasses
- **Inappropriate hierarchies**: Inheritance used for wrong reasons
- **Composition over inheritance**: Better alternatives exist

### Example Scenario
```php
// BAD: Inheritance for configuration
class ShapeRenderer {
    // Lots of conditional logic based on shape type
    public function renderCircle($radius) {
        $this->setColor('red');
        // Circle rendering logic
    }

    public function renderSquare($size) {
        $this->setColor('blue');
        // Square rendering logic
    }

    private function setColor($color) {
        // Color setting logic
    }
}
```

### Refactoring Solution
Use composition instead of inheritance, or proper inheritance hierarchies.

```php
// GOOD: Composition over inheritance
interface Renderer {
    public function render();
}

class CircleRenderer implements Renderer {
    private $colorSetter;

    public function __construct(ColorSetter $colorSetter) {
        $this->colorSetter = $colorSetter;
    }

    public function render() {
        $this->colorSetter->setColor('red');
        // Circle rendering logic
    }
}

class SquareRenderer implements Renderer {
    private $colorSetter;

    public function __construct(ColorSetter $colorSetter) {
        $this->colorSetter = $colorSetter;
    }

    public function render() {
        $this->colorSetter->setColor('blue');
        // Square rendering logic
    }
}
```

### Files to Check
- `renunciation-of-inheritance/bad/ShapeRenderer.php`
- `renunciation-of-inheritance/good/ShapeRenderer.php`

---

## Summary

Recognizing and refactoring code smells leads to:
- **More maintainable code**
- **Easier testing**
- **Better readability**
- **Reduced duplication**
- **Improved extensibility**

Remember: Code smells are indicators, not hard rules. Always consider the context and requirements of your project before refactoring.
