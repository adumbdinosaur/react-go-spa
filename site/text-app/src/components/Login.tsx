import { useState } from "react";
import { DefaultApiLoginPostRequest } from "../api/api";
import { useApi } from "../context/ApiContext";

interface LoginProps {
    onLoginSuccess: () => void;
    refreshFiles: () => Promise<void>;
}

export default function Login({ onLoginSuccess, refreshFiles }: LoginProps) {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [error, setError] = useState<string | null>(null);
    const { api } = useApi();

    const handleLogin = async () => {
        try {
            const params: DefaultApiLoginPostRequest = {
                loginPostRequest: { username, password },
            };
            const response = await api.loginPost(params);

            if (response.status === 200) {
                await refreshFiles(); // Fetch and update the file list
                onLoginSuccess();
            } else {
                setError("Login failed. Please check your credentials.");
            }
        } catch (err) {
            console.error("Login error:", err);
            setError("An error occurred during login. Please try again.");
        }
    };

    return (
        <div>
            {error && <p style={{ color: "red", textAlign: "center" }}>{error}</p>}
            <div style={{ marginBottom: "1rem" }}>
                <label>Username:</label>
                <input
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    style={{
                        width: "100%",
                        padding: "0.5rem",
                        marginTop: "0.5rem",
                        backgroundColor: "#000",
                        border: "1px solid #00FF00",
                        color: "#00FF00",
                    }}
                />
            </div>
            <div style={{ marginBottom: "1rem" }}>
                <label>Password:</label>
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    style={{
                        width: "100%",
                        padding: "0.5rem",
                        marginTop: "0.5rem",
                        backgroundColor: "#000",
                        border: "1px solid #00FF00",
                        color: "#00FF00",
                    }}
                />
            </div>
            <button
                onClick={handleLogin}
                style={{
                    backgroundColor: "#000",
                    border: "1px solid #00FF00",
                    color: "#00FF00",
                    padding: "0.5rem",
                    cursor: "pointer",
                    width: "100%",
                }}
            >
                Login
            </button>
        </div>
    );
}