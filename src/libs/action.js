import axios from 'axios'

const SERVER_URL = "http://localhost:8080"

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
export const articlesView = async (page_size, page_num) => {
    return await request.post('/api/articles/view', { page_size, page_num })
}

export const articleSave = async (title, content, tags, covers, status) => {
    return await request.post('/api/articles/create', { title, tags, content, covers, status })
}

export const postUpdate = async () => {
    return await request.put('/post/update', {})
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