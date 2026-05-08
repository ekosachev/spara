import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/api";
import { Excercise } from "@/types";

export const useExcercises = () => {
    return useQuery<Excercise[]>({
        queryKey: ['excercises'],
        queryFn: () => api.get('/excercise'),
    })
}
