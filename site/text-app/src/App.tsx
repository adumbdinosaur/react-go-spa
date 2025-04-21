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

  const { api } = useApi();

  const refreshFiles = async () => {
    try {
      const response = await api.userFilesGet();
      if (response.data.files) {
        setFiles(response.data.files);
      }
    } catch (error) {
      console.error("Error fetching files:", error);
    }
  };

  useEffect(() => {
    refreshFiles().catch(() => setShowAuthModal(true));
  }, [api]);

  const handleAuthSuccess = async () => {
    setShowAuthModal(false);
    await refreshFiles();
  };

  const handleLogout = async () => {
    try {
      await api.logoutPost();
      setShowAuthModal(true);
      setFiles([]);
      setSelectedFile(null);
    } catch (error) {
      console.error("Error during logout:", error);
    }
  };

  return (
    <div style={{ height: "100vh", display: "flex", flexDirection: "column" }}>
      {showAuthModal && <LoginRegister onAuthSuccess={handleAuthSuccess} refreshFiles={refreshFiles} />}
      <div style={{ display: "flex", flex: 1, overflow: "hidden" }}>
        <Sidebar setSelectedFile={setSelectedFile} files={files} onLogout={handleLogout} />
        <div className="main">
          <Display content={displayContent} />
          <CLI setDisplayContent={setDisplayContent} selectedFile={selectedFile} />
        </div>
      </div>
    </div>
  );
}