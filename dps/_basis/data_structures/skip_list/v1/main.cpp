#include <iostream>
#include <cstdlib>

#include "list.h"

using namespace std;

int main() {
   // std::srand(time(nullptr));

    List ll(5);
    ll.add_value(100);
    ll.add_value(105);
    //ll.add_value(95);
    ll.add_value(110);
//    ll.add_value(115);
//    ll.add_value(120);
//    ll.add_value(125);
    ll.println();
}
