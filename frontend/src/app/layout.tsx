import "./globals.css";
import Header from "@/components/layout/Header";
import { ToastProvider } from "@/components/providers/ToastProvider";
import { BooksUiProvider } from "@/context/BooksUiContext";

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en">
        <body className="bg-[var(--background)] text-[var(--foreground)] font-sans antialiased">
        <Header />
        <ToastProvider>
            <BooksUiProvider>
                {children}
            </BooksUiProvider>
        </ToastProvider>
        <div id="toast-root" />
        </body>
        </html>
    );
}
