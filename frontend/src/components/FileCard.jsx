const FileCard = ({ file, isSelected, onToggleSelection, onDelete }) => {
	// Compute the border color
	let borderColor = "";
	if (isSelected) {
		borderColor = "border-3";
	} else if (file.developer) {
		borderColor = "border";
	}
	return (
		<div>
		<div className={`p-8 m-8 cursor-pointer ${borderColor}`}
		onClick={onToggleSelection}
		>
		<h3 className="text-lg font-bold">{file.name}</h3>
		<p className="text-gray-600">Dev: {file.developer}</p>
		</div>
		{/* Delete Button */}
		<button
		onClick={() => onDelete(file.name)}
		className="mt-2 bg-red-500 text-white px-4 py-2 rounded"
		>
		Delete
		</button>
		</div>
  );
};

export default FileCard;

