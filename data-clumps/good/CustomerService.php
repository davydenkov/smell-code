<?php

class CustomerService {
    public function createCustomer(Person $person, Address $address): array {
        if (!$person->isValid()) {
            throw new Exception('Invalid person data');
        }

        if (!$address->isValid()) {
            throw new Exception('Invalid address data');
        }

        // Create customer record
        $customerData = [
            'first_name' => $person->getFirstName(),
            'last_name' => $person->getLastName(),
            'email' => $person->getEmail(),
            'phone' => $person->getPhone(),
            'date_of_birth' => $person->getDateOfBirth(),
            'street' => $address->getStreet(),
            'city' => $address->getCity(),
            'state' => $address->getState(),
            'zip_code' => $address->getZipCode()
        ];

        // Save to database (simulated)
        return $customerData;
    }

    public function updateCustomerAddress(int $customerId, Address $address): array {
        if (!$address->isValid()) {
            throw new Exception('Invalid address data');
        }

        // Update address
        $addressData = [
            'customer_id' => $customerId,
            'street' => $address->getStreet(),
            'city' => $address->getCity(),
            'state' => $address->getState(),
            'zip_code' => $address->getZipCode()
        ];

        // Save to database (simulated)
        return $addressData;
    }

    public function sendWelcomeEmail(Person $person, Address $address): array {
        $message = "Welcome {$person->getFullName()}!\n\n";
        $message .= "Your address: {$address->toString()}\n";

        // Send email (simulated)
        return ['to' => $person->getEmail(), 'message' => $message];
    }

    public function validateShippingAddress(Address $address): bool {
        return $address->isValid();
    }

    public function formatAddressLabel(Person $person, Address $address): string {
        return $person->getFullName() . "\n" . $address->toLabelFormat();
    }
}
