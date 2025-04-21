import React, { useRef } from "react";
import { useApi } from "../context/ApiContext";

interface FileUploadProps {
    onFileUpload: (fileName: string) => void;
}

export const FileUpload: React.FC<FileUploadProps> = ({ onFileUpload }) => {
    const fileInputRef = useRef<HTMLInputElement>(null);
    const { api } = useApi();

    const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (file) {
            try {
                const formData = new FormData();
                formData.append("file", file);
                const response = await api.uploadPost(
                    { file },
                );

                if (response.data.message) {
                    onFileUpload(file.name);
                } else {
                    alert("File upload failed. Please try again.");
                }
            } catch (error) {
                console.error("Error uploading file:", error);
                alert("An error occurred while uploading the file.");
            }
        }
    };

    const handleFileUploadClick = () => {
        fileInputRef.current?.click();
    };

    return (
        <div style={{ padding: "1rem", borderTop: "1px solid #00FF00" }}>
            <input
                type="file"
                ref={fileInputRef}
                style={{ display: "none" }}
                onChange={handleFileChange}
            />
            <button
                onClick={handleFileUploadClick}
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
                Upload File
            </button>
        </div>
    );
};