
// 读取file中的内容(图片)
export const getFileImageContent = async (file) => {
    let rd = new FileReader()
    rd.readAsDataURL(file)

    return await new Promise((resolve, reject) => {
        rd.onloadend = () => resolve(rd.result);
        rd.onerror = () => reject(new Error("Failed to read file."));
        rd.onabort = () => reject(new Error("File reading was aborted."));
    });
}

// 读取file中的内容(文本)
export const getFileContent = async (file) => {
    let rd = new FileReader()
    rd.readAsText(file);

    return await new Promise((resolve, reject) => {
        rd.onloadend = (e) => resolve(e.target.result);
        rd.onerror = () => reject(new Error("Failed to read file."));
        rd.onabort = () => reject(new Error("File reading was aborted."));
    });
}

// 将base64的内容转换为url类型
export const base64ToImageUrl = (base64) => {
    const base64Data = base64.split(',')[1];
    // 将 base64 编码的图像数据转换为 Blob 对象
    var byteCharacters = atob(base64Data);
    var byteNumbers = new Array(byteCharacters.length);
    for (var i = 0; i < byteCharacters.length; i++) {
        byteNumbers[i] = byteCharacters.charCodeAt(i);
    }
    var byteArray = new Uint8Array(byteNumbers);
    var blob = new Blob([byteArray], { type: 'image/png' });

    // 创建 URL
    return URL.createObjectURL(blob);
}
// 获取content中的所有image
export const getContentImage = (content) => {
    if (!content) return
    let cImages = content.match(/!\[.*?\]\(.*?\)/g)
    return cImages
}