export interface File {
    name: string;
    path: string;
    size: number;
    type: string;
    isDir: boolean;
}

export const getSizeInDisplayFormat = (file: File): string => {
    const sizeInKB = file.size / 1024;

    if (sizeInKB < 1024) {
        return `${sizeInKB.toFixed(2)} KB`;  // Display KB if less than 1MB
    }

    const sizeInMB = sizeInKB / 1024;
    if (sizeInMB < 1024) {
        return `${sizeInMB.toFixed(2)} MB`;  // Display MB if less than 1GB
    }

    const sizeInGB = sizeInMB / 1024;
    if (sizeInGB < 1024) {
        return `${sizeInGB.toFixed(2)} GB`;  // Display GB if less than 1TB
    }

    const sizeInTB = sizeInGB / 1024;
    return `${sizeInTB.toFixed(2)} TB`;  // Display TB for sizes larger than 1TB
}

export const getFileType = (name: String): string => {
    const extension = name.split(".").pop()
    return extension && extension !== name ? extension : ""
}

export const getIconSrc = (type: string): string => {
    switch (type) {
        case "txt":
            return "/assets/icons/file-txt.svg";
        case "png":
            return "/assets/icons/file-img.svg";
        case "jpg":
            return "/assets/icons/file-img.svg";
        case "jpeg":
            return "/assets/icons/file-img.svg";
        case "dir":
            return "/assets/icons/file-doc.svg";
        default:
            return "/assets/icons/uknown-file.svg";
    }
}