#include <iostream>

struct A
{
    virtual ~A() = default;
    virtual void do_all() = 0;

protected:
    void do_before() {
        std::cout << "A before" << std::endl;
    }
    void do_after() {
        std::cout << "A after" << std::endl;
    }
};

struct B : A
{
    void do_all() override {
        do_before();

        std::cout << "B do..." << std::endl;

        do_after();
    }
};

struct C : A
{
    void do_all() override {
        do_before();
        do_before();
        do_before();
        // ...

        std::cout << "C do..." << std::endl;

        do_after();
        do_after();
        do_after();
        // ...
    }
};

int main()
{
    {
        B b;
        b.do_all();
    }
    std::cout << std::endl;
    {
        C c;
        c.do_all();
    }
    std::cout << std::endl;

    // ***

    {
        A* b = new B;
        b->do_all();
        delete b;
    }
    std::cout << std::endl;
    {
        A* c = new C;
        c->do_all();
        delete c;
    }
}
