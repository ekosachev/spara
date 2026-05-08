'use client';

import { useExcercises } from "@/hooks/useExcercises";

export default function Home() {
    const { data: excercises, isLoading, error } = useExcercises();

    if (isLoading) return (<div>Грузим железки...</div>);

    if (error) return (<div>Бэк упал или CORS отвалился</div>);

    return (
        <main>
            <h1>Упражнения</h1>
            {excercises?.map((ex) => (
                <div key={ex.id}>
                    <h2>{ex.name}</h2>
                </div>
            ))}
        </main>
    )
}
