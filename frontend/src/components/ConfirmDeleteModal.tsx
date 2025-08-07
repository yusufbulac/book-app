"use client";

import Modal from "@/components/ui/Modal";

interface ConfirmDeleteModalProps {
    isOpen: boolean;
    onClose: () => void;
    onConfirm: () => void;
    bookTitle?: string;
}

export default function ConfirmDeleteModal({
                                               isOpen,
                                               onClose,
                                               onConfirm,
                                               bookTitle,
                                           }: ConfirmDeleteModalProps) {
    return (
        <Modal isOpen={isOpen} onClose={onClose} title="Delete Book">
            <p className="text-sm text-gray-700 mb-4">
                Are you sure you want to delete <strong>{bookTitle}</strong>? This action cannot be undone.
            </p>
            <div className="flex justify-end space-x-2">
                <button
                    onClick={onClose}
                    className="px-4 py-2 text-sm bg-gray-100 hover:bg-gray-200 text-gray-700 rounded-md transition"
                >
                    Cancel
                </button>
                <button
                    onClick={() => {
                        onConfirm();
                        onClose();
                    }}
                    className="px-4 py-2 text-sm bg-red-600 hover:bg-red-700 text-white rounded-md transition"
                >
                    Delete
                </button>
            </div>
        </Modal>
    );
}
