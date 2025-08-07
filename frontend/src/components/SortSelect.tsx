"use client";

import { useBooksUi } from "@/context/BooksUiContext";
import type { SortKey } from "@/context/BooksUiContext";

export default function SortSelect() {
    const { sort, setSort } = useBooksUi();
    return (
        <select
            className="border rounded-lg px-3 py-2"
            value={sort}
            onChange={(e: React.ChangeEvent<HTMLSelectElement>) =>
                setSort(e.target.value as SortKey)
            }
            aria-label="Sort books"
        >
            <option value="title-asc">Title ↑</option>
            <option value="title-desc">Title ↓</option>
            <option value="year-asc">Year ↑</option>
            <option value="year-desc">Year ↓</option>
        </select>
    );
}
