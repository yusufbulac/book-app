import Link from "next/link";

export default function Header() {
    return (
        <header className="sticky top-0 z-40 bg-white/90 backdrop-blur shadow-sm">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex h-14 items-center justify-between">
                    <Link
                        href="/"
                        className="text-lg font-semibold text-zinc-900"
                        aria-label="Go to homepage"
                    >
                        byFood • Library
                    </Link>

                    <nav className="hidden sm:flex items-center gap-6 text-sm">
                        <Link href="/" className="hover:text-[var(--brand-primary)]">
                            Books
                        </Link>
                    </nav>
                </div>
            </div>
        </header>
    );
}
