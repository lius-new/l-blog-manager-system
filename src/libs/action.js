import axios from 'axios'

const SERVER_URL = import.meta.env.VITE_API_SERVER_URI

const request = axios.create({
    withCredentials: true,
    baseURL: SERVER_URL,
})
request.interceptors.request.use(function (config) {
    return config
}, function (error) {
    return Promise.reject(error)
})
request.interceptors.response.use(function (response) {
    return response.data
}, function (error) {
    return Promise.reject(error)
})

export const login = async (username, password) => {
    return await request.post('/api/user/login', { username, password })
}
export const auth = async () => {
    return await request.post('/api/user/auth',)
}
export const articlesViews = async (page_size, page_num) => {
    return await request.post('/api/articles/views', { page_size, page_num })
}
export const articlesView = async (id) => {
    return await request.post('/api/articles/view', { id })
}

export const articleSave = async (title, content, tags, covers) => {
    return await request.postForm('/api/articles/create', {
        title, content, tags, covers
    }, {
        formSerializer: {
            indexes: null
        }
    })
}

export const articleModifyStatus = async (id, status) => {
    return await request.post('/api/articles/modify-status', { id, status })
}

export const articleModify = async (id, title, content, tags, covers, status) => {
    return await request.postForm('/api/articles/modify', { id, title, content, tags, covers, status }, {
        formSerializer: {
            indexes: null
        }
    })
}

export const postDisable = async () => {
    return await request.post('/post/disable', {})
}

export const tagView = async () => {
    return await request.get('/api/tag/view')
}

export const tagUpdate = async () => {
    return await request.put('/tag/update', {})
}