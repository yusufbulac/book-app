import AddBookButton from "@/components/actions/AddBookButton";
import BookCard from "@/components/BookCard";
import { fetchBooks } from "@/services/bookService";

export default async function Home() {
    const books = await fetchBooks();
    return (
        <main className="p-6 max-w-7xl mx-auto">
            <div className="flex items-end justify-between gap-4">
                <div>
                    <h1 className="text-3xl font-bold tracking-tight text-zinc-900">Books</h1>
                    <p className="mt-1 text-sm text-zinc-600">Manage your library — add, edit, and explore your collection.</p>
                </div>
                <AddBookButton />
            </div>

            {books.length === 0 ? (
                <div className="text-center py-20">
                    <h3 className="text-lg font-semibold text-gray-700">No books found</h3>
                    <p className="text-sm text-gray-500 mb-4">Start your collection by adding a new book now.</p>
                    <AddBookButton />
                </div>
            ) : (
                <ul className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mt-6">
                    {books.map((book) => (
                        <li key={book.id}>
                            <BookCard book={book} />
                        </li>
                    ))}
                </ul>
            )}
        </main>
    );
}
