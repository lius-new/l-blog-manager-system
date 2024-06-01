import axios from 'axios'

const SERVER_USER_URL = import.meta.env.VITE_USER_API_SERVER_URI
const SERVER_CONTENT_URL = import.meta.env.VITE_CONTENT_API_SERVER_URI


function buildRequest(url, interceptorsRequest) {
    if (typeof url != "string" && typeof interceptorsRequest != "function") {
        console.error("url & interceptorsRequest param must type")
        return
    }
    const request = axios.create({
        withCredentials: true,
        baseURL: url,
    })
    request.interceptors.request.use(function (config) {
        interceptorsRequest(config)
        return config
    }, function (error) {
        return Promise.reject(error)
    })
    request.interceptors.response.use(function (response) {
        return response.data
    }, function (error) {
        return Promise.reject(error)
    })
    return request
}

const userRequest = buildRequest(SERVER_USER_URL, function (config) {
    if (config.url == "/users/auth") {
        let token = localStorage.getItem("token")
        config.headers.Authorization = "Bearer " + token
    }
})
const contentRequest = buildRequest(SERVER_CONTENT_URL, function (config) {
    if (config.url.indexOf("backend") != -1) {
        let token = localStorage.getItem("token")
        config.headers.Authorization = "Bearer " + token
    }
})


// 登录
export const login = async (username, password) => {
    return await userRequest.post('/users/login', { username, password })
}
// 校验
export const auth = async () => {
    return await userRequest.get('/users/auth',)
}
// 查看文章列表
export const articlesViews = async (page_size, page_num) => {
    return await contentRequest.post('/articles/backend', { page_size, page_num })
}
// 查看文章
export const articlesView = async (id) => {
    return await contentRequest.post('/articles/backend/id', { id })
}
// 查看图片
export const coverView = async (hash) => {
    return await contentRequest.get(`/articles/image/${hash}`)
}
// 查看图片
export const coverViewBase64 = async (hash) => {
    return await contentRequest.get(`/articles/image/base64/${hash}`)
}
// 修改文件状态
export const articleModifyVisiable = async (id, visiable) => {
    return await contentRequest.post('/articles/backend/modify/visiable', { id, visiable })
}

// 创建文章
export const articleSave = async (title, content, description, tags, covers) => {
    return await contentRequest.post('/articles/backend/create', {
        title, content, description, tags, covers
    })
}



export const articleModify = async (id, title, content, description, tags, covers, visiable) => {
    return await contentRequest.post('/articles/backend/modify', { id, title, content, description, tags, covers, visiable })
}

export const uploadArticleInnerImages = async (images) => {
    return await contentRequest.post('/articles/backend/image/create', { contents: images })
}

export const tagView = async () => {
    return await contentRequest.get('/articles/backend/tag')
}

export const tagUpdate = async () => {
    return await contentRequest.put('/tag/update', {})
}