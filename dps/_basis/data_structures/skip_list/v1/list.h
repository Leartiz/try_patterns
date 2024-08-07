#ifndef LIST_H
#define LIST_H

#include <vector>
#include <cstdlib>

class List;
class Node;

// -----------------------------------------------------------------------

class Node {
    friend class ::List;

public:
    Node(int value, size_t level)
        : value{ value }
        , nodes{ level, nullptr } {}

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
        : beg{ nullptr }
        , max_level{ max_level } {}
    ~List() {

    }

    void add_value(int value) {
        if (!beg) {
            beg = new Node(value, max_level); // !
            return;
        }

        const auto level = random_level();
        const auto new_node = new Node(value, level);

        // because the level is different!
        auto prev_nodes = std::vector<Node*>(
            max_level, nullptr);

        // ***

        if (value < beg->value) {
            for (int i = level - 1; i >= 0; --i) {
                new_node->nodes[i] = beg->nodes[i];
            }

            beg = new_node;
            return;
        }

        Node* current_node = beg; // with max level.
        for (auto i = current_node->node_count() - 1; i >= 0; --i) {
            while (true) { // by level!
                if (i >= current_node->node_count()) break;
                if (!current_node->nodes[i]) break;

                if (value >= current_node->nodes[i]->value)
                    current_node = current_node->nodes[i];
                else break;
            }

            prev_nodes[i] = current_node;
        }

        // new_node->node_count() < max_level!

        for (auto i = level - 1; i >= 0; --i) {
            if (i >= prev_nodes[i]->node_count())
                continue;

            new_node->nodes[i] = prev_nodes[i]->nodes[i]; // next node!
            prev_nodes[i]->nodes[i] = new_node;
        }
    }

    void println() {
        if (!beg) {
            std::cout << "<empty>" << std::endl;
            return;
        }

        auto current_node = beg;
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
        return (rand() + 1) % max_level;
    }

    Node* beg;
    const int max_level;
};

#endif // LIST_H
