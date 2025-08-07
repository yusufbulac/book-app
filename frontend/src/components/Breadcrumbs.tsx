import Link from "next/link";

type Crumb = { label: string; href?: string };

export default function Breadcrumbs({ items }: { items: Crumb[] }) {
    return (
        <nav aria-label="Breadcrumb" className="mb-4">
            <ol className="flex flex-wrap items-center gap-1 text-sm text-zinc-600">
                {items.map((item, idx) => {
                    const isLast = idx === items.length - 1;
                    return (
                        <li key={`${item.label}-${idx}`} className="flex items-center">
                            {item.href && !isLast ? (
                                <Link
                                    href={item.href}
                                    className="hover:text-[var(--brand-primary)]"
                                >
                                    {item.label}
                                </Link>
                            ) : (
                                <span className="font-medium text-zinc-800">{item.label}</span>
                            )}
                            {!isLast && <span className="mx-2 text-zinc-400">/</span>}
                        </li>
                    );
                })}
            </ol>
        </nav>
    );
}
