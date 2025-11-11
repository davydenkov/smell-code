import secrets

class PaymentService:
    def __init__(self, payment_config):
        self.stripe_secret_key = payment_config.get('stripe_secret')
        self.paypal_client_id = payment_config.get('paypal_client_id')
        self.paypal_secret = payment_config.get('paypal_secret')

    def process_stripe_payment(self, amount: float, token: str):
        if token and amount > 0:
            return {'success': True, 'transaction_id': f'stripe_{secrets.token_hex(8)}'}
        return {'success': False, 'error': 'Invalid payment data'}

    def process_paypal_payment(self, amount: float, paypal_token: str):
        if paypal_token and amount > 0:
            return {'success': True, 'transaction_id': f'paypal_{secrets.token_hex(8)}'}
        return {'success': False, 'error': 'Invalid payment data'}

    def refund_payment(self, transaction_id: str, amount: float):
        if transaction_id.startswith('stripe_'):
            return {'success': True, 'refund_id': f'refund_{secrets.token_hex(8)}'}
        elif transaction_id.startswith('paypal_'):
            return {'success': True, 'refund_id': f'refund_{secrets.token_hex(8)}'}
        return {'success': False, 'error': 'Unknown transaction type'}
