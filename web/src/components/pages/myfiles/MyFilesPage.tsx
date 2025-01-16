import {File, getFileType} from "../../../types/File";
import {useEffect, useState} from "react";
import FileItem from "../../FileItem";
import {encodeStrBase64} from "../../../utils/FileUtils";

const FileExplorer = ({ setCurrentPath }: { setCurrentPath: (path: string) => void }) => {
    const [files, setFiles] = useState<File[]>([]);
    const [pathStack, setPathStack] = useState<string[]>(["root"]);
    const [loadingSuccessfully, setLoadingSuccessfully] = useState<boolean>(false);

    const currentPath = pathStack[pathStack.length - 1];

    const getFilesList = async (path: string) => {
        try {
            const encodedPath = encodeStrBase64(path);

            const response = await fetch(`http://localhost:8080/v1/myfiles/listfiles/${encodedPath}`, {
                method: 'GET',
            });

            if (response.ok) {
                const parsedData = await response.json();  // Handling response as binary data (Blob)

                if (parsedData && parsedData.data) {
                    const fileList: File[] = parsedData.data.map((item: {
                        name: string;
                        is_dir: boolean;
                        size: number
                    }) => ({
                        name: item.name,
                        size: item.size,
                        type: !item.is_dir ? getFileType(item.name) : "dir",
                        isDir: item.is_dir,
                        path: currentPath + "/" + item.name
                    }));
                    setFiles(fileList);
                } else {
                    setFiles([]);
                }
                setLoadingSuccessfully(true);

            } else {
                setLoadingSuccessfully(false);
                switch (response.status) {
                    case 404:
                        console.error('Not Found: The requested directory or file does not exist.');
                        break;
                    case 500:
                        console.error('Server Error: Something went wrong on the server.');
                        break;
                    case 400:
                        console.error('Bad Request: The request was malformed.');
                        break;
                    default:
                        console.error('Request failed with status:', response.status);
                        break;
                }

            }
        } catch (error) {
            console.error('Error during API request:', error);
        }
    };

    const handleDirectoryClick = (path: string) => {
        const updatedPath = `${currentPath}/${path}`;
        getFilesList(updatedPath);
        if (loadingSuccessfully) {
            setPathStack((prevStack) => [...prevStack, updatedPath]);
            setCurrentPath(updatedPath); // Update the current path in the parent component
        }
    };

    // Handle the back navigation
    const handlePopState = () => {
        if (pathStack.length > 1) {
            setPathStack((prevStack) => {
                const newStack = [...prevStack];
                newStack.pop();  // Pop the last path in the stack
                getFilesList(newStack[newStack.length - 1]);  // Fetch files for the previous path
                return newStack;
            });
        } else {
        }
    };

    useEffect(() => {
        window.addEventListener("popstate", handlePopState);
        getFilesList(currentPath);
        setCurrentPath(currentPath); // Update the current path in the parent component
        return () => {
            window.removeEventListener("popstate", handlePopState);
        };
    }, []);


    return (
            <div className="grid">
                {files.length === 0 ? (
                    // Show empty state if there are no files
                    <div className="empty-state">
                        <p>The directory is empty.</p>
                    </div>
                ) : (
                    files.map((file, _) => (
                        <FileItem
                            key={file.path}  // Added key prop for React to optimize list rendering
                            file={file}
                            onDirectoryClick={handleDirectoryClick}
                        />
                    ))
                )}
            </div>
    );
};

function MyFilesPage({ setCurrentPath }: { setCurrentPath: (path: string) => void }) {
    return (
        <div>
            <FileExplorer setCurrentPath={setCurrentPath} />
        </div>
    );
}

export default MyFilesPage;