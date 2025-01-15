import "./HomePage.css";
import "@styles/global.css";

import TopBar from "../../TopBar.tsx";
import Sidebar from "../../Sidebar.tsx";
import {Route, Routes} from "react-router-dom";
import MyFilesPage from "../myfiles/MyFilesPage.tsx";
import StarredPage from "../starred/StarredPage.tsx";
import TrashPage from "../trash/TrashPage.tsx";

const HomePageMainContent = () => {
    return (
        <div>
            This is home page
        </div>
    )
}


const HomePage = () => {
    return (
        <div className="home-page">
            <TopBar/>
            <div className="flex">
                <Sidebar/>
                <div className="main-content">
                    <Routes>
                        {/* Default route to show FileExplorer */}
                        <Route index element={<HomePageMainContent />} />

                        {/* Nested routes for specific content */}
                        <Route path="myfiles" element={<MyFilesPage />} />
                        <Route path="starred" element={<StarredPage />} />
                        <Route path="trash" element={<TrashPage />} />
                    </Routes>
                </div>

            </div>
        </div>
    );
};

export default HomePage;
