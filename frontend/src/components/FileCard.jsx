import { useState } from "react";

const FileCard = ({ file, onAssign, onDelete }) => {
	const [devName, setDevName] = useState(file.developer || "");

	const handleAssign = () => {
		if (!devName.trim()) return;
		onAssign(file.id, devName);
		// setDevName(""); // Clear input
	};

	return (
		<div className="bg-white p-4 shadow rounded m-2">
		<h3 className="text-lg font-bold">{file.name}</h3>
		<p className="text-gray-600">Branch: {file.branch}</p>
		<p className="text-gray-600">Dev: {file.developer}</p>
		{/* Assign button */}
		<input
		type="text"
		className="border p-1 w-full mt-2"
		placeholder="Enter your name"
		value={devName}
		onChange={(e) => setDevName(e.target.value)}
		/>
		<button
		onClick={handleAssign}
		className="bg-blue-500 text-white px-4 py-1 rounded mt-2 w-full"
		>
		Assign
		</button>

		{/* Delete Button */}
		<button
		onClick={() => onDelete(file.id)}
		className="mt-2 bg-red-500 text-white px-4 py-2 rounded"
		>
		Delete
		</button>
		</div>
	);
};

export default FileCard;

