# Instructions

In this exercise you are going to write some more code to compress your public key and generate compressed bitcoin addres from compressed public key

You have two tasks.
The first one is related to compress the public key and the second to generate the compressed address

## 1. Compress public key

Compress your public key depeding on whter y is even or odd

```python
(public_key_x, public_key_y) = public_key
compressed_prefix = '02' if (public_key_y % 2) == 0 else '03'
```

## 2. Generate address from compressed public key

```python
cryptos.pubkey_to_address(hex_compressed_public_key))
```