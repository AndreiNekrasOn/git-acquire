import { useEffect, useState } from "react";
import FileCard from "./FileCard";

const FileList = () => {
  const [files, setFiles] = useState([]);
  const [fileName, setFileName] = useState("");
  const [branch, setBranch] = useState("");

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
    if (!fileName.trim() || !branch.trim()) return;

    try {
      const response = await fetch("http://localhost:8080/files", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name: fileName, branch }),
      });

      if (!response.ok) throw new Error("Failed to add file");

      setFileName("");
      setBranch("");
      fetchFiles(); // Refresh file list
    } catch (error) {
      console.error("Error adding file:", error);
    }
  };

 const assignFile = async (fileId, devName) => {
   try {
     const response = await fetch(`http://localhost:8080/files/${fileId}`, {
       method: "PUT",
       headers: {
         "Content-Type": "application/json",
       },
       body: JSON.stringify({ developer: devName }),
     });
   } catch (error) {
     console.error("Error adding file:", error);
   }
 }

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
        <input
          type="text"
          placeholder="Branch"
          value={branch}
          onChange={(e) => setBranch(e.target.value)}
          className="p-2 border rounded mr-2"
        />
        <button
          onClick={addFile}
          className="bg-blue-500 text-white px-4 py-2 rounded"
        >
          Add File
        </button>
      </div>

      {/* File List */}
      <div className="flex flex-wrap">
        {files.length > 0 ? (
          files.map((file) => (
            <FileCard key={file.id} file={file} onAssign={assignFile} onDelete={deleteFile} />
          ))
        ) : (
          <p className="text-gray-500">No files found.</p>
        )}
      </div>
    </div>
  );
};

export default FileList;

