#include <cassert>

#include <iostream>
#include <memory>

// -----------------------------------------------------------------------

class Warrior
{
public:
    virtual void info() const = 0;
    virtual ~Warrior() {}
};

// -----------------------------------------------------------------------

class Infantryman final : public Warrior
{
public:
    void info() const override
    {
        std::cout << "Infantryman: {}" << std::endl;
    }

    ~Infantryman() override
    {
        std::cout << "~Infantryman()" << std::endl;
    }
};

class Horseman final : public Warrior
{
public:
    void info() const override
    {
        std::cout << "Horseman: {}" << std::endl;
    }

    ~Horseman() override
    {
        std::cout << "~Horseman()" << std::endl;
    }
};

class Archer final : public Warrior
{
public:
    void info() const override
    {
        std::cout << "Archer: {}" << std::endl;
    }

    ~Archer() override
    {
        std::cout << "~Archer()" << std::endl;
    }
};

// -----------------------------------------------------------------------

enum Warrior_ID
{
    Infantryman_ID = 0,
    Horseman_ID,
    Archer_ID,
};

// Фабричная функция.
/*

Решит проблему, когда нужно будет заменить возврат, для этого:
- изменить имя одного из элементов перечисления;
- вернуть новый созданный объект с другим типом.

Правда, в этом случае delete - явный!

*/
Warrior* createWarrior(const Warrior_ID id)
{
    Warrior *p;
    switch (id)
    {
        case Infantryman_ID:
            p = new Infantryman();
            break;
        case Horseman_ID:
            p = new Horseman();
            break;
        case Archer_ID:
            p = new Archer();
            break;
        default:
            assert(false);
    }
    return p;
}

std::shared_ptr<Warrior> createWarriorShPtr(const Warrior_ID id)
{
    switch (id) {
    case Infantryman_ID:
        return std::shared_ptr<Infantryman>(new Infantryman());
    case Horseman_ID:
        return std::shared_ptr<Horseman>(new Horseman());
    case Archer_ID:
        return std::shared_ptr<Archer>(new Archer());
    }

    assert(false);
    return nullptr;
}

// -----------------------------------------------------------------------

int main()
{
    // сreateWarrior
    {
        {
            auto w{ createWarrior(Infantryman_ID) };
            w->info();
            delete w;
        }
        {
            auto w{ createWarrior(Horseman_ID) };
            w->info();
            delete w;
        }
        {
            auto w{ createWarrior(Archer_ID) };
            w->info();
            delete w;
        }
    }

    // сreateWarriorShPtr
    {
        auto w1{ createWarriorShPtr(Infantryman_ID) };
        auto w2{ createWarriorShPtr(Horseman_ID) };
        auto w3{ createWarriorShPtr(Archer_ID) };
    }
    return 0;
}
