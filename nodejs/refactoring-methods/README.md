# Node.js Refactoring Methods Examples

This directory contains JavaScript examples for the refactoring techniques described in Martin Fowler's book "Refactoring: Improving the Design of Existing Code".

## Structure

The refactoring techniques are organized into 8 files based on categories:

### 01-extract-method.js
- **1. Method Extraction (Extract Method)** - Breaking down large methods into smaller, focused methods
- **2. Embedding a method (Inline Method)** - Removing unnecessary method indirection

### 02-variable-refactoring.js
- **3. Embedding a temporary variable (Inline Temp)** - Removing unnecessary temporary variables
- **4. Replacing a temporary variable with a method call (Replace Temp with Query)** - Converting temp variables to method calls
- **5. Introduction of an explanatory variable (Introduce Explaining Variable)** - Adding variables to clarify complex expressions
- **6. Splitting a Temporary Variable** - Separating variables used for different purposes
- **7. Removing parameter Assignments (Remove Assignments to Parameters)** - Avoiding parameter modification
- **8. Replacing a method with a method Object (Replace Method with Method Object)** - Converting methods to objects

### 03-moving-features.js
- **9. Substitution Algorithm** - Replacing complex algorithms with simpler ones
- **10. Moving functions between objects (Move Method)** - Relocating methods to appropriate classes
- **11. Moving the field (Move Field)** - Relocating fields to appropriate classes
- **12. Class Allocation (Extract Class)** - Splitting large classes into smaller ones
- **13. Embedding a class (Inline Class)** - Merging unnecessary classes
- **14. Hiding delegation (Hide Delegate)** - Encapsulating delegation relationships
- **15. Removing the intermediary (Remove Middle Man)** - Eliminating unnecessary delegation
- **16. Introduction of an external method (Introduce Foreign Method)** - Adding methods to external classes
- **17. The introduction of local extension (Introduce Local Extension)** - Extending external classes

### 04-data-organization.js
- **18. Self-Encapsulate Field** - Adding accessors for fields
- **19. Replacing the data value with an object (Replace Data Value with Object)** - Converting primitive values to objects
- **20. Replacing the value with a reference (Change Value to Reference)** - Using shared object references
- **21. Replacing a reference with a value (Change Reference to Value)** - Using independent object copies
- **22. Replacing an array with an object (Replace Array with Object)** - Converting arrays to structured objects
- **23. Duplication of visible data (Duplicate Observed Data)** - Separating domain and presentation data
- **24. Replacing Unidirectional communication with Bidirectional** - Adding bidirectional associations
- **25. Replacing Bidirectional communication with Unidirectional** - Removing bidirectional associations
- **26. Replacing the magic number with a symbolic constant** - Using named constants
- **27. Encapsulate Field** - Making fields private
- **28. Encapsulate Collection** - Protecting collection fields
- **29. Replacing a record with a Data Class** - Converting arrays to data classes
- **30. Replacing Type Code with Class** - Converting type codes to classes
- **31. Replacing Type Code with Subclasses** - Using inheritance for type codes
- **32. Replacing Type Code with State/Strategy** - Using composition for type codes
- **33. Replacing Subclass with Fields** - Converting subclasses to fields

### 05-conditional-expressions.js
- **34. Decomposition of a conditional operator (Decompose Conditional)** - Breaking down complex conditionals
- **35. Consolidation of a conditional expression (Consolidate Conditional Expression)** - Combining similar conditionals
- **36. Consolidation of duplicate conditional fragments** - Extracting common conditional code
- **37. Remove Control Flag** - Eliminating control flag variables
- **38. Replacing Nested Conditional statements with a boundary operator** - Using guard clauses
- **39. Replacing a conditional operator with polymorphism (Replace Conditional with Polymorphism)** - Using polymorphism instead of conditionals
- **40. Introduction of the object (Introduce Object)** - Creating result objects
- **41. Introduction of the statement (Introduction Statement)** - Adding assertion methods

### 06-method-calls.js
- **42. Renaming a method (Rename Method)** - Improving method names
- **43. Adding a parameter (Add Parameter)** - Extending method signatures
- **44. Deleting a parameter (Remove Parameter)** - Simplifying method signatures
- **45. Separation of Query and Modifier (Separate Query from Modifier)** - Splitting methods that query and modify
- **46. Parameterization of the method (Parameterize Method)** - Reducing similar methods
- **47. Replacing a parameter with explicit methods** - Using method names instead of parameters
- **48. Save the Whole Object** - Passing entire objects instead of individual fields
- **49. Replacing a parameter with a method call** - Using method calls instead of parameters
- **50. Introduction of the boundary object (Introduce Parameter Object)** - Grouping related parameters
- **51. Removing the Value Setting Method** - Removing unnecessary setters
- **52. Hiding a method (Hide Method)** - Reducing method visibility
- **53. Replacing the constructor with the factory method** - Using factory methods
- **54. Encapsulation of top-down type conversion** - Hiding type casting
- **55. Replacing the error code with an exceptional situation** - Using exceptions instead of error codes
- **56. Replacing an exceptional situation with a check** - Using checks instead of exceptions

### 07-generalization-problems.js
- **57. Lifting the field (Pull Up Field)** - Moving fields to superclasses
- **58. Lifting the method (Pull Up Method)** - Moving methods to superclasses
- **59. Lifting the constructor Body (Pull Up Constructor Body)** - Moving constructor code to superclasses
- **60. Method Descent (Push Down Method)** - Moving methods to subclasses
- **61. Field Descent (Push Down Field)** - Moving fields to subclasses
- **62. Subclass extraction (Extract Subclass)** - Creating subclasses for variations
- **63. Allocation of the parent class (Extract Superclass)** - Creating superclasses from common code
- **64. Interface extraction (Extract Interface)** - Creating interfaces from classes
- **65. Collapse Hierarchy** - Removing unnecessary inheritance levels
- **66. Formation of the method template (Form Template Method)** - Creating template methods
- **67. Replacement of inheritance by delegation** - Using composition instead of inheritance
- **68. Replacement of delegation by inheritance** - Using inheritance instead of composition

### 08-major-refactorings.js
- **69. Separation of inheritance (Tease Apart Inheritance)** - Splitting mixed inheritance hierarchies
- **70. Converting a procedural project into objects** - Object-oriented conversion
- **71. Separating the domain from the representation** - MVC-style separation
- **72. Hierarchy Extraction (Extract Hierarchy)** - Creating inheritance hierarchies

## Usage

Each file contains JavaScript classes demonstrating the "before" and "after" states of refactoring. The "before" examples show code that could benefit from refactoring, while the "after" examples show the improved code structure.

You can run these examples using Node.js:

```bash
node 01-extract-method.js
```

## Notes

- All examples are self-contained and don't require external dependencies (except where filesystem operations are demonstrated)
- The numbering corresponds to the refactoring techniques listed in the original book
- Each example includes comments explaining the refactoring technique
- Some examples may use simplified scenarios to focus on the refactoring concept
- JavaScript-specific adaptations have been made (e.g., using Map instead of associative arrays, classes instead of PHP-style classes)
