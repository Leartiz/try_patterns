'''
Декораторы в Python представляют функцию,
которая в качестве параметра получает функцию и 
    в качестве результата также возвращает функцию.
'''

def print_stars(num: int):
    print("*" * num)

def select(input_func):    
    def output_func(*args): 
        name: str = args[0] 
        stars_count = 7+len(name)  

        print_stars(stars_count) 
        input_func(*args)                
        print_stars(stars_count) 

    return output_func

@select         
def hello(name: str):
    print(f"hello {name}!")

hello("!!")
print()

# ------------------------------------------------------------------------

def fun(**kwargs):
    print(kwargs)
    
    for key in kwargs:
        print(f"{key} = {kwargs[key]}")
  
fun(name="Tom", age="38", company="Google")

# ------------------------------------------------------------------------

'''
Оператор * позволяет передать в функцию несколько значений, 
    и все они будут упакованы в кортеж.
Оператор ** упаковывает аргументы, переданные по имени, в словарь.
'''

def print_kw_and_args(input_func):
    def output_func(*args, **kwargs):
        if args:
            print("args:", sep=" ")
            print(args)
        else:
            print("args is none")

        if kwargs:
            print("kwargs:", sep=" ")
            print(kwargs)
        else:
            print("kwargs is none")

        input_func(*args, **kwargs) # Распаковка!
        
    return output_func

# ------------------------------------------------------------------------

@print_kw_and_args
def print_text(value: str):
    print(value)

print_text("Hello World!")

# ------------------------------------------------------------------------

def convert_to_upper(input_func):
    def output_func(*args, **kwargs):
        data: str = input_func(*args, **kwargs)
        return data.upper()
    return output_func

def remove_adj_spaces(input_func):
    def output_func(*args, **kwargs):
        data: str = input_func(*args, **kwargs)
        parts = data.split() 
        data = " ".join(parts) 
        return data
    return output_func

def remove_vowels(input_func) -> str:
    def output_func(*args, **kwargs):
        vowels = "aeiouAEIOU"
        data: str = input_func(*args, **kwargs)
        data: str = "".join(char for char in data \
                            if char not in vowels)
        return data
    return output_func

@remove_vowels
@remove_adj_spaces
@convert_to_upper
def remove_extra_chars(data: str) -> str:
    return data

print(remove_extra_chars("fa   fa   fa   fa"))

# ------------------------------------------------------------------------