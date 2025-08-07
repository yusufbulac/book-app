"use client";

import React, { useState, useEffect } from "react";
import Button from "@/components/ui/Button";
import Input from "@/components/ui/Input";
import Modal from "@/components/ui/Modal";
import { BookResponse, UpdateBookRequest } from "@/types/book";

interface EditBookModalProps {
    book: BookResponse;
    isOpen: boolean;
    onClose: () => void;
    onUpdate: (updatedBook: UpdateBookRequest) => void;
}

export default function EditBookModal({
                                          book,
                                          isOpen,
                                          onClose,
                                          onUpdate,
                                      }: EditBookModalProps) {
    const [title, setTitle] = useState(book.title);
    const [author, setAuthor] = useState(book.author);
    const [year, setYear] = useState(book.year);
    const [errors, setErrors] = useState<{ title?: string; author?: string; year?: string }>({});

    useEffect(() => {
        setTitle(book.title);
        setAuthor(book.author);
        setYear(book.year);
        setErrors({});
    }, [book]);

    const validate = () => {
        const v: typeof errors = {};
        if (!title.trim()) v.title = "Title is required";
        if (!author.trim()) v.author = "Author is required";
        if (!year || year < 1000 || year > 9999) v.year = "Enter a valid year";
        setErrors(v);
        return Object.keys(v).length === 0;
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (!validate()) return;
        onUpdate({ id: book.id, title, author, year });
        onClose();
    };

    return (
        <Modal isOpen={isOpen} onClose={onClose} title="Edit Book">
            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <Input
                        placeholder="Title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        className={errors.title ? "border-red-500" : ""}
                        aria-invalid={!!errors.title}
                        aria-describedby={errors.title ? "edit-title-error" : undefined}
                    />
                    {errors.title && (
                        <p id="edit-title-error" className="text-sm text-red-600 mt-1">
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
                        aria-describedby={errors.author ? "edit-author-error" : undefined}
                    />
                    {errors.author && (
                        <p id="edit-author-error" className="text-sm text-red-600 mt-1">
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
                        aria-describedby={errors.year ? "edit-year-error" : undefined}
                    />
                    {errors.year && (
                        <p id="edit-year-error" className="text-sm text-red-600 mt-1">
                            {errors.year}
                        </p>
                    )}
                </div>

                <div className="flex justify-end gap-2 pt-4">
                    <Button type="button" variant="secondary" onClick={onClose}>
                        Cancel
                    </Button>
                    <Button type="submit" variant="primary">
                        Update
                    </Button>
                </div>
            </form>
        </Modal>
    );
}
