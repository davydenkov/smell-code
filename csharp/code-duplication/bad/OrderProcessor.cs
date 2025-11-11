public class OrderProcessor
{
    public double CalculateTax(double price, string state)
    {
        double taxRate = 0.0;

        if (state == "CA")
        {
            taxRate = 0.0825;
        }
        else if (state == "NY")
        {
            taxRate = 0.04;
        }
        else if (state == "TX")
        {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    public double CalculateShipping(double weight, double distance)
    {
        double baseRate = 5.0;
        double weightRate = weight * 0.5;
        double distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

public class InvoiceGenerator
{
    public double CalculateTax(double price, string state)
    {
        double taxRate = 0.0;

        if (state == "CA")
        {
            taxRate = 0.0825;
        }
        else if (state == "NY")
        {
            taxRate = 0.04;
        }
        else if (state == "TX")
        {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    public double CalculateShipping(double weight, double distance)
    {
        double baseRate = 5.0;
        double weightRate = weight * 0.5;
        double distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

public class QuoteGenerator
{
    public double CalculateTax(double price, string state)
    {
        double taxRate = 0.0;

        if (state == "CA")
        {
            taxRate = 0.0825;
        }
        else if (state == "NY")
        {
            taxRate = 0.04;
        }
        else if (state == "TX")
        {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    public double CalculateShipping(double weight, double distance)
    {
        double baseRate = 5.0;
        double weightRate = weight * 0.5;
        double distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}
