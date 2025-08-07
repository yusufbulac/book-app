"use client";

import { createContext, useCallback, useContext, useEffect, useMemo, useState } from "react";
import { createPortal } from "react-dom";
import Toast, { ToastItem, ToastKind } from "@/components/ui/Toast";

type ToastContextValue = { show: (message: string, kind?: ToastKind) => void };
const ToastContext = createContext<ToastContextValue | null>(null);

export function ToastProvider({ children }: { children: React.ReactNode }) {
    const [items, setItems] = useState<ToastItem[]>([]);
    const [mounted, setMounted] = useState(false);
    const [container, setContainer] = useState<Element | null>(null);

    useEffect(() => {
        setMounted(true);
        setContainer(document.getElementById("toast-root"));
    }, []);

    const dismiss = useCallback((id: string) => {
        setItems((prev) => prev.filter((t) => t.id !== id));
    }, []);

    const show = useCallback((message: string, kind: ToastKind = "info") => {
        const id = Math.random().toString(36).slice(2);
        setItems((prev) => [...prev, { id, message, kind }]);
    }, []);

    const value = useMemo(() => ({ show }), [show]);

    return (
        <ToastContext.Provider value={value}>
            {children}
            {mounted && container
                ? createPortal(
                    <div className="fixed top-4 right-4 z-[100] flex flex-col gap-2">
                        {items.map((i) => (
                            <Toast key={i.id} item={i} onDoneAction={dismiss} />
                        ))}
                    </div>,
                    container
                )
                : null}
        </ToastContext.Provider>
    );
}

export function useToast() {
    const ctx = useContext(ToastContext);
    if (!ctx) throw new Error("useToast must be used within ToastProvider");
    return ctx;
}
