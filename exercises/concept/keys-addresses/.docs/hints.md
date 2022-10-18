# Hints

## 1. Compress public key

Compress your public key depending on whether y is even or odd

```python
(public_key_x, public_key_y) = public_key
compressed_prefix = '02' if (public_key_y % 2) == 0 else '03'
```

## 2. Generate address from compressed public key

Just call the cryptos.pubkey_to_address function

```python
cryptos.pubkey_to_address(public_key))
```