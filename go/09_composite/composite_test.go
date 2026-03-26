package composite

import "testing"

func TestComposite(t *testing.T) {
	
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	floder1 := &Floder{
		name: "Folder1",
	}

	floder1.add(file1)

	floder2 := &Floder{
		name: "Folder2",
	}
	floder2.add(file2)
	floder2.add(file3)
	floder2.add(floder1)

	floder2.search("rose")

}