# Code Smells and Refactoring Examples in Go

A comprehensive collection of Go examples demonstrating **code smells** (anti-patterns) and their corresponding **refactoring solutions**. This project serves as a practical guide for developers to recognize problematic code patterns and apply appropriate refactoring techniques in Go.

## 📚 What You'll Learn

- **Identify Code Smells**: Learn to recognize 10 common code anti-patterns
- **Apply Refactoring Techniques**: Master refactoring methods adapted for Go
- **See Real Examples**: Before/after code comparisons in Go
- **Understand Best Practices**: Learn maintainable, clean Go code principles

## 🏗️ Project Structure

```
golang/
├── code-duplication/           # DRY principle violations
│   ├── bad/                       # Duplicated calculation logic
│   └── good/                       # Extracted into reusable packages
├── data-classes/               # Structs with only data, no behavior
│   ├── bad/                       # User struct with only fields
│   └── good/                       # User struct with validation and behavior
├── data-clumps/                # Groups of data that should be together
│   ├── bad/                       # Individual parameters for address fields
│   └── good/                       # Address and Person structs
├── divergent-modifications/    # Structs changed for different reasons
│   ├── bad/                       # Single struct handling multiple concerns
│   └── good/                       # Split into focused structs with interfaces
├── feature-envy/               # Methods using other structs too much
│   ├── bad/                       # GeometryUtils accessing Rectangle data
│   └── good/                       # Methods moved to Rectangle struct
├── incompleteness-of-library-class/  # Incomplete utility structs
│   ├── bad/                       # HttpClient with limited functionality
│   └── good/                       # Complete HttpClient implementation
├── large-class/                # Structs doing too many things
│   ├── bad/                       # UserService handling everything
│   └── good/                       # Split into focused services
├── long-method/                # Functions that are too long
│   ├── bad/                       # Single function doing user registration
│   └── good/                       # Broken down into smaller functions
├── long-parameters/            # Functions with too many parameters
│   ├── bad/                       # Functions with 20+ parameters
│   └── good/                       # Grouped into parameter structs
├── refactoring-methods/        # 72 refactoring techniques in Go
│   ├── 01-extract-method.go      # Extract Method refactoring
│   ├── 02-variable-refactoring.go # Variable refactoring techniques
│   ├── 03-moving-features.go     # Moving methods and fields
│   ├── 04-data-organization.go   # Data restructuring
│   ├── 05-conditional-expressions.go # Simplifying conditionals
│   ├── 06-method-calls.go        # Method signature improvements
│   ├── 07-generalization-problems.go # Interface and composition issues
│   └── 08-major-refactorings.go  # Large-scale refactorings
└── renunciation-of-inheritance/ # Inheritance misused
    ├── bad/                       # Forced embedding hierarchy
    └── good/                       # Composition and interfaces
```

## 🔍 Code Smells Covered

Each smell directory contains:
- **`bad/`** - Examples of problematic code
- **`good/`** - Refactored solutions

### 1. Code Duplication (`code-duplication/`)
**Problem**: Same code appears in multiple places
**Solution**: Extract common functionality into reusable packages

### 2. Data Classes (`data-classes/`)
**Problem**: Structs with only fields and getters/setters, no behavior
**Solution**: Move behavior into data structs or extract logic elsewhere

### 3. Data Clumps (`data-clumps/`)
**Problem**: Groups of data items that always appear together
**Solution**: Create structs to hold the related data

### 4. Divergent Modifications (`divergent-modifications/`)
**Problem**: One struct modified for many different reasons
**Solution**: Split the struct into separate structs by responsibility

### 5. Feature Envy (`feature-envy/`)
**Problem**: Method uses more features of another struct than its own
**Solution**: Move the method to the struct it uses most

### 6. Incompleteness of Library Class (`incompleteness-of-library-class/`)
**Problem**: Library/utility structs lack needed functionality
**Solution**: Extend the struct or create wrapper methods

### 7. Large Class (`large-class/`)
**Problem**: Structs trying to do too many things
**Solution**: Extract structs and move methods to appropriate places

### 8. Long Method (`long-method/`)
**Problem**: Functions that are too long and complex
**Solution**: Break down into smaller, focused functions

### 9. Long Parameters (`long-parameters/`)
**Problem**: Functions with too many parameters
**Solution**: Create parameter structs or use variadic functions

### 10. Renunciation of Inheritance (`renunciation-of-inheritance/`)
**Problem**: Inheritance patterns used incorrectly in Go
**Solution**: Use composition and interfaces instead of embedding hierarchies

## 🔧 Refactoring Methods (`refactoring-methods/`)

Comprehensive examples of **72 refactoring techniques** adapted for Go from Martin Fowler's *Refactoring: Improving the Design of Existing Code*:

