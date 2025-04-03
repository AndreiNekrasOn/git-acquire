import { useEffect, useState } from "react";

const LoginForm = () => {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");

	const handleSubmit = async () => {
		localStorage.setItem('username', username)
		localStorage.setItem('password', password)
		let authData = window.btoa(username + ':' + password)
		try {
			const response = await fetch("http://localhost:8080/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
					"Authorization": 'Basic ' + authData,
				},
				body: ""
			});
			if (!response.ok) throw new Error("Failed to sign in");
		} catch (error) {
			console.error("Error signing in:", error);
		}
	}

	return <div>
		<h1 className="text-2xl font-bold mb-4">Files</h1>
		<div className="mb-4 p-4 bg-white shadow rounded">
		<h2 className="text-lg font-bold mb-2">Sign In</h2>
		<form onSubmit={handleSubmit}>
		<label htmlFor="username">Username:</label>
		<input type="text" id="username" name="username" value={username} required
		onChange={(e) => setUsername(e.target.value)}/>
		<br />
		<label htmlFor="password">Password:</label>
		<input type="password" id="password" name="password"  value={password} required
		onChange={(e) => setPassword(e.target.value)}/>
		<br />
		<button type="submit">Login</button>
		</form>
		</div>
		</div>
}


export default LoginForm;
