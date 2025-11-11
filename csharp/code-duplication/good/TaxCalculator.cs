using System.Collections.Generic;

public class TaxCalculator
{
    private static readonly Dictionary<string, double> TaxRates = new Dictionary<string, double>
    {
        ["CA"] = 0.0825,
        ["NY"] = 0.04,
        ["TX"] = 0.0625
    };

    public double CalculateTax(double price, string state)
    {
        double taxRate = TaxRates.ContainsKey(state) ? TaxRates[state] : 0.0;
        return price * taxRate;
    }
}

public class ShippingCalculator
{
    private const double BaseRate = 5.0;
    private const double WeightRatePerUnit = 0.5;
    private const double DistanceRatePerUnit = 0.1;

    public double CalculateShipping(double weight, double distance)
    {
        double weightRate = weight * WeightRatePerUnit;
        double distanceRate = distance * DistanceRatePerUnit;

        return BaseRate + weightRate + distanceRate;
    }
}
