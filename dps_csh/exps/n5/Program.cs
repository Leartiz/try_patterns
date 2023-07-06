using NUnit.Framework;
using System.IO;
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
    private Lazy<Stream> _stream;

    // ***

    protected LogFileReaderBase(string fileName)
    {
        _stream = new Lazy<Stream>(() => OpenFileStream(fileName));
    }

    // базовая реализация
    protected virtual Stream OpenFileStream(string fileName)
    {
        return new FileStream(fileName, FileMode.Open);
    }

    // ***

    public IEnumerable<string> ReadLogEntry()
    {
        StreamReader sr = new StreamReader(_stream.Value);
        return sr.ReadToEnd().Split(new[] { "\r\n" }, StringSplitOptions.None);
    }
}

public class FakeLogFileReader : LogFileReaderBase
{
    private readonly MemoryStream _mockStream;
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

// ***

public class TestLogFileReader
{
    public static MemoryStream GetMemoryStreamWithOneElement()
    {
        Random rnd = new Random();
        MemoryStream ms = new MemoryStream();

        byte[] oneByte = new byte[1];
        rnd.NextBytes(oneByte);

        ms.Write(oneByte, 0, oneByte.Length);
        return ms;
    }

    // ***

    [Test]
    public void TestFakedMemoryStreamProvidedOneElement()
    {
        // Arrange
        LogFileReaderBase cut = new FakeLogFileReader(
            GetMemoryStreamWithOneElement());

        // Act
        var logEntries = cut.ReadLogEntry();
        // Assert
        Assert.That(logEntries.Count(), Is.EqualTo(1));
    }
}