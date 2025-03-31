import { useEffect, useState } from "react";
import FileCard from "./FileCard";

const FileList = () => {
  const [files, setFiles] = useState([]);
  const [selectedFiles, setSelectedFiles] = useState([]); // Track selected files
  const [developer, setDeveloper] = useState("");
  const [fileName, setFileName] = useState("");

  useEffect(() => {
    fetchFiles();
  }, []);

  const fetchFiles = async () => {
    try {
      const response = await fetch("http://localhost:8080/files");
      if (!response.ok) throw new Error("Failed to fetch files");
      const data = await response.json();
      setFiles(data);
    } catch (error) {
      console.error("Error fetching files:", error);
    }
  };

  const addFile = async () => {
    if (!fileName.trim()) return;

    try {
      const response = await fetch("http://localhost:8080/files", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name: fileName }),
      });

      if (!response.ok) throw new Error("Failed to add file");

      setFileName("");
      fetchFiles(); // Refresh file list
    } catch (error) {
      console.error("Error adding file:", error);
    }
  };

  const deleteFile = async (fileId) => {
    try {
      const response = await fetch(`http://localhost:8080/files/${fileId}`, {
        method: "DELETE",
      });

      if (!response.ok) throw new Error("Failed to delete file");

      fetchFiles(); // Refresh after delete
    } catch (error) {
      console.error("Error deleting file:", error);
    }
  };

  const toggleSelection = (fileId) => {
    setSelectedFiles((prevSelected) =>
      prevSelected.includes(fileId)
        ? prevSelected.filter((id) => id !== fileId) // Unselect if already selected
        : [...prevSelected, fileId] // Select if not selected
    );
  };

  const assignFiles = async () => {
    if (selectedFiles.length === 0) return;
    try {
      const response = await fetch("http://localhost:8080/assign", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ developer, fileIds: selectedFiles }),
      });
      if (!response.ok) throw new Error("Failed to assign files");
      setDeveloper("");
      setSelectedFiles([]); // Clear selection
      fetchFiles(); // Refresh file list
    } catch (error) {
      console.error("Error assigning files:", error);
    }
  };

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Files</h1>

      {/* Add File Form */}
      <div className="mb-4 p-4 bg-white shadow rounded">
        <h2 className="text-lg font-bold mb-2">Add a New File</h2>
        <input
          type="text"
          placeholder="File name"
          value={fileName}
          onChange={(e) => setFileName(e.target.value)}
          className="p-2 border rounded mr-2"
        />
        <button
          onClick={addFile}
          className="bg-blue-500 text-white px-4 py-2 rounded"
        >
          Add File
        </button>
      </div>

      {/* Assign Developer Form */}
      <div className="mb-4 p-4 bg-white shadow rounded">
        <h2 className="text-lg font-bold mb-2">Assign Files</h2>
        <input
          type="text"
          placeholder="Developer name"
          value={developer}
          onChange={(e) => setDeveloper(e.target.value)}
          className="p-2 border rounded mr-2"
        />
        <button
          onClick={assignFiles}
          className="bg-green-500 text-white px-4 py-2 rounded"
        >
          Assign Selected Files
        </button>
      </div>

      {/* File List */}
      <div className="flex flex-wrap">
        {files.length > 0 ? (
          files.map((file) => (
            <FileCard
              key={file.id}
              file={file}
              isSelected={selectedFiles.includes(file.id)}
              onToggleSelection={() => toggleSelection(file.id)}
              onDelete={() => deleteFile(file.id)}
            />
          ))
        ) : (
          <p className="text-gray-500">No files found.</p>
        )}
      </div>
    </div>
  );
};

export default FileList;

