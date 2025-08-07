import Link from "next/link";
import Image from "next/image";
import { BookResponse } from "@/types/book";
import DeleteBookButton from "@/components/actions/DeleteBookButton";

export default function BookCard({ book }: { book: BookResponse }) {
    return (
        <div className="bg-white dark:bg-zinc-900 rounded-2xl shadow hover:shadow-xl overflow-hidden transition transform hover:-translate-y-0.5 focus-within:ring-2 focus-within:ring-[var(--brand-accent)]">
            <div className="relative w-full h-44">
                <Image
                    src={`https://picsum.photos/seed/${book.id}/600/400`}
                    alt={book.title}
                    fill
                    className="object-cover"
                    sizes="(max-width: 768px) 100vw, 33vw"
                    priority={false}
                />
            </div>

            <div className="p-4">
                <div className="text-xs uppercase text-[var(--brand-primary)] font-semibold tracking-wide">
                    Book
                </div>

                <h3 className="mt-1 text-lg font-semibold text-zinc-900">
                    <Link href={`/books/${book.id}`} className="hover:underline">
                        {book.title}
                    </Link>
                </h3>

                <p className="text-sm text-zinc-600">
                    by {book.author} ({book.year})
                </p>

                <div className="mt-3">
                    <DeleteBookButton id={book.id} title={book.title} />
                </div>
            </div>
        </div>
    );
}
