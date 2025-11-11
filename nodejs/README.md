# Code Smells and Refactoring Examples in Node.js

A comprehensive collection of Node.js/JavaScript examples demonstrating **code smells** (anti-patterns) and their corresponding **refactoring solutions**. This project serves as a practical guide for developers to recognize problematic code patterns and apply appropriate refactoring techniques in JavaScript.

## 📚 What You'll Learn

- **Identify Code Smells**: Learn to recognize 10 common code anti-patterns
- **Apply Refactoring Techniques**: Master refactoring methods adapted for JavaScript
- **See Real Examples**: Before/after code comparisons in Node.js
- **Understand Best Practices**: Learn maintainable, clean JavaScript code principles

## 🏗️ Project Structure

```
nodejs/
├── code-duplication/           # DRY principle violations
│   ├── bad/                       # Duplicated calculation logic
│   └── good/                       # Extracted into reusable modules
├── data-classes/               # Classes/objects with only data, no behavior
│   ├── bad/                       # User object with only properties
│   └── good/                       # User class with validation and behavior
├── data-clumps/                # Groups of data that should be together
│   ├── bad/                       # Individual parameters for address fields
│   └── good/                       # Address and Person classes
├── divergent-modifications/    # Classes changed for different reasons
│   ├── bad/                       # Single class handling multiple concerns
│   └── good/                       # Split into focused modules
├── feature-envy/               # Methods using other classes too much
│   ├── bad/                       # GeometryUtils accessing Rectangle data
│   └── good/                       # Methods moved to Rectangle class
├── incompleteness-of-library-class/  # Incomplete utility classes
│   ├── bad/                       # HttpClient with limited functionality
│   └── good/                       # Complete HttpClient implementation
├── large-class/                # Classes doing too many things
│   ├── bad/                       # UserService handling everything
│   └── good/                       # Split into focused services
├── long-method/                # Functions that are too long
│   ├── bad/                       # Single function doing user registration
│   └── good/                       # Broken down into smaller functions
├── long-parameters/            # Functions with too many parameters
│   ├── bad/                       # Functions with 20+ parameters
│   └── good/                       # Grouped into parameter objects
├── package.json                # Node.js dependencies
├── refactoring-methods/        # Individual refactoring techniques
│   ├── 01-extract-method.js      # Extract Method refactoring
│   ├── 02-variable-refactoring.js # Variable refactoring techniques
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
**Solution**: Extract common functionality into reusable modules

### 2. Data Classes (`data-classes/`)
**Problem**: Classes/objects with only data and getters/setters, no behavior
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
**Problem**: Functions that are too long and complex
**Solution**: Break down into smaller, focused functions

### 9. Long Parameters (`long-parameters/`)
**Problem**: Functions with too many parameters
**Solution**: Create parameter objects or use destructuring

### 10. Renunciation of Inheritance (`renunciation-of-inheritance/`)
**Problem**: Inheritance used incorrectly or not used when appropriate
**Solution**: Use composition instead of inheritance, or proper inheritance hierarchies

## 🔧 Refactoring Methods (`refactoring-methods/`)

Examples of **refactoring techniques** adapted for JavaScript from Martin Fowler's book:

### Categories:
- **Function Extraction** - Breaking down large functions
- **Variable Refactoring** - Improving variable usage
- **Moving Features** - Relocating code to appropriate places
- **Data Organization** - Restructuring data and classes
- **Conditional Expressions** - Simplifying complex conditionals
- **Function Calls** - Improving function signatures and calls
- **Generalization Problems** - Fixing inheritance issues
- **Major Refactorings** - Large-scale architectural changes

## 🚀 Getting Started

### Prerequisites
- Node.js 16.0 or higher
- npm or yarn
- Basic understanding of JavaScript OOP concepts

### Installation

```bash
cd nodejs
npm install
```

### Exploring Examples

1. **Choose a code smell** you're interested in
2. **Read the code comments** for explanation
3. **Compare bad/ vs good/** examples
4. **Run the examples** to see them in action

```bash
# Example: Run code duplication examples
cd nodejs/code-duplication/bad
node orderProcessor.js

cd ../good
node orderProcessor.js
node taxCalculator.js

# Example: Run refactoring method examples
cd ../refactoring-methods
node 01-extract-method.js
```

### Learning Path

1. **Start with Code Smells** - Learn to identify problems
2. **Study Refactoring Methods** - Learn specific techniques
3. **Practice** - Apply techniques to your own JavaScript code
4. **Review** - Compare before/after implementations

## 📖 Documentation

- **[Main Code Smells Guide](../CODE_SMELLS.md)** - Detailed explanations of each smell
- **[Refactoring Catalog](../refactoring-methods/README.md)** - All refactoring techniques
- **[Best Practices](../BEST_PRACTICES.md)** - Clean code principles

## 🎯 Learning Objectives

After studying these examples, you'll be able to:

- **Recognize** problematic code patterns in JavaScript quickly
- **Apply** appropriate refactoring techniques confidently in Node.js
- **Write** cleaner, more maintainable JavaScript code
- **Understand** the principles behind clean code in JavaScript
- **Refactor** legacy JavaScript code safely and effectively

## 📦 JavaScript/Node.js-Specific Features Used

These examples demonstrate JavaScript best practices and modern features:

- **ES6+ Classes**: Class syntax with proper inheritance
- **Modules**: ES6 import/export and CommonJS require/module.exports
- **Arrow Functions**: Concise function syntax and lexical `this`
- **Template Literals**: String interpolation with backticks
- **Destructuring**: Object and array destructuring
- **Spread/Rest Operators**: `...` for arrays and objects
- **Async/Await**: Asynchronous programming patterns
- **Promises**: Modern asynchronous handling
- **Map/Set**: Native collection types
- **Object Methods**: Modern object literal syntax

## 🌍 Translation Notes

The JavaScript translations follow these conventions:

- **Naming**: camelCase for variables/functions, PascalCase for classes
- **Modules**: ES6 import/export syntax where possible
- **Promises/Async**: Modern asynchronous patterns
- **Error Handling**: JavaScript Error objects and try/catch
- **Collections**: Arrays, Maps, Sets instead of PHP arrays
- **String Handling**: Template literals instead of concatenation
- **Null Safety**: Optional chaining (`?.`) and nullish coalescing (`??`)

## 📋 Dependencies

- **bcrypt** - For password hashing in authentication examples
- **nodemailer** - For email functionality in user registration examples

## 🤝 Contributing

Found a bug or want to add more Node.js examples?

1. Fork the repository
2. Create a feature branch
3. Add your examples with proper documentation
4. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](../LICENSE).

## 📚 Further Reading

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler
- **Clean Code: A Handbook of Agile Software Craftsmanship** by Robert C. Martin
- **Eloquent JavaScript** by Marijn Haverbeke
- **You Don't Know JS** series by Kyle Simpson
- **Node.js Design Patterns** by Mario Casciaro

---

**Happy refactoring with Node.js!** 📦
