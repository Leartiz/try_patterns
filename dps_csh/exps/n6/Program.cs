//#define CONTRACTS_FULL

using System;
using System.Diagnostics.CodeAnalysis;
using System.Diagnostics.Contracts;

// ***

try
{
    Account acc = new(-100.0);
    Console.WriteLine("Balance: " + acc.Balance);
}
catch (ApplicationException err) 
{
    Console.WriteLine("exc: " + err.Message);
}

// ***

Console.WriteLine(default(int));
Console.WriteLine(default(string));

Console.WriteLine(DefaultCreator<int>.Give());
Console.WriteLine(DefaultCreator<string>.Give());

// ***

Console.WriteLine(default(string) == null);
Console.WriteLine(DefaultCreator<string>.Give() == null);

Console.WriteLine(default(string)! == null);
Console.WriteLine(DefaultCreator<string>.Give() == null);

// -----------------------------------------------------------------------

[ContractClass(typeof(LogImporterContract))]
public abstract class LogImporter
{
    public abstract IEnumerable<string> ReadEntries(ref int position);
    public abstract string ParseLogEntry(string stringEntry);
}

// пишется отдельно (?), а также ограничения повторяются при реализации
[ExcludeFromCodeCoverage, ContractClassFor(typeof(LogImporter))]
public abstract class LogImporterContract : LogImporter
{
    public override IEnumerable<string> ReadEntries(ref int position)
    {
        Contract.Ensures(Contract.Result<IEnumerable<string>>() != null);

        Contract.Ensures(
            // внутри метода { position = ...; return position; }
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

public static class DefaultCreator<T>
{
    public static T Give()
    {
        // игнорировать NULL ссылку (?)
        return default!;
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
            if (balance < 0.0)
                throw new ApplicationException("Account, get, balance < 0");
            return balance;
        }
        set
        {
            if (value < 0.0)
                throw new ApplicationException("Account, get, value < 0");
            balance = value;
        }
    }

    public Account(double balance = 0)
    {
        Contract.Requires(balance >= 0);

        // ***

        if (balance < 0)
            throw new ApplicationException("Account, ctor, balance < 0");

        this.balance = balance;
    }

    public void Deposit(double amount)
    {
        Contract.Requires(amount >= 0);

        // ***

        if (amount < 0)
            throw new ApplicationException("Account, Deposit, amount < 0");

        Balance += amount;
    }

    public void Withdraw(double amount)
    {
        Contract.Requires(amount > 0);
        Contract.Requires(amount <= Balance);

        // ***

        if (amount <= 0)
            throw new ApplicationException("Account, Withdraw, amount <= 0");
        if (amount > Balance)
            throw new ApplicationException("Account, Withdraw, amount > Balance");

        Balance -= amount;
    }

    public Account Transfer()
    {
        Contract.Ensures(Balance == 0);

        Account newAccount = new(Balance);
        Balance = 0;

        return newAccount;
    }

    // ***

    // проверка инвариантов класса происходит
    // в конце каждого public метода класса
    [ContractInvariantMethod]
    protected void ValidAccount()
    {
        Contract.Invariant(Balance >= 0);
    }
}