### Categories:
- **01-extract-method.go** - Extract Method, Inline Method
- **02-variable-refactoring.go** - Inline Temp, Replace Temp with Query, Introduce Explaining Variable, Split Temporary Variable, Remove Assignments to Parameters, Replace Method with Method Object
- **03-moving-features.go** - Substitute Algorithm, Move Method, Move Field, Extract Class, Inline Class, Hide Delegate, Remove Middle Man, Introduce Foreign Method, Introduce Local Extension
- **04-data-organization.go** - Self-Encapsulate Field, Replace Data Value with Object, Change Value to Reference, Change Reference to Value, Replace Array with Object, Duplicate Observed Data, Change Unidirectional to Bidirectional Association, Change Bidirectional to Unidirectional Association, Replace Magic Number with Symbolic Constant
- **05-conditional-expressions.go** - Decompose Conditional, Consolidate Conditional Expression, Consolidate Duplicate Conditional Fragments, Remove Control Flag, Replace Nested Conditional with Guard Clauses, Replace Conditional with Polymorphism, Introduce Null Object, Introduce Assertion
- **06-method-calls.go** - Rename Method, Add Parameter, Remove Parameter, Separate Query from Modifier, Parameterize Method, Replace Parameter with Explicit Methods, Preserve Whole Object, Replace Parameter with Method, Introduce Parameter Object, Remove Setting Method, Hide Method, Replace Constructor with Factory Method
- **07-generalization-problems.go** - Pull Up Field, Pull Up Method, Pull Up Constructor Body, Push Down Method, Push Down Field, Extract Subclass, Extract Superclass, Extract Interface, Collapse Hierarchy
- **08-major-refactorings.go** - Tease Apart Inheritance, Convert Procedural Design to Objects, Separate Domain from Presentation, Extract Hierarchy

## 🚀 Getting Started

### Prerequisites
- Go 1.18 or higher
- Basic understanding of Go concepts

### Exploring Examples

1. **Choose a code smell** you're interested in
2. **Read the code comments** for explanation
3. **Compare bad/ vs good/** examples
4. **Run the examples** to see them in action

```bash
# Example: Run code duplication examples
cd golang/code-duplication/bad
go run order_processor.go

cd ../good
go run order_processor.go
go run tax_calculator.go

# Example: Run refactoring method examples
cd ../../refactoring-methods
go run 01-extract-method.go
go run 02-variable-refactoring.go
go run 03-moving-features.go
go run 04-data-organization.go
go run 05-conditional-expressions.go
go run 06-method-calls.go
go run 07-generalization-problems.go
go run 08-major-refactorings.go
```

### Learning Path

1. **Start with Code Smells** - Learn to identify problems
2. **Study Refactoring Methods** - Learn specific techniques
3. **Practice** - Apply techniques to your own Go code
4. **Review** - Compare before/after implementations

## 📖 Documentation

- **[Main Code Smells Guide](../CODE_SMELLS.md)** - Detailed explanations of each smell
- **[Refactoring Catalog](../refactoring-methods/README.md)** - All refactoring techniques
- **[Best Practices](../BEST_PRACTICES.md)** - Clean code principles

## 🎯 Learning Objectives

After studying these examples, you'll be able to:

- **Recognize** problematic code patterns in Go quickly
- **Apply** appropriate refactoring techniques confidently in Go
- **Write** cleaner, more maintainable Go code
- **Understand** the principles behind clean code in Go
- **Refactor** legacy Go code safely and effectively

## 🐹 Go-Specific Features Used

These examples demonstrate Go best practices and idioms:

- **Structs and Methods**: Go's approach to object-oriented programming
- **Interfaces**: Implicit interfaces and dependency injection
- **Composition**: Embedding structs for code reuse
- **Error Handling**: Go's explicit error handling patterns
- **Pointers**: Proper use of pointers for mutation
- **Slices and Maps**: Go's built-in collection types
- **Goroutines and Channels**: Concurrent programming patterns
- **Packages**: Modular code organization
- **Type Assertions**: Working with interface types safely

## 🌍 Translation Notes

The Go translations follow these conventions:

- **No Classes**: Used structs with methods instead of classes
- **No Inheritance**: Used interfaces and composition instead of inheritance
- **No Exceptions**: Used error return values instead of exceptions
- **No Abstract Classes**: Used interfaces to define contracts
- **Pointers**: Used pointers where mutation is needed
- **Embedding**: Used struct embedding for composition
- **Naming**: Go naming conventions (exported/unexported)
- **Error Handling**: Go's idiomatic error handling

## 🤝 Contributing

Found a bug or want to add more Go examples?

1. Fork the repository
2. Create a feature branch
3. Add your examples with proper documentation
4. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](../LICENSE).

## 📚 Further Reading

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler
- **Clean Code: A Handbook of Agile Software Craftsmanship** by Robert C. Martin
- **The Go Programming Language** by Alan Donovan and Brian Kernighan
- **Effective Go** - Official Go documentation
- **Go in Action** by William Kennedy

---

**Happy refactoring with Go!** 🐹
