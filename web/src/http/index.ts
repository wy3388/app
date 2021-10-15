import axios from "axios";
import {ElMessage} from "element-plus";
import {Search, Source} from "@/model";

const baseUrl = "/api"

const ax = axios.create({
    baseURL: baseUrl
})

interface ApiResult<T> {
    code: number
    message: string
    data: T
}

const get = async <T>(url: string): Promise<T | null> => {
    try {
        const response = await ax.get<ApiResult<T>>(url)
        const resp = response.data
        if (resp.code === 2000) {
            return Promise.resolve(resp.data)
        }
        ElMessage.error(resp.message)
    } catch (e) {
        ElMessage.error(e.message)
    }
    return null
}

// const post = async <T>(url: string, params?: any | null): Promise<T | null> => {
//     try {
//         const response = await ax.post<ApiResult<T>>(url, params)
//         const resp = response.data
//         if (resp.code === 2000) {
//             return Promise.resolve(resp.data)
//         }
//         ElMessage.error(resp.message)
//         return null
//     } catch (e) {
//         ElMessage.error(e.message)
//         return null
//     }
// }

const source = async (): Promise<Source[]> => {
    const resp = await get<Array<Source>>('source')
    return resp == null ? [] : resp
}

const search = async (key: string, id: number): Promise<Search[]> => {
    const resp = await get<Array<Search>>(`search?key=${key}&id=${id}`)
    return resp == null ? [] : resp
}

export {
    source,
    search
}
