import React, { useState } from "react";
import "@styles/global.css";

import {File, getIconSrc, getSizeInDisplayFormat} from "../types/File.tsx";
import {encodeStrBase64} from "../utils/FileUtils.tsx";

interface FileItemProps {
    file: File;
    onDirectoryClick: (path: string) => void;
}

const FileItem: React.FC<FileItemProps> = ({file, onDirectoryClick}) => {
    const [fileData, setFileData] = useState<Blob | null>(null);
    const [isDownloaded, setIsDownloaded] = useState(false);

    const handleClick = async () => {
        if (!isDownloaded) {
            try {
                const encodedPath = encodeStrBase64(file.path);

                const response = await fetch(`http://localhost:8080/v1/myfiles/download/${encodedPath}`, {
                    method: 'GET',
                });

                console.log(`Download file at ${file.path}`);

                if (response.ok) {
                    const blob = await response.blob();  // Handling response as binary data (Blob)
                    setFileData(blob);  // Store the binary data in state
                    setIsDownloaded(true);
                } else {
                    console.error('Request failed:', response.statusText);
                }
            } catch (error) {
                console.error('Error during API request:', error);
            }
        }
    };

    const handleDownload = () => {
        if (fileData) {
            const link = document.createElement("a");
            link.href = URL.createObjectURL(fileData);
            link.download = file.name;
            link.click();
            URL.revokeObjectURL(link.href); // Clean up the object URL
        }
    };

    return (
        <div className="file-item" onClick={file.isDir ? () => onDirectoryClick(file.name) : undefined}>
            <img src={getIconSrc(file.type)} className={"file-item-icon"} alt={`${file.name}`}></img>

            <p className="file-name">{file.name.length > 20 ? `${file.name.slice(0, 10)}...`: file.name}</p>

            {/* Download Button: Top Right */}
            { !file.isDir &&
                <div
                    onClick={isDownloaded ? handleDownload : handleClick}
                    className="download-btn">
                    <img src={isDownloaded ? "/assets/icons/download-blue.svg" : "/assets/icons/download.svg"} alt="Download" style={{ width: '20px', height: '20px' }} />
                </div>
            }

            { !file.isDir &&
                <p className="file-size">{getSizeInDisplayFormat(file)}</p>
            }

        </div>
    );
};

export default FileItem;
