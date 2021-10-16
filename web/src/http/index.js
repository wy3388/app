import axios from "axios";
import {ElMessage} from "element-plus";

const baseUrl = '/api'

const ax = axios.create({
    baseURL: baseUrl
})

const request = async (url, method = 'get', data = {}) => {
    let [resp, err] = [null, null]
    if (method === 'get') {
        [resp, err] = await ax.get(url).then(resp => [resp, null]).catch(err => [null, err])
    } else {
        [resp, err] = await ax.post(url, data).then(resp => [resp, null]).catch(err => [null, err])
    }
    if (err != null) {
        ElMessage.error(err.message)
        return
    }
    if (resp.data.code === 2000) {
        return Promise.resolve(resp.data.data)
    } else {
        return Promise.reject(resp.data.message)
    }
}

const get = async (url, isErr = false) => {
    const [resp, err] = await request(url).then(resp => [resp, null]).catch(err => [null, err])
    if (err != null) {
        if (isErr) {
            return Promise.reject(err)
        } else {
            ElMessage.error(err)
            return
        }
    }
    return Promise.resolve(resp)
}

const post = async (url, params, isErr = false) => {
    const [resp, err] = await request(url, 'post', params).then(resp => [resp, null]).catch(err => [null, err])
    if (err != null) {
        if (isErr) {
            return Promise.reject(err)
        } else {
            ElMessage.error(err)
            return
        }
    }
    return Promise.resolve(resp)
}

const source = async () => {
    return await get('source', true)
}

const search = async (key, id) => {
    return await get(`search?key=${key}&id=${id}`)
}

const info = async (url, id) => {
    return await get(`info?url=${url}&id=${id}`)
}

const content = async (url, id) => {
    return await get(`body?url=${url}&id=${id}`)
}

const addBookSelf = async (params) => {
    return await post('addBookSelf', params, true)
}

const listBookSelf = async () => {
    return await get('listBookSelf')
}

export {
    source,
    search,
    info,
    content,
    addBookSelf,
    listBookSelf
}