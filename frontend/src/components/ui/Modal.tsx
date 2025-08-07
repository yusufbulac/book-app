"use client";

import React, { ReactNode, useEffect, useId, useRef } from "react";
import clsx from "clsx";

interface ModalProps {
    isOpen: boolean;
    onClose: () => void;
    title: string;
    children: ReactNode;
    className?: string;
    closeOnOverlayClick?: boolean;
}

function getFocusable(container: HTMLElement | null): HTMLElement[] {
    if (!container) return [];
    const selectors =
        'a[href], button:not([disabled]), textarea, input, select, [tabindex]:not([tabindex="-1"])';
    const nodes = Array.from(container.querySelectorAll<HTMLElement>(selectors));
    return nodes.filter(
        (el) => !el.hasAttribute("disabled") && el.getAttribute("aria-hidden") !== "true",
    );
}

export default function Modal({
                                  isOpen,
                                  onClose,
                                  title,
                                  children,
                                  className,
                                  closeOnOverlayClick = true,
                              }: ModalProps) {
    const titleId = useId();
    const contentRef = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (!isOpen) return;
        // Body scroll lock
        document.body.classList.add("overflow-hidden");
        return () => {
            document.body.classList.remove("overflow-hidden");
        };
    }, [isOpen]);

    useEffect(() => {
        if (!isOpen) return;

        const handleKeyDown = (e: KeyboardEvent) => {
            if (e.key === "Escape") {
                e.preventDefault();
                onClose();
            }
            if (e.key === "Tab") {
                const focusables = getFocusable(contentRef.current);
                if (focusables.length === 0) return;

                const current = document.activeElement as HTMLElement | null;
                const index = focusables.indexOf(current || focusables[0]);

                if (e.shiftKey) {
                    // reverse tab
                    if (index <= 0) {
                        e.preventDefault();
                        focusables[focusables.length - 1].focus();
                    }
                } else {
                    // forward tab
                    if (index === focusables.length - 1) {
                        e.preventDefault();
                        focusables[0].focus();
                    }
                }
            }
        };

        document.addEventListener("keydown", handleKeyDown);
        return () => document.removeEventListener("keydown", handleKeyDown);
    }, [isOpen, onClose]);

    useEffect(() => {
        if (!isOpen) return;
        const focusables = getFocusable(contentRef.current);
        (focusables[0] || contentRef.current)?.focus();
    }, [isOpen]);

    if (!isOpen) return null;

    return (
        <div
            className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 px-4"
            onMouseDown={(e) => {
                if (!closeOnOverlayClick) return;
                if (e.target === e.currentTarget) onClose();
            }}
        >
            <div
                ref={contentRef}
                role="dialog"
                aria-modal="true"
                aria-labelledby={titleId}
                className={clsx("bg-white w-full max-w-md rounded-2xl shadow-xl p-6 outline-none", className)}
                tabIndex={-1}
                onMouseDown={(e) => e.stopPropagation()}
            >
                <h2 id={titleId} className="text-2xl font-semibold mb-6 text-gray-900">
                    {title}
                </h2>
                {children}
            </div>
        </div>
    );
}
