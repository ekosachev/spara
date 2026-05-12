import { api } from "@/lib/api"
import { LoginInput, RegisterInput } from "@/lib/validation/auth"
import { useMutation } from "@tanstack/react-query"
import { useRouter } from "next/navigation"

export const useAuth = () => {
    const router = useRouter()

    const loginMutation = useMutation({
        mutationFn: (data: LoginInput) => api.post("/auth/login", data),
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        onSuccess: (data: any) => {
            localStorage.setItem("token", data.token)
            router.push("/dashboard")
        },
    })

    const registerMutation = useMutation({
        mutationFn: (data: RegisterInput) => api.post("/user", data),
        onSuccess: () => {
            router.push("/auth/login")
        }
    })

    return { loginMutation, registerMutation }
}
