Test();
GC.Collect();   // очистка памяти под объект tom
Console.Read(); // ставим задержку

Console.WriteLine(GC.GetTotalMemory(false));

void Test()
{
    Person tom = new Person("Tom");

    Console.WriteLine("gen: " + GC.GetGeneration(tom));
}

public class Person
{
    public string Name { get; }
    public Person(string name) => Name = name;

    ~Person()
    {
        Console.WriteLine($"{Name} has been deleted");
    }
}