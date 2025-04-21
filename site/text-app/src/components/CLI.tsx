import React, { useState, useEffect, useRef } from "react";
import { DefaultApiQueryPostRequest } from "../api/api";
import { useApi } from "../context/ApiContext";

interface CLIProps {
    setDisplayContent: React.Dispatch<React.SetStateAction<string>>;
    selectedFile: string | null;
}

export default function CLI({ setDisplayContent, selectedFile }: CLIProps) {
    const [input, setInput] = useState<string>("");
    const typingTimeoutRef = useRef<NodeJS.Timeout | null>(null);
    const { api } = useApi();

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInput(e.target.value);

        if (typingTimeoutRef.current) {
            clearTimeout(typingTimeoutRef.current);
        }

        typingTimeoutRef.current = setTimeout(() => {
            if (e.target.value.trim() !== "" && selectedFile) {
                sendQuery(e.target.value);
            } else if (!selectedFile) {
                setDisplayContent("No file selected. Please select a file from the sidebar.");
            }
        }, 500);
    };

    const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === "Enter" && input.trim() !== "") {
            if (selectedFile) {
                sendQuery(input);
            } else {
                setDisplayContent("No file selected. Please select a file from the sidebar.");
            }
        }
    };

    const sendQuery = async (query: string) => {
        try {
            const params: DefaultApiQueryPostRequest = {
                queryPostRequest: { query, fileName: selectedFile! },
            };

            const response = await api.queryPost(params);

            setDisplayContent(response.data.results?.join("\n") || "No results found.");
        } catch (error) {
            console.error("Error sending request:", error);
            setDisplayContent("Error sending request. Please try again.");
        }
    };

    useEffect(() => {
        return () => {
            if (typingTimeoutRef.current) {
                clearTimeout(typingTimeoutRef.current);
            }
        };
    }, []);

    return (
        <div className="cli-container" style={{ padding: "10px", backgroundColor: "#000", color: "#00FF00" }}>
            <input
                type="text"
                value={input}
                onChange={handleInputChange}
                onKeyDown={handleKeyDown}
                placeholder="Type some text..."
                style={{
                    width: "100%",
                    padding: "10px",
                    color: "#00FF00",
                    backgroundColor: "#000",
                    border: "1px solid #00FF00",
                    fontFamily: "Courier New, monospace",
                }}
            />
        </div>
    );
}