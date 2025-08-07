"use client";

import { useEffect } from "react";
import { Bug } from "lucide-react";

export default function GlobalError({ error, reset }: { error: Error; reset: () => void }) {
    useEffect(() => {
        console.error("Unexpected error:", error);
    }, [error]);

    return (
        <div className="flex flex-col items-center justify-center h-[70vh] text-center space-y-4">
            <Bug className="w-12 h-12 text-red-600" />
            <h1 className="text-3xl font-bold text-red-600">Something went wrong</h1>
            <p className="text-zinc-600 dark:text-zinc-400">
                We encountered an unexpected error. Please try again later.
            </p>
            <button
                onClick={reset}
                className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700 text-sm"
            >
                Try Again
            </button>
        </div>
    );
}
