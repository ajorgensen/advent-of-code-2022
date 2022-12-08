package d02

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Cmd struct {
	Command string
	Output  []string
}

func (c *Cmd) Args() []string {
	args := strings.Split(c.Command, " ")
	if len(args) == 1 {
		return []string{}
	}

	return args[1:]
}

func (c *Cmd) AddOutput(output string) {
	c.Output = append(c.Output, output)
}

type File struct {
	Name     string
	Parent   *File
	Type     string
	Size     int
	Children []*File
}

func (f *File) TotalSize() int {
	total := f.Size

	for _, child := range f.Children {
		total += child.TotalSize()
	}

	return total
}

func (f *File) Cd(path string) (*File, error) {
	if path == ".." {
		return f.Parent, nil
	}

	for _, child := range f.Children {
		if child.Name == path {
			return child, nil
		}
	}

	return f, fmt.Errorf("NotFound %s in %s", path, f.Name)
}

func (f *File) AddChild(child *File) {
	child.Parent = f
	f.Children = append(f.Children, child)
}

func (f *File) String() string {
	out := fmt.Sprintf("- %s (%s", f.Name, f.Type)

	if f.Type == "file" {
		out = fmt.Sprintf("%s, size=%d", out, f.Size)
	}

	if f.Parent != nil {
		out = fmt.Sprintf("%s, parent=%s", out, f.Parent.Name)
	}

	return fmt.Sprintf("%s)", out)
}

func (f *File) Tree(depth int) string {
	indent := ""
	for i := 0; i < depth; i += 1 {
		indent += "."
	}

	lines := []string{
		fmt.Sprintf("%s%s", indent, f.String()),
	}

	for _, child := range f.Children {
		lines = append(lines, child.Tree(depth+1))
	}

	return strings.Join(lines, "\n")
}

func PartOne(r io.Reader) (int, error) {
	commands, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	fs, err := buildFilesystem(commands)
	if err != nil {
		return -1, err
	}

	var rec func(c *File) int
	rec = func(c *File) int {
		total := 0

		if c.Type == "dir" && c.TotalSize() <= 100000 {
			total += c.TotalSize()
		}

		for _, child := range c.Children {
			total += rec(child)
		}

		return total
	}

	total := rec(fs)

	return total, nil
}

func PartTwo(r io.Reader) (int, error) {
	commands, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	fs, err := buildFilesystem(commands)
	if err != nil {
		return -1, err
	}

	totalFsSpace := 70000000
	usedSpace := fs.TotalSize()
	need := 30000000

	type kv struct {
		Key   string
		Value int
	}

	var dirs []kv
	var rec func(*File)
	rec = func(f *File) {
		if f.Type == "dir" {
			dirs = append(dirs, kv{f.Name, f.TotalSize()})
		}

		for _, c := range f.Children {
			rec(c)
		}
	}

	rec(fs)

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Value < dirs[j].Value
	})

	for _, kv := range dirs {
		available := totalFsSpace - usedSpace + kv.Value
		if available >= need {
			return kv.Value, nil
		}
	}

	return -1, errors.New("something went wrong")
}

func buildFilesystem(commands []Cmd) (*File, error) {
	fs := &File{
		Name: "",
		Children: []*File{
			{
				Name:     "/",
				Type:     "dir",
				Children: make([]*File, 0),
			},
		},
	}

	curr := fs
	var err error
	for _, cmd := range commands {
		if strings.HasPrefix(cmd.Command, "cd") {
			path := cmd.Args()[0]

			curr, err = curr.Cd(path)
			if err != nil {
				return fs, err
			}
		}

		if strings.HasPrefix(cmd.Command, "ls") {
			for _, out := range cmd.Output {
				parts := strings.Split(out, " ")

				if strings.HasPrefix(out, "dir") {
					curr.AddChild(
						&File{
							Name:     parts[1],
							Type:     parts[0],
							Children: make([]*File, 0),
						},
					)
				} else {
					size, err := strconv.Atoi(parts[0])
					if err != nil {
						return fs, err
					}

					curr.AddChild(
						&File{
							Name:     parts[1],
							Type:     "file",
							Size:     size,
							Children: make([]*File, 0),
						},
					)
				}
			}
		}
	}

	return fs, nil
}

func parseInput(r io.Reader) ([]Cmd, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return []Cmd{}, err
	}

	lines := strings.Split(string(data), "\n")
	commands := make([]Cmd, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			commands = append(commands, Cmd{
				Command: line[2:],
				Output:  make([]string, 0),
			})
		} else {
			commands[len(commands)-1].AddOutput(line)
		}
	}

	return commands, nil
}
