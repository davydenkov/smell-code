# Data Classes

## Problem

Data classes are classes that contain only data (fields) and simple getters/setters, with no business logic. This creates **anemic domain models** where behavior is scattered across the application.

## Why It's Bad

- **Logic scattered**: Business rules spread across multiple classes
- **Encapsulation violation**: Data exposed without validation
- **Maintenance issues**: Changes require finding all usage locations
- **No data integrity**: No validation or business rules enforcement
- **Poor object-oriented design**: Objects without behavior

## Example Scenario

A `User` class that only holds data, while business logic is in services:

```php
// Data class - only holds data
class User {
    private $name;
    private $email;
    private $password;

    public function getName() { return $this->name; }
    public function setName($name) { $this->name = $name; }
    // ... more getters/setters
}

// Logic scattered elsewhere
class UserService {
    public function isValidEmail($email) { /* validation logic */ }
    public function hashPassword($password) { /* hashing logic */ }
}
```

## Before Refactoring

```php
<?php
class User {
    private $name;
    private $email;
    private $password;

    public function __construct($name, $email, $password) {
        $this->name = $name;
        $this->email = $email;
        $this->password = $password;
    }

    // Only getters and setters - no behavior
    public function getName() { return $this->name; }
    public function getEmail() { return $this->email; }
    public function getPassword() { return $this->password; }
}

// Business logic scattered in services
class UserService {
    public function createUser($data) {
        $user = new User($data['name'], $data['email'], $data['password']);
        // Validation logic here
        if (empty($data['name'])) {
            throw new Exception('Name required');
        }
        // More validation...
        return $user;
    }
}
```

## After Refactoring

```php
<?php
class User {
    private $name;
    private $email;
    private $password;

    public function __construct($name, $email, $password) {
        $this->setName($name);
        $this->setEmail($email);
        $this->setPassword($password);
    }

    public function getName() { return $this->name; }

    public function setName($name) {
        if (empty($name)) {
            throw new InvalidArgumentException('Name cannot be empty');
        }
        $this->name = $name;
    }

    public function getEmail() { return $this->email; }

    public function setEmail($email) {
        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            throw new InvalidArgumentException('Invalid email format');
        }
        $this->email = $email;
    }

    public function getPassword() { return $this->password; }

    public function setPassword($password) {
        if (strlen($password) < 8) {
            throw new InvalidArgumentException('Password must be at least 8 characters');
        }
        $this->password = $this->hashPassword($password);
    }

    private function hashPassword($password) {
        return password_hash($password, PASSWORD_DEFAULT);
    }

    public function changePassword($oldPassword, $newPassword) {
        if (!$this->verifyPassword($oldPassword)) {
            throw new InvalidArgumentException('Current password is incorrect');
        }
        $this->setPassword($newPassword);
    }

    public function verifyPassword($password) {
        return password_verify($password, $this->password);
    }

    public function getDisplayName() {
        return ucwords(strtolower($this->name));
    }
}
```

## Refactoring Techniques Used

1. **Move Method**: Move validation and business logic into the data class
2. **Encapsulate Field**: Make fields private and add validation to setters
3. **Introduce Assertions**: Add validation logic to constructors and setters

## Files in This Example

- `bad/User.php` - Anemic data class with no behavior
- `good/User.php` - Rich domain object with business logic and validation

## Signs of Data Class Smell

- **Class with only getters/setters**
- **Public fields** (should be private)
- **No validation in setters**
- **Business logic in service classes**
- **Data transfer objects** used for domain logic

## When Data Classes Are Acceptable

- **DTOs (Data Transfer Objects)**: For transferring data between layers
- **Configuration classes**: Simple key-value storage
- **External API responses**: When mapping external data
- **Database entities**: When using ORM frameworks

## Running the Examples

```bash
# Run the bad example
php bad/User.php

# Run the good example
php good/User.php
```

## Key Takeaways

- Move behavior close to the data it operates on
- Use constructors and setters for validation
- Create rich domain objects instead of anemic models
- Keep data encapsulated and protected
- Test domain logic within the objects themselves
