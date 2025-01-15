import {File, getFileType} from "../../../types/File";
import {useEffect, useState} from "react";
import FileItem from "../../FileItem";

const FileExplorer = () => {
    const [files, setFiles] = useState<File[]>([]);
    const [pathStack, _] = useState<string[]>(["root"]);

    const currentPath = pathStack[pathStack.length - 1];

    const getFilesList = async (path: string) => {
        try {
            const response = await fetch(`http://localhost:8080/v1/myfiles/listfiles/${path}`, {
                method: 'GET',
            });

            if (response.ok) {
                const parsedData = await response.json();  // Handling response as binary data (Blob)

                const fileList: File[] = parsedData.data.map((item: { name: string; is_dir: boolean; size: number }) => ({
                    name: item.name,
                    size: item.size,
                    type: !item.is_dir ? getFileType(item.name) : "dir",
                    isDir: item.is_dir,
                    path: item.name
                }));

                setFiles(fileList);
            } else {
                console.error('Request failed:', response.statusText);
            }
        } catch (error) {
            console.error('Error during API request:', error);
        }
    };

    useEffect(() => {
        getFilesList(currentPath); // Fetch files when the component mounts
    }, []); // Empty dependency array ensures this runs only once


    return (
        <div className="main-content">
            <div className="grid">
                {files.map((file, _) => (
                    <FileItem
                        file={file}/>
                ))}
            </div>
        </div>
    );
};

function MyFilesPage() {
    return (
        <div>
            <FileExplorer/>
        </div>
    );
}

export default MyFilesPage;