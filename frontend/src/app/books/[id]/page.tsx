import { notFound } from "next/navigation";
import EditBookButton from "@/components/actions/EditBookButton";
import Breadcrumbs from "@/components/Breadcrumbs";
import { fetchBookById } from "@/services/bookService";

export default async function BookDetailPage({
                                                 params,
                                             }: {
    params: Promise<{ id: string }>;
}) {
    const { id } = await params;
    const book = await fetchBookById(Number(id));
    if (!book) return notFound();

    const fmt = new Intl.DateTimeFormat("en-GB", {
        dateStyle: "medium",
        timeStyle: "short",
        timeZone: "UTC",
    });

    return (
        <main className="p-6 max-w-7xl mx-auto">
            <Breadcrumbs items={[{ label: "Home", href: "/" }, { label: "Books", href: "/" }, { label: book.title }]} />
            <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold text-zinc-900">Book Detail</h1>
                <EditBookButton book={book} />
            </div>
            <div className="bg-white rounded-2xl shadow p-6 space-y-3">
                <div><span className="font-medium text-zinc-700">Title:</span> {book.title}</div>
                <div><span className="font-medium text-zinc-700">Author:</span> {book.author}</div>
                <div><span className="font-medium text-zinc-700">Year:</span> {book.year}</div>
                <div><span className="font-medium text-zinc-700">Created At:</span> {fmt.format(new Date(book.created_at))}</div>
                <div><span className="font-medium text-zinc-700">Updated At:</span> {fmt.format(new Date(book.updated_at))}</div>
            </div>
        </main>
    );
}
