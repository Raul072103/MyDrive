import "./HomePage.css";
import "@styles/global.css";

import TopBar from "../../TopBar.tsx";
import Sidebar from "../../Sidebar.tsx";
import {Route, Routes} from "react-router-dom";
import MyFilesPage from "../myfiles/MyFilesPage.tsx";
import StarredPage from "../starred/StarredPage.tsx";
import TrashPage from "../trash/TrashPage.tsx";
import {useState} from "react";
import {encodeStrBase64} from "../../../utils/FileUtils.tsx";

const HomePageMainContent = () => {
    return (
        <div>
            This is home page
        </div>
    )
}

const HomePage = () => {
    const [currentPath, setCurrentPath] = useState<string>("root");

    const handleFileUpload = async (file: File) => {
        console.log("Upload at path ",currentPath, file.name);

        const filePath = encodeStrBase64(`${currentPath}/${file.name}`);

        const formData = new FormData();
        formData.append("myFile", file);

        try {
            const response = await fetch(`http://localhost:8080/v1/myfiles/upload?path=${filePath}&name=${file.name}&size=${file.size}`, {
                method: "POST",
                body: formData
            });

            if (response.ok) {
                console.log("Upload at path ",filePath);
            } else {
                console.error('Upload failed:', response.status);
            }
        } catch (error) {
            console.error('Error during file upload:', error);
        }
    };

    return (
        <div className="home-page">
            <TopBar onUpload={handleFileUpload}/>
                <div className="flex-container">
                    <Sidebar/>
                    <div className="content-placer">
                        <Routes>
                            {/* Default route to show FileExplorer */}
                            <Route index element={<HomePageMainContent />} />

                            {/* Nested routes for specific content */}
                            <Route path="myfiles" element={<MyFilesPage setCurrentPath={setCurrentPath} />} />
                            <Route path="starred" element={<StarredPage />} />
                            <Route path="trash" element={<TrashPage />} />
                        </Routes>
                    </div>
                </div>
        </div>
    );
};

export default HomePage;
