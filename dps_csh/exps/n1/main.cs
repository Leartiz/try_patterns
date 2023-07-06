namespace Delegate
{
    interface ILogSaver
    {
        void UploadLogEntries(IEnumerable<string> logEntries);
        void UploadExceptions(IEnumerable<string> exceptions);
    }

    class Channel : ILogSaver
    {
        public void UploadLogEntries(IEnumerable<string> logEntries)
        {
            Console.WriteLine("Channel, UploadLogEntries.");
        }
        public void UploadExceptions(IEnumerable<string> exceptions)
        {
            Console.WriteLine("Channel, UploadExceptions.");
        }
    }

    class CommunicationException : SystemException { }
    class OperationFailedException : SystemException
    {
        public OperationFailedException(SystemException se) 
        { 

        }
    }

    class ClientBase<T>
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


    class LogSaverProxy : ILogSaver
    {
        class LogSaverClient : ClientBase<ILogSaver>
        {
            public ILogSaver LogSaver
            {
                get { return Channel; }
            }
        }
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
            var client = new LogSaverClient();
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

    public class Exp
    {
        static int Main(string[] args)
        {
            {
                var foo = (int x) => Console.WriteLine(x);
                foo(100);
            }

            return 0;
        }
    }
}