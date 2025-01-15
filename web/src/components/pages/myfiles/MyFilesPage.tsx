import FileItem from "../../FileItem.tsx";

const FileExplorer = () => {
    const files = [
        { name: "Documents", iconClass: "", path: "large_file_1.txt" },
        { name: "Report.pdf", iconClass: "", path: "report.pdf"},
        { name: "Image.png", iconClass: "", path: "Image.png"},
    ];

    return (
        <div className="main-content">
            <div className="grid">
                {files.map((file, index) => (
                    <FileItem
                        key={index}
                        name={file.name}
                        iconClass={file.iconClass}
                        path={file.path} />
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