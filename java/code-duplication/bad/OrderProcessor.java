class OrderProcessor {
    public double calculateTax(double price, String state) {
        double taxRate = 0.0;

        if ("CA".equals(state)) {
            taxRate = 0.0825;
        } else if ("NY".equals(state)) {
            taxRate = 0.04;
        } else if ("TX".equals(state)) {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    public double calculateShipping(double weight, double distance) {
        double baseRate = 5.0;
        double weightRate = weight * 0.5;
        double distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

class InvoiceGenerator {
    public double calculateTax(double price, String state) {
        double taxRate = 0.0;

        if ("CA".equals(state)) {
            taxRate = 0.0825;
        } else if ("NY".equals(state)) {
            taxRate = 0.04;
        } else if ("TX".equals(state)) {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    public double calculateShipping(double weight, double distance) {
        double baseRate = 5.0;
        double weightRate = weight * 0.5;
        double distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}

class QuoteGenerator {
    public double calculateTax(double price, String state) {
        double taxRate = 0.0;

        if ("CA".equals(state)) {
            taxRate = 0.0825;
        } else if ("NY".equals(state)) {
            taxRate = 0.04;
        } else if ("TX".equals(state)) {
            taxRate = 0.0625;
        }

        return price * taxRate;
    }

    public double calculateShipping(double weight, double distance) {
        double baseRate = 5.0;
        double weightRate = weight * 0.5;
        double distanceRate = distance * 0.1;

        return baseRate + weightRate + distanceRate;
    }
}
