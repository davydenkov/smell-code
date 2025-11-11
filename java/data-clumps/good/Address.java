public class Address {
    private String street;
    private String city;
    private String state;
    private String zipCode;

    public Address(String street, String city, String state, String zipCode) {
        this.street = street;
        this.city = city;
        this.state = state;
        this.zipCode = zipCode;
    }

    public String getStreet() {
        return street;
    }

    public String getCity() {
        return city;
    }

    public String getState() {
        return state;
    }

    public String getZipCode() {
        return zipCode;
    }

    public boolean isValid() {
        if (street == null || street.isEmpty() ||
            city == null || city.isEmpty() ||
            state == null || state.isEmpty() ||
            zipCode == null || zipCode.isEmpty()) {
            return false;
        }

        if (zipCode.length() != 5) {
            return false;
        }

        return true;
    }

    public String toString() {
        return street + ", " + city + ", " + state + " " + zipCode;
    }

    public String toLabelFormat() {
        return street + "\n" + city + ", " + state + " " + zipCode;
    }
}
