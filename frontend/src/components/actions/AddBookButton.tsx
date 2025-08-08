"use client";

import { useState, useTransition, useCallback } from "react";
import { useRouter } from "next/navigation";
import AddBookModal from "@/components/AddBookModal";
import Button from "@/components/ui/Button";
import { CreateBookRequest } from "@/types/book";
import { createBook } from "@/services/bookService";
import { useToast } from "@/components/providers/ToastProvider";
import { ApiError } from "@/lib/fetcher";

export default function AddBookButton() {
    const [open, setOpen] = useState(false);
    const [submitting, setSubmitting] = useState(false);
    const [isPending, startTransition] = useTransition();
    const router = useRouter();
    const { show } = useToast();

    const handleAdd = useCallback(
        async (b: CreateBookRequest) => {
            if (submitting) return;
            setSubmitting(true);
            try {
                await createBook(b);
                show("Book created successfully.", "success");
                startTransition(() => router.refresh());
            } catch (e) {
                const msg =
                    e instanceof ApiError && e.message
                        ? e.message
                        : "Failed to create book.";
                show(msg, "error");
                throw e;
            } finally {
                setSubmitting(false);
            }
        },
        [router, show, submitting]
    );

    return (
        <>
            <Button
                type="button"
                onClick={() => setOpen(true)}
                aria-haspopup="dialog"
                aria-expanded={open}
                disabled={isPending || submitting}
            >
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
