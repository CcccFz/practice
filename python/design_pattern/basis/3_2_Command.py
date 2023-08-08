# -- coding: utf-8 --

import os

class MoveFileCommand(object):
    def __init__(self, src, dest):
        self.src = src
        self.dest = dest

    def __call__(self):
        print('renaming {} to {}'.format(self.src, self.dest))
        os.rename(self.src, self.dest)

    def execute(self):
        self()

    def undo(self):
        print('renaming {} to {}'.format(self.dest, self.src))
        os.rename(self.dest, self.src)


if __name__ == '__main__':
    command_stack = []

    command_stack.append(MoveFileCommand('foo.txt', 'bar.txt'))
    command_stack.append(MoveFileCommand('bar.txt', 'baz.txt'))

    try:
        with open("foo.txt", "w"):
            pass

        for cmd in command_stack:
            cmd.execute()

        for cmd in reversed(command_stack):
            cmd.undo()
    finally:
        os.unlink("foo.txt")
