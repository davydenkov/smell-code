# Code Smells and Refactoring Examples in PHP

A comprehensive collection of PHP examples demonstrating **code smells** (anti-patterns) and their corresponding **refactoring solutions**. This project serves as a practical guide for developers to recognize problematic code patterns and apply appropriate refactoring techniques in PHP.

## 📚 What You'll Learn

- **Identify Code Smells**: Learn to recognize 10 common code anti-patterns
- **Apply Refactoring Techniques**: Master 72 refactoring methods from Martin Fowler's book
- **See Real Examples**: Before/after code comparisons in PHP
- **Understand Best Practices**: Learn maintainable, clean PHP code principles

## 🏗️ Project Structure

```
php/
├── code-duplication/           # DRY principle violations
│   ├── bad/                       # Duplicated calculation logic
│   ├── good/                       # Extracted into reusable classes
│   └── README.md                  # Detailed explanation
├── data-classes/               # Classes with only data, no behavior
│   ├── bad/                       # User class with only getters/setters
│   ├── good/                       # User class with validation and behavior
│   └── README.md                  # Detailed explanation
├── data-clumps/                # Groups of data that should be together
│   ├── bad/                       # Individual parameters for address fields
│   └── good/                       # Address and Person classes
├── divergent-modifications/    # Classes changed for different reasons
│   ├── bad/                       # Single class handling multiple concerns
│   └── good/                       # Split into focused classes
├── feature-envy/               # Methods using other classes too much
│   ├── bad/                       # GeometryUtils accessing Rectangle data
│   └── good/                       # Methods moved to Rectangle class
├── incompleteness-of-library-class/  # Incomplete utility classes
│   ├── bad/                       # HttpClient with limited functionality
│   └── good/                       # Complete HttpClient implementation
├── large-class/                # Classes doing too many things
│   ├── bad/                       # UserService handling everything
│   ├── good/                       # Split into focused services
│   └── README.md                  # Detailed explanation
├── long-method/                # Methods that are too long
│   ├── bad/                       # Single method doing user registration
│   ├── good/                       # Broken down into smaller methods
│   └── README.md                  # Detailed explanation
├── long-parameters/            # Methods with too many parameters
│   ├── bad/                       # Methods with 20+ parameters
│   └── good/                       # Grouped into parameter objects
├── refactoring-methods/        # 72 refactoring techniques
│   ├── 01-extract-method.php     # Extract Method refactoring
│   ├── 02-variable-refactoring.php # Variable refactoring techniques
│   ├── 03-moving-features.php    # Moving methods and fields
│   ├── 04-data-organization.php  # Data restructuring
│   ├── 05-conditional-expressions.php # Simplifying conditionals
│   ├── 06-method-calls.php       # Method signature improvements
│   ├── 07-generalization-problems.php # Inheritance issues
│   ├── 08-major-refactorings.php # Large-scale refactorings
│   └── README.md                 # Complete catalog of techniques
└── renunciation-of-inheritance/ # Inheritance misused
    ├── bad/                       # Forced inheritance hierarchy
    └── good/                       # Composition over inheritance
```

## 🔍 Code Smells Covered

Each smell directory contains:
- **`bad/`** - Examples of problematic code
- **`good/`** - Refactored solutions
- **README.md** - Detailed explanation (in some directories)

### 1. Code Duplication (`code-duplication/`)
**Problem**: Same code appears in multiple places
**Solution**: Extract common functionality into reusable classes

### 2. Data Classes (`data-classes/`)
**Problem**: Classes with only fields and getters/setters, no behavior
**Solution**: Move behavior into data classes or extract logic elsewhere

### 3. Data Clumps (`data-clumps/`)
**Problem**: Groups of data items that always appear together
**Solution**: Create classes to hold the related data

### 4. Divergent Modifications (`divergent-modifications/`)
**Problem**: One class modified for many different reasons
**Solution**: Split the class into separate classes by responsibility

### 5. Feature Envy (`feature-envy/`)
**Problem**: Method uses more features of another class than its own
**Solution**: Move the method to the class it uses most

