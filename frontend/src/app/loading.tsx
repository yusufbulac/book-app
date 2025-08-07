export default function LoadingHome() {
    return (
        <main className="p-6 max-w-7xl mx-auto">
            <div className="flex justify-between items-center">
                <div className="h-8 w-40 bg-gray-200 rounded animate-pulse" />
                <div className="h-9 w-28 bg-gray-200 rounded animate-pulse" />
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mt-6">
                {Array.from({ length: 6 }).map((_, i) => (
                    <div key={i} className="bg-white rounded-lg shadow p-4 h-56 animate-pulse" />
                ))}
            </div>
        </main>
    );
}
