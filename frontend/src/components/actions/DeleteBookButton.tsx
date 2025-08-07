"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import ConfirmDeleteModal from "@/components/ConfirmDeleteModal";
import Button from "@/components/ui/Button";
import { deleteBook } from "@/services/bookService";
import { useToast } from "@/components/providers/ToastProvider";
import { ApiError } from "@/lib/fetcher";

export default function DeleteBookButton({
                                             id,
                                             title,
                                         }: {
    id: number;
    title: string;
}) {
    const [open, setOpen] = useState(false);
    const router = useRouter();
    const { show } = useToast();

    const confirm = async () => {
        try {
            await deleteBook(id);
            setOpen(false);
            show("Book deleted.", "success");
            router.refresh();
        } catch (e) {
            const msg = e instanceof ApiError ? e.message : "Failed to delete book.";
            show(msg, "error");
        }
    };

    return (
        <>
            <Button
                type="button"
                variant="destructive"
                size="sm"
                onClick={() => setOpen(true)}
            >
                Delete
            </Button>
            <ConfirmDeleteModal
                isOpen={open}
                onClose={() => setOpen(false)}
                onConfirm={confirm}
                bookTitle={title}
            />
        </>
    );
}
