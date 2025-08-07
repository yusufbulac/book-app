"use client";

import React, { useState } from "react";
import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Modal from "@/components/ui/Modal";

interface AddBookModalProps {
    onClose: () => void;
    onAdd: (book: { title: string; author: string; year: number }) => void;
    isOpen?: boolean;
}

export default function AddBookModal({ isOpen = false, onClose, onAdd }: AddBookModalProps) {
    const [title, setTitle] = useState("");
    const [author, setAuthor] = useState("");
    const [year, setYear] = useState(new Date().getFullYear());
    const [errors, setErrors] = useState<{ title?: string; author?: string; year?: string }>({});

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        const validationErrors: typeof errors = {};
        if (!title.trim()) validationErrors.title = "Title is required";
        if (!author.trim()) validationErrors.author = "Author is required";
        if (!year || year < 1000 || year > 9999) validationErrors.year = "Enter a valid year";

        if (Object.keys(validationErrors).length > 0) {
            setErrors(validationErrors);
            return;
        }

        onAdd({ title, author, year });
        setTitle("");
        setAuthor("");
        setYear(new Date().getFullYear());
        setErrors({});
        onClose();
    };

    return (
        <Modal isOpen={isOpen} onClose={onClose} title="Add Book">
            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <Input
                        placeholder="Title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        className={errors.title ? "border-red-500" : ""}
                        aria-invalid={!!errors.title}
                        aria-describedby={errors.title ? "add-title-error" : undefined}
                    />
                    {errors.title && (
                        <p id="add-title-error" className="text-sm text-red-600 mt-1">
                            {errors.title}
                        </p>
                    )}
                </div>

                <div>
                    <Input
                        placeholder="Author"
                        value={author}
                        onChange={(e) => setAuthor(e.target.value)}
                        className={errors.author ? "border-red-500" : ""}
                        aria-invalid={!!errors.author}
                        aria-describedby={errors.author ? "add-author-error" : undefined}
                    />
                    {errors.author && (
                        <p id="add-author-error" className="text-sm text-red-600 mt-1">
                            {errors.author}
                        </p>
                    )}
                </div>

                <div>
                    <Input
                        type="number"
                        placeholder="Year"
                        value={year}
                        onChange={(e) => setYear(Number(e.target.value))}
                        className={errors.year ? "border-red-500" : ""}
                        aria-invalid={!!errors.year}
                        aria-describedby={errors.year ? "add-year-error" : undefined}
                    />
                    {errors.year && (
                        <p id="add-year-error" className="text-sm text-red-600 mt-1">
                            {errors.year}
                        </p>
                    )}
                </div>

                <div className="flex justify-end gap-2 pt-4">
                    <Button type="button" variant="secondary" onClick={onClose}>
                        Cancel
                    </Button>
                    <Button type="submit" variant="primary">
                        Add
                    </Button>
                </div>
            </form>
        </Modal>
    );
}
