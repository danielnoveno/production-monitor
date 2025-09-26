# python_scripts/dummy_scan.py
import random
import string

def generate_barcode():
    return ''.join(random.choices(string.ascii_uppercase + string.digits, k=8))

if __name__ == "__main__":
    print(generate_barcode())
