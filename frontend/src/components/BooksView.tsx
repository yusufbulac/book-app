"use client";

import { useMemo } from "react";
import { useBooksUi } from "@/context/BooksUiContext";
import { BookResponse } from "@/types/book";
import BookCard from "@/components/BookCard";
import AddBookButton from "@/components/actions/AddBookButton";

export default function BooksView({ books }: { books: BookResponse[] }) {
    const { query, sort } = useBooksUi();

    const filtered = useMemo(() => {
        const q = query.trim().toLowerCase();
        let arr = !q
            ? books
            : books.filter(
                (b) =>
                    b.title.toLowerCase().includes(q) ||
                    b.author.toLowerCase().includes(q)
            );

        const cmp = new Intl.Collator("en").compare;
        switch (sort) {
            case "title-asc":
                arr = [...arr].sort((a, b) => cmp(a.title, b.title));
                break;
            case "title-desc":
                arr = [...arr].sort((a, b) => cmp(b.title, a.title));
                break;
            case "year-asc":
                arr = [...arr].sort((a, b) => a.year - b.year);
                break;
            case "year-desc":
                arr = [...arr].sort((a, b) => b.year - a.year);
                break;
        }
        return arr;
    }, [books, query, sort]);

    if (filtered.length === 0) {
        return (
            <div className="text-center py-20">
                <h3 className="text-lg font-semibold text-gray-700">No books found</h3>
                <p className="text-sm text-gray-500 mb-4">Try a different search.</p>
                <AddBookButton />
            </div>
        );
    }

    return (
        <ul className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mt-6">
            {filtered.map((book) => (
                <li key={book.id}>
                    <BookCard book={book} />
                </li>
            ))}
        </ul>
    );
}
