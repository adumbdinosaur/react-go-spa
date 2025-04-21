import { useState } from "react";
import Login from "./Login";
import Register from "./Register";

interface LoginRegisterProps {
    onAuthSuccess: () => void;
    refreshFiles: () => Promise<void>;
}

export default function LoginRegister({ onAuthSuccess, refreshFiles }: LoginRegisterProps) {
    const [isLogin, setIsLogin] = useState(true);

    const handleLoginSuccess = () => {
        onAuthSuccess();
    };

    const handleRegisterSuccess = () => {
        onAuthSuccess();
    };

    return (
        <div
            style={{
                position: "fixed",
                top: 0,
                left: 0,
                width: "100%",
                height: "100%",
                backgroundColor: "rgba(0, 0, 0, 0.8)",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                zIndex: 1000,
            }}
        >
            <div
                style={{
                    backgroundColor: "#000",
                    border: "1px solid #00FF00",
                    padding: "2rem",
                    borderRadius: "8px",
                    width: "300px",
                    color: "#00FF00",
                    fontFamily: "Courier New, monospace",
                }}
            >
                {isLogin ? (
                    <>
                        <h2 style={{ textAlign: "center" }}>Login</h2>
                        <Login onLoginSuccess={handleLoginSuccess} refreshFiles={refreshFiles} />
                    </>
                ) : (
                    <>
                        <h2 style={{ textAlign: "center" }}>Register</h2>
                        <Register onRegisterSuccess={handleRegisterSuccess} />
                    </>
                )}
                <button
                    onClick={() => setIsLogin(!isLogin)}
                    style={{
                        marginTop: "1rem",
                        backgroundColor: "#000",
                        border: "1px solid #00FF00",
                        color: "#00FF00",
                        padding: "0.5rem",
                        cursor: "pointer",
                        width: "100%",
                    }}
                >
                    {isLogin ? "Switch to Register" : "Switch to Login"}
                </button>
            </div>
        </div>
    );
}