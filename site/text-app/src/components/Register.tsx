import { useState } from "react";
import { DefaultApi, DefaultApiRegisterPostRequest } from "../api/api";

const API = new DefaultApi();

interface RegisterProps {
    onRegisterSuccess: (token: string) => void;
}

export default function Register({ onRegisterSuccess }: RegisterProps) {
    const [username, setUsername] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [error, setError] = useState<string | null>(null);

    const handleRegister = async () => {
        try {
            const params: DefaultApiRegisterPostRequest = {
                registerPostRequest: { username, password },
            };
            const { data } = await API.registerPost(params);

            if (data.token) {
                localStorage.setItem("authToken", data.token);
                window.dispatchEvent(new Event("authTokenChanged"));
                onRegisterSuccess(data.token);
            } else {
                setError("Registration failed. Please try again.");
            }
        } catch (err) {
            console.error("Registration error:", err);
            setError("An error occurred during registration. Please try again.");
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
                onClick={handleRegister}
                style={{
                    backgroundColor: "#000",
                    border: "1px solid #00FF00",
                    color: "#00FF00",
                    padding: "0.5rem",
                    cursor: "pointer",
                    width: "100%",
                }}
            >
                Register
            </button>
        </div>
    );
}