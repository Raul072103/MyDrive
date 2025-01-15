import React, { useState } from "react";
import "@styles/global.css";

import {File, getIconClass} from "../types/File.tsx";

interface FileItemProps {
    file: File; // Define the expected type of the prop
}

const FileItem: React.FC<FileItemProps> = ({file}) => {
    const [fileData, setFileData] = useState<Blob | null>(null);

    const handleClick = async () => {
        console.log("CLICKED ME!")
        try {
            const response = await fetch(`http://localhost:8080/v1/myfiles/download/${file.path}`, {
                method: 'GET',
            });

            if (response.ok) {
                const blob = await response.blob();  // Handling response as binary data (Blob)
                setFileData(blob);  // Store the binary data in state
            } else {
                console.error('Request failed:', response.statusText);
            }
        } catch (error) {
            console.error('Error during API request:', error);
        }
    };

    return (
        <div className="file-item" onClick={handleClick}>
            <i className={getIconClass()}></i>
            <p>{file.name}</p>

            {fileData && (
                <div className="file-data">
                    <a href={URL.createObjectURL(fileData)} download={file.name}>Download</a>
                </div>
            )}
        </div>
    );
};

export default FileItem;
