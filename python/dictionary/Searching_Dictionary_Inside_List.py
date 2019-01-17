"""
:: BAHASA
Dibawah ini terdapat kumpulan dictionary yang ada didalam list, akan tetapi tiap dictionary
memiliki key yang sama tetapi value-nya berbeda. Kasusnya adalah bagaimana caranya untuk mendapatkan
baris dictionary yang tepat hanya dengan menggunakan satu kata kunci dari key dan value yang berada
didalam dictionary.
=====================================================================================================
:: ENGLISH
There's bunch of dictionary inside list, however each dictionary has same key but different value.
In this case how to find the correct dictionary with just single keyword based on index and value
inside the dictionary.

:: HOW TO USE
$ python Searching_Dictionary_Inside_List.py 'your_key' 'your_value'
for example:
$ python Searching_Dictionary_Inside_List.py name 'John Doe'
"""

import sys

user = [
    {'name': 'John Doe', 'city': 'California', 'country': 'US'},
    {'name': 'Harry Brian', 'city': 'London', 'country': 'UK'},
    {'name': 'Jared Argus', 'city': 'Miami', 'country': 'US'},
    {'name': 'Brian Dom', 'city': 'Mexico City', 'country': 'Mexico'},
    {'name': 'Harry Brian', 'city': 'Jakarta', 'country': 'Indonesia'},
    {'name': 'Agus Suparjo', 'city': 'Malang', 'country': 'Indonesia'},
    {'name': 'Jane Doe', 'city': 'Kansas', 'country': 'US'},
    {'name': 'Jerome Dominique', 'city': 'California', 'country': 'US'},
    {'name': 'Tyrone', 'city': 'Detroit', 'country': 'US'},
    {'name': 'James Bond', 'city': 'London', 'country': 'UK'}
]

key = sys.argv[1]
value = sys.argv[2]

print(filter(lambda dct: dct[key] == value, user))
