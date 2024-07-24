import abc
from typing import List

'''

Абстрактные классы в объектно-ориентированном программировании — это базовые классы,
    которые можно наследовать, но нельзя реализовывать. 
То есть на их основе нельзя создать объект.

'''

# ------------------------------------------------------------------------

class Component(abc.ABC):
    @abc.abstractmethod
    def execute(self) -> int: pass


class Leaf(Component):
    def __init__(self) -> None:
        super().__init__()

    def execute(self) -> int:
        # print("Leaf.execute")
        return 1 


class Composite(Component):
    def __init__(self) -> None:
        self.children: List[Component] = list()
        super().__init__()

    # !!!
    def execute(self) -> int:
        # print("Composite.execute")  
        result: int = 0
        for c in self.children:
            result += c.execute()   
        return result

    def add(self, c: Component):
        self.children.append(c)

    def remove(self, c: Component):
        self.children.remove(c)

    def get_children(self) -> List[Component]:
        return self.children
    
# ------------------------------------------------------------------------

cs1 = Composite()
cs1.add(Leaf())
cs1.add(Leaf())
cs1.add(Leaf())

# ------------------------------------------------------------------------

cs = Composite()
print(cs.execute())

cs.add(Leaf())
print(cs.execute())

cs.add(Leaf())
print(cs.execute())

cs.add(cs1)
print(cs.execute())

# ------------------------------------------------------------------------