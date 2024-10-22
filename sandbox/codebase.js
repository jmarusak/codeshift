const codebase = {
  name: "root",
  type: "folder",
  children: [
    {
      name: "folder1",
      type: "folder",
      children: [
        { name: "file1.txt", type: "file" },
        { name: "file2.txt", type: "file" },
      ],
    },
    {
      name: "folder2",
      type: "folder",
      children: [
        { name: "file3.txt", type: "file" },
        {
          name: "subfolder1",
          type: "folder",
          children: [{ name: "file4.txt", type: "file" }],
        },
      ],
    },
  ],
};
