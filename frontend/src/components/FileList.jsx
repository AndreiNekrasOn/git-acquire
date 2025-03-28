import { useEffect, useState } from "react";
import FileCard from "./FileCard";
import DeveloperList from "./DeveloperList"; // Import developer list

const FileList = () => {
  const [files, setFiles] = useState([]);

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

  const assignDeveloper = async (fileId, developerName) => {
    try {
      const response = await fetch(`http://localhost:8080/developers/${fileId}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ developer: developerName }),
      });

      if (!response.ok) throw new Error("Failed to update developer");

      // Refresh both lists after update
      fetchFiles();
    } catch (error) {
      console.error("Error assigning developer:", error);
    }
  };

  return (
    <div className="p-4 bg-gray-100 min-h-screen">
      <h1 className="text-2xl font-bold mb-4">Dev File Tracker</h1>
      <div className="flex">
        {/* File Cards */}
        <div className="flex flex-wrap">
          {files.length > 0 ? (
            files.map((file) => (
              <FileCard key={file.id} file={file} onAssign={assignDeveloper} />
            ))
          ) : (
            <p className="text-gray-500">No files found.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default FileList;

