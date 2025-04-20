import { useState, useEffect } from "react";
import CLI from "./components/CLI";
import Sidebar from "./components/Sidebar";
import Display from "./components/Display";
import { ApiProvider, useApi } from "./context/ApiContext";
import LoginRegister from "./components/LoginRegister";
import "./App.css";

export default function App() {
  return (
    <ApiProvider>
      <MainApp />
    </ApiProvider>
  );
}

function MainApp() {
  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [displayContent, setDisplayContent] = useState<string>("");
  const [files, setFiles] = useState<string[]>([]);
  const [showAuthModal, setShowAuthModal] = useState<boolean>(false);

  const { api, setToken } = useApi();

  useEffect(() => {
    const authToken = localStorage.getItem("authToken");
    if (!authToken) {
      setShowAuthModal(true);
    } else {
      setToken(authToken);
    }
  }, [setToken]);

  useEffect(() => {
    const fetchFiles = async () => {
      try {
        const response = await api.userFilesGet();
        if (response.data.files) {
          setFiles(response.data.files);
        }
      } catch (error) {
        console.error("Error fetching files:", error);
      }
    };

    fetchFiles();
  }, [api]);

  const handleAuthSuccess = (token: string) => {
    localStorage.setItem("authToken", token);
    setToken(token);
    setShowAuthModal(false);
  };

  return (
    <div style={{ height: "100vh", display: "flex", flexDirection: "column" }}>
      {showAuthModal && (
        <LoginRegister onAuthSuccess={handleAuthSuccess} />
      )}
      <div style={{ display: "flex", flex: 1, overflow: "hidden" }}>
        <Sidebar setSelectedFile={setSelectedFile} files={files} />
        <div className="main">
          <Display content={displayContent} />
          <CLI setDisplayContent={setDisplayContent} selectedFile={selectedFile} />
        </div>
      </div>
    </div>
  );
}