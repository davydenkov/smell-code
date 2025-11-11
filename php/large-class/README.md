# Large Class

## Problem

A class that has grown too large and is trying to handle multiple responsibilities. This violates the **Single Responsibility Principle** and makes the code hard to understand, test, and maintain.

## Why It's Bad

- **Too many responsibilities**: Class does too many different things
- **Hard to understand**: Overwhelming amount of code in one place
- **Difficult to test**: Complex testing scenarios
- **Fragile**: Changes in one area can break other functionality
- **High coupling**: Many dependencies and interactions

## Example Scenario

A `UserService` class that handles:
- User CRUD operations
- Email sending
- Payment processing
- Report generation
- Data validation
- File operations

## Before Refactoring

```php
<?php
class UserService {
    // User management (CRUD operations)
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

    // File operations
    public function uploadAvatar($user, $file) { /* ... */ }
}
```

## After Refactoring

```php
<?php
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
    public function validateData($data) { /* ... */ }
}

class FileService {
    public function uploadAvatar($user, $file) { /* ... */ }
}
```

## Refactoring Techniques Used

1. **Extract Class**: Split large class into smaller, focused classes
2. **Move Method**: Move methods to classes that have the most relevant data
3. **Extract Interface**: Define clear contracts between classes
4. **Dependency Injection**: Pass dependencies rather than creating them

## Files in This Example

- `bad/UserService.php` - The original large class with multiple responsibilities
- `good/User.php` - User entity with validation behavior
- `good/UserRepository.php` - Data access layer
- `good/EmailService.php` - Email communication
- `good/PaymentService.php` - Payment processing
- `good/ReportService.php` - Report generation
- `good/Utility.php` - Shared utility functions

## Signs of Large Class Smell

- **Class with more than 200-300 lines**
- **Class with many instance variables**
- **Class with too many methods**
- **Class that uses "and" in its name** (e.g., "UserManagerAndEmailService")
- **Class that depends on many other classes**

## Refactoring Strategy

1. **Identify responsibilities**: List all the things the class does
2. **Group related methods**: Find methods that work together
3. **Extract classes**: Create new classes for each responsibility group
4. **Move fields**: Move relevant fields to the new classes
5. **Update dependencies**: Adjust classes that use the original class

## Running the Examples

```bash
# Run the bad example
php bad/UserService.php

# Run the good examples
php good/User.php
php good/UserRepository.php
php good/EmailService.php
php good/PaymentService.php
php good/ReportService.php
```

## Key Takeaways

- Aim for classes with single, well-defined responsibilities
- Use the "Extract Class" refactoring when a class grows too large
- Consider the Interface Segregation Principle
- Test each extracted class independently
- Use composition to combine small classes when needed
