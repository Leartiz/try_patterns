#include <iostream>
#include <functional>

struct A final
{
    void do_all(const std::function<void()>& concrete) {
        do_beg();

        concrete();

        do_end();
    }

private:
    void do_beg() {
        std::cout << "A do_beg" << std::endl;
    }
    void do_end() {
        std::cout << "A do_end" << std::endl;
    }
};

int main()
{
    A a;
    a.do_all([]() {
       std::cout << "do concrete..." << std::endl;
    });
    return 0;
}
