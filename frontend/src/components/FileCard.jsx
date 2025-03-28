import { useState } from "react";

const FileCard = ({ file, onAssign }) => {
	const [devName, setDevName] = useState(file.developer || "");

	const handleAssign = () => {
		if (!devName.trim()) return;
		onAssign(file.id, devName);
		setDevName(""); // Clear input
	};

	return (
		<div className="bg-white shadow-lg rounded-lg p-4 m-2 w-60">
		<h3 className="text-lg font-bold">{file.name}</h3>
		<p className="text-sm text-gray-500">Branch: {file.branch}</p>
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
		</div>
	);
};

export default FileCard;

