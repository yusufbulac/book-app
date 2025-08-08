import { BookResponse, CreateBookRequest, UpdateBookRequest } from "@/types/book";
import { apiBaseUrl, request } from "@/lib/fetcher";

const BASE = () => `${apiBaseUrl()}/api/v1/books`;

export async function fetchBooks(): Promise<BookResponse[]> {
    const data = await request<{ data: BookResponse[] }>(`${BASE()}`, { cache: "no-store" });
    return data.data ?? [];
}

export async function createBook(book: CreateBookRequest): Promise<BookResponse> {
    const data = await request<{ data: BookResponse }>(`${BASE()}`, {
        method: "POST",
        json: book
    });
    return data.data;
}

export async function fetchBookById(id: number): Promise<BookResponse | null> {
    try {
        const data = await request<{ data: BookResponse }>(`${BASE()}/${id}`, { cache: "no-store" });
        return data.data ?? null;
    } catch {
        return null;
    }
}

export async function updateBook(book: UpdateBookRequest): Promise<BookResponse> {
    const data = await request<{ data: BookResponse }>(`${BASE()}/${book.id}`, {
        method: "PUT",
        json: book,
    });
    return data.data;
}

export async function deleteBook(id: number): Promise<void> {
    await request(`${BASE()}/${id}`, { method: "DELETE" });
}
