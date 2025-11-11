/**
 * 57. Lifting the field (Pull Up Field)
 *
 * BEFORE: Duplicate fields in subclasses
 */
class EmployeePullBefore {
    protected String name;
}

class ManagerPullBefore extends EmployeePullBefore {
    protected String name; // Duplicate
    private double budget;
}

class EngineerPullBefore extends EmployeePullBefore {
    protected String name; // Duplicate
    private String skills;
}

/**
 * AFTER: Pull up field to superclass
 */
class EmployeePullAfter {
    protected String name;
}

class ManagerPullAfter extends EmployeePullAfter {
    private double budget;
}

class EngineerPullAfter extends EmployeePullAfter {
    private String skills;
}

/**
 * 58. Lifting the method (Pull Up Method)
 *
 * BEFORE: Duplicate methods in subclasses
 */
abstract class ShapeMethodBefore {
    abstract double area();
}

class CircleMethodBefore extends ShapeMethodBefore {
    private double radius;

    public double area() {
        return Math.PI * radius * radius;
    }

    public double circumference() {
        return 2 * Math.PI * radius;
    }
}

class SquareMethodBefore extends ShapeMethodBefore {
    private double side;

    public double area() {
        return side * side;
    }

    public double circumference() { // Could be duplicate logic
        return 4 * side;
    }
}

/**
 * AFTER: Pull up method to superclass
 */
abstract class ShapeMethodAfter {
    abstract double area();
    abstract double circumference();
}

class CircleMethodAfter extends ShapeMethodAfter {
    private double radius;

    public double area() {
        return Math.PI * radius * radius;
    }

    public double circumference() {
        return 2 * Math.PI * radius;
    }
}

class SquareMethodAfter extends ShapeMethodAfter {
    private double side;

    public double area() {
        return side * side;
    }

    public double circumference() {
        return 4 * side;
    }
}

/**
 * 59. Pushing down the field (Push Down Field)
 *
 * BEFORE: Field only used in one subclass
 */
class EmployeePushBefore {
    protected String name;
    protected double salary; // Only used by Manager
}

class ManagerPushBefore extends EmployeePushBefore {
    // Uses salary
}

class EngineerPushBefore extends EmployeePushBefore {
    // Doesn't use salary
}

/**
 * AFTER: Push down field to subclass that uses it
 */
class EmployeePushAfter {
    protected String name;
}

class ManagerPushAfter extends EmployeePushAfter {
    private double salary;
}

class EngineerPushAfter extends EmployeePushAfter {
    // No salary field
}

/**
 * 60. Pushing down the method (Push Down Method)
 *
 * BEFORE: Method only used in one subclass
 */
abstract class ShapePushBefore {
    abstract double area();

    protected double getPi() { // Only used by Circle
        return Math.PI;
    }
}

class CirclePushBefore extends ShapePushBefore {
    private double radius;

    public double area() {
        return getPi() * radius * radius;
    }
}

class SquarePushBefore extends ShapePushBefore {
    private double side;

    public double area() {
        return side * side;
    }
}

/**
 * AFTER: Push down method to subclass that uses it
 */
abstract class ShapePushAfter {
    abstract double area();
}

class CirclePushAfter extends ShapePushAfter {
    private double radius;

    public double area() {
        return Math.PI * radius * radius;
    }

    private double getPi() {
        return Math.PI;
    }
}

class SquarePushAfter extends ShapePushAfter {
    private double side;

    public double area() {
        return side * side;
    }
}

/**
 * 61. Extracting a subclass (Extract Subclass)
 *
 * BEFORE: Class has conditional behavior
 */
class EmployeeExtractBefore {
    private String name;
    private boolean isManager;
    private double budget; // Only for managers

    public double getSalary() {
        return isManager ? 50000 : 40000;
    }
}

/**
 * AFTER: Extract subclass for special behavior
 */
class EmployeeExtractAfter {
    protected String name;

    public double getSalary() {
        return 40000;
    }
}

class ManagerExtractAfter extends EmployeeExtractAfter {
    private double budget;

    public double getSalary() {
        return 50000;
    }
}

/**
 * 62. Extracting a superclass (Extract Superclass)
 *
 * BEFORE: Duplicate code in related classes
 */
class EmployeeSuperBefore {
    private String name;

    public String getName() {
        return name;
    }
}

class DepartmentSuperBefore {
    private String name;

    public String getName() {
        return name;
    }
}

/**
 * AFTER: Extract superclass
 */
abstract class NamedEntity {
    protected String name;

    public String getName() {
        return name;
    }
}

class EmployeeSuperAfter extends NamedEntity {
    // Employee-specific fields and methods
}

class DepartmentSuperAfter extends NamedEntity {
    // Department-specific fields and methods
}

/**
 * 63. Extracting an interface (Extract Interface)
 *
 * BEFORE: Clients depend on concrete classes
 */
class PrinterBefore {
    public void print(String document) {
        System.out.println("Printing: " + document);
    }
}

class ClientBefore {
    private PrinterBefore printer;

    public void doPrint(String doc) {
        printer.print(doc);
    }
}

/**
 * AFTER: Extract interface
 */
interface Printable {
    void print(String document);
}

class PrinterAfter implements Printable {
    public void print(String document) {
        System.out.println("Printing: " + document);
    }
}

class ClientAfter {
    private Printable printer;

    public void doPrint(String doc) {
        printer.print(doc);
    }
}

/**
 * 64. Folding the class (Collapse Hierarchy)
 *
 * BEFORE: Unnecessary class hierarchy
 */
abstract class EmployeeCollapseBefore {
    protected String name;
}

class ManagerCollapseBefore extends EmployeeCollapseBefore {
    private double budget;
}

/**
 * AFTER: Collapse hierarchy
 */
class EmployeeCollapseAfter {
    private String name;
    private double budget; // If manager
    private boolean isManager;
}

/**
 * 65. Creating a template method (Form Template Method)
 *
 * BEFORE: Duplicate algorithm structure
 */
abstract class DocumentProcessBefore {
    public final void process() {
        openFile();
        readData();
        processData();
        saveResults();
        closeFile();
    }

    abstract void readData();
    abstract void processData();

    private void openFile() { /* common */ }
    private void saveResults() { /* common */ }
    private void closeFile() { /* common */ }
}

class PDFProcessorBefore extends DocumentProcessBefore {
    void readData() { /* PDF specific */ }
    void processData() { /* PDF specific */ }
}

/**
 * AFTER: Template method pattern (same as before, just showing the pattern)
 */
abstract class DocumentProcessAfter {
    public final void process() {
        openFile();
        readData();
        processData();
        saveResults();
        closeFile();
    }

    abstract void readData();
    abstract void processData();

    private void openFile() { /* common */ }
    private void saveResults() { /* common */ }
    private void closeFile() { /* common */ }
}

class PDFProcessorAfter extends DocumentProcessAfter {
    void readData() { /* PDF specific */ }
    void processData() { /* PDF specific */ }
}
