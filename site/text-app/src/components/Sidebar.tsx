import { useState, useEffect } from "react";
import { FileUpload } from "./FileUpload";

interface SidebarProps {
    setSelectedFile: (file: string) => void;
    files: string[];
    onLogout: () => void;
}

export default function Sidebar({ setSelectedFile, files, onLogout }: SidebarProps) {
    const [items, setItems] = useState<string[]>([]);
    const [selectedIndex, setSelectedIndex] = useState<number | null>(null);

    useEffect(() => {
        setItems(files);
    }, [files]);

    const handleFileUpload = (fileName: string) => {
        setItems((prev) => [...prev, fileName]);
    };

    return (
        <div
            className="sidebar"
            style={{
                display: "flex",
                flexDirection: "column",
                height: "100vh",
                overflow: "hidden",
                backgroundColor: "#000",
                borderRight: "1px solid #00FF00",
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
                <button
                    onClick={onLogout}
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
            </div>
        </div>
    );
}