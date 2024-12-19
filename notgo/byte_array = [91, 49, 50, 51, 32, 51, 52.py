byte_array = [91, 49, 50, 51, 32, 51, 52, 32, 49, 49, 52, 32, 49, 48, 49, 32, 49, 49, 53, 32, 49, 49, 55, 32, 49, 48, 56, 32, 49, 49, 54, 32, 51, 52, 32, 53, 56, 32, 51, 52, 32, 53, 49, 32, 51, 52, 32, 49, 50, 53, 93]

# Convert each byte to a character using chr()
string_result = ''.join(chr(byte) for byte in byte_array)

print(string_result)