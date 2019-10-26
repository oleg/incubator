from typing import Tuple
from secrets import token_bytes

def encrypt(original: str) -> Tuple[int, int]:
    original_bytes: bytes = original.encode()
    original_key: int = int.from_bytes(original_bytes, "big")
    dummy: int = random_key(len(original_bytes))
    encrypted: int = original_key ^ dummy
    return dummy, encrypted

def random_key(length: int) -> int:
    tb: bytes = token_bytes(length)
    return int.from_bytes(tb, "big")

def decrypt(key1: int, key2: int) -> str:
    decrypted: int = key1 ^ key2
    temp: bytes = decrypted.to_bytes((decrypted.bit_length() + 7) // 8, "big")
    return temp.decode()

# Tests
def test_encode_decode_string():
    assert "hello".encode().decode() == "hello"

def test_encode_decode_key():
    phrase = "Hello, World!"
    key1, key2 = encrypt(phrase)
    result = decrypt(key1, key2)
    assert result == phrase
