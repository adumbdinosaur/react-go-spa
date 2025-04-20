import { useState, useEffect } from "react";
import { FileUpload } from "./FileUpload";

interface SidebarProps {
    setSelectedFile: (file: string) => void;
    files: string[];
}

export default function Sidebar({ setSelectedFile, files }: SidebarProps) {
    const [items, setItems] = useState<string[]>([]);
    const [selectedIndex, setSelectedIndex] = useState<number | null>(null);
    const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

    useEffect(() => {
        setItems(files);


        const authToken = localStorage.getItem("authToken");
        setIsLoggedIn(!!authToken);
    }, [files]);

    useEffect(() => {
        const updateLoginState = () => {
            const authToken = localStorage.getItem("authToken");
            setIsLoggedIn(!!authToken);
        };

        const handleAuthTokenChange = () => {
            updateLoginState();
        };

        window.addEventListener("authTokenChanged", handleAuthTokenChange);


        return () => {
            window.removeEventListener("authTokenChanged", handleAuthTokenChange);
        };
    }, []);

    const handleFileUpload = (fileName: string) => {
        setItems((prev) => [...prev, fileName]);
    };

    const handleLogout = () => {
        localStorage.removeItem("authToken");
        setIsLoggedIn(false);
        setSelectedFile("");
        window.dispatchEvent(new Event("authTokenChanged"));
        window.location.reload();
    };

    return (
        <div
            className="sidebar"
            style={{
                display: "flex",
                flexDirection: "column",
                height: "100vh",
                overflow: "hidden",
            }}
        >
            <div style={{ flex: 1, overflowY: "auto", padding: "1rem" }}>
                <h3 style={{ color: "#00FF00" }}>Files</h3>
                <ul style={{ listStyle: "none", padding: 0 }}>
                    {items.map((file, index) => (
                        <li
                            key={index}
                            style={{
                                padding: "0.5rem",
                                cursor: "pointer",
                                color: selectedIndex === index ? "#000" : "#00FF00",
                                backgroundColor: selectedIndex === index ? "#00FF00" : "transparent",
                                borderRadius: "4px",
                                marginBottom: "0.5rem",
                            }}
                            onClick={() => {
                                setSelectedIndex(index);
                                setSelectedFile(file);
                            }}
                        >
                            {file}
                        </li>
                    ))}
                </ul>
            </div>

            <FileUpload onFileUpload={handleFileUpload} />

            <div style={{ padding: "1rem", borderTop: "1px solid #00FF00" }}>
                {isLoggedIn && (
                    <button
                        onClick={handleLogout}
                        style={{
                            backgroundColor: "#000",
                            border: "1px solid #00FF00",
                            color: "#00FF00",
                            padding: "0.5rem",
                            width: "100%",
                            fontFamily: "Courier New, monospace",
                            cursor: "pointer",
                        }}
                    >
                        Logout
                    </button>
                )}
            </div>
        </div>
    );
}