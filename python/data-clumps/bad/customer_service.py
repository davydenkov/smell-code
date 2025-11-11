class CustomerService:
    def create_customer(self, first_name, last_name, email, street, city, state, zip_code, phone, date_of_birth):
        # Validate data
        if not first_name or not last_name:
            raise Exception('Name is required')
        if not self._is_valid_email(email):
            raise Exception('Invalid email')

        # Create customer record
        customer_data = {
            'first_name': first_name,
            'last_name': last_name,
            'email': email,
            'street': street,
            'city': city,
            'state': state,
            'zip_code': zip_code,
            'phone': phone,
            'date_of_birth': date_of_birth
        }

        # Save to database (simulated)
        return customer_data

    def update_customer_address(self, customer_id, street, city, state, zip_code):
        # Update address
        address_data = {
            'street': street,
            'city': city,
            'state': state,
            'zip_code': zip_code
        }

        # Save to database (simulated)
        return {'customer_id': customer_id, **address_data}

    def send_welcome_email(self, first_name, last_name, email, street, city, state, zip_code):
        full_name = f"{first_name} {last_name}"
        full_address = f"{street}, {city}, {state} {zip_code}"

        message = f"Welcome {full_name}!\n\n"
        message += f"Your address: {full_address}\n"

        # Send email (simulated)
        return {'to': email, 'message': message}

    def validate_shipping_address(self, street, city, state, zip_code):
        if not street or not city or not state or not zip_code:
            return False

        # Additional validation logic
        if len(zip_code) != 5:
            return False

        return True

    def format_address_label(self, first_name, last_name, street, city, state, zip_code):
        full_name = f"{first_name} {last_name}"
        full_address = f"{street}\n{city}, {state} {zip_code}"

        return f"{full_name}\n{full_address}"

    def _is_valid_email(self, email):
        import re
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None
