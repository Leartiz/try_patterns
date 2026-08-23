#include <iostream>
#include <string>

// Fields in "random" order -> more padding between members.
struct LayoutScattered {
    char ch;
    std::string str;
    double num;
    float num1;
    int num2;
    char ch1;
};

// Same layout as LayoutScattered; public: sections do not change size.
struct LayoutScatteredPublic {
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

// Same types, better order (small fields together) -> often less padding.
struct LayoutGrouped {
    char ch;
    char ch1;
    int num2;
    float num1;
    double num;
    std::string str;
};

// Same as LayoutGrouped; public: still does not change size.
struct LayoutGroupedPublic {
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

// Same order as LayoutGrouped, but pack(1) removes most padding.
#pragma pack(push, 1)
struct LayoutGroupedPack1 {
    char ch;
    char ch1;
    int num2;
    float num1;
    double num;
    std::string str;
};
#pragma pack(pop)

int main() {
    std::cout << "sizeof(LayoutScattered):       " << sizeof(LayoutScattered) << '\n';
    std::cout << "sizeof(LayoutScatteredPublic): " << sizeof(LayoutScatteredPublic) << '\n';
    std::cout << "sizeof(LayoutGrouped):         " << sizeof(LayoutGrouped) << '\n';
    std::cout << "sizeof(LayoutGroupedPublic):   " << sizeof(LayoutGroupedPublic) << '\n';
    std::cout << "sizeof(LayoutGroupedPack1):    " << sizeof(LayoutGroupedPack1) << '\n';

    // NOTE:
    /*
        sizeof(LayoutScattered):       64
        sizeof(LayoutScatteredPublic): 64
        sizeof(LayoutGrouped):         56
        sizeof(LayoutGroupedPublic):   56
        sizeof(LayoutGroupedPack1):    50
    */
    return 0;
}
