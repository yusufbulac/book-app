"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Button from "@/components/ui/Button";
import EditBookModal from "@/components/EditBookModal";
import { BookResponse, UpdateBookRequest } from "@/types/book";
import { updateBook } from "@/services/bookService";
import { useToast } from "@/components/providers/ToastProvider";
import { ApiError } from "@/lib/fetcher";

export default function EditBookButton({ book }: { book: BookResponse }) {
    const [open, setOpen] = useState(false);
    const router = useRouter();
    const { show } = useToast();

    const handleUpdate = async (updated: UpdateBookRequest) => {
        try {
            await updateBook(updated);
            setOpen(false);
            show("Book updated.", "success");
            router.refresh();
        } catch (e) {
            const msg = e instanceof ApiError ? e.message : "Failed to update book.";
            show(msg, "error");
        }
    };

    return (
        <>
            <Button
                type="button"
                onClick={() => setOpen(true)}
                className="bg-yellow-500 hover:bg-yellow-600"
            >
                Edit
            </Button>
            {open && (
                <EditBookModal
                    book={book}
                    isOpen={open}
                    onClose={() => setOpen(false)}
                    onUpdate={handleUpdate}
                />
            )}
        </>
    );
}
