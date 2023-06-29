#include <iostream>

struct A
{
    virtual ~A() = default;
    void do_all() {
        do_before();

        do_common();

        do_after();
    }

private:
    void do_before() {
        std::cout << "A before" << std::endl;
    }
    void do_after() {
        std::cout << "A after" << std::endl;
    }

protected:
    /* template method? */
    virtual void do_common() = 0;
};

struct B : A
{
// protected:
    void do_common() override {
        std::cout << "B do..." << std::endl;
    }
};

struct C : A
{
protected:
    void do_common() override {
        std::cout << "C do..." << std::endl;
    }
};

int main()
{
    {
        B b;
        b.do_common();
    }
    std::cout << std::endl;
    {
        C c;
        //c.do_common();
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
