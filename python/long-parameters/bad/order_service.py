from datetime import datetime

class OrderService:
    def create_order(
        self,
        customer_id,
        customer_name,
        customer_email,
        customer_phone,
        customer_address,
        customer_city,
        customer_state,
        customer_zip_code,
        product_id,
        product_name,
        product_price,
        quantity,
        tax_rate,
        discount_percent,
        shipping_method,
        shipping_cost,
        payment_method,
        billing_address,
        billing_city,
        billing_state,
        billing_zip_code,
        notes
    ):
        # Calculate totals
        subtotal = product_price * quantity
        discount_amount = subtotal * (discount_percent / 100)
        taxable_amount = subtotal - discount_amount
        tax_amount = taxable_amount * (tax_rate / 100)
        total = taxable_amount + tax_amount + shipping_cost

        # Create order record
        order_data = {
            'customer_id': customer_id,
            'customer_name': customer_name,
            'customer_email': customer_email,
            'customer_phone': customer_phone,
            'customer_address': customer_address,
            'customer_city': customer_city,
            'customer_state': customer_state,
            'customer_zip': customer_zip_code,
            'product_id': product_id,
            'product_name': product_name,
            'product_price': product_price,
            'quantity': quantity,
            'subtotal': subtotal,
            'discount_percent': discount_percent,
            'discount_amount': discount_amount,
            'tax_rate': tax_rate,
            'tax_amount': tax_amount,
            'shipping_method': shipping_method,
            'shipping_cost': shipping_cost,
            'payment_method': payment_method,
            'billing_address': billing_address,
            'billing_city': billing_city,
            'billing_state': billing_state,
            'billing_zip': billing_zip_code,
            'total': total,
            'notes': notes,
            'created_at': datetime.now().isoformat()
        }

        # In a real application, this would save to database
        return order_data

    def update_order(
        self,
        order_id,
        customer_id,
        customer_name,
        customer_email,
        customer_phone,
        customer_address,
        customer_city,
        customer_state,
        customer_zip_code,
        product_id,
        product_name,
        product_price,
        quantity,
        tax_rate,
        discount_percent,
        shipping_method,
        shipping_cost,
        payment_method,
        billing_address,
        billing_city,
        billing_state,
        billing_zip_code,
        notes
    ):
        # Similar logic but for updating
        subtotal = product_price * quantity
        discount_amount = subtotal * (discount_percent / 100)
        taxable_amount = subtotal - discount_amount
        tax_amount = taxable_amount * (tax_rate / 100)
        total = taxable_amount + tax_amount + shipping_cost

        order_data = {
            'id': order_id,
            'customer_id': customer_id,
            'customer_name': customer_name,
            'customer_email': customer_email,
            'customer_phone': customer_phone,
            'customer_address': customer_address,
            'customer_city': customer_city,
            'customer_state': customer_state,
            'customer_zip': customer_zip_code,
            'product_id': product_id,
            'product_name': product_name,
            'product_price': product_price,
            'quantity': quantity,
            'subtotal': subtotal,
            'discount_percent': discount_percent,
            'discount_amount': discount_amount,
            'tax_rate': tax_rate,
            'tax_amount': tax_amount,
            'shipping_method': shipping_method,
            'shipping_cost': shipping_cost,
            'payment_method': payment_method,
            'billing_address': billing_address,
            'billing_city': billing_city,
            'billing_state': billing_state,
            'billing_zip': billing_zip_code,
            'total': total,
            'notes': notes,
            'updated_at': datetime.now().isoformat()
        }

        return order_data
