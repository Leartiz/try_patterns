import os, abc

# ------------------------------------------------------------------------

current_directory = os.getcwd()
print(f"current directory: {current_directory}")

current_file_path = os.path.abspath(__file__)
print("Current file path:", current_file_path)

print(f"Os sep: {os.sep}")
fp_parts = current_file_path.split(os.sep)
current_file_path = (os.sep).join(fp_parts[:-1])
print("Current file path:", current_file_path)

# ------------------------------------------------------------------------

class DataSourceError(Exception):
    def __init__(self, message):
        super().__init__(message)

# ------------------------------------------------------------------------

class DataSource(abc.ABC):
    @abc.abstractmethod
    def read_data() -> str: pass

    @abc.abstractmethod
    def write_data(data): pass

'''
Composition and
    inheritance.
'''
class DataSourceDecorator(DataSource):
    def __init__(self, ds: DataSource) -> None:
        self.ds = ds

# ------------------------------------------------------------------------

class FileDataSource(DataSource):
    def __init__(self) -> None:
        try:
            fn = (os.sep).join([current_file_path, "data.txt"])
            self.f = open(fn, "w+")
        except OSError as e:
            raise DataSourceError(f"Open {fn} failed. Err: {str(e)}")
            
    def read_data(self) -> str:
        self.f.seek(0)
        return str(self.f.read())
    
    def write_data(self, data: str):
        self.f.write(data)
    
# ------------------------------------------------------------------------

class ConvertToUpperDataSourceDecorator(DataSourceDecorator):
    def __init__(self, ds: DataSource) -> None:
        super().__init__(ds)

    def read_data(self) -> str:
        data = self.ds.read_data()
        return data.upper()

    def write_data(self, data: str):
        self.ds.write_data(data.upper())


class ConvertToLowerDataSourceDecorator(DataSourceDecorator):
    def __init__(self, ds: DataSource) -> None:
        super().__init__(ds)

    def read_data(self) -> str:
        data = self.ds.read_data()
        return data.lower()

    def write_data(self, data: str):
        self.ds.write_data(data.lower())

# ------------------------------------------------------------------------

class VowelsRemoveDataSourceDecorator(DataSourceDecorator):
    vowels = "aeiouAEIOU"

    @staticmethod
    def remove_vowels(data: str) -> str:
        data: str = ''.join(char for char in data \
                       if char not in __class__.vowels)
        return data

    def __init__(self, ds: DataSource) -> None:
        super().__init__(ds)

    def read_data(self) -> str:
        data = self.ds.read_data()        
        data = self.remove_vowels(data)
        return data

    def write_data(self, data: str):
        data = self.remove_vowels(data)
        self.ds.write_data(data)


class AdjacentSpacesRemoveDataSourceDecorator(DataSourceDecorator):
    @staticmethod
    def remove_adj_spaces(data: str) -> str:
        words = data.split() 
        result = " ".join(words) 
        return result

    def __init__(self, ds: DataSource) -> None:
        super().__init__(ds)

    def read_data(self) -> str:
        data = self.ds.read_data()        
        data = self.remove_adj_spaces(data)
        return data

    def write_data(self, data: str):
        data = self.remove_adj_spaces(data)
        self.ds.write_data(data)

# ------------------------------------------------------------------------

assert issubclass(FileDataSource, DataSource)
assert isinstance(FileDataSource(), DataSource)

# ------------------------------------------------------------------------

if __name__ == '__main__':
    data_source = FileDataSource() # simple!
    data_source = ConvertToUpperDataSourceDecorator(data_source)
    data_source = VowelsRemoveDataSourceDecorator(data_source)
    data_source = ConvertToLowerDataSourceDecorator(data_source)
    data_source = AdjacentSpacesRemoveDataSourceDecorator(data_source)

    # ***

    print("read data: " + data_source.read_data()) # empty!
    data_source.write_data("What can AI do?")
    print("read data: " + data_source.read_data())

