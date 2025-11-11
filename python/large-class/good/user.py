class User:
    def __init__(self, id: int, email: str, name: str, balance: float = 0.0):
        self._id = id
        self._email = email
        self._name = name
        self._balance = balance

    def get_id(self) -> int:
        return self._id

    def get_email(self) -> str:
        return self._email

    def get_name(self) -> str:
        return self._name

    def get_balance(self) -> float:
        return self._balance

    def update_profile(self, name: str, email: str) -> None:
        self._name = name
        self._email = email
