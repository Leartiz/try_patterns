using NUnit.Framework;
using System.Text;

// ***

byte[] buffer = new byte[32];
Random rnd = new Random();
rnd.NextBytes(buffer);

buffer = Encoding.UTF8.GetBytes("ТестTest");
Console.WriteLine(Encoding.UTF8.GetString(buffer));

// ***

MemoryStream ms = new MemoryStream();
ms.Write(buffer, 0, buffer.Length);

byte[] newBuffer = new byte[32];
ms.Read(newBuffer, 0, newBuffer.Length);
Console.WriteLine(Encoding.UTF8.GetString(buffer));

// ***

A a = new A();
ms.Position = 0;
a.PrintToConsole(ms);

// -----------------------------------------------------------------------

public class A
{
    public void PrintToConsole(Stream s)
    {
        int one = s.ReadByte();
        while (one != -1)
        {
            StringBuilder sb = new StringBuilder();
            sb.AppendLine("Posn: " + s.Position);
            sb.AppendLine("Byte: " + (char)one);
            sb.Append("   ***   ");

            Console.WriteLine(sb.ToString());
            one = s.ReadByte();
        }
    }
}

// -----------------------------------------------------------------------

public abstract class LogFileReaderBase
{
    private readonly Lazy<Stream> _stream;

    // ***

    protected LogFileReaderBase(string fileName)
    {
        _stream = new Lazy<Stream>(() => OpenFileStream(fileName));
    }

    // базовая реализация (?)
    protected virtual Stream OpenFileStream(string fileName)
    {
        return new FileStream(fileName, FileMode.Open);
    }

    // ***

    public IEnumerable<string> ReadLogEntry()
    {
        StreamReader sr = new StreamReader(_stream.Value);
        List<string> lines = new List<string>(sr.ReadToEnd().Split(new[] { "\r\n" },
            StringSplitOptions.None));

        if (lines.Count > 0)
        {
            int endIndex = lines.Count - 1;
            if (lines[endIndex].Length == 0)
                lines.RemoveAt(endIndex);
        }
        return lines;
    }
}

public class FakeLogFileReader : LogFileReaderBase
{
    private readonly MemoryStream _mockStream;

    // так как ленивая инициализация потока...
    public FakeLogFileReader(MemoryStream mockStream)
        : base(string.Empty)
    {
        _mockStream = mockStream;
    }
    protected override Stream OpenFileStream(string fileName)
    {
        return _mockStream;
    }
}

// -----------------------------------------------------------------------

// атрибут, помечающий класс, содержащий тесты и,
// при необходимости, методы настройки или демонтажа.
// с NUnit 2.5 атрибут TestFixture является необязательным (?)
[TestFixture]
public class TestLogFileReader
{
    public static void NextLetters(byte[] buffer)
    {
        var rnd = new Random();
        string letters = new string("ABCDEabcde1234567890_");

        for (int i = 0; i < buffer.Length; i++)
        {
            var anyIndex = rnd.Next(letters.Length);
            buffer[i] = (byte)letters[anyIndex];
        }

        // ***

        Console.WriteLine(Encoding.UTF8.GetString(buffer));
    }

    public static MemoryStream GetMemoryStreamWithSomeElement(int n)
    {
        MemoryStream ms = new MemoryStream();
        for (int i = 0; i < n; ++i)
        {
            byte[] buffer = new byte[16];
            NextLetters(buffer);

            // ***

            var list = new List<byte>(buffer)
            {
                (byte)'\r',
                (byte)'\n'
            };

            ms.Write(list.ToArray(), 0, list.Count);
        }

        ms.Position = 0;
        return ms;
    }

    // ***

    [Test]
    public void TestFakedMemoryStreamProvidedOneElement()
    {
        // Arrange
        LogFileReaderBase cut = new FakeLogFileReader(
            GetMemoryStreamWithSomeElement(1));

        // Act
        var logEntries = cut.ReadLogEntry();
        // Assert
        Assert.That(logEntries.Count(), Is.EqualTo(1));
    }

    [Test]
    public void TestFakedMemoryStreamProvidedTwoElement()
    {
        // Arrange
        LogFileReaderBase cut = new FakeLogFileReader(
            GetMemoryStreamWithSomeElement(2));

        // Act
        var logEntries = cut.ReadLogEntry();
        // Assert
        Assert.That(logEntries.Count(), Is.EqualTo(2));
    }

    // ...
}