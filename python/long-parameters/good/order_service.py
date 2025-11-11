from datetime import datetime

class OrderService:
    def __init__(self):
        self.calculator = OrderCalculator()

    def create_order(self, customer, order_details):
        totals = self.calculator.calculate_totals(order_details)

        order_data = {
            'customer_id': customer.id,
            'customer_name': customer.name,
            'customer_email': customer.email,
            'customer_phone': customer.phone,
            'customer_address': customer.shipping_address.street,
            'customer_city': customer.shipping_address.city,
            'customer_state': customer.shipping_address.state,
            'customer_zip': customer.shipping_address.zip_code,
            'product_id': order_details.product.id,
            'product_name': order_details.product.name,
            'product_price': order_details.product.price,
            'quantity': order_details.quantity,
            'subtotal': totals['subtotal'],
            'discount_percent': order_details.discount_percent,
            'discount_amount': totals['discount_amount'],
            'tax_rate': order_details.tax_rate,
            'tax_amount': totals['tax_amount'],
            'shipping_method': order_details.shipping_method,
            'shipping_cost': totals['shipping_cost'],
            'payment_method': order_details.payment_method,
            'billing_address': customer.billing_address.street,
            'billing_city': customer.billing_address.city,
            'billing_state': customer.billing_address.state,
            'billing_zip': customer.billing_address.zip_code,
            'total': totals['total'],
            'notes': order_details.notes,
            'created_at': datetime.now().isoformat()
        }

        # In a real application, this would save to database
        return order_data

    def update_order(self, order_id, customer, order_details):
        totals = self.calculator.calculate_totals(order_details)

        order_data = {
            'id': order_id,
            'customer_id': customer.id,
            'customer_name': customer.name,
            'customer_email': customer.email,
            'customer_phone': customer.phone,
            'customer_address': customer.shipping_address.street,
            'customer_city': customer.shipping_address.city,
            'customer_state': customer.shipping_address.state,
            'customer_zip': customer.shipping_address.zip_code,
            'product_id': order_details.product.id,
            'product_name': order_details.product.name,
            'product_price': order_details.product.price,
            'quantity': order_details.quantity,
            'subtotal': totals['subtotal'],
            'discount_percent': order_details.discount_percent,
            'discount_amount': totals['discount_amount'],
            'tax_rate': order_details.tax_rate,
            'tax_amount': totals['tax_amount'],
            'shipping_method': order_details.shipping_method,
            'shipping_cost': totals['shipping_cost'],
            'payment_method': order_details.payment_method,
            'billing_address': customer.billing_address.street,
            'billing_city': customer.billing_address.city,
            'billing_state': customer.billing_address.state,
            'billing_zip': customer.billing_address.zip_code,
            'total': totals['total'],
            'notes': order_details.notes,
            'updated_at': datetime.now().isoformat()
        }

        return order_data
