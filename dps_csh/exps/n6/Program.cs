#define CONTRACTS_FULL

using System.Diagnostics.CodeAnalysis;
using System.Diagnostics.Contracts;

// ***

Account acc = new Account(100.0);
Console.WriteLine("Balance: " + acc.Balance);

// -----------------------------------------------------------------------

[ContractClass(typeof(LogImporterContract))]
public abstract class LogImporter
{
    public abstract IEnumerable<string> ReadEntries(ref int position);
    public abstract string ParseLogEntry(string stringEntry);
}

// пишется отдельно, а также ограничения повторяются при реализации
[ExcludeFromCodeCoverage, ContractClassFor(typeof(LogImporter))]
public abstract class LogImporterContract : LogImporter
{
    public override IEnumerable<string> ReadEntries(ref int position)
    {
        Contract.Ensures(Contract.Result<IEnumerable<string>>() != null);
        Contract.Ensures(
            Contract.ValueAtReturn(out position) >= 
            Contract.OldValue(position));
        
        throw new NotImplementedException();
    }
    public override string ParseLogEntry(string stringEntry)
    {
        Contract.Requires(stringEntry != null);
        Contract.Ensures(Contract.Result<string>() != null);

        throw new NotImplementedException();
    }
}

// -----------------------------------------------------------------------

public class MemLogImporter : LogImporter
{
    public override IEnumerable<string> ReadEntries(ref int position)
    {
        Contract.Ensures(Contract.Result<IEnumerable<string>>() != null);
        Contract.Ensures(
            Contract.ValueAtReturn(out position) >=
            Contract.OldValue(position));

        return new string[1];
    }

    public override string ParseLogEntry(string stringEntry)
    {


        throw new NotImplementedException();
    }
}

// -----------------------------------------------------------------------

public class Account
{
    private double balance = 0.0;
    public double Balance
    {
        get
        {
            var cond = () =>
            {
                Console.WriteLine(Contract.Result<double>());
                return Contract.Result<double>() >= 100.0;
            };
            Contract.Ensures(cond()) ;
            return balance;
        }
        set
        {
            Contract.Requires(value >= 10000.0);
            balance = value;
        }
    }

    public Account(double balance = 0)
    {
        Balance = balance;
    }

    public void Deposit(double amount)
    {
        Contract.Requires(amount >= 0);

        Balance += amount;
    }

    public void Withdraw(double amount)
    {
        Contract.Requires(amount > 0);
        Contract.Requires(amount <= Balance);

        Balance -= amount;
    }

    public Account Transfer()
    {
        Contract.Ensures(Balance == 0);

        Account newAccount = new Account()
        {
            Balance = this.Balance 
        };
        Balance = 0;
        return newAccount;
    }

    [ContractInvariantMethod]
    protected void ValidAccount()
    {
        Contract.Invariant(Balance >= 0);
    }
}