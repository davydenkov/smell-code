class TaxCalculator {
    private static final java.util.Map<String, Double> TAX_RATES = java.util.Map.of(
        "CA", 0.0825,
        "NY", 0.04,
        "TX", 0.0625
    );

    public double calculateTax(double price, String state) {
        double taxRate = TAX_RATES.getOrDefault(state, 0.0);
        return price * taxRate;
    }
}

class ShippingCalculator {
    private static final double BASE_RATE = 5.0;
    private static final double WEIGHT_RATE_PER_UNIT = 0.5;
    private static final double DISTANCE_RATE_PER_UNIT = 0.1;

    public double calculateShipping(double weight, double distance) {
        double weightRate = weight * WEIGHT_RATE_PER_UNIT;
        double distanceRate = distance * DISTANCE_RATE_PER_UNIT;

        return BASE_RATE + weightRate + distanceRate;
    }
}
