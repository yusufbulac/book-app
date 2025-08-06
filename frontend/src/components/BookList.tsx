"use client";

import { useEffect, useState } from "react";
import { fetchBooks } from "@/services/bookService";
import {BookResponse} from "@/types/book";

export default function BookList() {
    const [books, setBooks] = useState<BookResponse[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchBooks()
            .then(setBooks)
            .finally(() => setLoading(false));
    }, []);

    if (loading) return <p className="text-gray-500">Loading...</p>;

    return (
        <div className="space-y-4">
            <h1 className="text-2xl font-bold">Books</h1>
            {books.length === 0 ? (
                <p>No books found.</p>
            ) : (
                <ul className="space-y-2">
                    {books.map((book) => (
                        <li key={book.id} className="border p-4 rounded shadow">
                            <p><strong>{book.title}</strong> by {book.author} ({book.year})</p>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
}
