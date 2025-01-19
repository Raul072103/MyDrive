import "./Topbar.css";
import {useLocation} from "react-router-dom";
import {useState} from "react";

const TopBar = ({ onUpload }: { onUpload: (file: File) => void }) => {
    const [_, setSelectedFile] = useState<File | null>(null);
    const location = useLocation();

    const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files?.[0]) {
            const file = event.target.files[0];
            setSelectedFile(file);
            onUpload(file);
        }
    };

    return (
        <div className="top-bar">
            <div className="logo">
                <img src="/assets/icons/logo.svg" alt={`Logo icon`} className="logo-icon"/>
                <span className="logo-text">MyDrive</span>
            </div>

            {/* File input and upload button */}
            {location.pathname === "/home/myfiles" && (
                 <div>
                    <input
                        type="file"
                        onChange={handleFileChange}
                        style={{ display: "none" }}
                        id="file-input"
                    />
                    <button className="upload-btn"
                        onClick={() => document.getElementById("file-input")?.click()}
                    >
                        Upload
                    </button>
                </div>
            )}
        </div>
    );
};

export default TopBar;