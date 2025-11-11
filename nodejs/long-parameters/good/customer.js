class Address {
    constructor(street = '', city = '', state = '', zipCode = '') {
        this.street = street;
        this.city = city;
        this.state = state;
        this.zipCode = zipCode;
    }
}

class Customer {
    constructor(id, name, email, phone = null, shippingAddress = null, billingAddress = null) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.phone = phone;
        this.shippingAddress = shippingAddress || new Address();
        this.billingAddress = billingAddress || new Address();
    }
}

class Product {
    constructor(id, name, price) {
        this.id = id;
        this.name = name;
        this.price = price;
    }
}

class OrderDetails {
    constructor(
        product,
        quantity,
        taxRate = 0.0,
        discountPercent = 0.0,
        shippingMethod = 'standard',
        shippingCost = 0.0,
        paymentMethod = 'credit_card',
        notes = null
    ) {
        this.product = product;
        this.quantity = quantity;
        this.taxRate = taxRate;
        this.discountPercent = discountPercent;
        this.shippingMethod = shippingMethod;
        this.shippingCost = shippingCost;
        this.paymentMethod = paymentMethod;
        this.notes = notes;
    }
}

module.exports = {
    Address,
    Customer,
    Product,
    OrderDetails
};
