import { useState } from "react";
import { DefaultApi, DefaultApiLoginPostRequest } from "../api/api";

const API = new DefaultApi();

interface LoginProps {
    onLoginSuccess: () => void;
}

export default function Login({ onLoginSuccess }: LoginProps) {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [error, setError] = useState<string | null>(null);

    const handleLogin = async () => {
        try {
            const params: DefaultApiLoginPostRequest = {
                loginPostRequest: { username, password },
            };
            const { data } = await API.loginPost(params);

            if (data.token) {
                localStorage.setItem("authToken", data.token);
                window.dispatchEvent(new Event("authTokenChanged"));
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