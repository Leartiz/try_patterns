#include <iostream>

class Warrior
{
  public:
    virtual void info() = 0;
    virtual ~Warrior() {}
};

class Infantryman: public Warrior
{
  public:
      void info() { std::cout << "Infantryman" << std::endl; }
};

class Archer: public Warrior
{
  public:
    void info() { std::cout << "Archer" << std::endl; }
};

class Horseman: public Warrior
{
  public:
    void info() { std::cout << "Horseman" << std::endl; }
};

int main()
{
    std::cout << "Hello World!" << std::endl;
    return 0;
}
