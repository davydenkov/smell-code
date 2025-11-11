# C++ Refactoring Examples

A collection of C++ examples demonstrating **refactoring techniques** from Martin Fowler's book "Refactoring: Improving the Design of Existing Code". While this directory currently focuses on refactoring methods, the concepts can be applied to address the 10 common code smells covered in the main project.

## 📚 What You'll Learn

- **Apply Refactoring Techniques**: Master refactoring methods adapted for C++
- **See Real Examples**: C++ implementations of refactoring patterns
- **Understand Best Practices**: Learn maintainable, clean C++ code principles
- **Modern C++ Features**: Examples using C++17 features and idioms

## 🏗️ Project Structure

```
cpp/
├── refactoring-methods/       # Refactoring techniques in C++
│   └── 08-major-refactorings.cpp # Major refactoring patterns
└── README.md                 # This file
```

## 🔧 Refactoring Methods (`refactoring-methods/`)

Examples of **refactoring techniques** adapted for C++ from Martin Fowler's book:

### Categories:
- **Method Extraction** - Breaking down large methods
- **Variable Refactoring** - Improving variable usage
- **Moving Features** - Relocating code to appropriate places
- **Data Organization** - Restructuring data and classes
- **Conditional Expressions** - Simplifying complex conditionals
- **Method Calls** - Improving method signatures and calls
- **Generalization Problems** - Fixing inheritance issues
- **Major Refactorings** - Large-scale architectural changes

### Available Examples:
- **08-major-refactorings.cpp** - Large-scale refactoring patterns including:
  - Tease Apart Inheritance
  - Convert Procedural Design to Objects
  - Separate Domain from Presentation
  - Extract Hierarchy

## 🔍 Code Smells Coverage

While C++-specific code smell examples are not yet available in this directory, you can apply the refactoring techniques here to address the 10 common code smells:

1. **Code Duplication** - Extract common functionality into reusable classes
2. **Data Classes** - Add behavior to classes that only contain data
3. **Data Clumps** - Create structs/classes for related data groups
4. **Divergent Modifications** - Split classes by responsibility
5. **Feature Envy** - Move methods to classes they use most
6. **Incompleteness of Library Class** - Extend or wrap utility classes
7. **Large Class** - Extract smaller, focused classes
8. **Long Method** - Break down into smaller methods
9. **Long Parameters** - Create parameter objects
10. **Renunciation of Inheritance** - Use composition over inheritance

## 🚀 Getting Started

### Prerequisites
- C++17 or higher
- Basic understanding of OOP concepts in C++

### Exploring Examples

1. **Choose a refactoring technique** you're interested in
2. **Read the code comments** for explanation
3. **Compare before/after** implementations
4. **Compile and run the examples** to see them in action

```bash
# Example: Compile the major refactorings example (as library)
cd cpp/refactoring-methods
g++ -std=c++17 -c 08-major-refactorings.cpp -o major_refactorings.o
# Note: This creates object files with class definitions.
# To create executable programs, you would add a main() function.
```

### Learning Path

1. **Study Refactoring Methods** - Learn specific techniques in C++
2. **Practice** - Apply techniques to your own C++ code
3. **Review** - Compare before/after implementations

## 📖 Documentation

- **[Main Code Smells Guide](../CODE_SMELLS.md)** - Detailed explanations of each smell
- **[Refactoring Catalog](../refactoring-methods/README.md)** - All refactoring techniques
- **[Best Practices](../BEST_PRACTICES.md)** - Clean code principles

## 🎯 Learning Objectives

After studying these examples, you'll be able to:

- **Recognize** problematic code patterns in C++ quickly
- **Apply** appropriate refactoring techniques confidently in C++
- **Write** cleaner, more maintainable C++ code
- **Understand** the principles behind clean code in C++
- **Refactor** legacy C++ code safely and effectively

## 📚 Further Reading

- **Refactoring: Improving the Design of Existing Code** by Martin Fowler
- **Clean Code: A Handbook of Agile Software Craftsmanship** by Robert C. Martin
- **Effective Modern C++** by Scott Meyers
- **The C++ Programming Language** by Bjarne Stroustrup

---

**Happy refactoring with C++!** 🎉
