"use client";

import { createContext, useContext, useState, useMemo } from "react";

type SortKey = "title-asc" | "title-desc" | "year-asc" | "year-desc";

type BooksUiState = {
    query: string;
    sort: SortKey;
    setQuery: (q: string) => void;
    setSort: (s: SortKey) => void;
};

const BooksUiContext = createContext<BooksUiState | null>(null);

export function BooksUiProvider({ children }: { children: React.ReactNode }) {
    const [query, setQuery] = useState("");
    const [sort, setSort] = useState<SortKey>("title-asc");

    const value = useMemo(
        () => ({ query, sort, setQuery, setSort }),
        [query, sort]
    );

    return <BooksUiContext.Provider value={value}>{children}</BooksUiContext.Provider>;
}

export function useBooksUi() {
    const ctx = useContext(BooksUiContext);
    if (!ctx) throw new Error("useBooksUi must be used within BooksUiProvider");
    return ctx;
}
