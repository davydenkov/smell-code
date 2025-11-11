import java.util.*;

public class CustomerService {
    public Map<String, Object> createCustomer(
            String firstName, String lastName, String email,
            String street, String city, String state, String zipCode,
            String phone, String dateOfBirth
    ) throws Exception {
        // Validate data
        if (firstName == null || firstName.isEmpty() || lastName == null || lastName.isEmpty()) {
            throw new Exception("Name is required");
        }
        if (!isValidEmail(email)) {
            throw new Exception("Invalid email");
        }

        // Create customer record
        Map<String, Object> customerData = new HashMap<>();
        customerData.put("first_name", firstName);
        customerData.put("last_name", lastName);
        customerData.put("email", email);
        customerData.put("street", street);
        customerData.put("city", city);
        customerData.put("state", state);
        customerData.put("zip_code", zipCode);
        customerData.put("phone", phone);
        customerData.put("date_of_birth", dateOfBirth);

        // Save to database (simulated)
        return customerData;
    }

    public Map<String, Object> updateCustomerAddress(
            int customerId,
            String street, String city, String state, String zipCode
    ) {
        // Update address
        Map<String, Object> addressData = new HashMap<>();
        addressData.put("street", street);
        addressData.put("city", city);
        addressData.put("state", state);
        addressData.put("zip_code", zipCode);

        // Save to database (simulated)
        Map<String, Object> result = new HashMap<>(addressData);
        result.put("customer_id", customerId);
        return result;
    }

    public Map<String, String> sendWelcomeEmail(
            String firstName, String lastName, String email,
            String street, String city, String state, String zipCode
    ) {
        String fullName = firstName + " " + lastName;
        String fullAddress = street + ", " + city + ", " + state + " " + zipCode;

        String message = "Welcome " + fullName + "!\n\n";
        message += "Your address: " + fullAddress + "\n";

        // Send email (simulated)
        Map<String, String> result = new HashMap<>();
        result.put("to", email);
        result.put("message", message);
        return result;
    }

    public boolean validateShippingAddress(
            String street, String city, String state, String zipCode
    ) {
        if (street == null || street.isEmpty() ||
            city == null || city.isEmpty() ||
            state == null || state.isEmpty() ||
            zipCode == null || zipCode.isEmpty()) {
            return false;
        }

        // Additional validation logic
        if (zipCode.length() != 5) {
            return false;
        }

        return true;
    }

    public String formatAddressLabel(
            String firstName, String lastName,
            String street, String city, String state, String zipCode
    ) {
        String fullName = firstName + " " + lastName;
        String fullAddress = street + "\n" + city + ", " + state + " " + zipCode;

        return fullName + "\n" + fullAddress;
    }

    private boolean isValidEmail(String email) {
        // Simple email validation
        return email != null && email.contains("@") && email.contains(".");
    }
}
