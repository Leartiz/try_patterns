#ifndef LIST_H
#define LIST_H

#include <vector>
#include <cstdlib>
#include <stdexcept>

class List;
class Node;

// -----------------------------------------------------------------------

class Node {
    friend class ::List;

public:
    Node(int value, int level)
        : value{ value } {
        if (level < 1)
            throw std::runtime_error("node level less than one");

        nodes.resize(level, nullptr);
    }

    int node_count() const {
        return static_cast<int>(nodes.size());
    }

private:
    int value;
    std::vector<Node*> nodes;
};

// -----------------------------------------------------------------------

class List {
public:
    explicit List(int max_level)
        : max_level{ max_level } {
        if (max_level < 1)
            throw std::runtime_error("list maximum level less than one");

        beg_nodes.resize(max_level, nullptr);
        // front is at the bottom!
    }
    ~List() {
        auto step = beg_nodes.front();
        while (step) {
            auto next = step->nodes.front();
            delete step; step = next;
        }
    }

    // ***

    void add_value(int value) {
        const auto level = random_level();
        const auto new_node = new Node(value, level);

        if (!beg_nodes.front()) { // empty list!
            for (int i = level - 1; i >= 0; --i)
                beg_nodes[i] = new_node;
            cur_level = level;
            return;
        }

        // ***

        if (value < beg_nodes.front()->value) { // add value to begin list!
            for (int i = level - 1; i >= 0; --i) {
                new_node->nodes[i] = beg_nodes[i]; // mb nullptr.
                beg_nodes[i]       = new_node; // move some of the pointers!
            }

            if (level > cur_level)
                cur_level = level;
            return;
        }
        if (level > cur_level) {
            for (int i = cur_level - 1; i < level; ++i)
                beg_nodes[i] = new_node;
            cur_level = level;
        }

        // ***

        // because the level is different!
        auto prev_nodes = std::vector<Node*>(cur_level, nullptr);
        auto cur_node = beg_nodes[cur_level - 1];
        for (auto i = cur_level - 1; i >= 0; --i) {
            while (cur_node->nodes[i] && value > cur_node->value)
                cur_node = (cur_node->nodes[i]);

            prev_nodes[i] = cur_node;
        }

        for (auto i = level - 1; i >= 0; --i) {
            new_node->nodes[i] = prev_nodes[i]->nodes[i];
            prev_nodes[i]->nodes[i] = new_node; // replace!
        }
    }

    void println() {
        if (!beg_nodes.front()) {
            std::cout << "<empty>" << std::endl;
            return;
        }

        auto current_node = beg_nodes.front(); // bottom line!
        while (current_node != nullptr) {
            std::cout
                << "value: " << current_node->value << " | "
                << "level: " << current_node->nodes.size() << " | "
                << std::endl;
            current_node = current_node->nodes.front();
        }
    }

private:
    int random_level() {
        return (rand() % max_level) + 1;
    }

    std::vector<Node*> beg_nodes;
    const int max_level;
    int cur_level = -1; // )
};

#endif // LIST_H
