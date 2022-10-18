import unittest
import pytest

class KeytoAddressTest(unittest.TestCase):
    def test_key_to_address(self):
        compressed_addr = '1ME6yf9teG2wDJLigdFWGFeXYsk1nGZk45'
        compressed_public_addr = '02a7802cf1eb0e09a474d0da6c997ffde828c562e5b488f4e09bc57c5217bbe6c3'

        error = ("Expected same output")

        self.assertEqual(compressed_addr, cryptos.pubkey_to_address(compressed_public_addr), msg=error)