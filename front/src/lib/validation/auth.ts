import * as z from "zod"

export const loginSchema = z.object({
    email: z.string().email({ message: "Incorrect email" }),
    password: z.string().min(6, { message: "Password too short. Use at least 6 characters" })
})

export const registerSchema = z.object({
    email: z.string().email({ message: "Incorrect email" }),
    username: z.string().min(6, { message: "Username too short. Use at least 6 letters" }),
    password: z.string().min(6, { message: "Password too short. Use at least 6 characters" })
})

export type LoginInput = z.infer<typeof loginSchema>
export type RegisterInput = z.infer<typeof registerSchema>
