import React from "react";

interface DisplayProps {
    content: string;
}

const Display: React.FC<DisplayProps> = ({ content }) => {
    return (
        <div className="display" style={{ flex: 1, padding: "10px", color: "#00FF00", backgroundColor: "#000" }}>
            <pre>{content}</pre>
        </div>
    );
};

export default Display;