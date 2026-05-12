import { ReactNode } from "react";

export default function DashboardLayout({ children }: { children: ReactNode }) {
    return (
        <div className="flex min-h-screen">
            <aside className="hidden md:flex w-64 flex-col border-r bg-card">
                <span>Sidebar</span>
            </aside>

            <main className="flex-1 bg-muted/30">
                <header className="h-16 border-b bg-card flex items-center px-8 md:hidden">
                    <span className="font-bold">Spara</span>
                </header>

                <div className="p-4 md:p-8">
                    {children}
                </div>
            </main>
        </div>
    )
}
