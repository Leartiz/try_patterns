# Composite

**Компоновщик** — это структурный паттерн проектирования, который позволяет сгруппировать множество объектов в древовидную структуру, а затем работать с ней так, как будто это единичный объект.

**Компоновщик** и **Декоратор** имеют похожие структуры классов из-за того, что оба построены на рекурсивной вложенности. 

<!-- --------------------------------------------------------------------- -->

```mermaid
classDiagram

%% nodes

    class Client{
    }

    class Component{
        +int execute()
    }

    class Leaf{
        +int execute()
    }

    class Composite{
        -children List~Component~
        +int execute()
        +add(c: Component)
        +remove(c: Component)
        +List~Component~ get_children()
    }

%% edges

    Composite o-- Component
    Composite ..|> Component
    Leaf ..|> Component
    Component <-- Client 
```