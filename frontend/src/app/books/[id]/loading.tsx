export default function LoadingBookDetail() {
    return (
        <main className="p-6 max-w-7xl mx-auto">
            <div className="h-8 w-40 bg-gray-200 rounded mb-6 animate-pulse" />
            <div className="bg-white rounded-lg shadow p-6 space-y-3">
                <div className="h-4 w-64 bg-gray-200 rounded animate-pulse" />
                <div className="h-4 w-52 bg-gray-200 rounded animate-pulse" />
                <div className="h-4 w-36 bg-gray-200 rounded animate-pulse" />
                <div className="h-4 w-72 bg-gray-200 rounded animate-pulse" />
                <div className="h-4 w-80 bg-gray-200 rounded animate-pulse" />
            </div>
        </main>
    );
}
