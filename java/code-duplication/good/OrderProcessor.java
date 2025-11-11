class OrderProcessor {
    private TaxCalculator taxCalculator;
    private ShippingCalculator shippingCalculator;

    public OrderProcessor() {
        this.taxCalculator = new TaxCalculator();
        this.shippingCalculator = new ShippingCalculator();
    }

    public double calculateTax(double price, String state) {
        return taxCalculator.calculateTax(price, state);
    }

    public double calculateShipping(double weight, double distance) {
        return shippingCalculator.calculateShipping(weight, distance);
    }
}

class InvoiceGenerator {
    private TaxCalculator taxCalculator;
    private ShippingCalculator shippingCalculator;

    public InvoiceGenerator() {
        this.taxCalculator = new TaxCalculator();
        this.shippingCalculator = new ShippingCalculator();
    }

    public double calculateTax(double price, String state) {
        return taxCalculator.calculateTax(price, state);
    }

    public double calculateShipping(double weight, double distance) {
        return shippingCalculator.calculateShipping(weight, distance);
    }
}

class QuoteGenerator {
    private TaxCalculator taxCalculator;
    private ShippingCalculator shippingCalculator;

    public QuoteGenerator() {
        this.taxCalculator = new TaxCalculator();
        this.shippingCalculator = new ShippingCalculator();
    }

    public double calculateTax(double price, String state) {
        return taxCalculator.calculateTax(price, state);
    }

    public double calculateShipping(double weight, double distance) {
        return shippingCalculator.calculateShipping(weight, distance);
    }
}
