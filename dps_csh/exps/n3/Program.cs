A b = new B();
b.WriteAbout();

// ***

public abstract class A 
{
    // абстрактные члены также, как и виртуальные,
    // являются частью полиморфного интерфейса.
    public abstract void WriteAbout();
}

public class B : A 
{
    public override void WriteAbout()
    {
        Console.WriteLine("B");
    }
}

