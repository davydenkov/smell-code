public class OrderProcessor
{
    private TaxCalculator taxCalculator;
    private ShippingCalculator shippingCalculator;

    public OrderProcessor()
    {
        taxCalculator = new TaxCalculator();
        shippingCalculator = new ShippingCalculator();
    }

    public double CalculateTax(double price, string state)
    {
        return taxCalculator.CalculateTax(price, state);
    }

    public double CalculateShipping(double weight, double distance)
    {
        return shippingCalculator.CalculateShipping(weight, distance);
    }
}

public class InvoiceGenerator
{
    private TaxCalculator taxCalculator;
    private ShippingCalculator shippingCalculator;

    public InvoiceGenerator()
    {
        taxCalculator = new TaxCalculator();
        shippingCalculator = new ShippingCalculator();
    }

    public double CalculateTax(double price, string state)
    {
        return taxCalculator.CalculateTax(price, state);
    }

    public double CalculateShipping(double weight, double distance)
    {
        return shippingCalculator.CalculateShipping(weight, distance);
    }
}

public class QuoteGenerator
{
    private TaxCalculator taxCalculator;
    private ShippingCalculator shippingCalculator;

    public QuoteGenerator()
    {
        taxCalculator = new TaxCalculator();
        shippingCalculator = new ShippingCalculator();
    }

    public double CalculateTax(double price, string state)
    {
        return taxCalculator.CalculateTax(price, state);
    }

    public double CalculateShipping(double weight, double distance)
    {
        return shippingCalculator.CalculateShipping(weight, distance);
    }
}
