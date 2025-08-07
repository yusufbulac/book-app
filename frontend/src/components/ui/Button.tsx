import clsx from "clsx";
import React from "react";

type Variant = "primary" | "secondary" | "outline" | "destructive";
type Size = "sm" | "md" | "lg";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    variant?: Variant;
    size?: Size;
}

const variantClasses: Record<Variant, string> = {
    primary:
        "bg-[var(--brand-primary)] text-white hover:bg-[var(--brand-primary-hover)]",
    secondary:
        "bg-gray-100 text-gray-900 hover:bg-gray-200",
    outline:
        "border border-[var(--brand-primary)] text-[var(--brand-primary)] hover:bg-[color-mix(in_hsl,var(--brand-accent)_12%,transparent)]",
    destructive:
        "bg-red-600 text-white hover:bg-red-700",
};

const sizeClasses: Record<Size, string> = {
    sm: "px-3 py-1.5 text-sm rounded-md",
    md: "px-4 py-2 text-base rounded-lg",
    lg: "px-5 py-3 text-lg rounded-xl",
};

export default function Button({
                                   children,
                                   className,
                                   variant = "primary",
                                   size = "md",
                                   ...props
                               }: ButtonProps) {
    return (
        <button
            className={clsx(
                "font-medium transition-colors duration-200",
                variantClasses[variant],
                sizeClasses[size],
                className
            )}
            {...props}
        >
            {children}
        </button>
    );
}
