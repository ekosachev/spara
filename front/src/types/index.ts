export interface ApiResponse<T> {
    success: boolean;
    data?: T;
    error?: string;
}

export interface Excercise {
    id: number;
    name: string;
    description: string;
}

export interface User {
    id: number;
    email: string;
    username: string;
}
