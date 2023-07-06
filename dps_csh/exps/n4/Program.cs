[assembly: CLSCompliant(true)]
namespace SomeLibrary
{
    public class Program
    {
        static public void Main(string[] args)
        {
            Console.WriteLine("Main");
        }
    }

    // возникает предупреждение поскольку тип открытый
    public sealed class SomeLibraryType
    {
        // тип, возвращаемый функцией не соответсвует CLS
        public UInt32 Abc() { return 0; }

        // идентификатор abc() отличается от предыдущего, только если
        // не выдерживается соответсвие
        public void abc() { }

        // ошибки нет, метод закрытый
        private UInt32 ABC() { return 0; }
    }
}