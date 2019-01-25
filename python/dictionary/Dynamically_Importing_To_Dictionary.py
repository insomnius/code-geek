"""
:: BAHASA
Disini terdapat kasus dimana anda harus melakukan import module secara dinamis dan lalu mendapatkan
property yang ada didalam, setelah itu anda harus merubah import tersebut menjadi sebuah dictionary
yang dimana didalamnya terdapat sebuah list yang berisikan variable yang telah diolah menjadi key
dan value.
=====================================================================================================
:: ENGLISH
In this case you will be doing import module dynamically and get the property inside it, the you must
change the import into a dictionary where inside the dictionary there's a list contain variable that
already changed into key and value

:: HOW TO USE
$ python Dynamically_Importing_To_Dictionary.py 'your_python_module' 'your_variable', 'your_variable' ...
for example:
$ python Dynamically_Importing_To_Dictionary.py movies iron_man superman
"""

import sys, imp


my_find = sys.argv[1]
my_dict = {my_find: []}
my_import = 'sample/%s.py' % my_find
my_data = sys.argv[2:]

for i in my_data:
    src = imp.load_source(i, my_import)
    my_dict[my_find].append({i : getattr(src, i)})

print(my_dict)
