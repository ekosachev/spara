import axios from "axios";
import { ApiResponse } from "@/types";

export const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL
});

api.interceptors.request.use((config) => {
    const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;

    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }

    return config
})

api.interceptors.response.use((response) => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const res = response.data as ApiResponse<any>;

    return res.data
},
    (error) => {
        return Promise.reject(error)
    })
