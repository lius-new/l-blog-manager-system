import axios from 'axios'

const SERVER_URL = "http://localhost:8080"

const request = axios.create({
    baseURL: SERVER_URL,
})
request.interceptors.request.use(function (config) {
    return config
}, function (error) {
    return Promise.reject(error)
})

export const login = async (username, password) => {
    return await request.post('/api/user/login', { username, password })
}

export const postSave = async () => {
    return await request.post('/post/save', {})
}

export const postUpdate = async () => {
    return await request.put('/post/update', {})
}

export const postDisable = async () => {
    return await request.post('/post/disable', {})
}

export const tagSave = async () => {
    return await request.post('/tag/save', {})
}

export const tagUpdate = async () => {
    return await request.put('/tag/update', {})
}