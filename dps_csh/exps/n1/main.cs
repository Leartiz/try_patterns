namespace N1
{
    interface ILogSaver
    {
        void UploadLogEntries(IEnumerable<string> logEntries);
        void UploadExceptions(IEnumerable<string> exceptions);
    }

    // *** конкретная реализация интерфейса

    public class Channel : ILogSaver
    {
        private Random rnd = new Random();
        private bool hasError()
        {
            return rnd.Next(0, 2) == 1;
        }

        // ***

        public void UploadLogEntries(IEnumerable<string> logEntries)
        {
            if (hasError())
            {
                throw new CommunicationException(new SystemException());
            }

            Console.WriteLine("Channel, UploadLogEntries:");
            Console.WriteLine("{ ");
            foreach (var one in logEntries)
            {
                Console.WriteLine(one);
            }
            Console.WriteLine("}");
        }
        public void UploadExceptions(IEnumerable<string> exceptions)
        {
            if (hasError())
            {
                throw new CommunicationException(new SystemException());
            }

            Console.WriteLine("Channel, UploadExceptions:");
            Console.WriteLine("{ ");
            foreach (var one in exceptions)
            {
                Console.WriteLine(one);
            }
            Console.WriteLine("}");
        }
    }

    // *** исключения

    public class CommunicationException : SystemException 
    {
        public CommunicationException(SystemException se)
        {
            Console.WriteLine("CommunicationException, ctor, msg: " + se.Message);
        }
    }
    public class OperationFailedException : SystemException
    {
        public OperationFailedException(SystemException se) 
        {
            Console.WriteLine("OperationFailedException, ctor, msg: " + se.Message);
        }
    }

    // ***
    /* заметки:
       - возможно не хватает виртуального метода (реализация ILogSaver без внедрения)
    */

    public class ClientBase<T>
    {
        protected Channel Channel = new Channel();

        public void Close()
        {
            Console.WriteLine("ClientBase, Close.");
        }
        public void Abort()
        {
            Console.WriteLine("ClientBase, Abort.");
        }
    }

    // ***
    /* заметки:
       - не вижу в коде внедрения зависимостей (конкретно делегата/функции)
    */

    public class LogSaverProxy : ILogSaver
    {
        class LogSaverClient : ClientBase<ILogSaver>
        {
            // должен быть виртуальным?
            public ILogSaver LogSaver
            {
                get { return Channel; }
            }
        }

        // ***

        public void UploadLogEntries(IEnumerable<string> logEntries)
        {
            UseProxyClient((ILogSaver c) => c.UploadLogEntries(logEntries));
        }

        public void UploadExceptions(IEnumerable<string> exceptions)
        {
            UseProxyClient((ILogSaver c) => c.UploadExceptions(exceptions));
        }

        // ***

        private void UseProxyClient(Action<ILogSaver> accessor)
        {
            var client = new LogSaverClient(); // DI?
            try
            {
                accessor(client.LogSaver);
                client.Close();
            }
            catch (CommunicationException e)
            {
                client.Abort();
                throw new OperationFailedException(e);
            }
        }
    }

    public class Program
    {
        static int Main(string[] args)
        {
            {
                var foo = (int x) => Console.WriteLine(x);
                foo(100);
            }
            // ***
            {
                try
                {
                    LogSaverProxy logSaverProxy = new LogSaverProxy();
                    string[] days = {
                        "Monday", "Tuesday", "Wednesday",
                        "Thursday", "Friday", "Saturday",
                        "Sunday"
                    };
                    logSaverProxy.UploadLogEntries(days);
                }
                catch(SystemException se)
                {
                    Console.WriteLine("Program, Main, msg: " + se.Message);
                }                
            }
            {
                try
                {
                    LogSaverProxy logSaverProxy = new LogSaverProxy();
                    string[] days = {
                        "Monday", "Tuesday", "Wednesday",
                        "Thursday", "Friday", "Saturday",
                        "Sunday"
                    };
                    logSaverProxy.UploadExceptions(days);
                }
                catch (SystemException se)
                {
                    Console.WriteLine("Program, Main, msg: " + se.Message);
                }
            }

            return 0;
        }
    }
}