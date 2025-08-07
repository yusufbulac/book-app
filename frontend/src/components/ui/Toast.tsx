"use client";

import { useEffect } from "react";
import clsx from "clsx";

export type ToastKind = "success" | "error" | "info";

export interface ToastItem {
    id: string;
    message: string;
    kind: ToastKind;
}

export default function Toast({
                                  item,
                                  onDoneAction
                              }: {
    item: ToastItem;
    onDoneAction: (id: string) => void;
}) {
    useEffect(() => {
        const t = setTimeout(() => onDoneAction(item.id), 3000);
        return () => clearTimeout(t);
    }, [item.id, onDoneAction]);

    return (
        <div
            className={clsx(
                "rounded-lg shadow px-4 py-2 text-sm text-white",
                item.kind === "success" && "bg-green-600",
                item.kind === "error" && "bg-red-600",
                item.kind === "info" && "bg-zinc-800"
            )}
            role="status"
            aria-live="polite"
        >
            {item.message}
        </div>
    );
}
