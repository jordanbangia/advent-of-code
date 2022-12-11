from functools import cached_property


class DirNode:    
    def __init__(self, name, parent):
        self.name = name
        self.contents = []
        self.parent = parent
    
    def add_content(self, node):
        self.contents.append(node)

    @cached_property
    def size(self):
        return sum([s.size for s in self.contents])

    def dot_dot(self):
        return self.parent

    def __str__(self):
        return f'Dir {self.name}: {",".join(str(n) for n in self.contents)}'

    def find_dir(self, dir_name):
        return [s for s in self.contents if s.name == dir_name][0]

    
class FileNode:
    def __init__(self, name, size):
        self.name = name
        self.size = size
    
    def __str__(self):
        return self.name
    
    
def parse(file_name: str) -> list[str]:
    with open(file_name, "r") as f:
        return f.readlines()


directories = []

directories.append(DirNode("/", None))

active_dir = directories[0]
for line in parse("input.txt"):
    parts = line.strip().split(' ')
    if parts[0] == "$":
        if parts[1] == "cd":
            if parts[2] == "..":
                active_dir = active_dir.dot_dot()
            elif parts[2] == "/":
                active_dir = directories[0]
            else:
                # cd <dir name>
                active_dir = active_dir.find_dir(parts[2])
        elif parts[1] == "ls":
            # don't really need to handle anything, we just know
            # that everything after is going to be for the active dir
            pass
    elif parts[0] == "dir":
        node = DirNode(name=parts[1], parent=active_dir)
        # print("adding dir", node.name, "to", active_dir.name)
        active_dir.add_content(node)
        directories.append(node)
    else:
        node = FileNode(parts[1], int(parts[0]))
        # print("adding file", node.name, "to", active_dir.name)
        active_dir.add_content(node)


size = 0
for n in directories:
    if n.size <= 100000:
        size += n.size
print(size)

system_space = 70000000
required_space = 30000000

current_free = system_space - directories[0].size
required_free = required_space - current_free

node_to_delete = directories[0]
for n in directories:
    if n.size >= required_free:
        # this is a candidate for deletion
        if n.size < node_to_delete.size:
            node_to_delete = n

print(node_to_delete.size)
