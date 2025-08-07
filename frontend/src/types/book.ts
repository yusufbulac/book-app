export interface BookResponse {
    id: number;
    title: string;
    author: string;
    year: number;
    created_at: string;
    updated_at: string;
}

export interface CreateBookRequest {
    title: string;
    author: string;
    year: number;
}

export interface UpdateBookRequest {
    id: number;
    title: string;
    author: string;
    year: number;
}
