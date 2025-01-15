export interface File {
    name: string;
    path: string;
    size: number;
    type: string;
    isDir: boolean;
}

export const getSizeInKB = (file: File): number => file.size / 1024;

export const getFileType = (name: String): string => {
    const extension = name.split(".").pop()
    return extension && extension !== name ? extension : ""
}

export const getIconClass = (): string => {
    return ""
}