#include <iostream>
#include <random>

#include <vector>
#include <string>

using namespace std;

/*

Рассмотрим реализацию паттерна Command на примере игры «Шахматы».
Имитируем возможность выполнения следующих операций:

1. Создать новую игру.
2. Открыть существующую игру.
3. Сохранить игру.
4. Сделать очередной ход.
5. Отменить последний ход.

*/

class Game final
{
public:
    void create()
    {
        cout << "Create game " << endl;
    }
    void open(string file)
    {
        cout << "Open game from " << file << endl;
    }
    void save(string file)
    {
        cout << "Save game in " << file << endl;
    }
    void make_move(string move)
    {
        cout << "Make move " << move << endl;
    }
};

// -----------------------------------------------------------------------

string generateFileName(size_t len = 10)
{
    static const std::string chars("0123456789"
                                   "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
                                   "abcdefghijklmnopqrstuvwxyz"
                                   );
    static std::default_random_engine eng{ std::random_device{}() };
    static std::uniform_int_distribution<size_t> dist(0, chars.size() - 1); // [beg, end]

    std::string result(len, 0);
    for (auto& ch : result) {
        ch = chars[dist(eng)];
    }
    return result + ".sav";
}

string generateChessStep()
{
    static std::default_random_engine eng{ std::random_device{}() };
    static std::uniform_int_distribution<size_t> dist(0, 7);

    static const std::string letters{ "ABCDEFGH" };
    static const std::string numbers{ "12345678" };

    std::string result;
    result.reserve(5);

    result.push_back(letters[dist(eng)]);
    result.push_back(numbers[dist(eng)]);
    result += '-';
    result.push_back(letters[dist(eng)]);
    result.push_back(numbers[dist(eng)]);
    return result;
}

string getPlayerFileName(string prompt)
{
    string input{ generateFileName() };
    cout << prompt << " " << input << endl;
    return input;
}

string getPlayerChessStep(string prompt)
{
    string input{ generateChessStep() };
    cout << prompt << " " << input << endl;
    return input;
}


// -----------------------------------------------------------------------

// Базовый класс
class Command
{
public:
    virtual ~Command() {}
    virtual void execute() = 0;

protected:
    Command(Game* p)
        : pgame(p)
    {
    }

    Game* pgame;
};

class CreateGameCommand : public Command
{
public:
    CreateGameCommand(Game* p)
        : Command(p)
    {
    }
    void execute()
    {
        pgame->create();
    }
};

class OpenGameCommand : public Command
{
public:
    OpenGameCommand(Game* p)
        : Command(p)
    {
    }
    void execute()
    {
        string file_name;
        file_name = getPlayerFileName("Enter file name:");
        pgame->open(file_name);
    }
};

class SaveGameCommand : public Command
{
public:
    SaveGameCommand(Game* p)
        : Command(p)
    {
    }
    void execute()
    {
        string file_name;
        file_name = getPlayerFileName("Enter file name:");
        pgame->save(file_name);
    }
};

class MakeMoveCommand : public Command {
public:
    MakeMoveCommand(Game* p)
        : Command(p)
    {
    }
    void execute()
    {
        // Сохраним игру для возможного последующего отката
        pgame->save("TEMP_FILE");
        string move;
        move = getPlayerChessStep("Enter your move:");
        pgame->make_move(move);
    }
};

class UndoCommand : public Command {
public:
    UndoCommand(Game* p)
        : Command(p)
    {
    }
    void execute()
    {
        // Восстановим игру из временного файла
        pgame->open("TEMP_FILE");
    }
};

// -----------------------------------------------------------------------

int main()
{
    Game game;
    // Имитация действий игрока (обычно это одна команда?)
    vector<Command*> v;

    // Создаем новую игру
    v.push_back(new CreateGameCommand(&game));
    // Делаем несколько ходов
    v.push_back(new MakeMoveCommand(&game));
    v.push_back(new MakeMoveCommand(&game));
    // Последний ход отменяем
    v.push_back(new UndoCommand(&game));
    // Сохраняем игру
    v.push_back(new SaveGameCommand(&game));

    for (size_t i = 0; i < v.size(); ++i)
        v[i]->execute();

    for (size_t i = 0; i < v.size(); ++i)
        delete v[i];

    return 0;
}
