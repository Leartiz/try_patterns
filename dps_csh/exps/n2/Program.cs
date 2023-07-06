using System.Text;

string s = "Работать надо и все будет?";
char c = 'а';

int res = s.CharCount(c);
Console.WriteLine(res);

B b = new B();
Console.WriteLine(b.StrMult(5));

// ***

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

