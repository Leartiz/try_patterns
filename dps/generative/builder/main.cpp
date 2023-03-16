#include <iostream>
#include <vector>

using namespace std;

// -----------------------------------------------------------------------

// Классы всех возможных родов войск
class Infantryman final
{
public:
    void info() {
        cout << "Infantryman" << endl;
    }
};

class Archer final
{
public:
    void info() {
        cout << "Archer" << endl;
    }
};

class Horseman final
{
public:
    void info() {
        cout << "Horseman" << endl;
    }
};

class Catapult final
{
public:
    void info() {
        cout << "Catapult" << endl;
    }
};

class Elephant final
{
public:
    void info() {
        cout << "Elephant" << endl;
    }
};

// -----------------------------------------------------------------------

// Класс "Армия", содержащий все типы боевых единиц
class Army final
{
public:
    vector<Infantryman> vi;
    vector<Archer>      va;
    vector<Horseman>    vh;
    vector<Catapult>    vc;
    vector<Elephant>    ve;

public:
    void info() {
        size_t i{ 0 };
        for(i=0; i < vi.size(); ++i) vi[i].info();
        for(i=0; i < va.size(); ++i) va[i].info();
        for(i=0; i < vh.size(); ++i) vh[i].info();
        for(i=0; i < vc.size(); ++i) vc[i].info();
        for(i=0; i < ve.size(); ++i) ve[i].info();
    }
};

// Базовый класс ArmyBuilder объявляет интерфейс для поэтапного
// построения армии и предусматривает его реализацию по умолчанию

class ArmyBuilder
{
public:
    virtual ~ArmyBuilder() {}

public:
    virtual void createArmy() {}
    virtual void buildInfantryman() {}
    virtual void buildArcher() {}
    virtual void buildHorseman() {}

    virtual void buildCatapult() {}
    virtual void buildElephant() {}

public:
    Army* releaseArmy()
    {
        auto army{ p };
        p = nullptr;
        return army;
    }

protected:
    Army* p{ nullptr };
};

// -----------------------------------------------------------------------

class Director;

// Римская армия имеет все типы боевых единиц кроме боевых слонов
class RomanArmyBuilder : public ArmyBuilder
{
    friend Director;
protected:
    void createArmy() override { p = new Army; }
    void buildInfantryman() override { p->vi.push_back(Infantryman()); }
    void buildArcher() override { p->va.push_back(Archer()); }
    void buildHorseman() override { p->vh.push_back(Horseman()); }

    void buildCatapult() override { p->vc.push_back(Catapult()); }
};

// Армия Карфагена имеет все типы боевых единиц кроме катапульт
class CarthaginianArmyBuilder : public ArmyBuilder
{
    friend Director;
protected:
    void createArmy() override { p = new Army; }
    void buildInfantryman() override { p->vi.push_back(Infantryman()); }
    void buildArcher() override { p->va.push_back(Archer()); }
    void buildHorseman() override { p->vh.push_back(Horseman()); }

    void buildElephant() override { p->ve.push_back(Elephant()); }
};

// Класс-распорядитель, поэтапно создающий армию той или иной стороны.
// Именно здесь определен алгоритм построения армии.
class Director
{
public:
    Army* createArmy(ArmyBuilder &builder)
    {
        builder.createArmy();

        builder.buildInfantryman();
        builder.buildArcher();
        builder.buildHorseman();

        builder.buildCatapult();
        builder.buildElephant();

        return builder.releaseArmy();
    }
};

// -----------------------------------------------------------------------

int main()
{
    Director dir;

    RomanArmyBuilder ra_builder;
    CarthaginianArmyBuilder ca_builder;

    Army *ra = dir.createArmy(ra_builder);
    Army *ca = dir.createArmy(ca_builder);

    {
        cout << "Roman army:" << endl;
        ra->info();

        cout << "\nCarthaginian army:" << endl;
        ca->info();
        // ...
    }

    delete ra;
    delete ca;

    return 0;
}
