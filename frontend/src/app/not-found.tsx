import Link from "next/link";
import { AlertTriangle } from "lucide-react";

export default function NotFound() {
    return (
        <div className="flex flex-col items-center justify-center h-[70vh] text-center space-y-4">
            <AlertTriangle className="w-12 h-12 text-yellow-500" />
            <h1 className="text-4xl font-bold text-zinc-800 dark:text-zinc-100">404 - Page Not Found</h1>
            <p className="text-zinc-600 dark:text-zinc-400">
                The page you’re looking for doesn’t exist or has been moved.
            </p>
            <Link href="/" className="text-sm text-blue-600 hover:underline">
                ← Go back to Home
            </Link>
        </div>
    );
}
