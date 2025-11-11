# 3. Embedding a temporary variable (Inline Temp)
#
# BEFORE: Unnecessary temporary variable
class PriceCalculatorBefore:
    def __init__(self, quantity, item_price):
        self.quantity = quantity
        self.item_price = item_price

    def get_price(self):
        base_price = self.quantity * self.item_price
        if base_price > 1000:
            return base_price * 0.95
        else:
            return base_price * 0.98

# AFTER: Inline the temporary variable
class PriceCalculatorAfter:
    def __init__(self, quantity, item_price):
        self.quantity = quantity
        self.item_price = item_price

    def get_price(self):
        if self.quantity * self.item_price > 1000:
            return self.quantity * self.item_price * 0.95
        else:
            return self.quantity * self.item_price * 0.98

# 4. Replacing a temporary variable with a method call (Replace Temp with Query)
#
# BEFORE: Temporary variable used multiple times
class OrderBefore:
    def __init__(self, quantity, item_price):
        self._quantity = quantity
        self._item_price = item_price

    def get_price(self):
        base_price = self._quantity * self._item_price
        return base_price - self.get_discount(base_price)

    def get_discount(self, base_price):
        return max(0, base_price - 500) * 0.05

# AFTER: Replace temp with query
class OrderAfter:
    def __init__(self, quantity, item_price):
        self._quantity = quantity
        self._item_price = item_price

    def get_price(self):
        return self.get_base_price() - self.get_discount()

    def get_base_price(self):
        return self._quantity * self._item_price

    def get_discount(self):
        return max(0, self.get_base_price() - 500) * 0.05

# 5. Introduction of an explanatory variable (Introduce Explaining Variable)
#
# BEFORE: Complex expression hard to understand
class PerformanceCalculatorBefore:
    def __init__(self, goals, assists, minutes_played):
        self.goals = goals
        self.assists = assists
        self.minutes_played = minutes_played

    def get_performance(self):
        return (self.goals * 2) + (self.assists * 1.5) + (self.minutes_played / 60) * 0.1

# AFTER: Introduce explaining variables for clarity
class PerformanceCalculatorAfter:
    def __init__(self, goals, assists, minutes_played):
        self.goals = goals
        self.assists = assists
        self.minutes_played = minutes_played

    def get_performance(self):
        goal_points = self.goals * 2
        assist_points = self.assists * 1.5
        playing_time_bonus = (self.minutes_played / 60) * 0.1

        return goal_points + assist_points + playing_time_bonus

# 6. Splitting a Temporary Variable
#
# BEFORE: Same variable used for different purposes
class TemperatureMonitorBefore:
    def __init__(self):
        self._current_temp = 20.0
        self._adjustment = 2.0

    def get_current_temperature(self):
        return self._current_temp

    def get_adjustment(self):
        return self._adjustment

    def get_reading(self):
        temp = self.get_current_temperature()

        # First use: get initial reading
        initial_temp = temp

        # Later: temp is reused for different calculation
        temp = temp + self.get_adjustment()
        adjusted_temp = temp

        return {'initial': initial_temp, 'adjusted': adjusted_temp}

# AFTER: Split the temporary variable
class TemperatureMonitorAfter:
    def __init__(self):
        self._current_temp = 20.0
        self._adjustment = 2.0

    def get_current_temperature(self):
        return self._current_temp

    def get_adjustment(self):
        return self._adjustment

    def get_reading(self):
        temp = self.get_current_temperature()
        initial_temp = temp

        adjusted_temp = temp + self.get_adjustment()

        return {'initial': initial_temp, 'adjusted': adjusted_temp}

# 7. Removing parameter Assignments (Remove Assignments to Parameters)
#
# BEFORE: Parameter is modified inside method
class DiscountCalculatorBefore:
    def apply_discount(self, price):
        if price > 100:
            price = price * 0.9  # Modifying parameter
        return price

# AFTER: Use a local variable instead
class DiscountCalculatorAfter:
    def apply_discount(self, price):
        result = price
        if price > 100:
            result = price * 0.9
        return result

# 8. Replacing a method with a method Object (Replace Method with Method Object)
#
# BEFORE: Method with many parameters and local variables
class AccountBefore:
    def calculate_interest(self, principal, rate, time, compounding_frequency):
        amount = principal * (1 + (rate / compounding_frequency)) ** (compounding_frequency * time)
        interest = amount - principal
        return interest

# AFTER: Extract to a method object
class InterestCalculation:
    def __init__(self, principal, rate, time, compounding_frequency):
        self._principal = principal
        self._rate = rate
        self._time = time
        self._compounding_frequency = compounding_frequency

    def calculate(self):
        amount = self._principal * (1 + (self._rate / self._compounding_frequency)) ** \
                 (self._compounding_frequency * self._time)
        return amount - self._principal

class AccountAfter:
    def calculate_interest(self, principal, rate, time, compounding_frequency):
        calculation = InterestCalculation(principal, rate, time, compounding_frequency)
        return calculation.calculate()
