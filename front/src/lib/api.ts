import axios from "axios";
import { ApiResponse } from "@/types";
import { error } from "console";

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

    if (!res.success) {
        return Promise.reject(new Error(res.error || 'Something went wrong'))
    }

    return res.data
},
    (error) => {
        const message = error.response?.data?.error || 'Server error'

        return Promise.reject(new Error(message))
    })