### 6. Incompleteness of Library Class (`incompleteness-of-library-class/`)
**Problem**: Library/utility classes lack needed functionality
**Solution**: Extend the class or create wrapper methods

### 7. Large Class (`large-class/`)
**Problem**: Classes trying to do too many things
**Solution**: Extract classes and move methods to appropriate places

### 8. Long Method (`long-method/`)
**Problem**: Methods that are too long and complex
**Solution**: Break down into smaller, focused methods

### 9. Long Parameters (`long-parameters/`)
**Problem**: Methods with too many parameters
**Solution**: Create parameter objects or use method chaining

### 10. Renunciation of Inheritance (`renunciation-of-inheritance/`)
**Problem**: Inheritance used incorrectly or not used when appropriate
**Solution**: Use composition instead of inheritance, or proper inheritance hierarchies

## 🔧 Refactoring Methods (`refactoring-methods/`)

Comprehensive examples of **72 refactoring techniques** from Martin Fowler's *Refactoring: Improving the Design of Existing Code*:

### Categories:
- **Method Extraction** - Breaking down large methods
- **Variable Refactoring** - Improving variable usage
- **Moving Features** - Relocating code to appropriate places
- **Data Organization** - Restructuring data and classes
- **Conditional Expressions** - Simplifying complex conditionals
- **Method Calls** - Improving method signatures and calls
- **Generalization Problems** - Fixing inheritance issues
- **Major Refactorings** - Large-scale architectural changes

## 🚀 Getting Started

### Prerequisites
- PHP 7.4 or higher
- Basic understanding of OOP concepts

### Exploring Examples

1. **Choose a code smell** you're interested in
2. **Read the README.md** in that directory for explanation (where available)
3. **Compare bad/ vs good/** examples
4. **Run the examples** to see them in action

```bash
# Example: Run code duplication examples
cd php/code-duplication/bad
php OrderProcessor.php

cd ../good
php OrderProcessor.php
php TaxCalculator.php

# Example: Run refactoring method examples
cd ../refactoring-methods
php 01-extract-method.php
```

### Learning Path

1. **Start with Code Smells** - Learn to identify problems
2. **Study Refactoring Methods** - Learn specific techniques
3. **Practice** - Apply techniques to your own PHP code
4. **Review** - Compare before/after implementations

## 📖 Documentation

- **[Main Code Smells Guide](../CODE_SMELLS.md)** - Detailed explanations of each smell
- **[Refactoring Catalog](./refactoring-methods/README.md)** - All 72 refactoring techniques
- **[Best Practices](../BEST_PRACTICES.md)** - Clean code principles

## 🎯 Learning Objectives

After studying these examples, you'll be able to:

- **Recognize** problematic code patterns in PHP quickly
- **Apply** appropriate refactoring techniques confidently in PHP
- **Write** cleaner, more maintainable PHP code
- **Understand** the principles behind clean code in PHP
- **Refactor** legacy PHP code safely and effectively

## 🐘 PHP-Specific Features Used

These examples demonstrate PHP best practices and modern features:

- **Type Declarations**: Scalar types, return types, and strict typing
- **Namespaces**: Proper organization and autoloading
- **Traits**: Code reuse without inheritance
- **Anonymous Classes**: For testing and flexibility
- **Null Coalescing Operator**: `??` for cleaner null handling
- **Spaceship Operator**: `<=>` for comparisons
- **Array Destructuring**: List assignment and spread operator
- **Generators**: Memory-efficient iteration
- **PSR Standards**: Following PHP community standards

## 🤝 Contributing

Found a bug or want to add more PHP examples?

1. Fork the repository
2. Create a feature branch
3. Add your examples with proper documentation
4. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](../LICENSE).

## 📚 Further Reading

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler
- **Clean Code: A Handbook of Agile Software Craftsmanship** by Robert C. Martin
- **PHP: The Right Way** - Community best practices
- **Modern PHP** by Josh Lockhart
- **PHP 8 Objects, Patterns, and Practice** by Matt Zandstra

---

**Happy refactoring with PHP!** 🐘
