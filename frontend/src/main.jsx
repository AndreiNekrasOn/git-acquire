import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import FileList from "./components/FileList";
import DeveloperList from "./components/DeveloperList";
import './index.css'

const App = () => {
  return (
    <Router>
      <div className="p-4 bg-gray-100 min-h-screen">
        <nav className="mb-4">
          <Link to="/" className="mr-4 text-blue-500">Files</Link>
          <Link to="/developers" className="text-blue-500">Developers</Link>
        </nav>

        <Routes>
          <Route path="/" element={<FileList />} />
          <Route path="/developers" element={<DeveloperList />} />
        </Routes>
      </div>
    </Router>
  );
};

ReactDOM.createRoot(document.getElementById("root")).render(<App />);

