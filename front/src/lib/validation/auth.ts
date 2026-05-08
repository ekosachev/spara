import * as z from "zod"

export const loginSchema = z.object({
    email: z.string().email({ message: "Некорректный формат email-адреса" }),
    password: z.string().min(6, { message: "Пароль должен быть не менее чем из 6 символов" })
})

export const registerSchema = z.object({
    email: z.string().email({ message: "Некорректный формат email-адреса" }),
    username: z.string().min(6, { message: "Имя пользователя должно быть длиннее 5 символов" }),
    password: z.string().min(6, { message: "Пароль должен быть не менее чем из 6 символов" })
})

export type LoginInput = z.infer<typeof loginSchema>
