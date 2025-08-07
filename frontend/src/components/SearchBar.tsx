"use client";

import { useBooksUi } from "@/context/BooksUiContext";
import Input from "@/components/ui/Input";

export default function SearchBar() {
    const { query, setQuery } = useBooksUi();
    return (
        <div className="w-full max-w-sm">
            <Input
                placeholder="Search by title or author..."
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                aria-label="Search books"
            />
        </div>
    );
}
