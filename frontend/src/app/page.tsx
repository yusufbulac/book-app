import { fetchBooks } from "@/services/bookService";
import SearchBar from "@/components/SearchBar";
import SortSelect from "@/components/SortSelect";
import AddBookButton from "@/components/actions/AddBookButton";
import BooksView from "@/components/BooksView";

export default async function Home() {
    const books = await fetchBooks();

    return (
        <main className="p-6 max-w-7xl mx-auto">
            <div className="flex items-end justify-between gap-4">
                <div>
                    <h1 className="text-3xl font-bold tracking-tight text-zinc-900">Books</h1>
                    <p className="mt-1 text-sm text-zinc-600">
                        Manage your library — add, edit, and explore your collection.
                    </p>
                </div>
                <AddBookButton />
            </div>

            <div className="mt-4 flex items-center gap-3">
                <SearchBar />
                <SortSelect />
            </div>

            <BooksView books={books} />
        </main>
    );
}
