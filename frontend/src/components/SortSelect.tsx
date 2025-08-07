"use client";

import { useBooksUi } from "@/context/BooksUiContext";

export default function SortSelect() {
    const { sort, setSort } = useBooksUi();
    return (
        <select
            className="border rounded-lg px-3 py-2"
            value={sort}
            onChange={(e) => setSort(e.target.value as any)}
            aria-label="Sort books"
        >
            <option value="title-asc">Title ↑</option>
            <option value="title-desc">Title ↓</option>
            <option value="year-asc">Year ↑</option>
            <option value="year-desc">Year ↓</option>
        </select>
    );
}
