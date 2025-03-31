import { useEffect, useState } from "react";

const DeveloperList = () => {
  const [developers, setDevelopers] = useState([]);

  useEffect(() => {
    fetchDevelopers();
  }, []);

  const fetchDevelopers = async () => {
    try {
      const response = await fetch("http://localhost:8080/developers");
      if (!response.ok) throw new Error("Failed to fetch developers");
      const data = await response.json();
      setDevelopers(data);
    } catch (error) {
      console.error("Error fetching developers:", error);
    }
  };

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Developers</h1>
      {developers.length === 0 ? (
        <p>No developers assigned yet.</p>
      ) : (
        developers.map((dev) => (
          <div key={dev.name} className="bg-white p-4 shadow rounded my-2 border">
            <h3 className="text-lg font-bold">{dev.name}</h3>
            <p className="text-gray-600">
              <strong>Files:</strong> {dev.files?.length > 0 ? dev.files.join(", ") : "None"}
            </p>
          </div>
        ))
      )}
    </div>
  );
};

export default DeveloperList;

