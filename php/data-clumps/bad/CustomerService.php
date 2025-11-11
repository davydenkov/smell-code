<?php

class CustomerService {
    public function createCustomer(
        $firstName, $lastName, $email,
        $street, $city, $state, $zipCode,
        $phone, $dateOfBirth
    ) {
        // Validate data
        if (empty($firstName) || empty($lastName)) {
            throw new Exception('Name is required');
        }
        if (!filter_var($email, FILTER_VALIDATE_EMAIL)) {
            throw new Exception('Invalid email');
        }

        // Create customer record
        $customerData = [
            'first_name' => $firstName,
            'last_name' => $lastName,
            'email' => $email,
            'street' => $street,
            'city' => $city,
            'state' => $state,
            'zip_code' => $zipCode,
            'phone' => $phone,
            'date_of_birth' => $dateOfBirth
        ];

        // Save to database (simulated)
        return $customerData;
    }

    public function updateCustomerAddress(
        $customerId,
        $street, $city, $state, $zipCode
    ) {
        // Update address
        $addressData = [
            'street' => $street,
            'city' => $city,
            'state' => $state,
            'zip_code' => $zipCode
        ];

        // Save to database (simulated)
        return array_merge(['customer_id' => $customerId], $addressData);
    }

    public function sendWelcomeEmail(
        $firstName, $lastName, $email,
        $street, $city, $state, $zipCode
    ) {
        $fullName = $firstName . ' ' . $lastName;
        $fullAddress = $street . ', ' . $city . ', ' . $state . ' ' . $zipCode;

        $message = "Welcome {$fullName}!\n\n";
        $message .= "Your address: {$fullAddress}\n";

        // Send email (simulated)
        return ['to' => $email, 'message' => $message];
    }

    public function validateShippingAddress(
        $street, $city, $state, $zipCode
    ) {
        if (empty($street) || empty($city) || empty($state) || empty($zipCode)) {
            return false;
        }

        // Additional validation logic
        if (strlen($zipCode) !== 5) {
            return false;
        }

        return true;
    }

    public function formatAddressLabel(
        $firstName, $lastName,
        $street, $city, $state, $zipCode
    ) {
        $fullName = $firstName . ' ' . $lastName;
        $fullAddress = $street . "\n" . $city . ', ' . $state . ' ' . $zipCode;

        return $fullName . "\n" . $fullAddress;
    }
}
