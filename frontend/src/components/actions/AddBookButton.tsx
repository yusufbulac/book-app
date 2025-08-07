"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import AddBookModal from "@/components/AddBookModal";
import Button from "@/components/ui/Button";
import { CreateBookRequest } from "@/types/book";
import { createBook } from "@/services/bookService";
import { useToast } from "@/components/providers/ToastProvider";
import { ApiError } from "@/lib/fetcher";

export default function AddBookButton() {
    const [open, setOpen] = useState(false);
    const router = useRouter();
    const { show } = useToast();

    const handleAdd = async (b: CreateBookRequest) => {
        try {
            await createBook(b);
            setOpen(false);
            show("Book created successfully.", "success");
            router.refresh();
        } catch (e) {
            const msg = e instanceof ApiError ? e.message : "Failed to create book.";
            show(msg, "error");
        }
    };

    return (
        <>
            <Button type="button" onClick={() => setOpen(true)}>
                Add Book
            </Button>
            {open && (
                <AddBookModal
                    isOpen={open}
                    onClose={() => setOpen(false)}
                    onAdd={handleAdd}
                />
            )}
        </>
    );
}
