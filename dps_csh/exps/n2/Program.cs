using System.Text;

string s = "Работать надо и все будет?";
char c = 'а';

int res = s.CharCount(c);
Console.WriteLine(res);

B b = new B();
Console.WriteLine(b.StrMult(5));

// ***

NS_1.A a = new NS_1.A("A From NS_1");

NS_1.SUB_1.AEx.WriteAbout(a);
NS_1.SUB_2.AEx.WriteAbout(a);

// -----------------------------------------------------------------------

public static class StringExtension
{
    // открытый метод
    public static int CharCount(this string str, char c)
    {
        int counter = 0;
        for (int i = 0; i < str.Length; i++)
        {
            if (str[i] == c)
            {
                counter++;
            }
        }
        return counter;
    }
}

// -----------------------------------------------------------------------

public class A
{
    public double NumVal { get; set; }
}

public class B : A
{
    private string strVal = new string("");
    public string StrVal { get => strVal; set => strVal = value; }
}

// ***

public static class AEx
{
    public static double NumMult(this A a, uint count)
    {
        return a.NumVal * count;
    }
}

public static class BEx
{
    public static string StrMult(this B b, uint count)
    {
        var stringBuilder = new StringBuilder();
        for (int i = 0; i < count; ++i)
        {
            stringBuilder.Append(b.StrVal);
        }
        return stringBuilder.ToString();
    }
}

// -----------------------------------------------------------------------

namespace NS_1
{
    public class A
    {
        public A(string name)
        {
            Name = name;
        }
        public string Name { get; set; }
    }

    // ***

    namespace SUB_1
    {
        public static class AEx
        {
            public static void WriteAbout(this A a)
            {
                Console.WriteLine("AEx +++ From SUB_1. Has Name: " + a.Name);
            }
        }
    }

    namespace SUB_2
    {
        public static class AEx
        {
            public static void WriteAbout(this A a)
            {
                Console.WriteLine("AEx --- From SUB_2. Has Name: " + a.Name);
            }
        }
    }
}

