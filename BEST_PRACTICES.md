# Clean Code Best Practices

This guide summarizes the best practices learned from studying code smells and refactoring techniques.

## Table of Contents

1. [Class Design Principles](#class-design-principles)
2. [Method Design Principles](#method-design-principles)
3. [Code Organization](#code-organization)
4. [SOLID Principles](#solid-principles)
5. [Testing Best Practices](#testing-best-practices)
6. [Refactoring Guidelines](#refactoring-guidelines)

---

## Class Design Principles

### Single Responsibility Principle (SRP)
```php
// ❌ BAD: Multiple responsibilities
class UserService {
    public function createUser($data) { /* ... */ }
    public function sendEmail($user) { /* ... */ }
    public function processPayment($user) { /* ... */ }
}

// ✅ GOOD: Single responsibility per class
class UserRepository {
    public function createUser($data) { /* ... */ }
}

class EmailService {
    public function sendEmail($user) { /* ... */ }
}
```

### Class Size Guidelines
- **Small classes**: 100-200 lines maximum
- **Few instance variables**: 5-10 fields maximum
- **Clear naming**: Class name should describe its single responsibility

### Composition Over Inheritance
```php
// ❌ BAD: Inheritance for code reuse
class UserManager extends DatabaseConnection {
    // Inherits database functionality
}

// ✅ GOOD: Composition
class UserManager {
    private $database;

    public function __construct(DatabaseConnection $db) {
        $this->database = $db;
    }
}
```

---

## Method Design Principles

### Method Size Guidelines
- **1-5 lines**: Simple getters/setters, basic operations
- **6-10 lines**: Straightforward business logic
- **11-20 lines**: Consider refactoring
- **20+ lines**: Strong candidate for decomposition

### Method Naming
```php
// ❌ BAD: Unclear names
public function process() { /* ... */ }
public function doStuff() { /* ... */ }

// ✅ GOOD: Descriptive names
public function processUserRegistration() { /* ... */ }
public function calculateTotalPrice() { /* ... */ }
```

### Parameter Guidelines
- **0 parameters**: Ideal for simple operations
- **1-3 parameters**: Acceptable for most methods
- **4+ parameters**: Consider refactoring

```php
// ❌ BAD: Too many parameters
public function createOrder($id, $name, $email, $address, $payment, $items) { /* ... */ }

// ✅ GOOD: Parameter object
public function createOrder(OrderRequest $request) { /* ... */ }
```

### Early Returns
```php
// ❌ BAD: Deep nesting
public function processOrder($order) {
    if ($order->isValid()) {
        if ($order->hasItems()) {
            if ($order->canShip()) {
                // Process order
            }
        }
    }
}

// ✅ GOOD: Early returns
public function processOrder($order) {
    if (!$order->isValid()) return false;
    if (!$order->hasItems()) return false;
    if (!$order->canShip()) return false;

    // Process order
    return true;
}
```

---

## Code Organization

### Package Structure
```
src/
├── Domain/
│   ├── Entity/
│   ├── ValueObject/
│   └── Service/
├── Application/
│   ├── Command/
│   ├── Query/
│   └── Handler/
├── Infrastructure/
│   ├── Repository/
│   ├── ExternalService/
│   └── Database/
└── Presentation/
    ├── Controller/
    ├── View/
    └── DTO/
```

### File Organization
- **One class per file**
- **Related classes in same package**
- **Clear package naming conventions**

### Import Organization
```php
// ❌ BAD: Unorganized imports
use App\Service\UserService;
use App\Repository\OrderRepository;
use App\Model\User;
use App\Util\StringHelper;

// ✅ GOOD: Grouped and ordered imports
use App\Model\User;
use App\Repository\OrderRepository;
use App\Service\UserService;
use App\Util\StringHelper;
```

---

## SOLID Principles

### 1. Single Responsibility Principle
*"A class should have only one reason to change"*

### 2. Open/Closed Principle
*"Software entities should be open for extension, but closed for modification"*

```php
// ✅ GOOD: Open for extension
interface PaymentProcessor {
    public function process(Payment $payment): bool;
}

class CreditCardProcessor implements PaymentProcessor {
    public function process(Payment $payment): bool { /* ... */ }
}

class PayPalProcessor implements PaymentProcessor {
    public function process(Payment $payment): bool { /* ... */ }
}
```

### 3. Liskov Substitution Principle
*"Subtypes must be substitutable for their base types"*

### 4. Interface Segregation Principle
*"Clients should not be forced to depend on interfaces they don't use"*

```php
// ❌ BAD: Fat interface
interface Worker {
    public function work();
    public function eat();
    public function sleep();
}

// ✅ GOOD: Segregated interfaces
interface Workable {
    public function work();
}

interface Eatable {
    public function eat();
}

interface Sleepable {
    public function sleep();
}
```

### 5. Dependency Inversion Principle
*"High-level modules should not depend on low-level modules"*

```php
// ✅ GOOD: Depend on abstractions
interface DatabaseConnection {
    public function connect(): PDO;
}

class UserRepository {
    private $connection;

    public function __construct(DatabaseConnection $connection) {
        $this->connection = $connection;
    }
}
```

---

## Testing Best Practices

### Unit Testing
- **Test one thing per test method**
- **Use descriptive test names**
- **Test both success and failure cases**
- **Mock external dependencies**

```php
// ✅ GOOD: Focused test
public function testUserCreationWithValidData() {
    $userData = ['name' => 'John', 'email' => 'john@example.com'];
    $user = $this->userService->createUser($userData);

    $this->assertEquals('John', $user->getName());
    $this->assertEquals('john@example.com', $user->getEmail());
}
```

### Test Structure
```php
class UserServiceTest extends TestCase {
    private $userService;
    private $userRepository;

    protected function setUp(): void {
        $this->userRepository = $this->createMock(UserRepository::class);
        $this->userService = new UserService($this->userRepository);
    }

    public function testCreatesUserSuccessfully() {
        // Arrange
        $userData = ['name' => 'John', 'email' => 'john@example.com'];

        // Act
        $user = $this->userService->createUser($userData);

        // Assert
        $this->assertInstanceOf(User::class, $user);
    }
}
```

---

## Refactoring Guidelines

### When to Refactor
- **Before adding new features**: Clean up code first
- **When fixing bugs**: Improve code quality while fixing
- **During code reviews**: Address smells immediately
- **When code becomes hard to understand**

### Refactoring Steps
1. **Identify the smell**: Recognize the problem
2. **Understand the code**: Know what it does
3. **Write tests**: Ensure safety during refactoring
4. **Apply refactoring**: Use appropriate techniques
5. **Test**: Verify behavior is preserved
6. **Clean up**: Remove any temporary code

### Safe Refactoring Checklist
- [ ] All existing tests pass
- [ ] New tests added for new code
- [ ] Code style standards maintained
- [ ] Documentation updated
- [ ] No performance regressions
- [ ] Team reviewed changes

### Common Refactoring Patterns

#### Extract Method
```php
// Before
public function calculateTotal() {
    $subtotal = $this->getSubtotal();
    $tax = $subtotal * 0.08;
    $shipping = $this->calculateShipping();
    return $subtotal + $tax + $shipping;
}

// After
public function calculateTotal() {
    $subtotal = $this->getSubtotal();
    $tax = $this->calculateTax($subtotal);
    $shipping = $this->calculateShipping();
    return $subtotal + $tax + $shipping;
}

private function calculateTax($amount) {
    return $amount * 0.08;
}
```

#### Extract Class
```php
// Before: Large class
class OrderProcessor {
    public function process($orderData) {
        $this->validate($orderData);
        $this->calculateTotal($orderData);
        $this->save($orderData);
    }
}

// After: Split responsibilities
class OrderValidator {
    public function validate($orderData) { /* ... */ }
}

class OrderCalculator {
    public function calculateTotal($orderData) { /* ... */ }
}

class OrderRepository {
    public function save($orderData) { /* ... */ }
}
```

---

## Naming Conventions

### Classes
- **PascalCase**: `UserService`, `OrderRepository`
- **Descriptive**: `PaymentProcessor`, not `Processor`
- **No abbreviations**: `HttpClient`, not `HTTPClient`

### Methods
- **camelCase**: `getUserById()`, `processPayment()`
- **Verbs for actions**: `save()`, `delete()`, `process()`
- **Questions for booleans**: `isValid()`, `hasPermission()`

### Variables
- **camelCase**: `$userName`, `$orderTotal`
- **Descriptive**: `$customerEmail`, not `$ce`
- **No single letters** (except loops): `$i`, `$j`

### Constants
- **UPPER_SNAKE_CASE**: `DEFAULT_TIMEOUT`, `MAX_RETRIES`

---

## Error Handling

### Exceptions vs Error Codes
```php
// ❌ BAD: Error codes
public function saveUser($user) {
    if (!$this->isValid($user)) {
        return ERROR_INVALID_USER;
    }
    // Save logic...
    return SUCCESS;
}

// ✅ GOOD: Exceptions
public function saveUser($user) {
    if (!$this->isValid($user)) {
        throw new InvalidUserException('User data is invalid');
    }
    // Save logic...
    return true;
}
```

### Custom Exceptions
```php
class UserNotFoundException extends Exception {}
class InvalidEmailException extends Exception {}
class PaymentFailedException extends Exception {}
```

### Exception Handling
```php
try {
    $user = $userService->createUser($userData);
    $paymentService->chargeUser($user, $amount);
} catch (InvalidUserException $e) {
    // Handle validation errors
    $this->showValidationErrors($e->getMessage());
} catch (PaymentFailedException $e) {
    // Handle payment errors
    $this->showPaymentError($e->getMessage());
} catch (Exception $e) {
    // Handle unexpected errors
    $this->logError($e);
    $this->showGenericError();
}
```

---

## Performance Considerations

### Avoid Premature Optimization
1. **Write clean code first**
2. **Measure performance**
3. **Optimize bottlenecks only**

### Common Performance Smells
- **N+1 queries**: Multiple database calls in loops
- **Large objects**: Loading unnecessary data
- **Tight loops**: Inefficient algorithms
- **Memory leaks**: Not releasing resources

### Efficient Patterns
```php
// ✅ GOOD: Batch operations
public function updateUsers($userIds, $data) {
    return $this->userRepository->batchUpdate($userIds, $data);
}

// ✅ GOOD: Lazy loading
class Order {
    private $items;
    private $itemsLoaded = false;

    public function getItems() {
        if (!$this->itemsLoaded) {
            $this->items = $this->loadItemsFromDatabase();
            $this->itemsLoaded = true;
        }
        return $this->items;
    }
}
```

---

## Documentation Best Practices

### Code Comments
```php
// ❌ BAD: Obvious comments
$i = 0; // Set i to 0

// ✅ GOOD: Explain why and what
// Calculate total price including tax and shipping
$total = $subtotal + $tax + $shipping;
```

### DocBlocks
```php
/**
 * Creates a new user account
 *
 * @param array $userData User information
 * @return User The created user object
 * @throws InvalidUserException When user data is invalid
 * @throws EmailExistsException When email is already registered
 */
public function createUser(array $userData): User {
    // Implementation...
}
```

### README Files
- **Project overview**: What the project does
- **Installation**: How to set up
- **Usage examples**: How to use the code
- **API documentation**: Key classes and methods
- **Contributing guidelines**: How to contribute

---

## Code Review Checklist

### Design
- [ ] Single Responsibility Principle followed
- [ ] Appropriate abstractions used
- [ ] SOLID principles applied
- [ ] Code duplication eliminated

### Implementation
- [ ] Methods are small and focused
- [ ] Variables and methods named clearly
- [ ] Error handling implemented
- [ ] Edge cases considered

### Quality
- [ ] Unit tests written and passing
- [ ] Code style standards followed
- [ ] Security considerations addressed
- [ ] Performance impact assessed

### Documentation
- [ ] Code is self-documenting
- [ ] Necessary comments added
- [ ] Documentation updated
- [ ] Breaking changes communicated

---

## Final Thoughts

Clean code is not a set of rigid rules, but a mindset. Focus on:
- **Readability**: Code should be easy to understand
- **Maintainability**: Code should be easy to modify
- **Testability**: Code should be easy to test
- **Extensibility**: Code should be easy to extend

Remember: **"Any fool can write code that a computer can understand. Good programmers write code that humans can understand."** - Martin Fowler
