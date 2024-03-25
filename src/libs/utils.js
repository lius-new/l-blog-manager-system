
export const getFileContent = async (file) => {
    let rd = new FileReader()
    rd.readAsDataURL(file)

    return await new Promise((resolve, reject) => {
        rd.onloadend = () => resolve(rd.result);
        rd.onerror = () => reject(new Error("Failed to read file."));
        rd.onabort = () => reject(new Error("File reading was aborted."));
    });
}