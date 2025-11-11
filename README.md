# Code Smells and Refactoring Examples

A comprehensive collection of **PHP** and **Python** examples demonstrating **code smells** (anti-patterns) and their corresponding **refactoring solutions**. This project serves as a practical guide for developers to recognize problematic code patterns and apply appropriate refactoring techniques.

## 📚 What You'll Learn

- **Identify Code Smells**: Learn to recognize 9 common code anti-patterns
- **Apply Refactoring Techniques**: Master 72 refactoring methods from Martin Fowler's book
- **See Real Examples**: Before/after code comparisons in PHP and Python
- **Understand Best Practices**: Learn maintainable, clean code principles

## 🏗️ Project Structure

```
smell-code/
├── php/                       # PHP implementations
│   ├── code-duplication/          # DRY principle violations
│   ├── data-classes/              # Classes with only data, no behavior
│   ├── data-clumps/               # Groups of data that should be together
│   ├── divergent-modifications/   # Classes changed for different reasons
│   ├── feature-envy/              # Methods using other classes too much
│   ├── incompleteness-of-library-class/  # Incomplete utility classes
│   ├── large-class/               # Classes doing too many things
│   ├── long-method/               # Methods that are too long
│   ├── long-parameters/           # Methods with too many parameters
│   ├── refactoring-methods/       # 72 refactoring techniques
│   └── renunciation-of-inheritance/  # Inheritance misused
├── python/                    # Python implementations (translated from PHP)
│   ├── code-duplication/          # DRY principle violations
│   ├── data-classes/              # Classes with only data, no behavior
│   ├── data-clumps/               # Groups of data that should be together
│   ├── divergent-modifications/   # Classes changed for different reasons
│   ├── feature-envy/              # Methods using other classes too much
│   ├── incompleteness-of-library-class/  # Incomplete utility classes
│   ├── large-class/               # Classes doing too many things
│   ├── long-method/               # Methods that are too long
│   ├── long-parameters/           # Methods with too many parameters
│   ├── refactoring-methods/       # 72 refactoring techniques
│   └── renunciation-of-inheritance/  # Inheritance misused
├── golang/                    # Go implementations (existing)
└── README.md                  # Documentation
```

## 🔍 Code Smells Covered

Each smell directory contains:
- **`bad/`** - Examples of problematic code
- **`good/`** - Refactored solutions
- **README.md** - Detailed explanation of the smell and fixes

### 1. Code Duplication (`code-duplication/`)
**Problem**: Same code appears in multiple places
**Solution**: Extract common functionality into reusable methods/classes

### 2. Data Classes (`data-classes/`)
**Problem**: Classes with only fields and getters/setters, no behavior
**Solution**: Move behavior into data classes or extract logic elsewhere

### 3. Data Clumps (`data-clumps/`)
**Problem**: Groups of data items that always appear together
**Solution**: Create a class to hold the related data

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
- PHP 7.4 or higher and/or Python 3.6 or higher
- Basic understanding of OOP concepts

### Exploring Examples

1. **Choose a code smell** you're interested in
2. **Read the README.md** in that directory for explanation
3. **Compare bad/ vs good/** examples
4. **Run the examples** to see them in action

```bash
# Example: Run PHP code duplication examples
cd php/code-duplication/bad
php OrderProcessor.php

cd ../good
php OrderProcessor.php
php TaxCalculator.php

# Example: Run Python code duplication examples
cd ../../../python/code-duplication/bad
python order_processor.py

cd ../good
python order_processor.py
python tax_calculator.py
```

### Learning Path

1. **Start with Code Smells** - Learn to identify problems
2. **Study Refactoring Methods** - Learn specific techniques
3. **Practice** - Apply techniques to your own code
4. **Review** - Compare before/after implementations

## 📖 Documentation

- **[Code Smells Guide](./CODE_SMELLS.md)** - Detailed explanations of each smell
- **[Refactoring Catalog](./refactoring-methods/README.md)** - All 72 refactoring techniques
- **[Best Practices](./BEST_PRACTICES.md)** - Clean code principles

## 🎯 Learning Objectives

After studying these examples, you'll be able to:

- **Recognize** problematic code patterns quickly
- **Apply** appropriate refactoring techniques confidently
- **Write** cleaner, more maintainable code
- **Understand** the principles behind clean code
- **Refactor** legacy code safely and effectively

## 🤝 Contributing

Found a bug or want to add more examples?

1. Fork the repository
2. Create a feature branch
3. Add your examples with proper documentation
4. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 📚 Further Reading

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler
- **Clean Code: A Handbook of Agile Software Craftsmanship** by Robert C. Martin
- **The Pragmatic Programmer** by Andrew Hunt and David Thomas

---

**Happy refactoring!** 🎉
