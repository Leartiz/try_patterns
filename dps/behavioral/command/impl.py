from sys import stdout as console

# Обработка команды exit
class SessionClosed(Exception):
    def __init__(self, value):
        self.value = value

# Интерфейс команды
class Command:
    def execute(self):
        raise NotImplementedError()

    def cancel(self):
        raise NotImplementedError()		

    def name():
        raise NotImplementedError()

# Команда help
class HelpCommand(Command):
    def execute(self):
        text = "available commands:\n"
        for k in COMMANDS.keys():
            text += ('   ' + k + '\n')
        console.write(text)

    def cancel(self):
        console.write("You are canceled \"help\" command\n")

    def name(self):
        return "help"


# Команда rm
class RmCommand(Command):
    def execute(self):
        console.write("You are executed \"rm\" command\n")

    def cancel(self):
        console.write("You are canceled \"rm\" command\n")

    def name(self):
        return "rm"

# Команда uptime
class UptimeCommand(Command):
    def execute(self):
        console.write("You are executed \"uptime\" command\n")

    def cancel(self):
        console.write("You are canceled \"uptime\" command\n")
    
    def name(self):
        return "uptime"

# Команда undo
class UndoCommand(Command):
    def execute(self):
        try:
            cmd = HISTORY.pop()
            TRASH.append(cmd)
            console.write("Undo command \"{0}\"\n".format(cmd.name()))
            cmd.cancel()

        except IndexError:
            console.write("ERROR: HISTORY is empty\n")

    def name(self):
        return "undo"

# Команда redo
class RedoCommand(Command):
    def execute(self):
        try:
            cmd = TRASH.pop()
            HISTORY.append(cmd)
            console.write("Redo command \"{0}\"\n".format(cmd.name()))
            cmd.execute()
            
        except IndexError:
            console.write("ERROR: TRASH is empty\n")
            
    def name(self):
        return "redo"

# Команда history
class HistoryCommand(Command):
    def execute(self):
        i = 0
        for cmd in HISTORY:
            console.write("{0}: {1}\n".format(i, cmd.name()))
            i = i + 1
            
    def name(self):
        print("history")

# Команда trash
class TrashCommand(Command):
    def execute(self):
        i = 0
        for cmd in TRASH:
            console.write("{0}: {1}\n".format(i, cmd.name()))
            i = i + 1

    def name(self):
        print("trash")

# Команда exit
class ExitCommand(Command):
    def execute(self):
        raise SessionClosed("Good bay!")

    def name(self):
        return "exit"

# Словарь доступных команд
COMMANDS = {
    'help': HelpCommand(),
    'rm': RmCommand(),
    'uptime': UptimeCommand(),
    'undo': UndoCommand(), 
    'redo': RedoCommand(),
    'history': HistoryCommand(),
    'trash': TrashCommand(),
    'exit': ExitCommand()
}   

HISTORY = list()
TRASH = list()

# Шелл
def main():
    try:
        while True:
            console.flush()
            console.write("push >> ")

            cmd = input()

            try:
                command = COMMANDS[cmd]
                command.execute() 

                if (not isinstance(command, UndoCommand) and 
                    not isinstance(command, RedoCommand) and 
                    not isinstance(command, HistoryCommand) and
                    not isinstance(command, TrashCommand)
                    ):    

                    global TRASH 
                    TRASH = list()
                    HISTORY.append(command)

            except KeyError:
                console.write("ERROR: Command \"%s\" not found\n" % cmd)
                
    except SessionClosed as e:
        console.write(e.value)

if __name__ == "__main__":
    main()