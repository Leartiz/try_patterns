import abc

class Component(abc.ABC):
    @abc.abstractmethod
    def execute():
        pass