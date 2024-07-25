#include <string>
#include <iostream>

struct Av1 final
{
    char ch;
    std::string str;
    double num;
    float num1;
    int num2;
    char ch1;
};

struct Av2 final
{
public:
    char ch;
public:
    std::string str;
public:
    double num;
public:
    float num1;
public:
    int num2;
public:
    char ch1;
};

struct Av3 final
{
    char ch;
    char ch1;
    int num2;
    float num1;
    double num;
    std::string str;
};

struct Av4 final
{
public:
    char ch;
public:
    char ch1;
public:
    int num2;
public:
    float num1;
public:
    double num;
public:
    std::string str;
};

#pragma pack(push, 1)
struct Av5
{
    char ch;
    char ch1;
    int num2;
    float num1;
    double num;
    std::string str;
};
#pragma pack(pop)

// ------------------------------------------------------------------

int main()
{
    std::cout << "sizeof(Av1): " << sizeof(Av1) << std::endl;
    std::cout << "sizeof(Av2): " << sizeof(Av2) << std::endl;
    std::cout << "sizeof(Av3): " << sizeof(Av3) << std::endl;
    std::cout << "sizeof(Av4): " << sizeof(Av4) << std::endl;
    std::cout << "sizeof(Av5): " << sizeof(Av5) << std::endl;
    return 0;
}
