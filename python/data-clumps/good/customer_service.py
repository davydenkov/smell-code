class CustomerService:
    def create_customer(self, person, address):
        if not person.is_valid():
            raise Exception('Invalid person data')

        if not address.is_valid():
            raise Exception('Invalid address data')

        # Create customer record
        customer_data = {
            'first_name': person.get_first_name(),
            'last_name': person.get_last_name(),
            'email': person.get_email(),
            'phone': person.get_phone(),
            'date_of_birth': person.get_date_of_birth(),
            'street': address.get_street(),
            'city': address.get_city(),
            'state': address.get_state(),
            'zip_code': address.get_zip_code()
        }

        # Save to database (simulated)
        return customer_data

    def update_customer_address(self, customer_id, address):
        if not address.is_valid():
            raise Exception('Invalid address data')

        # Update address
        address_data = {
            'customer_id': customer_id,
            'street': address.get_street(),
            'city': address.get_city(),
            'state': address.get_state(),
            'zip_code': address.get_zip_code()
        }

        # Save to database (simulated)
        return address_data

    def send_welcome_email(self, person, address):
        message = f"Welcome {person.get_full_name()}!\n\n"
        message += f"Your address: {address.to_string()}\n"

        # Send email (simulated)
        return {'to': person.get_email(), 'message': message}

    def validate_shipping_address(self, address):
        return address.is_valid()

    def format_address_label(self, person, address):
        return f"{person.get_full_name()}\n{address.to_label_format()}"
