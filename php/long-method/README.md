# Long Method

## Problem

A method that has grown too long and complex, making it hard to understand, test, and maintain. Long methods often violate the **Single Responsibility Principle** at the method level.

## Why It's Bad

- **Hard to understand**: Too much logic in one place
- **Difficult to test**: Complex testing scenarios
- **Error prone**: Easy to introduce bugs when modifying
- **Poor reusability**: Can't reuse parts of the method logic
- **Maintenance nightmare**: Changes affect large portions of code

## Example Scenario

A user registration method that:
- Validates input data
- Creates user object
- Hashes password
- Saves to database
- Sends welcome email
- Logs the registration
- Handles error cases

## Before Refactoring

```php
<?php
class UserManager {
    public function processUserRegistration($userData) {
        // Validate input (10+ lines)
        if (empty($userData['name'])) {
            throw new Exception('Name is required');
        }
        if (!filter_var($userData['email'], FILTER_VALIDATE_EMAIL)) {
            throw new Exception('Invalid email');
        }
        if (strlen($userData['password']) < 8) {
            throw new Exception('Password too short');
        }

        // Create user object (5+ lines)
        $user = new User();
        $user->setName($userData['name']);
        $user->setEmail($userData['email']);

        // Hash password (3+ lines)
        $hashedPassword = password_hash($userData['password'], PASSWORD_DEFAULT);
        $user->setPassword($hashedPassword);

        // Save to database (5+ lines)
        $db = new Database();
        $userId = $db->insert('users', [
            'name' => $user->getName(),
            'email' => $user->getEmail(),
            'password' => $user->getPassword()
        ]);

        // Send welcome email (10+ lines)
        $emailService = new EmailService();
        $subject = 'Welcome to our platform!';
        $body = "Hello {$user->getName()}, welcome...";
        $emailService->sendEmail($user->getEmail(), $subject, $body);

        // Log registration (3+ lines)
        $logger = new Logger();
        $logger->info("User registered: {$user->getEmail()}");

        return $userId;
    }
}
```

## After Refactoring

```php
<?php
class UserManager {
    private $userRepository;
    private $emailService;
    private $logger;

    public function __construct(UserRepository $userRepo, EmailService $emailSvc, Logger $logger) {
        $this->userRepository = $userRepo;
        $this->emailService = $emailSvc;
        $this->logger = $logger;
    }

    public function processUserRegistration($userData) {
        $user = $this->createUser($userData);
        $this->saveUser($user);
        $this->sendWelcomeEmail($user);
        $this->logRegistration($user);
        return $user->getId();
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
        if (!filter_var($userData['email'], FILTER_VALIDATE_EMAIL)) {
            throw new Exception('Invalid email');
        }
        if (strlen($userData['password']) < 8) {
            throw new Exception('Password too short');
        }
    }

    private function hashPassword($password) {
        return password_hash($password, PASSWORD_DEFAULT);
    }

    private function saveUser(User $user) {
        return $this->userRepository->save($user);
    }

    private function sendWelcomeEmail(User $user) {
        $this->emailService->sendWelcomeEmail($user);
    }

    private function logRegistration(User $user) {
        $this->logger->info("User registered: {$user->getEmail()}");
    }
}
```

## Refactoring Techniques Used

1. **Extract Method**: Break down long method into smaller methods
2. **Replace Temp with Query**: Replace temporary variables with method calls
3. **Introduce Explaining Variable**: Add variables to clarify complex expressions
4. **Move Method**: Move methods to more appropriate classes

## Files in This Example

- `bad/UserManager.php` - The original long method
- `good/UserManager.php` - Refactored with extracted methods
- `good/EmailService.php` - Email functionality
- `good/NotificationService.php` - Notification handling
- `good/UserRepository.php` - Data access
- `good/UserValidator.php` - Validation logic

## Signs of Long Method Smell

- **Method longer than 10-15 lines**
- **Method with multiple responsibilities**
- **Method with many local variables**
- **Method with complex nested conditionals**
- **Method that's hard to name** (needs "and" or "or" in name)

## Refactoring Strategy

1. **Identify responsibilities**: What does this method do?
2. **Look for comments**: Comments often indicate extractable methods
3. **Extract methods**: Create smaller methods for each responsibility
4. **Replace conditionals**: Use early returns or polymorphism
5. **Test frequently**: Ensure behavior is preserved after each extraction

## Method Size Guidelines

- **1-5 lines**: Ideal for simple methods
- **6-10 lines**: Acceptable for straightforward logic
- **11-20 lines**: Consider refactoring
- **20+ lines**: Strong candidate for refactoring

## Running the Examples

```bash
# Run the bad example
php bad/UserManager.php

# Run the good examples
php good/UserManager.php
php good/UserRepository.php
php good/EmailService.php
php good/UserValidator.php
```

## Key Takeaways

- Keep methods focused on a single responsibility
- Use descriptive names for extracted methods
- Consider method length as a warning sign, not a hard rule
- Extract methods when you see comments explaining sections of code
- Test after each refactoring to ensure functionality is preserved
