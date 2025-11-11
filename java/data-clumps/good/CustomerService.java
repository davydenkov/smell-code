import java.util.*;

public class CustomerService {
    public Map<String, Object> createCustomer(Person person, Address address) throws Exception {
        if (!person.isValid()) {
            throw new Exception("Invalid person data");
        }

        if (!address.isValid()) {
            throw new Exception("Invalid address data");
        }

        // Create customer record
        Map<String, Object> customerData = new HashMap<>();
        customerData.put("first_name", person.getFirstName());
        customerData.put("last_name", person.getLastName());
        customerData.put("email", person.getEmail());
        customerData.put("phone", person.getPhone());
        customerData.put("date_of_birth", person.getDateOfBirth());
        customerData.put("street", address.getStreet());
        customerData.put("city", address.getCity());
        customerData.put("state", address.getState());
        customerData.put("zip_code", address.getZipCode());

        // Save to database (simulated)
        return customerData;
    }

    public Map<String, Object> updateCustomerAddress(int customerId, Address address) throws Exception {
        if (!address.isValid()) {
            throw new Exception("Invalid address data");
        }

        // Update address
        Map<String, Object> addressData = new HashMap<>();
        addressData.put("customer_id", customerId);
        addressData.put("street", address.getStreet());
        addressData.put("city", address.getCity());
        addressData.put("state", address.getState());
        addressData.put("zip_code", address.getZipCode());

        // Save to database (simulated)
        return addressData;
    }

    public Map<String, String> sendWelcomeEmail(Person person, Address address) {
        String message = "Welcome " + person.getFullName() + "!\n\n";
        message += "Your address: " + address.toString() + "\n";

        // Send email (simulated)
        Map<String, String> result = new HashMap<>();
        result.put("to", person.getEmail());
        result.put("message", message);
        return result;
    }

    public boolean validateShippingAddress(Address address) {
        return address.isValid();
    }

    public String formatAddressLabel(Person person, Address address) {
        return person.getFullName() + "\n" + address.toLabelFormat();
    }
}
