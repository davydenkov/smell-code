# Code Smells and Refactoring Examples in Python

A comprehensive collection of Python examples demonstrating **code smells** (anti-patterns) and their corresponding **refactoring solutions**. This project serves as a practical guide for developers to recognize problematic code patterns and apply appropriate refactoring techniques in Python.

## 📚 What You'll Learn

- **Identify Code Smells**: Learn to recognize 10 common code anti-patterns
- **Apply Refactoring Techniques**: Master refactoring methods adapted for Python
- **See Real Examples**: Before/after code comparisons in Python
- **Understand Best Practices**: Learn maintainable, clean Python code principles

## 🏗️ Project Structure

```
python/
├── code-duplication/           # DRY principle violations
│   ├── bad/                       # Duplicated calculation logic
│   └── good/                       # Extracted into reusable classes
├── data-classes/               # Classes with only data, no behavior
│   ├── bad/                       # User class with only attributes
│   └── good/                       # User class with validation and behavior
├── data-clumps/                # Groups of data that should be together
│   ├── bad/                       # Individual parameters for address fields
│   └── good/                       # Address and Person dataclasses
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
│   └── good/                       # Split into focused services
├── long-method/                # Methods that are too long
│   ├── bad/                       # Single method doing user registration
│   └── good/                       # Broken down into smaller methods
├── long-parameters/            # Methods with too many parameters
│   ├── bad/                       # Methods with 20+ parameters
│   └── good/                       # Grouped into parameter objects
├── refactoring-methods/        # Individual refactoring techniques
│   ├── bad/                       # Before refactoring examples
│   └── good/                       # After refactoring examples
└── renunciation-of-inheritance/ # Inheritance misused
    ├── bad/                       # Forced inheritance hierarchy
    └── good/                       # Composition over inheritance
```

## 🔍 Code Smells Covered

Each smell directory contains:
- **`bad/`** - Examples of problematic code
- **`good/`** - Refactored solutions

### 1. Code Duplication (`code-duplication/`)
**Problem**: Same code appears in multiple places
**Solution**: Extract common functionality into reusable classes

### 2. Data Classes (`data-classes/`)
**Problem**: Classes with only fields and getters/setters, no behavior
**Solution**: Move behavior into data classes or extract logic elsewhere

### 3. Data Clumps (`data-clumps/`)
**Problem**: Groups of data items that always appear together
**Solution**: Create dataclasses to hold the related data

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

Examples of **refactoring techniques** adapted for Python from Martin Fowler's book:

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
- Python 3.6 or higher
- Basic understanding of OOP concepts

### Exploring Examples

1. **Choose a code smell** you're interested in
2. **Read the code comments** for explanation
3. **Compare bad/ vs good/** examples
4. **Run the examples** to see them in action

```bash
# Example: Run code duplication examples
cd python/code-duplication/bad
python order_processor.py

cd ../good
python order_processor.py
python tax_calculator.py
```

### Learning Path

1. **Start with Code Smells** - Learn to identify problems
2. **Study Refactoring Methods** - Learn specific techniques
3. **Practice** - Apply techniques to your own Python code
4. **Review** - Compare before/after implementations

## 📖 Documentation

- **[Main Code Smells Guide](../CODE_SMELLS.md)** - Detailed explanations of each smell
- **[Refactoring Catalog](../refactoring-methods/README.md)** - All refactoring techniques
- **[Best Practices](../BEST_PRACTICES.md)** - Clean code principles

## 🎯 Learning Objectives

After studying these examples, you'll be able to:

- **Recognize** problematic code patterns in Python quickly
- **Apply** appropriate refactoring techniques confidently in Python
- **Write** cleaner, more maintainable Python code
- **Understand** the principles behind clean code in Python
- **Refactor** legacy Python code safely and effectively

## 🐍 Python-Specific Features Used

These examples demonstrate Python best practices and idioms:

- **Type Hints**: Using `typing` module for better code documentation
- **Dataclasses**: `@dataclass` decorator for clean data structures
- **Dictionaries**: Pythonic data structures for tax rates and configurations
- **List/Dict Comprehensions**: Clean, readable data transformations
- **Context Managers**: Proper resource management
- **Duck Typing**: Python's flexible type system
- **Properties**: `@property` decorator for computed attributes

## 🤝 Contributing

Found a bug or want to add more Python examples?

1. Fork the repository
2. Create a feature branch
3. Add your examples with proper documentation
4. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](../LICENSE).

## 📚 Further Reading

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler
- **Clean Code: A Handbook of Agile Software Craftsmanship** by Robert C. Martin
- **Python Cookbook** by David Beazley and Brian K. Jones
- **Fluent Python** by Luciano Ramalho
- **Effective Python** by Brett Slatkin

---

**Happy refactoring with Python!** 🐍
